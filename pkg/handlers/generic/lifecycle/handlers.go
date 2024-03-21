// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package lifecycle

import (
	"github.com/spf13/pflag"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/api/v1alpha1"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/capi/clustertopology/handlers"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/clusterautoscaler"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/cni/calico"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/cni/cilium"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/cpi"
	awscpi "github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/cpi/aws"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/csi"
	awsebs "github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/csi/aws-ebs"
	nutanixcsi "github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/csi/nutanix-csi"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/nfd"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/generic/lifecycle/servicelbgc"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/pkg/handlers/options"
)

type Handlers struct {
	calicoCNIConfig         *calico.CNIConfig
	ciliumCNIConfig         *cilium.CNIConfig
	nfdConfig               *nfd.Config
	clusterAutoscalerConfig *clusterautoscaler.Config
	ebsConfig               *awsebs.AWSEBSConfig
	nutnaixCSIConfig        *nutanixcsi.NutanixCSIConfig
	awsCPIConfig            *awscpi.AWSCPIConfig
}

func New(globalOptions *options.GlobalOptions) *Handlers {
	return &Handlers{
		calicoCNIConfig:         &calico.CNIConfig{GlobalOptions: globalOptions},
		ciliumCNIConfig:         &cilium.CNIConfig{GlobalOptions: globalOptions},
		nfdConfig:               &nfd.Config{GlobalOptions: globalOptions},
		clusterAutoscalerConfig: &clusterautoscaler.Config{GlobalOptions: globalOptions},
		ebsConfig:               &awsebs.AWSEBSConfig{GlobalOptions: globalOptions},
		awsCPIConfig:            &awscpi.AWSCPIConfig{GlobalOptions: globalOptions},
		nutnaixCSIConfig:        &nutanixcsi.NutanixCSIConfig{GlobalOptions: globalOptions},
	}
}

func (h *Handlers) AllHandlers(mgr manager.Manager) []handlers.Named {
	csiHandlers := map[string]csi.CSIProvider{
		v1alpha1.CSIProviderAWSEBS:  awsebs.New(mgr.GetClient(), h.ebsConfig),
		v1alpha1.CSIProviderNutanix: nutanixcsi.New(mgr.GetClient(), h.nutnaixCSIConfig),
	}
	cpiHandlers := map[string]cpi.CPIProvider{
		v1alpha1.CPIProviderAWS: awscpi.New(mgr.GetClient(), h.awsCPIConfig),
	}

	return []handlers.Named{
		calico.New(mgr.GetClient(), h.calicoCNIConfig),
		cilium.New(mgr.GetClient(), h.ciliumCNIConfig),
		nfd.New(mgr.GetClient(), h.nfdConfig),
		clusterautoscaler.New(mgr.GetClient(), h.clusterAutoscalerConfig),
		servicelbgc.New(mgr.GetClient()),
		csi.New(mgr.GetClient(), csiHandlers),
		cpi.New(mgr.GetClient(), cpiHandlers),
	}
}

func (h *Handlers) AddFlags(flagSet *pflag.FlagSet) {
	h.nfdConfig.AddFlags("nfd", flagSet)
	h.clusterAutoscalerConfig.AddFlags("cluster-autoscaler", flagSet)
	h.calicoCNIConfig.AddFlags("cni.calico", flagSet)
	h.ciliumCNIConfig.AddFlags("cni.cilium", flagSet)
	h.ebsConfig.AddFlags("awsebs", pflag.CommandLine)
	h.nutnaixCSIConfig.AddFlags("nutanixcsi", flagSet)
	h.awsCPIConfig.AddFlags("awscpi", pflag.CommandLine)
}
