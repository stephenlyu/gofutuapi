package main

import (
	. "github.com/stephenlyu/gofutuapi/futuapi"
	"github.com/Sirupsen/logrus"
	"time"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetTradeDate"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_Common"
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
}

func (cb *ConnCallback) OnInitConnect(conn FTAPIConn, errCode int64, desc string) {
	logrus.Infof("OnInitConnect errCode: %d desc: %s", errCode, desc)

	client := conn.(FTAPIConnQot)

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

func (conn *ConnCallback) OnDisconnect(client FTAPIConn, errCode int64) {
	logrus.Infof("OnDisconnect errCode: %d", errCode)
}

type QuoteSpiCallback struct {
	FTSPIQotBase
}

func (spi *QuoteSpiCallback) OnReply_GetTradeDate(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetTradeDate.Response) {
	logrus.Infof("QuoteSpi.OnReply_GetTradeDate nSerialNo: %d response: %s", nSerialNo, rsp.String())
}

func main() {
	Init()
	client := NewFTAPIConnQot()
	client.SetSpi(&ConnCallback{})
	client.SetQotSpi(&QuoteSpiCallback{})
	client.SetClientInfo("FTAPITest", 1);
	client.SetRSAPrivateKey(PrivateKey)
	client.InitConnect("118.190.77.238", 11111, true)

	time.Sleep(time.Second * 10)

	client.Disconnect()

	time.Sleep(time.Second * 10)
}
