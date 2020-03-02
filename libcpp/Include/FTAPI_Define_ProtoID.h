#pragma once

/** 交易、行情号段范围预留定义 */
#define FTAPI_ProtoID_Trd_Begin									2000 /**< 交易相关协议号段开始 */
#define FTAPI_ProtoID_Trd_End									2999 /**< 交易相关协议号段结束 */
#define FTAPI_ProtoID_Qot_Begin									3000 /**< 行情相关协议号段开始 */
#define FTAPI_ProtoID_Qot_End									3999 /**< 行情相关协议号段结束 */

// 全局协议
#define FTAPI_ProtoID_InitConnect								1001 /**< 初始化连接 */
#define FTAPI_ProtoID_GetGlobalState							1002 /**< 获取全局状态 */
#define FTAPI_ProtoID_Notify									1003 /**< 推送通知 */
#define FTAPI_ProtoID_KeepAlive									1004 /**< 心跳 */
#define FTAPI_ProtoID_GetUserInfo								1005 /**< 获取用户信息 */
#define FTAPI_ProtoID_Verification								1006 /**< 请求或输入验证码 */
#define FTAPI_ProtoID_GetDelayStatistics						1007 /**< 获取延迟统计 */

// 交易协议
#define FTAPI_ProtoID_Trd_GetAccList							2001 /**< 获取交易账户列表 */
#define FTAPI_ProtoID_Trd_UnlockTrade							2005 /**< 解锁或锁定交易 */
#define FTAPI_ProtoID_Trd_SubAccPush							2008 /**< 订阅接收推送数据的交易账户 */

#define FTAPI_ProtoID_Trd_GetFunds								2101 /**< 获取账户资金 */
#define FTAPI_ProtoID_Trd_GetPositionList						2102 /**< 获取账户持仓 */

#define FTAPI_ProtoID_Trd_GetMaxTrdQtys							2111 /**< 获取最大交易数量 */

#define FTAPI_ProtoID_Trd_GetOrderList							2201 /**< 获取订单列表 */
#define FTAPI_ProtoID_Trd_PlaceOrder							2202 /**< 下单 */
#define FTAPI_ProtoID_Trd_ModifyOrder							2205 /**< 修改订单 */
#define FTAPI_ProtoID_Trd_UpdateOrder							2208 /**< 订单状态变动通知(推送) */

#define FTAPI_ProtoID_Trd_GetOrderFillList						2211 /**< 获取成交列表 */
#define FTAPI_ProtoID_Trd_UpdateOrderFill						2218 /**< 成交通知(推送) */

#define FTAPI_ProtoID_Trd_GetHistoryOrderList					2221 /**< 获取历史订单列表 */
#define FTAPI_ProtoID_Trd_GetHistoryOrderFillList				2222 /**< 获取历史成交列表 */

// 行情-实时数据协议
#define FTAPI_ProtoID_Qot_Sub									3001 /**< 订阅或者反订阅 */
#define FTAPI_ProtoID_Qot_RegQotPush							3002 /**< 注册推送 */
#define FTAPI_ProtoID_Qot_GetSubInfo							3003 /**< 获取订阅信息 */
#define FTAPI_ProtoID_Qot_GetBasicQot							3004 /**< 获取基本行情 */
#define FTAPI_ProtoID_Qot_UpdateBasicQot						3005 /**< 推送基本行情 */
#define FTAPI_ProtoID_Qot_GetKL									3006 /**< 获取K线 */
#define FTAPI_ProtoID_Qot_UpdateKL								3007 /**< 推送K线 */
#define FTAPI_ProtoID_Qot_GetRT									3008 /**< 获取分时 */
#define FTAPI_ProtoID_Qot_UpdateRT								3009 /**< 推送分时 */
#define FTAPI_ProtoID_Qot_GetTicker								3010 /**< 获取逐笔 */
#define FTAPI_ProtoID_Qot_UpdateTicker							3011 /**< 推送逐笔 */
#define FTAPI_ProtoID_Qot_GetOrderBook							3012 /**< 获取买卖盘 */
#define FTAPI_ProtoID_Qot_UpdateOrderBook						3013 /**< 推送买卖盘 */
#define FTAPI_ProtoID_Qot_GetBroker								3014 /**< 获取经纪队列 */
#define FTAPI_ProtoID_Qot_UpdateBroker							3015 /**< 推送经纪队列 */
#define FTAPI_ProtoID_Qot_GetOrderDetail						3016 /**< 获取委托明细 */
#define FTAPI_ProtoID_Qot_UpdateOrderDetail						3017 /**< 推送委托明细 */

//	行情-历史数据协议
#define FTAPI_ProtoID_Qot_GetHistoryKL							3100 /**< 获取历史K线 */
#define FTAPI_ProtoID_Qot_GetHistoryKLPoints					3101 /**< 获取多只股票历史单点K线 */
#define FTAPI_ProtoID_Qot_GetRehab								3102 /**< 获取复权信息 */
#define FTAPI_ProtoID_Qot_RequestHistoryKL						3103 /**< 拉取历史K线，不读本地历史数据DB */
#define FTAPI_ProtoID_Qot_RequestHistoryKLQuota					3104 /**< 拉取历史K线已经用掉的额度 */
#define FTAPI_ProtoID_Qot_RequestRehab							3105 /**< 拉取复权信息，不读本地历史数据DB */

// 行情-其他数据协议
#define FTAPI_ProtoID_Qot_GetTradeDate							3200 /**< 获取市场交易日 */
#define FTAPI_ProtoID_Qot_GetSuspend							3201 /**< 获取股票停牌信息 */
#define FTAPI_ProtoID_Qot_GetStaticInfo							3202 /**< 获取股票静态信息 */
#define FTAPI_ProtoID_Qot_GetSecuritySnapshot					3203 /**< 获取股票快照 */
#define FTAPI_ProtoID_Qot_GetPlateSet							3204 /**< 获取板块集合下的板块 */
#define FTAPI_ProtoID_Qot_GetPlateSecurity						3205 /**< 获取板块下的股票 */
#define FTAPI_ProtoID_Qot_GetReference							3206 /**< 获取正股相关股票，包括窝轮和期货 */
#define FTAPI_ProtoID_Qot_GetOwnerPlate							3207 /**< 获取股票所属板块 */
#define FTAPI_ProtoID_Qot_GetHoldingChangeList					3208 /**< 获取大股东持股变化列表 */
#define FTAPI_ProtoID_Qot_GetOptionChain						3209 /**< 获取期权链 */
#define FTAPI_ProtoID_Qot_GetWarrant							3210 /**< 获取涡轮 */
#define FTAPI_ProtoID_Qot_GetCapitalFlow						3211 /**< 获取资金流向 */
#define FTAPI_ProtoID_Qot_GetCapitalDistribution				3212 /**< 获取资金分布 */
#define FTAPI_ProtoID_Qot_GetUserSecurity						3213 /**< 获取自选股分组下的股票 */
#define FTAPI_ProtoID_Qot_ModifyUserSecurity					3214 /**< 修改自选股分组下的股票 */
#define FTAPI_ProtoID_Qot_StockFilter							3215 /**< 获取条件选股 */
#define FTAPI_ProtoID_Qot_GetCodeChange							3216 /**< 获取股票代码变化信息 */
#define FTAPI_ProtoID_Qot_GetIpoList							3217 /**< 新股IPO */
#define FTAPI_ProtoID_Qot_GetFutureInfo						    3218 /**< 获取期货合约资料 */
