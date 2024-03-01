package protoservice

import (
	"math/rand/v2"
)

func TyrInitHelper(hostServer HostServer, pluginServer PluginServer) {
	if x, ok := pluginServer.(*pluginXServerAdapter); ok {
		x.InitHelper(NewHostXClient(NewHostClientFromServer(hostServer)))
	}
}
func (x *pluginXServerAdapter) InitHelper(c *HostXClient) {
	x.Helper = &Helper{HostXClient: c}
}

type Helper struct {
	*HostXClient
}

func (h *Helper) NewRand(seed uint64) *Rand { return rand.New(uint64Value(seed)) }

type Rand = rand.Rand
type uint64Value uint64

func (u uint64Value) Uint64() uint64 { return uint64(u) }
