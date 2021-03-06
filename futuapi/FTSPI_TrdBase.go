package futuapi

import (
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetAccList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_UnlockTrade"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_SubAccPush"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetFunds"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetPositionList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetMaxTrdQtys"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetOrderList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_PlaceOrder"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_ModifyOrder"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetOrderFillList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetHistoryOrderList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_GetHistoryOrderFillList"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_UpdateOrder"
	"github.com/stephenlyu/gofutuapi/futuproto/Trd_UpdateOrderFill"
)

type FTSPI_TrdBase struct {
}

func (spi *FTSPI_TrdBase) OnReply_GetAccList(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_GetAccList.Response) {} //获取交易账户列表回调
func (spi *FTSPI_TrdBase) OnReply_UnlockTrade(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_UnlockTrade.Response) {} //解锁回调
func (spi *FTSPI_TrdBase) OnReply_SubAccPush(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_SubAccPush.Response) {} //订阅接收推送数据的交易账户回调
func (spi *FTSPI_TrdBase) OnReply_GetFunds(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_GetFunds.Response) {} //获取账户资金回调
func (spi *FTSPI_TrdBase) OnReply_GetPositionList(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_GetPositionList.Response) {} //获取账户持仓回调
func (spi *FTSPI_TrdBase) OnReply_GetMaxTrdQtys(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_GetMaxTrdQtys.Response) {} //获取最大交易数量回调
func (spi *FTSPI_TrdBase) OnReply_GetOrderList(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_GetOrderList.Response) {} //获取当日订单列表回调
func (spi *FTSPI_TrdBase) OnReply_PlaceOrder(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_PlaceOrder.Response) {} //下单回调
func (spi *FTSPI_TrdBase) OnReply_ModifyOrder(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_ModifyOrder.Response) {} //修改订单回调
func (spi *FTSPI_TrdBase) OnReply_GetOrderFillList(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_GetOrderFillList.Response) {} //获取当日成交列表回调
func (spi *FTSPI_TrdBase) OnReply_GetHistoryOrderList(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_GetHistoryOrderList.Response) {} //获取历史订单列表回调
func (spi *FTSPI_TrdBase) OnReply_GetHistoryOrderFillList(client FTAPIConnTrd, nSerialNo uint, rsp *Trd_GetHistoryOrderFillList.Response) {} //获取历史成交列表回调
func (spi *FTSPI_TrdBase) OnPush_UpdateOrder(client FTAPIConnTrd, rsp *Trd_UpdateOrder.Response) {} //推送订单变化
func (spi *FTSPI_TrdBase) OnPush_UpdateOrderFill(client FTAPIConnTrd, rsp *Trd_UpdateOrderFill.Response) {} //推送订单成交