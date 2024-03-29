// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package tests

import (
	"context"
	"testing"

	"github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apiserver/pkg/storage/names"
	runtimehooksv1 "sigs.k8s.io/cluster-api/exp/runtime/hooks/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/api/v1alpha1"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/capi/clustertopology/handlers/mutation"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/testutils/capitest"
	"github.com/d2iq-labs/cluster-api-runtime-extensions-nutanix/common/pkg/testutils/capitest/request"
)

const (
	validSecretName                       = "myregistry-credentials"
	registryStaticCredentialsSecretSuffix = "registry-creds"
)

func TestGeneratePatches(
	t *testing.T,
	generatorFunc func() mutation.GeneratePatches,
	fakeClient client.Client,
	variableName string,
	variablePath ...string,
) {
	t.Helper()

	require.NoError(
		t,
		fakeClient.Create(
			context.Background(),
			newRegistryCredentialsSecret(validSecretName, request.Namespace),
		),
	)

	capitest.ValidateGeneratePatches(
		t,
		generatorFunc,
		capitest.PatchTestDef{
			Name: "unset variable",
		},
		capitest.PatchTestDef{
			Name: "files added in KubeadmControlPlaneTemplate for ECR without a Secret",
			Vars: []runtimehooksv1.Variable{
				capitest.VariableWithValue(
					variableName,
					v1alpha1.ImageRegistries{
						v1alpha1.ImageRegistry{
							URL: "https://123456789.dkr.ecr.us-east-1.amazonaws.com",
						},
					},
					variablePath...,
				),
			},
			RequestItem: request.NewKubeadmControlPlaneTemplateRequestItem(""),
			ExpectedPatchMatchers: []capitest.JSONPatchMatcher{
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/files",
					ValueMatcher: gomega.ContainElements(
						gomega.HaveKeyWithValue(
							"path", "/etc/cre/install-kubelet-credential-providers.sh",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/image-credential-provider-config.yaml",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/dynamic-credential-provider-config.yaml",
						),
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/preKubeadmCommands",
					ValueMatcher: gomega.ContainElement(
						"/bin/bash /etc/cre/install-kubelet-credential-providers.sh",
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/initConfiguration/nodeRegistration/kubeletExtraArgs",
					ValueMatcher: gomega.HaveKeyWithValue(
						"image-credential-provider-bin-dir",
						"/etc/kubernetes/image-credential-provider/",
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/joinConfiguration/nodeRegistration/kubeletExtraArgs",
					ValueMatcher: gomega.HaveKeyWithValue(
						"image-credential-provider-config",
						"/etc/kubernetes/image-credential-provider-config.yaml",
					),
				},
			},
		},
		capitest.PatchTestDef{
			Name: "files added in KubeadmControlPlaneTemplate for registry with a Secret",
			Vars: []runtimehooksv1.Variable{
				capitest.VariableWithValue(
					variableName,
					v1alpha1.ImageRegistries{
						v1alpha1.ImageRegistry{
							URL: "https://registry.example.com",
							Credentials: &v1alpha1.RegistryCredentials{
								SecretRef: &corev1.LocalObjectReference{
									Name: validSecretName,
								},
							},
						},
					},
					variablePath...,
				),
			},
			RequestItem: request.NewKubeadmControlPlaneTemplateRequest(
				"",
				"test-kubeadmconfigtemplate",
			),
			ExpectedPatchMatchers: []capitest.JSONPatchMatcher{
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/files",
					ValueMatcher: gomega.ContainElements(
						gomega.HaveKeyWithValue(
							"path", "/etc/cre/install-kubelet-credential-providers.sh",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/image-credential-provider-config.yaml",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/dynamic-credential-provider-config.yaml",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/static-image-credentials.json",
						),
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/preKubeadmCommands",
					ValueMatcher: gomega.ContainElement(
						"/bin/bash /etc/cre/install-kubelet-credential-providers.sh",
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/initConfiguration/nodeRegistration/kubeletExtraArgs",
					ValueMatcher: gomega.HaveKeyWithValue(
						"image-credential-provider-bin-dir",
						"/etc/kubernetes/image-credential-provider/",
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/kubeadmConfigSpec/joinConfiguration/nodeRegistration/kubeletExtraArgs",
					ValueMatcher: gomega.HaveKeyWithValue(
						"image-credential-provider-config",
						"/etc/kubernetes/image-credential-provider-config.yaml",
					),
				},
			},
		},
		capitest.PatchTestDef{
			Name: "files added in KubeadmConfigTemplate for ECR without a Secret",
			Vars: []runtimehooksv1.Variable{
				capitest.VariableWithValue(
					variableName,
					v1alpha1.ImageRegistries{
						v1alpha1.ImageRegistry{
							URL: "https://123456789.dkr.ecr.us-east-1.amazonaws.com",
						},
					},
					variablePath...,
				),
				capitest.VariableWithValue(
					"builtin",
					map[string]any{
						"machineDeployment": map[string]any{
							"class": names.SimpleNameGenerator.GenerateName("worker-"),
						},
					},
				),
			},
			RequestItem: request.NewKubeadmConfigTemplateRequestItem(""),
			ExpectedPatchMatchers: []capitest.JSONPatchMatcher{
				{
					Operation: "add",
					Path:      "/spec/template/spec/files",
					ValueMatcher: gomega.ContainElements(
						gomega.HaveKeyWithValue(
							"path", "/etc/cre/install-kubelet-credential-providers.sh",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/image-credential-provider-config.yaml",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/dynamic-credential-provider-config.yaml",
						),
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/preKubeadmCommands",
					ValueMatcher: gomega.ContainElement(
						"/bin/bash /etc/cre/install-kubelet-credential-providers.sh",
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/joinConfiguration/nodeRegistration/kubeletExtraArgs",
					ValueMatcher: gomega.HaveKeyWithValue(
						"image-credential-provider-bin-dir",
						"/etc/kubernetes/image-credential-provider/",
					),
				},
			},
		},
		capitest.PatchTestDef{
			Name: "files added in KubeadmConfigTemplate for registry with a Secret",
			Vars: []runtimehooksv1.Variable{
				capitest.VariableWithValue(
					variableName,
					v1alpha1.ImageRegistries{
						v1alpha1.ImageRegistry{
							URL: "https://registry.example.com",
							Credentials: &v1alpha1.RegistryCredentials{
								SecretRef: &corev1.LocalObjectReference{
									Name: validSecretName,
								},
							},
						},
					},
					variablePath...,
				),
				capitest.VariableWithValue(
					"builtin",
					map[string]any{
						"machineDeployment": map[string]any{
							"class": names.SimpleNameGenerator.GenerateName("worker-"),
						},
					},
				),
			},
			RequestItem: request.NewKubeadmConfigTemplateRequest("", "test-kubeadmconfigtemplate"),
			ExpectedPatchMatchers: []capitest.JSONPatchMatcher{
				{
					Operation: "add",
					Path:      "/spec/template/spec/files",
					ValueMatcher: gomega.ContainElements(
						gomega.HaveKeyWithValue(
							"path", "/etc/cre/install-kubelet-credential-providers.sh",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/image-credential-provider-config.yaml",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/dynamic-credential-provider-config.yaml",
						),
						gomega.HaveKeyWithValue(
							"path", "/etc/kubernetes/static-image-credentials.json",
						),
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/preKubeadmCommands",
					ValueMatcher: gomega.ContainElement(
						"/bin/bash /etc/cre/install-kubelet-credential-providers.sh",
					),
				},
				{
					Operation: "add",
					Path:      "/spec/template/spec/joinConfiguration/nodeRegistration/kubeletExtraArgs",
					ValueMatcher: gomega.HaveKeyWithValue(
						"image-credential-provider-bin-dir",
						"/etc/kubernetes/image-credential-provider/",
					),
				},
			},
		},
		capitest.PatchTestDef{
			Name: "error for a registry with no credentials",
			Vars: []runtimehooksv1.Variable{
				capitest.VariableWithValue(
					variableName,
					v1alpha1.ImageRegistries{
						v1alpha1.ImageRegistry{
							URL: "https://registry.example.com",
						},
					},
					variablePath...,
				),
			},
			RequestItem:     request.NewKubeadmControlPlaneTemplateRequestItem(""),
			ExpectedFailure: true,
		},
	)
}

func newRegistryCredentialsSecret(name, namespace string) *corev1.Secret {
	secretData := map[string][]byte{
		"username": []byte("myuser"),
		"password": []byte("mypassword"),
	}
	return &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: secretData,
		Type: corev1.SecretTypeOpaque,
	}
}
