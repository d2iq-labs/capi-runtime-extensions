---
title: "Cluster Config"
---

The Cluster Config handlers wrap all the other mutation handlers in a convenient single patch for inclusion in your
ClusterClasses, allowing for a single configuration variable with nested values. This provides the most flexibility
with the least configuration. The included patches are usable individually, but require declaring all the individual
patch and variable handlers in the ClusterClass.

To enable the API server certificate SANs enable the `clusterconfigvars` and `clusterconfigpatch`
external patches on `ClusterClass`.

```yaml
apiVersion: cluster.x-k8s.io/v1beta1
kind: ClusterClass
metadata:
  name: <NAME>
spec:
  patches:
    - name: cluster-config
      external:
        generateExtension: "clusterconfigpatch.<external-config-name>"
        discoverVariablesExtension: "clusterconfigvars.<external-config-name>"
```

On the cluster resource then specify desired certificate SANs values:

```yaml
apiVersion: cluster.x-k8s.io/v1beta1
kind: Cluster
metadata:
  name: <NAME>
spec:
  topology:
    variables:
      - name: clusterConfig
        value:
          extraAPIServerCertSANs:
            - a.b.c.example.com
            - d.e.f.example.com
          proxy:
            http: http://example.com
            https: https://example.com
            additionalNo:
              - no-proxy-1.example.com
              - no-proxy-2.example.com
```

Applying this configuration will result in the certificate SANs being correctly set in the
`KubeadmControlPlaneTemplate`.