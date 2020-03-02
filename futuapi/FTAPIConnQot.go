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
	. "github.com/stephenlyu/gofutuapi"
	"github.com/golang/protobuf/proto"
	"github.com/stephenlyu/gofutuapi/futuproto/Common"
	"github.com/stephenlyu/gofutuapi/futuproto/Notify"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateBasicQot"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateBroker"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateKL"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateOrderBook"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateRT"
	"github.com/stephenlyu/gofutuapi/futuproto/Qot_UpdateTicker"
)

type FTAPIConnQot struct {
	*FTAPIConn
	qotSpi FTSPIQot
}

func NewFTAPIConnQot() *FTAPIConnQot {
	conn := new(FTAPIConnQot)
	conn.FTAPIConn = NewFTAPIConn()
	FTAPIChannelSetCallbacks(conn.channel, NewDirectorFTAPIChannel_Callback(conn))

	return conn
}

func (conn *FTAPIConnQot) SetQotSpi(qotSpi FTSPIQot) {
	conn.qotSpi = qotSpi
}

/***
 * 获取全局状态，具体字段请参考GetGlobalState.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetGlobalState(req *GetGlobalState.Request) uint {
	return conn.SendProto(GETGLOBALSTATE, req)
}

/***
 * 订阅或者反订阅，具体字段请参考Qot_Sub.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) Sub(req *Qot_Sub.Request) uint {
	return conn.SendProto(QOT_SUB, req)
}

/***
 * 注册推送，具体字段请参考Qot_RegQotPush.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) RegQotPush(req *Qot_RegQotPush.Request) uint {
	return conn.SendProto(QOT_REGQOTPUSH, req)
}

/***
 * 获取订阅信息，具体字段请参考Qot_GetSubInfo.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetSubInfo(req *Qot_GetSubInfo.Request) uint {
	return conn.SendProto(QOT_GETSUBINFO, req)
}

/***
 * 获取逐笔,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Ticker)，具体字段请参考Qot_GetTicker.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetTicker(req *Qot_GetTicker.Request) uint {
	return conn.SendProto(QOT_GETTICKER, req)
}

/***
 * 获取基本行情,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Basic)，具体字段请参考Qot_GetBasicQot.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetBasicQot(req *Qot_GetBasicQot.Request) uint {
	return conn.SendProto(QOT_GETBASICQOT, req)
}

/***
 * 获取摆盘,调用该接口前需要先订阅(订阅位：Qot_Common.SubType_OrderBook)，具体字段请参考Qot_GetOrderBook.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetOrderBook(req *Qot_GetOrderBook.Request) uint {
	return conn.SendProto(QOT_GETORDERBOOK, req)
}

/***
 * 获取K线，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_KL_XXX)，具体字段请参考Qot_GetKL.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetKL(req *Qot_GetKL.Request) uint {
	return conn.SendProto(QOT_GETKL, req)
}

/***
 * 获取分时，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_RT)，具体字段请参考Qot_GetRT.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetRT(req *Qot_GetRT.Request) uint {
	return conn.SendProto(QOT_GETRT, req)
}

/***
 * 获取经纪队列，调用该接口前需要先订阅(订阅位：Qot_Common.SubType_Broker)，具体字段请参考Qot_GetBroker.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetBroker(req *Qot_GetBroker.Request) uint {
	return conn.SendProto(QOT_GETBROKER, req)
}

/***
 * 在线请求历史复权信息，不读本地历史数据DB，具体字段请参考Qot_RequestRehab.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) RequestRehab(req *Qot_RequestRehab.Request) uint {
	return conn.SendProto(QOT_REQUESTREHAB, req)
}

/***
 * 在线请求历史K线，不读本地历史数据DB，具体字段请参考Qot_RequestHistoryKL.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) RequestHistoryKL(req *Qot_RequestHistoryKL.Request) uint {
	return conn.SendProto(QOT_REQUESTHISTORYKL, req)
}

/***
 * 获取历史K线已经用掉的额度，具体字段请参考Qot_RequestHistoryKLQuota.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) RequestHistoryKLQuota(req *Qot_RequestHistoryKLQuota.Request) uint {
	return conn.SendProto(QOT_REQUESTHISTORYKLQUOTA, req)
}

/***
 * 获取交易日，具体字段请参考Qot_GetTradeDate.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetTradeDate(req *Qot_GetTradeDate.Request) uint {
	return conn.SendProto(QOT_GETTRADEDATE, req)
}

/***
 * 获取静态信息，具体字段请参考Qot_GetStaticInfo.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetStaticInfo(req *Qot_GetStaticInfo.Request) uint {
	return conn.SendProto(QOT_GETSTATICINFO, req)
}

/***
 * 获取股票快照，具体字段请参考Qot_GetSecuritySnapshot.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetSecuritySnapshot(req *Qot_GetSecuritySnapshot.Request) uint {
	return conn.SendProto(QOT_GETSECURITYSNAPSHOT, req)
}

/***
 * 获取板块集合下的板块，具体字段请参考Qot_GetPlateSet.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetPlateSet(req *Qot_GetPlateSet.Request) uint {
	return conn.SendProto(QOT_GETPLATESET, req)
}

/***
 * 获取板块下的股票，具体字段请参考Qot_GetPlateSecurity.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetPlateSecurity(req *Qot_GetPlateSecurity.Request) uint {
	return conn.SendProto(QOT_GETPLATESECURITY, req)
}

/***
 * 获取相关股票，具体字段请参考Qot_GetReference.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetReference(req *Qot_GetReference.Request) uint {
	return conn.SendProto(QOT_GETREFERENCE, req)
}

/***
 * 获取股票所属的板块，具体字段请参考Qot_GetOwnerPlate.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetOwnerPlate(req *Qot_GetOwnerPlate.Request) uint {
	return conn.SendProto(QOT_GETOWNERPLATE, req)
}

/***
 * 获取大股东持股变化列表，具体字段请参考Qot_GetHoldingChangeList.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetHoldingChangeList(req *Qot_GetHoldingChangeList.Request) uint {
	return conn.SendProto(QOT_GETHOLDINGCHANGELIST, req)
}

/***
 * 筛选期权，具体字段请参考Qot_GetOptionChain.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetOptionChain(req *Qot_GetOptionChain.Request) uint {
	return conn.SendProto(QOT_GETOPTIONCHAIN, req)
}

/***
 * 筛选窝轮，具体字段请参考Qot_GetWarrant.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetWarrant(req *Qot_GetWarrant.Request) uint {
	return conn.SendProto(QOT_GETWARRANT, req)
}

/***
 * 获取资金流向，具体字段请参考Qot_GetCapitalFlow.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetCapitalFlow(req *Qot_GetCapitalFlow.Request) uint {
	return conn.SendProto(QOT_GETCAPITALFLOW, req)
}

/***
 * 获取资金分布，具体字段请参考Qot_GetCapitalDistribution.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetCapitalDistribution(req *Qot_GetCapitalDistribution.Request) uint {
	return conn.SendProto(QOT_GETCAPITALDISTRIBUTION, req)
}

/***
 * 获取自选股分组下的股票，具体字段请参考Qot_GetUserSecurity.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetUserSecurity(req *Qot_GetUserSecurity.Request) uint {
	return conn.SendProto(QOT_GETUSERSECURITY, req)
}

/***
 * 修改自选股分组下的股票，具体字段请参考Qot_ModifyUserSecurity.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) ModifyUserSecurity(req *Qot_ModifyUserSecurity.Request) uint {
	return conn.SendProto(QOT_MODIFYUSERSECURITY, req)
}

/***
 * 条件选股，具体字段请参考Qot_StockFilter.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) StockFilter(req *Qot_StockFilter.Request) uint {
	return conn.SendProto(QOT_STOCKFILTER, req)
}

/***
 * 获取股票代码变化信息，具体字段请参考Qot_GetCodeChange.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetCodeChange(req *Qot_GetCodeChange.Request) uint {
	return conn.SendProto(QOT_GETCODECHANGE, req)
}
/***
 * 获取IPO列表，具体字段请参考Qot_GetIpoList.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetIpoList(req *Qot_GetIpoList.Request) uint {
	return conn.SendProto(QOT_GETIPOLIST, req)
}
/***
 * 获取获取期货合约资料, 具体字段请参考Qot_GetFutureInfo.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnQot) GetFutureInfo(req *Qot_GetFutureInfo.Request) uint {
	return conn.SendProto(QOT_GETFUTUREINFO, req)
}

func (conn *FTAPIConnQot) OnReply(replyType FTAPI_ReqReplyType, protoHeader FTAPI_ProtoHeader, str string) {
	data := []byte(str)
	protoID := protoHeader.GetNProtoID()
	serialNo := protoHeader.GetNSerialNo()
	if (conn.qotSpi == nil) {
		return
	}
	var retType int32
	switch (protoID) {
	case GETGLOBALSTATE://获取全局状态
		{
			var rsp GetGlobalState.Response
			if (replyType == FTAPI_ReqReplyType_SvrReply) {
				err := proto.Unmarshal(data, &rsp)
				if err != nil {
					retType = int32(Common.RetType_RetType_Invalid)
					rsp.RetType = &retType
				}
			} else {
				retType = int32(replyType)
				rsp.RetType = &retType
			}
			conn.qotSpi.OnReply_GetGlobalState(conn.FTAPIConn, serialNo, &rsp)
		}
		break

	case QOT_GETBASICQOT: {
		var rsp Qot_GetBasicQot.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetBasicQot(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETBROKER: {
		var rsp Qot_GetBroker.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetBroker(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETCAPITALDISTRIBUTION: {
		var rsp Qot_GetCapitalDistribution.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetCapitalDistribution(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETCAPITALFLOW: {
		var rsp Qot_GetCapitalFlow.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetCapitalFlow(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETCODECHANGE: {
		var rsp Qot_GetCodeChange.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetCodeChange(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETHISTORYKL: {
		//	var rsp Qot_GetHistoryKL.Response
		//	if (replyType == FTAPI_ReqReplyType_SvrReply) {
		//		err := proto.Unmarshal(data, &rsp)
		//		if err != nil {
		//			retType = int32(Common.RetType_RetType_Invalid)
		//			rsp.RetType = &retType
		//		}
		//	} else {
		//		retType = int32(replyType)
		//		rsp.RetType = &retType
		//	}
		//
		//	conn.qotSpi.OnReply_GetHistoryKL(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETHISTORYKLPOINTS: {
		//var rsp Qot_GetHistoryKLPoints.Response
		//if (replyType == FTAPI_ReqReplyType_SvrReply) {
		//	err := proto.Unmarshal(data, &rsp)
		//	if err != nil {
		//		retType = int32(Common.RetType_RetType_Invalid)
		//		rsp.RetType = &retType
		//	}
		//} else {
		//	retType = int32(replyType)
		//	rsp.RetType = &retType
		//}
		//
		//conn.qotSpi.OnReply_GetHistoryKLPoints(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETHOLDINGCHANGELIST: {
		var rsp Qot_GetHoldingChangeList.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetHoldingChangeList(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETKL: {
		var rsp Qot_GetKL.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetKL(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETOPTIONCHAIN: {
		var rsp Qot_GetOptionChain.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetOptionChain(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETORDERBOOK: {
		var rsp Qot_GetOrderBook.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetOrderBook(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETOWNERPLATE: {
		var rsp Qot_GetOwnerPlate.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetOwnerPlate(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETPLATESECURITY: {
		var rsp Qot_GetPlateSecurity.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetPlateSecurity(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETPLATESET: {
		var rsp Qot_GetPlateSet.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetPlateSet(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETREFERENCE: {
		var rsp Qot_GetReference.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetReference(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETREHAB: {
		//var rsp Qot_GetRehab.Response
		//if (replyType == FTAPI_ReqReplyType_SvrReply) {
		//	err := proto.Unmarshal(data, &rsp)
		//	if err != nil {
		//		retType = int32(Common.RetType_RetType_Invalid)
		//		rsp.RetType = &retType
		//	}
		//} else {
		//	retType = int32(replyType)
		//	rsp.RetType = &retType
		//}
		//
		//conn.qotSpi.OnReply_GetRehab(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETRT: {
		var rsp Qot_GetRT.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetRT(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETSECURITYSNAPSHOT: {
		var rsp Qot_GetSecuritySnapshot.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetSecuritySnapshot(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETSTATICINFO: {
		var rsp Qot_GetStaticInfo.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetStaticInfo(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETSUBINFO: {
		var rsp Qot_GetSubInfo.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetSubInfo(conn.FTAPIConn, serialNo, &rsp)
	}
		break


	case QOT_GETTICKER: {
		var rsp Qot_GetTicker.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetTicker(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETTRADEDATE: {
		var rsp Qot_GetTradeDate.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetTradeDate(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETUSERSECURITY: {
		var rsp Qot_GetUserSecurity.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetUserSecurity(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_GETWARRANT: {
		var rsp Qot_GetWarrant.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetWarrant(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_MODIFYUSERSECURITY: {
		var rsp Qot_ModifyUserSecurity.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_ModifyUserSecurity(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_REGQOTPUSH: {
		var rsp Qot_RegQotPush.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_RegQotPush(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_REQUESTHISTORYKL: {
		var rsp Qot_RequestHistoryKL.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_RequestHistoryKL(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_REQUESTHISTORYKLQUOTA: {
		var rsp Qot_RequestHistoryKLQuota.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_RequestHistoryKLQuota(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_REQUESTREHAB: {
		var rsp Qot_RequestRehab.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_RequestRehab(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_STOCKFILTER: {
		var rsp Qot_StockFilter.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_StockFilter(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case QOT_SUB: {
		var rsp Qot_Sub.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_Sub(conn.FTAPIConn, serialNo, &rsp)
	}
		break
	case QOT_GETIPOLIST: {
		var rsp Qot_GetIpoList.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetIpoList(conn.FTAPIConn, serialNo, &rsp)
	}
		break
	case QOT_GETFUTUREINFO: {
		var rsp Qot_GetFutureInfo.Response
		if (replyType == FTAPI_ReqReplyType_SvrReply) {
			err := proto.Unmarshal(data, &rsp)
			if err != nil {
				retType = int32(Common.RetType_RetType_Invalid)
				rsp.RetType = &retType
			}
		} else {
			retType = int32(replyType)
			rsp.RetType = &retType
		}

		conn.qotSpi.OnReply_GetFutureInfo(conn.FTAPIConn, serialNo, &rsp)
	}
		break
	}
}

func (conn *FTAPIConnQot) OnPush(arg2 uintptr, arg3 FTAPI_ProtoHeader, arg4 string) {
	protoID := arg3.GetNProtoID()

	if (conn.qotSpi == nil) {
		return
	}

	var retType int32
	switch (protoID) {
	case NOTIFY: {
		var rsp Notify.Response
		err := proto.Unmarshal([]byte(arg4), &rsp)
		if err != nil {
			retType = int32(Common.RetType_RetType_Invalid)
			rsp.RetType = &retType
		}
		conn.qotSpi.OnPush_Notify(conn.FTAPIConn, &rsp)
	}
		break
	case QOT_UPDATEBASICQOT: {
		var rsp Qot_UpdateBasicQot.Response
		err := proto.Unmarshal([]byte(arg4), &rsp)
		if err != nil {
			retType = int32(Common.RetType_RetType_Invalid)
			rsp.RetType = &retType
		}
		conn.qotSpi.OnPush_UpdateBasicQuote(conn.FTAPIConn, &rsp)
	}
		break
	case QOT_UPDATEBROKER: {
		var rsp Qot_UpdateBroker.Response
		err := proto.Unmarshal([]byte(arg4), &rsp)
		if err != nil {
			retType = int32(Common.RetType_RetType_Invalid)
			rsp.RetType = &retType
		}
		conn.qotSpi.OnPush_UpdateBroker(conn.FTAPIConn, &rsp)
	}
		break
	case QOT_UPDATEKL: {
		var rsp Qot_UpdateKL.Response
		err := proto.Unmarshal([]byte(arg4), &rsp)
		if err != nil {
			retType = int32(Common.RetType_RetType_Invalid)
			rsp.RetType = &retType
		}
		conn.qotSpi.OnPush_UpdateKL(conn.FTAPIConn, &rsp)
	}
		break
	case QOT_UPDATEORDERBOOK: {
		var rsp Qot_UpdateOrderBook.Response
		err := proto.Unmarshal([]byte(arg4), &rsp)
		if err != nil {
			retType = int32(Common.RetType_RetType_Invalid)
			rsp.RetType = &retType
		}
		conn.qotSpi.OnPush_UpdateOrderBook(conn.FTAPIConn, &rsp)
	}
		break
	case QOT_UPDATERT: {
		var rsp Qot_UpdateRT.Response
		err := proto.Unmarshal([]byte(arg4), &rsp)
		if err != nil {
			retType = int32(Common.RetType_RetType_Invalid)
			rsp.RetType = &retType
		}
		conn.qotSpi.OnPush_UpdateRT(conn.FTAPIConn, &rsp)
	}
		break
	case QOT_UPDATETICKER: {
		var rsp Qot_UpdateTicker.Response
		err := proto.Unmarshal([]byte(arg4), &rsp)
		if err != nil {
			retType = int32(Common.RetType_RetType_Invalid)
			rsp.RetType = &retType
		}
		conn.qotSpi.OnPush_UpdateTicker(conn.FTAPIConn, &rsp)
	}
		break
	}
}