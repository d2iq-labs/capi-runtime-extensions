// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package instancetype

import (
	"context"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	runtimehooksv1 "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1"

	"github.com/d2iq-labs/capi-runtime-extensions/api/v1alpha1"
	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/clustertopology/handlers"
	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/capi/clustertopology/handlers/mutation"
)

var (
	_ handlers.Named             = &awsInstanceTypeVariableHandler{}
	_ mutation.DiscoverVariables = &awsInstanceTypeVariableHandler{}
)

const (
	// VariableName is the external patch variable name.
	VariableName = "instanceType"

	// HandlerNameVariable is the name of the variable handler.
	HandlerNameVariable = "AWSInstanceTypeVars"
)

func NewVariable() *awsInstanceTypeVariableHandler {
	return &awsInstanceTypeVariableHandler{}
}

type awsInstanceTypeVariableHandler struct{}

func (h *awsInstanceTypeVariableHandler) Name() string {
	return HandlerNameVariable
}

func (h *awsInstanceTypeVariableHandler) DiscoverVariables(
	ctx context.Context,
	_ *runtimehooksv1.DiscoverVariablesRequest,
	resp *runtimehooksv1.DiscoverVariablesResponse,
) {
	resp.Variables = append(resp.Variables, clusterv1.ClusterClassVariable{
		Name:     VariableName,
		Required: false,
		Schema:   v1alpha1.InstanceType("").VariableSchema(),
	})
	resp.SetStatus(runtimehooksv1.ResponseStatusSuccess)
}