CREATE TABLE public.workspace_namespace_bindings (
    workspace_id INT REFERENCES workspaces(id) ON DELETE CASCADE,
    cluster_name text NOT NULL,
    namespace_name text NOT NULL
);
