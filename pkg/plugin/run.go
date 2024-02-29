package plugin

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/ohanan/card_proto/pkg/protoservice"
	"google.golang.org/grpc"
)

func ServePlugin(name string, server protoservice.PluginXServer) {
	xServer := protoservice.NewPluginServerFromXServer(server)
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
	xServer *protoservice.PluginXServerAdapter
}

func (p *xPlugin) GRPCServer(broker *plugin.GRPCBroker, server *grpc.Server) error {
	protoservice.RegisterPluginServer(server, p.xServer)
	return nil
}

func (p *xPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, conn *grpc.ClientConn) (interface{}, error) {
	client := protoservice.NewCardClient(conn)
	p.xServer.Client = *protoservice.NewCardXClient(client)
	return client, nil
}
func CreateHandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "CARD",
		MagicCookieValue: "card",
	}
}
