// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: GetDelayStatistics.proto

package getdelaystatistics

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "nitrohsu.com/futu/api/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DelayStatisticsType int32

const (
	DelayStatisticsType_DelayStatisticsType_Unkonw     DelayStatisticsType = 0 //未知类型
	DelayStatisticsType_DelayStatisticsType_QotPush    DelayStatisticsType = 1 //行情推送统计
	DelayStatisticsType_DelayStatisticsType_ReqReply   DelayStatisticsType = 2 //请求回应统计
	DelayStatisticsType_DelayStatisticsType_PlaceOrder DelayStatisticsType = 3 //下单统计
)

// Enum value maps for DelayStatisticsType.
var (
	DelayStatisticsType_name = map[int32]string{
		0: "DelayStatisticsType_Unkonw",
		1: "DelayStatisticsType_QotPush",
		2: "DelayStatisticsType_ReqReply",
		3: "DelayStatisticsType_PlaceOrder",
	}
	DelayStatisticsType_value = map[string]int32{
		"DelayStatisticsType_Unkonw":     0,
		"DelayStatisticsType_QotPush":    1,
		"DelayStatisticsType_ReqReply":   2,
		"DelayStatisticsType_PlaceOrder": 3,
	}
)

func (x DelayStatisticsType) Enum() *DelayStatisticsType {
	p := new(DelayStatisticsType)
	*p = x
	return p
}

func (x DelayStatisticsType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DelayStatisticsType) Descriptor() protoreflect.EnumDescriptor {
	return file_GetDelayStatistics_proto_enumTypes[0].Descriptor()
}

func (DelayStatisticsType) Type() protoreflect.EnumType {
	return &file_GetDelayStatistics_proto_enumTypes[0]
}

func (x DelayStatisticsType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DelayStatisticsType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DelayStatisticsType(num)
	return nil
}

// Deprecated: Use DelayStatisticsType.Descriptor instead.
func (DelayStatisticsType) EnumDescriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{0}
}

//某段时间的统计数据
//SR表示服务器收到数据，目前只有港股支持SR字段，SS表示服务器发出数据
//CR表示OpenD收到数据，CS表示OpenD发出数据
type QotPushStage int32

const (
	QotPushStage_QotPushStage_Unkonw QotPushStage = 0 // 未知
	QotPushStage_QotPushStage_SR2SS  QotPushStage = 1 //统计服务端处理耗时
	QotPushStage_QotPushStage_SS2CR  QotPushStage = 2 //统计网络耗时
	QotPushStage_QotPushStage_CR2CS  QotPushStage = 3 //统计OpenD处理耗时
	QotPushStage_QotPushStage_SS2CS  QotPushStage = 4 //统计服务器发出到OpenD发出的处理耗时
	QotPushStage_QotPushStage_SR2CS  QotPushStage = 5 //统计服务器收到数据到OpenD发出的处理耗时
)

// Enum value maps for QotPushStage.
var (
	QotPushStage_name = map[int32]string{
		0: "QotPushStage_Unkonw",
		1: "QotPushStage_SR2SS",
		2: "QotPushStage_SS2CR",
		3: "QotPushStage_CR2CS",
		4: "QotPushStage_SS2CS",
		5: "QotPushStage_SR2CS",
	}
	QotPushStage_value = map[string]int32{
		"QotPushStage_Unkonw": 0,
		"QotPushStage_SR2SS":  1,
		"QotPushStage_SS2CR":  2,
		"QotPushStage_CR2CS":  3,
		"QotPushStage_SS2CS":  4,
		"QotPushStage_SR2CS":  5,
	}
)

func (x QotPushStage) Enum() *QotPushStage {
	p := new(QotPushStage)
	*p = x
	return p
}

func (x QotPushStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QotPushStage) Descriptor() protoreflect.EnumDescriptor {
	return file_GetDelayStatistics_proto_enumTypes[1].Descriptor()
}

func (QotPushStage) Type() protoreflect.EnumType {
	return &file_GetDelayStatistics_proto_enumTypes[1]
}

func (x QotPushStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *QotPushStage) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = QotPushStage(num)
	return nil
}

// Deprecated: Use QotPushStage.Descriptor instead.
func (QotPushStage) EnumDescriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{1}
}

//行情推送类型
type QotPushType int32

const (
	QotPushType_QotPushType_Unkonw    QotPushType = 0 // 未知
	QotPushType_QotPushType_Price     QotPushType = 1 //最新价
	QotPushType_QotPushType_Ticker    QotPushType = 2 //逐笔
	QotPushType_QotPushType_OrderBook QotPushType = 3 //摆盘
	QotPushType_QotPushType_Broker    QotPushType = 4 //经纪队列
)

// Enum value maps for QotPushType.
var (
	QotPushType_name = map[int32]string{
		0: "QotPushType_Unkonw",
		1: "QotPushType_Price",
		2: "QotPushType_Ticker",
		3: "QotPushType_OrderBook",
		4: "QotPushType_Broker",
	}
	QotPushType_value = map[string]int32{
		"QotPushType_Unkonw":    0,
		"QotPushType_Price":     1,
		"QotPushType_Ticker":    2,
		"QotPushType_OrderBook": 3,
		"QotPushType_Broker":    4,
	}
)

func (x QotPushType) Enum() *QotPushType {
	p := new(QotPushType)
	*p = x
	return p
}

func (x QotPushType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QotPushType) Descriptor() protoreflect.EnumDescriptor {
	return file_GetDelayStatistics_proto_enumTypes[2].Descriptor()
}

func (QotPushType) Type() protoreflect.EnumType {
	return &file_GetDelayStatistics_proto_enumTypes[2]
}

func (x QotPushType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *QotPushType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = QotPushType(num)
	return nil
}

// Deprecated: Use QotPushType.Descriptor instead.
func (QotPushType) EnumDescriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{2}
}

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeList     []int32 `protobuf:"varint,1,rep,name=typeList" json:"typeList,omitempty"`         //统计数据类型，DelayStatisticsType
	QotPushStage *int32  `protobuf:"varint,2,opt,name=qotPushStage" json:"qotPushStage,omitempty"` //行情推送统计的区间，行情推送统计时有效，QotPushStage
	SegmentList  []int32 `protobuf:"varint,3,rep,name=segmentList" json:"segmentList,omitempty"`   //统计分段，默认100ms以下以2ms分段，100ms以上以500，1000，2000，-1分段，-1表示无穷大。
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDelayStatistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_GetDelayStatistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S.ProtoReflect.Descriptor instead.
func (*C2S) Descriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetTypeList() []int32 {
	if x != nil {
		return x.TypeList
	}
	return nil
}

func (x *C2S) GetQotPushStage() int32 {
	if x != nil && x.QotPushStage != nil {
		return *x.QotPushStage
	}
	return 0
}

func (x *C2S) GetSegmentList() []int32 {
	if x != nil {
		return x.SegmentList
	}
	return nil
}

type DelayStatisticsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//范围左闭右开，[begin,end)
	Begin           *int32   `protobuf:"varint,1,req,name=begin" json:"begin,omitempty"`                      //耗时范围起点，毫秒单位
	End             *int32   `protobuf:"varint,2,req,name=end" json:"end,omitempty"`                          //耗时范围结束，毫秒单位
	Count           *int32   `protobuf:"varint,3,req,name=count" json:"count,omitempty"`                      //个数
	Proportion      *float32 `protobuf:"fixed32,4,req,name=proportion" json:"proportion,omitempty"`           //占比, %
	CumulativeRatio *float32 `protobuf:"fixed32,5,req,name=cumulativeRatio" json:"cumulativeRatio,omitempty"` //累计占比, %
}

func (x *DelayStatisticsItem) Reset() {
	*x = DelayStatisticsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDelayStatistics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelayStatisticsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelayStatisticsItem) ProtoMessage() {}

func (x *DelayStatisticsItem) ProtoReflect() protoreflect.Message {
	mi := &file_GetDelayStatistics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelayStatisticsItem.ProtoReflect.Descriptor instead.
func (*DelayStatisticsItem) Descriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{1}
}

func (x *DelayStatisticsItem) GetBegin() int32 {
	if x != nil && x.Begin != nil {
		return *x.Begin
	}
	return 0
}

func (x *DelayStatisticsItem) GetEnd() int32 {
	if x != nil && x.End != nil {
		return *x.End
	}
	return 0
}

func (x *DelayStatisticsItem) GetCount() int32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *DelayStatisticsItem) GetProportion() float32 {
	if x != nil && x.Proportion != nil {
		return *x.Proportion
	}
	return 0
}

func (x *DelayStatisticsItem) GetCumulativeRatio() float32 {
	if x != nil && x.CumulativeRatio != nil {
		return *x.CumulativeRatio
	}
	return 0
}

type DelayStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QotPushType *int32                 `protobuf:"varint,1,req,name=qotPushType" json:"qotPushType,omitempty"` //行情推送类型,QotPushType
	ItemList    []*DelayStatisticsItem `protobuf:"bytes,2,rep,name=itemList" json:"itemList,omitempty"`        //统计信息
	DelayAvg    *float32               `protobuf:"fixed32,3,req,name=delayAvg" json:"delayAvg,omitempty"`      //平均延迟
	Count       *int32                 `protobuf:"varint,4,req,name=count" json:"count,omitempty"`             //总包数
}

func (x *DelayStatistics) Reset() {
	*x = DelayStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDelayStatistics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelayStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelayStatistics) ProtoMessage() {}

func (x *DelayStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_GetDelayStatistics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelayStatistics.ProtoReflect.Descriptor instead.
func (*DelayStatistics) Descriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{2}
}

func (x *DelayStatistics) GetQotPushType() int32 {
	if x != nil && x.QotPushType != nil {
		return *x.QotPushType
	}
	return 0
}

func (x *DelayStatistics) GetItemList() []*DelayStatisticsItem {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *DelayStatistics) GetDelayAvg() float32 {
	if x != nil && x.DelayAvg != nil {
		return *x.DelayAvg
	}
	return 0
}

func (x *DelayStatistics) GetCount() int32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

type ReqReplyStatisticsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoID      *int32   `protobuf:"varint,1,req,name=protoID" json:"protoID,omitempty"`            //协议ID
	Count        *int32   `protobuf:"varint,2,req,name=count" json:"count,omitempty"`                //请求个数
	TotalCostAvg *float32 `protobuf:"fixed32,3,req,name=totalCostAvg" json:"totalCostAvg,omitempty"` //平均总耗时，毫秒单位
	OpenDCostAvg *float32 `protobuf:"fixed32,4,req,name=openDCostAvg" json:"openDCostAvg,omitempty"` //平均OpenD耗时，毫秒单位
	NetDelayAvg  *float32 `protobuf:"fixed32,5,req,name=netDelayAvg" json:"netDelayAvg,omitempty"`   //平均网络耗时，非当时实际请求网络耗时，毫秒单位
	IsLocalReply *bool    `protobuf:"varint,6,req,name=isLocalReply" json:"isLocalReply,omitempty"`  //是否本地直接回包，没有向服务器请求数据
}

func (x *ReqReplyStatisticsItem) Reset() {
	*x = ReqReplyStatisticsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDelayStatistics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqReplyStatisticsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqReplyStatisticsItem) ProtoMessage() {}

func (x *ReqReplyStatisticsItem) ProtoReflect() protoreflect.Message {
	mi := &file_GetDelayStatistics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqReplyStatisticsItem.ProtoReflect.Descriptor instead.
func (*ReqReplyStatisticsItem) Descriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{3}
}

func (x *ReqReplyStatisticsItem) GetProtoID() int32 {
	if x != nil && x.ProtoID != nil {
		return *x.ProtoID
	}
	return 0
}

func (x *ReqReplyStatisticsItem) GetCount() int32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *ReqReplyStatisticsItem) GetTotalCostAvg() float32 {
	if x != nil && x.TotalCostAvg != nil {
		return *x.TotalCostAvg
	}
	return 0
}

func (x *ReqReplyStatisticsItem) GetOpenDCostAvg() float32 {
	if x != nil && x.OpenDCostAvg != nil {
		return *x.OpenDCostAvg
	}
	return 0
}

func (x *ReqReplyStatisticsItem) GetNetDelayAvg() float32 {
	if x != nil && x.NetDelayAvg != nil {
		return *x.NetDelayAvg
	}
	return 0
}

func (x *ReqReplyStatisticsItem) GetIsLocalReply() bool {
	if x != nil && x.IsLocalReply != nil {
		return *x.IsLocalReply
	}
	return false
}

type PlaceOrderStatisticsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID    *string  `protobuf:"bytes,1,req,name=orderID" json:"orderID,omitempty"`         //订单ID
	TotalCost  *float32 `protobuf:"fixed32,2,req,name=totalCost" json:"totalCost,omitempty"`   //总耗时，毫秒单位
	OpenDCost  *float32 `protobuf:"fixed32,3,req,name=openDCost" json:"openDCost,omitempty"`   //OpenD耗时，毫秒单位
	NetDelay   *float32 `protobuf:"fixed32,4,req,name=netDelay" json:"netDelay,omitempty"`     //网络耗时，非当时实际请求网络耗时，毫秒单位
	UpdateCost *float32 `protobuf:"fixed32,5,req,name=updateCost" json:"updateCost,omitempty"` //订单回包后到接收到订单下到交易所的耗时，毫秒单位
}

func (x *PlaceOrderStatisticsItem) Reset() {
	*x = PlaceOrderStatisticsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDelayStatistics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceOrderStatisticsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceOrderStatisticsItem) ProtoMessage() {}

func (x *PlaceOrderStatisticsItem) ProtoReflect() protoreflect.Message {
	mi := &file_GetDelayStatistics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceOrderStatisticsItem.ProtoReflect.Descriptor instead.
func (*PlaceOrderStatisticsItem) Descriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{4}
}

func (x *PlaceOrderStatisticsItem) GetOrderID() string {
	if x != nil && x.OrderID != nil {
		return *x.OrderID
	}
	return ""
}

func (x *PlaceOrderStatisticsItem) GetTotalCost() float32 {
	if x != nil && x.TotalCost != nil {
		return *x.TotalCost
	}
	return 0
}

func (x *PlaceOrderStatisticsItem) GetOpenDCost() float32 {
	if x != nil && x.OpenDCost != nil {
		return *x.OpenDCost
	}
	return 0
}

func (x *PlaceOrderStatisticsItem) GetNetDelay() float32 {
	if x != nil && x.NetDelay != nil {
		return *x.NetDelay
	}
	return 0
}

func (x *PlaceOrderStatisticsItem) GetUpdateCost() float32 {
	if x != nil && x.UpdateCost != nil {
		return *x.UpdateCost
	}
	return 0
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QotPushStatisticsList    []*DelayStatistics          `protobuf:"bytes,1,rep,name=qotPushStatisticsList" json:"qotPushStatisticsList,omitempty"`       //行情推送延迟统计
	ReqReplyStatisticsList   []*ReqReplyStatisticsItem   `protobuf:"bytes,2,rep,name=reqReplyStatisticsList" json:"reqReplyStatisticsList,omitempty"`     //请求延迟统计
	PlaceOrderStatisticsList []*PlaceOrderStatisticsItem `protobuf:"bytes,3,rep,name=placeOrderStatisticsList" json:"placeOrderStatisticsList,omitempty"` //下单延迟统计
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDelayStatistics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_GetDelayStatistics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C.ProtoReflect.Descriptor instead.
func (*S2C) Descriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{5}
}

func (x *S2C) GetQotPushStatisticsList() []*DelayStatistics {
	if x != nil {
		return x.QotPushStatisticsList
	}
	return nil
}

func (x *S2C) GetReqReplyStatisticsList() []*ReqReplyStatisticsItem {
	if x != nil {
		return x.ReqReplyStatisticsList
	}
	return nil
}

func (x *S2C) GetPlaceOrderStatisticsList() []*PlaceOrderStatisticsItem {
	if x != nil {
		return x.PlaceOrderStatisticsList
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2S *C2S `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDelayStatistics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_GetDelayStatistics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{6}
}

func (x *Request) GetC2S() *C2S {
	if x != nil {
		return x.C2S
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //返回结果，参见Common.RetType的枚举定义
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`             //返回结果描述
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`          //错误码，客户端一般通过retType和retMsg来判断结果和详情，errCode只做日志记录，仅在个别协议失败时对账用
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetDelayStatistics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_GetDelayStatistics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_GetDelayStatistics_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetRetType() int32 {
	if x != nil && x.RetType != nil {
		return *x.RetType
	}
	return Default_Response_RetType
}

func (x *Response) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *Response) GetErrCode() int32 {
	if x != nil && x.ErrCode != nil {
		return *x.ErrCode
	}
	return 0
}

func (x *Response) GetS2C() *S2C {
	if x != nil {
		return x.S2C
	}
	return nil
}

var File_GetDelayStatistics_proto protoreflect.FileDescriptor

var file_GetDelayStatistics_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x0c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x03,
	0x43, 0x32, 0x53, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x71, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x71, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x05, 0x62, 0x65,
	0x67, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x02, 0x28, 0x02, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x05,
	0x20, 0x02, 0x28, 0x02, 0x52, 0x0f, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x52, 0x61, 0x74, 0x69, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x6f, 0x74,
	0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0b,
	0x71, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x69,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x67, 0x18, 0x03, 0x20, 0x02,
	0x28, 0x02, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xd6, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x76, 0x67, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x02, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x76,
	0x67, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x43, 0x6f, 0x73, 0x74, 0x41, 0x76,
	0x67, 0x18, 0x04, 0x20, 0x02, 0x28, 0x02, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x43, 0x6f,
	0x73, 0x74, 0x41, 0x76, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x41, 0x76, 0x67, 0x18, 0x05, 0x20, 0x02, 0x28, 0x02, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0c, 0x69,
	0x73, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xac, 0x01, 0x0a, 0x18,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x02, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x02, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x04, 0x20, 0x02, 0x28, 0x02,
	0x52, 0x08, 0x6e, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28, 0x02, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x22, 0xae, 0x02, 0x0a, 0x03, 0x53,
	0x32, 0x43, 0x12, 0x59, 0x0a, 0x15, 0x71, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x15, 0x71, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x62, 0x0a,
	0x16, 0x72, 0x65, 0x71, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x16, 0x72, 0x65, 0x71, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x68, 0x0a, 0x18, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x18, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32,
	0x73, 0x22, 0x87, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a,
	0x04, 0x2d, 0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x29, 0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x53, 0x32, 0x43, 0x52, 0x03, 0x73, 0x32, 0x63, 0x2a, 0x9c, 0x01, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6f, 0x6e,
	0x77, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x51, 0x6f, 0x74, 0x50, 0x75,
	0x73, 0x68, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x10, 0x02, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x10, 0x03, 0x2a, 0x9f, 0x01, 0x0a, 0x0c, 0x51,
	0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x51,
	0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6f,
	0x6e, 0x77, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x53, 0x52, 0x32, 0x53, 0x53, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x53, 0x53, 0x32,
	0x43, 0x52, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x43, 0x52, 0x32, 0x43, 0x53, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12,
	0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x53, 0x53, 0x32,
	0x43, 0x53, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x53, 0x52, 0x32, 0x43, 0x53, 0x10, 0x05, 0x2a, 0x87, 0x01, 0x0a,
	0x0b, 0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x55, 0x6e, 0x6b, 0x6f,
	0x6e, 0x77, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x5f, 0x50, 0x72, 0x69, 0x63, 0x65, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x51,
	0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x72, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x5f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x10, 0x03, 0x12, 0x16,
	0x0a, 0x12, 0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x10, 0x04, 0x42, 0x49, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75,
	0x74, 0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x74, 0x75, 0x6f, 0x70,
	0x65, 0x6e, 0x2f, 0x66, 0x74, 0x61, 0x70, 0x69, 0x34, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x67,
	0x65, 0x74, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73,
}

var (
	file_GetDelayStatistics_proto_rawDescOnce sync.Once
	file_GetDelayStatistics_proto_rawDescData = file_GetDelayStatistics_proto_rawDesc
)

func file_GetDelayStatistics_proto_rawDescGZIP() []byte {
	file_GetDelayStatistics_proto_rawDescOnce.Do(func() {
		file_GetDelayStatistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetDelayStatistics_proto_rawDescData)
	})
	return file_GetDelayStatistics_proto_rawDescData
}

var file_GetDelayStatistics_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_GetDelayStatistics_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_GetDelayStatistics_proto_goTypes = []interface{}{
	(DelayStatisticsType)(0),         // 0: GetDelayStatistics.DelayStatisticsType
	(QotPushStage)(0),                // 1: GetDelayStatistics.QotPushStage
	(QotPushType)(0),                 // 2: GetDelayStatistics.QotPushType
	(*C2S)(nil),                      // 3: GetDelayStatistics.C2S
	(*DelayStatisticsItem)(nil),      // 4: GetDelayStatistics.DelayStatisticsItem
	(*DelayStatistics)(nil),          // 5: GetDelayStatistics.DelayStatistics
	(*ReqReplyStatisticsItem)(nil),   // 6: GetDelayStatistics.ReqReplyStatisticsItem
	(*PlaceOrderStatisticsItem)(nil), // 7: GetDelayStatistics.PlaceOrderStatisticsItem
	(*S2C)(nil),                      // 8: GetDelayStatistics.S2C
	(*Request)(nil),                  // 9: GetDelayStatistics.Request
	(*Response)(nil),                 // 10: GetDelayStatistics.Response
}
var file_GetDelayStatistics_proto_depIdxs = []int32{
	4, // 0: GetDelayStatistics.DelayStatistics.itemList:type_name -> GetDelayStatistics.DelayStatisticsItem
	5, // 1: GetDelayStatistics.S2C.qotPushStatisticsList:type_name -> GetDelayStatistics.DelayStatistics
	6, // 2: GetDelayStatistics.S2C.reqReplyStatisticsList:type_name -> GetDelayStatistics.ReqReplyStatisticsItem
	7, // 3: GetDelayStatistics.S2C.placeOrderStatisticsList:type_name -> GetDelayStatistics.PlaceOrderStatisticsItem
	3, // 4: GetDelayStatistics.Request.c2s:type_name -> GetDelayStatistics.C2S
	8, // 5: GetDelayStatistics.Response.s2c:type_name -> GetDelayStatistics.S2C
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_GetDelayStatistics_proto_init() }
func file_GetDelayStatistics_proto_init() {
	if File_GetDelayStatistics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetDelayStatistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetDelayStatistics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelayStatisticsItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetDelayStatistics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelayStatistics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetDelayStatistics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqReplyStatisticsItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetDelayStatistics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaceOrderStatisticsItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetDelayStatistics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetDelayStatistics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_GetDelayStatistics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GetDelayStatistics_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetDelayStatistics_proto_goTypes,
		DependencyIndexes: file_GetDelayStatistics_proto_depIdxs,
		EnumInfos:         file_GetDelayStatistics_proto_enumTypes,
		MessageInfos:      file_GetDelayStatistics_proto_msgTypes,
	}.Build()
	File_GetDelayStatistics_proto = out.File
	file_GetDelayStatistics_proto_rawDesc = nil
	file_GetDelayStatistics_proto_goTypes = nil
	file_GetDelayStatistics_proto_depIdxs = nil
}
