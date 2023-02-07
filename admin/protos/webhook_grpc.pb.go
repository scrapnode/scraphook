// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: webhook.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WebhookClient is the client API for Webhook service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebhookClient interface {
	Save(ctx context.Context, in *WebhookSaveReq, opts ...grpc.CallOption) (*WebhookRecord, error)
	Get(ctx context.Context, in *WebhookGetReq, opts ...grpc.CallOption) (*WebhookRecord, error)
	List(ctx context.Context, in *WebhookListReq, opts ...grpc.CallOption) (*WebhookListRes, error)
	Delete(ctx context.Context, in *WebhookDeleteReq, opts ...grpc.CallOption) (*WebhookDeleteRes, error)
	AddTokens(ctx context.Context, in *WebhookAddTokensReq, opts ...grpc.CallOption) (*WebhookAddTokensRes, error)
	DeleteToken(ctx context.Context, in *WebhookDeleteTokenReq, opts ...grpc.CallOption) (*WebhookDeleteTokenRes, error)
}

type webhookClient struct {
	cc grpc.ClientConnInterface
}

func NewWebhookClient(cc grpc.ClientConnInterface) WebhookClient {
	return &webhookClient{cc}
}

func (c *webhookClient) Save(ctx context.Context, in *WebhookSaveReq, opts ...grpc.CallOption) (*WebhookRecord, error) {
	out := new(WebhookRecord)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Webhook/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) Get(ctx context.Context, in *WebhookGetReq, opts ...grpc.CallOption) (*WebhookRecord, error) {
	out := new(WebhookRecord)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Webhook/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) List(ctx context.Context, in *WebhookListReq, opts ...grpc.CallOption) (*WebhookListRes, error) {
	out := new(WebhookListRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Webhook/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) Delete(ctx context.Context, in *WebhookDeleteReq, opts ...grpc.CallOption) (*WebhookDeleteRes, error) {
	out := new(WebhookDeleteRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Webhook/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) AddTokens(ctx context.Context, in *WebhookAddTokensReq, opts ...grpc.CallOption) (*WebhookAddTokensRes, error) {
	out := new(WebhookAddTokensRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Webhook/AddTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webhookClient) DeleteToken(ctx context.Context, in *WebhookDeleteTokenReq, opts ...grpc.CallOption) (*WebhookDeleteTokenRes, error) {
	out := new(WebhookDeleteTokenRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Webhook/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebhookServer is the server API for Webhook service.
// All implementations must embed UnimplementedWebhookServer
// for forward compatibility
type WebhookServer interface {
	Save(context.Context, *WebhookSaveReq) (*WebhookRecord, error)
	Get(context.Context, *WebhookGetReq) (*WebhookRecord, error)
	List(context.Context, *WebhookListReq) (*WebhookListRes, error)
	Delete(context.Context, *WebhookDeleteReq) (*WebhookDeleteRes, error)
	AddTokens(context.Context, *WebhookAddTokensReq) (*WebhookAddTokensRes, error)
	DeleteToken(context.Context, *WebhookDeleteTokenReq) (*WebhookDeleteTokenRes, error)
	mustEmbedUnimplementedWebhookServer()
}

// UnimplementedWebhookServer must be embedded to have forward compatible implementations.
type UnimplementedWebhookServer struct {
}

func (UnimplementedWebhookServer) Save(context.Context, *WebhookSaveReq) (*WebhookRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedWebhookServer) Get(context.Context, *WebhookGetReq) (*WebhookRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedWebhookServer) List(context.Context, *WebhookListReq) (*WebhookListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedWebhookServer) Delete(context.Context, *WebhookDeleteReq) (*WebhookDeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedWebhookServer) AddTokens(context.Context, *WebhookAddTokensReq) (*WebhookAddTokensRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTokens not implemented")
}
func (UnimplementedWebhookServer) DeleteToken(context.Context, *WebhookDeleteTokenReq) (*WebhookDeleteTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (UnimplementedWebhookServer) mustEmbedUnimplementedWebhookServer() {}

// UnsafeWebhookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebhookServer will
// result in compilation errors.
type UnsafeWebhookServer interface {
	mustEmbedUnimplementedWebhookServer()
}

func RegisterWebhookServer(s grpc.ServiceRegistrar, srv WebhookServer) {
	s.RegisterService(&Webhook_ServiceDesc, srv)
}

func _Webhook_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Webhook/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Save(ctx, req.(*WebhookSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Webhook/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Get(ctx, req.(*WebhookGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Webhook/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).List(ctx, req.(*WebhookListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Webhook/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).Delete(ctx, req.(*WebhookDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_AddTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookAddTokensReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).AddTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Webhook/AddTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).AddTokens(ctx, req.(*WebhookAddTokensReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Webhook_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookDeleteTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Webhook/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookServer).DeleteToken(ctx, req.(*WebhookDeleteTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Webhook_ServiceDesc is the grpc.ServiceDesc for Webhook service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Webhook_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scraphook.admin.dashboard.v1.Webhook",
	HandlerType: (*WebhookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Webhook_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Webhook_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Webhook_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Webhook_Delete_Handler,
		},
		{
			MethodName: "AddTokens",
			Handler:    _Webhook_AddTokens_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Webhook_DeleteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "webhook.proto",
}
