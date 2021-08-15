package quote

import (
	"github.com/golang/protobuf/proto"
	"log"
	"nitrohsu.com/futu/api/Qot_Common"
	"nitrohsu.com/futu/api/Qot_GetOrderBook"
	"nitrohsu.com/futu/api/Qot_GetSubInfo"
	"nitrohsu.com/futu/api/Qot_Sub"
	"nitrohsu.com/futu/api/Trd_Common"
	"nitrohsu.com/futu/api/Trd_GetHistoryOrderList"
	"nitrohsu.com/futu/api/Trd_UnlockTrade"
	"nitrohsu.com/futu/conf"
	"nitrohsu.com/futu/netmanager"
	"nitrohsu.com/futu/protocol"
	"time"
)

type Quote struct {
	Config    conf.Config
	NetMgr    *netmanager.NetManager
	FutuState *protocol.Handler
	SubStock  []*Qot_Common.Security
	account   map[Trd_Common.TrdMarket]*Trd_Common.TrdHeader
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

	quote.account = make(map[Trd_Common.TrdMarket]*Trd_Common.TrdHeader)
	quote.account[Trd_Common.TrdMarket_TrdMarket_US] = &Trd_Common.TrdHeader{
		TrdEnv:    proto.Int32(int32(Trd_Common.TrdEnv_TrdEnv_Real)),
		AccID:     proto.Uint64(0),
		TrdMarket: proto.Int32(int32(Trd_Common.TrdMarket_TrdMarket_US)),
	}
	quote.account[Trd_Common.TrdMarket_TrdMarket_HK] = &Trd_Common.TrdHeader{
		TrdEnv:    proto.Int32(int32(Trd_Common.TrdEnv_TrdEnv_Real)),
		AccID:     proto.Uint64(0),
		TrdMarket: proto.Int32(int32(Trd_Common.TrdMarket_TrdMarket_HK)),
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
					quote.subResponse(msg.Body.(*Qot_Sub.Response))
				case protocol.P_Qot_GetSubInfo:
					quote.subInfoListResponse(msg.Body.(*Qot_GetSubInfo.Response))
				case protocol.P_Trd_GetOrderList:
					quote.realOrderListResponse(msg.Body.(*Qot_GetOrderBook.Response))
				case protocol.P_Trd_GetHistoryOrderList:
					quote.historyOrderListResponse(msg.Body.(*Trd_GetHistoryOrderList.Response))
				case protocol.P_Trd_UnlockTrade:
					quote.unlockTradeResponse(msg.Body.(*Trd_UnlockTrade.Response))
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

	stock := &Qot_Common.Security{
		Market: proto.Int32(int32(Qot_Common.QotMarket_QotMarket_US_Security)),
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
		Body: &Qot_GetOrderBook.Request{
			C2S: &Qot_GetOrderBook.C2S{
				Num:      proto.Int32(10),
				Security: stock,
			}},
	}
}

func (quote *Quote) historyOrderListRequest(markets []Trd_Common.TrdMarket, writer chan *protocol.Message) {

	if markets == nil {
		markets = make([]Trd_Common.TrdMarket, 2)
		markets[0] = Trd_Common.TrdMarket_TrdMarket_US
		markets[1] = Trd_Common.TrdMarket_TrdMarket_HK
	}
	for _, market := range markets {
		writer <- &protocol.Message{
			ProtoID: protocol.P_Trd_GetHistoryOrderList,
			Body: &Trd_GetHistoryOrderList.Request{
				C2S: &Trd_GetHistoryOrderList.C2S{
					Header: quote.account[market],
					FilterConditions: &Trd_Common.TrdFilterConditions{
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
func (quote *Quote) realOrderListResponse(msg *Qot_GetOrderBook.Response) {
	if msg.GetErrCode() == 0 {
		orderList := msg.S2C.OrderBookAskList
		log.Printf("%x", orderList)
	} else {
		log.Printf("orderList err. err=%s", *msg.RetMsg)
	}
}

func (quote *Quote) historyOrderListResponse(msg *Trd_GetHistoryOrderList.Response) {
	if msg.GetErrCode() == 0 {
		header := msg.S2C.GetHeader()
		log.Printf("accId=%d, market=%s, env=%s", header.GetAccID(), Trd_Common.TrdMarket_name[header.GetTrdMarket()], Trd_Common.TrdEnv_name[header.GetTrdEnv()])
		orderList := msg.S2C.GetOrderList()
		for _, order := range orderList {
			log.Printf("%s,%s,%s,%d,%s,%s,%s,%s,%f,%f",
				order.GetCreateTime(),
				Trd_Common.TrdMarket_name[header.GetTrdMarket()],
				order.GetUpdateTime(),
				order.GetOrderID(),
				Trd_Common.OrderStatus_name[order.GetOrderStatus()],
				Trd_Common.TrdSide_name[order.GetTrdSide()],
				order.GetCode(),
				order.GetName(),
				order.GetPrice(),
				order.GetQty())
		}
	} else {
		log.Printf("orderList err. err=%s", *msg.RetMsg)
	}
}

func subRequest(stock *Qot_Common.Security, writer chan *protocol.Message) {
	writer <- &protocol.Message{
		ProtoID: protocol.P_Qot_Sub,
		Body: &Qot_Sub.Request{
			C2S: &Qot_Sub.C2S{
				IsFirstPush:  proto.Bool(false),
				IsSubOrUnSub: proto.Bool(true),
				SecurityList: []*Qot_Common.Security{stock},
				SubTypeList: []int32{
					int32(Qot_Common.SubType_SubType_OrderBook),
					int32(Qot_Common.SubType_SubType_OrderDetail)},
			}},
	}
}

func (quote *Quote) subResponse(msg *Qot_Sub.Response) {
	if *msg.ErrCode == 0 {
		log.Printf("sub succes")
	} else {
		log.Printf("sub err. err=%s", *msg.RetMsg)
	}
}

func (quote *Quote) subInfoListResponse(msg *Qot_GetSubInfo.Response) {
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

func (quote *Quote) unlockTradeResponse(msg *Trd_UnlockTrade.Response) {
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
