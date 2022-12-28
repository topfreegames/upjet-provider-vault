/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	identitygroup "github.com/topfreegames/upjet-provider-vault/config/identity-group"
	kubernetesauthbackendrole "github.com/topfreegames/upjet-provider-vault/config/kubernetes-auth-backend-role"
	ujconfig "github.com/upbound/upjet/pkg/config"
)

const (
	resourcePrefix = "vault"
	modulePath     = "github.com/topfreegames/upjet-provider-vault"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		kubernetesauthbackendrole.Configure,
		identitygroup.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
