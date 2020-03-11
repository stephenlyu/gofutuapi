package futuapi

type FTSpiConn interface {
	/***
	 * InitConnect协议返回。如果返回成功，则可以发起行情或交易请求。
	 * @param client
	 * @param errCode 错误码。当高32位为ConnectFailType类型时，低32位为系统错误码；当高32位等于FTAPI.INIT_FAIL，则低32位为InitFailType类型。
	 * @param desc 错误描述
	 */
	OnInitConnect(client FTAPIConn, errCode int64, desc string)

	/***
	 * 与OpenD的连接断开
	 * @param client
	 * @param errCode
	 */
	OnDisconnect(client FTAPIConn, errCode int64)
}
