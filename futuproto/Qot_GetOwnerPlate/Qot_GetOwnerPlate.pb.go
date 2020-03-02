// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetOwnerPlate.proto

/*
Package Qot_GetOwnerPlate is a generated protocol buffer package.

It is generated from these files:
	Qot_GetOwnerPlate.proto

It has these top-level messages:
	C2S
	SecurityOwnerPlate
	S2C
	Request
	Response
*/
package Qot_GetOwnerPlate

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
	SecurityList     []*Qot_Common.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S) GetSecurityList() []*Qot_Common.Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

type SecurityOwnerPlate struct {
	Security         *Qot_Common.Security    `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	PlateInfoList    []*Qot_Common.PlateInfo `protobuf:"bytes,2,rep,name=plateInfoList" json:"plateInfoList,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *SecurityOwnerPlate) Reset()                    { *m = SecurityOwnerPlate{} }
func (m *SecurityOwnerPlate) String() string            { return proto.CompactTextString(m) }
func (*SecurityOwnerPlate) ProtoMessage()               {}
func (*SecurityOwnerPlate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SecurityOwnerPlate) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *SecurityOwnerPlate) GetPlateInfoList() []*Qot_Common.PlateInfo {
	if m != nil {
		return m.PlateInfoList
	}
	return nil
}

type S2C struct {
	OwnerPlateList   []*SecurityOwnerPlate `protobuf:"bytes,1,rep,name=ownerPlateList" json:"ownerPlateList,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *S2C) GetOwnerPlateList() []*SecurityOwnerPlate {
	if m != nil {
		return m.OwnerPlateList
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
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	proto.RegisterType((*C2S)(nil), "Qot_GetOwnerPlate.C2S")
	proto.RegisterType((*SecurityOwnerPlate)(nil), "Qot_GetOwnerPlate.SecurityOwnerPlate")
	proto.RegisterType((*S2C)(nil), "Qot_GetOwnerPlate.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetOwnerPlate.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetOwnerPlate.Response")
}

func init() { proto.RegisterFile("Qot_GetOwnerPlate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4b, 0xf3, 0x40,
	0x10, 0xc5, 0x49, 0xd3, 0x7e, 0xed, 0x37, 0x56, 0xd1, 0x55, 0x6b, 0xe8, 0x41, 0x42, 0x40, 0xc8,
	0xc5, 0xa4, 0x44, 0x0f, 0x62, 0x0f, 0x82, 0x39, 0x88, 0x60, 0xb1, 0x6e, 0x7a, 0xf2, 0x22, 0x6d,
	0x9c, 0xb6, 0x81, 0x26, 0xbb, 0x66, 0x37, 0x48, 0xcf, 0x82, 0x7f, 0xb7, 0x6c, 0x6c, 0x6a, 0x6a,
	0xec, 0x69, 0x99, 0x99, 0xdf, 0x7b, 0xfb, 0x66, 0x59, 0x38, 0x79, 0x62, 0xf2, 0xe5, 0x0e, 0xe5,
	0xe3, 0x7b, 0x82, 0xe9, 0x70, 0x31, 0x96, 0xe8, 0xf0, 0x94, 0x49, 0x46, 0x0e, 0x2a, 0x83, 0x6e,
	0xdb, 0x67, 0x71, 0xcc, 0x92, 0x6f, 0xa0, 0xbb, 0xaf, 0x80, 0x72, 0xc7, 0xba, 0x01, 0xdd, 0xf7,
	0x02, 0x72, 0x05, 0x6d, 0x81, 0x61, 0x96, 0x46, 0x72, 0xf9, 0x10, 0x09, 0x69, 0x68, 0xa6, 0x6e,
	0xef, 0x78, 0x47, 0x4e, 0x89, 0x0f, 0x56, 0x73, 0xba, 0x41, 0x5a, 0x1f, 0x1a, 0x90, 0x62, 0xf4,
	0x73, 0x2f, 0xe9, 0x41, 0xab, 0xc0, 0x0c, 0xcd, 0xac, 0x6d, 0x35, 0x5b, 0x53, 0xa4, 0x0f, 0xbb,
	0x5c, 0x49, 0xef, 0x93, 0x29, 0xcb, 0x33, 0xd4, 0xf2, 0x0c, 0xc7, 0x65, 0xd9, 0xb0, 0x00, 0xe8,
	0x26, 0x6b, 0x8d, 0x40, 0x0f, 0x3c, 0x9f, 0x0c, 0x60, 0x8f, 0xad, 0x33, 0x94, 0x16, 0x39, 0x73,
	0xaa, 0x4f, 0x56, 0x0d, 0x4d, 0x7f, 0x89, 0xad, 0x0b, 0x68, 0x52, 0x7c, 0xcb, 0x50, 0x48, 0x62,
	0x83, 0x1e, 0x7a, 0x62, 0xb5, 0x4a, 0xe7, 0x0f, 0x3b, 0xdf, 0x0b, 0xa8, 0x42, 0xac, 0x4f, 0x0d,
	0x5a, 0x14, 0x05, 0x67, 0x89, 0x40, 0x72, 0x0a, 0xcd, 0x14, 0xe5, 0x68, 0xc9, 0x31, 0x97, 0x36,
	0xae, 0xeb, 0xe7, 0x97, 0xbd, 0x1e, 0x2d, 0x9a, 0xa4, 0x03, 0xff, 0x52, 0x94, 0x03, 0x31, 0x33,
	0x6a, 0xa6, 0x66, 0xff, 0xa7, 0xab, 0x8a, 0x18, 0xd0, 0xc4, 0x34, 0xf5, 0xd9, 0x2b, 0x1a, 0xba,
	0xa9, 0xd9, 0x0d, 0x5a, 0x94, 0x2a, 0x88, 0xf0, 0x42, 0xa3, 0x6e, 0x6a, 0x5b, 0x82, 0x04, 0x9e,
	0x4f, 0x15, 0x72, 0x4b, 0xe1, 0x30, 0x64, 0xb1, 0x33, 0xcd, 0x64, 0xe6, 0x30, 0x8e, 0xc9, 0x98,
	0x47, 0x0e, 0x9f, 0x3c, 0xf7, 0x67, 0x91, 0x9c, 0x67, 0x13, 0x27, 0x64, 0xb1, 0x2b, 0x24, 0xf2,
	0x39, 0x26, 0x8b, 0x65, 0xe6, 0xce, 0x98, 0x02, 0xc7, 0x3c, 0x72, 0xd5, 0x99, 0x7f, 0x0e, 0xb7,
	0x62, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xc5, 0xef, 0x37, 0x7b, 0x02, 0x00, 0x00,
}