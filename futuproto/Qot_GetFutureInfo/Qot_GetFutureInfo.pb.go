// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetFutureInfo.proto

/*
Package Qot_GetFutureInfo is a generated protocol buffer package.

It is generated from these files:
	Qot_GetFutureInfo.proto

It has these top-level messages:
	TradeTime
	FutureInfo
	C2S
	S2C
	Request
	Response
*/
package Qot_GetFutureInfo

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

// 交易时间
type TradeTime struct {
	Begin            *float64 `protobuf:"fixed64,1,opt,name=begin" json:"begin,omitempty"`
	End              *float64 `protobuf:"fixed64,2,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TradeTime) Reset()                    { *m = TradeTime{} }
func (m *TradeTime) String() string            { return proto.CompactTextString(m) }
func (*TradeTime) ProtoMessage()               {}
func (*TradeTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TradeTime) GetBegin() float64 {
	if m != nil && m.Begin != nil {
		return *m.Begin
	}
	return 0
}

func (m *TradeTime) GetEnd() float64 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

// 期货合约资料的列表
type FutureInfo struct {
	Name               *string              `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Security           *Qot_Common.Security `protobuf:"bytes,2,req,name=security" json:"security,omitempty"`
	LastTradeTime      *string              `protobuf:"bytes,3,req,name=lastTradeTime" json:"lastTradeTime,omitempty"`
	LastTradeTimestamp *float64             `protobuf:"fixed64,4,opt,name=lastTradeTimestamp" json:"lastTradeTimestamp,omitempty"`
	Owner              *Qot_Common.Security `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`
	OwnerOther         *string              `protobuf:"bytes,6,req,name=ownerOther" json:"ownerOther,omitempty"`
	Exchange           *string              `protobuf:"bytes,7,req,name=exchange" json:"exchange,omitempty"`
	ContractType       *string              `protobuf:"bytes,8,req,name=contractType" json:"contractType,omitempty"`
	ContractSize       *float64             `protobuf:"fixed64,9,req,name=contractSize" json:"contractSize,omitempty"`
	ContractSizeUnit   *string              `protobuf:"bytes,10,req,name=contractSizeUnit" json:"contractSizeUnit,omitempty"`
	QuoteCurrency      *string              `protobuf:"bytes,11,req,name=quoteCurrency" json:"quoteCurrency,omitempty"`
	MinVar             *float64             `protobuf:"fixed64,12,req,name=minVar" json:"minVar,omitempty"`
	MinVarUnit         *string              `protobuf:"bytes,13,req,name=minVarUnit" json:"minVarUnit,omitempty"`
	QuoteUnit          *string              `protobuf:"bytes,14,opt,name=quoteUnit" json:"quoteUnit,omitempty"`
	TradeTime          []*TradeTime         `protobuf:"bytes,15,rep,name=tradeTime" json:"tradeTime,omitempty"`
	TimeZone           *string              `protobuf:"bytes,16,req,name=timeZone" json:"timeZone,omitempty"`
	ExchangeFormatUrl  *string              `protobuf:"bytes,17,req,name=exchangeFormatUrl" json:"exchangeFormatUrl,omitempty"`
	XXX_unrecognized   []byte               `json:"-"`
}

func (m *FutureInfo) Reset()                    { *m = FutureInfo{} }
func (m *FutureInfo) String() string            { return proto.CompactTextString(m) }
func (*FutureInfo) ProtoMessage()               {}
func (*FutureInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FutureInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FutureInfo) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *FutureInfo) GetLastTradeTime() string {
	if m != nil && m.LastTradeTime != nil {
		return *m.LastTradeTime
	}
	return ""
}

func (m *FutureInfo) GetLastTradeTimestamp() float64 {
	if m != nil && m.LastTradeTimestamp != nil {
		return *m.LastTradeTimestamp
	}
	return 0
}

func (m *FutureInfo) GetOwner() *Qot_Common.Security {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *FutureInfo) GetOwnerOther() string {
	if m != nil && m.OwnerOther != nil {
		return *m.OwnerOther
	}
	return ""
}

func (m *FutureInfo) GetExchange() string {
	if m != nil && m.Exchange != nil {
		return *m.Exchange
	}
	return ""
}

func (m *FutureInfo) GetContractType() string {
	if m != nil && m.ContractType != nil {
		return *m.ContractType
	}
	return ""
}

func (m *FutureInfo) GetContractSize() float64 {
	if m != nil && m.ContractSize != nil {
		return *m.ContractSize
	}
	return 0
}

func (m *FutureInfo) GetContractSizeUnit() string {
	if m != nil && m.ContractSizeUnit != nil {
		return *m.ContractSizeUnit
	}
	return ""
}

func (m *FutureInfo) GetQuoteCurrency() string {
	if m != nil && m.QuoteCurrency != nil {
		return *m.QuoteCurrency
	}
	return ""
}

func (m *FutureInfo) GetMinVar() float64 {
	if m != nil && m.MinVar != nil {
		return *m.MinVar
	}
	return 0
}

func (m *FutureInfo) GetMinVarUnit() string {
	if m != nil && m.MinVarUnit != nil {
		return *m.MinVarUnit
	}
	return ""
}

func (m *FutureInfo) GetQuoteUnit() string {
	if m != nil && m.QuoteUnit != nil {
		return *m.QuoteUnit
	}
	return ""
}

func (m *FutureInfo) GetTradeTime() []*TradeTime {
	if m != nil {
		return m.TradeTime
	}
	return nil
}

func (m *FutureInfo) GetTimeZone() string {
	if m != nil && m.TimeZone != nil {
		return *m.TimeZone
	}
	return ""
}

func (m *FutureInfo) GetExchangeFormatUrl() string {
	if m != nil && m.ExchangeFormatUrl != nil {
		return *m.ExchangeFormatUrl
	}
	return ""
}

type C2S struct {
	SecurityList     []*Qot_Common.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *C2S) Reset()                    { *m = C2S{} }
func (m *C2S) String() string            { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()               {}
func (*C2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *C2S) GetSecurityList() []*Qot_Common.Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

type S2C struct {
	FutureInfoList   []*FutureInfo `protobuf:"bytes,1,rep,name=futureInfoList" json:"futureInfoList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *S2C) Reset()                    { *m = S2C{} }
func (m *S2C) String() string            { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()               {}
func (*S2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *S2C) GetFutureInfoList() []*FutureInfo {
	if m != nil {
		return m.FutureInfoList
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
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

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
	proto.RegisterType((*TradeTime)(nil), "Qot_GetFutureInfo.TradeTime")
	proto.RegisterType((*FutureInfo)(nil), "Qot_GetFutureInfo.FutureInfo")
	proto.RegisterType((*C2S)(nil), "Qot_GetFutureInfo.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetFutureInfo.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetFutureInfo.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetFutureInfo.Response")
}

func init() { proto.RegisterFile("Qot_GetFutureInfo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x51, 0x6f, 0xd3, 0x40,
	0x0c, 0x56, 0xda, 0x75, 0x6d, 0xdc, 0x6e, 0x74, 0xc7, 0x34, 0x4e, 0xd3, 0x98, 0xa2, 0x88, 0x87,
	0x68, 0x82, 0x74, 0xca, 0x78, 0x40, 0xe3, 0x01, 0x89, 0x88, 0x21, 0xa4, 0x21, 0xc4, 0x65, 0xe3,
	0x61, 0x2f, 0x28, 0xcb, 0xbc, 0x36, 0xd2, 0x72, 0x97, 0x5d, 0x2e, 0x82, 0xf2, 0x03, 0xf6, 0xbb,
	0xd1, 0x5d, 0xda, 0x34, 0xa1, 0xe5, 0xa9, 0xf6, 0xe7, 0x2f, 0xf6, 0x67, 0xfb, 0x5c, 0x78, 0xf1,
	0x5d, 0xa8, 0x9f, 0x9f, 0x51, 0x5d, 0x94, 0xaa, 0x94, 0xf8, 0x85, 0xdf, 0x0b, 0x3f, 0x97, 0x42,
	0x09, 0xb2, 0xb7, 0x16, 0x38, 0x1c, 0x85, 0x22, 0xcb, 0x04, 0xaf, 0x08, 0x87, 0x63, 0x4d, 0x68,
	0x22, 0xee, 0x19, 0xd8, 0x57, 0x32, 0xbe, 0xc3, 0xab, 0x34, 0x43, 0xb2, 0x0f, 0xbd, 0x5b, 0x9c,
	0xa6, 0x9c, 0x5a, 0x8e, 0xe5, 0x59, 0xac, 0x72, 0xc8, 0x18, 0xba, 0xc8, 0xef, 0x68, 0xc7, 0x60,
	0xda, 0x74, 0x9f, 0x7a, 0x00, 0xab, 0x1a, 0x84, 0xc0, 0x16, 0x8f, 0x33, 0xa4, 0x96, 0xd3, 0xf1,
	0x6c, 0x66, 0x6c, 0x72, 0x0a, 0x83, 0x02, 0x93, 0x52, 0xa6, 0x6a, 0x4e, 0x3b, 0x4e, 0xc7, 0x1b,
	0x06, 0xfb, 0x7e, 0xa3, 0x78, 0xb4, 0x88, 0xb1, 0x9a, 0x45, 0x5e, 0xc1, 0xce, 0x43, 0x5c, 0xa8,
	0x5a, 0x0d, 0xed, 0x9a, 0x74, 0x6d, 0x90, 0xf8, 0x40, 0x5a, 0x40, 0xa1, 0xe2, 0x2c, 0xa7, 0x5b,
	0x46, 0xdb, 0x86, 0x08, 0x39, 0x81, 0x9e, 0xf8, 0xc5, 0x51, 0xd2, 0x9e, 0x63, 0xfd, 0x57, 0x44,
	0x45, 0x21, 0xc7, 0x00, 0xc6, 0xf8, 0xa6, 0x66, 0x28, 0xe9, 0xb6, 0x29, 0xdf, 0x40, 0xc8, 0x21,
	0x0c, 0xf0, 0x77, 0x32, 0x8b, 0xf9, 0x14, 0x69, 0xdf, 0x44, 0x6b, 0x9f, 0xb8, 0x30, 0x4a, 0x04,
	0x57, 0x32, 0x4e, 0xd4, 0xd5, 0x3c, 0x47, 0x3a, 0x30, 0xf1, 0x16, 0xd6, 0xe4, 0x44, 0xe9, 0x1f,
	0xa4, 0xb6, 0xd3, 0xf1, 0x2c, 0xd6, 0xc2, 0xc8, 0x09, 0x8c, 0x9b, 0xfe, 0x35, 0x4f, 0x15, 0x05,
	0x93, 0x6b, 0x0d, 0xd7, 0x13, 0x7b, 0x2c, 0x85, 0xc2, 0xb0, 0x94, 0x12, 0x79, 0x32, 0xa7, 0xc3,
	0x6a, 0x62, 0x2d, 0x90, 0x1c, 0xc0, 0x76, 0x96, 0xf2, 0x1f, 0xb1, 0xa4, 0x23, 0x53, 0x6f, 0xe1,
	0xe9, 0x6e, 0x2b, 0xcb, 0xd4, 0xd8, 0xa9, 0xba, 0x5d, 0x21, 0xe4, 0x08, 0x6c, 0x93, 0xc8, 0x84,
	0x77, 0x1d, 0xcb, 0xb3, 0xd9, 0x0a, 0x20, 0xe7, 0x60, 0xab, 0x7a, 0x53, 0xcf, 0x9c, 0xae, 0x37,
	0x0c, 0x8e, 0xfc, 0xf5, 0x77, 0x59, 0x6f, 0x83, 0xad, 0xe8, 0x7a, 0x8e, 0x2a, 0xcd, 0xf0, 0x46,
	0x70, 0xa4, 0xe3, 0x6a, 0x8e, 0x4b, 0x9f, 0xbc, 0x86, 0xbd, 0xe5, 0x4c, 0x2f, 0x84, 0xcc, 0x62,
	0x75, 0x2d, 0x1f, 0xe8, 0x9e, 0x21, 0xad, 0x07, 0xdc, 0x0f, 0xd0, 0x0d, 0x83, 0x88, 0xbc, 0x83,
	0xd1, 0xf2, 0x19, 0x5d, 0xa6, 0x85, 0xa2, 0x96, 0xd1, 0xb3, 0x79, 0xd7, 0x2d, 0xa6, 0x7b, 0x09,
	0xdd, 0x28, 0x08, 0xc9, 0x27, 0xd8, 0xbd, 0xaf, 0x45, 0x37, 0x52, 0xbc, 0xdc, 0xd0, 0xd2, 0xca,
	0x64, 0xff, 0x7c, 0xe4, 0x9e, 0x41, 0x9f, 0xe1, 0x63, 0x89, 0x85, 0x22, 0x1e, 0x74, 0x93, 0xa0,
	0x30, 0x27, 0x31, 0x0c, 0x0e, 0x36, 0xa4, 0x09, 0x83, 0x88, 0x69, 0x8a, 0xfb, 0x64, 0xc1, 0x80,
	0x61, 0x91, 0x0b, 0x5e, 0x20, 0x39, 0x86, 0xbe, 0xc4, 0xea, 0x05, 0xe9, 0x4f, 0x7b, 0xe7, 0x5b,
	0x6f, 0xde, 0x9e, 0x9e, 0xb2, 0x25, 0xa8, 0x97, 0x29, 0x51, 0x7d, 0x2d, 0xa6, 0xe6, 0x1c, 0x6d,
	0xb6, 0xf0, 0x08, 0x85, 0x3e, 0x4a, 0x19, 0x8a, 0x3b, 0x7d, 0x36, 0x96, 0xd7, 0x63, 0x4b, 0x57,
	0x0b, 0x29, 0x82, 0xc4, 0x5c, 0xc8, 0x66, 0x21, 0x51, 0x10, 0x32, 0x4d, 0xf9, 0xc8, 0xe0, 0x79,
	0x22, 0x32, 0x5f, 0xf7, 0xe4, 0x8b, 0x1c, 0x79, 0x9c, 0xa7, 0x7e, 0x7e, 0x7b, 0xf3, 0x7e, 0x9a,
	0xaa, 0x59, 0x79, 0xeb, 0x27, 0x22, 0x9b, 0x14, 0x0a, 0xf3, 0x19, 0xf2, 0x87, 0x79, 0x39, 0x99,
	0x0a, 0x4d, 0x8c, 0xf3, 0x74, 0xa2, 0x7f, 0xcd, 0x9f, 0xc9, 0x64, 0x2d, 0xf9, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7d, 0x93, 0x3f, 0xa6, 0xab, 0x04, 0x00, 0x00,
}