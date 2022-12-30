package kubernetesauthbackendrole

import (
	"context"
	"fmt"
	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
)

var (
	ErrFmtNoAttribute    = ""
	ErrFmtUnexpectedType = ""
)

func getFullyQualifiedIDfunc(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
	backend, ok := parameters["backend"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "backend")
	}
	backendStr, ok := backend.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "backend")
	}

	return fmt.Sprintf("auth/%s/role/%s", backendStr, externalName), nil
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vault_kubernetes_auth_backend_role", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "kubernetesAuthBackendRole"
		r.ExternalName = config.NameAsIdentifier
		r.ExternalName.SetIdentifierArgumentFn = func(base map[string]any, externalName string) {
			base["role_name"] = externalName
		}
		r.ExternalName.OmittedFields = []string{
			"role_name",
		}
		r.ExternalName.GetIDFn = getFullyQualifiedIDfunc
		r.ExternalName.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
			if name, ok := tfstate["role_name"].(string); ok && name != "" {
				return name, nil
			}
			return "", errors.New("cannot find role_name in tfstate")
		}
	})
}
