package identitygroup

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_identity_group", func(r *config.Resource) {
		r.ShortGroup = "identityGroup"
		r.ExternalName = config.IdentifierFromProvider
		config.MarkAsRequired(r.TerraformResource, "type")
	})
}
