package futuapi

import (
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
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetCapitalDistribution"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetUserSecurity"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_ModifyUserSecurity"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_StockFilter"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetCodeChange"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetIpoList"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_GetFutureInfo"
	"github.com/stephenlyu/gofutuapi/futuproto/Notify"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateOrderBook"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateBasicQot"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateKL"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateRT"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateTicker"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateBroker"
)

type FTSPIQotBase struct {
}

func (spi *FTSPIQotBase) OnReply_GetGlobalState(client FTAPIConnQot, nSerialNo uint, resp *GetGlobalState.Response) {} //获取全局状态回调
func (spi *FTSPIQotBase) OnReply_Sub(client FTAPIConnQot, nSerialNo uint, rsp *Qot_Sub.Response) {}//订阅或者反订阅回调
func (spi *FTSPIQotBase) OnReply_RegQotPush(client FTAPIConnQot, nSerialNo uint, rsp *Qot_RegQotPush.Response) {} //注册推送回调
func (spi *FTSPIQotBase) OnReply_GetSubInfo(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetSubInfo.Response) {} //获取订阅信息回调
func (spi *FTSPIQotBase) OnReply_GetTicker(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetTicker.Response) {} //获取逐笔,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Ticker)回调
func (spi *FTSPIQotBase) OnReply_GetBasicQot(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetBasicQot.Response) {} //获取基本行情,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Basic)回调
func (spi *FTSPIQotBase) OnReply_GetOrderBook(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetOrderBook.Response) {} //获取摆盘,调用该接口前需要先订阅(订阅位：rsp *Qot__Common.SubType_OrderBook)回调
func (spi *FTSPIQotBase) OnReply_GetKL(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetKL.Response) {} //获取K线，调用该接口前需要先订阅(订阅位：rsp *Qot__Common.SubType_KL_XXX)回调
func (spi *FTSPIQotBase) OnReply_GetRT(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetRT.Response) {} //获取分时，调用该接口前需要先订阅(订阅位：rsp *Qot__Common.SubType_RT)回调
func (spi *FTSPIQotBase) OnReply_GetBroker(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetBroker.Response) {} //获取经纪队列，调用该接口前需要先订阅(订阅位：rsp *Qot__Common.SubType_Broker)回调
//OnReply_GetHistoryKL(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetHistoryKL.Response){} //获取本地历史K线回调
//OnReply_GetHistoryKLPoints(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetHistoryKLPoints.Response){} //获取多股票多点本地历史K线回调
//OnReply_GetRehab(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetRehab.Response){} //获取本地历史复权信息回调
func (spi *FTSPIQotBase) OnReply_RequestRehab(client FTAPIConnQot, nSerialNo uint, rsp *Qot_RequestRehab.Response) {} //在线请求历史复权信息，不读本地历史数据DB回调
func (spi *FTSPIQotBase) OnReply_RequestHistoryKL(client FTAPIConnQot, nSerialNo uint, rsp *Qot_RequestHistoryKL.Response) {} //在线请求历史K线，不读本地历史数据DB回调
func (spi *FTSPIQotBase) OnReply_RequestHistoryKLQuota(client FTAPIConnQot, nSerialNo uint, rsp *Qot_RequestHistoryKLQuota.Response) {} //获取历史K线已经用掉的额度回调
func (spi *FTSPIQotBase) OnReply_GetTradeDate(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetTradeDate.Response) {} //获取交易日回调
func (spi *FTSPIQotBase) OnReply_GetStaticInfo(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetStaticInfo.Response) {} //获取静态信息回调
func (spi *FTSPIQotBase) OnReply_GetSecuritySnapshot(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetSecuritySnapshot.Response) {} //获取股票快照回调
func (spi *FTSPIQotBase) OnReply_GetPlateSet(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetPlateSet.Response) {} //获取板块集合下的板块回调
func (spi *FTSPIQotBase) OnReply_GetPlateSecurity(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetPlateSecurity.Response) {} //获取板块下的股票回调
func (spi *FTSPIQotBase) OnReply_GetReference(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetReference.Response) {} //获取相关股票回调
func (spi *FTSPIQotBase) OnReply_GetOwnerPlate(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetOwnerPlate.Response) {} //获取股票所属的板块回调
func (spi *FTSPIQotBase) OnReply_GetHoldingChangeList(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetHoldingChangeList.Response) {} //获取大股东持股变化列表回调
func (spi *FTSPIQotBase) OnReply_GetOptionChain(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetOptionChain.Response) {} //筛选期权回调
func (spi *FTSPIQotBase) OnReply_GetWarrant(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetWarrant.Response) {} //筛选窝轮回调
func (spi *FTSPIQotBase) OnReply_GetCapitalFlow(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetCapitalFlow.Response) {} //获取资金流向回调
func (spi *FTSPIQotBase) OnReply_GetCapitalDistribution(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetCapitalDistribution.Response) {} //获取资金分布回调
func (spi *FTSPIQotBase) OnReply_GetUserSecurity(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetUserSecurity.Response) {} //获取自选股分组下的股票回调
func (spi *FTSPIQotBase) OnReply_ModifyUserSecurity(client FTAPIConnQot, nSerialNo uint, rsp *Qot_ModifyUserSecurity.Response) {} //修改自选股分组下的股票回调
func (spi *FTSPIQotBase) OnReply_StockFilter(client FTAPIConnQot, nSerialNo uint, rsp *Qot_StockFilter.Response) {} //条件选股
func (spi *FTSPIQotBase) OnReply_GetCodeChange(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetCodeChange.Response) {} //获取股票代码变化信息
func (spi *FTSPIQotBase) OnReply_GetIpoList(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetIpoList.Response) {} //获取IPO列表
func (spi *FTSPIQotBase) OnReply_GetFutureInfo(client FTAPIConnQot, nSerialNo uint, rsp *Qot_GetFutureInfo.Response) {} //获取期货合约资料
func (spi *FTSPIQotBase) OnPush_Notify(client FTAPIConnQot, rsp *Notify.Response) {}//推送系统通知
func (spi *FTSPIQotBase) OnPush_UpdateOrderBook(client FTAPIConnQot, rsp *Qot_UpdateOrderBook.Response) {} //推送摆盘
func (spi *FTSPIQotBase) OnPush_UpdateBasicQuote(client FTAPIConnQot, rsp *Qot_UpdateBasicQot.Response) {} //推送报价
func (spi *FTSPIQotBase) OnPush_UpdateKL(client FTAPIConnQot, rsp *Qot_UpdateKL.Response) {} //推送K线
func (spi *FTSPIQotBase) OnPush_UpdateRT(client FTAPIConnQot, rsp *Qot_UpdateRT.Response) {} //推送分时
func (spi *FTSPIQotBase) OnPush_UpdateTicker(client FTAPIConnQot, rsp *Qot_UpdateTicker.Response) {} //推送逐笔
func (spi *FTSPIQotBase) OnPush_UpdateBroker(client FTAPIConnQot, rsp *Qot_UpdateBroker.Response) {}