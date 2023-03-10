<!--
 Copyright 2023 D2iQ, Inc. All rights reserved.
 SPDX-License-Identifier: Apache-2.0
 -->

# CAPI Runtime Extensions Server

See [upstream documentation](https://cluster-api.sigs.k8s.io/tasks/experimental-features/runtime-sdk/index.html).

## Development

To deploy a local build, either initial install to update an existing deployment, run:

```shell
make dev.run-on-kind
eval $(make kind.kubeconfig)
```

By default this will use the `ClusterResourceSet` addons provider. To use the `FluxHelmRelease` addons provider run:

```shell
make ADDONS_PROVIDER=FluxHelmRelease dev.run-on-kind
eval $(make kind.kubeconfig)
```

Pro-tip: to redeploy without rebuilding the binaries, images, etc (useful if you have only changed the Helm chart for
example), run:

```shell
make SKIP_BUILD=true dev.run-on-kind
```

To create a cluster with [clusterctl](https://cluster-api.sigs.k8s.io/user/quick-start.html), run:

```shell
clusterctl generate cluster capi-quickstart \
  --flavor development \
  --kubernetes-version v1.26.0 \
  --control-plane-machine-count=1 \
  --worker-machine-count=1 | \
  gojq --yaml-input --yaml-output --slurp \
    '.[] | (select( .kind=="Cluster").metadata.labels += {"capi-runtime-extensions.d2iq-labs.com/cni": "calico"})' | \
  kubectl apply -f -
```

Wait until control plane is ready:

```shell
kubectl wait clusters/capi-quickstart --for=condition=ControlPlaneInitialized --timeout=5m
```

To get the kubeconfig for the new cluster, run:

```shell
clusterctl get kubeconfig capi-quickstart > capd-kubeconfig
```

If you are not on Linux, you will also need to fix the generated kubeconfig's `server`, run:

```shell
kubectl config set-cluster capi-quickstart \
  --kubeconfig capd-kubeconfig \
  --server=https://$(docker port capi-quickstart-lb 6443/tcp)
```

Wait until all nodes are ready (this indicates that CNI has been deployed successfully):

```shell
kubectl --kubeconfig capd-kubeconfig wait nodes --all --for=condition=Ready --timeout=5m
```

Show that Calico is running successfully on the workload cluster:

```shell
kubectl --kubeconfig capd-kubeconfig get daemonsets -n calico-system
```

To delete the workload cluster, run:

```shell
kubectl delete cluster capi-quickstart
```

To delete the dev KinD cluster, run:

```shell
make kind.delete
```
