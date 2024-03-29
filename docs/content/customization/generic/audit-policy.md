+++
title = "Audit policy"
+++

Kubernetes auditing provides a security-relevant, chronological set of records documenting the sequence of actions in a
cluster. The cluster audits the activities generated by users, by applications that use the Kubernetes API, and by the
control plane itself.

There are currently no configuration options for the Audit Policy customization and this customization will be
automatically applied when the [provider-specific cluster configuration patch]({{< ref ".." >}}) is included in the
`ClusterClass`.
