/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	database "github.com/saagie/provider-ovh/internal/controller/database/database"
	iprestrictions "github.com/saagie/provider-ovh/internal/controller/kube/iprestrictions"
	kube "github.com/saagie/provider-ovh/internal/controller/kube/kube"
	nodepool "github.com/saagie/provider-ovh/internal/controller/kube/nodepool"
	providerconfig "github.com/saagie/provider-ovh/internal/controller/providerconfig"
	s3credentials "github.com/saagie/provider-ovh/internal/controller/user/s3credentials"
	s3policy "github.com/saagie/provider-ovh/internal/controller/user/s3policy"
	user "github.com/saagie/provider-ovh/internal/controller/user/user"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		database.Setup,
		iprestrictions.Setup,
		kube.Setup,
		nodepool.Setup,
		providerconfig.Setup,
		s3credentials.Setup,
		s3policy.Setup,
		user.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
