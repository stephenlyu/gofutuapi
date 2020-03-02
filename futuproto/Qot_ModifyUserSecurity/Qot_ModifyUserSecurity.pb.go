// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_ModifyUserSecurity.proto

/*
Package Qot_ModifyUserSecurity is a generated protocol buffer package.

It is generated from these files:
	Qot_ModifyUserSecurity.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package Qot_ModifyUserSecurity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stephenlyu/gofutuapi/futuproto/Common"
import Qot_Common "github.com/stephenlyu/gofutuapi/futuproto/Qot_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ModifyUserSecurityOp int32

const (
	ModifyUserSecurityOp_ModifyUserSecurityOp_Unknown ModifyUserSecurityOp = 0
	ModifyUserSecurityOp_ModifyUserSecurityOp_Add     ModifyUserSecurityOp = 1
	ModifyUserSecurityOp_ModifyUserSecurityOp_Del     ModifyUserSecurityOp = 2
	ModifyUserSecurityOp_ModifyUserSecurityOp_MoveOut ModifyUserSecurityOp = 3
)

var ModifyUserSecurityOp_name = map[int32]string{
	0: "ModifyUserSecurityOp_Unknown",
	1: "ModifyUserSecurityOp_Add",
	2: "ModifyUserSecurityOp_Del",
	3: "ModifyUserSecurityOp_MoveOut",
}
var ModifyUserSecurityOp_value = map[string]int32{
	"ModifyUserSecurityOp_Unknown": 0,
	"ModifyUserSecurityOp_Add":     1,
	"ModifyUserSecurityOp_Del":     2,
	"ModifyUserSecurityOp_MoveOut": 3,
}

func (x ModifyUserSecurityOp) Enum() *ModifyUserSecurityOp {
	p := new(ModifyUserSecurityOp)
	*p = x
	return p
}
func (x ModifyUserSecurityOp) String() string {
	return proto.EnumName(ModifyUserSecurityOp_name, int32(x))
}
func (x *ModifyUserSecurityOp) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ModifyUserSecurityOp_value, data, "ModifyUserSecurityOp")
	if err != nil {
		return err
	}
	*x = ModifyUserSecurityOp(value)
	return nil
}
func (ModifyUserSecurityOp) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type C2S struct {
	GroupName        *string                `protobuf:"bytes,1,req,name=groupName" json:"groupName,omitempty"`
	Op               *int32                 `protobuf:"varint,2,req,name=op" json:"op,omitempty"`
	SecurityList     []*Qot_Common.Security `protobuf:"bytes,3,rep,name=securityList" json:"securityList,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *C2S) GetOp() int32 {
	if m != nil && m.Op != nil {
		return *m.Op
	}
	return 0
}

func (m *C2S) GetSecurityList() []*Qot_Common.Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

type S2C struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Request struct {
	C2S              *C2S   `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType          *int32  `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg           *string `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode          *int32  `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C              *S2C    `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
	proto.RegisterType((*C2S)(nil), "Qot_ModifyUserSecurity.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_ModifyUserSecurity.S2C")
	proto.RegisterType((*Request)(nil), "Qot_ModifyUserSecurity.Request")
	proto.RegisterType((*Response)(nil), "Qot_ModifyUserSecurity.Response")
	proto.RegisterEnum("Qot_ModifyUserSecurity.ModifyUserSecurityOp", ModifyUserSecurityOp_name, ModifyUserSecurityOp_value)
}

func init() { proto.RegisterFile("Qot_ModifyUserSecurity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6b, 0xd5, 0x40,
	0x14, 0xc5, 0x4d, 0xd2, 0xf8, 0x7c, 0xb7, 0x45, 0xc2, 0x58, 0xca, 0x50, 0x83, 0x84, 0xac, 0x82,
	0xd0, 0xa4, 0x0c, 0x2e, 0x8a, 0x2b, 0x35, 0x2e, 0x7d, 0x16, 0x27, 0x16, 0xc1, 0x4d, 0x79, 0x4d,
	0x6e, 0xd3, 0x60, 0x93, 0x3b, 0xce, 0x1f, 0x25, 0x9f, 0xc1, 0xbd, 0x9f, 0x57, 0xd2, 0xf6, 0x61,
	0xc5, 0xbc, 0xae, 0x86, 0x7b, 0xce, 0x8f, 0x7b, 0xcf, 0x81, 0x81, 0xf8, 0x13, 0xd9, 0xf3, 0x15,
	0x35, 0xdd, 0xe5, 0x78, 0x66, 0x50, 0x57, 0x58, 0x3b, 0xdd, 0xd9, 0x31, 0x57, 0x9a, 0x2c, 0xb1,
	0x83, 0x79, 0xf7, 0x70, 0xaf, 0xa4, 0xbe, 0xa7, 0xe1, 0x96, 0x3a, 0x8c, 0x26, 0xea, 0xbe, 0x92,
	0xf6, 0x10, 0x94, 0xa2, 0x62, 0x31, 0x2c, 0x5b, 0x4d, 0x4e, 0x7d, 0x5c, 0xf7, 0xc8, 0xbd, 0xc4,
	0xcf, 0x96, 0xf2, 0xaf, 0xc0, 0x9e, 0x82, 0x4f, 0x8a, 0xfb, 0x89, 0x9f, 0x85, 0xd2, 0x27, 0xc5,
	0x4e, 0x60, 0xcf, 0xdc, 0x1d, 0xf8, 0xd0, 0x19, 0xcb, 0x83, 0x24, 0xc8, 0x76, 0xc5, 0x7e, 0x7e,
	0x6f, 0xfb, 0x26, 0x80, 0xfc, 0x87, 0x4c, 0x43, 0x08, 0x2a, 0x51, 0xa6, 0x27, 0xb0, 0x90, 0xf8,
	0xdd, 0xa1, 0xb1, 0xec, 0x08, 0x82, 0x5a, 0x98, 0x9b, 0x9b, 0xbb, 0xe2, 0x79, 0xbe, 0xa5, 0x64,
	0x29, 0x2a, 0x39, 0x71, 0xe9, 0x2f, 0x0f, 0x9e, 0x48, 0x34, 0x8a, 0x06, 0x83, 0xec, 0x05, 0x2c,
	0x34, 0xda, 0xcf, 0xa3, 0xba, 0xcd, 0x1c, 0xbe, 0xde, 0x39, 0x7a, 0x75, 0x7c, 0x2c, 0x37, 0x22,
	0x3b, 0x80, 0xc7, 0x1a, 0xed, 0xca, 0xb4, 0xdc, 0x4f, 0xbc, 0x6c, 0x29, 0xef, 0x26, 0xc6, 0x61,
	0x81, 0x5a, 0x97, 0xd4, 0x20, 0x0f, 0x12, 0x2f, 0x0b, 0xe5, 0x66, 0x9c, 0xd2, 0x18, 0x51, 0xf3,
	0x9d, 0xc4, 0x7b, 0x28, 0x4d, 0x25, 0x4a, 0x39, 0x71, 0x2f, 0x7f, 0x7b, 0xb0, 0xff, 0xbf, 0x7f,
	0xaa, 0x58, 0x02, 0xf1, 0x9c, 0x7e, 0x7e, 0x36, 0x7c, 0x1b, 0xe8, 0xe7, 0x10, 0x3d, 0x62, 0x31,
	0xf0, 0x59, 0xe2, 0x6d, 0xd3, 0x44, 0xde, 0x56, 0xf7, 0x3d, 0x5e, 0x47, 0xfe, 0xd6, 0xed, 0x2b,
	0xfa, 0x81, 0xa7, 0xce, 0x46, 0xc1, 0xbb, 0x2f, 0xf0, 0xac, 0xa6, 0x3e, 0xbf, 0x74, 0xd6, 0xe5,
	0xa4, 0x70, 0x58, 0xab, 0x2e, 0x57, 0x17, 0x5f, 0xdf, 0xb4, 0x9d, 0xbd, 0x72, 0x17, 0x79, 0x4d,
	0x7d, 0x61, 0x2c, 0xaa, 0x2b, 0x1c, 0xae, 0x47, 0x57, 0xb4, 0x34, 0x81, 0x6b, 0xd5, 0x15, 0xd3,
	0x7b, 0xf3, 0x31, 0x8a, 0xf9, 0xea, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x45, 0xd2, 0x78, 0x6b,
	0x86, 0x02, 0x00, 0x00,
}