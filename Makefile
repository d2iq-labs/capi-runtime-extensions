# Copyright 2023 Nutanix. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

REPO_ROOT := $(CURDIR)

# Versions for tools that are not managed by devbox.
ENVTEST_VERSION=1.29.x

include make/all.mk
