package futuapi

import (
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetAccList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_SubAccPush"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_UnlockTrade"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetOrderList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetMaxTrdQtys"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_PlaceOrder"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_ModifyOrder"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetOrderFillList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetHistoryOrderList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetHistoryOrderFillList"
	. "github.com/stephenlyu/gofutuapi"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetFunds"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetPositionList"
	"github.com/stephenlyu/gofutuapi/futuproto/Common"
	"github.com/golang/protobuf/proto"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_UpdateOrder"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_UpdateOrderFill"
)

type FTAPIConnTrd struct {
	*FTAPIConn
	trdSpi FTSPITrd
}

func NewFTAPIConnTrd() *FTAPIConnTrd {
	conn := new(FTAPIConnTrd)
	conn.FTAPIConn = NewFTAPIConn()
	FTAPIChannelSetCallbacks(conn.channel, NewDirectorFTAPIChannel_Callback(conn))
	return conn
}

func (conn *FTAPIConnTrd) SetTrdSpi(trdSpi FTSPITrd) {
	conn.trdSpi = trdSpi
}


/***
 * 获取交易账户列表，具体字段请参考Trd_GetAccList.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) GetAccList(req *Trd_GetAccList.Request) uint {
	return conn.SendProto(TRD_GETACCLIST, req)
}

/***
 * 解锁，具体字段请参考Trd_UnlockTrade.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) UnlockTrade(req *Trd_UnlockTrade.Request) uint {
	return conn.SendProto(TRD_UNLOCKTRADE, req)
}

/***
 * 订阅接收推送数据的交易账户，具体字段请参考Trd_SubAccPush.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) SubAccPush(req *Trd_SubAccPush.Request) uint {
	return conn.SendProto(TRD_SUBACCPUSH, req)
}

/***
 * 获取账户资金，具体字段请参考Trd_GetFunds.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) GetFunds(req *Trd_GetFunds.Request) uint {
	return conn.SendProto(TRD_GETFUNDS, req)
}

/***
 * 获取账户持仓，具体字段请参考Trd_GetPositionList.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) GetPositionList(req *Trd_GetPositionList.Request) uint {
	return conn.SendProto(TRD_GETPOSITIONLIST, req)
}

/***
 * 获取最大交易数量，具体字段请参考Trd_GetMaxTrdQtys.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) GetMaxTrdQtys(req *Trd_GetMaxTrdQtys.Request) uint {
	return conn.SendProto(TRD_GETMAXTRDQTYS, req)
}

/***
 * 获取当日订单列表，具体字段请参考Trd_GetOrderList.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) GetOrderList(req *Trd_GetOrderList.Request) uint {
	return conn.SendProto(TRD_GETORDERLIST, req)
}

/***
 * 下单，具体字段请参考Trd_PlaceOrder.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) PlaceOrder(req *Trd_PlaceOrder.Request) uint {
	//    	if (req.hasC2S()) {
	//    	    Common.PacketID packetID = nextPacketID()
	//    	    req = req.toBuilder().setC2S(req.getC2S().toBuilder().setPacketID(packetID)).build()
	//        }
	return conn.SendProto(TRD_PLACEORDER, req)
}

/***
 * 修改订单，具体字段请参考Trd_ModifyOrder.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) ModifyOrder(req *Trd_ModifyOrder.Request) uint {
	//        if (req.hasC2S()) {
	//            Common.PacketID packetID = nextPacketID()
	//            req = req.toBuilder().setC2S(req.getC2S().toBuilder().setPacketID(packetID)).build()
	//        }
	return conn.SendProto(TRD_MODIFYORDER, req)
}

/***
 * 获取当日成交列表，具体字段请参考Trd_GetOrderFillList.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) GetOrderFillList(req *Trd_GetOrderFillList.Request) uint {
	return conn.SendProto(TRD_GETORDERFILLLIST, req)
}

/***
 * 获取历史订单列表，具体字段请参考Trd_GetHistoryOrderList.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) GetHistoryOrderList(req *Trd_GetHistoryOrderList.Request) uint {
	return conn.SendProto(TRD_GETHISTORYORDERLIST, req)
}

/***
 * 获取历史成交列表，具体字段请参考Trd_GetHistoryOrderFillList.proto协议
 * @param req
 * @return 请求的序列号
 */
func (conn *FTAPIConnTrd) GetHistoryOrderFillList(req *Trd_GetHistoryOrderFillList.Request) uint {
	return conn.SendProto(TRD_GETHISTORYORDERFILLLIST, req)
}

func (conn *FTAPIConnTrd) OnReply(arg2 uintptr, replyType FTAPI_ReqReplyType, arg4 FTAPI_ProtoHeader, arg5 string) {
	data := []byte(arg5)
	protoID := arg4.GetNProtoID()
	serialNo := arg4.GetNSerialNo()
	if (conn.trdSpi == nil) {
		return
	}

	var retType int32
	switch (protoID) {
	case TRD_GETACCLIST: {
		var rsp Trd_GetAccList.Response
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

		conn.trdSpi.OnReply_GetAccList(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case TRD_GETFUNDS: {
		var rsp Trd_GetFunds.Response
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

		conn.trdSpi.OnReply_GetFunds(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case TRD_GETHISTORYORDERFILLLIST: {
		var rsp Trd_GetHistoryOrderFillList.Response
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

		conn.trdSpi.OnReply_GetHistoryOrderFillList(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case TRD_GETHISTORYORDERLIST: {
		var rsp Trd_GetHistoryOrderList.Response
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

		conn.trdSpi.OnReply_GetHistoryOrderList(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case TRD_GETMAXTRDQTYS: {
		var rsp Trd_GetMaxTrdQtys.Response
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

		conn.trdSpi.OnReply_GetMaxTrdQtys(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case TRD_GETORDERFILLLIST: {
		var rsp Trd_GetOrderFillList.Response
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

		conn.trdSpi.OnReply_GetOrderFillList(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case TRD_GETORDERLIST: {
		var rsp Trd_GetOrderList.Response
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

		conn.trdSpi.OnReply_GetOrderList(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case TRD_GETPOSITIONLIST: {
		var rsp Trd_GetPositionList.Response
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

		conn.trdSpi.OnReply_GetPositionList(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case TRD_MODIFYORDER: {
		var rsp Trd_ModifyOrder.Response
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

		conn.trdSpi.OnReply_ModifyOrder(conn.FTAPIConn, serialNo, &rsp)
	}
		break


	case TRD_PLACEORDER: {
		var rsp Trd_PlaceOrder.Response
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

		conn.trdSpi.OnReply_PlaceOrder(conn.FTAPIConn, serialNo, &rsp)
	}
		break


	case TRD_SUBACCPUSH: {
		var rsp Trd_SubAccPush.Response
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

		conn.trdSpi.OnReply_SubAccPush(conn.FTAPIConn, serialNo, &rsp)
	}
		break

	case TRD_UNLOCKTRADE: {
		var rsp Trd_UnlockTrade.Response
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

		conn.trdSpi.OnReply_UnlockTrade(conn.FTAPIConn, serialNo, &rsp)
	}
		break
	}
}

func (conn *FTAPIConnTrd) OnPush(arg2 uintptr, arg3 FTAPI_ProtoHeader, arg4 string) {
	protoID := arg3.GetNProtoID()

	if (conn.trdSpi == nil) {
		return
	}

	var retType int32
	switch (protoID) {
	case TRD_UPDATEORDER: {
		var rsp Trd_UpdateOrder.Response
		err := proto.Unmarshal([]byte(arg4), &rsp)
		if err != nil {
			retType = int32(Common.RetType_RetType_Invalid)
			rsp.RetType = &retType
		}
		conn.trdSpi.OnPush_UpdateOrder(conn.FTAPIConn, &rsp)
	}
		break
	case TRD_UPDATEORDERFILL: {
		var rsp Trd_UpdateOrderFill.Response
		err := proto.Unmarshal([]byte(arg4), &rsp)
		if err != nil {
			retType = int32(Common.RetType_RetType_Invalid)
			rsp.RetType = &retType
		}
		conn.trdSpi.OnPush_UpdateOrderFill(conn.FTAPIConn, &rsp)
	}
		break
	}
}