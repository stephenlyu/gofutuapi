// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetKL.proto

/*
Package Qot_GetKL is a generated protocol buffer package.

It is generated from these files:
	Qot_GetKL.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package Qot_GetKL

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

type C2S struct {
	RehabType        *int32               `protobuf:"varint,1,req,name=rehabType" json:"rehabType,omitempty"`
	KlType           *int32               `protobuf:"varint,2,req,name=klType" json:"klType,omitempty"`
	Security         *Qot_Common.Security `protobuf:"bytes,3,req,name=security" json:"security,omitempty"`
	ReqNum           *int32               `protobuf:"varint,4,req,name=reqNum" json:"reqNum,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S) GetRehabType() int32 {
	if m != nil && m.RehabType != nil {
		return *m.RehabType
	}
	return 0
}

func (m *C2S) GetKlType() int32 {
	if m != nil && m.KlType != nil {
		return *m.KlType
	}
	return 0
}

func (m *C2S) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *C2S) GetReqNum() int32 {
	if m != nil && m.ReqNum != nil {
		return *m.ReqNum
	}
	return 0
}

type S2C struct {
	Security         *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	KlList           []*Qot_Common.KLine  `protobuf:"bytes,2,rep,name=klList" json:"klList,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2C) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetKlList() []*Qot_Common.KLine {
	if m != nil {
		return m.KlList
	}
	return nil
}

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
	proto.RegisterType((*C2S)(nil), "Qot_GetKL.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetKL.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetKL.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetKL.Response")
}

func init() { proto.RegisterFile("Qot_GetKL.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4f, 0xb3, 0x40,
	0x18, 0xc4, 0x03, 0xb4, 0x6f, 0xdb, 0xed, 0x1b, 0x3f, 0xd0, 0x98, 0x4d, 0x63, 0x0c, 0xe1, 0x84,
	0x31, 0x2e, 0xcd, 0xea, 0xc9, 0xa3, 0x1c, 0x4c, 0x6c, 0x35, 0x71, 0xf1, 0xe4, 0xc5, 0x14, 0x7c,
	0x6c, 0x89, 0x85, 0xdd, 0xee, 0xc7, 0xa1, 0x17, 0x6f, 0xfe, 0xdf, 0x66, 0x0b, 0xfd, 0xd0, 0x93,
	0x27, 0x32, 0xf3, 0x9b, 0xcc, 0xc0, 0x03, 0xda, 0x7f, 0xe2, 0xfa, 0xf5, 0x0e, 0xf4, 0x68, 0x4c,
	0x84, 0xe4, 0x9a, 0xfb, 0xbd, 0x8d, 0x31, 0xf8, 0x9f, 0xf0, 0xb2, 0xe4, 0x55, 0x0d, 0x06, 0x07,
	0x16, 0xec, 0x3a, 0xe1, 0x97, 0x83, 0xbc, 0x84, 0xa6, 0xfe, 0x29, 0xea, 0x49, 0x98, 0x4d, 0xb2,
	0xe7, 0xa5, 0x00, 0xec, 0x04, 0x6e, 0xd4, 0x66, 0x5b, 0xc3, 0x3f, 0x41, 0xff, 0x3e, 0xe6, 0x2b,
	0xe4, 0xae, 0x50, 0xa3, 0xfc, 0x21, 0xea, 0x2a, 0xc8, 0x8d, 0x2c, 0xf4, 0x12, 0x7b, 0x81, 0x1b,
	0xf5, 0xe9, 0x31, 0xd9, 0x99, 0x48, 0x1b, 0xc6, 0x36, 0x29, 0xdb, 0x24, 0x61, 0xf1, 0x68, 0x4a,
	0xdc, 0xaa, 0x9b, 0x6a, 0x15, 0x66, 0xc8, 0x4b, 0x69, 0xf2, 0xa3, 0xd0, 0xf9, 0x53, 0xe1, 0xb9,
	0x7d, 0xb5, 0x71, 0xa1, 0x34, 0x76, 0x03, 0x2f, 0xea, 0xd3, 0xc3, 0xdd, 0xfc, 0x68, 0x5c, 0x54,
	0xc0, 0x9a, 0x40, 0x78, 0x81, 0x3a, 0x0c, 0x16, 0x06, 0x94, 0xf6, 0x03, 0xe4, 0xe5, 0x54, 0x35,
	0x13, 0x7b, 0x64, 0x7b, 0xc0, 0x84, 0xa6, 0xcc, 0xa2, 0xf0, 0x13, 0x75, 0x19, 0x28, 0xc1, 0x2b,
	0x05, 0xfe, 0x19, 0xea, 0x48, 0xd0, 0xdb, 0xd3, 0xdc, 0xb4, 0x2e, 0xaf, 0x87, 0x43, 0xb6, 0x36,
	0xeb, 0x8f, 0xd2, 0x0f, 0x6a, 0x8a, 0xdd, 0xc0, 0x89, 0x7a, 0xac, 0x51, 0x3e, 0x46, 0x1d, 0x90,
	0x32, 0xe1, 0x6f, 0x80, 0xbd, 0xc0, 0x89, 0xda, 0x6c, 0x2d, 0xed, 0xbe, 0xa2, 0x39, 0x6e, 0x05,
	0xce, 0xaf, 0xfd, 0x94, 0x26, 0xcc, 0xa2, 0xdb, 0x7b, 0x74, 0x94, 0xf3, 0x92, 0xbc, 0x1b, 0x6d,
	0x08, 0x17, 0x50, 0x4d, 0x44, 0x41, 0x44, 0xf6, 0x72, 0x35, 0x2d, 0xf4, 0xcc, 0x64, 0x24, 0xe7,
	0x65, 0xac, 0x34, 0x88, 0x19, 0x54, 0xf3, 0xa5, 0x89, 0xa7, 0xdc, 0x06, 0x27, 0xa2, 0x88, 0xed,
	0x73, 0xf5, 0x6b, 0xe3, 0x4d, 0xe9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x62, 0x20, 0x79,
	0x21, 0x02, 0x00, 0x00,
}
