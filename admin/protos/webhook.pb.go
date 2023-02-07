// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: webhook.proto

package protos

import (
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

type WebhookSaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AddTokenCount int32  `protobuf:"varint,3,opt,name=add_token_count,json=addTokenCount,proto3" json:"add_token_count,omitempty"`
}

func (x *WebhookSaveReq) Reset() {
	*x = WebhookSaveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookSaveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookSaveReq) ProtoMessage() {}

func (x *WebhookSaveReq) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookSaveReq.ProtoReflect.Descriptor instead.
func (*WebhookSaveReq) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *WebhookSaveReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WebhookSaveReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WebhookSaveReq) GetAddTokenCount() int32 {
	if x != nil {
		return x.AddTokenCount
	}
	return 0
}

type WebhookRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkspaceId string                `protobuf:"bytes,1,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Id          string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt   int64                 `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   int64                 `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Tokens      []*WebhookTokenRecord `protobuf:"bytes,6,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *WebhookRecord) Reset() {
	*x = WebhookRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookRecord) ProtoMessage() {}

func (x *WebhookRecord) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookRecord.ProtoReflect.Descriptor instead.
func (*WebhookRecord) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{1}
}

func (x *WebhookRecord) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *WebhookRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WebhookRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WebhookRecord) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *WebhookRecord) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *WebhookRecord) GetTokens() []*WebhookTokenRecord {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type WebhookTokenRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebhookId string `protobuf:"bytes,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Token     string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	CreatedAt int64  `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *WebhookTokenRecord) Reset() {
	*x = WebhookTokenRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookTokenRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookTokenRecord) ProtoMessage() {}

func (x *WebhookTokenRecord) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookTokenRecord.ProtoReflect.Descriptor instead.
func (*WebhookTokenRecord) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{2}
}

func (x *WebhookTokenRecord) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

func (x *WebhookTokenRecord) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WebhookTokenRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WebhookTokenRecord) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *WebhookTokenRecord) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type WebhookGetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WebhookGetReq) Reset() {
	*x = WebhookGetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookGetReq) ProtoMessage() {}

func (x *WebhookGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookGetReq.ProtoReflect.Descriptor instead.
func (*WebhookGetReq) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{3}
}

func (x *WebhookGetReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type WebhookListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor string `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Size   int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *WebhookListReq) Reset() {
	*x = WebhookListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookListReq) ProtoMessage() {}

func (x *WebhookListReq) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookListReq.ProtoReflect.Descriptor instead.
func (*WebhookListReq) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{4}
}

func (x *WebhookListReq) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *WebhookListReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *WebhookListReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type WebhookListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor string           `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Data   []*WebhookRecord `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *WebhookListRes) Reset() {
	*x = WebhookListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookListRes) ProtoMessage() {}

func (x *WebhookListRes) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookListRes.ProtoReflect.Descriptor instead.
func (*WebhookListRes) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{5}
}

func (x *WebhookListRes) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *WebhookListRes) GetData() []*WebhookRecord {
	if x != nil {
		return x.Data
	}
	return nil
}

type WebhookDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *WebhookDeleteReq) Reset() {
	*x = WebhookDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookDeleteReq) ProtoMessage() {}

func (x *WebhookDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookDeleteReq.ProtoReflect.Descriptor instead.
func (*WebhookDeleteReq) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{6}
}

func (x *WebhookDeleteReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type WebhookDeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WebhookDeleteRes) Reset() {
	*x = WebhookDeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookDeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookDeleteRes) ProtoMessage() {}

func (x *WebhookDeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookDeleteRes.ProtoReflect.Descriptor instead.
func (*WebhookDeleteRes) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{7}
}

type WebhookAddTokensReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebhookId string `protobuf:"bytes,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	Count     int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *WebhookAddTokensReq) Reset() {
	*x = WebhookAddTokensReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookAddTokensReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookAddTokensReq) ProtoMessage() {}

func (x *WebhookAddTokensReq) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookAddTokensReq.ProtoReflect.Descriptor instead.
func (*WebhookAddTokensReq) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{8}
}

func (x *WebhookAddTokensReq) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

func (x *WebhookAddTokensReq) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type WebhookAddTokensRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []*WebhookTokenRecord `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *WebhookAddTokensRes) Reset() {
	*x = WebhookAddTokensRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookAddTokensRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookAddTokensRes) ProtoMessage() {}

func (x *WebhookAddTokensRes) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookAddTokensRes.ProtoReflect.Descriptor instead.
func (*WebhookAddTokensRes) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{9}
}

func (x *WebhookAddTokensRes) GetTokens() []*WebhookTokenRecord {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type WebhookDeleteTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WebhookId string   `protobuf:"bytes,1,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	Id        []string `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *WebhookDeleteTokenReq) Reset() {
	*x = WebhookDeleteTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookDeleteTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookDeleteTokenReq) ProtoMessage() {}

func (x *WebhookDeleteTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookDeleteTokenReq.ProtoReflect.Descriptor instead.
func (*WebhookDeleteTokenReq) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{10}
}

func (x *WebhookDeleteTokenReq) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

func (x *WebhookDeleteTokenReq) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

type WebhookDeleteTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WebhookDeleteTokenRes) Reset() {
	*x = WebhookDeleteTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookDeleteTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookDeleteTokenRes) ProtoMessage() {}

func (x *WebhookDeleteTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookDeleteTokenRes.ProtoReflect.Descriptor instead.
func (*WebhookDeleteTokenRes) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{11}
}

var File_webhook_proto protoreflect.FileDescriptor

var file_webhook_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x5c, 0x0a,
	0x0e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x64,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xde, 0x01, 0x0a, 0x0d,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x8c, 0x01, 0x0a,
	0x12, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x0e,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x22, 0x69, 0x0a, 0x0e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x63, 0x72,
	0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a,
	0x10, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x12, 0x0a, 0x10, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x13, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x41, 0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x5f, 0x0a, 0x13, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x41, 0x64, 0x64, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x22, 0x46, 0x0a, 0x15, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x32, 0x93, 0x05, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12,
	0x63, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x2c, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68,
	0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x61,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f,
	0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x73, 0x63,
	0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x2b, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x00, 0x12, 0x64, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2c, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x2c, 0x2e,
	0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x6a, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2e, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68,
	0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68,
	0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x09, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x31, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f,
	0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x41, 0x64, 0x64,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x73, 0x63, 0x72, 0x61,
	0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x41, 0x64, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x79,
	0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x33, 0x2e,
	0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x33, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_webhook_proto_rawDescOnce sync.Once
	file_webhook_proto_rawDescData = file_webhook_proto_rawDesc
)

func file_webhook_proto_rawDescGZIP() []byte {
	file_webhook_proto_rawDescOnce.Do(func() {
		file_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_webhook_proto_rawDescData)
	})
	return file_webhook_proto_rawDescData
}

var file_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_webhook_proto_goTypes = []interface{}{
	(*WebhookSaveReq)(nil),        // 0: scraphook.admin.dashboard.v1.WebhookSaveReq
	(*WebhookRecord)(nil),         // 1: scraphook.admin.dashboard.v1.WebhookRecord
	(*WebhookTokenRecord)(nil),    // 2: scraphook.admin.dashboard.v1.WebhookTokenRecord
	(*WebhookGetReq)(nil),         // 3: scraphook.admin.dashboard.v1.WebhookGetReq
	(*WebhookListReq)(nil),        // 4: scraphook.admin.dashboard.v1.WebhookListReq
	(*WebhookListRes)(nil),        // 5: scraphook.admin.dashboard.v1.WebhookListRes
	(*WebhookDeleteReq)(nil),      // 6: scraphook.admin.dashboard.v1.WebhookDeleteReq
	(*WebhookDeleteRes)(nil),      // 7: scraphook.admin.dashboard.v1.WebhookDeleteRes
	(*WebhookAddTokensReq)(nil),   // 8: scraphook.admin.dashboard.v1.WebhookAddTokensReq
	(*WebhookAddTokensRes)(nil),   // 9: scraphook.admin.dashboard.v1.WebhookAddTokensRes
	(*WebhookDeleteTokenReq)(nil), // 10: scraphook.admin.dashboard.v1.WebhookDeleteTokenReq
	(*WebhookDeleteTokenRes)(nil), // 11: scraphook.admin.dashboard.v1.WebhookDeleteTokenRes
}
var file_webhook_proto_depIdxs = []int32{
	2,  // 0: scraphook.admin.dashboard.v1.WebhookRecord.tokens:type_name -> scraphook.admin.dashboard.v1.WebhookTokenRecord
	1,  // 1: scraphook.admin.dashboard.v1.WebhookListRes.data:type_name -> scraphook.admin.dashboard.v1.WebhookRecord
	2,  // 2: scraphook.admin.dashboard.v1.WebhookAddTokensRes.tokens:type_name -> scraphook.admin.dashboard.v1.WebhookTokenRecord
	0,  // 3: scraphook.admin.dashboard.v1.Webhook.Save:input_type -> scraphook.admin.dashboard.v1.WebhookSaveReq
	3,  // 4: scraphook.admin.dashboard.v1.Webhook.Get:input_type -> scraphook.admin.dashboard.v1.WebhookGetReq
	4,  // 5: scraphook.admin.dashboard.v1.Webhook.List:input_type -> scraphook.admin.dashboard.v1.WebhookListReq
	6,  // 6: scraphook.admin.dashboard.v1.Webhook.Delete:input_type -> scraphook.admin.dashboard.v1.WebhookDeleteReq
	8,  // 7: scraphook.admin.dashboard.v1.Webhook.AddTokens:input_type -> scraphook.admin.dashboard.v1.WebhookAddTokensReq
	10, // 8: scraphook.admin.dashboard.v1.Webhook.DeleteToken:input_type -> scraphook.admin.dashboard.v1.WebhookDeleteTokenReq
	1,  // 9: scraphook.admin.dashboard.v1.Webhook.Save:output_type -> scraphook.admin.dashboard.v1.WebhookRecord
	1,  // 10: scraphook.admin.dashboard.v1.Webhook.Get:output_type -> scraphook.admin.dashboard.v1.WebhookRecord
	5,  // 11: scraphook.admin.dashboard.v1.Webhook.List:output_type -> scraphook.admin.dashboard.v1.WebhookListRes
	7,  // 12: scraphook.admin.dashboard.v1.Webhook.Delete:output_type -> scraphook.admin.dashboard.v1.WebhookDeleteRes
	9,  // 13: scraphook.admin.dashboard.v1.Webhook.AddTokens:output_type -> scraphook.admin.dashboard.v1.WebhookAddTokensRes
	11, // 14: scraphook.admin.dashboard.v1.Webhook.DeleteToken:output_type -> scraphook.admin.dashboard.v1.WebhookDeleteTokenRes
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_webhook_proto_init() }
func file_webhook_proto_init() {
	if File_webhook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_webhook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookSaveReq); i {
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
		file_webhook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookRecord); i {
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
		file_webhook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookTokenRecord); i {
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
		file_webhook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookGetReq); i {
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
		file_webhook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookListReq); i {
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
		file_webhook_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookListRes); i {
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
		file_webhook_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookDeleteReq); i {
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
		file_webhook_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookDeleteRes); i {
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
		file_webhook_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookAddTokensReq); i {
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
		file_webhook_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookAddTokensRes); i {
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
		file_webhook_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookDeleteTokenReq); i {
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
		file_webhook_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookDeleteTokenRes); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_webhook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_webhook_proto_goTypes,
		DependencyIndexes: file_webhook_proto_depIdxs,
		MessageInfos:      file_webhook_proto_msgTypes,
	}.Build()
	File_webhook_proto = out.File
	file_webhook_proto_rawDesc = nil
	file_webhook_proto_goTypes = nil
	file_webhook_proto_depIdxs = nil
}
