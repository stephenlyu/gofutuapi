// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetHoldingChangeList.proto

/*
Package Qot_GetHoldingChangeList is a generated protocol buffer package.

It is generated from these files:
	Qot_GetHoldingChangeList.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package Qot_GetHoldingChangeList

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
	Security       *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	HolderCategory *int32               `protobuf:"varint,2,req,name=holderCategory" json:"holderCategory,omitempty"`
	// 以下是发布时间筛选，不传返回所有数据，传了返回发布时间属于开始时间到结束时间段内的数据
	BeginTime        *string `protobuf:"bytes,3,opt,name=beginTime" json:"beginTime,omitempty"`
	EndTime          *string `protobuf:"bytes,4,opt,name=endTime" json:"endTime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
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

func (m *C2S) GetHolderCategory() int32 {
	if m != nil && m.HolderCategory != nil {
		return *m.HolderCategory
	}
	return 0
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

type S2C struct {
	Security          *Qot_Common.Security             `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	HoldingChangeList []*Qot_Common.ShareHoldingChange `protobuf:"bytes,2,rep,name=holdingChangeList" json:"holdingChangeList,omitempty"`
	XXX_unrecognized  []byte                           `json:"-"`
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

func (m *S2C) GetHoldingChangeList() []*Qot_Common.ShareHoldingChange {
	if m != nil {
		return m.HoldingChangeList
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
	proto.RegisterType((*C2S)(nil), "Qot_GetHoldingChangeList.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetHoldingChangeList.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetHoldingChangeList.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetHoldingChangeList.Response")
}

func init() { proto.RegisterFile("Qot_GetHoldingChangeList.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x8f, 0xd3, 0x30,
	0x18, 0xc5, 0x95, 0xa4, 0xa5, 0xad, 0x8b, 0x10, 0x18, 0x84, 0xa2, 0x0a, 0xaa, 0x28, 0x03, 0xca,
	0x42, 0x52, 0x59, 0x4c, 0x1d, 0xeb, 0x01, 0x86, 0x32, 0xe0, 0x74, 0x81, 0x05, 0xa5, 0xc9, 0x47,
	0x62, 0xa9, 0xb1, 0x8d, 0xed, 0x0c, 0xf9, 0x03, 0x6e, 0xbc, 0xf9, 0xfe, 0xdd, 0x53, 0xd2, 0xf6,
	0xae, 0xed, 0xa9, 0x37, 0xdc, 0x14, 0x7d, 0xef, 0xfd, 0xbe, 0xf8, 0xe9, 0xd9, 0x68, 0xfe, 0x4b,
	0xda, 0xbf, 0xdf, 0xc1, 0xfe, 0x90, 0xbb, 0x82, 0x8b, 0x92, 0x56, 0x99, 0x28, 0x61, 0xcd, 0x8d,
	0x8d, 0x95, 0x96, 0x56, 0x62, 0xff, 0x9a, 0x3f, 0x7b, 0x4d, 0x65, 0x5d, 0x4b, 0xb1, 0xe7, 0x66,
	0x6f, 0x3b, 0xee, 0x54, 0x09, 0xef, 0x1c, 0xe4, 0x51, 0x92, 0xe2, 0x05, 0x1a, 0x1b, 0xc8, 0x1b,
	0xcd, 0x6d, 0xeb, 0x3b, 0x81, 0x1b, 0x4d, 0xc9, 0x87, 0xf8, 0x04, 0x4e, 0x0f, 0x1e, 0x7b, 0xa0,
	0xf0, 0x17, 0xf4, 0xa6, 0x92, 0xbb, 0x02, 0x34, 0xcd, 0x2c, 0x94, 0x52, 0xb7, 0xbe, 0x1b, 0xb8,
	0xd1, 0x90, 0x5d, 0xa8, 0xf8, 0x13, 0x9a, 0x6c, 0xa1, 0xe4, 0x62, 0xc3, 0x6b, 0xf0, 0xbd, 0xc0,
	0x89, 0x26, 0xec, 0x51, 0xc0, 0x3e, 0x1a, 0x81, 0x28, 0x7a, 0x6f, 0xd0, 0x7b, 0xc7, 0x31, 0xbc,
	0x71, 0x90, 0x97, 0x12, 0xfa, 0x82, 0x64, 0x6b, 0xf4, 0xae, 0xba, 0x2c, 0xc2, 0x77, 0x03, 0x2f,
	0x9a, 0x92, 0xf9, 0xd9, 0x6a, 0x95, 0x69, 0x38, 0xab, 0x8c, 0x3d, 0x5d, 0x0c, 0x97, 0x68, 0xc4,
	0xe0, 0x7f, 0x03, 0xc6, 0xe2, 0x04, 0x79, 0x39, 0x31, 0x87, 0x14, 0x9f, 0xe3, 0xab, 0x97, 0x42,
	0x49, 0xca, 0x3a, 0x32, 0xbc, 0x75, 0xd0, 0x98, 0x81, 0x51, 0x52, 0x18, 0xc0, 0x73, 0x34, 0xd2,
	0x60, 0x37, 0xad, 0x82, 0xfe, 0x0f, 0xc3, 0xe5, 0xe0, 0xeb, 0xb7, 0xc5, 0x82, 0x1d, 0x45, 0xfc,
	0x11, 0xbd, 0xd2, 0x60, 0x7f, 0x9a, 0xd2, 0x77, 0xfb, 0x26, 0x0e, 0x53, 0x5f, 0x91, 0xd6, 0x54,
	0x16, 0xfb, 0xfa, 0x86, 0xec, 0x38, 0x76, 0x79, 0x0c, 0xc9, 0xfb, 0xe2, 0x9e, 0xcd, 0x93, 0x12,
	0xca, 0x3a, 0x72, 0xf5, 0x1b, 0xbd, 0xcf, 0x65, 0x1d, 0xff, 0x6b, 0x6c, 0x13, 0x4b, 0x05, 0x22,
	0x53, 0x3c, 0x56, 0xdb, 0x3f, 0xab, 0x92, 0xdb, 0xaa, 0xd9, 0xc6, 0xb9, 0xac, 0x13, 0x63, 0x41,
	0x55, 0x20, 0x76, 0x6d, 0x93, 0x94, 0xb2, 0x03, 0x33, 0xc5, 0x93, 0xee, 0xdb, 0xbf, 0x97, 0xe4,
	0xda, 0x19, 0xf7, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xd4, 0x0a, 0xa5, 0xa3, 0x02, 0x00, 0x00,
}
