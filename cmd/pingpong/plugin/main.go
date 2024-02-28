package main

import (
	"context"

	"github.com/ohanan/card_proto/pkg/plugin"
	"github.com/ohanan/card_proto/pkg/proto"
)

type server struct {
}

func (s server) PingPong(ctx context.Context, req *proto.PingPongReq) (*proto.PingPongResp, error) {
	println("invoke ping pong")
	return &proto.PingPongResp{Pong: req.Ping}, nil
}

func (s server) Hello(ctx context.Context, req *proto.HelloReq) (*proto.HelloResp, error) {
	println("invoke hello")
	return &proto.HelloResp{Pong: req.Ping}, nil
}

func main() {
	plugin.ServePlugin("demo", &plugin.PluginService{
		Card:  &server{},
		Hello: &server{},
	})
}
