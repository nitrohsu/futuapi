package protocol

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"log"
	"nitrohsu.com/futu/api/common"
	"nitrohsu.com/futu/api/initconnect"
	"nitrohsu.com/futu/api/keepalive"
	"nitrohsu.com/futu/api/qotgetsubinfo"
	"nitrohsu.com/futu/api/trdunlocktrade"
	"time"
)

type Handler struct {
	wakeup chan bool
	quit   chan bool
	Resp   chan *Message
}

func (handler *Handler) Init(writer chan *Message, reader chan *Message, pwd string) error {
	if writer == nil || reader == nil {
		return errors.New("not init netManager")
	}
	handler.wakeup = make(chan bool, 1)
	handler.quit = make(chan bool, 1)
	handler.Resp = make(chan *Message, 10)
	// init connect
	initConnectRequest(writer, "123", 300, true, 0, common.PacketEncAlgo_PacketEncAlgo_None)
	// response
	go handler.response(writer, reader, pwd)
	return nil
}
func (handler Handler) Close() {
	handler.quit <- true
	close(handler.Resp)
	time.Sleep(10 * time.Second)
}

func initConnectRequest(writer chan *Message, clientId string, clientVersion int32, notify bool, format int32, encodeAlgo common.PacketEncAlgo) {
	ftProto := Message{
		ProtoID: uint32(P_InitConnect),
		Body: &initconnect.Request{
			C2S: &initconnect.C2S{
				ClientID:      proto.String(clientId),
				ClientVer:     proto.Int32(clientVersion),
				RecvNotify:    proto.Bool(notify),
				PushProtoFmt:  proto.Int32(format),
				PacketEncAlgo: proto.Int32(int32(encodeAlgo)),
			},
		},
	}
	writer <- &ftProto
}

func (handler *Handler) keepAliveFunction(writer chan *Message, intervalSeconds int32) {
	go func() {
		for {
			select {
			case <-time.After(time.Duration(intervalSeconds-2) * time.Second):
				go func() {
					handler.wakeup <- true
				}()
			case <-handler.quit:
				log.Printf("keepAlive quit")
			case <-handler.wakeup:
				writer <- &Message{
					ProtoID: uint32(P_KeepAlive),
					Body: &keepalive.Request{
						C2S: &keepalive.C2S{
							Time: proto.Int64(time.Now().Unix()),
						},
					},
				}
			}
		}
	}()
}

func (handler *Handler) response(writer chan *Message, reader chan *Message, pwd string) {
	for {
		select {
		case msg := <-reader:
			if msg == nil {
				log.Printf("response is nil")
				return
			}

			switch msg.ProtoID {
			case P_InitConnect:
				{
					body, ok := msg.Body.(*initconnect.Response)
					if !ok || *body.ErrCode != 0 {
						log.Printf("init connect resp failed, err=%s", *body.RetMsg)
						return
					}
					// request basic info
					handler.keepAliveFunction(writer, *body.S2C.KeepAliveInterval)
					handler.unlockAccount(pwd, writer)
					//handler.getSubInfo(writer)
				}
			case P_KeepAlive:
				{
					body, ok := msg.Body.(*keepalive.Response)
					log.Printf("read content, retErr=%d, retType=%d, retMsg=%s", *body.ErrCode, *body.RetType, *body.RetMsg)
					if !ok || *body.ErrCode != 0 {
						log.Printf("keepAlive resp failed, err=%s", *body.RetMsg)
						return
					}
				}
			default:
				handler.Resp <- msg
			}
		}
	}
}

func (handler *Handler) unlockAccount(pwd string, writer chan *Message) {
	writer <- &Message{
		ProtoID: uint32(P_Trd_UnlockTrade),
		Body: &trdunlocktrade.Request{
			C2S: &trdunlocktrade.C2S{
				Unlock: proto.Bool(true),
				PwdMD5: proto.String(pwd),
			},
		},
	}
}
func (handler *Handler) getSubInfo(writer chan *Message) {
	writer <- &Message{
		ProtoID: uint32(P_Qot_GetSubInfo),
		Body: &qotgetsubinfo.Request{
			C2S: &qotgetsubinfo.C2S{
				IsReqAllConn: proto.Bool(false),
			},
		},
	}
}
