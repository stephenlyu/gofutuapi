// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_RequestHistoryKL.proto

/*
Package Qot_RequestHistoryKL is a generated protocol buffer package.

It is generated from these files:
	Qot_RequestHistoryKL.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package Qot_RequestHistoryKL

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
	BeginTime        *string              `protobuf:"bytes,4,req,name=beginTime" json:"beginTime,omitempty"`
	EndTime          *string              `protobuf:"bytes,5,req,name=endTime" json:"endTime,omitempty"`
	MaxAckKLNum      *int32               `protobuf:"varint,6,opt,name=maxAckKLNum" json:"maxAckKLNum,omitempty"`
	NeedKLFieldsFlag *int64               `protobuf:"varint,7,opt,name=needKLFieldsFlag" json:"needKLFieldsFlag,omitempty"`
	NextReqKey       []byte               `protobuf:"bytes,8,opt,name=nextReqKey" json:"nextReqKey,omitempty"`
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

func (m *C2S) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *C2S) GetEndTime() string {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return ""
}

func (m *C2S) GetMaxAckKLNum() int32 {
	if m != nil && m.MaxAckKLNum != nil {
		return *m.MaxAckKLNum
	}
	return 0
}

func (m *C2S) GetNeedKLFieldsFlag() int64 {
	if m != nil && m.NeedKLFieldsFlag != nil {
		return *m.NeedKLFieldsFlag
	}
	return 0
}

func (m *C2S) GetNextReqKey() []byte {
	if m != nil {
		return m.NextReqKey
	}
	return nil
}

type S2C struct {
	Security         *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	KlList           []*Qot_Common.KLine  `protobuf:"bytes,2,rep,name=klList" json:"klList,omitempty"`
	NextReqKey       []byte               `protobuf:"bytes,3,opt,name=nextReqKey" json:"nextReqKey,omitempty"`
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

func (m *S2C) GetNextReqKey() []byte {
	if m != nil {
		return m.NextReqKey
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
	proto.RegisterType((*C2S)(nil), "Qot_RequestHistoryKL.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_RequestHistoryKL.S2C")
	proto.RegisterType((*Request)(nil), "Qot_RequestHistoryKL.Request")
	proto.RegisterType((*Response)(nil), "Qot_RequestHistoryKL.Response")
}

func init() { proto.RegisterFile("Qot_RequestHistoryKL.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x8b, 0xd3, 0x40,
	0x18, 0xc5, 0x99, 0xa4, 0xdd, 0xb6, 0xd3, 0x3d, 0xac, 0xe3, 0x22, 0x63, 0x91, 0x65, 0xe8, 0x29,
	0x2a, 0xa6, 0x25, 0x88, 0x07, 0x0f, 0x82, 0x06, 0x16, 0x21, 0x51, 0x70, 0xb2, 0x5e, 0xbc, 0x48,
	0x9a, 0x7e, 0xa6, 0xc3, 0x26, 0x33, 0xd9, 0xcc, 0x04, 0x36, 0x57, 0x6f, 0x9e, 0xfd, 0x87, 0x65,
	0x62, 0x6a, 0xa3, 0x5b, 0x61, 0x4f, 0xe1, 0x7b, 0xef, 0x47, 0xe6, 0xcd, 0xbc, 0x0f, 0x2f, 0x3e,
	0x29, 0xf3, 0x95, 0xc3, 0x4d, 0x03, 0xda, 0xbc, 0x17, 0xda, 0xa8, 0xba, 0x8d, 0x62, 0xbf, 0xaa,
	0x95, 0x51, 0xe4, 0xfc, 0x98, 0xb7, 0x38, 0x0d, 0x55, 0x59, 0x2a, 0xf9, 0x9b, 0x59, 0x9c, 0x59,
	0x66, 0xa8, 0x2c, 0x7f, 0x3a, 0xd8, 0x0d, 0x83, 0x84, 0x3c, 0xc1, 0xb3, 0x1a, 0x76, 0xe9, 0xe6,
	0xaa, 0xad, 0x80, 0x22, 0xe6, 0x78, 0x63, 0x7e, 0x10, 0xc8, 0x23, 0x7c, 0x72, 0x5d, 0x74, 0x96,
	0xd3, 0x59, 0xfd, 0x44, 0xd6, 0x78, 0xaa, 0x21, 0x6b, 0x6a, 0x61, 0x5a, 0xea, 0x32, 0xc7, 0x9b,
	0x07, 0xe7, 0xfe, 0xe0, 0x88, 0xa4, 0xf7, 0xf8, 0x1f, 0xca, 0x9e, 0xb3, 0x81, 0x5c, 0xc8, 0x2b,
	0x51, 0x02, 0x1d, 0x31, 0xc7, 0x9b, 0xf1, 0x83, 0x40, 0x28, 0x9e, 0x80, 0xdc, 0x76, 0xde, 0xb8,
	0xf3, 0xf6, 0x23, 0x61, 0x78, 0x5e, 0xa6, 0xb7, 0x6f, 0xb3, 0xeb, 0x28, 0xfe, 0xd8, 0x94, 0xf4,
	0x84, 0x21, 0x6f, 0xcc, 0x87, 0x12, 0x79, 0x86, 0xcf, 0x24, 0xc0, 0x36, 0x8a, 0x2f, 0x05, 0x14,
	0x5b, 0x7d, 0x59, 0xa4, 0x39, 0x9d, 0x30, 0xe4, 0xb9, 0xfc, 0x8e, 0x4e, 0x2e, 0x30, 0x96, 0x70,
	0x6b, 0x38, 0xdc, 0x44, 0xd0, 0xd2, 0x29, 0x43, 0xde, 0x29, 0x1f, 0x28, 0xcb, 0xef, 0x08, 0xbb,
	0x49, 0x10, 0xfe, 0x75, 0x3f, 0x74, 0xaf, 0xfb, 0x3d, 0xb5, 0x2f, 0x15, 0x0b, 0x6d, 0xa8, 0xc3,
	0x5c, 0x6f, 0x1e, 0x3c, 0x18, 0xf2, 0x51, 0x2c, 0x24, 0xf0, 0x1e, 0xf8, 0x27, 0x84, 0x7b, 0x27,
	0xc4, 0x2b, 0x3c, 0xe9, 0xeb, 0x24, 0xcf, 0xb1, 0x9b, 0x05, 0xba, 0x8f, 0xf0, 0xd8, 0x3f, 0xba,
	0x05, 0x61, 0x90, 0x70, 0x4b, 0x2d, 0x7f, 0x20, 0x3c, 0xe5, 0xa0, 0x2b, 0x25, 0x35, 0x90, 0x0b,
	0x3c, 0xa9, 0xc1, 0x1c, 0x5a, 0x7d, 0x3d, 0x7a, 0xf1, 0x72, 0xbd, 0xe6, 0x7b, 0xd1, 0x36, 0x5b,
	0x83, 0xf9, 0xa0, 0x73, 0xea, 0x30, 0xe4, 0xcd, 0x78, 0x3f, 0x75, 0x4d, 0xd4, 0x75, 0xa8, 0xb6,
	0xd0, 0x25, 0x1b, 0xf3, 0xfd, 0x68, 0xb3, 0xe8, 0x20, 0xa3, 0x23, 0x86, 0xfe, 0x9f, 0x25, 0x09,
	0x42, 0x6e, 0xa9, 0x77, 0x9f, 0xf1, 0xc3, 0x4c, 0x95, 0xfe, 0xb7, 0xc6, 0x34, 0xbe, 0xaa, 0x40,
	0xa6, 0x95, 0xf0, 0xab, 0xcd, 0x97, 0x37, 0xb9, 0x30, 0xbb, 0x66, 0xe3, 0x67, 0xaa, 0x5c, 0x69,
	0x03, 0xd5, 0x0e, 0x64, 0xd1, 0x36, 0xab, 0x5c, 0x59, 0x30, 0xad, 0xc4, 0xca, 0x7e, 0xbb, 0x05,
	0x5d, 0x1d, 0xfb, 0xff, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0x34, 0xc2, 0x54, 0x08, 0x03,
	0x00, 0x00,
}
