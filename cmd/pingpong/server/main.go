package main

import (
	"context"
	"path/filepath"

	"github.com/ohanan/card_proto/pkg/proto"
	"github.com/ohanan/cs/pkg/service"
)

func main() {
	println("start")
	abs, _ := filepath.Abs(".")
	println(abs)
	plugin, err := service.RunPlugin("./demo.exe", filepath.Join(abs, "./demo"))
	if err != nil {
		panic(err)
	}
	pong, err := plugin.PingPong(context.Background(), &proto.PingPongReq{Ping: "hello"})
	if err != nil {
		panic(err)
	}
	println(pong.Pong)
	pong, err = plugin.PingPong(context.Background(), &proto.PingPongReq{Ping: "world"})
	if err != nil {
		panic(err)
	}
	println(pong.Pong)
}
