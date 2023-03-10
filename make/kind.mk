# Copyright 2023 D2iQ, Inc. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

KIND_DIR := $(REPO_ROOT)/.local/kind

KIND_CLUSTER_NAME ?= $(GITHUB_REPOSITORY)-dev
KIND_KUBECONFIG ?= $(KIND_DIR)/$(KIND_CLUSTER_NAME)/kubeconfig

KINDEST_NODE_IMAGE ?= kindest/node
KINDEST_NODE_VERSION_v1.24 ?= v1.24.7@sha256:577c630ce8e509131eab1aea12c022190978dd2f745aac5eb1fe65c0807eb315
KINDEST_NODE_VERSION_v1.25 ?= v1.25.3@sha256:f52781bc0d7a19fb6c405c2af83abfeb311f130707a0e219175677e366cc45d1
KINDEST_NODE_VERSION_v1.26 ?= v1.26.0@sha256:691e24bd2417609db7e589e1a479b902d2e209892a10ce375fab60a8407c7352
# Allow easy override of Kubernetes version to use via `make KIND_KUBERNETES_VERSION=v1.23` to use in CI
KIND_KUBERNETES_VERSION ?= v1.26
ifndef KINDEST_NODE_VERSION_$(KIND_KUBERNETES_VERSION)
  $(error Unsupported Kind Kubernetes version: $(KIND_KUBERNETES_VERSION) (use on of: [$(patsubst KINDEST_NODE_VERSION_%,%,$(filter KINDEST_NODE_VERSION_%,$(.VARIABLES)))]))
endif

KINDEST_IMAGE = $(KINDEST_NODE_IMAGE):$(KINDEST_NODE_VERSION_$(KIND_KUBERNETES_VERSION))

.PHONY: kind.recreate
kind.recreate: ## Re-creates new KinD cluster if necessary
kind.recreate: kind.delete kind.create

.PHONY: kind.create
kind.create: ## Creates new KinD cluster
kind.create: install-tool.kubectl install-tool.kind install-tool.go.envsubst ; $(info $(M) creating kind cluster - $(KIND_CLUSTER_NAME))
	(kind get clusters | grep -Eq '^$(KIND_CLUSTER_NAME)$$' && echo '$(KIND_CLUSTER_NAME) already exists') || \
		env KUBECONFIG=$(KIND_KUBECONFIG) $(REPO_ROOT)/hack/kind/create-cluster.sh \
		  --cluster-name $(KIND_CLUSTER_NAME) \
		  --kindest-image $(KINDEST_IMAGE) \
		  --output-dir $(KIND_DIR)/$(KIND_CLUSTER_NAME) \
		  --base-config $(REPO_ROOT)/hack/kind/kind-base-config.yaml

.PHONY: kind.delete
kind.delete: ## Deletes KinD cluster
kind.delete: install-tool.kind ; $(info $(M) deleting kind cluster - $(KIND_CLUSTER_NAME))
	(kind get clusters | grep -Eq '^$(KIND_CLUSTER_NAME)$$' && kind delete cluster --name $(KIND_CLUSTER_NAME)) || \
	  echo '$(KIND_CLUSTER_NAME) does not exist'
	rm -rf $(KIND_DIR)/$(KIND_CLUSTER_NAME)

.PHONY: kind.kubeconfig
kind.kubeconfig: ## Prints export definition for kubeconfig
	echo "export KUBECONFIG=$(KIND_KUBECONFIG)"
