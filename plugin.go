package main

import (
	"os"

	"log"

	"github.com/grafana/grafana_plugin_model/go/datasource"
	"github.com/hashicorp/go-plugin"
)

var pluginConfig = &plugin.ServeConfig{
	HandshakeConfig: plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "grafana_plugin_type",
		MagicCookieValue: "datasource",
	},
	Plugins: map[string]plugin.Plugin{
		"backend-datasource": &datasource.DatasourcePluginImpl{Plugin: &ConsulDatasource{}},
	},

	// A non-nil value here enables gRPC serving for this plugin...
	GRPCServer: plugin.DefaultGRPCServer,
}

func main() {
	log.SetOutput(os.Stderr) // the plugin sends logs to the host process on strErr

	plugin.Serve(pluginConfig)
}
