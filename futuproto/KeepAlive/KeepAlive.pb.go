// Code generated by protoc-gen-go. DO NOT EDIT.
// source: KeepAlive.proto

/*
Package KeepAlive is a generated protocol buffer package.

It is generated from these files:
	KeepAlive.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package KeepAlive

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/stephenlyu/gofutuapi/futuproto/Common"

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
	Time             *int64 `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type S2C struct {
	Time             *int64 `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2C) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
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
	proto.RegisterType((*C2S)(nil), "KeepAlive.C2S")
	proto.RegisterType((*S2C)(nil), "KeepAlive.S2C")
	proto.RegisterType((*Request)(nil), "KeepAlive.Request")
	proto.RegisterType((*Response)(nil), "KeepAlive.Response")
}

func init() { proto.RegisterFile("KeepAlive.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x93, 0x1a, 0x3b, 0x8a, 0xc2, 0x0a, 0x12, 0x3d, 0x48, 0xe8, 0x29, 0x20, 0x6e,
	0xca, 0xea, 0xc9, 0x9b, 0xee, 0x4d, 0xf1, 0xb2, 0xf1, 0xe4, 0xad, 0x8d, 0x63, 0x1a, 0x68, 0x32,
	0xeb, 0xfe, 0x11, 0x7a, 0xf1, 0xb3, 0xcb, 0xa6, 0xb6, 0x82, 0xf4, 0x34, 0xf3, 0xde, 0x6f, 0x78,
	0x0f, 0x06, 0x4e, 0x9f, 0x11, 0xf5, 0xc3, 0xaa, 0xfd, 0x42, 0xae, 0x0d, 0x39, 0x62, 0x93, 0x9d,
	0x71, 0x79, 0x2c, 0xa9, 0xeb, 0xa8, 0xdf, 0x80, 0xe9, 0x05, 0xc4, 0x52, 0x54, 0x8c, 0x41, 0xe2,
	0xda, 0x0e, 0xb3, 0x28, 0x1f, 0x15, 0xb1, 0x1a, 0xf6, 0x80, 0x2a, 0x21, 0xf7, 0xa2, 0x6b, 0x48,
	0x15, 0x7e, 0x7a, 0xb4, 0x8e, 0xe5, 0x10, 0xd7, 0xc2, 0x0e, 0xf4, 0x48, 0x9c, 0xf0, 0xbf, 0x62,
	0x29, 0x2a, 0x15, 0xd0, 0xf4, 0x1b, 0x0e, 0x15, 0x5a, 0x4d, 0xbd, 0x45, 0x76, 0x05, 0xa9, 0x41,
	0xf7, 0xba, 0xd6, 0x9b, 0xbc, 0xf1, 0x7d, 0x72, 0x73, 0x37, 0x9b, 0xa9, 0xad, 0xc9, 0xce, 0xe1,
	0xc0, 0xa0, 0x7b, 0xb1, 0x4d, 0x36, 0xca, 0xa3, 0x62, 0xa2, 0x7e, 0x15, 0xcb, 0x20, 0x45, 0x63,
	0x24, 0xbd, 0x63, 0x16, 0xe7, 0x51, 0x31, 0x56, 0x5b, 0x19, 0xfa, 0xad, 0xa8, 0xb3, 0x24, 0x8f,
	0xfe, 0xf5, 0x57, 0x42, 0xaa, 0x80, 0x1e, 0x9f, 0xe0, 0xac, 0xa6, 0x8e, 0x7f, 0x78, 0xe7, 0x39,
	0x69, 0xec, 0xe7, 0xba, 0xe5, 0x7a, 0xf1, 0x76, 0xdb, 0xb4, 0x6e, 0xe9, 0x17, 0xbc, 0xa6, 0xae,
	0xb4, 0x0e, 0xf5, 0x12, 0xfb, 0xd5, 0xda, 0x97, 0x0d, 0x85, 0xc3, 0xb9, 0x6e, 0xcb, 0x30, 0x87,
	0x27, 0x95, 0xbb, 0xd0, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x9b, 0x2a, 0x83, 0x59, 0x01,
	0x00, 0x00,
}
