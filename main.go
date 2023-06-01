package main

import (
	"github.com/TheNicholi/terraform-provider-json-formatter/json_formatter"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return json_formatter.Provider()
		},
	})
}
