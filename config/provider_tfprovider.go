//go:build !ci

package config

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
	cfprovider "github.com/prolixalias/terraform-provider-cloudflare/v5/provider"
)

// getTerraformProvider returns the actual Cloudflare terraform provider.
// This is excluded during CI builds to avoid compiling the massive
// terraform-provider-cloudflare dependency.
func getTerraformProvider() provider.Provider {
	return cfprovider.NewProvider("v5.16.0")()
}
