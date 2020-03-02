// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Qot_UpdateKL.proto

#include "Qot_UpdateKL.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/stubs/port.h>
#include <google/protobuf/stubs/once.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// This is a temporary google only hack
#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
#include "third_party/protobuf/version.h"
#endif
// @@protoc_insertion_point(includes)
namespace Qot_UpdateKL {
class S2CDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<S2C>
      _instance;
} _S2C_default_instance_;
class ResponseDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<Response>
      _instance;
} _Response_default_instance_;
}  // namespace Qot_UpdateKL
namespace protobuf_Qot_5fUpdateKL_2eproto {
void InitDefaultsS2CImpl() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  ::google::protobuf::internal::InitProtobufDefaultsForceUnique();
#else
  ::google::protobuf::internal::InitProtobufDefaults();
#endif  // GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  protobuf_Qot_5fCommon_2eproto::InitDefaultsSecurity();
  protobuf_Qot_5fCommon_2eproto::InitDefaultsKLine();
  {
    void* ptr = &::Qot_UpdateKL::_S2C_default_instance_;
    new (ptr) ::Qot_UpdateKL::S2C();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::Qot_UpdateKL::S2C::InitAsDefaultInstance();
}

void InitDefaultsS2C() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &InitDefaultsS2CImpl);
}

void InitDefaultsResponseImpl() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

#ifdef GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  ::google::protobuf::internal::InitProtobufDefaultsForceUnique();
#else
  ::google::protobuf::internal::InitProtobufDefaults();
#endif  // GOOGLE_PROTOBUF_ENFORCE_UNIQUENESS
  protobuf_Qot_5fUpdateKL_2eproto::InitDefaultsS2C();
  {
    void* ptr = &::Qot_UpdateKL::_Response_default_instance_;
    new (ptr) ::Qot_UpdateKL::Response();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::Qot_UpdateKL::Response::InitAsDefaultInstance();
}

void InitDefaultsResponse() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &InitDefaultsResponseImpl);
}

::google::protobuf::Metadata file_level_metadata[2];

const ::google::protobuf::uint32 TableStruct::offsets[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::S2C, _has_bits_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::S2C, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::S2C, rehabtype_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::S2C, kltype_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::S2C, security_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::S2C, kllist_),
  1,
  2,
  0,
  ~0u,
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::Response, _has_bits_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::Response, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::Response, rettype_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::Response, retmsg_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::Response, errcode_),
  GOOGLE_PROTOBUF_GENERATED_MESSAGE_FIELD_OFFSET(::Qot_UpdateKL::Response, s2c_),
  3,
  0,
  2,
  1,
};
static const ::google::protobuf::internal::MigrationSchema schemas[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
  { 0, 9, sizeof(::Qot_UpdateKL::S2C)},
  { 13, 22, sizeof(::Qot_UpdateKL::Response)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&::Qot_UpdateKL::_S2C_default_instance_),
  reinterpret_cast<const ::google::protobuf::Message*>(&::Qot_UpdateKL::_Response_default_instance_),
};

void protobuf_AssignDescriptors() {
  AddDescriptors();
  ::google::protobuf::MessageFactory* factory = NULL;
  AssignDescriptors(
      "Qot_UpdateKL.proto", schemas, file_default_instances, TableStruct::offsets, factory,
      file_level_metadata, NULL, NULL);
}

void protobuf_AssignDescriptorsOnce() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &protobuf_AssignDescriptors);
}

void protobuf_RegisterTypes(const ::std::string&) GOOGLE_PROTOBUF_ATTRIBUTE_COLD;
void protobuf_RegisterTypes(const ::std::string&) {
  protobuf_AssignDescriptorsOnce();
  ::google::protobuf::internal::RegisterAllTypes(file_level_metadata, 2);
}

void AddDescriptorsImpl() {
  InitDefaults();
  static const char descriptor[] GOOGLE_PROTOBUF_ATTRIBUTE_SECTION_VARIABLE(protodesc_cold) = {
      "\n\022Qot_UpdateKL.proto\022\014Qot_UpdateKL\032\014Comm"
      "on.proto\032\020Qot_Common.proto\"s\n\003S2C\022\021\n\treh"
      "abType\030\001 \002(\005\022\016\n\006klType\030\002 \002(\005\022&\n\010security"
      "\030\003 \002(\0132\024.Qot_Common.Security\022!\n\006klList\030\004"
      " \003(\0132\021.Qot_Common.KLine\"b\n\010Response\022\025\n\007r"
      "etType\030\001 \002(\005:\004-400\022\016\n\006retMsg\030\002 \001(\t\022\017\n\007er"
      "rCode\030\003 \001(\005\022\036\n\003s2c\030\004 \001(\0132\021.Qot_UpdateKL."
      "S2CB\025\n\023com.futu.openapi.pb"
  };
  ::google::protobuf::DescriptorPool::InternalAddGeneratedFile(
      descriptor, 306);
  ::google::protobuf::MessageFactory::InternalRegisterGeneratedFile(
    "Qot_UpdateKL.proto", &protobuf_RegisterTypes);
  ::protobuf_Common_2eproto::AddDescriptors();
  ::protobuf_Qot_5fCommon_2eproto::AddDescriptors();
}

void AddDescriptors() {
  static GOOGLE_PROTOBUF_DECLARE_ONCE(once);
  ::google::protobuf::GoogleOnceInit(&once, &AddDescriptorsImpl);
}
// Force AddDescriptors() to be called at dynamic initialization time.
struct StaticDescriptorInitializer {
  StaticDescriptorInitializer() {
    AddDescriptors();
  }
} static_descriptor_initializer;
}  // namespace protobuf_Qot_5fUpdateKL_2eproto
namespace Qot_UpdateKL {

// ===================================================================

void S2C::InitAsDefaultInstance() {
  ::Qot_UpdateKL::_S2C_default_instance_._instance.get_mutable()->security_ = const_cast< ::Qot_Common::Security*>(
      ::Qot_Common::Security::internal_default_instance());
}
void S2C::clear_security() {
  if (security_ != NULL) security_->Clear();
  clear_has_security();
}
void S2C::clear_kllist() {
  kllist_.Clear();
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int S2C::kRehabTypeFieldNumber;
const int S2C::kKlTypeFieldNumber;
const int S2C::kSecurityFieldNumber;
const int S2C::kKlListFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

S2C::S2C()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  if (GOOGLE_PREDICT_TRUE(this != internal_default_instance())) {
    ::protobuf_Qot_5fUpdateKL_2eproto::InitDefaultsS2C();
  }
  SharedCtor();
  // @@protoc_insertion_point(constructor:Qot_UpdateKL.S2C)
}
S2C::S2C(const S2C& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL),
      _has_bits_(from._has_bits_),
      _cached_size_(0),
      kllist_(from.kllist_) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  if (from.has_security()) {
    security_ = new ::Qot_Common::Security(*from.security_);
  } else {
    security_ = NULL;
  }
  ::memcpy(&rehabtype_, &from.rehabtype_,
    static_cast<size_t>(reinterpret_cast<char*>(&kltype_) -
    reinterpret_cast<char*>(&rehabtype_)) + sizeof(kltype_));
  // @@protoc_insertion_point(copy_constructor:Qot_UpdateKL.S2C)
}

void S2C::SharedCtor() {
  _cached_size_ = 0;
  ::memset(&security_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&kltype_) -
      reinterpret_cast<char*>(&security_)) + sizeof(kltype_));
}

S2C::~S2C() {
  // @@protoc_insertion_point(destructor:Qot_UpdateKL.S2C)
  SharedDtor();
}

void S2C::SharedDtor() {
  if (this != internal_default_instance()) delete security_;
}

void S2C::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* S2C::descriptor() {
  ::protobuf_Qot_5fUpdateKL_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_Qot_5fUpdateKL_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const S2C& S2C::default_instance() {
  ::protobuf_Qot_5fUpdateKL_2eproto::InitDefaultsS2C();
  return *internal_default_instance();
}

S2C* S2C::New(::google::protobuf::Arena* arena) const {
  S2C* n = new S2C;
  if (arena != NULL) {
    arena->Own(n);
  }
  return n;
}

void S2C::Clear() {
// @@protoc_insertion_point(message_clear_start:Qot_UpdateKL.S2C)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  kllist_.Clear();
  cached_has_bits = _has_bits_[0];
  if (cached_has_bits & 0x00000001u) {
    GOOGLE_DCHECK(security_ != NULL);
    security_->Clear();
  }
  if (cached_has_bits & 6u) {
    ::memset(&rehabtype_, 0, static_cast<size_t>(
        reinterpret_cast<char*>(&kltype_) -
        reinterpret_cast<char*>(&rehabtype_)) + sizeof(kltype_));
  }
  _has_bits_.Clear();
  _internal_metadata_.Clear();
}

bool S2C::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:Qot_UpdateKL.S2C)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // required int32 rehabType = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(8u /* 8 & 0xFF */)) {
          set_has_rehabtype();
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &rehabtype_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // required int32 klType = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(16u /* 16 & 0xFF */)) {
          set_has_kltype();
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &kltype_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // required .Qot_Common.Security security = 3;
      case 3: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(26u /* 26 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessage(
               input, mutable_security()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // repeated .Qot_Common.KLine klList = 4;
      case 4: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(34u /* 34 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessage(input, add_kllist()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:Qot_UpdateKL.S2C)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:Qot_UpdateKL.S2C)
  return false;
#undef DO_
}

void S2C::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:Qot_UpdateKL.S2C)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  cached_has_bits = _has_bits_[0];
  // required int32 rehabType = 1;
  if (cached_has_bits & 0x00000002u) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(1, this->rehabtype(), output);
  }

  // required int32 klType = 2;
  if (cached_has_bits & 0x00000004u) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(2, this->kltype(), output);
  }

  // required .Qot_Common.Security security = 3;
  if (cached_has_bits & 0x00000001u) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      3, *this->security_, output);
  }

  // repeated .Qot_Common.KLine klList = 4;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->kllist_size()); i < n; i++) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      4, this->kllist(static_cast<int>(i)), output);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        _internal_metadata_.unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:Qot_UpdateKL.S2C)
}

::google::protobuf::uint8* S2C::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:Qot_UpdateKL.S2C)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  cached_has_bits = _has_bits_[0];
  // required int32 rehabType = 1;
  if (cached_has_bits & 0x00000002u) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(1, this->rehabtype(), target);
  }

  // required int32 klType = 2;
  if (cached_has_bits & 0x00000004u) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(2, this->kltype(), target);
  }

  // required .Qot_Common.Security security = 3;
  if (cached_has_bits & 0x00000001u) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageToArray(
        3, *this->security_, deterministic, target);
  }

  // repeated .Qot_Common.KLine klList = 4;
  for (unsigned int i = 0,
      n = static_cast<unsigned int>(this->kllist_size()); i < n; i++) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageToArray(
        4, this->kllist(static_cast<int>(i)), deterministic, target);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:Qot_UpdateKL.S2C)
  return target;
}

size_t S2C::RequiredFieldsByteSizeFallback() const {
// @@protoc_insertion_point(required_fields_byte_size_fallback_start:Qot_UpdateKL.S2C)
  size_t total_size = 0;

  if (has_security()) {
    // required .Qot_Common.Security security = 3;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::MessageSize(
        *this->security_);
  }

  if (has_rehabtype()) {
    // required int32 rehabType = 1;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->rehabtype());
  }

  if (has_kltype()) {
    // required int32 klType = 2;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->kltype());
  }

  return total_size;
}
size_t S2C::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:Qot_UpdateKL.S2C)
  size_t total_size = 0;

  if (_internal_metadata_.have_unknown_fields()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        _internal_metadata_.unknown_fields());
  }
  if (((_has_bits_[0] & 0x00000007) ^ 0x00000007) == 0) {  // All required fields are present.
    // required .Qot_Common.Security security = 3;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::MessageSize(
        *this->security_);

    // required int32 rehabType = 1;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->rehabtype());

    // required int32 klType = 2;
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->kltype());

  } else {
    total_size += RequiredFieldsByteSizeFallback();
  }
  // repeated .Qot_Common.KLine klList = 4;
  {
    unsigned int count = static_cast<unsigned int>(this->kllist_size());
    total_size += 1UL * count;
    for (unsigned int i = 0; i < count; i++) {
      total_size +=
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          this->kllist(static_cast<int>(i)));
    }
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = cached_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void S2C::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:Qot_UpdateKL.S2C)
  GOOGLE_DCHECK_NE(&from, this);
  const S2C* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const S2C>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:Qot_UpdateKL.S2C)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:Qot_UpdateKL.S2C)
    MergeFrom(*source);
  }
}

void S2C::MergeFrom(const S2C& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:Qot_UpdateKL.S2C)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  kllist_.MergeFrom(from.kllist_);
  cached_has_bits = from._has_bits_[0];
  if (cached_has_bits & 7u) {
    if (cached_has_bits & 0x00000001u) {
      mutable_security()->::Qot_Common::Security::MergeFrom(from.security());
    }
    if (cached_has_bits & 0x00000002u) {
      rehabtype_ = from.rehabtype_;
    }
    if (cached_has_bits & 0x00000004u) {
      kltype_ = from.kltype_;
    }
    _has_bits_[0] |= cached_has_bits;
  }
}

void S2C::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:Qot_UpdateKL.S2C)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void S2C::CopyFrom(const S2C& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:Qot_UpdateKL.S2C)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool S2C::IsInitialized() const {
  if ((_has_bits_[0] & 0x00000007) != 0x00000007) return false;
  if (!::google::protobuf::internal::AllAreInitialized(this->kllist())) return false;
  if (has_security()) {
    if (!this->security_->IsInitialized()) return false;
  }
  return true;
}

void S2C::Swap(S2C* other) {
  if (other == this) return;
  InternalSwap(other);
}
void S2C::InternalSwap(S2C* other) {
  using std::swap;
  kllist_.InternalSwap(&other->kllist_);
  swap(security_, other->security_);
  swap(rehabtype_, other->rehabtype_);
  swap(kltype_, other->kltype_);
  swap(_has_bits_[0], other->_has_bits_[0]);
  _internal_metadata_.Swap(&other->_internal_metadata_);
  swap(_cached_size_, other->_cached_size_);
}

::google::protobuf::Metadata S2C::GetMetadata() const {
  protobuf_Qot_5fUpdateKL_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_Qot_5fUpdateKL_2eproto::file_level_metadata[kIndexInFileMessages];
}


// ===================================================================

void Response::InitAsDefaultInstance() {
  ::Qot_UpdateKL::_Response_default_instance_._instance.get_mutable()->s2c_ = const_cast< ::Qot_UpdateKL::S2C*>(
      ::Qot_UpdateKL::S2C::internal_default_instance());
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int Response::kRetTypeFieldNumber;
const int Response::kRetMsgFieldNumber;
const int Response::kErrCodeFieldNumber;
const int Response::kS2CFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

Response::Response()
  : ::google::protobuf::Message(), _internal_metadata_(NULL) {
  if (GOOGLE_PREDICT_TRUE(this != internal_default_instance())) {
    ::protobuf_Qot_5fUpdateKL_2eproto::InitDefaultsResponse();
  }
  SharedCtor();
  // @@protoc_insertion_point(constructor:Qot_UpdateKL.Response)
}
Response::Response(const Response& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(NULL),
      _has_bits_(from._has_bits_),
      _cached_size_(0) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  retmsg_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (from.has_retmsg()) {
    retmsg_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.retmsg_);
  }
  if (from.has_s2c()) {
    s2c_ = new ::Qot_UpdateKL::S2C(*from.s2c_);
  } else {
    s2c_ = NULL;
  }
  ::memcpy(&errcode_, &from.errcode_,
    static_cast<size_t>(reinterpret_cast<char*>(&rettype_) -
    reinterpret_cast<char*>(&errcode_)) + sizeof(rettype_));
  // @@protoc_insertion_point(copy_constructor:Qot_UpdateKL.Response)
}

void Response::SharedCtor() {
  _cached_size_ = 0;
  retmsg_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  ::memset(&s2c_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&errcode_) -
      reinterpret_cast<char*>(&s2c_)) + sizeof(errcode_));
  rettype_ = -400;
}

Response::~Response() {
  // @@protoc_insertion_point(destructor:Qot_UpdateKL.Response)
  SharedDtor();
}

void Response::SharedDtor() {
  retmsg_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (this != internal_default_instance()) delete s2c_;
}

void Response::SetCachedSize(int size) const {
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
}
const ::google::protobuf::Descriptor* Response::descriptor() {
  ::protobuf_Qot_5fUpdateKL_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_Qot_5fUpdateKL_2eproto::file_level_metadata[kIndexInFileMessages].descriptor;
}

const Response& Response::default_instance() {
  ::protobuf_Qot_5fUpdateKL_2eproto::InitDefaultsResponse();
  return *internal_default_instance();
}

Response* Response::New(::google::protobuf::Arena* arena) const {
  Response* n = new Response;
  if (arena != NULL) {
    arena->Own(n);
  }
  return n;
}

void Response::Clear() {
// @@protoc_insertion_point(message_clear_start:Qot_UpdateKL.Response)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  cached_has_bits = _has_bits_[0];
  if (cached_has_bits & 3u) {
    if (cached_has_bits & 0x00000001u) {
      GOOGLE_DCHECK(!retmsg_.IsDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited()));
      (*retmsg_.UnsafeRawStringPointer())->clear();
    }
    if (cached_has_bits & 0x00000002u) {
      GOOGLE_DCHECK(s2c_ != NULL);
      s2c_->Clear();
    }
  }
  if (cached_has_bits & 12u) {
    errcode_ = 0;
    rettype_ = -400;
  }
  _has_bits_.Clear();
  _internal_metadata_.Clear();
}

bool Response::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!GOOGLE_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:Qot_UpdateKL.Response)
  for (;;) {
    ::std::pair< ::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // required int32 retType = 1 [default = -400];
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(8u /* 8 & 0xFF */)) {
          set_has_rettype();
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &rettype_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // optional string retMsg = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(18u /* 18 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_retmsg()));
          ::google::protobuf::internal::WireFormat::VerifyUTF8StringNamedField(
            this->retmsg().data(), static_cast<int>(this->retmsg().length()),
            ::google::protobuf::internal::WireFormat::PARSE,
            "Qot_UpdateKL.Response.retMsg");
        } else {
          goto handle_unusual;
        }
        break;
      }

      // optional int32 errCode = 3;
      case 3: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(24u /* 24 & 0xFF */)) {
          set_has_errcode();
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &errcode_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // optional .Qot_UpdateKL.S2C s2c = 4;
      case 4: {
        if (static_cast< ::google::protobuf::uint8>(tag) ==
            static_cast< ::google::protobuf::uint8>(34u /* 34 & 0xFF */)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessage(
               input, mutable_s2c()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:Qot_UpdateKL.Response)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:Qot_UpdateKL.Response)
  return false;
#undef DO_
}

void Response::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:Qot_UpdateKL.Response)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  cached_has_bits = _has_bits_[0];
  // required int32 retType = 1 [default = -400];
  if (cached_has_bits & 0x00000008u) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(1, this->rettype(), output);
  }

  // optional string retMsg = 2;
  if (cached_has_bits & 0x00000001u) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8StringNamedField(
      this->retmsg().data(), static_cast<int>(this->retmsg().length()),
      ::google::protobuf::internal::WireFormat::SERIALIZE,
      "Qot_UpdateKL.Response.retMsg");
    ::google::protobuf::internal::WireFormatLite::WriteStringMaybeAliased(
      2, this->retmsg(), output);
  }

  // optional int32 errCode = 3;
  if (cached_has_bits & 0x00000004u) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(3, this->errcode(), output);
  }

  // optional .Qot_UpdateKL.S2C s2c = 4;
  if (cached_has_bits & 0x00000002u) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      4, *this->s2c_, output);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        _internal_metadata_.unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:Qot_UpdateKL.Response)
}

::google::protobuf::uint8* Response::InternalSerializeWithCachedSizesToArray(
    bool deterministic, ::google::protobuf::uint8* target) const {
  (void)deterministic; // Unused
  // @@protoc_insertion_point(serialize_to_array_start:Qot_UpdateKL.Response)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  cached_has_bits = _has_bits_[0];
  // required int32 retType = 1 [default = -400];
  if (cached_has_bits & 0x00000008u) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(1, this->rettype(), target);
  }

  // optional string retMsg = 2;
  if (cached_has_bits & 0x00000001u) {
    ::google::protobuf::internal::WireFormat::VerifyUTF8StringNamedField(
      this->retmsg().data(), static_cast<int>(this->retmsg().length()),
      ::google::protobuf::internal::WireFormat::SERIALIZE,
      "Qot_UpdateKL.Response.retMsg");
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        2, this->retmsg(), target);
  }

  // optional int32 errCode = 3;
  if (cached_has_bits & 0x00000004u) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(3, this->errcode(), target);
  }

  // optional .Qot_UpdateKL.S2C s2c = 4;
  if (cached_has_bits & 0x00000002u) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageToArray(
        4, *this->s2c_, deterministic, target);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:Qot_UpdateKL.Response)
  return target;
}

size_t Response::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:Qot_UpdateKL.Response)
  size_t total_size = 0;

  if (_internal_metadata_.have_unknown_fields()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        _internal_metadata_.unknown_fields());
  }
  // required int32 retType = 1 [default = -400];
  if (has_rettype()) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->rettype());
  }
  if (_has_bits_[0 / 32] & 7u) {
    // optional string retMsg = 2;
    if (has_retmsg()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::StringSize(
          this->retmsg());
    }

    // optional .Qot_UpdateKL.S2C s2c = 4;
    if (has_s2c()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::MessageSize(
          *this->s2c_);
    }

    // optional int32 errCode = 3;
    if (has_errcode()) {
      total_size += 1 +
        ::google::protobuf::internal::WireFormatLite::Int32Size(
          this->errcode());
    }

  }
  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  GOOGLE_SAFE_CONCURRENT_WRITES_BEGIN();
  _cached_size_ = cached_size;
  GOOGLE_SAFE_CONCURRENT_WRITES_END();
  return total_size;
}

void Response::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:Qot_UpdateKL.Response)
  GOOGLE_DCHECK_NE(&from, this);
  const Response* source =
      ::google::protobuf::internal::DynamicCastToGenerated<const Response>(
          &from);
  if (source == NULL) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:Qot_UpdateKL.Response)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:Qot_UpdateKL.Response)
    MergeFrom(*source);
  }
}

void Response::MergeFrom(const Response& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:Qot_UpdateKL.Response)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  cached_has_bits = from._has_bits_[0];
  if (cached_has_bits & 15u) {
    if (cached_has_bits & 0x00000001u) {
      set_has_retmsg();
      retmsg_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.retmsg_);
    }
    if (cached_has_bits & 0x00000002u) {
      mutable_s2c()->::Qot_UpdateKL::S2C::MergeFrom(from.s2c());
    }
    if (cached_has_bits & 0x00000004u) {
      errcode_ = from.errcode_;
    }
    if (cached_has_bits & 0x00000008u) {
      rettype_ = from.rettype_;
    }
    _has_bits_[0] |= cached_has_bits;
  }
}

void Response::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:Qot_UpdateKL.Response)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void Response::CopyFrom(const Response& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:Qot_UpdateKL.Response)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Response::IsInitialized() const {
  if ((_has_bits_[0] & 0x00000008) != 0x00000008) return false;
  if (has_s2c()) {
    if (!this->s2c_->IsInitialized()) return false;
  }
  return true;
}

void Response::Swap(Response* other) {
  if (other == this) return;
  InternalSwap(other);
}
void Response::InternalSwap(Response* other) {
  using std::swap;
  retmsg_.Swap(&other->retmsg_);
  swap(s2c_, other->s2c_);
  swap(errcode_, other->errcode_);
  swap(rettype_, other->rettype_);
  swap(_has_bits_[0], other->_has_bits_[0]);
  _internal_metadata_.Swap(&other->_internal_metadata_);
  swap(_cached_size_, other->_cached_size_);
}

::google::protobuf::Metadata Response::GetMetadata() const {
  protobuf_Qot_5fUpdateKL_2eproto::protobuf_AssignDescriptorsOnce();
  return ::protobuf_Qot_5fUpdateKL_2eproto::file_level_metadata[kIndexInFileMessages];
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace Qot_UpdateKL

// @@protoc_insertion_point(global_scope)
