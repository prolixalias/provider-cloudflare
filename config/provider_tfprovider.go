//go:build !ci

package config

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	cfprovider "github.com/prolixalias/terraform-provider-cloudflare/v5/provider"
)

// getTerraformProvider returns the actual Cloudflare terraform provider.
// This uses the forked provider which exports the provider via provider/provider.go
// wrapper around internal.NewProvider() to work around Go's internal package restrictions.
func getTerraformProvider() provider.Provider {
	return cfprovider.NewProvider("v5.16.0")()
}
