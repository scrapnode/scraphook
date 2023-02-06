// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: endpoint.proto

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

// EndpointClient is the client API for Endpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndpointClient interface {
	Save(ctx context.Context, in *EndpointSaveReq, opts ...grpc.CallOption) (*EndpointSaveRes, error)
	Get(ctx context.Context, in *EndpointGetReq, opts ...grpc.CallOption) (*EndpointGetRes, error)
	List(ctx context.Context, in *EndpointListReq, opts ...grpc.CallOption) (*EndpointListRes, error)
	Delete(ctx context.Context, in *EndpointDeleteReq, opts ...grpc.CallOption) (*EndpointDeleteRes, error)
}

type endpointClient struct {
	cc grpc.ClientConnInterface
}

func NewEndpointClient(cc grpc.ClientConnInterface) EndpointClient {
	return &endpointClient{cc}
}

func (c *endpointClient) Save(ctx context.Context, in *EndpointSaveReq, opts ...grpc.CallOption) (*EndpointSaveRes, error) {
	out := new(EndpointSaveRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Endpoint/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) Get(ctx context.Context, in *EndpointGetReq, opts ...grpc.CallOption) (*EndpointGetRes, error) {
	out := new(EndpointGetRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Endpoint/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) List(ctx context.Context, in *EndpointListReq, opts ...grpc.CallOption) (*EndpointListRes, error) {
	out := new(EndpointListRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Endpoint/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endpointClient) Delete(ctx context.Context, in *EndpointDeleteReq, opts ...grpc.CallOption) (*EndpointDeleteRes, error) {
	out := new(EndpointDeleteRes)
	err := c.cc.Invoke(ctx, "/scraphook.admin.dashboard.v1.Endpoint/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointServer is the server API for Endpoint service.
// All implementations must embed UnimplementedEndpointServer
// for forward compatibility
type EndpointServer interface {
	Save(context.Context, *EndpointSaveReq) (*EndpointSaveRes, error)
	Get(context.Context, *EndpointGetReq) (*EndpointGetRes, error)
	List(context.Context, *EndpointListReq) (*EndpointListRes, error)
	Delete(context.Context, *EndpointDeleteReq) (*EndpointDeleteRes, error)
	mustEmbedUnimplementedEndpointServer()
}

// UnimplementedEndpointServer must be embedded to have forward compatible implementations.
type UnimplementedEndpointServer struct {
}

func (UnimplementedEndpointServer) Save(context.Context, *EndpointSaveReq) (*EndpointSaveRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedEndpointServer) Get(context.Context, *EndpointGetReq) (*EndpointGetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEndpointServer) List(context.Context, *EndpointListReq) (*EndpointListRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedEndpointServer) Delete(context.Context, *EndpointDeleteReq) (*EndpointDeleteRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedEndpointServer) mustEmbedUnimplementedEndpointServer() {}

// UnsafeEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndpointServer will
// result in compilation errors.
type UnsafeEndpointServer interface {
	mustEmbedUnimplementedEndpointServer()
}

func RegisterEndpointServer(s grpc.ServiceRegistrar, srv EndpointServer) {
	s.RegisterService(&Endpoint_ServiceDesc, srv)
}

func _Endpoint_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Endpoint/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).Save(ctx, req.(*EndpointSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Endpoint/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).Get(ctx, req.(*EndpointGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Endpoint/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).List(ctx, req.(*EndpointListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endpoint_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scraphook.admin.dashboard.v1.Endpoint/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointServer).Delete(ctx, req.(*EndpointDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Endpoint_ServiceDesc is the grpc.ServiceDesc for Endpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Endpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scraphook.admin.dashboard.v1.Endpoint",
	HandlerType: (*EndpointServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Endpoint_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Endpoint_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Endpoint_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Endpoint_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "endpoint.proto",
}
