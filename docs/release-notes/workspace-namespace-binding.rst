:orphan:

**Breaking Changes** 

-   Cluster: The ``kubernetes_namespace`` field under the resource pool config is
    no longer supported. Instead, the users can submit workloads to particular namespaces, by binding
    workspaces to namespaces using the CLI or API.

**New Features** 

-   Cluster: The ``namespace`` field under the Kubernetes Resource Manager config has
    been deprecated, and a new field called ``default_namespace`` has been added. This field will be the
    default namespace where Determined will deploy Pods and ConfigMaps, if the workspace that a workload
    is submitted in, is not bound to a particular namespace. Currently, the master config will accept
    either ``namespace`` or ``default_namespace`` fields, but if both are provided an error will be
    thrown.
