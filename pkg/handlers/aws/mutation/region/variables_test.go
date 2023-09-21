// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package region

import (
	"testing"

	"k8s.io/utils/ptr"

	"github.com/d2iq-labs/capi-runtime-extensions/api/v1alpha1"
	"github.com/d2iq-labs/capi-runtime-extensions/common/pkg/testutils/capitest"
)

func TestVariableValidation(t *testing.T) {
	capitest.ValidateDiscoverVariables(
		t,
		variableName,
		ptr.To(v1alpha1.Region("").VariableSchema()),
		false,
		NewVariable,
		capitest.VariableTestDef{
			Name: "valid",
			Vals: "us-west-2",
		},
		capitest.VariableTestDef{
			Name:        "invalid",
			Vals:        "invalid region",
			ExpectError: true,
		},
	)
}
