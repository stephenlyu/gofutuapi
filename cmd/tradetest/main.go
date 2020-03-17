package main

import (
	. "github.com/stephenlyu/gofutuapi/futuapi"
	"github.com/Sirupsen/logrus"
	"time"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetAccList"
	"runtime"
	"fmt"
)

type ConnCallback struct {
	ch chan bool
}

func (cb *ConnCallback) OnInitConnect(conn FTAPIConn, errCode int64, desc string) {
	logrus.Infof("OnInitConnect errCode: %d desc: %s connectionId: %d", errCode, desc, conn.GetConnectID())
	cb.ch <- true
}

func (conn *ConnCallback) OnDisconnect(client FTAPIConn, errCode int64) {
	logrus.Infof("OnDisconnect errCode: %d", errCode)
}

type TradeSpiCallback struct {
	FTSPI_TrdBase
	ch chan bool
}

func (spi *TradeSpiCallback) OnReply_GetAccList(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_GetAccList.Response) {
	logrus.Infof("TradeSpi.OnReply_GetAccList nSerialNo: %d response: %s", nSerialNo, rsp.String())
}

func main() {
	var config Config
	err := config.Load("../config.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", config)

	ch := make(chan bool)

	Init()

	client := NewFTAPIConnTrd()
	client.SetSpi(&ConnCallback{ch:ch})
	client.SetTrdSpi(&TradeSpiCallback{ch:ch})
	client.SetClientInfo(config.ClientInfo, 1);
	client.SetRSAPrivateKey(config.PrivateKey)
	client.InitConnect(config.IP, config.Port, true)

	<-ch

	req := &Trd_GetAccList.Request{
		C2S: &Trd_GetAccList.C2S{
			UserID: MakeUInt64Pointer(config.UserId),
		},
	}
	ret := client.GetAccList(req)
	logrus.Infof("GetAccList ret: %d", ret)

	time.Sleep(time.Second * 10)

	client.Disconnect()

	time.Sleep(time.Second * 10)
	runtime.UnlockOSThread()
}
