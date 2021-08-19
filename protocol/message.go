package protocol

import (
	"bufio"
	"crypto/sha1"
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"log"
	"nitrohsu.com/futu/api/initconnect"
	"nitrohsu.com/futu/api/keepalive"
	"nitrohsu.com/futu/api/qotgetorderbook"
	"nitrohsu.com/futu/api/qotgetsubinfo"
	"nitrohsu.com/futu/api/qotsub"
	"nitrohsu.com/futu/api/trdgethistoryorderlist"
	"nitrohsu.com/futu/api/trdunlocktrade"
	"reflect"
)

type Message struct {
	HeaderFlag   [2]uint8  // default is 'FT'
	ProtoID      uint32    // protocol id, little Endian
	protoFmtType uint8     // 0-Protobuf,1-Json
	protoVer     uint8     // default is 0
	serialNo     uint32    // auto increment
	bodyLen      uint32    // package length
	arrBodySHA1  [20]uint8 // package body sha1
	arrReserved  [8]uint8  // keep extra
	Body         interface{}
}

const (
	HEADER_LENGTH = 2 + 4 + 1 + 1 + 4 + 4 + 20 + 8
)

var (
	jsonPb                = jsonpb.Marshaler{}
	globalSerialId uint32 = 10
)

const (
	P_InitConnect                 = 1001 //初始化连接
	P_GetGlobalState              = 1002 //获取全局状态
	P_Notify                      = 1003 //系统通知推送
	P_KeepAlive                   = 1004 //保活心跳
	P_Trd_GetAccList              = 2001 //获取业务账户列表
	P_Trd_UnlockTrade             = 2005 //解锁或锁定交易
	P_Trd_SubAccPush              = 2008 //订阅业务账户的交易推送数据
	P_Trd_GetFunds                = 2101 //获取账户资金
	P_Trd_GetPositionList         = 2102 //获取账户持仓
	P_Trd_GetMaxTrdQtys           = 2111 //获取最大交易数量
	P_Trd_GetOrderList            = 2201 //获取订单列表
	P_Trd_PlaceOrder              = 2202 //下单
	P_Trd_ModifyOrder             = 2205 //修改订单
	P_Trd_UpdateOrder             = 2208 //推送订单状态变动通知
	P_Trd_GetOrderFillList        = 2211 //获取成交列表
	P_Trd_UpdateOrderFill         = 2218 //推送成交通知
	P_Trd_GetHistoryOrderList     = 2221 //获取历史订单列表
	P_Trd_GetHistoryOrderFillList = 2222 //获取历史成交列表
	P_Qot_Sub                     = 3001 //订阅或者反订阅
	P_Qot_RegQotPush              = 3002 //注册推送
	P_Qot_GetSubInfo              = 3003 //获取订阅信息
	P_Qot_GetBasicQot             = 3004 //获取股票基本报价
	P_Qot_UpdateBasicQot          = 3005 //推送股票基本报价
	P_Qot_GetKL                   = 3006 //获取K线
	P_Qot_UpdateKL                = 3007 //推送K线
	P_Qot_GetRT                   = 3008 //获取分时
	P_Qot_UpdateRT                = 3009 //推送分时
	P_Qot_GetTicker               = 3010 //获取逐笔
	P_Qot_UpdateTicker            = 3011 //推送逐笔
	P_Qot_GetOrderBook            = 3012 //获取买卖盘
	P_Qot_UpdateOrderBook         = 3013 //推送买卖盘
	P_Qot_GetBroker               = 3014 //获取经纪队列
	P_Qot_UpdateBroker            = 3015 //推送经纪队列
	P_Qot_GetHistoryKL            = 3100 //从本地下载历史数据获取单只股票一段历史K线
	P_Qot_GetHistoryKLPoints      = 3101 //从本地下载历史数据获取多只股票多点历史K线
	P_Qot_GetRehab                = 3102 //从本地下载历史数据获取复权信息
	P_Qot_RequestHistoryKL        = 3103 //在线获取单只股票一段历史K线
	P_Qot_RequestRehab            = 3105 //在线获取单只股票复权信息
	P_Qot_GetTradeDate            = 3200 //获取市场交易日
	P_Qot_GetStaticInfo           = 3202 //获取股票静态信息
	P_Qot_GetSecuritySnapshot     = 3203 //获取股票快照
	P_Qot_GetPlateSet             = 3204 //获取板块集合下的板块
	P_Qot_GetPlateSecurity        = 3205 //获取板块下的股票
	P_Qot_GetReference            = 3206 //获取正股相关股票
	P_Qot_GetOwnerPlate           = 3207 //获取股票所属板块
	P_Qot_GetHoldingChangeList    = 3208 //获取持股变化列表
	P_Qot_GetOptionChain          = 3209 //获取期权链
	P_Qot_GetWarrant              = 3210 //获取窝轮
	P_Qot_GetCapitalFlow          = 3211 //获取资金流向
	P_Qot_GetCapitalDistribution  = 3212 //获取资金分布
	P_Qot_GetUserSecurity         = 3213 //获取自选股分组下的股票
	P_Qot_ModifyUserSecurity      = 3214 //修改自选股分组下的股票
	P_Qot_StockFilter             = 3215 //获取条件选股
	P_Qot_GetIpoList              = 3217 //获取新股
	P_Qot_GetFutureInfo           = 3218 //获取期货合约资料
	P_Qot_RequestTradeDate        = 3219 //获取市场交易日，在线拉取不在本地计算
)

func (message *Message) Write(writer *bufio.Writer) (nn int, err error) {
	//
	if message.Body == nil {
		return -1, errors.New("no target pb")
	}
	pb, ok := message.Body.(proto.Message)
	if !ok {
		log.Printf("write err. %d, %s", -1, errors.New("not pb message"))
	}
	body, err := proto.Marshal(pb)
	if err != nil {
		log.Printf("write err. %d, %s", -1, errors.Unwrap(err))
	}

	message.serialNo = globalSerialId
	globalSerialId += 1
	message.bodyLen = uint32(len(body))
	//
	content := make([]byte, HEADER_LENGTH)
	offset := 0
	content[offset] = 'F'
	offset += 1
	content[offset] = 'T'
	offset += 1
	binary.LittleEndian.PutUint32(content[offset:], message.ProtoID)
	offset += 4
	content[offset] = uint8(0)
	offset += 1
	content[offset] = uint8(0)
	offset += 1
	binary.LittleEndian.PutUint32(content[offset:], message.serialNo)
	offset += 4
	binary.LittleEndian.PutUint32(content[offset:], message.bodyLen)
	offset += 4
	//
	sum := sha1.Sum(body)
	for _, b := range sum {
		content[offset] = b
		offset += 1
	}
	binary.LittleEndian.PutUint64(content[offset:], uint64(0))
	//
	content = append(content, body...)
	headerN, headerErr := writer.Write(content)
	headerErr = writer.Flush()
	if headerErr != nil {
		return headerN, headerErr
	}
	jsonStr, _ := jsonPb.MarshalToString(pb)
	log.Printf("w serial=%2d, id=%4d, packLen=%2d, sockLen=%3d, body=%s",
		message.serialNo,
		message.ProtoID,
		message.bodyLen,
		headerN,
		jsonStr)
	//
	return headerN, nil
}

func (message *Message) Read(reader *bufio.Reader) (int, error) {
	// header
	header := make([]byte, HEADER_LENGTH)
	nn, err := reader.Read(header)
	if err != nil {
		return nn, err
	}
	message.HeaderFlag[0] = header[0]
	message.HeaderFlag[1] = header[1]
	message.ProtoID = binary.LittleEndian.Uint32(header[2:6])
	message.protoFmtType = header[6]
	message.protoVer = header[7]
	message.serialNo = binary.LittleEndian.Uint32(header[8:12])
	message.bodyLen = binary.LittleEndian.Uint32(header[12:16])

	i := 0
	var sha1Var [20]uint8
	for i = 0; i < 20; i++ {
		sha1Var[i] = header[16+i]
	}
	message.arrBodySHA1 = sha1Var
	// body
	body := make([]byte, message.bodyLen)
	var readTotal uint32 = 0
	for {
		read, connErr := reader.Read(body[readTotal:])
		if connErr != nil {
			return -1, connErr
		}
		readTotal += uint32(read)
		if readTotal == message.bodyLen {
			break
		}
	}

	if !reflect.DeepEqual(sha1.Sum(body), sha1Var) {
		return -1, errors.New("sha1 error")
	}

	pb := newResp(message.ProtoID)
	pbErr := proto.Unmarshal(body, pb)
	if pbErr != nil {
		return -1, pbErr
	}
	message.Body = pb

	jsonStr, _ := jsonPb.MarshalToString(pb)
	log.Printf("r serial=%2d, id=%4d, packLen=%2d, sockLen=%3d, body=%s",
		message.serialNo,
		message.ProtoID,
		message.bodyLen,
		readTotal,
		jsonStr)

	return int(readTotal), nil
}

func newResp(protoId uint32) proto.Message {
	switch protoId {
	case P_InitConnect:
		return &initconnect.Response{}
	case P_KeepAlive:
		return &keepalive.Response{}
	case P_Qot_GetOrderBook:
		return &qotgetorderbook.Response{}
	case P_Qot_Sub:
		return &qotsub.Response{}
	case P_Qot_GetSubInfo:
		return &qotgetsubinfo.Response{}
	case P_Trd_GetHistoryOrderList:
		return &trdgethistoryorderlist.Response{}
	case P_Trd_UnlockTrade:
		return &trdunlocktrade.Response{}
	}
	return nil
}
