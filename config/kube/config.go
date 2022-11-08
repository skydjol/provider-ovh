package kube

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure kube resources
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_kube", func(r *config.Resource) {
		r.UseAsync = true

		r.ShortGroup = "kube.cloud"
		r.Kind = "Kube"
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_nodepool", func(r *config.Resource) {
		r.References["kube_id"] = config.Reference{
			Type: "github.com/saagie/provider-ovh/apis/kube/v1alpha1.Kube",
		}
		r.UseAsync = true

		r.ShortGroup = "kube.cloud"
		r.Kind = "NodePool"
	})
	p.AddResourceConfigurator("ovh_cloud_project_kube_iprestrictions", func(r *config.Resource) {
		r.References["kube_id"] = config.Reference{
			Type: "github.com/saagie/provider-ovh/apis/kube/v1alpha1.Kube",
		}

		r.ShortGroup = "kube.cloud"
		r.Kind = "IpRestrictions"
	})
}
