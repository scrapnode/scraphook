// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: endpoint_rule.proto

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

type EndpointRuleSaveReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId string `protobuf:"bytes,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Rule       string `protobuf:"bytes,3,opt,name=rule,proto3" json:"rule,omitempty"`
	Negative   bool   `protobuf:"varint,4,opt,name=negative,proto3" json:"negative,omitempty"`
	Priority   int32  `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *EndpointRuleSaveReq) Reset() {
	*x = EndpointRuleSaveReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRuleSaveReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRuleSaveReq) ProtoMessage() {}

func (x *EndpointRuleSaveReq) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRuleSaveReq.ProtoReflect.Descriptor instead.
func (*EndpointRuleSaveReq) Descriptor() ([]byte, []int) {
	return file_endpoint_rule_proto_rawDescGZIP(), []int{0}
}

func (x *EndpointRuleSaveReq) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

func (x *EndpointRuleSaveReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EndpointRuleSaveReq) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *EndpointRuleSaveReq) GetNegative() bool {
	if x != nil {
		return x.Negative
	}
	return false
}

func (x *EndpointRuleSaveReq) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type EndpointRuleGetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId string `protobuf:"bytes,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EndpointRuleGetReq) Reset() {
	*x = EndpointRuleGetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRuleGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRuleGetReq) ProtoMessage() {}

func (x *EndpointRuleGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRuleGetReq.ProtoReflect.Descriptor instead.
func (*EndpointRuleGetReq) Descriptor() ([]byte, []int) {
	return file_endpoint_rule_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointRuleGetReq) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

func (x *EndpointRuleGetReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EndpointRuleListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId string `protobuf:"bytes,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Cursor     string `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Size       int32  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Search     string `protobuf:"bytes,4,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *EndpointRuleListReq) Reset() {
	*x = EndpointRuleListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRuleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRuleListReq) ProtoMessage() {}

func (x *EndpointRuleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRuleListReq.ProtoReflect.Descriptor instead.
func (*EndpointRuleListReq) Descriptor() ([]byte, []int) {
	return file_endpoint_rule_proto_rawDescGZIP(), []int{2}
}

func (x *EndpointRuleListReq) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

func (x *EndpointRuleListReq) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *EndpointRuleListReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *EndpointRuleListReq) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type EndpointRuleListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor string                `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Data   []*EndpointRuleRecord `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *EndpointRuleListRes) Reset() {
	*x = EndpointRuleListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_rule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRuleListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRuleListRes) ProtoMessage() {}

func (x *EndpointRuleListRes) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_rule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRuleListRes.ProtoReflect.Descriptor instead.
func (*EndpointRuleListRes) Descriptor() ([]byte, []int) {
	return file_endpoint_rule_proto_rawDescGZIP(), []int{3}
}

func (x *EndpointRuleListRes) GetCursor() string {
	if x != nil {
		return x.Cursor
	}
	return ""
}

func (x *EndpointRuleListRes) GetData() []*EndpointRuleRecord {
	if x != nil {
		return x.Data
	}
	return nil
}

type EndpointRuleDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointId string `protobuf:"bytes,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EndpointRuleDeleteReq) Reset() {
	*x = EndpointRuleDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_rule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRuleDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRuleDeleteReq) ProtoMessage() {}

func (x *EndpointRuleDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_rule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRuleDeleteReq.ProtoReflect.Descriptor instead.
func (*EndpointRuleDeleteReq) Descriptor() ([]byte, []int) {
	return file_endpoint_rule_proto_rawDescGZIP(), []int{4}
}

func (x *EndpointRuleDeleteReq) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

func (x *EndpointRuleDeleteReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EndpointRuleDeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EndpointRuleDeleteRes) Reset() {
	*x = EndpointRuleDeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_endpoint_rule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRuleDeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRuleDeleteRes) ProtoMessage() {}

func (x *EndpointRuleDeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_endpoint_rule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRuleDeleteRes.ProtoReflect.Descriptor instead.
func (*EndpointRuleDeleteRes) Descriptor() ([]byte, []int) {
	return file_endpoint_rule_proto_rawDescGZIP(), []int{5}
}

var File_endpoint_rule_proto protoreflect.FileDescriptor

var file_endpoint_rule_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x1a, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x13, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x45, 0x0a, 0x12, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x7a, 0x0a, 0x13, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x73, 0x0a, 0x13, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x44, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x48, 0x0a, 0x15, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x32, 0xd0, 0x03, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x12, 0x6d, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x31, 0x2e, 0x73,
	0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x30, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x30, 0x2e, 0x73, 0x63, 0x72,
	0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x30, 0x2e, 0x73,
	0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x00,
	0x12, 0x6e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x31, 0x2e, 0x73, 0x63,
	0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x74, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x33, 0x2e, 0x73, 0x63, 0x72,
	0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x33, 0x2e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x73,
	0x63, 0x72, 0x61, 0x70, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_endpoint_rule_proto_rawDescOnce sync.Once
	file_endpoint_rule_proto_rawDescData = file_endpoint_rule_proto_rawDesc
)

func file_endpoint_rule_proto_rawDescGZIP() []byte {
	file_endpoint_rule_proto_rawDescOnce.Do(func() {
		file_endpoint_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_endpoint_rule_proto_rawDescData)
	})
	return file_endpoint_rule_proto_rawDescData
}

var file_endpoint_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_endpoint_rule_proto_goTypes = []interface{}{
	(*EndpointRuleSaveReq)(nil),   // 0: scraphook.admin.dashboard.v1.EndpointRuleSaveReq
	(*EndpointRuleGetReq)(nil),    // 1: scraphook.admin.dashboard.v1.EndpointRuleGetReq
	(*EndpointRuleListReq)(nil),   // 2: scraphook.admin.dashboard.v1.EndpointRuleListReq
	(*EndpointRuleListRes)(nil),   // 3: scraphook.admin.dashboard.v1.EndpointRuleListRes
	(*EndpointRuleDeleteReq)(nil), // 4: scraphook.admin.dashboard.v1.EndpointRuleDeleteReq
	(*EndpointRuleDeleteRes)(nil), // 5: scraphook.admin.dashboard.v1.EndpointRuleDeleteRes
	(*EndpointRuleRecord)(nil),    // 6: scraphook.admin.dashboard.v1.EndpointRuleRecord
}
var file_endpoint_rule_proto_depIdxs = []int32{
	6, // 0: scraphook.admin.dashboard.v1.EndpointRuleListRes.data:type_name -> scraphook.admin.dashboard.v1.EndpointRuleRecord
	0, // 1: scraphook.admin.dashboard.v1.EndpointRule.Save:input_type -> scraphook.admin.dashboard.v1.EndpointRuleSaveReq
	1, // 2: scraphook.admin.dashboard.v1.EndpointRule.Get:input_type -> scraphook.admin.dashboard.v1.EndpointRuleGetReq
	2, // 3: scraphook.admin.dashboard.v1.EndpointRule.List:input_type -> scraphook.admin.dashboard.v1.EndpointRuleListReq
	4, // 4: scraphook.admin.dashboard.v1.EndpointRule.Delete:input_type -> scraphook.admin.dashboard.v1.EndpointRuleDeleteReq
	6, // 5: scraphook.admin.dashboard.v1.EndpointRule.Save:output_type -> scraphook.admin.dashboard.v1.EndpointRuleRecord
	6, // 6: scraphook.admin.dashboard.v1.EndpointRule.Get:output_type -> scraphook.admin.dashboard.v1.EndpointRuleRecord
	3, // 7: scraphook.admin.dashboard.v1.EndpointRule.List:output_type -> scraphook.admin.dashboard.v1.EndpointRuleListRes
	5, // 8: scraphook.admin.dashboard.v1.EndpointRule.Delete:output_type -> scraphook.admin.dashboard.v1.EndpointRuleDeleteRes
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_endpoint_rule_proto_init() }
func file_endpoint_rule_proto_init() {
	if File_endpoint_rule_proto != nil {
		return
	}
	file_entities_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_endpoint_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRuleSaveReq); i {
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
		file_endpoint_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRuleGetReq); i {
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
		file_endpoint_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRuleListReq); i {
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
		file_endpoint_rule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRuleListRes); i {
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
		file_endpoint_rule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRuleDeleteReq); i {
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
		file_endpoint_rule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRuleDeleteRes); i {
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
			RawDescriptor: file_endpoint_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_endpoint_rule_proto_goTypes,
		DependencyIndexes: file_endpoint_rule_proto_depIdxs,
		MessageInfos:      file_endpoint_rule_proto_msgTypes,
	}.Build()
	File_endpoint_rule_proto = out.File
	file_endpoint_rule_proto_rawDesc = nil
	file_endpoint_rule_proto_goTypes = nil
	file_endpoint_rule_proto_depIdxs = nil
}
