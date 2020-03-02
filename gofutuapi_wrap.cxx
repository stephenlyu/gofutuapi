/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: gofutuapi.i

#define SWIGMODULE gofutuapi
#define SWIG_DIRECTORS

#ifdef __cplusplus
/* SwigValueWrapper is described in swig.swg */
template<typename T> class SwigValueWrapper {
  struct SwigMovePointer {
    T *ptr;
    SwigMovePointer(T *p) : ptr(p) { }
    ~SwigMovePointer() { delete ptr; }
    SwigMovePointer& operator=(SwigMovePointer& rhs) { T* oldptr = ptr; ptr = 0; delete oldptr; ptr = rhs.ptr; rhs.ptr = 0; return *this; }
  } pointer;
  SwigValueWrapper& operator=(const SwigValueWrapper<T>& rhs);
  SwigValueWrapper(const SwigValueWrapper<T>& rhs);
public:
  SwigValueWrapper() : pointer(0) { }
  SwigValueWrapper& operator=(const T& t) { SwigMovePointer tmp(new T(t)); pointer = tmp; return *this; }
  operator T&() const { return *pointer.ptr; }
  T *operator&() { return pointer.ptr; }
};

template <typename T> T SwigValueInit() {
  return T();
}
#endif

/* -----------------------------------------------------------------------------
 *  This section contains generic SWIG labels for method/variable
 *  declarations/attributes, and other compiler dependent labels.
 * ----------------------------------------------------------------------------- */

/* template workaround for compilers that cannot correctly implement the C++ standard */
#ifndef SWIGTEMPLATEDISAMBIGUATOR
# if defined(__SUNPRO_CC) && (__SUNPRO_CC <= 0x560)
#  define SWIGTEMPLATEDISAMBIGUATOR template
# elif defined(__HP_aCC)
/* Needed even with `aCC -AA' when `aCC -V' reports HP ANSI C++ B3910B A.03.55 */
/* If we find a maximum version that requires this, the test would be __HP_aCC <= 35500 for A.03.55 */
#  define SWIGTEMPLATEDISAMBIGUATOR template
# else
#  define SWIGTEMPLATEDISAMBIGUATOR
# endif
#endif

/* inline attribute */
#ifndef SWIGINLINE
# if defined(__cplusplus) || (defined(__GNUC__) && !defined(__STRICT_ANSI__))
#   define SWIGINLINE inline
# else
#   define SWIGINLINE
# endif
#endif

/* attribute recognised by some compilers to avoid 'unused' warnings */
#ifndef SWIGUNUSED
# if defined(__GNUC__)
#   if !(defined(__cplusplus)) || (__GNUC__ > 3 || (__GNUC__ == 3 && __GNUC_MINOR__ >= 4))
#     define SWIGUNUSED __attribute__ ((__unused__))
#   else
#     define SWIGUNUSED
#   endif
# elif defined(__ICC)
#   define SWIGUNUSED __attribute__ ((__unused__))
# else
#   define SWIGUNUSED
# endif
#endif

#ifndef SWIG_MSC_UNSUPPRESS_4505
# if defined(_MSC_VER)
#   pragma warning(disable : 4505) /* unreferenced local function has been removed */
# endif
#endif

#ifndef SWIGUNUSEDPARM
# ifdef __cplusplus
#   define SWIGUNUSEDPARM(p)
# else
#   define SWIGUNUSEDPARM(p) p SWIGUNUSED
# endif
#endif

/* internal SWIG method */
#ifndef SWIGINTERN
# define SWIGINTERN static SWIGUNUSED
#endif

/* internal inline SWIG method */
#ifndef SWIGINTERNINLINE
# define SWIGINTERNINLINE SWIGINTERN SWIGINLINE
#endif

/* exporting methods */
#if defined(__GNUC__)
#  if (__GNUC__ >= 4) || (__GNUC__ == 3 && __GNUC_MINOR__ >= 4)
#    ifndef GCC_HASCLASSVISIBILITY
#      define GCC_HASCLASSVISIBILITY
#    endif
#  endif
#endif

#ifndef SWIGEXPORT
# if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__)
#   if defined(STATIC_LINKED)
#     define SWIGEXPORT
#   else
#     define SWIGEXPORT __declspec(dllexport)
#   endif
# else
#   if defined(__GNUC__) && defined(GCC_HASCLASSVISIBILITY)
#     define SWIGEXPORT __attribute__ ((visibility("default")))
#   else
#     define SWIGEXPORT
#   endif
# endif
#endif

/* calling conventions for Windows */
#ifndef SWIGSTDCALL
# if defined(_WIN32) || defined(__WIN32__) || defined(__CYGWIN__)
#   define SWIGSTDCALL __stdcall
# else
#   define SWIGSTDCALL
# endif
#endif

/* Deal with Microsoft's attempt at deprecating C standard runtime functions */
#if !defined(SWIG_NO_CRT_SECURE_NO_DEPRECATE) && defined(_MSC_VER) && !defined(_CRT_SECURE_NO_DEPRECATE)
# define _CRT_SECURE_NO_DEPRECATE
#endif

/* Deal with Microsoft's attempt at deprecating methods in the standard C++ library */
#if !defined(SWIG_NO_SCL_SECURE_NO_DEPRECATE) && defined(_MSC_VER) && !defined(_SCL_SECURE_NO_DEPRECATE)
# define _SCL_SECURE_NO_DEPRECATE
#endif

/* Deal with Apple's deprecated 'AssertMacros.h' from Carbon-framework */
#if defined(__APPLE__) && !defined(__ASSERT_MACROS_DEFINE_VERSIONS_WITHOUT_UNDERSCORES)
# define __ASSERT_MACROS_DEFINE_VERSIONS_WITHOUT_UNDERSCORES 0
#endif

/* Intel's compiler complains if a variable which was never initialised is
 * cast to void, which is a common idiom which we use to indicate that we
 * are aware a variable isn't used.  So we just silence that warning.
 * See: https://github.com/swig/swig/issues/192 for more discussion.
 */
#ifdef __INTEL_COMPILER
# pragma warning disable 592
#endif


#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>



typedef long long intgo;
typedef unsigned long long uintgo;


# if !defined(__clang__) && (defined(__i386__) || defined(__x86_64__))
#   define SWIGSTRUCTPACKED __attribute__((__packed__, __gcc_struct__))
# else
#   define SWIGSTRUCTPACKED __attribute__((__packed__))
# endif



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;




#define swiggo_size_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];
#define swiggo_size_assert(t, n) swiggo_size_assert_eq(sizeof(t), n, swiggo_sizeof_##t##_is_not_##n)

swiggo_size_assert(char, 1)
swiggo_size_assert(short, 2)
swiggo_size_assert(int, 4)
typedef long long swiggo_long_long;
swiggo_size_assert(swiggo_long_long, 8)
swiggo_size_assert(float, 4)
swiggo_size_assert(double, 8)

#ifdef __cplusplus
extern "C" {
#endif
extern void crosscall2(void (*fn)(void *, int), void *, int);
extern char* _cgo_topofstack(void) __attribute__ ((weak));
extern void _cgo_allocate(void *, int);
extern void _cgo_panic(void *, int);
#ifdef __cplusplus
}
#endif

static char *_swig_topofstack() {
  if (_cgo_topofstack) {
    return _cgo_topofstack();
  } else {
    return 0;
  }
}

static void _swig_gopanic(const char *p) {
  struct {
    const char *p;
  } SWIGSTRUCTPACKED a;
  a.p = p;
  crosscall2(_cgo_panic, &a, (int) sizeof a);
}




#define SWIG_contract_assert(expr, msg) \
  if (!(expr)) { _swig_gopanic(msg); } else


static _gostring_ Swig_AllocateString(const char *p, size_t l) {
  _gostring_ ret;
  ret.p = (char*)malloc(l);
  memcpy(ret.p, p, l);
  ret.n = l;
  return ret;
}


static void Swig_free(void* p) {
  free(p);
}

static void* Swig_malloc(int c) {
  return malloc(c);
}


#include "./libcpp/Include/FTAPIChannel_Define.h"
#include "./libcpp/Include/FTAPIChannel.h"
#include "FTAPIChannel_Callback.h"



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



#include <string>


// C++ director class methods.
#include "gofutuapi_wrap.h"


#include <map>

namespace {
  struct GCItem {
    virtual ~GCItem() {}
  };

  struct GCItem_var {
    GCItem_var(GCItem *item = 0) : _item(item) {
    }

    GCItem_var& operator=(GCItem *item) {
      GCItem *tmp = _item;
      _item = item;
      delete tmp;
      return *this;
    }

    ~GCItem_var() {
      delete _item;
    }

    GCItem* operator->() {
      return _item;
    }

    private:
      GCItem *_item;
  };

  template <typename Type>
  struct GCItem_T : GCItem {
    GCItem_T(Type *ptr) : _ptr(ptr) {
    }

    virtual ~GCItem_T() {
      delete _ptr;
    }

  private:
    Type *_ptr;
  };
}

class Swig_memory {
public:
  template <typename Type>
  void swig_acquire_pointer(Type* vptr) {
    if (vptr) {
      swig_owner[vptr] = new GCItem_T<Type>(vptr);
    }
  }
private:
  typedef std::map<void *, GCItem_var> swig_ownership_map;
  swig_ownership_map swig_owner;
};

template <typename Type>
static void swig_acquire_pointer(Swig_memory** pmem, Type* ptr) {
  if (!pmem) {
    *pmem = new Swig_memory;
  }
  (*pmem)->swig_acquire_pointer(ptr);
}

SwigDirector_FTAPIChannel_Callback::SwigDirector_FTAPIChannel_Callback(int swig_p)
    : FTAPIChannel_Callback(),
      go_val(swig_p), swig_mem(0)
{ }

extern "C" void Swig_DirectorFTAPIChannel_Callback_callback_OnDisConnect_gofutuapi_d667299e3498665a(int, FTAPIChannelPtr pChannel, long long nErrCode);
void SwigDirector_FTAPIChannel_Callback::OnDisConnect(FTAPIChannelPtr pChannel, Futu::i64_t nErrCode) {
  FTAPIChannelPtr swig_pChannel;
  long long swig_nErrCode;
  
  *(FTAPIChannelPtr *)&swig_pChannel = (FTAPIChannelPtr)pChannel; 
  swig_nErrCode = (Futu::i64_t)nErrCode; 
  Swig_DirectorFTAPIChannel_Callback_callback_OnDisConnect_gofutuapi_d667299e3498665a(go_val, swig_pChannel, swig_nErrCode);
}

extern "C" void Swig_DirectorFTAPIChannel_Callback_callback_OnInitConnect_gofutuapi_d667299e3498665a(int, FTAPIChannelPtr pChannel, long long nErrCode, _gostring_ strDesc);
void SwigDirector_FTAPIChannel_Callback::OnInitConnect(FTAPIChannelPtr pChannel, Futu::i64_t nErrCode, char const *strDesc) {
  FTAPIChannelPtr swig_pChannel;
  long long swig_nErrCode;
  _gostring_ swig_strDesc;
  
  *(FTAPIChannelPtr *)&swig_pChannel = (FTAPIChannelPtr)pChannel; 
  swig_nErrCode = (Futu::i64_t)nErrCode; 
  
  swig_strDesc = Swig_AllocateString((char*)strDesc, strDesc ? strlen((char*)strDesc) : 0);
  
  Swig_DirectorFTAPIChannel_Callback_callback_OnInitConnect_gofutuapi_d667299e3498665a(go_val, swig_pChannel, swig_nErrCode, swig_strDesc);
}

extern "C" void Swig_DirectorFTAPIChannel_Callback_callback_OnReply_gofutuapi_d667299e3498665a(int, FTAPIChannelPtr pChannel, intgo enReqReplyType, FTAPI_ProtoHeader *pProtoHeader, _gostring_ pProtoData, intgo nDataLen);
void SwigDirector_FTAPIChannel_Callback::OnReply(FTAPIChannelPtr pChannel, FTAPI_ReqReplyType enReqReplyType, FTAPI_ProtoHeader const *pProtoHeader, Futu::i8_t const *pProtoData, Futu::i32_t nDataLen) {
  FTAPIChannelPtr swig_pChannel;
  intgo swig_enReqReplyType;
  FTAPI_ProtoHeader *swig_pProtoHeader;
  _gostring_ swig_pProtoData;
  intgo swig_nDataLen;
  
  *(FTAPIChannelPtr *)&swig_pChannel = (FTAPIChannelPtr)pChannel; 
  swig_enReqReplyType = (intgo)enReqReplyType; 
  *(FTAPI_ProtoHeader **)&swig_pProtoHeader = (FTAPI_ProtoHeader *)pProtoHeader; 
  
  swig_pProtoData = Swig_AllocateString((char*)pProtoData, pProtoData ? strlen((char*)pProtoData) : 0);
  
  swig_nDataLen = (Futu::i32_t)nDataLen; 
  Swig_DirectorFTAPIChannel_Callback_callback_OnReply_gofutuapi_d667299e3498665a(go_val, swig_pChannel, swig_enReqReplyType, swig_pProtoHeader, swig_pProtoData, swig_nDataLen);
}

extern "C" void Swig_DirectorFTAPIChannel_Callback_callback_OnPush_gofutuapi_d667299e3498665a(int, FTAPIChannelPtr pChannel, FTAPI_ProtoHeader *pProtoHeader, _gostring_ pProtoData, intgo nDataLen);
void SwigDirector_FTAPIChannel_Callback::OnPush(FTAPIChannelPtr pChannel, FTAPI_ProtoHeader const *pProtoHeader, Futu::i8_t const *pProtoData, Futu::i32_t nDataLen) {
  FTAPIChannelPtr swig_pChannel;
  FTAPI_ProtoHeader *swig_pProtoHeader;
  _gostring_ swig_pProtoData;
  intgo swig_nDataLen;
  
  *(FTAPIChannelPtr *)&swig_pChannel = (FTAPIChannelPtr)pChannel; 
  *(FTAPI_ProtoHeader **)&swig_pProtoHeader = (FTAPI_ProtoHeader *)pProtoHeader; 
  
  swig_pProtoData = Swig_AllocateString((char*)pProtoData, pProtoData ? strlen((char*)pProtoData) : 0);
  
  swig_nDataLen = (Futu::i32_t)nDataLen; 
  Swig_DirectorFTAPIChannel_Callback_callback_OnPush_gofutuapi_d667299e3498665a(go_val, swig_pChannel, swig_pProtoHeader, swig_pProtoData, swig_nDataLen);
}

extern "C" void Swiggo_DeleteDirector_FTAPIChannel_Callback_gofutuapi_d667299e3498665a(intgo);
SwigDirector_FTAPIChannel_Callback::~SwigDirector_FTAPIChannel_Callback()
{
  Swiggo_DeleteDirector_FTAPIChannel_Callback_gofutuapi_d667299e3498665a(go_val);
  delete swig_mem;
}

#ifdef __cplusplus
extern "C" {
#endif

void _wrap_Swig_free_gofutuapi_d667299e3498665a(void *_swig_go_0) {
  void *arg1 = (void *) 0 ;
  
  arg1 = *(void **)&_swig_go_0; 
  
  Swig_free(arg1);
  
}


void *_wrap_Swig_malloc_gofutuapi_d667299e3498665a(intgo _swig_go_0) {
  int arg1 ;
  void *result = 0 ;
  void *_swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = (void *)Swig_malloc(arg1);
  *(void **)&_swig_go_result = (void *)result; 
  return _swig_go_result;
}


void _wrap_FTAPI_ProtoHeader_szHeaderFlag_set_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0, char *_swig_go_1) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t *arg2 ;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  arg2 = *(Futu::u8_t **)&_swig_go_1; 
  
  {
    size_t ii;
    Futu::u8_t *b = (Futu::u8_t *) arg1->szHeaderFlag;
    for (ii = 0; ii < (size_t)2; ii++) b[ii] = *((Futu::u8_t *) arg2 + ii);
  }
  
}


char *_wrap_FTAPI_ProtoHeader_szHeaderFlag_get_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t *result = 0 ;
  char *_swig_go_result;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  
  result = (Futu::u8_t *)(Futu::u8_t *) ((arg1)->szHeaderFlag);
  *(Futu::u8_t **)&_swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_FTAPI_ProtoHeader_nProtoID_set_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0, intgo _swig_go_1) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u32_t arg2 ;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  arg2 = (Futu::u32_t)_swig_go_1; 
  
  if (arg1) (arg1)->nProtoID = arg2;
  
}


intgo _wrap_FTAPI_ProtoHeader_nProtoID_get_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u32_t result;
  intgo _swig_go_result;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  
  result = (Futu::u32_t) ((arg1)->nProtoID);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_FTAPI_ProtoHeader_nProtoFmtType_set_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0, char _swig_go_1) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t arg2 ;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  arg2 = (Futu::u8_t)_swig_go_1; 
  
  if (arg1) (arg1)->nProtoFmtType = arg2;
  
}


char _wrap_FTAPI_ProtoHeader_nProtoFmtType_get_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t result;
  char _swig_go_result;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  
  result = (Futu::u8_t) ((arg1)->nProtoFmtType);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_FTAPI_ProtoHeader_nProtoVer_set_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0, char _swig_go_1) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t arg2 ;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  arg2 = (Futu::u8_t)_swig_go_1; 
  
  if (arg1) (arg1)->nProtoVer = arg2;
  
}


char _wrap_FTAPI_ProtoHeader_nProtoVer_get_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t result;
  char _swig_go_result;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  
  result = (Futu::u8_t) ((arg1)->nProtoVer);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_FTAPI_ProtoHeader_nSerialNo_set_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0, intgo _swig_go_1) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u32_t arg2 ;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  arg2 = (Futu::u32_t)_swig_go_1; 
  
  if (arg1) (arg1)->nSerialNo = arg2;
  
}


intgo _wrap_FTAPI_ProtoHeader_nSerialNo_get_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u32_t result;
  intgo _swig_go_result;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  
  result = (Futu::u32_t) ((arg1)->nSerialNo);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_FTAPI_ProtoHeader_nBodyLen_set_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0, intgo _swig_go_1) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u32_t arg2 ;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  arg2 = (Futu::u32_t)_swig_go_1; 
  
  if (arg1) (arg1)->nBodyLen = arg2;
  
}


intgo _wrap_FTAPI_ProtoHeader_nBodyLen_get_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u32_t result;
  intgo _swig_go_result;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  
  result = (Futu::u32_t) ((arg1)->nBodyLen);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_FTAPI_ProtoHeader_arrBodySHA1_set_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0, char *_swig_go_1) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t *arg2 ;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  arg2 = *(Futu::u8_t **)&_swig_go_1; 
  
  {
    size_t ii;
    Futu::u8_t *b = (Futu::u8_t *) arg1->arrBodySHA1;
    for (ii = 0; ii < (size_t)20; ii++) b[ii] = *((Futu::u8_t *) arg2 + ii);
  }
  
}


char *_wrap_FTAPI_ProtoHeader_arrBodySHA1_get_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t *result = 0 ;
  char *_swig_go_result;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  
  result = (Futu::u8_t *)(Futu::u8_t *) ((arg1)->arrBodySHA1);
  *(Futu::u8_t **)&_swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_FTAPI_ProtoHeader_arrReserved_set_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0, char *_swig_go_1) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t *arg2 ;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  arg2 = *(Futu::u8_t **)&_swig_go_1; 
  
  {
    size_t ii;
    Futu::u8_t *b = (Futu::u8_t *) arg1->arrReserved;
    for (ii = 0; ii < (size_t)8; ii++) b[ii] = *((Futu::u8_t *) arg2 + ii);
  }
  
}


char *_wrap_FTAPI_ProtoHeader_arrReserved_get_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  Futu::u8_t *result = 0 ;
  char *_swig_go_result;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  
  result = (Futu::u8_t *)(Futu::u8_t *) ((arg1)->arrReserved);
  *(Futu::u8_t **)&_swig_go_result = result; 
  return _swig_go_result;
}


FTAPI_ProtoHeader *_wrap_new_FTAPI_ProtoHeader_gofutuapi_d667299e3498665a() {
  FTAPI_ProtoHeader *result = 0 ;
  FTAPI_ProtoHeader *_swig_go_result;
  
  
  result = (FTAPI_ProtoHeader *)new FTAPI_ProtoHeader();
  *(FTAPI_ProtoHeader **)&_swig_go_result = (FTAPI_ProtoHeader *)result; 
  return _swig_go_result;
}


void _wrap_delete_FTAPI_ProtoHeader_gofutuapi_d667299e3498665a(FTAPI_ProtoHeader *_swig_go_0) {
  FTAPI_ProtoHeader *arg1 = (FTAPI_ProtoHeader *) 0 ;
  
  arg1 = *(FTAPI_ProtoHeader **)&_swig_go_0; 
  
  delete arg1;
  
}


intgo _wrap_FTAPI_ReqReplyType_SvrReply_gofutuapi_d667299e3498665a() {
  FTAPI_ReqReplyType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ReqReplyType_SvrReply;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ReqReplyType_Timeout_gofutuapi_d667299e3498665a() {
  FTAPI_ReqReplyType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ReqReplyType_Timeout;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ReqReplyType_DisConnect_gofutuapi_d667299e3498665a() {
  FTAPI_ReqReplyType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ReqReplyType_DisConnect;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_InitFailType_Unknow_gofutuapi_d667299e3498665a() {
  FTAPI_InitFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_InitFailType_Unknow;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_InitFailType_Timeout_gofutuapi_d667299e3498665a() {
  FTAPI_InitFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_InitFailType_Timeout;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_InitFailType_DisConnect_gofutuapi_d667299e3498665a() {
  FTAPI_InitFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_InitFailType_DisConnect;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_InitFailType_SeriaNoNotMatch_gofutuapi_d667299e3498665a() {
  FTAPI_InitFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_InitFailType_SeriaNoNotMatch;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_InitFailType_SendInitReqFailed_gofutuapi_d667299e3498665a() {
  FTAPI_InitFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_InitFailType_SendInitReqFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_InitFailType_OpenDReject_gofutuapi_d667299e3498665a() {
  FTAPI_InitFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_InitFailType_OpenDReject;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_Unknown_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_Unknown;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_None_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_None;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_CreateFailed_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_CreateFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_CloseFailed_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_CloseFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_ShutdownFailed_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_ShutdownFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_GetHostByNameFailed_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_GetHostByNameFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_GetHostByNameWrong_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_GetHostByNameWrong;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_ConnectFailed_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_ConnectFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_BindFailed_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_BindFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_ListenFailed_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_ListenFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_SelectReturnError_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_SelectReturnError;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_SendFailed_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_SendFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


intgo _wrap_FTAPI_ConnectFailType_RecvFailed_gofutuapi_d667299e3498665a() {
  FTAPI_ConnectFailType result;
  intgo _swig_go_result;
  
  
  result = FTAPI_ConnectFailType_RecvFailed;
  
  _swig_go_result = (intgo)result; 
  return _swig_go_result;
}


void _wrap_FTAPIChannel_Init_gofutuapi_d667299e3498665a() {
  FTAPIChannel_Init();
  
}


void _wrap_FTAPIChannel_UnInit_gofutuapi_d667299e3498665a() {
  FTAPIChannel_UnInit();
  
}


FTAPIChannelPtr _wrap_CreateFTAPIChannel_gofutuapi_d667299e3498665a() {
  FTAPIChannelPtr result;
  FTAPIChannelPtr _swig_go_result;
  
  
  result = (FTAPIChannelPtr)CreateFTAPIChannel();
  *(FTAPIChannelPtr *)&_swig_go_result = (FTAPIChannelPtr)result; 
  return _swig_go_result;
}


void _wrap_ReleaseFTAPIChannel_gofutuapi_d667299e3498665a(FTAPIChannelPtr _swig_go_0) {
  FTAPIChannelPtr arg1 = (FTAPIChannelPtr) 0 ;
  
  arg1 = *(FTAPIChannelPtr *)&_swig_go_0; 
  
  ReleaseFTAPIChannel(arg1);
  
}


void _wrap_FTAPIChannel_SetClientInfo_gofutuapi_d667299e3498665a(FTAPIChannelPtr _swig_go_0, _gostring_ _swig_go_1, intgo _swig_go_2) {
  FTAPIChannelPtr arg1 = (FTAPIChannelPtr) 0 ;
  char *arg2 = (char *) 0 ;
  Futu::i32_t arg3 ;
  
  arg1 = *(FTAPIChannelPtr *)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  arg3 = (Futu::i32_t)_swig_go_2; 
  
  FTAPIChannel_SetClientInfo(arg1,(char const *)arg2,arg3);
  
  free(arg2); 
}


void _wrap_FTAPIChannel_SetRSAPrivateKey_gofutuapi_d667299e3498665a(FTAPIChannelPtr _swig_go_0, _gostring_ _swig_go_1) {
  FTAPIChannelPtr arg1 = (FTAPIChannelPtr) 0 ;
  char *arg2 = (char *) 0 ;
  
  arg1 = *(FTAPIChannelPtr *)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  
  FTAPIChannel_SetRSAPrivateKey(arg1,(char const *)arg2);
  
  free(arg2); 
}


intgo _wrap_FTAPIChannel_InitConnect_gofutuapi_d667299e3498665a(FTAPIChannelPtr _swig_go_0, _gostring_ _swig_go_1, short _swig_go_2, intgo _swig_go_3) {
  FTAPIChannelPtr arg1 = (FTAPIChannelPtr) 0 ;
  char *arg2 = (char *) 0 ;
  Futu::u16_t arg3 ;
  Futu::i32_t arg4 ;
  Futu::i32_t result;
  intgo _swig_go_result;
  
  arg1 = *(FTAPIChannelPtr *)&_swig_go_0; 
  
  arg2 = (char *)malloc(_swig_go_1.n + 1);
  memcpy(arg2, _swig_go_1.p, _swig_go_1.n);
  arg2[_swig_go_1.n] = '\0';
  
  arg3 = (Futu::u16_t)_swig_go_2; 
  arg4 = (Futu::i32_t)_swig_go_3; 
  
  result = (Futu::i32_t)FTAPIChannel_InitConnect(arg1,(char const *)arg2,arg3,arg4);
  _swig_go_result = result; 
  free(arg2); 
  return _swig_go_result;
}


long long _wrap_FTAPIChannel_GetConnectID_gofutuapi_d667299e3498665a(FTAPIChannelPtr _swig_go_0) {
  FTAPIChannelPtr arg1 = (FTAPIChannelPtr) 0 ;
  Futu::u64_t result;
  long long _swig_go_result;
  
  arg1 = *(FTAPIChannelPtr *)&_swig_go_0; 
  
  result = (Futu::u64_t)FTAPIChannel_GetConnectID(arg1);
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_FTAPIChannel_SendProto_gofutuapi_d667299e3498665a(FTAPIChannelPtr _swig_go_0, intgo _swig_go_1, char _swig_go_2, _gostring_ _swig_go_3) {
  FTAPIChannelPtr arg1 = (FTAPIChannelPtr) 0 ;
  Futu::u32_t arg2 ;
  Futu::u8_t arg3 ;
  Futu::i8_t *arg4 = (Futu::i8_t *) 0 ;
  Futu::i32_t arg5 ;
  Futu::u32_t result;
  intgo _swig_go_result;
  
  arg1 = *(FTAPIChannelPtr *)&_swig_go_0; 
  arg2 = (Futu::u32_t)_swig_go_1; 
  arg3 = (Futu::u8_t)_swig_go_2; 
  
  arg4 = (Futu::i8_t *) _swig_go_3.p;
  arg5 = (Futu::i32_t) _swig_go_3.n;
  
  
  result = (Futu::u32_t)FTAPIChannel_SendProto(arg1,arg2,arg3,(char const *)arg4,arg5);
  _swig_go_result = result; 
  return _swig_go_result;
}


intgo _wrap_FTAPIChannel_Close_gofutuapi_d667299e3498665a(FTAPIChannelPtr _swig_go_0) {
  FTAPIChannelPtr arg1 = (FTAPIChannelPtr) 0 ;
  Futu::i32_t result;
  intgo _swig_go_result;
  
  arg1 = *(FTAPIChannelPtr *)&_swig_go_0; 
  
  result = (Futu::i32_t)FTAPIChannel_Close(arg1);
  _swig_go_result = result; 
  return _swig_go_result;
}


void _wrap_wrap_InitializeCallbacks_gofutuapi_d667299e3498665a(FTAPIChannelPtr _swig_go_0) {
  FTAPIChannelPtr arg1 = (FTAPIChannelPtr) 0 ;
  
  arg1 = *(FTAPIChannelPtr *)&_swig_go_0; 
  
  wrap_InitializeCallbacks(arg1);
  
}


FTAPIChannel_Callback *_wrap__swig_NewDirectorFTAPIChannel_CallbackFTAPIChannel_Callback_gofutuapi_d667299e3498665a(intgo _swig_go_0) {
  int arg1 ;
  FTAPIChannel_Callback *result = 0 ;
  FTAPIChannel_Callback *_swig_go_result;
  
  arg1 = (int)_swig_go_0; 
  
  result = new SwigDirector_FTAPIChannel_Callback(arg1);
  *(FTAPIChannel_Callback **)&_swig_go_result = (FTAPIChannel_Callback *)result; 
  return _swig_go_result;
}


void _wrap_DeleteDirectorFTAPIChannel_Callback_gofutuapi_d667299e3498665a(FTAPIChannel_Callback *_swig_go_0) {
  FTAPIChannel_Callback *arg1 = (FTAPIChannel_Callback *) 0 ;
  
  arg1 = *(FTAPIChannel_Callback **)&_swig_go_0; 
  
  delete arg1;
  
}


void _wrap_FTAPIChannel_Callback_OnDisConnect_gofutuapi_d667299e3498665a(FTAPIChannel_Callback *_swig_go_0, FTAPIChannelPtr _swig_go_1, long long _swig_go_2) {
  FTAPIChannel_Callback *arg1 = (FTAPIChannel_Callback *) 0 ;
  FTAPIChannelPtr arg2 = (FTAPIChannelPtr) 0 ;
  Futu::i64_t arg3 ;
  
  arg1 = *(FTAPIChannel_Callback **)&_swig_go_0; 
  arg2 = *(FTAPIChannelPtr *)&_swig_go_1; 
  arg3 = (Futu::i64_t)_swig_go_2; 
  
  (arg1)->OnDisConnect(arg2,arg3);
  
}


void _wrap_FTAPIChannel_Callback_OnInitConnect_gofutuapi_d667299e3498665a(FTAPIChannel_Callback *_swig_go_0, FTAPIChannelPtr _swig_go_1, long long _swig_go_2, _gostring_ _swig_go_3) {
  FTAPIChannel_Callback *arg1 = (FTAPIChannel_Callback *) 0 ;
  FTAPIChannelPtr arg2 = (FTAPIChannelPtr) 0 ;
  Futu::i64_t arg3 ;
  char *arg4 = (char *) 0 ;
  
  arg1 = *(FTAPIChannel_Callback **)&_swig_go_0; 
  arg2 = *(FTAPIChannelPtr *)&_swig_go_1; 
  arg3 = (Futu::i64_t)_swig_go_2; 
  
  arg4 = (char *)malloc(_swig_go_3.n + 1);
  memcpy(arg4, _swig_go_3.p, _swig_go_3.n);
  arg4[_swig_go_3.n] = '\0';
  
  
  (arg1)->OnInitConnect(arg2,arg3,(char const *)arg4);
  
  free(arg4); 
}


void _wrap_FTAPIChannel_Callback_OnReply_gofutuapi_d667299e3498665a(FTAPIChannel_Callback *_swig_go_0, FTAPIChannelPtr _swig_go_1, intgo _swig_go_2, FTAPI_ProtoHeader *_swig_go_3, _gostring_ _swig_go_4) {
  FTAPIChannel_Callback *arg1 = (FTAPIChannel_Callback *) 0 ;
  FTAPIChannelPtr arg2 = (FTAPIChannelPtr) 0 ;
  FTAPI_ReqReplyType arg3 ;
  FTAPI_ProtoHeader *arg4 = (FTAPI_ProtoHeader *) 0 ;
  Futu::i8_t *arg5 = (Futu::i8_t *) 0 ;
  Futu::i32_t arg6 ;
  
  arg1 = *(FTAPIChannel_Callback **)&_swig_go_0; 
  arg2 = *(FTAPIChannelPtr *)&_swig_go_1; 
  arg3 = (FTAPI_ReqReplyType)_swig_go_2; 
  arg4 = *(FTAPI_ProtoHeader **)&_swig_go_3; 
  
  arg5 = (Futu::i8_t *) _swig_go_4.p;
  arg6 = (Futu::i32_t) _swig_go_4.n;
  
  
  (arg1)->OnReply(arg2,arg3,(FTAPI_ProtoHeader const *)arg4,(Futu::i8_t const *)arg5,arg6);
  
}


void _wrap_FTAPIChannel_Callback_OnPush_gofutuapi_d667299e3498665a(FTAPIChannel_Callback *_swig_go_0, FTAPIChannelPtr _swig_go_1, FTAPI_ProtoHeader *_swig_go_2, _gostring_ _swig_go_3) {
  FTAPIChannel_Callback *arg1 = (FTAPIChannel_Callback *) 0 ;
  FTAPIChannelPtr arg2 = (FTAPIChannelPtr) 0 ;
  FTAPI_ProtoHeader *arg3 = (FTAPI_ProtoHeader *) 0 ;
  Futu::i8_t *arg4 = (Futu::i8_t *) 0 ;
  Futu::i32_t arg5 ;
  
  arg1 = *(FTAPIChannel_Callback **)&_swig_go_0; 
  arg2 = *(FTAPIChannelPtr *)&_swig_go_1; 
  arg3 = *(FTAPI_ProtoHeader **)&_swig_go_2; 
  
  arg4 = (Futu::i8_t *) _swig_go_3.p;
  arg5 = (Futu::i32_t) _swig_go_3.n;
  
  
  (arg1)->OnPush(arg2,(FTAPI_ProtoHeader const *)arg3,(Futu::i8_t const *)arg4,arg5);
  
}


void _wrap_delete_FTAPIChannel_Callback_gofutuapi_d667299e3498665a(FTAPIChannel_Callback *_swig_go_0) {
  FTAPIChannel_Callback *arg1 = (FTAPIChannel_Callback *) 0 ;
  
  arg1 = *(FTAPIChannel_Callback **)&_swig_go_0; 
  
  delete arg1;
  
}


#ifdef __cplusplus
}
#endif

 
  //printf("Initialization gofutuapi done.\n");
