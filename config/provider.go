/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"github.com/saagie/provider-ovh/config/database"
	"github.com/saagie/provider-ovh/config/kube"
	"github.com/saagie/provider-ovh/config/user"

	ujconfig "github.com/upbound/upjet/pkg/config"

	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
)

const (
	resourcePrefix = "ovh"
	modulePath     = "github.com/saagie/provider-ovh"
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
		database.Configure,
		kube.Configure,
		user.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
