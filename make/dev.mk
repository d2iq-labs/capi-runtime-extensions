# Copyright 2023 Nutanix. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

ADDONS_PROVIDER := ClusterResourceSet

.PHONY: dev.run-on-kind
dev.run-on-kind: export KUBECONFIG := $(KIND_KUBECONFIG)
dev.run-on-kind: kind.create clusterctl.init
ifndef SKIP_BUILD
	$(MAKE) release-snapshot
endif
	kind load docker-image --name $(KIND_CLUSTER_NAME) \
		ko.local/cluster-api-runtime-extensions-nutanix:$$(gojq -r .version dist/metadata.json)
	helm upgrade --install cluster-api-runtime-extensions-nutanix ./charts/cluster-api-runtime-extensions-nutanix \
		--set-string image.repository=ko.local/cluster-api-runtime-extensions-nutanix \
		--set-string image.tag=$$(gojq -r .version dist/metadata.json) \
		--wait --wait-for-jobs
	kubectl rollout restart deployment cluster-api-runtime-extensions-nutanix
	kubectl rollout status deployment cluster-api-runtime-extensions-nutanix

dev.update-webhook-image-on-kind:
ifndef SKIP_BUILD
	$(MAKE) release-snapshot
endif
	kind load docker-image --name $(KIND_CLUSTER_NAME) \
		ko.local/cluster-api-runtime-extensions-nutanix:$$(gojq -r .version dist/metadata.json)
	kubectl set image deployment cluster-api-runtime-extensions-nutanix webhook=ko.local/cluster-api-runtime-extensions-nutanix:$$(gojq -r .version dist/metadata.json)
	kubectl rollout restart deployment cluster-api-runtime-extensions-nutanix
	kubectl rollout status deployment cluster-api-runtime-extensions-nutanix
