// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package instancetype

import (
	"context"
	_ "embed"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	runtimehooksv1 "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/d2iq-labs/capi-runtime-extensions/api/v1alpha1"
	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/apis"
	commonhandlers "github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/clustertopology/handlers"
	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/clustertopology/handlers/mutation"
	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/clustertopology/patches"
	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/clustertopology/patches/selectors"
	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/clustertopology/variables"
	capav1 "github.com/d2iq-labs/capi-runtime-extensions/common/pkg/external/sigs.k8s.io/cluster-api-provider-aws/v2/api/v1beta2"
	"github.com/d2iq-labs/capi-runtime-extensions/pkg/handlers"
	awsclusterconfig "github.com/d2iq-labs/capi-runtime-extensions/pkg/handlers/aws/clusterconfig"
	"github.com/d2iq-labs/capi-runtime-extensions/pkg/handlers/generic/clusterconfig"
)

const (
	// VariableName is the external patch variable name.
	VariableName = "instanceType"

	// WorkerHandlerNamePatch is the name of the inject handler.
	ControlPlaneHandlerNamePatch = "AWSInstanceTypeControlPlanePatch"
)

type awsInstanceTypeControlPlanePatchHandler struct {
	variableName      string
	variableFieldPath []string
}

var (
	_ commonhandlers.Named     = &awsInstanceTypeControlPlanePatchHandler{}
	_ mutation.GeneratePatches = &awsInstanceTypeControlPlanePatchHandler{}
	_ mutation.MetaMutator     = &awsInstanceTypeControlPlanePatchHandler{}
)

func NewControlPlaneMetaPatch() *awsInstanceTypeControlPlanePatchHandler {
	return newAWSInstanceTypeControlPlanePatchHandler(
		clusterconfig.MetaVariableName,
		clusterconfig.MetaControlPlaneConfigName,
		awsclusterconfig.AWSVariableName,
		VariableName,
	)
}

func newAWSInstanceTypeControlPlanePatchHandler(
	variableName string,
	variableFieldPath ...string,
) *awsInstanceTypeControlPlanePatchHandler {
	return &awsInstanceTypeControlPlanePatchHandler{
		variableName:      variableName,
		variableFieldPath: variableFieldPath,
	}
}

func (h *awsInstanceTypeControlPlanePatchHandler) Name() string {
	return ControlPlaneHandlerNamePatch
}

func (h *awsInstanceTypeControlPlanePatchHandler) Mutate(
	ctx context.Context,
	obj runtime.Object,
	vars map[string]apiextensionsv1.JSON,
	holderRef runtimehooksv1.HolderReference,
	_ client.ObjectKey,
) error {
	log := ctrl.LoggerFrom(ctx).WithValues(
		"holderRef", holderRef,
	)

	instanceTypeVar, found, err := variables.Get[v1alpha1.InstanceType](
		vars,
		h.variableName,
		h.variableFieldPath...,
	)
	if err != nil {
		return err
	}
	if !found {
		log.V(5).Info("AWS instance type variable for control-plane not defined")
		return nil
	}

	log = log.WithValues(
		"variableName",
		h.variableName,
		"variableFieldPath",
		h.variableFieldPath,
		"variableValue",
		instanceTypeVar,
	)

	return patches.Generate(
		obj,
		vars,
		&holderRef,
		selectors.InfrastructureControlPlaneMachines(
			"v1beta2",
			"AWSMachineTemplate",
		),
		log,
		func(obj *capav1.AWSMachineTemplate) error {
			log.WithValues(
				"patchedObjectKind", obj.GetObjectKind().GroupVersionKind().String(),
				"patchedObjectName", client.ObjectKeyFromObject(obj),
			).Info("setting instance type in control plane AWSMachineTemplate spec")

			obj.Spec.Template.Spec.InstanceType = string(instanceTypeVar)

			return nil
		},
	)
}

func (h *awsInstanceTypeControlPlanePatchHandler) GeneratePatches(
	ctx context.Context,
	req *runtimehooksv1.GeneratePatchesRequest,
	resp *runtimehooksv1.GeneratePatchesResponse,
) {
	handlers.GeneratePatches(ctx, req, resp, apis.CAPADecoder(), h.Mutate)
}
