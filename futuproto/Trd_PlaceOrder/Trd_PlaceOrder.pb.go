// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_PlaceOrder.proto

/*
Package Trd_PlaceOrder is a generated protocol buffer package.

It is generated from these files:
	Trd_PlaceOrder.proto

It has these top-level messages:
	C2S
	S2C
	Request
	Response
*/
package Trd_PlaceOrder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import Common "github.com/stephenlyu/gofutuapi/futuproto/Common"
import Trd_Common "github.com/stephenlyu/gofutuapi/futuproto/Trd_Common"

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
	PacketID  *Common.PacketID      `protobuf:"bytes,1,req,name=packetID" json:"packetID,omitempty"`
	Header    *Trd_Common.TrdHeader `protobuf:"bytes,2,req,name=header" json:"header,omitempty"`
	TrdSide   *int32                `protobuf:"varint,3,req,name=trdSide" json:"trdSide,omitempty"`
	OrderType *int32                `protobuf:"varint,4,req,name=orderType" json:"orderType,omitempty"`
	Code      *string               `protobuf:"bytes,5,req,name=code" json:"code,omitempty"`
	Qty       *float64              `protobuf:"fixed64,6,req,name=qty" json:"qty,omitempty"`
	Price     *float64              `protobuf:"fixed64,7,opt,name=price" json:"price,omitempty"`
	// 以下2个为调整价格使用，都传才有效，对港、A股有意义，因为港股有价位，A股2位精度，美股可不传
	AdjustPrice        *bool    `protobuf:"varint,8,opt,name=adjustPrice" json:"adjustPrice,omitempty"`
	AdjustSideAndLimit *float64 `protobuf:"fixed64,9,opt,name=adjustSideAndLimit" json:"adjustSideAndLimit,omitempty"`
	SecMarket          *int32   `protobuf:"varint,10,opt,name=secMarket" json:"secMarket,omitempty"`
	Remark             *string  `protobuf:"bytes,11,opt,name=remark" json:"remark,omitempty"`
	XXX_unrecognized   []byte   `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2S) GetPacketID() *Common.PacketID {
	if m != nil {
		return m.PacketID
	}
	return nil
}

func (m *C2S) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *C2S) GetTrdSide() int32 {
	if m != nil && m.TrdSide != nil {
		return *m.TrdSide
	}
	return 0
}

func (m *C2S) GetOrderType() int32 {
	if m != nil && m.OrderType != nil {
		return *m.OrderType
	}
	return 0
}

func (m *C2S) GetCode() string {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return ""
}

func (m *C2S) GetQty() float64 {
	if m != nil && m.Qty != nil {
		return *m.Qty
	}
	return 0
}

func (m *C2S) GetPrice() float64 {
	if m != nil && m.Price != nil {
		return *m.Price
	}
	return 0
}

func (m *C2S) GetAdjustPrice() bool {
	if m != nil && m.AdjustPrice != nil {
		return *m.AdjustPrice
	}
	return false
}

func (m *C2S) GetAdjustSideAndLimit() float64 {
	if m != nil && m.AdjustSideAndLimit != nil {
		return *m.AdjustSideAndLimit
	}
	return 0
}

func (m *C2S) GetSecMarket() int32 {
	if m != nil && m.SecMarket != nil {
		return *m.SecMarket
	}
	return 0
}

func (m *C2S) GetRemark() string {
	if m != nil && m.Remark != nil {
		return *m.Remark
	}
	return ""
}

type S2C struct {
	Header           *Trd_Common.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	OrderID          *uint64               `protobuf:"varint,2,opt,name=orderID" json:"orderID,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2C) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *S2C) GetOrderID() uint64 {
	if m != nil && m.OrderID != nil {
		return *m.OrderID
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
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
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
	proto.RegisterType((*C2S)(nil), "Trd_PlaceOrder.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_PlaceOrder.S2C")
	proto.RegisterType((*Request)(nil), "Trd_PlaceOrder.Request")
	proto.RegisterType((*Response)(nil), "Trd_PlaceOrder.Response")
}

func init() { proto.RegisterFile("Trd_PlaceOrder.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x96, 0x93, 0x74, 0x6d, 0x5f, 0x11, 0xaa, 0xbc, 0x81, 0xac, 0x09, 0x21, 0xab, 0x12, 0x52,
	0x0e, 0x2c, 0xad, 0x22, 0x0e, 0x88, 0x1b, 0x64, 0x07, 0x26, 0x31, 0x56, 0x39, 0x3d, 0x71, 0x41,
	0x59, 0xf2, 0x68, 0x43, 0x97, 0xd8, 0xb3, 0x9d, 0x43, 0xcf, 0xfc, 0x9f, 0xfc, 0x2d, 0xc8, 0x6e,
	0x3a, 0x3a, 0xb4, 0x03, 0xa7, 0xf8, 0xfb, 0x91, 0x97, 0xef, 0x7d, 0x31, 0x9c, 0xad, 0x74, 0xf5,
	0x7d, 0x79, 0x57, 0x94, 0x78, 0xa3, 0x2b, 0xd4, 0x89, 0xd2, 0xd2, 0x4a, 0xfa, 0xfc, 0x31, 0x7b,
	0xfe, 0x2c, 0x93, 0x4d, 0x23, 0xdb, 0xbd, 0x7a, 0x3e, 0x75, 0xea, 0x31, 0x33, 0xfb, 0x1d, 0x40,
	0x98, 0xa5, 0x39, 0x7d, 0x0b, 0x23, 0x55, 0x94, 0x5b, 0xb4, 0x57, 0x97, 0x8c, 0xf0, 0x20, 0x9e,
	0xa4, 0xd3, 0xa4, 0x37, 0x2e, 0x7b, 0x5e, 0x3c, 0x38, 0xe8, 0x05, 0x9c, 0x6c, 0xb0, 0xa8, 0x50,
	0xb3, 0xc0, 0x7b, 0x5f, 0x24, 0x47, 0x83, 0x57, 0xba, 0xfa, 0xec, 0x45, 0xd1, 0x9b, 0x28, 0x83,
	0xa1, 0xd5, 0x55, 0x5e, 0x57, 0xc8, 0x42, 0x1e, 0xc4, 0x03, 0x71, 0x80, 0xf4, 0x15, 0x8c, 0xa5,
	0xcb, 0xb9, 0xda, 0x29, 0x64, 0x91, 0xd7, 0xfe, 0x12, 0x94, 0x42, 0x54, 0xca, 0x0a, 0xd9, 0x80,
	0x07, 0xf1, 0x58, 0xf8, 0x33, 0x9d, 0x42, 0x78, 0x6f, 0x77, 0xec, 0x84, 0x07, 0x31, 0x11, 0xee,
	0x48, 0xcf, 0x60, 0xa0, 0x74, 0x5d, 0x22, 0x1b, 0x72, 0x12, 0x13, 0xb1, 0x07, 0x94, 0xc3, 0xa4,
	0xa8, 0x7e, 0x76, 0xc6, 0x2e, 0xbd, 0x36, 0xe2, 0x24, 0x1e, 0x89, 0x63, 0x8a, 0x26, 0x40, 0xf7,
	0xd0, 0x25, 0xf9, 0xd8, 0x56, 0x5f, 0xea, 0xa6, 0xb6, 0x6c, 0xec, 0x87, 0x3c, 0xa1, 0xb8, 0xac,
	0x06, 0xcb, 0xeb, 0x42, 0x6f, 0xd1, 0x32, 0xe0, 0xc4, 0x65, 0x7d, 0x20, 0xe8, 0x4b, 0x38, 0xd1,
	0xd8, 0x14, 0x7a, 0xcb, 0x26, 0x9c, 0xc4, 0x63, 0xd1, 0xa3, 0xd9, 0x57, 0x08, 0xf3, 0x34, 0x3b,
	0x6a, 0x8c, 0xfc, 0x67, 0x63, 0xbe, 0x86, 0xab, 0x4b, 0x16, 0x70, 0x12, 0x47, 0xe2, 0x00, 0x67,
	0x0b, 0x18, 0x0a, 0xbc, 0xef, 0xd0, 0x58, 0xfa, 0x06, 0xc2, 0x32, 0x35, 0xfd, 0xc0, 0xd3, 0xe4,
	0x9f, 0xfb, 0x90, 0xa5, 0xb9, 0x70, 0xfa, 0xec, 0x17, 0x81, 0x91, 0x40, 0xa3, 0x64, 0x6b, 0x90,
	0xbe, 0x86, 0xa1, 0x46, 0xeb, 0xeb, 0x76, 0xef, 0x0d, 0x3e, 0x44, 0x17, 0xef, 0x16, 0x0b, 0x71,
	0x20, 0xf7, 0x6b, 0xd8, 0x6b, 0xb3, 0xf6, 0xdf, 0xf5, 0x6b, 0x38, 0xe4, 0x02, 0xa1, 0xd6, 0x99,
	0xf4, 0xbf, 0xd0, 0xad, 0x7e, 0x80, 0x2e, 0x85, 0x49, 0x4b, 0x16, 0x71, 0xf2, 0x54, 0x8a, 0x3c,
	0xcd, 0x84, 0xd3, 0x3f, 0xdd, 0xc0, 0x69, 0x29, 0x9b, 0xe4, 0x47, 0x67, 0xbb, 0x44, 0x2a, 0x6c,
	0x0b, 0x55, 0x27, 0xea, 0xf6, 0xdb, 0xfb, 0x75, 0x6d, 0x37, 0xdd, 0x6d, 0x52, 0xca, 0x66, 0x6e,
	0x2c, 0xaa, 0x0d, 0xb6, 0x77, 0xbb, 0x6e, 0xbe, 0x96, 0xce, 0x58, 0xa8, 0x7a, 0xee, 0x9e, 0xfe,
	0xaa, 0xce, 0x1f, 0x4f, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x55, 0xa6, 0x0b, 0x00, 0x03,
	0x00, 0x00,
}
