#pragma once

/**
* @file FTAPIChannel_Define.h
* @brief FTAPIChannel一些定义文件
* @author Lin
* @date 2018/11/11 11:11
* @version 1.0
* @copyright Copyright (C) 2018 富途
*/
#include <stdint.h>

namespace Futu {

typedef char i8_t; // arm-linux-gcc把char默认定义为了unsigned char，用-fsigned-char参数编译变成signed char
typedef short i16_t;
typedef int i32_t;
#ifdef __linux__
typedef int64_t i64_t;
#else
typedef long long i64_t;
#endif

typedef unsigned char u8_t;
typedef unsigned short u16_t;
typedef unsigned int u32_t;
#ifdef __linux__
typedef uint64_t u64_t;
#else
typedef unsigned long long u64_t;
#endif
}
#if defined _WIN32
#	ifdef DllXPort_FTAPICHANNEL_Import
#		define DllXPort_FTAPICHANNEL __declspec(dllimport)
#	else
#		define DllXPort_FTAPICHANNEL __declspec(dllexport)
#	endif
#else
#	define DllXPort_FTAPICHANNEL
#endif

#pragma pack(push, FTAPI_ProtoHeader, 1)
/**
* @struct FTAPI_ProtoHeader
* @brief FTAPI协议包头结构定义
*/
struct FTAPI_ProtoHeader
{
	Futu::u8_t szHeaderFlag[2]; /**< 协议包头标识，固定为"FT" */
	Futu::u32_t nProtoID;	 /**< 协议号 */
	Futu::u8_t nProtoFmtType; /**< 协议格式，0代表Protobuf格式，1代表Json格式 */
	Futu::u8_t nProtoVer; /**< 协议版本，用于后续迭代升级兼容 */
	Futu::u32_t nSerialNo; /**< 包序列号 */
	Futu::u32_t nBodyLen; /**< 包体长度 */
	Futu::u8_t arrBodySHA1[20]; /**< 包体原始(加密前)数据的SHA1哈希值 */
	Futu::u8_t arrReserved[8]; /**< 保留8字节扩展 */
};
#pragma pack(pop, FTAPI_ProtoHeader)

#define FTAPI_ProtoFmtType_Protobuf 0 /**< Protobuf协议格式 */
#define FTAPI_ProtoFmtType_Json 1 /**< Json协议格式 */

/**
* @enum FTAPI_ReqReplyType
* @brief 请求应答类型枚举
*/
enum FTAPI_ReqReplyType
{
	FTAPI_ReqReplyType_SvrReply = 0, /**< 来自服务器的应答 */
	FTAPI_ReqReplyType_Timeout = -100, /**< 等待服务器应答超时 */
	FTAPI_ReqReplyType_DisConnect = -200, /**< 因连接已断开(被动断开或主动关闭)的应答 */
};

enum FTAPI_InitFailType
{
	FTAPI_InitFailType_Unknow = 0,  /**< 未知 */
	FTAPI_InitFailType_Timeout = 1,  /**< 超时 */
	FTAPI_InitFailType_DisConnect = 3,  /**< 连接断开 */
	FTAPI_InitFailType_SeriaNoNotMatch = 4, /**< 序列号不符 */
	FTAPI_InitFailType_SendInitReqFailed = 4, /**< 发送初始化协议失败 */
	FTAPI_InitFailType_OpenDReject = 5, /**< FutuOpenD回包指定错误，具体错误看描述*/
};

enum FTAPI_ConnectFailType
{
	FTAPI_ConnectFailType_Unknown = -1,
	FTAPI_ConnectFailType_None = 0,
	FTAPI_ConnectFailType_CreateFailed = 1,
	FTAPI_ConnectFailType_CloseFailed = 2,
	FTAPI_ConnectFailType_ShutdownFailed = 3,
	FTAPI_ConnectFailType_GetHostByNameFailed = 4,
	FTAPI_ConnectFailType_GetHostByNameWrong = 5,
	FTAPI_ConnectFailType_ConnectFailed = 6,
	FTAPI_ConnectFailType_BindFailed = 7,
	FTAPI_ConnectFailType_ListenFailed = 8,
	FTAPI_ConnectFailType_SelectReturnError = 9,
	FTAPI_ConnectFailType_SendFailed = 10,
	FTAPI_ConnectFailType_RecvFailed = 11,
};

#define FTAPI_InitFail		100
