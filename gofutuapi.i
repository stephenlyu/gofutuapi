/* Copyright 2011 The Go Authors.  All rights reserved.
   Use of this source code is governed by a BSD-style
   license that can be found in the LICENSE file.  */

%module(directors="1") gofutuapi

%init %{ 
  //printf("Initialization gofutuapi done.\n");
%}

%typemap(gotype) (const Futu::i8_t* pProtoData, Futu::i32_t nDataLen) "string"

%typemap(in) (const Futu::i8_t* pProtoData, Futu::i32_t nDataLen)
%{
  $1 = (Futu::i8_t *) malloc($input.n);
  memcpy($1, $input.p, $input.n);
  $2 = (Futu::i32_t) $input.n;
%}

%typemap(freearg) (const Futu::i8_t* pProtoData, Futu::i32_t nDataLen)
%{
  {
    free($1);
  }
%}


/* Includes the header files in the wrapper code */
%header %{
#include "./libcpp/Include/FTAPIChannel_Define.h"
#include "./libcpp/Include/FTAPIChannel.h"
#include "FTAPIChannel_Callback.h"
%}

%{

extern "C" void* Wrap_GetCallback(FTAPIChannelPtr pChannel);
static void FTAPIChannel_OnDisConnectCallback_impl(FTAPIChannelPtr pChannel, Futu::i64_t nErrCode) {
  FTAPIChannel_Callback* callback = (FTAPIChannel_Callback*) Wrap_GetCallback(pChannel);
  if (callback != 0) {
    callback->OnDisConnect(pChannel, nErrCode);
  }
}

static void FTAPIChannel_OnInitConnectCallback_impl(FTAPIChannelPtr pChannel, Futu::i64_t nErrCode, const char* strDesc) {
  FTAPIChannel_Callback* callback = (FTAPIChannel_Callback*) Wrap_GetCallback(pChannel);
  if (callback != 0) {
    callback->OnInitConnect(pChannel, nErrCode, strDesc);
  }
}

static void FTAPIChannel_OnReplyCallback_impl(FTAPIChannelPtr pChannel, FTAPI_ReqReplyType enReqReplyType, const FTAPI_ProtoHeader* pProtoHeader, const Futu::i8_t* pProtoData, Futu::i32_t nDataLen) {
  FTAPIChannel_Callback* callback = (FTAPIChannel_Callback*) Wrap_GetCallback(pChannel);
  if (callback != 0) {
    callback->OnReply(pChannel, enReqReplyType, pProtoHeader, pProtoData, nDataLen);
  }
}

static void FTAPIChannel_OnPushCallback_impl(FTAPIChannelPtr pChannel, const FTAPI_ProtoHeader* pProtoHeader, const Futu::i8_t* pProtoData, Futu::i32_t nDataLen) {
  FTAPIChannel_Callback* callback = (FTAPIChannel_Callback*) Wrap_GetCallback(pChannel);
  if (callback != 0) {
    callback->OnPush(pChannel, pProtoHeader, pProtoData, nDataLen);
  }
}

void wrap_InitializeCallbacks(FTAPIChannelPtr pChannel) {
  FTAPIChannel_SetOnDisconnectCallback(pChannel, FTAPIChannel_OnDisConnectCallback_impl);
  FTAPIChannel_SetOnInitConnectCallback(pChannel, FTAPIChannel_OnInitConnectCallback_impl);
  FTAPIChannel_SetOnReplyCallback(pChannel, FTAPIChannel_OnReplyCallback_impl);
  FTAPIChannel_SetOnPushCallback(pChannel, FTAPIChannel_OnPushCallback_impl);
}

%}

%insert(go_wrapper) %{

var gChannelCallbacks = make(map[uintptr]FTAPIChannel_Callback)

//export Wrap_GetCallback
func Wrap_GetCallback(pChannel uintptr) uintptr {
  cb, ok := gChannelCallbacks[pChannel]
  if !ok {
    return 0
  }
  return cb.Swigcptr()
}

func FTAPIChannelSetCallbacks(pChannel uintptr, callback FTAPIChannel_Callback) {
  gChannelCallbacks[pChannel] = callback
  Wrap_InitializeCallbacks(pChannel)
}

func FTAPIChannelDestroy(pChannel uintptr) {
  delete(gChannelCallbacks, pChannel)
  ReleaseFTAPIChannel(pChannel)
}

%}

/* Parse the header files to generate wrappers */
%include "std_string.i"

%feature("director") FTAPIChannel_Callback;

%include "./libcpp/Include/FTAPIChannel_Define.h"
%include "./libcpp/Include/FTAPIChannel_swig.h"
%include "FTAPIChannel_Callback.h"
