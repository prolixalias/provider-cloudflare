//go:build ci

package clients

import fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"

func getFrameworkProvider() fwprovider.Provider {
	return nil
}
