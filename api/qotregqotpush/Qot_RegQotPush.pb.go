// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: Qot_RegQotPush.proto

package qotregqotpush

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "nitrohsu.com/futu/api/common"
	qotcommon "nitrohsu.com/futu/api/qotcommon"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type C2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityList  []*qotcommon.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`    //股票
	SubTypeList   []int32               `protobuf:"varint,2,rep,name=subTypeList" json:"subTypeList,omitempty"`     //Qot_Common.SubType,要注册到该连接的订阅类型
	RehabTypeList []int32               `protobuf:"varint,3,rep,name=rehabTypeList" json:"rehabTypeList,omitempty"` //Qot_Common.RehabType,复权类型,注册K线类型才生效,其他订阅类型忽略该参数,注册K线时该参数不指定默认前复权
	IsRegOrUnReg  *bool                 `protobuf:"varint,4,req,name=isRegOrUnReg" json:"isRegOrUnReg,omitempty"`   //注册或取消
	IsFirstPush   *bool                 `protobuf:"varint,5,opt,name=isFirstPush" json:"isFirstPush,omitempty"`     //注册后如果本地已有数据是否首推一次已存在数据,该参数不指定则默认true
}

func (x *C2S) Reset() {
	*x = C2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_RegQotPush_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S) ProtoMessage() {}

func (x *C2S) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_RegQotPush_proto_msgTypes[0]
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
	return file_Qot_RegQotPush_proto_rawDescGZIP(), []int{0}
}

func (x *C2S) GetSecurityList() []*qotcommon.Security {
	if x != nil {
		return x.SecurityList
	}
	return nil
}

func (x *C2S) GetSubTypeList() []int32 {
	if x != nil {
		return x.SubTypeList
	}
	return nil
}

func (x *C2S) GetRehabTypeList() []int32 {
	if x != nil {
		return x.RehabTypeList
	}
	return nil
}

func (x *C2S) GetIsRegOrUnReg() bool {
	if x != nil && x.IsRegOrUnReg != nil {
		return *x.IsRegOrUnReg
	}
	return false
}

func (x *C2S) GetIsFirstPush() bool {
	if x != nil && x.IsFirstPush != nil {
		return *x.IsFirstPush
	}
	return false
}

type S2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *S2C) Reset() {
	*x = S2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_RegQotPush_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C) ProtoMessage() {}

func (x *S2C) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_RegQotPush_proto_msgTypes[1]
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
	return file_Qot_RegQotPush_proto_rawDescGZIP(), []int{1}
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
		mi := &file_Qot_RegQotPush_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_RegQotPush_proto_msgTypes[2]
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
	return file_Qot_RegQotPush_proto_rawDescGZIP(), []int{2}
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

	RetType *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"` //RetType,返回结果
	RetMsg  *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C     *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
}

// Default values for Response fields.
const (
	Default_Response_RetType = int32(-400)
)

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Qot_RegQotPush_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_Qot_RegQotPush_proto_msgTypes[3]
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
	return file_Qot_RegQotPush_proto_rawDescGZIP(), []int{3}
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

var File_Qot_RegQotPush_proto protoreflect.FileDescriptor

var file_Qot_RegQotPush_proto_rawDesc = []byte{
	0x0a, 0x14, 0x51, 0x6f, 0x74, 0x5f, 0x52, 0x65, 0x67, 0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x51, 0x6f, 0x74, 0x5f, 0x52, 0x65, 0x67, 0x51,
	0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x03, 0x43, 0x32, 0x53, 0x12, 0x38,
	0x0a, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x54,
	0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x73,
	0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65,
	0x68, 0x61, 0x62, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0d, 0x72, 0x65, 0x68, 0x61, 0x62, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x4f, 0x72, 0x55, 0x6e, 0x52, 0x65, 0x67,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x4f, 0x72, 0x55,
	0x6e, 0x52, 0x65, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74, 0x50,
	0x75, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x50, 0x75, 0x73, 0x68, 0x22, 0x05, 0x0a, 0x03, 0x53, 0x32, 0x43, 0x22, 0x30, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x63, 0x32, 0x73, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x51, 0x6f, 0x74, 0x5f, 0x52, 0x65, 0x67, 0x51,
	0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x2e, 0x43, 0x32, 0x53, 0x52, 0x03, 0x63, 0x32, 0x73, 0x22,
	0x83, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x3a, 0x04, 0x2d,
	0x34, 0x30, 0x30, 0x52, 0x07, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25,
	0x0a, 0x03, 0x73, 0x32, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x51, 0x6f,
	0x74, 0x5f, 0x52, 0x65, 0x67, 0x51, 0x6f, 0x74, 0x50, 0x75, 0x73, 0x68, 0x2e, 0x53, 0x32, 0x43,
	0x52, 0x03, 0x73, 0x32, 0x63, 0x42, 0x44, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74,
	0x75, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x62, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x74, 0x75, 0x6f, 0x70, 0x65,
	0x6e, 0x2f, 0x66, 0x74, 0x61, 0x70, 0x69, 0x34, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x71, 0x6f,
	0x74, 0x72, 0x65, 0x67, 0x71, 0x6f, 0x74, 0x70, 0x75, 0x73, 0x68,
}

var (
	file_Qot_RegQotPush_proto_rawDescOnce sync.Once
	file_Qot_RegQotPush_proto_rawDescData = file_Qot_RegQotPush_proto_rawDesc
)

func file_Qot_RegQotPush_proto_rawDescGZIP() []byte {
	file_Qot_RegQotPush_proto_rawDescOnce.Do(func() {
		file_Qot_RegQotPush_proto_rawDescData = protoimpl.X.CompressGZIP(file_Qot_RegQotPush_proto_rawDescData)
	})
	return file_Qot_RegQotPush_proto_rawDescData
}

var file_Qot_RegQotPush_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Qot_RegQotPush_proto_goTypes = []interface{}{
	(*C2S)(nil),                // 0: Qot_RegQotPush.C2S
	(*S2C)(nil),                // 1: Qot_RegQotPush.S2C
	(*Request)(nil),            // 2: Qot_RegQotPush.Request
	(*Response)(nil),           // 3: Qot_RegQotPush.Response
	(*qotcommon.Security)(nil), // 4: Qot_Common.Security
}
var file_Qot_RegQotPush_proto_depIdxs = []int32{
	4, // 0: Qot_RegQotPush.C2S.securityList:type_name -> Qot_Common.Security
	0, // 1: Qot_RegQotPush.Request.c2s:type_name -> Qot_RegQotPush.C2S
	1, // 2: Qot_RegQotPush.Response.s2c:type_name -> Qot_RegQotPush.S2C
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Qot_RegQotPush_proto_init() }
func file_Qot_RegQotPush_proto_init() {
	if File_Qot_RegQotPush_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Qot_RegQotPush_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_RegQotPush_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_RegQotPush_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_Qot_RegQotPush_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_Qot_RegQotPush_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Qot_RegQotPush_proto_goTypes,
		DependencyIndexes: file_Qot_RegQotPush_proto_depIdxs,
		MessageInfos:      file_Qot_RegQotPush_proto_msgTypes,
	}.Build()
	File_Qot_RegQotPush_proto = out.File
	file_Qot_RegQotPush_proto_rawDesc = nil
	file_Qot_RegQotPush_proto_goTypes = nil
	file_Qot_RegQotPush_proto_depIdxs = nil
}
