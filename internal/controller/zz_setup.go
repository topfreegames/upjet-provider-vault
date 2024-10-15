/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	authbackendrole "github.com/topfreegames/upjet-provider-vault/internal/controller/aws/authbackendrole"
	group "github.com/topfreegames/upjet-provider-vault/internal/controller/identitygroup/group"
	authbackendrolejwt "github.com/topfreegames/upjet-provider-vault/internal/controller/jwt/authbackendrole"
	authbackendrolekubernetesauthbackendrole "github.com/topfreegames/upjet-provider-vault/internal/controller/kubernetesauthbackendrole/authbackendrole"
	providerconfig "github.com/topfreegames/upjet-provider-vault/internal/controller/providerconfig"
	policy "github.com/topfreegames/upjet-provider-vault/internal/controller/vault/policy"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		authbackendrole.Setup,
		group.Setup,
		authbackendrolejwt.Setup,
		authbackendrolekubernetesauthbackendrole.Setup,
		providerconfig.Setup,
		policy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
