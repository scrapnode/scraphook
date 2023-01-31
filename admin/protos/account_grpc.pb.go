// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: account.proto

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

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountClient interface {
	Sign(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignRes, error)
	Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRes, error)
	Refresh(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*RefreshRes, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Sign(ctx context.Context, in *SignReq, opts ...grpc.CallOption) (*SignRes, error) {
	out := new(SignRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Account/Sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Verify(ctx context.Context, in *VerifyReq, opts ...grpc.CallOption) (*VerifyRes, error) {
	out := new(VerifyRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Account/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Refresh(ctx context.Context, in *RefreshReq, opts ...grpc.CallOption) (*RefreshRes, error) {
	out := new(RefreshRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Account/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
// All implementations must embed UnimplementedAccountServer
// for forward compatibility
type AccountServer interface {
	Sign(context.Context, *SignReq) (*SignRes, error)
	Verify(context.Context, *VerifyReq) (*VerifyRes, error)
	Refresh(context.Context, *RefreshReq) (*RefreshRes, error)
	mustEmbedUnimplementedAccountServer()
}

// UnimplementedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (UnimplementedAccountServer) Sign(context.Context, *SignReq) (*SignRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedAccountServer) Verify(context.Context, *VerifyReq) (*VerifyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedAccountServer) Refresh(context.Context, *RefreshReq) (*RefreshRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedAccountServer) mustEmbedUnimplementedAccountServer() {}

// UnsafeAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServer will
// result in compilation errors.
type UnsafeAccountServer interface {
	mustEmbedUnimplementedAccountServer()
}

func RegisterAccountServer(s grpc.ServiceRegistrar, srv AccountServer) {
	s.RegisterService(&Account_ServiceDesc, srv)
}

func _Account_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Account/Sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Sign(ctx, req.(*SignReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Account/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Verify(ctx, req.(*VerifyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Account/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Refresh(ctx, req.(*RefreshReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Account_ServiceDesc is the grpc.ServiceDesc for Account service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Account_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scraphook.admin.dashboard.v1.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sign",
			Handler:    _Account_Sign_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Account_Verify_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _Account_Refresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}