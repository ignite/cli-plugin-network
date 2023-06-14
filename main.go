package main

import (
	"context"
	"os"

	hplugin "github.com/hashicorp/go-plugin"

	"github.com/ignite/cli/ignite/services/plugin"

	"github.com/ignite/cli-plugin-network/cmd"
)

type p struct {
	plugin.Interface
}

func (p) Manifest(_ context.Context) (*plugin.Manifest, error) {
	m := &plugin.Manifest{
		Name: "network",
	}
	m.ImportCobraCommand(cmd.NewNetwork(), "ignite")
	return m, nil
}

func (p) Execute(_ context.Context, c *plugin.ExecutedCommand) error {
	// Instead of a switch on c.Use, we run the root command like if
	// we were in a command line context. This implies to set os.Args
	// correctly.
	// Remove the first arg "ignite" from OsArgs because our network
	// command root is "network" not "ignite".
	os.Args = c.OsArgs[1:]
	return cmd.NewNetwork().Execute()
}

func main() {
	hplugin.Serve(&hplugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig(),
		Plugins: map[string]hplugin.Plugin{
			"cli-plugin-network": plugin.NewGRPC(&p{}),
		},
		GRPCServer: hplugin.DefaultGRPCServer,
	})
}
