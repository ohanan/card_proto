// Code generated by protoc-gen-card-proto-plugin. DO NOT EDIT.

package protoservice

import (
	context "context"
	proto "github.com/ohanan/card_proto/pkg/protoservice/proto"
	grpc "google.golang.org/grpc"
)

type CardXClient struct {
	c CardClient
}

func NewCardXClient(c CardClient) *CardXClient {
	return &CardXClient{c: c}
}
func (x CardXClient) Hello(abcd string) *proto.Hello_Resp {
	r, err := x.c.Hello(context.Background(), &proto.Hello_Req{
		Abcd: abcd,
	})
	if err != nil {
		panic(err)
	}
	return r
}
func (x CardXClient) Hello0(req *proto.Hello_Req) *proto.Hello_Resp {
	r, err := x.c.Hello(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return r
}
func (x CardXClient) GetPlayerInfo(userId string) *proto.GetPlayerInfo_Resp {
	r, err := x.c.GetPlayerInfo(context.Background(), &proto.GetPlayerInfo_Req{
		UserId: userId,
	})
	if err != nil {
		panic(err)
	}
	return r
}
func (x CardXClient) GetPlayerInfo0(req *proto.GetPlayerInfo_Req) *proto.GetPlayerInfo_Resp {
	r, err := x.c.GetPlayerInfo(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return r
}
func (x CardXClient) RegisterNotify(event []string) *proto.RegisterNotify_Resp {
	r, err := x.c.RegisterNotify(context.Background(), &proto.RegisterNotify_Req{
		Event: event,
	})
	if err != nil {
		panic(err)
	}
	return r
}
func (x CardXClient) RegisterNotify0(req *proto.RegisterNotify_Req) *proto.RegisterNotify_Resp {
	r, err := x.c.RegisterNotify(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return r
}

type CardXServer interface {
	Hello(req *proto.Hello_Req, resp *proto.Hello_Resp)
	GetPlayerInfo(req *proto.GetPlayerInfo_Req, resp *proto.GetPlayerInfo_Resp)
	RegisterNotify(req *proto.RegisterNotify_Req, resp *proto.RegisterNotify_Resp)
}

func NewCardClientFromServer(x CardServer) CardClient {
	return &cardClientByServer{s: x}
}

type cardClientByServer struct{ s CardServer }

func (x cardClientByServer) Hello(ctx context.Context, req *proto.Hello_Req, opts ...grpc.CallOption) (*proto.Hello_Resp, error) {
	return x.s.Hello(ctx, req)
}
func (x cardClientByServer) GetPlayerInfo(ctx context.Context, req *proto.GetPlayerInfo_Req, opts ...grpc.CallOption) (*proto.GetPlayerInfo_Resp, error) {
	return x.s.GetPlayerInfo(ctx, req)
}
func (x cardClientByServer) RegisterNotify(ctx context.Context, req *proto.RegisterNotify_Req, opts ...grpc.CallOption) (*proto.RegisterNotify_Resp, error) {
	return x.s.RegisterNotify(ctx, req)
}
func NewCardServerFromXServer(server CardXServer) *CardXServerAdapter {
	return &CardXServerAdapter{Server: server}
}

type CardXServerAdapter struct {
	Server CardXServer
}

func (x CardXServerAdapter) Hello(ctx context.Context, req *proto.Hello_Req) (*proto.Hello_Resp, error) {
	resp := &proto.Hello_Resp{}
	x.Server.Hello(req, resp)
	return resp, nil
}
func (x CardXServerAdapter) GetPlayerInfo(ctx context.Context, req *proto.GetPlayerInfo_Req) (*proto.GetPlayerInfo_Resp, error) {
	resp := &proto.GetPlayerInfo_Resp{}
	x.Server.GetPlayerInfo(req, resp)
	return resp, nil
}
func (x CardXServerAdapter) RegisterNotify(ctx context.Context, req *proto.RegisterNotify_Req) (*proto.RegisterNotify_Resp, error) {
	resp := &proto.RegisterNotify_Resp{}
	x.Server.RegisterNotify(req, resp)
	return resp, nil
}

type PluginXClient struct {
	c PluginClient
}

func NewPluginXClient(c PluginClient) *PluginXClient {
	return &PluginXClient{c: c}
}
func (x PluginXClient) GetPluginInfo() *proto.GetPluginInfo_Resp {
	r, err := x.c.GetPluginInfo(context.Background(), &proto.GetPluginInfo_Req{})
	if err != nil {
		panic(err)
	}
	return r
}
func (x PluginXClient) GetPluginInfo0(req *proto.GetPluginInfo_Req) *proto.GetPluginInfo_Resp {
	r, err := x.c.GetPluginInfo(context.Background(), req)
	if err != nil {
		panic(err)
	}
	return r
}

type PluginXServer interface {
	GetPluginInfo(remote CardXClient, req *proto.GetPluginInfo_Req, resp *proto.GetPluginInfo_Resp)
}

func NewPluginClientFromServer(x PluginServer) PluginClient {
	return &pluginClientByServer{s: x}
}

type pluginClientByServer struct{ s PluginServer }

func (x pluginClientByServer) GetPluginInfo(ctx context.Context, req *proto.GetPluginInfo_Req, opts ...grpc.CallOption) (*proto.GetPluginInfo_Resp, error) {
	return x.s.GetPluginInfo(ctx, req)
}
func NewPluginServerFromXServer(server PluginXServer) *PluginXServerAdapter {
	return &PluginXServerAdapter{Server: server}
}

type PluginXServerAdapter struct {
	Server PluginXServer
	Client CardXClient
}

func (x PluginXServerAdapter) GetPluginInfo(ctx context.Context, req *proto.GetPluginInfo_Req) (*proto.GetPluginInfo_Resp, error) {
	resp := &proto.GetPluginInfo_Resp{}
	x.Server.GetPluginInfo(x.Client, req, resp)
	return resp, nil
}