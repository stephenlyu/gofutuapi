package main

import (
	. "github.com/stephenlyu/gofutuapi/futuapi"
	"github.com/Sirupsen/logrus"
	"time"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetTradeDate"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_Common"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetStaticInfo"
	"encoding/json"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_Sub"
	"github.com/stephenlyu/gofutuapi/futuproto/Notify"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateOrderBook"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateBasicQot"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateKL"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateRT"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateTicker"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateBroker"
	"fmt"
)

func toString(data interface{}) string {
	bytes, _ := json.MarshalIndent(data, "", "  ")
	return string(bytes)
}

type ConnCallback struct {
}

func (cb *ConnCallback) OnInitConnect(conn FTAPIConn, errCode int64, desc string) {
	logrus.Infof("OnInitConnect errCode: %d desc: %s", errCode, desc)

	client := conn.(FTAPIConnQot)
	GetTradeDate(client)
	//GetStaticInfo(client)
	//Sub(client)
}

func (conn *ConnCallback) OnDisconnect(client FTAPIConn, errCode int64) {
	logrus.Infof("OnDisconnect errCode: %d", errCode)
}

type QuoteSpiCallback struct {
	FTSPIQotBase
}

func (spi *QuoteSpiCallback) OnReply_GetTradeDate(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetTradeDate.Response) {
	logrus.Infof("QuoteSpi.OnReply_GetTradeDate nSerialNo: %d response: %s", nSerialNo, rsp.String())
}

func (spi *QuoteSpiCallback) OnReply_GetStaticInfo(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetStaticInfo.Response) {
	logrus.Infof("QuoteSpi.OnReply_GetStaticInfo nSerialNo: %d", nSerialNo)
	println(toString(rsp))
}

func (spi *QuoteSpiCallback) OnReply_Sub(client FTAPIConnQot, nSerialNo uint, rsp *Qot_Sub.Response) {
	logrus.Infof("QuoteSpi.OnReply_Sub nSerialNo: %d", nSerialNo)
	println(toString(rsp))
}

func (spi *QuoteSpiCallback) OnPush_Notify(client FTAPIConnQot, rsp *Notify.Response) {
	logrus.Infof("QuoteSpi.OnPush_Notify")
	println(toString(rsp))
}

func (spi *QuoteSpiCallback) OnPush_UpdateOrderBook(client FTAPIConnQot, rsp *Qot_UpdateOrderBook.Response) {
	logrus.Infof("QuoteSpi.OnPush_UpdateOrderBook")
	println(toString(rsp))
}

func (spi *QuoteSpiCallback) OnPush_UpdateBasicQuote(client FTAPIConnQot, rsp *Qot_UpdateBasicQot.Response) {
	logrus.Infof("QuoteSpi.OnPush_UpdateBasicQuote")
	println(toString(rsp))
}

func (spi *QuoteSpiCallback) OnPush_UpdateKL(client FTAPIConnQot, rsp *Qot_UpdateKL.Response) {
	logrus.Infof("QuoteSpi.OnPush_UpdateKL")
	println(toString(rsp))
}

func (spi *QuoteSpiCallback) OnPush_UpdateRT(client FTAPIConnQot, rsp *Qot_UpdateRT.Response) {
	logrus.Infof("QuoteSpi.OnPush_UpdateRT")
	println(toString(rsp))
}

func (spi *QuoteSpiCallback) OnPush_UpdateTicker(client FTAPIConnQot, rsp *Qot_UpdateTicker.Response) {
	logrus.Infof("QuoteSpi.OnPush_UpdateTicker")
	println(toString(rsp))
}

func (spi *QuoteSpiCallback) OnPush_UpdateBroker(client FTAPIConnQot, rsp *Qot_UpdateBroker.Response) {
	logrus.Infof("QuoteSpi.OnPush_UpdateBroker")
	println(toString(rsp))
}

func GetStaticInfo(client FTAPIConnQot) {
	req := &Qot_GetStaticInfo.Request{
		C2S:&Qot_GetStaticInfo.C2S{
			Market: MakeInt32Pointer(int32(Qot_Common.QotMarket_QotMarket_HK_Security)),
			SecType: MakeInt32Pointer(int32(Qot_Common.SecurityType_SecurityType_Future)),
			SecurityList: []*Qot_Common.Security{
				{
					Market: MakeInt32Pointer(int32(Qot_Common.QotMarket_QotMarket_HK_Security)),
					Code: MakeStringPointer("999010"),
				},
			},
		},
	}
	ret := client.GetStaticInfo(req)
	logrus.Infof("GetStaticInfo ret: %d", ret)
}

func GetTradeDate(client FTAPIConnQot) {
	req := &Qot_GetTradeDate.Request{
		C2S: &Qot_GetTradeDate.C2S{
			Market: MakeInt32Pointer(int32(Qot_Common.QotMarket_QotMarket_HK_Security)),
			BeginTime: MakeStringPointer("2020-01-01"),
			EndTime: MakeStringPointer("2020-03-01"),
		},
	}
	ret := client.GetTradeDate(req)
	logrus.Infof("GetTradeDate ret: %d", ret)
}

func Sub(client FTAPIConnQot) {
	req := &Qot_Sub.Request{
		C2S: &Qot_Sub.C2S{
			SecurityList: []*Qot_Common.Security{
				{
					Market: MakeInt32Pointer(int32(Qot_Common.QotMarket_QotMarket_HK_Security)),
					Code: MakeStringPointer("HSImain"),
				},
			},
			SubTypeList: []int32{
				int32(Qot_Common.SubType_SubType_Basic),
				int32(Qot_Common.SubType_SubType_Ticker),
				int32(Qot_Common.SubType_SubType_OrderBook),
			},
			IsSubOrUnSub: MakeBoolPointer(true),
			IsRegOrUnRegPush: MakeBoolPointer(true),
		},
	}
	ret := client.Sub(req)
	logrus.Infof("Sub ret: %d", ret)
}

//func GetFutureInfo(client FTAPIConnQot) {
//	req := &Qot_GetFutureInfo.Request{
//		C2S: &Qot_GetFutureInfo.C2S{
//			SecurityList: []*Qot_Common.Security{
//				{
//					Market: MakeInt32Pointer(int32(Qot_Common.QotMarket_QotMarket_HK_Security)),
//					Code: MakeStringPointer("999010"),
//				},
//			},
//		},
//	}
//	ret := client.Sub(req)
//	logrus.Infof("Sub ret: %d", ret)
//}

func main() {
	var config Config
	err := config.Load("../config.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", config)

	Init()
	client := NewFTAPIConnQot()
	client.SetSpi(&ConnCallback{})
	client.SetQotSpi(&QuoteSpiCallback{})
	client.SetClientInfo(config.ClientInfo, 1);
	client.SetRSAPrivateKey(config.PrivateKey)
	client.InitConnect(config.IP, config.Port, true)

	time.Sleep(time.Second * 10)

	client.Disconnect()

	time.Sleep(time.Second * 10)
}
