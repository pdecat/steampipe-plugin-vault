package main

import (
	"github.com/theapsgroup/steampipe-plugin-vault/vault"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: vault.Plugin})
}
