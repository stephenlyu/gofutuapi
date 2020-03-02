#include "./libcpp/Include/FTAPIChannel_Define.h"
#include "./libcpp/Include/FTAPIChannel.h"

class FTAPIChannel_Callback {
  public:
  virtual void OnDisConnect(FTAPIChannelPtr pChannel, Futu::i64_t nErrCode) = 0;
  virtual void OnInitConnect(FTAPIChannelPtr pChannel, Futu::i64_t nErrCode, const char* strDesc) = 0;
  virtual void OnReply(FTAPIChannelPtr pChannel, FTAPI_ReqReplyType enReqReplyType, const FTAPI_ProtoHeader* pProtoHeader, const Futu::i8_t* pProtoData, Futu::i32_t nDataLen) = 0;
  virtual void OnPush(FTAPIChannelPtr pChannel, const FTAPI_ProtoHeader* pProtoHeader, const Futu::i8_t* pProtoData, Futu::i32_t nDataLen) = 0;
  virtual ~FTAPIChannel_Callback() {};
};