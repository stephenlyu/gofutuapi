// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetPlateSecurity.proto

/*
Package Qot_GetPlateSecurity is a generated protocol buffer package.

It is generated from these files:
	Qot_GetPlateSecurity.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package Qot_GetPlateSecurity

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
	Plate            *Qot_Common.Security `protobuf:"bytes,1,req,name=plate" json:"plate,omitempty"`
	SortField        *int32               `protobuf:"varint,2,opt,name=sortField" json:"sortField,omitempty"`
	Ascend           *bool                `protobuf:"varint,3,opt,name=ascend" json:"ascend,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S) GetPlate() *Qot_Common.Security {
	if m != nil {
		return m.Plate
	}
	return nil
}

func (m *C2S) GetSortField() int32 {
	if m != nil && m.SortField != nil {
		return *m.SortField
	}
	return 0
}

func (m *C2S) GetAscend() bool {
	if m != nil && m.Ascend != nil {
		return *m.Ascend
	}
	return false
}

type S2C struct {
	StaticInfoList   []*Qot_Common.SecurityStaticInfo `protobuf:"bytes,1,rep,name=staticInfoList" json:"staticInfoList,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2C) GetStaticInfoList() []*Qot_Common.SecurityStaticInfo {
	if m != nil {
		return m.StaticInfoList
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
	proto.RegisterType((*C2S)(nil), "Qot_GetPlateSecurity.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetPlateSecurity.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetPlateSecurity.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetPlateSecurity.Response")
}

func init() { proto.RegisterFile("Qot_GetPlateSecurity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x49, 0xd3, 0xbc, 0xb6, 0xd3, 0xc7, 0xe3, 0x31, 0x16, 0x89, 0x45, 0x4a, 0xc8, 0x2a,
	0x28, 0x26, 0x65, 0x10, 0x17, 0x2e, 0x5c, 0x18, 0xa8, 0x08, 0x16, 0x74, 0xa2, 0x1b, 0x37, 0x92,
	0xa6, 0xb7, 0x69, 0xa0, 0xc9, 0x8c, 0x33, 0x37, 0x8b, 0x7e, 0x04, 0xbf, 0xb5, 0x4c, 0xff, 0x68,
	0x91, 0xb8, 0x0a, 0xe7, 0xe4, 0x77, 0xef, 0x9c, 0x33, 0x43, 0x86, 0x4f, 0x02, 0xdf, 0xee, 0x00,
	0x1f, 0x57, 0x29, 0x42, 0x02, 0x59, 0xad, 0x0a, 0x5c, 0x87, 0x52, 0x09, 0x14, 0x74, 0xd0, 0xf4,
	0x6f, 0xf8, 0x37, 0x16, 0x65, 0x29, 0xaa, 0x2d, 0x33, 0xfc, 0x6f, 0x98, 0x43, 0xc7, 0xcf, 0x89,
	0x1d, 0xb3, 0x84, 0x9e, 0x11, 0x47, 0x9a, 0x39, 0xd7, 0xf2, 0x5a, 0x41, 0x9f, 0x0d, 0xc2, 0x03,
	0x70, 0xbf, 0x8b, 0x6f, 0x11, 0x7a, 0x4a, 0x7a, 0x5a, 0x28, 0x9c, 0x14, 0xb0, 0x9a, 0xbb, 0x2d,
	0xcf, 0x0a, 0x1c, 0xfe, 0x6d, 0xd0, 0x63, 0xf2, 0x27, 0xd5, 0x19, 0x54, 0x73, 0xd7, 0xf6, 0xac,
	0xa0, 0xcb, 0x77, 0xca, 0x9f, 0x12, 0x3b, 0x61, 0x31, 0x9d, 0x90, 0x7f, 0x1a, 0x53, 0x2c, 0xb2,
	0xfb, 0x6a, 0x21, 0x1e, 0x0a, 0x8d, 0xae, 0xe5, 0xd9, 0x41, 0x9f, 0x8d, 0x9a, 0x4e, 0x4c, 0xbe,
	0x48, 0xfe, 0x63, 0xca, 0xbf, 0x22, 0x1d, 0x0e, 0xef, 0x35, 0x68, 0xa4, 0xe7, 0xc4, 0xce, 0x98,
	0xde, 0x25, 0x3f, 0x09, 0x1b, 0xaf, 0x28, 0x66, 0x09, 0x37, 0x94, 0xff, 0x61, 0x91, 0x2e, 0x07,
	0x2d, 0x45, 0xa5, 0x81, 0x8e, 0x48, 0x47, 0x01, 0x3e, 0xaf, 0xe5, 0xb6, 0xb7, 0x73, 0xdd, 0xbe,
	0xb8, 0x1c, 0x8f, 0xf9, 0xde, 0x34, 0x5d, 0x14, 0xe0, 0x54, 0xe7, 0x9b, 0x9a, 0x3d, 0xbe, 0x53,
	0xd4, 0x25, 0x1d, 0x50, 0x2a, 0x16, 0x73, 0xd8, 0x94, 0x74, 0xf8, 0x5e, 0x9a, 0x2c, 0x9a, 0x65,
	0x6e, 0xdb, 0xb3, 0x7e, 0xcf, 0x92, 0xb0, 0x98, 0x1b, 0xea, 0xf6, 0x85, 0x1c, 0x65, 0xa2, 0x0c,
	0x17, 0x35, 0xd6, 0xa1, 0x90, 0x50, 0xa5, 0xb2, 0x08, 0xe5, 0xec, 0xf5, 0x26, 0x2f, 0x70, 0x59,
	0xcf, 0xc2, 0x4c, 0x94, 0x91, 0x46, 0x90, 0x4b, 0xa8, 0x56, 0xeb, 0x3a, 0xca, 0x85, 0x01, 0x53,
	0x59, 0x44, 0xe6, 0xbb, 0x79, 0xbd, 0xa8, 0x69, 0xff, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee,
	0x59, 0x80, 0x1b, 0x25, 0x02, 0x00, 0x00,
}
