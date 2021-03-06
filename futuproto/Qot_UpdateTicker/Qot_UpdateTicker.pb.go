// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_UpdateTicker.proto

/*
Package Qot_UpdateTicker is a generated protocol buffer package.

It is generated from these files:
	Qot_UpdateTicker.proto

It has these top-level messages:
	S2C
	Response
*/
package Qot_UpdateTicker

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

type S2C struct {
	Security         *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	TickerList       []*Qot_Common.Ticker `protobuf:"bytes,2,rep,name=tickerList" json:"tickerList,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *S2C) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetTickerList() []*Qot_Common.Ticker {
	if m != nil {
		return m.TickerList
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
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
	proto.RegisterType((*S2C)(nil), "Qot_UpdateTicker.S2C")
	proto.RegisterType((*Response)(nil), "Qot_UpdateTicker.Response")
}

func init() { proto.RegisterFile("Qot_UpdateTicker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x31, 0x4f, 0xf3, 0x30,
	0x18, 0x84, 0x95, 0xa6, 0xfd, 0xda, 0xcf, 0x61, 0x40, 0x06, 0x2a, 0xab, 0x03, 0x8a, 0xba, 0x90,
	0x05, 0x27, 0xb2, 0x98, 0x3a, 0x92, 0x15, 0x86, 0x3a, 0x65, 0x61, 0x41, 0x69, 0xfa, 0x92, 0x5a,
	0x25, 0xb1, 0x65, 0xbf, 0x19, 0xf2, 0x03, 0xf8, 0xdf, 0x28, 0x69, 0x8b, 0xa2, 0x32, 0x59, 0x77,
	0xf7, 0xd8, 0x67, 0x1d, 0x99, 0xaf, 0x35, 0x7e, 0xbc, 0x99, 0x5d, 0x8e, 0xb0, 0x51, 0xc5, 0x01,
	0x2c, 0x37, 0x56, 0xa3, 0xa6, 0xd7, 0x97, 0xfe, 0xe2, 0x2a, 0xd5, 0x55, 0xa5, 0xeb, 0x63, 0xbe,
	0xe8, 0xf3, 0xa1, 0xb3, 0x3c, 0x10, 0x3f, 0x13, 0x29, 0x4d, 0xc8, 0xcc, 0x41, 0xd1, 0x58, 0x85,
	0x2d, 0xf3, 0xc2, 0x51, 0x14, 0x88, 0x5b, 0x3e, 0x60, 0xb3, 0x53, 0x26, 0x7f, 0x29, 0x2a, 0x08,
	0xc1, 0xbe, 0xe2, 0x45, 0x39, 0x64, 0xa3, 0xd0, 0x8f, 0x02, 0x41, 0x87, 0x77, 0x8e, 0x1f, 0x90,
	0x03, 0x6a, 0xf9, 0xed, 0x91, 0x99, 0x04, 0x67, 0x74, 0xed, 0x80, 0xde, 0x93, 0xa9, 0x05, 0xdc,
	0xb4, 0x06, 0xfa, 0xc6, 0xc9, 0x6a, 0xfc, 0xf8, 0x94, 0x24, 0xf2, 0x6c, 0xd2, 0x39, 0xf9, 0x67,
	0x01, 0x5f, 0x5d, 0xc9, 0x46, 0xa1, 0x17, 0xfd, 0x97, 0x27, 0x45, 0x19, 0x99, 0x82, 0xb5, 0xa9,
	0xde, 0x01, 0xf3, 0x43, 0x2f, 0x9a, 0xc8, 0xb3, 0xa4, 0x0f, 0xc4, 0x77, 0xa2, 0x60, 0xe3, 0xd0,
	0x8b, 0x02, 0x71, 0xc7, 0xff, 0x6c, 0x94, 0x89, 0x54, 0x76, 0xc4, 0xf3, 0x9a, 0xdc, 0x14, 0xba,
	0xe2, 0x9f, 0x0d, 0x36, 0x5c, 0x1b, 0xa8, 0x73, 0xa3, 0xb8, 0xd9, 0xbe, 0xaf, 0x4a, 0x85, 0xfb,
	0x66, 0xcb, 0x0b, 0x5d, 0xc5, 0x0e, 0xc1, 0xec, 0xa1, 0xfe, 0x6a, 0x9b, 0xb8, 0xd4, 0x1d, 0x98,
	0x1b, 0x15, 0x77, 0x67, 0x3f, 0x5b, 0x7c, 0xf9, 0xf6, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32,
	0x24, 0xdc, 0x72, 0x92, 0x01, 0x00, 0x00,
}
