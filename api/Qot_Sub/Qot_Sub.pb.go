// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_Sub.proto

package Qot_Sub

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	"nitrohsu.com/futu/api/Qot_Common"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type C2S struct {
	SecurityList         []*Qot_Common.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`
	SubTypeList          []int32                `protobuf:"varint,2,rep,name=subTypeList" json:"subTypeList,omitempty"`
	IsSubOrUnSub         *bool                  `protobuf:"varint,3,req,name=isSubOrUnSub" json:"isSubOrUnSub,omitempty"`
	IsRegOrUnRegPush     *bool                  `protobuf:"varint,4,opt,name=isRegOrUnRegPush" json:"isRegOrUnRegPush,omitempty"`
	RegPushRehabTypeList []int32                `protobuf:"varint,5,rep,name=regPushRehabTypeList" json:"regPushRehabTypeList,omitempty"`
	IsFirstPush          *bool                  `protobuf:"varint,6,opt,name=isFirstPush" json:"isFirstPush,omitempty"`
	IsUnsubAll           *bool                  `protobuf:"varint,7,opt,name=isUnsubAll" json:"isUnsubAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_532ae7909aa52c5a, []int{0}
}

func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (m *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(m, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetSecurityList() []*Qot_Common.Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

func (m *C2S) GetSubTypeList() []int32 {
	if m != nil {
		return m.SubTypeList
	}
	return nil
}

func (m *C2S) GetIsSubOrUnSub() bool {
	if m != nil && m.IsSubOrUnSub != nil {
		return *m.IsSubOrUnSub
	}
	return false
}

func (m *C2S) GetIsRegOrUnRegPush() bool {
	if m != nil && m.IsRegOrUnRegPush != nil {
		return *m.IsRegOrUnRegPush
	}
	return false
}

func (m *C2S) GetRegPushRehabTypeList() []int32 {
	if m != nil {
		return m.RegPushRehabTypeList
	}
	return nil
}

func (m *C2S) GetIsFirstPush() bool {
	if m != nil && m.IsFirstPush != nil {
		return *m.IsFirstPush
	}
	return false
}

func (m *C2S) GetIsUnsubAll() bool {
	if m != nil && m.IsUnsubAll != nil {
		return *m.IsUnsubAll
	}
	return false
}

type S2C struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_532ae7909aa52c5a, []int{1}
}

func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (m *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(m, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_532ae7909aa52c5a, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_532ae7909aa52c5a, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Qot_Sub.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_Sub.S2C")
	proto.RegisterType((*Request)(nil), "Qot_Sub.Request")
	proto.RegisterType((*Response)(nil), "Qot_Sub.Response")
}

func init() {
	proto.RegisterFile("Qot_Sub.proto", fileDescriptor_532ae7909aa52c5a)
}

var fileDescriptor_532ae7909aa52c5a = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x8b, 0xdb, 0x30,
	0x10, 0x85, 0xb1, 0x1d, 0xc7, 0xe9, 0xc4, 0x85, 0xa0, 0xa6, 0x45, 0xe4, 0x10, 0x84, 0x4f, 0x6e,
	0xa1, 0x26, 0x88, 0x1e, 0x4a, 0x6f, 0xad, 0xa1, 0xa7, 0x96, 0xb6, 0x52, 0x73, 0x5e, 0x62, 0xaf,
	0x36, 0x11, 0x24, 0x96, 0x57, 0x23, 0x1d, 0x02, 0xfb, 0x7f, 0xf6, 0x6f, 0x2e, 0x76, 0x12, 0xd6,
	0x61, 0xf7, 0xe6, 0xf7, 0xbd, 0x37, 0xf3, 0x06, 0x0b, 0xde, 0xfe, 0x33, 0xee, 0x46, 0xfa, 0xaa,
	0x68, 0xad, 0x71, 0x86, 0x24, 0x67, 0xb9, 0x48, 0x4b, 0x73, 0x38, 0x98, 0xe6, 0x84, 0x17, 0xb3,
	0x0e, 0x0f, 0x49, 0xf6, 0x18, 0x42, 0x54, 0x72, 0x49, 0xbe, 0x42, 0x8a, 0xaa, 0xf6, 0x56, 0xbb,
	0xe3, 0x2f, 0x8d, 0x8e, 0x06, 0x2c, 0xca, 0xa7, 0x7c, 0x5e, 0x0c, 0x06, 0xe4, 0xd9, 0x17, 0x57,
	0x49, 0xc2, 0x60, 0x8a, 0xbe, 0xfa, 0x7f, 0x6c, 0x55, 0x3f, 0x18, 0xb2, 0x28, 0x8f, 0xc5, 0x10,
	0x91, 0x0c, 0x52, 0x8d, 0xd2, 0x57, 0x7f, 0xec, 0xba, 0x91, 0xbe, 0xa2, 0x11, 0x0b, 0xf3, 0x89,
	0xb8, 0x62, 0xe4, 0x13, 0xcc, 0x34, 0x0a, 0xb5, 0xed, 0xb4, 0x50, 0xdb, 0xbf, 0x1e, 0x77, 0x74,
	0xc4, 0x82, 0x7c, 0x22, 0x5e, 0x70, 0xc2, 0x61, 0x6e, 0x4f, 0x9f, 0x42, 0xed, 0x36, 0xcf, 0xd5,
	0x71, 0x5f, 0xfd, 0xaa, 0xd7, 0x5d, 0xa9, 0xf1, 0xa7, 0xb6, 0xe8, 0xfa, 0xd5, 0xe3, 0x7e, 0xf5,
	0x10, 0x91, 0x25, 0x80, 0xc6, 0x75, 0x83, 0xbe, 0xfa, 0xbe, 0xdf, 0xd3, 0xa4, 0x0f, 0x0c, 0x48,
	0x16, 0x43, 0x24, 0x79, 0x99, 0x7d, 0x84, 0x44, 0xa8, 0x7b, 0xaf, 0xd0, 0x91, 0x25, 0x44, 0x35,
	0x47, 0x1a, 0xb0, 0x30, 0x9f, 0xf2, 0xb4, 0xb8, 0xbc, 0x40, 0xc9, 0xa5, 0xe8, 0x8c, 0xec, 0x01,
	0x26, 0x42, 0x61, 0x6b, 0x1a, 0x54, 0x64, 0x09, 0x89, 0x55, 0xae, 0x3b, 0xa7, 0xcf, 0xc7, 0xdf,
	0x46, 0x9f, 0xbf, 0xac, 0x56, 0xe2, 0x02, 0xc9, 0x07, 0x18, 0x5b, 0xe5, 0x7e, 0xe3, 0x96, 0x86,
	0x2c, 0xc8, 0xdf, 0x88, 0xb3, 0x22, 0x14, 0x12, 0x65, 0x6d, 0x69, 0x6e, 0x15, 0x8d, 0x58, 0x90,
	0xc7, 0xe2, 0x22, 0xbb, 0x76, 0xe4, 0x75, 0xff, 0x93, 0x86, 0xed, 0x92, 0x97, 0xa2, 0x33, 0x7e,
	0xbc, 0x87, 0x77, 0xb5, 0x39, 0x14, 0x77, 0xde, 0xf9, 0xc2, 0xb4, 0xaa, 0xd9, 0xb4, 0xba, 0x68,
	0xab, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xab, 0xb0, 0x67, 0x29, 0x02, 0x00, 0x00,
}
