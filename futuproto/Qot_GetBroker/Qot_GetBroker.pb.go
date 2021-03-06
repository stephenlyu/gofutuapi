// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetBroker.proto

/*
Package Qot_GetBroker is a generated protocol buffer package.

It is generated from these files:
	Qot_GetBroker.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package Qot_GetBroker

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
	Security         *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

type S2C struct {
	Security         *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	BrokerAskList    []*Qot_Common.Broker `protobuf:"bytes,2,rep,name=brokerAskList" json:"brokerAskList,omitempty"`
	BrokerBidList    []*Qot_Common.Broker `protobuf:"bytes,3,rep,name=brokerBidList" json:"brokerBidList,omitempty"`
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

func (m *S2C) GetBrokerAskList() []*Qot_Common.Broker {
	if m != nil {
		return m.BrokerAskList
	}
	return nil
}

func (m *S2C) GetBrokerBidList() []*Qot_Common.Broker {
	if m != nil {
		return m.BrokerBidList
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
	proto.RegisterType((*C2S)(nil), "Qot_GetBroker.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetBroker.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetBroker.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetBroker.Response")
}

func init() { proto.RegisterFile("Qot_GetBroker.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0x49, 0xd3, 0x7e, 0xed, 0x37, 0xb5, 0x20, 0x5b, 0x91, 0xd0, 0x83, 0x84, 0xe0, 0x21,
	0x17, 0x93, 0xb2, 0x08, 0x15, 0x6f, 0x36, 0x07, 0x2f, 0x2a, 0xb8, 0xf1, 0xe4, 0x45, 0xda, 0x74,
	0x6c, 0x43, 0x4d, 0x66, 0xdd, 0xdd, 0x1c, 0x7a, 0xf5, 0xef, 0xf8, 0x27, 0x25, 0x69, 0x6a, 0x1b,
	0x10, 0xc4, 0xd3, 0x32, 0xf3, 0xbe, 0xcf, 0xf2, 0xb0, 0x0b, 0xc3, 0x47, 0x32, 0x2f, 0xb7, 0x68,
	0xa6, 0x8a, 0xd6, 0xa8, 0x02, 0xa9, 0xc8, 0x10, 0x1b, 0x34, 0x96, 0xa3, 0xa3, 0x88, 0xb2, 0x8c,
	0xf2, 0x6d, 0x38, 0x3a, 0x2e, 0xc3, 0xc3, 0x8d, 0x37, 0x01, 0x3b, 0xe2, 0x31, 0x1b, 0x43, 0x4f,
	0x63, 0x52, 0xa8, 0xd4, 0x6c, 0x1c, 0xcb, 0x6d, 0xf9, 0x7d, 0x7e, 0x12, 0x1c, 0x74, 0xe3, 0x3a,
	0x13, 0xdf, 0x2d, 0xef, 0xd3, 0x02, 0x3b, 0xe6, 0xd1, 0xdf, 0x49, 0x76, 0x05, 0x83, 0x79, 0x25,
	0x77, 0xa3, 0xd7, 0x77, 0xa9, 0x36, 0x4e, 0xcb, 0xb5, 0xfd, 0x3e, 0x67, 0x87, 0xd8, 0xd6, 0x5e,
	0x34, 0x8b, 0x7b, 0x72, 0x9a, 0x2e, 0x2a, 0xd2, 0xfe, 0x8d, 0xac, 0x8b, 0x5e, 0x08, 0x5d, 0x81,
	0xef, 0x05, 0x6a, 0xc3, 0xce, 0xc1, 0x4e, 0xb8, 0xae, 0x5d, 0xb7, 0xe8, 0xfe, 0x0d, 0x23, 0x1e,
	0x8b, 0x32, 0xf6, 0x3e, 0x2c, 0xe8, 0x09, 0xd4, 0x92, 0x72, 0x8d, 0xec, 0x0c, 0xba, 0x0a, 0xcd,
	0xd3, 0x46, 0x62, 0x85, 0x75, 0xae, 0xdb, 0x17, 0x97, 0xe3, 0xb1, 0xd8, 0x2d, 0xd9, 0x29, 0xfc,
	0x53, 0x68, 0xee, 0xf5, 0xd2, 0x69, 0xb9, 0x96, 0xff, 0x5f, 0xd4, 0x13, 0x73, 0xa0, 0x8b, 0x4a,
	0x45, 0xb4, 0x40, 0xc7, 0x76, 0x2d, 0xbf, 0x23, 0x76, 0x63, 0x29, 0xa1, 0x79, 0xe2, 0xb4, 0x5d,
	0xeb, 0x07, 0x89, 0x98, 0x47, 0xa2, 0x8c, 0xa7, 0x0f, 0x30, 0x4c, 0x28, 0x0b, 0x5e, 0x0b, 0x53,
	0x04, 0x24, 0x31, 0x9f, 0xc9, 0x34, 0x90, 0xf3, 0xe7, 0xc9, 0x32, 0x35, 0xab, 0x62, 0x1e, 0x24,
	0x94, 0x85, 0xda, 0xa0, 0x5c, 0x61, 0xfe, 0xb6, 0x29, 0xc2, 0x25, 0x95, 0xc5, 0x99, 0x4c, 0xc3,
	0xf2, 0xac, 0xbe, 0x37, 0x6c, 0x5c, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x31, 0x45, 0x7d, 0x09,
	0x31, 0x02, 0x00, 0x00,
}
