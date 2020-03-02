package futuapi

import (
	. "github.com/stephenlyu/gofutuapi"
	"github.com/golang/protobuf/proto"
)

type FTAPIConn struct {
	channel uintptr
	connSpi FTSpiConn
}

func NewFTAPIConn() *FTAPIConn {
	conn := new(FTAPIConn)

	conn.channel = CreateFTAPIChannel()
	FTAPIChannelSetCallbacks(conn.channel, NewDirectorFTAPIChannel_Callback(conn))

	return conn
}

func (conn *FTAPIConn) SetSpi(connSpi FTSpiConn) {
	conn.connSpi = connSpi
}

func (conn *FTAPIConn) OnDisConnect(arg2 uintptr, arg3 int64) {
	if conn.connSpi != nil {
		conn.connSpi.OnDisconnect(conn, arg3)
	}
}

func (conn *FTAPIConn) OnInitConnect(arg2 uintptr, arg3 int64, arg4 string) {
	if conn.connSpi != nil {
		conn.connSpi.OnInitConnect(conn, arg3, arg4)
	}
}

func (conn *FTAPIConn) OnReply(arg2 uintptr, arg3 FTAPI_ReqReplyType, arg4 FTAPI_ProtoHeader, arg5 string) {

}

func (conn *FTAPIConn) OnPush(arg2 uintptr, arg3 FTAPI_ProtoHeader, arg4 string) {

}

func (conn *FTAPIConn) Close() {
	if conn.channel != 0 {
		FTAPIChannelDestroy(conn.channel)
		conn.channel = 0
	}
}

func (conn *FTAPIConn) Disconnect() bool {
	if conn.channel != 0 {
		return FTAPIChannel_Close(conn.channel) == 0
	}
	return false
}

func (conn *FTAPIConn) SetClientInfo(clientID string, clientVer int) {
	if (conn.channel != 0) {
		FTAPIChannel_SetClientInfo(conn.channel, clientID, clientVer)
	}
}

/***
 * 设置加密私钥
 * @param key
 */
func (conn *FTAPIConn) SetRSAPrivateKey(key string) {
	if (conn.channel != 0) {
		FTAPIChannel_SetRSAPrivateKey(conn.channel, key);
	}
}

/**
 * @param ip              地址
 * @param port            端口
 * @param isEnableEncrypt 启用加密
 * @return bool 是否启动了执行，不代表连接结果，结果通过OnInitConnect回调
 * @brief 初始化连接，连接并初始化
 */
func (conn *FTAPIConn) InitConnect(ip string, port uint16, isEnableEncrypt bool) bool {
	if conn.channel != 0 {
		var arg4 int
		if isEnableEncrypt {
			arg4 = 1
		}
		ret := FTAPIChannel_InitConnect(conn.channel, ip, port, arg4)
		return ret == 0
	}
	return false;
}

/**
 * 此连接的连接ID，连接的唯一标识，InitConnect协议返回，没有初始化前为0
 *
 * @return
 */
func (conn *FTAPIConn) GetConnectID() uint64 {
	if (conn.channel != 0) {
		return FTAPIChannel_GetConnectID(conn.channel)
	}
	return 0
}

func (conn *FTAPIConn) SendProto(protoID uint, req proto.Message) uint {
	if (conn.channel != 0) {
		data, _ := proto.Marshal(req)
		return FTAPIChannel_SendProto(conn.channel, protoID, 0, string(data))
	}
	return 0
}

