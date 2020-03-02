// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_UpdateKL.proto

/*
Package Qot_UpdateKL is a generated protocol buffer package.

It is generated from these files:
	Qot_UpdateKL.proto

It has these top-level messages:
	S2C
	Response
*/
package Qot_UpdateKL

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
	RehabType        *int32               `protobuf:"varint,1,req,name=rehabType" json:"rehabType,omitempty"`
	KlType           *int32               `protobuf:"varint,2,req,name=klType" json:"klType,omitempty"`
	Security         *Qot_Common.Security `protobuf:"bytes,3,req,name=security" json:"security,omitempty"`
	KlList           []*Qot_Common.KLine  `protobuf:"bytes,4,rep,name=klList" json:"klList,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *S2C) GetRehabType() int32 {
	if m != nil && m.RehabType != nil {
		return *m.RehabType
	}
	return 0
}

func (m *S2C) GetKlType() int32 {
	if m != nil && m.KlType != nil {
		return *m.KlType
	}
	return 0
}

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
	proto.RegisterType((*S2C)(nil), "Qot_UpdateKL.S2C")
	proto.RegisterType((*Response)(nil), "Qot_UpdateKL.Response")
}

func init() { proto.RegisterFile("Qot_UpdateKL.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4f, 0xc2, 0x30,
	0x1c, 0xc5, 0xd3, 0x0d, 0x04, 0x0a, 0x07, 0xad, 0xc6, 0x34, 0xc4, 0x98, 0x06, 0x2f, 0xf3, 0x60,
	0x47, 0x16, 0xe3, 0xc1, 0xa3, 0x3b, 0x02, 0x07, 0x8b, 0x5e, 0xbc, 0x98, 0x31, 0xfe, 0x42, 0x23,
	0x5b, 0x9b, 0xb6, 0x3b, 0x70, 0xf4, 0x5b, 0xf8, 0x71, 0x4d, 0x71, 0x8c, 0x9d, 0x9a, 0xf7, 0x7e,
	0x2f, 0xff, 0xbe, 0x3c, 0x4c, 0x5e, 0x95, 0xfb, 0x7c, 0xd7, 0xeb, 0xcc, 0xc1, 0x6c, 0xce, 0xb5,
	0x51, 0x4e, 0x91, 0x51, 0xdb, 0x1b, 0x8f, 0x52, 0x55, 0x14, 0xaa, 0xfc, 0x67, 0xe3, 0x73, 0xcf,
	0xda, 0xce, 0xe4, 0x17, 0xe1, 0x70, 0x99, 0xa4, 0xe4, 0x06, 0x0f, 0x0c, 0x6c, 0xb3, 0xd5, 0xdb,
	0x5e, 0x03, 0x45, 0x2c, 0x88, 0xba, 0xe2, 0x64, 0x90, 0x6b, 0x7c, 0xf6, 0xbd, 0x3b, 0xa0, 0xe0,
	0x80, 0x6a, 0x45, 0xa6, 0xb8, 0x6f, 0x21, 0xaf, 0x8c, 0x74, 0x7b, 0x1a, 0xb2, 0x20, 0x1a, 0x26,
	0x57, 0xbc, 0xf5, 0xc5, 0xb2, 0x66, 0xa2, 0x49, 0x91, 0x7b, 0x7f, 0x69, 0x2e, 0xad, 0xa3, 0x1d,
	0x16, 0x46, 0xc3, 0xe4, 0xa2, 0x9d, 0x9f, 0xcd, 0x65, 0x09, 0xa2, 0x0e, 0x4c, 0x7e, 0x10, 0xee,
	0x0b, 0xb0, 0x5a, 0x95, 0x16, 0xc8, 0x2d, 0xee, 0x19, 0x70, 0xa7, 0x76, 0xcf, 0x9d, 0x87, 0xc7,
	0xe9, 0x54, 0x1c, 0x4d, 0xdf, 0xd0, 0x80, 0x5b, 0xd8, 0x0d, 0x0d, 0x18, 0x8a, 0x06, 0xa2, 0x56,
	0x84, 0xe2, 0x1e, 0x18, 0x93, 0xaa, 0x35, 0xd0, 0x90, 0xa1, 0xa8, 0x2b, 0x8e, 0x92, 0xdc, 0xe1,
	0xd0, 0x26, 0x39, 0xed, 0x30, 0xd4, 0xd4, 0x68, 0x96, 0x5c, 0x26, 0xa9, 0xf0, 0xf4, 0x65, 0x81,
	0x2f, 0x73, 0x55, 0xf0, 0xaf, 0xca, 0x55, 0x5c, 0x69, 0x28, 0x33, 0x2d, 0xb9, 0x5e, 0x7d, 0x3c,
	0x6d, 0xa4, 0xdb, 0x56, 0x2b, 0x9e, 0xab, 0x22, 0xb6, 0x0e, 0xf4, 0x16, 0xca, 0xdd, 0xbe, 0x8a,
	0x37, 0xca, 0x07, 0x33, 0x2d, 0x63, 0xff, 0x1e, 0x06, 0x8e, 0xdb, 0x77, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x5a, 0x9b, 0x16, 0xd7, 0xb0, 0x01, 0x00, 0x00,
}
