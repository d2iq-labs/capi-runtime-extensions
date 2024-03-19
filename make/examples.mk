# Copyright 2023 Nutanix. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

export KUBERNETES_VERSION := v1.27.5

export CLUSTERCTL_VERSION := $(shell clusterctl version -o short 2>/dev/null)
export CAPA_VERSION := $(shell cd hack/third-party/capa && go list -m -f '{{ .Version }}' sigs.k8s.io/cluster-api-provider-aws/v2)

.PHONY: examples.sync
examples.sync: ## Syncs the examples by fetching upstream examples and applying kustomize patches
	hack/examples/sync.sh
