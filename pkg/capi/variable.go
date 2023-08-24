// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package capi

import (
	"encoding/json"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"sigs.k8s.io/cluster-api/exp/runtime/topologymutation"
)

// GetVariable finds and parses variable to given type.
func GetVariable[T any](variables map[string]apiextensionsv1.JSON, name string) (T, bool, error) {
	var v T
	variable, found, err := topologymutation.GetVariable(variables, name)
	if err != nil || !found {
		return v, found, err
	}

	err = json.Unmarshal(variable.Raw, &v)
	return v, err == nil, err
}