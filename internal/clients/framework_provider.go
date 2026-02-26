//go:build !ci

package clients

import (
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	cfprovider "github.com/prolixalias/terraform-provider-cloudflare/v5/provider"
)

func getFrameworkProvider() fwprovider.Provider {
	return cfprovider.NewProvider("v5.16.0")()
}
