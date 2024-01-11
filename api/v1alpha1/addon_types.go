// Copyright 2024 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"

	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/clustertopology/variables"
)

type Addons struct {
	// +optional
	CNI *CNI `json:"cni,omitempty"`

	// +optional
	NFD *NFD `json:"nfd,omitempty"`

	// +optional
	CPI *CPI `json:"cpi,omitempty"`

	// +optional
	CSIProviders *CSIProviders `json:"csi,omitempty"`
}

func (Addons) VariableSchema() clusterv1.VariableSchema {
	return clusterv1.VariableSchema{
		OpenAPIV3Schema: clusterv1.JSONSchemaProps{
			Description: "Cluster configuration",
			Type:        "object",
			Properties: map[string]clusterv1.JSONSchemaProps{
				"cni": CNI{}.VariableSchema().OpenAPIV3Schema,
				"nfd": NFD{}.VariableSchema().OpenAPIV3Schema,
				"csi": CSIProviders{}.VariableSchema().OpenAPIV3Schema,
				"cpi": CPI{}.VariableSchema().OpenAPIV3Schema,
			},
		},
	}
}

type AddonStrategy string

const (
	AddonStrategyClusterResourceSet AddonStrategy = "ClusterResourceSet"
)

// CNI required for providing CNI configuration.
type CNI struct {
	// +optional
	Provider string `json:"provider,omitempty"`
	// +optional
	Strategy AddonStrategy `json:"strategy,omitempty"`
}

func (CNI) VariableSchema() clusterv1.VariableSchema {
	supportedCNIProviders := []string{CNIProviderCalico}

	return clusterv1.VariableSchema{
		OpenAPIV3Schema: clusterv1.JSONSchemaProps{
			Type: "object",
			Properties: map[string]clusterv1.JSONSchemaProps{
				"provider": {
					Description: "CNI provider to deploy",
					Type:        "string",
					Enum:        variables.MustMarshalValuesToEnumJSON(supportedCNIProviders...),
				},
				"strategy": {
					Description: "Addon strategy used to deploy the CNI provider to the workload cluster",
					Type:        "string",
					Enum: variables.MustMarshalValuesToEnumJSON(
						AddonStrategyClusterResourceSet,
					),
				},
			},
			Required: []string{"provider", "strategy"},
		},
	}
}

// NFD tells us to enable or disable the node feature discovery addon.
type NFD struct{}

func (NFD) VariableSchema() clusterv1.VariableSchema {
	return clusterv1.VariableSchema{
		OpenAPIV3Schema: clusterv1.JSONSchemaProps{
			Type: "object",
		},
	}
}

type CSIProviders struct {
	// +optional
	Providers []CSIProvider `json:"providers,omitempty"`
	// +optional
	DefaultClassName string `json:"defaultClassName,omitempty"`
}

type CSIProvider struct {
	Name string `json:"name,omitempty"`
}

func (CSIProviders) VariableSchema() clusterv1.VariableSchema {
	supportedCSIProviders := []string{CSIProviderAWSEBS}
	return clusterv1.VariableSchema{
		OpenAPIV3Schema: clusterv1.JSONSchemaProps{
			Type: "object",
			Properties: map[string]clusterv1.JSONSchemaProps{
				"providers": {
					Type: "array",
					Items: &clusterv1.JSONSchemaProps{
						Type: "object",
						Properties: map[string]clusterv1.JSONSchemaProps{
							"name": {
								Type: "string",
								Enum: variables.MustMarshalValuesToEnumJSON(
									supportedCSIProviders...),
							},
						},
					},
				},
				"defaultClassName": {
					Type: "string",
					Enum: variables.MustMarshalValuesToEnumJSON(supportedCSIProviders...),
				},
			},
		},
	}
}

// CPI tells us to enable or disable the cloud provider interface.
type CPI struct{}

func (CPI) VariableSchema() clusterv1.VariableSchema {
	return clusterv1.VariableSchema{
		OpenAPIV3Schema: clusterv1.JSONSchemaProps{
			Type: "object",
		},
	}
}
