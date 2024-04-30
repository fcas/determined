//go:build integration
// +build integration

package workspace

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/pkg/etc"
	"github.com/determined-ai/determined/proto/pkg/workspacev1"
)

const (
	cluster1 = "C1"
	cluster2 = "C2"
)

func TestMain(m *testing.M) {
	pgDB, err := db.ResolveTestPostgres()
	if err != nil {
		log.Panicln(err)
	}

	err = db.MigrateTestPostgres(pgDB, "file://../../static/migrations", "up")
	if err != nil {
		log.Panicln(err)
	}

	err = etc.SetRootPath("../../static/srv")
	if err != nil {
		log.Panicln(err)
	}

	os.Exit(m.Run())
}

type Bindings struct {
	bun.BaseModel `bun:"table:workspace_namespace_bindings"`
	WorkspaceID   int    `bun:"workspace_id"`
	ClusterName   string `bun:"cluster_name"`
	NamespaceName string `bun:"namespace_name"`
}

func TestGetNamespaceFromWorkspace(t *testing.T) {
	ctx := context.Background()
	wksp1 := "wksp_1"
	wkspID1, _ := db.RequireMockWorkspaceID(t, db.SingleDB(), wksp1)
	wksp2 := "wksp_2"
	wkspID2, _ := db.RequireMockWorkspaceID(t, db.SingleDB(), wksp2)
	wksp3 := "wksp_3"
	wkspID3, _ := db.RequireMockWorkspaceID(t, db.SingleDB(), wksp3)
	cluster1 := "C1"
	cluster2 := "C2"

	b1 := Bindings{WorkspaceID: wkspID1, ClusterName: cluster1, NamespaceName: "n1"}
	b2 := Bindings{WorkspaceID: wkspID1, ClusterName: cluster2, NamespaceName: "n2"}
	b3 := Bindings{WorkspaceID: wkspID2, ClusterName: cluster1, NamespaceName: "n3"}

	bindings := []Bindings{b1, b2, b3}

	_, err := db.Bun().NewInsert().Model(&bindings).Exec(ctx)
	require.NoError(t, err)

	tests := []struct {
		name          string
		binding       Bindings
		workspaceName string
		expectNoError bool
	}{
		{"simple-1", b1, wksp1, true},
		{"simple-2", b2, wksp1, true},
		{"simple-3", b3, wksp2, true},
		{"fail-1", Bindings{WorkspaceID: wkspID2, ClusterName: cluster2, NamespaceName: ""}, wksp2, false},
		{"fail-2", Bindings{WorkspaceID: wkspID3, ClusterName: cluster1, NamespaceName: ""}, wksp3, false},
	}

	for _, tt := range tests {
		ns, err := GetNamespaceFromWorkspace(ctx, tt.workspaceName, tt.binding.ClusterName)
		if tt.expectNoError {
			require.NoError(t, err)
			require.Equal(t, ns, tt.binding.NamespaceName)
		} else {
			require.ErrorContains(t, err, "no rows")
		}
	}

	// Clean-up
	holder := &workspacev1.Workspace{}
	err = db.SingleDB().QueryProto("delete_workspace", holder, wkspID1)
	require.NoError(t, err)
	err = db.SingleDB().QueryProto("delete_workspace", holder, wkspID2)
	require.NoError(t, err)
	err = db.SingleDB().QueryProto("delete_workspace", holder, wkspID3)
	require.NoError(t, err)
}

func TestGetAllNamespacesForRM(t *testing.T) {
	ctx := context.Background()
	wkspID1, _ := db.RequireMockWorkspaceID(t, db.SingleDB(), "wksp1")
	wkspID2, _ := db.RequireMockWorkspaceID(t, db.SingleDB(), "wksp2")

	b1 := Bindings{WorkspaceID: wkspID1, ClusterName: cluster1, NamespaceName: "n1"}
	b2 := Bindings{WorkspaceID: wkspID1, ClusterName: cluster2, NamespaceName: "n2"}
	b3 := Bindings{WorkspaceID: wkspID2, ClusterName: cluster1, NamespaceName: "n3"}

	bindings := []Bindings{b1, b2, b3}

	_, err := db.Bun().NewInsert().Model(&bindings).Exec(ctx)
	require.NoError(t, err)

	ns, err := GetAllNamespacesForRM(ctx, cluster1, "")
	require.NoError(t, err)
	require.Equal(t, ns, []string{"n1", "n3", "default"})

	ns, err = GetAllNamespacesForRM(ctx, cluster2, "test")
	require.NoError(t, err)
	require.Equal(t, ns, []string{"n2", "test"})

	ns, err = GetAllNamespacesForRM(ctx, "cluster3", "test2")
	require.NoError(t, err)
	require.Equal(t, ns, []string{"test2"})

	// Clean-up
	holder := &workspacev1.Workspace{}
	err = db.SingleDB().QueryProto("delete_workspace", holder, wkspID1)
	require.NoError(t, err)
	err = db.SingleDB().QueryProto("delete_workspace", holder, wkspID2)
	require.NoError(t, err)
}
