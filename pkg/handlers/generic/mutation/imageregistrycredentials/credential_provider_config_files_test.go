// Copyright 2023 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package imageregistrycredentials

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cabpkv1 "sigs.k8s.io/cluster-api/bootstrap/kubeadm/api/v1beta1"
)

func Test_templateKubeletCredentialProviderConfig(t *testing.T) {
	tests := []struct {
		name        string
		credentials imageRegistryCredentials
		want        *cabpkv1.File
		wantErr     error
	}{
		{
			name:        "ECR image registry",
			credentials: imageRegistryCredentials{URL: "https://123456789.dkr.ecr.us-east-1.amazonaws.com"},
			want: &cabpkv1.File{
				Path:        "/etc/kubernetes/image-credential-provider-config.yaml",
				Owner:       "",
				Permissions: "0600",
				Encoding:    "",
				Append:      false,
				Content: `apiVersion: kubelet.config.k8s.io/v1beta1
kind: CredentialProviderConfig
providers:
- name: dynamic-credential-provider
  args:
  - get-credentials
  - -c
  - /etc/kubernetes/dynamic-credential-provider-config.yaml
  matchImages:
  - "123456789.dkr.ecr.us-east-1.amazonaws.com"
  - "*"
  - "*.*"
  - "*.*.*"
  - "*.*.*.*"
  - "*.*.*.*.*"
  - "*.*.*.*.*.*"
  defaultCacheDuration: "0s"
  apiVersion: credentialprovider.kubelet.k8s.io/v1beta1
`,
			},
		},
		{
			name: "image registry with static credentials",
			credentials: imageRegistryCredentials{
				URL:      "https://myregistry.com",
				Username: "myuser",
				Password: "mypassword",
			},
			want: &cabpkv1.File{
				Path:        "/etc/kubernetes/image-credential-provider-config.yaml",
				Owner:       "",
				Permissions: "0600",
				Encoding:    "",
				Append:      false,
				Content: `apiVersion: kubelet.config.k8s.io/v1beta1
kind: CredentialProviderConfig
providers:
- name: dynamic-credential-provider
  args:
  - get-credentials
  - -c
  - /etc/kubernetes/dynamic-credential-provider-config.yaml
  matchImages:
  - "myregistry.com"
  - "*"
  - "*.*"
  - "*.*.*"
  - "*.*.*.*"
  - "*.*.*.*.*"
  - "*.*.*.*.*.*"
  defaultCacheDuration: "0s"
  apiVersion: credentialprovider.kubelet.k8s.io/v1beta1
`,
			},
		},
	}
	for idx := range tests {
		tt := tests[idx]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			file, err := templateKubeletCredentialProviderConfig(tt.credentials)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, file)
		})
	}
}

func Test_templateDynamicCredentialProviderConfig(t *testing.T) {
	tests := []struct {
		name        string
		credentials imageRegistryCredentials
		want        *cabpkv1.File
		wantErr     error
	}{
		{
			name:        "ECR image registry",
			credentials: imageRegistryCredentials{URL: "https://123456789.dkr.ecr.us-east-1.amazonaws.com"},
			want: &cabpkv1.File{
				Path:        "/etc/kubernetes/dynamic-credential-provider-config.yaml",
				Owner:       "",
				Permissions: "0600",
				Encoding:    "",
				Append:      false,
				Content: `apiVersion: credentialprovider.d2iq.com/v1alpha1
kind: DynamicCredentialProviderConfig
credentialProviderPluginBinDir: /etc/kubernetes/image-credential-provider/
mirror:
  endpoint: "123456789.dkr.ecr.us-east-1.amazonaws.com"
  credentialsStrategy: MirrorCredentialsOnly
credentialProviders:
  apiVersion: kubelet.config.k8s.io/v1beta1
  kind: CredentialProviderConfig
  providers:
  - name: ecr-credential-provider
    args:
    - get-credentials
    matchImages:
    - "123456789.dkr.ecr.us-east-1.amazonaws.com"
    defaultCacheDuration: "0s"
    apiVersion: credentialprovider.kubelet.k8s.io/v1alpha1
`,
			},
		},
		{
			name: "image registry with static credentials",
			credentials: imageRegistryCredentials{
				URL:      "https://myregistry.com",
				Username: "myuser",
				Password: "mypassword",
			},
			want: &cabpkv1.File{
				Path:        "/etc/kubernetes/dynamic-credential-provider-config.yaml",
				Owner:       "",
				Permissions: "0600",
				Encoding:    "",
				Append:      false,
				Content: `apiVersion: credentialprovider.d2iq.com/v1alpha1
kind: DynamicCredentialProviderConfig
credentialProviderPluginBinDir: /etc/kubernetes/image-credential-provider/
mirror:
  endpoint: "myregistry.com"
  credentialsStrategy: MirrorCredentialsOnly
credentialProviders:
  apiVersion: kubelet.config.k8s.io/v1beta1
  kind: CredentialProviderConfig
  providers:
  - name: static-credential-provider
    args:
    - /etc/kubernetes/static-image-credentials.json
    matchImages:
    - "myregistry.com"
    defaultCacheDuration: "0s"
    apiVersion: credentialprovider.kubelet.k8s.io/v1beta1
`,
			},
		},
	}
	for idx := range tests {
		tt := tests[idx]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			file, err := templateDynamicCredentialProviderConfig(tt.credentials)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, file)
		})
	}
}
