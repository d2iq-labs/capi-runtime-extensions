# Copyright 2023 D2iQ, Inc. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

export CALICO_VERSION := $(shell goprintconst -file pkg/handlers/generic/lifecycle/cni/calico/strategy_helmaddon.go -name defaultCalicoHelmChartVersion)
export CILIUM_VERSION := $(shell goprintconst -file pkg/handlers/generic/lifecycle/cni/cilium/strategy_helmaddon.go -name defaultCiliumHelmChartVersion)
export NODE_FEATURE_DISCOVERY_VERSION := $(shell goprintconst -file pkg/handlers/generic/lifecycle/nfd/strategy_helmaddon.go -name defaultHelmChartVersion)
export CLUSTER_AUTOSCALER_VERSION := 9.35.0
export AWS_CSI_SNAPSHOT_CONTROLLER_VERSION := v6.3.3
export AWS_EBS_CSI_CHART_VERSION := v2.28.1
# a map of AWS CPI versions
export AWS_CPI_VERSION_127 := v1.27.1
export AWS_CPI_CHART_VERSION_127 := 0.0.8
export AWS_CPI_VERSION_128 := v1.28.1
export AWS_CPI_CHART_VERSION_128 := 0.0.8

.PHONY: addons.sync
addons.sync: $(addprefix update-addon.,calico cilium nfd cluster-autoscaler aws-ebs-csi aws-cpi.127 aws-cpi.128)

.PHONY: update-addon.calico
update-addon.calico: ; $(info $(M) updating calico manifests)
	./hack/addons/update-calico-manifests.sh

.PHONY: update-addon.cilium
update-addon.cilium: ; $(info $(M) updating cilium manifests)
	./hack/addons/update-cilium-manifests.sh

.PHONY: update-addon.nfd
update-addon.nfd: ; $(info $(M) updating node feature discovery manifests)
	./hack/addons/update-node-feature-discovery-manifests.sh

.PHONY: update-addon.cluster-autoscaler
update-addon.cluster-autoscaler: ; $(info $(M) updating cluster-autoscaler manifests)
	./hack/addons/update-cluster-autoscaler.sh

.PHONY: update-addon.aws-ebs-csi
update-addon.aws-ebs-csi: ; $(info $(M) updating aws ebs csi manifests)
	./hack/addons/update-aws-ebs-csi.sh

.PHONY: update-addon.aws-cpi.%
update-addon.aws-cpi.%: ; $(info $(M) updating aws cpi $* manifests)
	./hack/addons/update-aws-cpi.sh $(AWS_CPI_VERSION_$*) $(AWS_CPI_CHART_VERSION_$*)
