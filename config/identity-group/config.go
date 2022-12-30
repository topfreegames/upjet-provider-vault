package identitygroup

import (
	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_identity_group", func(r *config.Resource) {
		r.ShortGroup = "identityGroup"
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if name, ok := tfstate["name"].(string); ok && name != "" {
				return name, nil
			}
			return "", errors.New("cannot find name in tfstate")
		}
		config.MarkAsRequired(r.TerraformResource, "type")
	})
}
