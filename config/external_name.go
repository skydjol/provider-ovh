/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/errors"
	"github.com/upbound/upjet/pkg/config"
)

const (
	// ErrFmtNoAttribute is an error string for not-found attributes
	ErrFmtNoAttribute = `"attribute not found: %s`
	// ErrFmtUnexpectedType is an error string for attribute map values of unexpected type
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

var kubeIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}

		return fmt.Sprintf("%s/%s", serviceName, externalName), nil
	},
	DisableNameInitializer: true,
}

var kubePoolIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}

		kubeID, ok := parameters["kube_id"]
		if !ok {
			return "", errors.Errorf(ErrFmtNoAttribute, "kube_id")
		}
		kubeIDStr, ok := kubeID.(string)
		if !ok {
			return "", errors.Errorf(ErrFmtUnexpectedType, "kube_id")
		}

		return fmt.Sprintf("%s/%s/%s", serviceName, kubeIDStr, externalName), nil
	},
	DisableNameInitializer: true,
}

var databaseIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}

		engine, ok := parameters["engine"]
		if !ok {
			return "", errors.Errorf(ErrFmtNoAttribute, "engine")
		}
		engineStr, ok := engine.(string)
		if !ok {
			return "", errors.Errorf(ErrFmtUnexpectedType, "engine")
		}

		return fmt.Sprintf("%s/%s/%s", serviceName, engineStr, externalName), nil
	},
	DisableNameInitializer: true,
}

var s3CredentialsIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}

		userID, ok := parameters["user_id"]
		if !ok {
			return "", errors.Errorf(ErrFmtNoAttribute, "user_id")
		}
		userIDStr, ok := userID.(string)
		if !ok {
			return "", errors.Errorf(ErrFmtUnexpectedType, "user_id")
		}

		return fmt.Sprintf("%s/%s/%s", serviceName, userIDStr, externalName), nil
	},
	DisableNameInitializer: true,
}

var userIdentifierFromProvider = config.ExternalName{
	SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
	GetExternalNameFn:       config.IDAsExternalName,
	GetIDFn: func(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {

		// Hack when externalName is null on creation i don't know why
		if externalName == "" {
			return externalName, nil
		}

		serviceName, err := serviceName(parameters)
		if err != nil {
			return serviceName, err
		}

		return fmt.Sprintf("%s/%s", serviceName, externalName), nil
	},
	DisableNameInitializer: true,
}

func serviceName(parameters map[string]any) (string, error) {
	serviceName, ok := parameters["service_name"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "service_name")
	}
	serviceNameStr, ok := serviceName.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "service_name")
	}
	return serviceNameStr, nil
}

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"ovh_cloud_project_database":            databaseIdentifierFromProvider,
	"ovh_cloud_project_kube":                kubeIdentifierFromProvider,
	"ovh_cloud_project_kube_nodepool":       kubePoolIdentifierFromProvider,
	"ovh_cloud_project_kube_iprestrictions": kubeIdentifierFromProvider,
	"ovh_cloud_project_user":                userIdentifierFromProvider,
	"ovh_cloud_project_user_s3_policy":      config.IdentifierFromProvider,
	"ovh_cloud_project_user_s3_credential":  s3CredentialsIdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
