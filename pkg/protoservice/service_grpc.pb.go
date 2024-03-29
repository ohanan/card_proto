// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: service.proto

package protoservice

import (
	context "context"
	proto "github.com/ohanan/card_proto/pkg/protoservice/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Host_GetPlayerInfo_FullMethodName  = "/proto.Host/GetPlayerInfo"
	Host_RegisterNotify_FullMethodName = "/proto.Host/RegisterNotify"
	Host_AskAction_FullMethodName      = "/proto.Host/AskAction"
)

// HostClient is the client API for Host service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostClient interface {
	GetPlayerInfo(ctx context.Context, in *proto.GetPlayerInfo_Req, opts ...grpc.CallOption) (*proto.GetPlayerInfo_Resp, error)
	RegisterNotify(ctx context.Context, in *proto.RegisterNotify_Req, opts ...grpc.CallOption) (*proto.RegisterNotify_Resp, error)
	AskAction(ctx context.Context, in *proto.AskAction_Req, opts ...grpc.CallOption) (*proto.AskAction_Resp, error)
}

type hostClient struct {
	cc grpc.ClientConnInterface
}

func NewHostClient(cc grpc.ClientConnInterface) HostClient {
	return &hostClient{cc}
}

func (c *hostClient) GetPlayerInfo(ctx context.Context, in *proto.GetPlayerInfo_Req, opts ...grpc.CallOption) (*proto.GetPlayerInfo_Resp, error) {
	out := new(proto.GetPlayerInfo_Resp)
	err := c.cc.Invoke(ctx, Host_GetPlayerInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostClient) RegisterNotify(ctx context.Context, in *proto.RegisterNotify_Req, opts ...grpc.CallOption) (*proto.RegisterNotify_Resp, error) {
	out := new(proto.RegisterNotify_Resp)
	err := c.cc.Invoke(ctx, Host_RegisterNotify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostClient) AskAction(ctx context.Context, in *proto.AskAction_Req, opts ...grpc.CallOption) (*proto.AskAction_Resp, error) {
	out := new(proto.AskAction_Resp)
	err := c.cc.Invoke(ctx, Host_AskAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServer is the server API for Host service.
// All implementations should embed UnimplementedHostServer
// for forward compatibility
type HostServer interface {
	GetPlayerInfo(context.Context, *proto.GetPlayerInfo_Req) (*proto.GetPlayerInfo_Resp, error)
	RegisterNotify(context.Context, *proto.RegisterNotify_Req) (*proto.RegisterNotify_Resp, error)
	AskAction(context.Context, *proto.AskAction_Req) (*proto.AskAction_Resp, error)
}

// UnimplementedHostServer should be embedded to have forward compatible implementations.
type UnimplementedHostServer struct {
}

func (UnimplementedHostServer) GetPlayerInfo(context.Context, *proto.GetPlayerInfo_Req) (*proto.GetPlayerInfo_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerInfo not implemented")
}
func (UnimplementedHostServer) RegisterNotify(context.Context, *proto.RegisterNotify_Req) (*proto.RegisterNotify_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNotify not implemented")
}
func (UnimplementedHostServer) AskAction(context.Context, *proto.AskAction_Req) (*proto.AskAction_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskAction not implemented")
}

// UnsafeHostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostServer will
// result in compilation errors.
type UnsafeHostServer interface {
	mustEmbedUnimplementedHostServer()
}

func RegisterHostServer(s grpc.ServiceRegistrar, srv HostServer) {
	s.RegisterService(&Host_ServiceDesc, srv)
}

func _Host_GetPlayerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.GetPlayerInfo_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServer).GetPlayerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Host_GetPlayerInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServer).GetPlayerInfo(ctx, req.(*proto.GetPlayerInfo_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Host_RegisterNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.RegisterNotify_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServer).RegisterNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Host_RegisterNotify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServer).RegisterNotify(ctx, req.(*proto.RegisterNotify_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Host_AskAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.AskAction_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServer).AskAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Host_AskAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServer).AskAction(ctx, req.(*proto.AskAction_Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Host_ServiceDesc is the grpc.ServiceDesc for Host service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Host_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Host",
	HandlerType: (*HostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerInfo",
			Handler:    _Host_GetPlayerInfo_Handler,
		},
		{
			MethodName: "RegisterNotify",
			Handler:    _Host_RegisterNotify_Handler,
		},
		{
			MethodName: "AskAction",
			Handler:    _Host_AskAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

const (
	Plugin_GetPluginInfo_FullMethodName = "/proto.Plugin/GetPluginInfo"
	Plugin_StartGame_FullMethodName     = "/proto.Plugin/StartGame"
)

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	GetPluginInfo(ctx context.Context, in *proto.GetPluginInfo_Req, opts ...grpc.CallOption) (*proto.GetPluginInfo_Resp, error)
	StartGame(ctx context.Context, in *proto.StartGame_Req, opts ...grpc.CallOption) (*proto.StartGame_Resp, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) GetPluginInfo(ctx context.Context, in *proto.GetPluginInfo_Req, opts ...grpc.CallOption) (*proto.GetPluginInfo_Resp, error) {
	out := new(proto.GetPluginInfo_Resp)
	err := c.cc.Invoke(ctx, Plugin_GetPluginInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) StartGame(ctx context.Context, in *proto.StartGame_Req, opts ...grpc.CallOption) (*proto.StartGame_Resp, error) {
	out := new(proto.StartGame_Resp)
	err := c.cc.Invoke(ctx, Plugin_StartGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
// All implementations should embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	GetPluginInfo(context.Context, *proto.GetPluginInfo_Req) (*proto.GetPluginInfo_Resp, error)
	StartGame(context.Context, *proto.StartGame_Req) (*proto.StartGame_Resp, error)
}

// UnimplementedPluginServer should be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (UnimplementedPluginServer) GetPluginInfo(context.Context, *proto.GetPluginInfo_Req) (*proto.GetPluginInfo_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}
func (UnimplementedPluginServer) StartGame(context.Context, *proto.StartGame_Req) (*proto.StartGame_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}

// UnsafePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServer will
// result in compilation errors.
type UnsafePluginServer interface {
	mustEmbedUnimplementedPluginServer()
}

func RegisterPluginServer(s grpc.ServiceRegistrar, srv PluginServer) {
	s.RegisterService(&Plugin_ServiceDesc, srv)
}

func _Plugin_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.GetPluginInfo_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_GetPluginInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetPluginInfo(ctx, req.(*proto.GetPluginInfo_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.StartGame_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Plugin_StartGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).StartGame(ctx, req.(*proto.StartGame_Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Plugin_ServiceDesc is the grpc.ServiceDesc for Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPluginInfo",
			Handler:    _Plugin_GetPluginInfo_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _Plugin_StartGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
