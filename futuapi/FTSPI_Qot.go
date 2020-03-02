package futuapi

import (
	"github.com/stephenlyu/gofutuapi/futuproto/GetGlobalState"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_RegQotPush"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_Sub"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetSubInfo"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetTicker"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetBasicQot"
	"github.com/stephenlyu/gofutuapi/futuproto/Notify"
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
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetCapitalDistribution"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetUserSecurity"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_ModifyUserSecurity"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_StockFilter"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetCodeChange"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetIpoList"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetFutureInfo"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateOrderBook"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateKL"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateTicker"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateBroker"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateRT"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateBasicQot"
)

type FTSPIQot interface {
	OnReply_GetGlobalState(client *FTAPIConn, nSerialNo uint, resp *GetGlobalState.Response) //获取全局状态回调
	OnReply_Sub(client *FTAPIConn, nSerialNo uint, rsp *Qot_Sub.Response)FTAPIConn //订阅或者反订阅回调
	OnReply_RegQotPush(client *FTAPIConn, nSerialNo uint, rsp *Qot_RegQotPush.Response)FTAPIConn //注册推送回调
	OnReply_GetSubInfo(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetSubInfo.Response)FTAPIConn //获取订阅信息回调
	OnReply_GetTicker(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetTicker.Response)FTAPIConn //获取逐笔,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Ticker)回调
	OnReply_GetBasicQot(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetBasicQot.Response)FTAPIConn //获取基本行情,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Basic)回调
	OnReply_GetOrderBook(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetOrderBook.Response)FTAPIConn //获取摆盘,调用该接口前需要先订阅(订阅位：rsp *Qot__Common.SubType_OrderBook)回调
	OnReply_GetKL(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetKL.Response)FTAPIConn //获取K线，调用该接口前需要先订阅(订阅位：rsp *Qot__Common.SubType_KL_XXX)回调
	OnReply_GetRT(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetRT.Response)FTAPIConn //获取分时，调用该接口前需要先订阅(订阅位：rsp *Qot__Common.SubType_RT)回调
	OnReply_GetBroker(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetBroker.Response)FTAPIConn //获取经纪队列，调用该接口前需要先订阅(订阅位：rsp *Qot__Common.SubType_Broker)回调
	//OnReply_GetHistoryKL(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetHistoryKL.Response)FTAPIConn //获取本地历史K线回调
	//OnReply_GetHistoryKLPoints(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetHistoryKLPoints.Response)FTAPIConn //获取多股票多点本地历史K线回调
	//OnReply_GetRehab(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetRehab.Response)FTAPIConn //获取本地历史复权信息回调
	OnReply_RequestRehab(client *FTAPIConn, nSerialNo uint, rsp *Qot_RequestRehab.Response)FTAPIConn //在线请求历史复权信息，不读本地历史数据DB回调
	OnReply_RequestHistoryKL(client *FTAPIConn, nSerialNo uint, rsp *Qot_RequestHistoryKL.Response)FTAPIConn //在线请求历史K线，不读本地历史数据DB回调
	OnReply_RequestHistoryKLQuota(client *FTAPIConn, nSerialNo uint, rsp *Qot_RequestHistoryKLQuota.Response)FTAPIConn //获取历史K线已经用掉的额度回调
	OnReply_GetTradeDate(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetTradeDate.Response)FTAPIConn //获取交易日回调
	OnReply_GetStaticInfo(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetStaticInfo.Response)FTAPIConn //获取静态信息回调
	OnReply_GetSecuritySnapshot(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetSecuritySnapshot.Response)FTAPIConn //获取股票快照回调
	OnReply_GetPlateSet(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetPlateSet.Response)FTAPIConn //获取板块集合下的板块回调
	OnReply_GetPlateSecurity(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetPlateSecurity.Response)FTAPIConn //获取板块下的股票回调
	OnReply_GetReference(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetReference.Response)FTAPIConn //获取相关股票回调
	OnReply_GetOwnerPlate(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetOwnerPlate.Response)FTAPIConn //获取股票所属的板块回调
	OnReply_GetHoldingChangeList(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetHoldingChangeList.Response)FTAPIConn //获取大股东持股变化列表回调
	OnReply_GetOptionChain(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetOptionChain.Response)FTAPIConn //筛选期权回调
	OnReply_GetWarrant(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetWarrant.Response)FTAPIConn //筛选窝轮回调
	OnReply_GetCapitalFlow(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetCapitalFlow.Response)FTAPIConn //获取资金流向回调
	OnReply_GetCapitalDistribution(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetCapitalDistribution.Response)FTAPIConn //获取资金分布回调
	OnReply_GetUserSecurity(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetUserSecurity.Response)FTAPIConn //获取自选股分组下的股票回调
	OnReply_ModifyUserSecurity(client *FTAPIConn, nSerialNo uint, rsp *Qot_ModifyUserSecurity.Response)FTAPIConn //修改自选股分组下的股票回调
	OnReply_StockFilter(client *FTAPIConn, nSerialNo uint, rsp *Qot_StockFilter.Response)FTAPIConn //条件选股
	OnReply_GetCodeChange(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetCodeChange.Response)FTAPIConn //获取股票代码变化信息
	OnReply_GetIpoList(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetIpoList.Response)FTAPIConn //获取IPO列表
	OnReply_GetFutureInfo(client *FTAPIConn, nSerialNo uint, rsp *Qot_GetFutureInfo.Response)FTAPIConn //获取期货合约资料
	OnPush_Notify(client *FTAPIConn, rsp *Notify.Response)FTAPIConn  //推送系统通知
	OnPush_UpdateOrderBook(client *FTAPIConn, rsp *Qot_UpdateOrderBook.Response)FTAPIConn //推送摆盘
	OnPush_UpdateBasicQuote(client *FTAPIConn, rsp *Qot_UpdateBasicQot.Response)FTAPIConn //推送报价
	OnPush_UpdateKL(client *FTAPIConn, rsp *Qot_UpdateKL.Response)FTAPIConn //推送K线
	OnPush_UpdateRT(client *FTAPIConn, rsp *Qot_UpdateRT.Response)FTAPIConn //推送分时
	OnPush_UpdateTicker(client *FTAPIConn, rsp *Qot_UpdateTicker.Response)FTAPIConn //推送逐笔
	OnPush_UpdateBroker(client *FTAPIConn, rsp *Qot_UpdateBroker.Response)FTAPIConn //推送经纪队列
}