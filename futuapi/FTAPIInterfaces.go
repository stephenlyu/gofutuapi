package futuapi

import (
	"github.com/golang/protobuf/proto"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetCapitalDistribution"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetUserSecurity"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_ModifyUserSecurity"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_StockFilter"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetCodeChange"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetIpoList"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetFutureInfo"
	"github.com/stephenlyu/gofutuapi/futuproto/GetGlobalState"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_Sub"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_RegQotPush"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetSubInfo"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetTicker"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetBasicQot"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetOrderBook"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetKL"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetRT"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetBroker"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_RequestRehab"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_RequestHistoryKL"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_RequestHistoryKLQuota"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetTradeDate"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetStaticInfo"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetSecuritySnapshot"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetPlateSet"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetPlateSecurity"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetReference"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetOwnerPlate"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetHoldingChangeList"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetOptionChain"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetWarrant"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetCapitalFlow"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetOrderList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_PlaceOrder"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_ModifyOrder"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetOrderFillList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetHistoryOrderList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetHistoryOrderFillList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetAccList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_UnlockTrade"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_SubAccPush"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetFunds"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetPositionList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetMaxTrdQtys"
)

type FTAPIConn interface {
	SetSpi(connSpi FTSpiConn)
	Close()

	Disconnect() bool
	SetClientInfo(clientID string, clientVer int)

	/***
	 * 设置加密私钥
	 * @param key
	 */
	SetRSAPrivateKey(key string)

	/**
	 * @param ip              地址
	 * @param port            端口
	 * @param isEnableEncrypt 启用加密
	 * @return bool 是否启动了执行，不代表连接结果，结果通过OnInitConnect回调
	 * @brief 初始化连接，连接并初始化
	 */
	InitConnect(ip string, port uint16, isEnableEncrypt bool) bool

	/**
	 * 此连接的连接ID，连接的唯一标识，InitConnect协议返回，没有初始化前为0
	 *
	 * @return
	 */
	GetConnectID() uint64

	SendProto(protoID uint, req proto.Message) uint
}

type FTAPIConnQot interface {
	FTAPIConn

	SetQotSpi(qotSpi FTSPIQot)

	GetGlobalState(req *GetGlobalState.Request) uint
	/***
	 * 订阅或者反订阅，具体字段请参考Qot_Sub.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	Sub(req *Qot_Sub.Request) uint

	/***
	 * 注册推送，具体字段请参考Qot_RegQotPush.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	RegQotPush(req *Qot_RegQotPush.Request) uint
	/***
	 * 获取订阅信息，具体字段请参考Qot_GetSubInfo.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetSubInfo(req *Qot_GetSubInfo.Request) uint
	/***
	 * 获取逐笔,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Ticker)，具体字段请参考Qot_GetTicker.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetTicker(req *Qot_GetTicker.Request) uint
	/***
	 * 获取基本行情,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Basic)，具体字段请参考Qot_GetBasicQot.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetBasicQot(req *Qot_GetBasicQot.Request) uint
	/***
	 * 获取摆盘,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_OrderBook)，具体字段请参考Qot_GetOrderBook.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetOrderBook(req *Qot_GetOrderBook.Request) uint
	/***
	 * 获取K线，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_KL_XXX)，具体字段请参考Qot_GetKL.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetKL(req *Qot_GetKL.Request) uint

	/***
	 * 获取分时，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_RT)，具体字段请参考Qot_GetRT.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetRT(req *Qot_GetRT.Request) uint
	/***
	 * 获取经纪队列，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Broker)，具体字段请参考Qot_GetBroker.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetBroker(req *Qot_GetBroker.Request) uint
	/***
	 * 在线请求历史复权信息，不读本地历史数据DB，具体字段请参考Qot_RequestRehab.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	RequestRehab(req *Qot_RequestRehab.Request) uint

	/***
	 * 在线请求历史K线，不读本地历史数据DB，具体字段请参考Qot_RequestHistoryKL.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	RequestHistoryKL(req *Qot_RequestHistoryKL.Request) uint

	/***
	 * 获取历史K线已经用掉的额度，具体字段请参考Qot_RequestHistoryKLQuota.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	RequestHistoryKLQuota(req *Qot_RequestHistoryKLQuota.Request) uint

	/***
	 * 获取交易日，具体字段请参考Qot_GetTradeDate.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetTradeDate(req *Qot_GetTradeDate.Request) uint

	/***
	 * 获取静态信息，具体字段请参考Qot_GetStaticInfo.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetStaticInfo(req *Qot_GetStaticInfo.Request) uint

	/***
	 * 获取股票快照，具体字段请参考Qot_GetSecuritySnapshot.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetSecuritySnapshot(req *Qot_GetSecuritySnapshot.Request) uint

	/***
	 * 获取板块集合下的板块，具体字段请参考Qot_GetPlateSet.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetPlateSet(req *Qot_GetPlateSet.Request) uint

	/***
	 * 获取板块下的股票，具体字段请参考Qot_GetPlateSecurity.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetPlateSecurity(req *Qot_GetPlateSecurity.Request) uint
	/***
	 * 获取相关股票，具体字段请参考Qot_GetReference.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetReference(req *Qot_GetReference.Request) uint

	/***
	 * 获取股票所属的板块，具体字段请参考Qot_GetOwnerPlate.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetOwnerPlate(req *Qot_GetOwnerPlate.Request) uint

	/***
	 * 获取大股东持股变化列表，具体字段请参考Qot_GetHoldingChangeList.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetHoldingChangeList(req *Qot_GetHoldingChangeList.Request) uint
	/***
	 * 筛选期权，具体字段请参考Qot_GetOptionChain.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetOptionChain(req *Qot_GetOptionChain.Request) uint

	/***
	 * 筛选窝轮，具体字段请参考Qot_GetWarrant.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetWarrant(req *Qot_GetWarrant.Request) uint

	/***
	 * 获取资金流向，具体字段请参考Qot_GetCapitalFlow.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetCapitalFlow(req *Qot_GetCapitalFlow.Request) uint
	/***
	 * 获取资金分布，具体字段请参考Qot_GetCapitalDistribution.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetCapitalDistribution(req *Qot_GetCapitalDistribution.Request) uint
	/***
	 * 获取自选股分组下的股票，具体字段请参考Qot_GetUserSecurity.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetUserSecurity(req *Qot_GetUserSecurity.Request) uint

	/***
	 * 修改自选股分组下的股票，具体字段请参考Qot_ModifyUserSecurity.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	ModifyUserSecurity(req *Qot_ModifyUserSecurity.Request) uint
	/***
	 * 条件选股，具体字段请参考Qot_StockFilter.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	StockFilter(req *Qot_StockFilter.Request) uint

	/***
	 * 获取股票代码变化信息，具体字段请参考Qot_GetCodeChange.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetCodeChange(req *Qot_GetCodeChange.Request) uint
	/***
	 * 获取IPO列表，具体字段请参考Qot_GetIpoList.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetIpoList(req *Qot_GetIpoList.Request) uint
	/***
	 * 获取获取期货合约资料, 具体字段请参考Qot_GetFutureInfo.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetFutureInfo(req *Qot_GetFutureInfo.Request) uint
}

type FTAPIConnTrd interface {
	FTAPIConn

	SetTrdSpi(trdSpi FTSPITrd)

	/***
	* 获取交易账户列表，具体字段请参考Trd_GetAccList.proto协议
	* @param req
	* @return 请求的序列号
	*/
	GetAccList(req *Trd_GetAccList.Request) uint
	/***
	 * 解锁，具体字段请参考Trd_UnlockTrade.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	UnlockTrade(req *Trd_UnlockTrade.Request) uint
	/***
	 * 订阅接收推送数据的交易账户，具体字段请参考Trd_SubAccPush.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	SubAccPush(req *Trd_SubAccPush.Request) uint
	/***
	 * 获取账户资金，具体字段请参考Trd_GetFunds.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetFunds(req *Trd_GetFunds.Request) uint

	/***
	 * 获取账户持仓，具体字段请参考Trd_GetPositionList.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetPositionList(req *Trd_GetPositionList.Request) uint

	/***
	 * 获取最大交易数量，具体字段请参考Trd_GetMaxTrdQtys.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetMaxTrdQtys(req *Trd_GetMaxTrdQtys.Request) uint
	/***
	 * 获取当日订单列表，具体字段请参考Trd_GetOrderList.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetOrderList(req *Trd_GetOrderList.Request) uint

	/***
	 * 下单，具体字段请参考Trd_PlaceOrder.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	PlaceOrder(req *Trd_PlaceOrder.Request) uint
	/***
	 * 修改订单，具体字段请参考Trd_ModifyOrder.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	ModifyOrder(req *Trd_ModifyOrder.Request) uint

	/***
	 * 获取当日成交列表，具体字段请参考Trd_GetOrderFillList.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetOrderFillList(req *Trd_GetOrderFillList.Request) uint

	/***
	 * 获取历史订单列表，具体字段请参考Trd_GetHistoryOrderList.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetHistoryOrderList(req *Trd_GetHistoryOrderList.Request) uint
	/***
	 * 获取历史成交列表，具体字段请参考Trd_GetHistoryOrderFillList.proto协议
	 * @param req
	 * @return 请求的序列号
	 */
	GetHistoryOrderFillList(req *Trd_GetHistoryOrderFillList.Request) uint
}

