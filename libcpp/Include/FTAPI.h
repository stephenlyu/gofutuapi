#pragma once
#include "FTAPI_Define.h"

class FTAPIImp;
namespace Futu {
class FTSPI_Conn;
class FTSPI_Qot;
class FTSPI_Trd;
class  FTAPI_Conn
{
public:
	FTAPI_Conn();
	virtual ~FTAPI_Conn();

public:
	/**
	* @brief 设置客户端信息
	* @param [in] strClientID 客户端标识
	* @param [in] nClientVer 客户端版本
	*/
	void SetClientInfo(const char* szClientID, Futu::i32_t nClientVer);

	/**
	* @brief 设置密钥
	* @param [in] strRSAPrivateKey 密钥
	*/
	void SetRSAPrivateKey(const char* szRSAPrivateKey);
	
	/**
	* @brief 初始化连接，连接并初始化
	* @param [in] szIPAddr 地址
	* @param [in] nPort 端口
	* @param [in] bEnableEncrypt 启用加密
	* @return bool 是否启动了执行，不代表连接结果，结果通过OnInitConnect回调
	*/
	bool InitConnect(const char* szIPAddr, Futu::u16_t nPort, bool bEnableEncrypt);

	/**
	* @brief 此连接的连接ID，连接的唯一标识，InitConnect协议返回，没有初始化前为0
	* @return 连接ID
	*/
	Futu::u64_t GetConnectID();

	/**
	* @brief 关闭连接
	* @return 是否成功
	*/
	bool Close();

	/**
	* @brief 注册回调
	* @param [in] pSpi 回调，该对象没有反注册前不可销毁
	*/
	void RegisterConnSpi(FTSPI_Conn* pSpi);

	/**
	* @brief 反注册回调
	*/
	void UnregisterConnSpi();

protected:
	FTAPIImp *m_pImp{ nullptr };
};

/**
* @brief 行情接口
*/
class  FTAPI_Qot : public FTAPI_Conn
{
public:
	FTAPI_Qot();
	virtual ~FTAPI_Qot();

public:
	/**
	* @brief 注册回调
	* @param [in] pSpi 回调，该对象没有反注册前不可销毁
	*/
	void RegisterQotSpi(FTSPI_Qot* pSpi);

	/**
	* @brief 反注册回调
	*/
	void UnregisterQotSpi();

	/**
	* @brief 订阅，反订阅
	* @param [in] stReq 请求包，具体字段请参考Qot_Sub.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t Sub(const Qot_Sub::Request &stReq);
	/**
	* @brief 获取全局信息
	* @param [in] stReq 请求包，具体字段请参考GetGlobalState.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetGlobalState(const GetGlobalState::Request &stReq);
	/**
	* @brief 注册推送，用于当前连接订阅，但是在其他连接获取推送的场景
	* @param [in] stReq 请求包，具体字段请参考Qot_RegQotPush.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t RegQotPush(const Qot_RegQotPush::Request &stReq);
	/**
	* @brief 获取订阅信息
	* @param [in] stReq 请求包，具体字段请参考Qot_GetSubInfo.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetSubInfo(const Qot_GetSubInfo::Request &stReq);
	/**
	* @brief 获取逐笔，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Ticker)
	* @param [in] stReq 请求包，具体字段请参考Qot_GetTicker.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetTicker(const Qot_GetTicker::Request &stReq);
	/**
	* @brief 获取报价，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Basic)
	* @param [in] stReq 请求包，具体字段请参考Qot_GetBasicQot.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetBasicQot(const Qot_GetBasicQot::Request &stReq);
	/**
	* @brief 获取摆盘，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_OrderBook)
	* @param [in] stReq 请求包，具体字段请参考Qot_GetOrderBook.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetOrderBook(const Qot_GetOrderBook::Request &stReq);
	/**
	* @brief 获取K线，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_KL_XXX)
	* @param [in] stReq 请求包，具体字段请参考Qot_GetKL.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetKL(const Qot_GetKL::Request &stReq);
	/**
	* @brief 获取分时，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_RT)
	* @param [in] stReq 请求包，具体字段请参考GetRT.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetRT(const Qot_GetRT::Request &stReq);
	/**
	* @brief 获取经纪队列，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Broker)
	* @param [in] stReq 请求包，具体字段请参考Qot_GetBroker.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetBroker(const Qot_GetBroker::Request &stReq);
	/**
	* @brief 在线请求历史复权信息，不读本地历史数据DB
	* @param [in] stReq 请求包，具体字段请参考Qot_RequestRehab.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t RequestRehab(const Qot_RequestRehab::Request &stReq);
	/**
	* @brief 在线请求历史K线，不读本地历史数据DB
	* @param [in] stReq 请求包，具体字段请参考Qot_RequestHistoryKL.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t RequestHistoryKL(const Qot_RequestHistoryKL::Request &stReq);
	/**
	* @brief 获取历史K线已经用掉的额度 
	* @param [in] stReq 请求包，具体字段请参考Qot_RequestHistoryKLQuota.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t RequestHistoryKLQuota(const Qot_RequestHistoryKLQuota::Request &stReq);
	/**
	* @brief 获取交易日
	* @param [in] stReq 请求包，具体字段请参考Qot_GetTradeDate.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetTradeDate(const Qot_GetTradeDate::Request &stReq);
	/**
	* @brief 获取静态信息
	* @param [in] stReq 请求包，具体字段请参考Qot_GetStaticInfo.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetStaticInfo(const Qot_GetStaticInfo::Request &stReq);
	/**
	* @brief 获取股票快照
	* @param [in] stReq 请求包，具体字段请参考Qot_GetSecuritySnapshot.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetSecuritySnapshot(const Qot_GetSecuritySnapshot::Request &stReq);
	/**
	* @brief 获取板块集合下的板块
	* @param [in] stReq 请求包，具体字段请参考Qot_GetPlateSet.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetPlateSet(const Qot_GetPlateSet::Request &stReq);
	/**
	* @brief 获取板块下的股票
	* @param [in] stReq 请求包，具体字段请参考Qot_GetPlateSecurity.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetPlateSecurity(const Qot_GetPlateSecurity::Request &stReq);
	/**
	* @brief 获取相关股票
	* @param [in] stReq 请求包，具体字段请参考Qot_GetReference.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetReference(const Qot_GetReference::Request &stReq);
	/**
	* @brief 获取股票所属的板块
	* @param [in] stReq 请求包，具体字段请参考Qot_GetOwnerPlate.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetOwnerPlate(const Qot_GetOwnerPlate::Request &stReq);
	/**
	* @brief 获取大股东持股变化列表
	* @param [in] stReq 请求包，具体字段请参考Qot_GetHoldingChangeList.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetHoldingChangeList(const Qot_GetHoldingChangeList::Request &stReq);
	/**
	* @brief 筛选期权
	* @param [in] stReq 请求包，具体字段请参考Qot_GetOptionChain.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetOptionChain(const Qot_GetOptionChain::Request &stReq);
	/**
	* @brief 筛选窝轮
	* @param [in] stReq 请求包，具体字段请参考Qot_GetWarrant.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetWarrant(const Qot_GetWarrant::Request &stReq);
	/**
	* @brief 获取资金流向
	* @param [in] stReq 请求包，具体字段请参考Qot_GetCapitalFlow.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetCapitalFlow(const Qot_GetCapitalFlow::Request &stReq);
	/**
	* @brief 获取资金分布
	* @param [in] stReq 请求包，具体字段请参考Qot_GetCapitalDistribution.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetCapitalDistribution(const Qot_GetCapitalDistribution::Request &stReq);
	/**
	* @brief 获取自选股分组下的股票
	* @param [in] stReq 请求包，具体字段请参考Qot_GetUserSecurity.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetUserSecurity(const Qot_GetUserSecurity::Request &stReq);
	/**
	* @brief 修改自选股分组下的股票
	* @param [in] stReq 请求包，具体字段请参考Qot_ModifyUserSecurity.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t ModifyUserSecurity(const Qot_ModifyUserSecurity::Request &stReq);
	/**
	* @brief 条件选股
	* @param [in] stReq 请求包，具体字段请参考Qot_StockFilter.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t StockFilter(const Qot_StockFilter::Request &stReq);

	/**
	* @brief 获取股票代码变化
	* @param [in] stReq 请求包，具体字段请参考Qot_GetCodeChange.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetCodeChange(const Qot_GetCodeChange::Request &stReq);

	/**
	* @brief 新股IPO
	* @param [in] stReq 请求包，具体字段请参考Qot_GetIpoList.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetIpoList(const Qot_GetIpoList::Request &stReq);

	/**
	* @brief 期货合约资料
	* @param [in] stReq 请求包，具体字段请参考Qot_GetFutureInfo.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetFutureInfo(const Qot_GetFutureInfo::Request &stReq);
	
};

/**
* @brief 交易接口
*/
class  FTAPI_Trd : public FTAPI_Conn
{
public:
	FTAPI_Trd();
	virtual ~FTAPI_Trd();

public:
	/**
	* @brief 注册回调
	* @param [in] pSpi 回调，该对象没有反注册前不可销毁
	*/
	void RegisterTrdSpi(FTSPI_Trd* pSpi);

	/**
	* @brief 反注册回调
	*/
	void UnregisterTrdSpi();

	/**
	* @brief 获取交易包标识
	* @param [out] 获取到的交易包标识
	* @return bool  是否成功，没有初始化前调用会返回失败
	*/
	bool GetPacketID(Common::PacketID &stPacketID);
	
	/**
	* @brief 获取交易帐号列表
	* @param [in] stReq 请求包，具体字段请参考Trd_GetAccList.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetAccList(const Trd_GetAccList::Request &stReq);
	/**
	* @brief 解锁，针对OpenD解锁一次即可
	* @param [in] stReq 请求包，具体字段请参考Trd_UnlockTrade.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t UnlockTrade(const Trd_UnlockTrade::Request &stReq);
	/**
	* @brief 订阅接收推送数据的交易账户
	* @param [in] stReq 请求包，具体字段请参考Trd_SubAccPush.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t SubAccPush(const Trd_SubAccPush::Request &stReq);
	/**
	* @brief 获取账户资金
	* @param [in] stReq 请求包，具体字段请参考Trd_GetFunds.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetFunds(const Trd_GetFunds::Request &stReq);
	/**
	* @brief 获取账户持仓
	* @param [in] stReq 请求包，具体字段请参考Trd_GetPositionList.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetPositionList(const Trd_GetPositionList::Request &stReq);
	/**
	* @brief 获取最大交易数量
	* @param [in] stReq 请求包，具体字段请参考Trd_GetMaxTrdQtys.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetMaxTrdQtys(const Trd_GetMaxTrdQtys::Request &stReq);
	/**
	* @brief 获取当日订单列表
	* @param [in] stReq 请求包，具体字段请参考Trd_GetOrderList.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetOrderList(const Trd_GetOrderList::Request &stReq);
	/**
	* @brief 下单
	* @param [in] stReq 请求包，具体字段请参考Trd_PlaceOrder.proto协议，PacketID不需填写，发送时接口会填
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t PlaceOrder(const Trd_PlaceOrder::Request &stReq);
	/**
	* @brief 修改订单
	* @param [in] stReq 请求包，具体字段请参考Trd_ModifyOrder.proto协议，PacketID不需填写，发送时接口会填
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t ModifyOrder(const Trd_ModifyOrder::Request &stReq);
	/**
	* @brief 获取当日成交列表
	* @param [in] stReq 请求包，具体字段请参考Trd_GetOrderFillList.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetOrderFillList(const Trd_GetOrderFillList::Request &stReq);
	/**
	* @brief 获取历史订单列表
	* @param [in] stReq 请求包，具体字段请参考Trd_GetHistoryOrderList.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetHistoryOrderList(const Trd_GetHistoryOrderList::Request &stReq);
	/**
	* @brief 获取历史成交列表
	* @param [in] stReq 请求包，具体字段请参考Trd_GetHistoryOrderFillList.proto协议
	* @return Futu::u32_t 包序列号， 0表示请求发送失败
	*/
	Futu::u32_t GetHistoryOrderFillList(const Trd_GetHistoryOrderFillList::Request &stReq);
};

/**
* @brief 全局工具接口
*/
class FTAPI
{
public:
	static FTAPI_Qot* CreateQotApi();
	static void ReleaseQotApi(FTAPI_Qot* pQot);

	static FTAPI_Trd* CreateTrdApi();
	static void ReleaseTrdApi(FTAPI_Trd* pTrd);

	static void Init();
	static void UnInit();
};

}