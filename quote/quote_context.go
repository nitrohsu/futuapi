package quote

import (
	"github.com/golang/protobuf/proto"
	"log"
	"nitrohsu.com/futu/api/qotcommon"
	"nitrohsu.com/futu/api/qotgetorderbook"
	"nitrohsu.com/futu/api/qotgetsubinfo"
	"nitrohsu.com/futu/api/qotsub"
	"nitrohsu.com/futu/api/trdcommon"
	"nitrohsu.com/futu/api/trdgethistoryorderlist"
	"nitrohsu.com/futu/api/trdunlocktrade"
	"nitrohsu.com/futu/conf"
	"nitrohsu.com/futu/netmanager"
	"nitrohsu.com/futu/protocol"
	"time"
)

type Quote struct {
	Config    conf.Config
	NetMgr    *netmanager.NetManager
	FutuState *protocol.Handler
	SubStock  []*qotcommon.Security
	account   map[trdcommon.TrdMarket]*trdcommon.TrdHeader
	password  string
	unlock    bool
}

func (quote *Quote) Start() (err error) {
	//
	quote.NetMgr = &netmanager.NetManager{
		Host: quote.Config.Host,
		Port: quote.Config.Port,
	}
	err = quote.NetMgr.Connect()
	if err != nil {
		return
	}

	quote.account = make(map[trdcommon.TrdMarket]*trdcommon.TrdHeader)
	quote.account[trdcommon.TrdMarket_TrdMarket_US] = &trdcommon.TrdHeader{
		TrdEnv:    proto.Int32(int32(trdcommon.TrdEnv_TrdEnv_Real)),
		AccID:     proto.Uint64(0),
		TrdMarket: proto.Int32(int32(trdcommon.TrdMarket_TrdMarket_US)),
	}
	quote.account[trdcommon.TrdMarket_TrdMarket_HK] = &trdcommon.TrdHeader{
		TrdEnv:    proto.Int32(int32(trdcommon.TrdEnv_TrdEnv_Real)),
		AccID:     proto.Uint64(0),
		TrdMarket: proto.Int32(int32(trdcommon.TrdMarket_TrdMarket_HK)),
	}
	quote.password = "MD5 YOUR PASSWORD"
	quote.FutuState = &protocol.Handler{}
	//
	if futuErr := quote.FutuState.Init(quote.NetMgr.WriteBuffer, quote.NetMgr.ReadBuffer, quote.password); futuErr != nil {
		return futuErr
	}
	//
	go func() {
		for {
			select {
			case msg := <-quote.FutuState.Resp:
				switch msg.ProtoID {
				case protocol.P_Qot_Sub:
					quote.subResponse(msg.Body.(*qotsub.Response))
				case protocol.P_Qot_GetSubInfo:
					quote.subInfoListResponse(msg.Body.(*qotgetsubinfo.Response))
				case protocol.P_Trd_GetOrderList:
					quote.realOrderListResponse(msg.Body.(*qotgetorderbook.Response))
				case protocol.P_Trd_GetHistoryOrderList:
					quote.historyOrderListResponse(msg.Body.(*trdgethistoryorderlist.Response))
				case protocol.P_Trd_UnlockTrade:
					quote.unlockTradeResponse(msg.Body.(*trdunlocktrade.Response))
				}
			}
		}
	}()
	time.Sleep(time.Duration(10) * time.Second)
	//quote.realOrderListRequest(quote.NetMgr.WriteBuffer)
	quote.historyOrderListRequest(nil, quote.NetMgr.WriteBuffer)

	return
}

func (quote *Quote) realOrderListRequest(writer chan *protocol.Message) {

	stock := &qotcommon.Security{
		Market: proto.Int32(int32(qotcommon.QotMarket_QotMarket_US_Security)),
		Code:   proto.String("BA"),
	}
	exist := false
	for _, item := range quote.SubStock {
		if stock.GetCode() == item.GetCode() && stock.GetMarket() == item.GetMarket() {
			exist = true
		}
	}
	if !exist {
		subRequest(stock, writer)
	} else {
		log.Printf("subed stock, %d=%s", stock.GetMarket(), stock.GetCode())
	}
	writer <- &protocol.Message{
		ProtoID: protocol.P_Qot_GetOrderBook,
		Body: &qotgetorderbook.Request{
			C2S: &qotgetorderbook.C2S{
				Num:      proto.Int32(10),
				Security: stock,
			}},
	}
}

func (quote *Quote) historyOrderListRequest(markets []trdcommon.TrdMarket, writer chan *protocol.Message) {

	if markets == nil {
		markets = make([]trdcommon.TrdMarket, 2)
		markets[0] = trdcommon.TrdMarket_TrdMarket_US
		markets[1] = trdcommon.TrdMarket_TrdMarket_HK
	}
	for _, market := range markets {
		writer <- &protocol.Message{
			ProtoID: protocol.P_Trd_GetHistoryOrderList,
			Body: &trdgethistoryorderlist.Request{
				C2S: &trdgethistoryorderlist.C2S{
					Header: quote.account[market],
					FilterConditions: &trdcommon.TrdFilterConditions{
						CodeList: []string{},
						IdList:   []uint64{},
						//YYYY-MM-DD HH:MM:SS
						BeginTime: proto.String("2020-01-01 00:00:00"),
						EndTime:   proto.String("2020-03-15 00:00:00"),
					},
				}},
		}
	}
}
func (quote *Quote) realOrderListResponse(msg *qotgetorderbook.Response) {
	if msg.GetErrCode() == 0 {
		orderList := msg.S2C.OrderBookAskList
		log.Printf("%x", orderList)
	} else {
		log.Printf("orderList err. err=%s", *msg.RetMsg)
	}
}

func (quote *Quote) historyOrderListResponse(msg *trdgethistoryorderlist.Response) {
	if msg.GetErrCode() == 0 {
		header := msg.S2C.GetHeader()
		log.Printf("accId=%d, market=%s, env=%s", header.GetAccID(), trdcommon.TrdMarket_name[header.GetTrdMarket()], trdcommon.TrdEnv_name[header.GetTrdEnv()])
		orderList := msg.S2C.GetOrderList()
		for _, order := range orderList {
			log.Printf("%s,%s,%s,%d,%s,%s,%s,%s,%f,%f",
				order.GetCreateTime(),
				trdcommon.TrdMarket_name[header.GetTrdMarket()],
				order.GetUpdateTime(),
				order.GetOrderID(),
				trdcommon.OrderStatus_name[order.GetOrderStatus()],
				trdcommon.TrdSide_name[order.GetTrdSide()],
				order.GetCode(),
				order.GetName(),
				order.GetPrice(),
				order.GetQty())
		}
	} else {
		log.Printf("orderList err. err=%s", *msg.RetMsg)
	}
}

func subRequest(stock *qotcommon.Security, writer chan *protocol.Message) {
	writer <- &protocol.Message{
		ProtoID: protocol.P_Qot_Sub,
		Body: &qotsub.Request{
			C2S: &qotsub.C2S{
				IsFirstPush:  proto.Bool(false),
				IsSubOrUnSub: proto.Bool(true),
				SecurityList: []*qotcommon.Security{stock},
				SubTypeList: []int32{
					int32(qotcommon.SubType_SubType_OrderBook),
				},
			}},
	}
}

func (quote *Quote) subResponse(msg *qotsub.Response) {
	if *msg.ErrCode == 0 {
		log.Printf("sub succes")
	} else {
		log.Printf("sub err. err=%s", *msg.RetMsg)
	}
}

func (quote *Quote) subInfoListResponse(msg *qotgetsubinfo.Response) {
	if *msg.ErrCode == 0 {
		connSubInfoList := msg.S2C.ConnSubInfoList
		for _, connSubInfo := range connSubInfoList {
			for _, subInfo := range connSubInfo.SubInfoList {
				for _, stock := range subInfo.SecurityList {
					quote.SubStock = append(quote.SubStock, stock)
				}
			}
			log.Printf("getSubInfoList used=%d", connSubInfo.GetUsedQuota())
		}
	} else {
		log.Printf("getSubInfoList err. err=%s", *msg.RetMsg)
	}
}

func (quote *Quote) unlockTradeResponse(msg *trdunlocktrade.Response) {
	if *msg.ErrCode == 0 {
		quote.unlock = true
		log.Printf("unlockTrade success")
	} else {
		log.Printf("unlockTrade err. err=%s", *msg.RetMsg)
	}
}

func (quote *Quote) Close() {
	quote.NetMgr.Close()
	quote.FutuState.Close()
}
