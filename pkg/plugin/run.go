package plugin

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/ohanan/card_proto/pkg/proto"
	"google.golang.org/grpc"
)

func ServePlugin(name string, server proto.PluginXServer) {
	xServer := proto.NewPluginServerFromXServer(server)
	p := &xPlugin{
		xServer: xServer,
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: CreateHandshakeConfig(),
		Plugins:         map[string]plugin.Plugin{name: p},
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}

type xPlugin struct {
	plugin.Plugin
	xServer *proto.PluginXServerAdapter
}

func (p *xPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	proto.RegisterPluginServer(server, p.xServer)
	return nil
}

func (p *xPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	client := proto.NewCardClient(conn)
	p.xServer.Client = *proto.NewCardXClient(client)
	return client, nil
}
func CreateHandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "CARD",
		MagicCookieValue: "card",
	}
}
