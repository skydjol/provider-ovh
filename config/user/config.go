package user

import (
	"github.com/upbound/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ovh_cloud_project_user", func(r *config.Resource) {
		r.ShortGroup = "user.cloud"
		r.Kind = "User"
	})
	p.AddResourceConfigurator("ovh_cloud_project_user_s3_policy", func(r *config.Resource) {
		r.References["user_id"] = config.Reference{
			Type: "github.com/saagie/provider-ovh/apis/user/v1alpha1.User",
		}
		r.ShortGroup = "user.cloud"
		r.Kind = "S3Policy"
	})
	p.AddResourceConfigurator("ovh_cloud_project_user_s3_credential", func(r *config.Resource) {
		r.References["user_id"] = config.Reference{
			Type: "github.com/saagie/provider-ovh/apis/user/v1alpha1.User",
		}
		r.ShortGroup = "user.cloud"
		r.Kind = "S3Credentials"
	})
}
