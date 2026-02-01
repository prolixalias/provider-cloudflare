//go:build ci

package config

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"
)

// getTerraformProvider returns nil during CI builds.
// This avoids compiling the massive terraform-provider-cloudflare
// dependency which requires 10GB+ of disk space.
// The actual provider is only needed at runtime, not for linting or testing.
func getTerraformProvider() provider.Provider {
	return nil
}
