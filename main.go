package main

import (
	"github.com/jaczkal/terraform-provider-linux/linux"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: linux.Provider,
	})
}
