#pragma once
#include "FTAPIChannel_Define.h"

#ifdef __cplusplus
extern "C" {
#endif

	typedef void *FTAPIChannelPtr;
	typedef void(*FTAPIChannel_OnDisConnectCallback)(FTAPIChannelPtr pChannel, Futu::i64_t nErrCode);
	typedef void(*FTAPIChannel_OnInitConnectCallback)(FTAPIChannelPtr pChannel, Futu::i64_t nErrCode, const char* strDesc);
	typedef void(*FTAPIChannel_OnReplyCallback)(FTAPIChannelPtr pChannel, FTAPI_ReqReplyType enReqReplyType, const FTAPI_ProtoHeader* pProtoHeader, const Futu::i8_t* pProtoData, Futu::i32_t nDataLen);
	typedef void(*FTAPIChannel_OnPushCallback)(FTAPIChannelPtr pChannel, const FTAPI_ProtoHeader* pProtoHeader, const Futu::i8_t* pProtoData, Futu::i32_t nDataLen);

	DllXPort_FTAPICHANNEL void FTAPIChannel_Init();
	DllXPort_FTAPICHANNEL void FTAPIChannel_UnInit();
	DllXPort_FTAPICHANNEL FTAPIChannelPtr CreateFTAPIChannel();
	DllXPort_FTAPICHANNEL void ReleaseFTAPIChannel(FTAPIChannelPtr pChannel);

	/* 设置客户端信息 */
	DllXPort_FTAPICHANNEL void FTAPIChannel_SetClientInfo(FTAPIChannelPtr pChannel, const char* szClientID, Futu::i32_t nClientVer);

	/* 设置密钥 */
	DllXPort_FTAPICHANNEL void FTAPIChannel_SetRSAPrivateKey(FTAPIChannelPtr pChannel, const char* strRSAPrivateKey);

	/* 初始化连接，连接并初始化 */
	DllXPort_FTAPICHANNEL Futu::i32_t FTAPIChannel_InitConnect(FTAPIChannelPtr pChannel, const char* szIPAddr, Futu::u16_t nPort, Futu::i32_t nEnableEncrypt);

	/* 此连接的连接ID，连接的唯一标识，InitConnect协议返回，没有初始化前为0 */
	DllXPort_FTAPICHANNEL Futu::u64_t FTAPIChannel_GetConnectID(FTAPIChannelPtr pChannel);

	DllXPort_FTAPICHANNEL Futu::u32_t FTAPIChannel_SendProto(FTAPIChannelPtr pChannel, Futu::u32_t nProtoID, Futu::u8_t nProtoVer, const Futu::i8_t* pProtoData, Futu::i32_t nDataLen);

	/* 关闭连接 */
	DllXPort_FTAPICHANNEL Futu::i32_t FTAPIChannel_Close(FTAPIChannelPtr pChannel);

  void wrap_InitializeCallbacks(FTAPIChannelPtr pChannel);

#ifdef __cplusplus
}
#endif