#pragma once
#include "FTAPI_Define.h"

namespace Futu {
class FTAPI_Conn;
class FTSPI_Conn
{
public:
	/**
	* @brief 初始化连接
	* @praram pConn 对应连接实例指针
	* @param nErrCode 当高32位为FTAPI_ConnectFailType类型，则低32位为系统错误码；
	*				  当高32位为FTAPI_InitFail，则低32位为FTAPI_InitFailType类型。
	* @param strDesc 错误描述
	*/
	virtual void OnInitConnect(FTAPI_Conn* pConn, Futu::i64_t nErrCode, const char* strDesc) = 0;
	/**
	* @brief 连接断开
	* @praram pConn 对应连接实例指针
	* @param nErrCode 高32位为FTAPI_ConnectFailType类型，低32位为系统错误码；
	*/
	virtual void OnDisConnect(FTAPI_Conn* pConn, Futu::i64_t nErrCode) = 0;
};

class FTSPI_Qot
{
public:
	/**
	* @brief 请求全局状态
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考GetGlobalState.proto协议
	*/
	virtual void OnReply_GetGlobalState(Futu::u32_t nSerialNo, const GetGlobalState::Response &stRsp) = 0;
	/**
	* @brief 订阅
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_Sub.proto协议
	*/
	virtual void OnReply_Sub(Futu::u32_t nSerialNo, const Qot_Sub::Response &stRsp) = 0;
	/**
	* @brief 注册推送
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_RegQotPush.proto协议
	*/
	virtual void OnReply_RegQotPush(Futu::u32_t nSerialNo, const Qot_RegQotPush::Response &stRsp) = 0;
	/**
	* @brief 获取订阅信息
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetSubInfo.proto协议
	*/
	virtual void OnReply_GetSubInfo(Futu::u32_t nSerialNo, const Qot_GetSubInfo::Response &stRsp) = 0;
	/**
	* @brief 获取逐笔
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetTicker.proto协议
	*/
	virtual void OnReply_GetTicker(Futu::u32_t nSerialNo, const Qot_GetTicker::Response &stRsp) = 0;
	/**
	* @brief 获取报价
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetBasicQot.proto协议
	*/
	virtual void OnReply_GetBasicQot(Futu::u32_t nSerialNo, const Qot_GetBasicQot::Response &stRsp) = 0;
	/**
	* @brief 获取摆盘
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetOrderBook.proto协议
	*/
	virtual void OnReply_GetOrderBook(Futu::u32_t nSerialNo, const Qot_GetOrderBook::Response &stRsp) = 0;
	/**
	* @brief 获取K线
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetKL.proto协议
	*/
	virtual void OnReply_GetKL(Futu::u32_t nSerialNo, const Qot_GetKL::Response &stRsp) = 0;
	/**
	* @brief 获取分时
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetRT.proto协议
	*/
	virtual void OnReply_GetRT(Futu::u32_t nSerialNo, const Qot_GetRT::Response &stRsp) = 0;
	/**
	* @brief 获取经纪队列
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetBroker.proto协议
	*/
	virtual void OnReply_GetBroker(Futu::u32_t nSerialNo, const Qot_GetBroker::Response &stRsp) = 0;
	/**
	* @brief 在线请求历史复权信息，不读本地历史数据DB
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_RequestRehab.proto协议
	*/
	virtual void OnReply_RequestRehab(Futu::u32_t nSerialNo, const Qot_RequestRehab::Response &stRsp) = 0;
	/**
	* @brief 在线请求历史K线，不读本地历史数据DB
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_RequestHistoryKL.proto协议
	*/
	virtual void OnReply_RequestHistoryKL(Futu::u32_t nSerialNo, const Qot_RequestHistoryKL::Response &stRsp) = 0;
	/**
	* @brief 获取历史K线已经用掉的额度 
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_RequestHistoryKLQuota.proto协议
	*/
	virtual void OnReply_RequestHistoryKLQuota(Futu::u32_t nSerialNo, const Qot_RequestHistoryKLQuota::Response &stRsp) = 0;
	/**
	* @brief 获取交易日
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetTradeDate.proto协议
	*/
	virtual void OnReply_GetTradeDate(Futu::u32_t nSerialNo, const Qot_GetTradeDate::Response &stRsp) = 0;
	/**
	* @brief 获取静态信息
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetStaticInfo.proto协议
	*/
	virtual void OnReply_GetStaticInfo(Futu::u32_t nSerialNo, const Qot_GetStaticInfo::Response &stRsp) = 0;
	/**
	* @brief 获取股票快照
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetSecuritySnapshot.proto协议
	*/
	virtual void OnReply_GetSecuritySnapshot(Futu::u32_t nSerialNo, const Qot_GetSecuritySnapshot::Response &stRsp) = 0;
	/**
	* @brief 获取板块集合下的板块
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetPlateSet.proto协议
	*/
	virtual void OnReply_GetPlateSet(Futu::u32_t nSerialNo, const Qot_GetPlateSet::Response &stRsp) = 0;
	/**
	* @brief 获取板块下的股票
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetPlateSecurity.proto协议
	*/
	virtual void OnReply_GetPlateSecurity(Futu::u32_t nSerialNo, const Qot_GetPlateSecurity::Response &stRsp) = 0;
	/**
	* @brief 获取相关股票
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetReference.proto协议
	*/
	virtual void OnReply_GetReference(Futu::u32_t nSerialNo, const Qot_GetReference::Response &stRsp) = 0;
	/**
	* @brief 获取股票所属的板块
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetOwnerPlate.proto协议
	*/
	virtual void OnReply_GetOwnerPlate(Futu::u32_t nSerialNo, const Qot_GetOwnerPlate::Response &stRsp) = 0;
	/**
	* @brief 获取大股东持股变化列表
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetHoldingChangeList.proto协议
	*/
	virtual void OnReply_GetHoldingChangeList(Futu::u32_t nSerialNo, const Qot_GetHoldingChangeList::Response &stRsp) = 0;
	/**
	* @brief 筛选期权
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetOptionChain.proto协议
	*/
	virtual void OnReply_GetOptionChain(Futu::u32_t nSerialNo, const Qot_GetOptionChain::Response &stRsp) = 0;
	/**
	* @brief 筛选窝轮
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetWarrant.proto协议
	*/
	virtual void OnReply_GetWarrant(Futu::u32_t nSerialNo, const Qot_GetWarrant::Response &stRsp) = 0;
	/**
	* @brief 获取资金流向
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetCapitalFlow.proto协议
	*/
	virtual void OnReply_GetCapitalFlow(Futu::u32_t nSerialNo, const Qot_GetCapitalFlow::Response &stRsp) = 0;
	/**
	* @brief 获取资金分布
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetCapitalDistribution.proto协议
	*/
	virtual void OnReply_GetCapitalDistribution(Futu::u32_t nSerialNo, const Qot_GetCapitalDistribution::Response &stRsp) = 0;
	/**
	* @brief 获取自选股分组下的股票
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetUserSecurity.proto协议
	*/
	virtual void OnReply_GetUserSecurity(Futu::u32_t nSerialNo, const Qot_GetUserSecurity::Response &stRsp) = 0;
	/**
	* @brief 修改自选股分组下的股票
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_ModifyUserSecurity.proto协议
	*/
	virtual void OnReply_ModifyUserSecurity(Futu::u32_t nSerialNo, const Qot_ModifyUserSecurity::Response &stRsp) = 0;
	
	/**
	* @brief 条件选股
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_StockFilter.proto协议
	*/
	virtual void OnReply_StockFilter(Futu::u32_t nSerialNo, const Qot_StockFilter::Response &stRsp) = 0;
	/**
	* @brief 获取股票代码变化信息
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetCodeChange.proto协议
	*/
	virtual void OnReply_GetCodeChange(Futu::u32_t nSerialNo, const Qot_GetCodeChange::Response &stRsp) = 0;
	/**
	* @brief 新股IPO
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetIpoList.proto协议
	*/
	virtual void OnReply_GetIpoList(Futu::u32_t nSerialNo, const Qot_GetIpoList::Response &stRsp) = 0;
	/**
	* @brief 期货合约资料
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_GetFutureInfo.proto协议
	*/
	virtual void OnReply_GetFutureInfo(Futu::u32_t nSerialNo, const Qot_GetFutureInfo::Response &stRsp) = 0;

	/**
	* @brief 推送通知
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Notify.proto协议
	*/
	virtual void OnPush_Notify(const Notify::Response &stRsp) = 0;
	/**
	* @brief 推送逐笔
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_UpdateTicker.proto协议
	*/
	virtual void OnPush_UpdateTicker(const Qot_UpdateTicker::Response &stRsp) = 0;
	/**
	* @brief 推送报价
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_UpdateBasicQot.proto协议
	*/
	virtual void OnPush_UpdateBasicQot(const Qot_UpdateBasicQot::Response &stRsp) = 0;
	/**
	* @brief 推送摆盘
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_UpdateOrderBook.proto协议
	*/
	virtual void OnPush_UpdateOrderBook(const Qot_UpdateOrderBook::Response &stRsp) = 0;
	/**
	* @brief 推送K线
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_UpdateKL.proto协议
	*/
	virtual void OnPush_UpdateKL(const Qot_UpdateKL::Response &stRsp) = 0;
	/**
	* @brief 推送分时
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_UpdateRT.proto协议
	*/
	virtual void OnPush_UpdateRT(const Qot_UpdateRT::Response &stRsp) = 0;
	/**
	* @brief 推送经纪队列
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Qot_UpdateBroker.proto协议
	*/
	virtual void OnPush_UpdateBroker(const Qot_UpdateBroker::Response &stRsp) = 0;
};

class FTSPI_Trd
{
public:
	/**
	* @brief 获取交易账户列表
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_GetAccList.proto协议
	*/
	virtual void OnReply_GetAccList(Futu::u32_t nSerialNo, const Trd_GetAccList::Response &stRsp) = 0;
	/**
	* @brief 解锁
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_UnlockTrade.proto协议
	*/
	virtual void OnReply_UnlockTrade(Futu::u32_t nSerialNo, const Trd_UnlockTrade::Response &stRsp) = 0;
	/**
	* @brief 订阅接收推送数据的交易账户
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_SubAccPush.proto协议
	*/
	virtual void OnReply_SubAccPush(Futu::u32_t nSerialNo, const Trd_SubAccPush::Response &stRsp) = 0;
	/**
	* @brief 获取账户资金
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_GetFunds.proto协议
	*/
	virtual void OnReply_GetFunds(Futu::u32_t nSerialNo, const Trd_GetFunds::Response &stRsp) = 0;
	/**
	* @brief 获取账户持仓
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_GetPositionList.proto协议
	*/
	virtual void OnReply_GetPositionList(Futu::u32_t nSerialNo, const Trd_GetPositionList::Response &stRsp) = 0;
	/**
	* @brief 获取最大交易数量
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_GetMaxTrdQtys.proto协议
	*/
	virtual void OnReply_GetMaxTrdQtys(Futu::u32_t nSerialNo, const Trd_GetMaxTrdQtys::Response &stRsp) = 0;
	/**
	* @brief 获取当日订单列表
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_GetOrderList.proto协议
	*/
	virtual void OnReply_GetOrderList(Futu::u32_t nSerialNo, const Trd_GetOrderList::Response &stRsp) = 0;
	/**
	* @brief 下单
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_PlaceOrder.proto协议
	*/
	virtual void OnReply_PlaceOrder(Futu::u32_t nSerialNo, const Trd_PlaceOrder::Response &stRsp) = 0;
	/**
	* @brief 修改订单
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_ModifyOrder.proto协议
	*/
	virtual void OnReply_ModifyOrder(Futu::u32_t nSerialNo, const Trd_ModifyOrder::Response &stRsp) = 0;
	/**
	* @brief 获取当日成交列表
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_GetOrderFillList.proto协议
	*/
	virtual void OnReply_GetOrderFillList(Futu::u32_t nSerialNo, const Trd_GetOrderFillList::Response &stRsp) = 0;
	/**
	* @brief 获取历史订单列表
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_GetHistoryOrderList.proto协议
	*/
	virtual void OnReply_GetHistoryOrderList(Futu::u32_t nSerialNo, const Trd_GetHistoryOrderList::Response &stRsp) = 0;
	/**
	* @brief 获取历史成交列表
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_GetHistoryOrderFillList.proto协议
	*/
	virtual void OnReply_GetHistoryOrderFillList(Futu::u32_t nSerialNo, const Trd_GetHistoryOrderFillList::Response &stRsp) = 0;
	/**
	* @brief 推送订单变化
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_UpdateOrder.proto协议
	*/
	virtual void OnPush_UpdateOrder(const Trd_UpdateOrder::Response &stRsp) = 0;
	/**
	* @brief 订单成交推送
	* @praram nSerialNo 包序列号
	* @param stRsp 回包，具体字段请参考Trd_UpdateOrderFill.proto协议
	*/
	virtual void OnPush_UpdateOrderFill(const Trd_UpdateOrderFill::Response &stRsp) = 0;
};
}