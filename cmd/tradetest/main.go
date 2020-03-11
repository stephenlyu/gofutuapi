package main

import (
	. "github.com/stephenlyu/gofutuapi/futuapi"
	"github.com/Sirupsen/logrus"
	"time"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetAccList"
	"runtime"
)

const PrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQColSw8thhKlEY/yYmJ7bhgMqO3QGQhwOFNIT5PoF/diw+mm0Hh
tQn3mXMHCD1hQxUhFIICLj2174c1ZRBeDYJVNZtEsNddKKlKe092GAKQPLCNYpqr
P49sl+FwROWIHEAcIr2aYO7CHnXCRN1uIXf2arXixElwL1kdWpAG/COTMwIDAQAB
AoGAJjL2/SK9ylhiup1uHuTQvGt9EU7z4XoVEycPOXe7gTW7bCMOAJjHE2Wf3N4P
GnTa2s4Mz3Wu4gTOfFjUJpulBk+WNEi8jDMI5FWHMLh/PqswAbGFj5lH4DyPzjmp
wDbzGndhhMB/T3GdIeM1g5Cc+GvGxrINOFAe4agx0pGkb3ECQQDW2fhdf+nao3Ok
N4UEMiPSGAbW6ggy1XsWxjv2yGSdRtfaDfj4PRHQWvQDaDpyafnxYSnwXaXSlKl+
vALQE3z1AkEAyN6s/GhulqY+BqWkePBmKxwl17oiMMPYz3lrYDcaUA+Mts81zXAl
lW4UNLYTjEqdivaN2pOoe1QyFyO3bt02hwJAX4YV6OxAOxdFCRQuLcllJ7nLAK6Y
6pED4wJMEtLR+SNQQQDJWwU78Fkf+IvUwJ3hpLJAhT/9w/yYx2IsFfs0KQJABDCY
3R70h5HqI0tbNeaVyvpoU6qfQfMj15gJxFUB6H+aiMmjrqhTMF2+cCcIG1oHFTn1
VYTU89Wawd7N2bMliwJBAJmYeO7DD+7gJ9rj8+LmtDKEdtYyPDCoq98471Fn+nF4
3kdn76A4adzHzMHSO5rLq4zgtB1RRGN90RhCwmIwYs0=
-----END RSA PRIVATE KEY-----`

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

	ch := make(chan bool)

	runtime.LockOSThread()

	println("here")
	Init()

	client := NewFTAPIConnTrd()
	client.SetSpi(&ConnCallback{ch:ch})
	client.SetTrdSpi(&TradeSpiCallback{ch:ch})
	client.SetClientInfo("javaclient", 1);
	client.SetRSAPrivateKey(PrivateKey)
	client.InitConnect("118.190.77.238", 11111, true)

	<- ch

	req := &Trd_GetAccList.Request{
		C2S: &Trd_GetAccList.C2S{
			UserID: MakeUInt64Pointer(5883618),
		},
	}
	ret := client.GetAccList(req)
	logrus.Infof("GetAccList ret: %d", ret)

	time.Sleep(time.Second * 10)

	client.Disconnect()

	time.Sleep(time.Second * 10)
	runtime.UnlockOSThread()
}
