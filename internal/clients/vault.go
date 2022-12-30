/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"strings"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"github.com/topfreegames/upjet-provider-vault/apis/v1beta1"
	"github.com/upbound/upjet/pkg/terraform"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// error messages
	errNoProviderConfig   = "no providerConfigRef provided"
	errGetProviderConfig  = "cannot get referenced ProviderConfig"
	errTrackUsage         = "cannot track ProviderConfig usage"
	errExtractCredentials = "cannot extract credentials"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		token, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		ps.Configuration = map[string]any{}
		ps.Configuration["address"] = pc.Spec.Address
		ps.Configuration["skip_tls_verify"] = pc.Spec.SkipTLSVerify
		ps.Configuration["token"] = strings.TrimSuffix(string(token), "\n")

		return ps, nil
	}
}
