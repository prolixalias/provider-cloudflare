package address

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_address_map", func(r *config.Resource) {
		// Use "addressmap" as short group to avoid Go reserved keyword "map"
		r.ShortGroup = "addressmap"
		r.Kind = "AddressMap"
	})
}
