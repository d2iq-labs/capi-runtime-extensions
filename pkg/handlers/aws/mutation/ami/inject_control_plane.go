// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package ami

import (
	"github.com/d2iq-labs/capi-runtime-extensions/api/v1alpha1"
	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/clustertopology/patches/selectors"
	capav1 "github.com/d2iq-labs/capi-runtime-extensions/common/pkg/external/sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	"github.com/d2iq-labs/capi-runtime-extensions/pkg/handlers/generic/clusterconfig"
)

func NewControlPlanePatch() *awsAMISpecPatchHandler {
	return newAWSAMISpecPatchHandler(
		clusterconfig.MetaVariableName,
		[]string{
			clusterconfig.MetaControlPlaneConfigName,
			v1alpha1.AWSVariableName,
			VariableName,
		},
		selectors.InfrastructureControlPlaneMachines(
			capav1.GroupVersion.Version,
			"AWSMachineTemplate",
		),
	)
}