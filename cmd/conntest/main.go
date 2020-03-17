package main

import (
	. "github.com/stephenlyu/gofutuapi/futuapi"
	"github.com/Sirupsen/logrus"
	"time"
	"fmt"
)

type ConnCallback struct {
}

func (conn *ConnCallback) OnInitConnect(_ FTAPIConn, errCode int64, desc string) {
	logrus.Infof("OnInitConnect errCode: %d desc: %s", errCode, desc)
}

func (conn *ConnCallback) OnDisconnect(_ FTAPIConn, errCode int64) {
	logrus.Infof("OnDisconnect errCode: %d", errCode)
}

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

	client.SetClientInfo(config.ClientInfo, 1);
	client.SetRSAPrivateKey(config.PrivateKey)
	client.InitConnect(config.IP, config.Port, true)

	time.Sleep(time.Second * 10)

	client.Disconnect()

	time.Sleep(time.Second * 10)
}
