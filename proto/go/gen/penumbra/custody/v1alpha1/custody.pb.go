// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: penumbra/custody/v1alpha1/custody.proto

package custodyv1alpha1

import (
	v1alpha11 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/crypto/v1alpha1"
	v1alpha1 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/transaction/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthorizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The transaction plan to authorize.
	Plan *v1alpha1.TransactionPlan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
	// Identifies the FVK (and hence the spend authorization key) to use for signing.
	AccountGroupId *v1alpha11.AccountGroupId `protobuf:"bytes,2,opt,name=account_group_id,json=accountGroupId,proto3" json:"account_group_id,omitempty"`
	// Optionally, pre-authorization data, if required by the custodian.
	//
	// Multiple `PreAuthorization` packets can be included in a single request,
	// to support multi-party pre-authorizations.
	PreAuthorizations []*PreAuthorization `protobuf:"bytes,3,rep,name=pre_authorizations,json=preAuthorizations,proto3" json:"pre_authorizations,omitempty"`
}

func (x *AuthorizeRequest) Reset() {
	*x = AuthorizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeRequest) ProtoMessage() {}

func (x *AuthorizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizeRequest) GetPlan() *v1alpha1.TransactionPlan {
	if x != nil {
		return x.Plan
	}
	return nil
}

func (x *AuthorizeRequest) GetAccountGroupId() *v1alpha11.AccountGroupId {
	if x != nil {
		return x.AccountGroupId
	}
	return nil
}

func (x *AuthorizeRequest) GetPreAuthorizations() []*PreAuthorization {
	if x != nil {
		return x.PreAuthorizations
	}
	return nil
}

type AuthorizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *v1alpha1.AuthorizationData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AuthorizeResponse) Reset() {
	*x = AuthorizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeResponse) ProtoMessage() {}

func (x *AuthorizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeResponse) GetData() *v1alpha1.AuthorizationData {
	if x != nil {
		return x.Data
	}
	return nil
}

// A pre-authorization packet.  This allows a custodian to delegate (partial)
// signing authority to other authorization mechanisms.  Details of how a
// custodian manages those keys are out-of-scope for the custody protocol and
// are custodian-specific.
type PreAuthorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PreAuthorization:
	//
	//	*PreAuthorization_Ed25519_
	PreAuthorization isPreAuthorization_PreAuthorization `protobuf_oneof:"pre_authorization"`
}

func (x *PreAuthorization) Reset() {
	*x = PreAuthorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreAuthorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreAuthorization) ProtoMessage() {}

func (x *PreAuthorization) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreAuthorization.ProtoReflect.Descriptor instead.
func (*PreAuthorization) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{2}
}

func (m *PreAuthorization) GetPreAuthorization() isPreAuthorization_PreAuthorization {
	if m != nil {
		return m.PreAuthorization
	}
	return nil
}

func (x *PreAuthorization) GetEd25519() *PreAuthorization_Ed25519 {
	if x, ok := x.GetPreAuthorization().(*PreAuthorization_Ed25519_); ok {
		return x.Ed25519
	}
	return nil
}

type isPreAuthorization_PreAuthorization interface {
	isPreAuthorization_PreAuthorization()
}

type PreAuthorization_Ed25519_ struct {
	Ed25519 *PreAuthorization_Ed25519 `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}

func (*PreAuthorization_Ed25519_) isPreAuthorization_PreAuthorization() {}

// An Ed25519-based preauthorization, containing an Ed25519 signature over the
// `TransactionPlan`.
type PreAuthorization_Ed25519 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Ed25519 verification key used to verify the signature.
	Vk []byte `protobuf:"bytes,1,opt,name=vk,proto3" json:"vk,omitempty"`
	// The Ed25519 signature over the `TransactionPlan`.
	Sig []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *PreAuthorization_Ed25519) Reset() {
	*x = PreAuthorization_Ed25519{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreAuthorization_Ed25519) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreAuthorization_Ed25519) ProtoMessage() {}

func (x *PreAuthorization_Ed25519) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_custody_v1alpha1_custody_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreAuthorization_Ed25519.ProtoReflect.Descriptor instead.
func (*PreAuthorization_Ed25519) Descriptor() ([]byte, []int) {
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PreAuthorization_Ed25519) GetVk() []byte {
	if x != nil {
		return x.Vk
	}
	return nil
}

func (x *PreAuthorization_Ed25519) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

var File_penumbra_custody_v1alpha1_custody_proto protoreflect.FileDescriptor

var file_penumbra_custody_v1alpha1_custody_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x64, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x70, 0x65, 0x6e, 0x75, 0x6d,
	0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x2a, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x34, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x02, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x70,
	0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04,
	0x70, 0x6c, 0x61, 0x6e, 0x12, 0x57, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x52, 0x0e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x5a, 0x0a,
	0x12, 0x70, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x70, 0x72, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x5e, 0x0a, 0x11, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70,
	0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa5, 0x01, 0x0a, 0x10, 0x50, 0x72,
	0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f,
	0x0a, 0x07, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x64, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x64, 0x32,
	0x35, 0x35, 0x31, 0x39, 0x48, 0x00, 0x52, 0x07, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x1a,
	0x2b, 0x0a, 0x07, 0x45, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x76, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x42, 0x13, 0x0a, 0x11,
	0x70, 0x72, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0x80, 0x01, 0x0a, 0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x09,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x2b, 0x2e, 0x70, 0x65, 0x6e, 0x75,
	0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x8d, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x65, 0x6e,
	0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2d, 0x7a, 0x6f, 0x6e, 0x65,
	0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xa2, 0x02, 0x03, 0x50, 0x43, 0x58, 0xaa, 0x02, 0x19, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x19, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x64, 0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x25, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64,
	0x79, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x3a, 0x3a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_penumbra_custody_v1alpha1_custody_proto_rawDescOnce sync.Once
	file_penumbra_custody_v1alpha1_custody_proto_rawDescData = file_penumbra_custody_v1alpha1_custody_proto_rawDesc
)

func file_penumbra_custody_v1alpha1_custody_proto_rawDescGZIP() []byte {
	file_penumbra_custody_v1alpha1_custody_proto_rawDescOnce.Do(func() {
		file_penumbra_custody_v1alpha1_custody_proto_rawDescData = protoimpl.X.CompressGZIP(file_penumbra_custody_v1alpha1_custody_proto_rawDescData)
	})
	return file_penumbra_custody_v1alpha1_custody_proto_rawDescData
}

var file_penumbra_custody_v1alpha1_custody_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_penumbra_custody_v1alpha1_custody_proto_goTypes = []interface{}{
	(*AuthorizeRequest)(nil),           // 0: penumbra.custody.v1alpha1.AuthorizeRequest
	(*AuthorizeResponse)(nil),          // 1: penumbra.custody.v1alpha1.AuthorizeResponse
	(*PreAuthorization)(nil),           // 2: penumbra.custody.v1alpha1.PreAuthorization
	(*PreAuthorization_Ed25519)(nil),   // 3: penumbra.custody.v1alpha1.PreAuthorization.Ed25519
	(*v1alpha1.TransactionPlan)(nil),   // 4: penumbra.core.transaction.v1alpha1.TransactionPlan
	(*v1alpha11.AccountGroupId)(nil),   // 5: penumbra.core.crypto.v1alpha1.AccountGroupId
	(*v1alpha1.AuthorizationData)(nil), // 6: penumbra.core.transaction.v1alpha1.AuthorizationData
}
var file_penumbra_custody_v1alpha1_custody_proto_depIdxs = []int32{
	4, // 0: penumbra.custody.v1alpha1.AuthorizeRequest.plan:type_name -> penumbra.core.transaction.v1alpha1.TransactionPlan
	5, // 1: penumbra.custody.v1alpha1.AuthorizeRequest.account_group_id:type_name -> penumbra.core.crypto.v1alpha1.AccountGroupId
	2, // 2: penumbra.custody.v1alpha1.AuthorizeRequest.pre_authorizations:type_name -> penumbra.custody.v1alpha1.PreAuthorization
	6, // 3: penumbra.custody.v1alpha1.AuthorizeResponse.data:type_name -> penumbra.core.transaction.v1alpha1.AuthorizationData
	3, // 4: penumbra.custody.v1alpha1.PreAuthorization.ed25519:type_name -> penumbra.custody.v1alpha1.PreAuthorization.Ed25519
	0, // 5: penumbra.custody.v1alpha1.CustodyProtocolService.Authorize:input_type -> penumbra.custody.v1alpha1.AuthorizeRequest
	1, // 6: penumbra.custody.v1alpha1.CustodyProtocolService.Authorize:output_type -> penumbra.custody.v1alpha1.AuthorizeResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_penumbra_custody_v1alpha1_custody_proto_init() }
func file_penumbra_custody_v1alpha1_custody_proto_init() {
	if File_penumbra_custody_v1alpha1_custody_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreAuthorization); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_custody_v1alpha1_custody_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreAuthorization_Ed25519); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_penumbra_custody_v1alpha1_custody_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PreAuthorization_Ed25519_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_penumbra_custody_v1alpha1_custody_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_penumbra_custody_v1alpha1_custody_proto_goTypes,
		DependencyIndexes: file_penumbra_custody_v1alpha1_custody_proto_depIdxs,
		MessageInfos:      file_penumbra_custody_v1alpha1_custody_proto_msgTypes,
	}.Build()
	File_penumbra_custody_v1alpha1_custody_proto = out.File
	file_penumbra_custody_v1alpha1_custody_proto_rawDesc = nil
	file_penumbra_custody_v1alpha1_custody_proto_goTypes = nil
	file_penumbra_custody_v1alpha1_custody_proto_depIdxs = nil
}