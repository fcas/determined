//go:build integration
// +build integration

package multirm

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/internal/config"
	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/internal/mocks"
	"github.com/determined-ai/determined/master/internal/rm"
	"github.com/determined-ai/determined/master/internal/rm/rmerrors"
	"github.com/determined-ai/determined/master/internal/sproto"
	"github.com/determined-ai/determined/master/pkg/etc"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/proto/pkg/agentv1"
	"github.com/determined-ai/determined/proto/pkg/apiv1"
	"github.com/determined-ai/determined/proto/pkg/jobv1"
	"github.com/determined-ai/determined/proto/pkg/resourcepoolv1"
)

const rp = "resource-pool"

func TestMain(m *testing.M) {
	pgDB, err := db.ResolveTestPostgres()
	if err != nil {
		log.Panicln(err)
	}

	err = db.MigrateTestPostgres(pgDB, "file://../../../static/migrations", "up")
	if err != nil {
		log.Panicln(err)
	}

	err = etc.SetRootPath("../../../static/srv")
	if err != nil {
		log.Panicln(err)
	}

	os.Exit(m.Run())
}

func TestNewMultiRM(t *testing.T) {
	cases := []struct {
		name string
		cfgs []*config.ResourceManagerWithPoolsConfig
	}{
		{"simple", []*config.ResourceManagerWithPoolsConfig{mockConfig(uuid.NewString())}},
		{"no-name", []*config.ResourceManagerWithPoolsConfig{mockConfig("")}},
		{"multirm", []*config.ResourceManagerWithPoolsConfig{
			mockConfig(uuid.NewString()),
			mockConfig(uuid.NewString()),
			mockConfig(uuid.NewString()),
		}},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMultiRM(db.SingleDB(), echo.New(), tt.cfgs, nil, nil, nil)
			require.NotNil(t, m)
			require.Equal(t, len(tt.cfgs), len(m.rms))

			// If the cfg name is originally "", then NewMultiRM will use
			// config.DefaultRMName.
			if tt.cfgs[0].ResourceManager.Name() == "" {
				require.Equal(t, config.DefaultRMName, m.defaultRMName)
			} else {
				require.Equal(t, tt.cfgs[0].ResourceManager.Name(), m.defaultRMName)
			}
		})
	}
}

func TestGetAllocationSummaries(t *testing.T) {
	cases := []struct {
		name       string
		allocNames []string
		managers   int
	}{
		{"simple", []string{uuid.NewString(), uuid.NewString()}, 1},
		{"multirm", []string{uuid.NewString(), uuid.NewString()}, 3},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			rms := map[string]rm.ResourceManager{}
			for i := 1; i <= tt.managers; i++ {
				ret := map[model.AllocationID]sproto.AllocationSummary{}
				for _, alloc := range tt.allocNames {
					a := alloc + fmt.Sprint(i)
					ret[*model.NewAllocationID(&a)] = sproto.AllocationSummary{}
				}
				require.Equal(t, len(ret), len(tt.allocNames))

				mockRM := mocks.ResourceManager{}
				mockRM.On("GetAllocationSummaries").Return(ret, nil)

				rms[uuid.NewString()] = &mockRM
			}

			m := &MultiRMRouter{rms: rms}

			allocs, err := m.GetAllocationSummaries()
			require.NoError(t, err)
			require.Equal(t, tt.managers*len(tt.allocNames), len(allocs))
			require.NotNil(t, allocs)

			bogus := "bogus"
			require.Empty(t, allocs[*model.NewAllocationID(&bogus)])

			for _, name := range tt.allocNames {
				n := fmt.Sprintf(name + "0")
				tmpName := name

				require.NotNil(t, allocs[*model.NewAllocationID(&n)])
				require.Empty(t, allocs[*model.NewAllocationID(&tmpName)])
			}
		})
	}
}

func TestAllocate(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")

	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	allocReq := sproto.AllocateRequest{ResourceManager: manager}
	mockRM.On("Allocate", manager, allocReq).Return(&sproto.ResourcesSubscription{}, nil)

	res, err := m.Allocate(manager, allocReq)
	require.NoError(t, err)
	require.Equal(t, res, &sproto.ResourcesSubscription{})

	// Check that bogus RM calls get routed to default RM.
	res, err = m.Allocate("bogus", allocReq)
	require.Equal(t, err, nil)
	require.Equal(t, res, &sproto.ResourcesSubscription{})
}

func TestValidateResources(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := sproto.ValidateResourcesRequest{
		ResourcePool: "",
		Slots:        0,
		IsSingleNode: true,
	}

	mockRM.On("ValidateResources", manager, req).Return(nil, nil)

	_, err := m.ValidateResources(manager, req)
	require.NoError(t, err)

	// Check that bogus RM calls get routed to default RM.
	_, err = m.ValidateResources("bogus", req)
	require.NoError(t, err)
}

func TestDeleteJob(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}
	job1 := sproto.DeleteJob{JobID: model.JobID("job1")}

	mockRM.On("DeleteJob", job1).Return(sproto.EmptyDeleteJobResponse(), nil)

	_, err := m.DeleteJob(job1)
	require.NoError(t, err)
}

func TestNotifyContainerRunning(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	mockRM.On("NotifyContainerRunning", manager, sproto.NotifyContainerRunning{}).Return(nil)

	err := m.NotifyContainerRunning(sproto.NotifyContainerRunning{})
	require.Equal(t, err, rmerrors.ErrNotSupported)
}

func TestSetGroupWeight(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req1 := sproto.SetGroupWeight{ResourcePool: "rp1"}

	mockRM.On("SetGroupWeight", manager, req1).Return(nil)

	err := m.SetGroupWeight(manager, req1)
	require.NoError(t, err)

	// Check that bogus RM calls get routed to default RM.
	err = m.SetGroupWeight("bogus", req1)
	require.NoError(t, err)
}

func TestSetGroupPriority(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req1 := sproto.SetGroupPriority{ResourcePool: "rp1"}

	mockRM.On("SetGroupPriority", manager, req1).Return(nil)

	err := m.SetGroupPriority(manager, req1)
	require.NoError(t, err)

	// Check that bogus RM calls get routed to default RM.
	err = m.SetGroupPriority("bogus", req1)
	require.NoError(t, err)
}

func TestExternalPreemptionPending(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	alloc1 := model.AllocationID("alloc1")

	mockRM.On("ExternalPreemptionPending", manager, alloc1).Return(nil)

	err := m.ExternalPreemptionPending(alloc1)
	require.Equal(t, err, rmerrors.ErrNotSupported)
}

func TestIsReattachable(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	mockRM.On("IsReattachableOnlyAfterStarted", mock.Anything).Return(true)

	val := m.IsReattachableOnlyAfterStarted(manager)
	require.Equal(t, true, val)

	// Check that bogus RM calls get routed to default RM.
	val = m.IsReattachableOnlyAfterStarted("bogus")
	require.Equal(t, true, val)
}

func TestGetResourcePools(t *testing.T) {
	cases := []struct {
		name     string
		rpNames  []string
		managers int
	}{
		{"simple", []string{uuid.NewString(), uuid.NewString()}, 1},
		{"multirm", []string{uuid.NewString(), uuid.NewString()}, 5},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			rms := map[string]rm.ResourceManager{}
			for i := 1; i <= tt.managers; i++ {
				ret := []*resourcepoolv1.ResourcePool{}
				for _, n := range tt.rpNames {
					ret = append(ret, &resourcepoolv1.ResourcePool{Name: n})
				}

				mockRM := mocks.ResourceManager{}
				mockRM.On("GetResourcePools").Return(&apiv1.GetResourcePoolsResponse{ResourcePools: ret}, nil)

				rms[uuid.NewString()] = &mockRM
			}

			m := &MultiRMRouter{rms: rms}

			rps, err := m.GetResourcePools()
			require.NoError(t, err)
			require.Equal(t, tt.managers*len(tt.rpNames), len(rps.ResourcePools))

			for _, rp := range rps.ResourcePools {
				require.Contains(t, tt.rpNames, rp.Name)
			}
		})
	}
}

func TestGetDefaultResourcePools(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	res1 := sproto.GetDefaultComputeResourcePoolResponse{PoolName: "default"}
	res2 := sproto.GetDefaultAuxResourcePoolResponse{PoolName: "default"}

	mockRM.On("GetDefaultComputeResourcePool", manager).Return(res1, nil)
	mockRM.On("GetDefaultAuxResourcePool", manager).Return(res2, nil)

	actual1, err := m.GetDefaultComputeResourcePool(manager)
	require.NoError(t, err)
	require.Equal(t, res1, actual1)

	actual2, err := m.GetDefaultAuxResourcePool(manager)
	require.NoError(t, err)
	require.Equal(t, res2, actual2)

	// Check that bogus RM calls get routed to default RM.
	actual1, err = m.GetDefaultComputeResourcePool("bogus")
	require.NoError(t, err)
	require.Equal(t, res1, actual1)

	// Check that bogus RM calls get routed to default RM.
	actual2, err = m.GetDefaultAuxResourcePool("bogus")
	require.NoError(t, err)
	require.Equal(t, res2, actual2)
}

func TestValidateResourcePool(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	mockRM.On("ValidateResourcePool", manager, rp).Return(nil)

	err := m.ValidateResourcePool(manager, rp)
	require.NoError(t, err)

	// Check that bogus RM calls get routed to default RM.
	err = m.ValidateResourcePool("bogus", rp)
	require.NoError(t, err)
}

func TestResolveResourcePool(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := sproto.ResolveResourcesRequest{ResourcePool: rp}

	mockRM.On("ResolveResourcePool", manager, req).Return(manager, rp, nil)

	resolvedRM, resolvedRP, err := m.ResolveResourcePool(manager, req)
	require.NoError(t, err)
	require.Equal(t, manager, resolvedRM)
	require.Equal(t, rp, resolvedRP)

	// Check that bogus RM calls get routed to default RM.
	resolvedRM, resolvedRP, err = m.ResolveResourcePool("bogus", req)
	require.NoError(t, err)
	require.Equal(t, manager, resolvedRM)
	require.Equal(t, rp, resolvedRP)
}

func TestTaskContainerDefaults(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	res := model.TaskContainerDefaultsConfig{}

	mockRM.On("TaskContainerDefaults", manager, rp, res).Return(res, nil)

	_, err := m.TaskContainerDefaults(manager, rp, res)
	require.NoError(t, err)

	// Check that bogus RM calls get routed to default RM.
	_, err = m.TaskContainerDefaults("bogus", rp, res)
	require.NoError(t, err)
}

func TestGetJobQ(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	res := map[model.JobID]*sproto.RMJobInfo{}

	mockRM.On("GetJobQ", manager, rp).Return(res, nil)

	ret, err := m.GetJobQ(manager, rp)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.GetJobQ("bogus", rp)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

func TestGetJobQueueStatsRequest(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := &apiv1.GetJobQueueStatsRequest{ResourcePools: []string{rp}}
	res := &apiv1.GetJobQueueStatsResponse{}

	mockRM.On("GetJobQueueStatsRequest", manager, req).Return(res, nil)

	ret, err := m.GetJobQueueStatsRequest(manager, req)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.GetJobQueueStatsRequest("bogus", req)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

func TestMoveJob(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := sproto.MoveJob{ResourcePool: rp}

	mockRM.On("MoveJob", manager, req).Return(nil)

	err := m.MoveJob(manager, req)
	require.NoError(t, err)

	// Check that bogus RM calls get routed to default RM.
	err = m.MoveJob("bogus", req)
	require.NoError(t, err)
}

func TestGetExternalJobs(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	res := []*jobv1.Job{}

	mockRM.On("GetExternalJobs", manager, rp).Return(res, nil)

	ret, err := m.GetExternalJobs(manager, rp)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.GetExternalJobs("bogus", rp)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

func TestGetAgents(t *testing.T) {
	cases := []struct {
		name       string
		agentNames []string
		managers   int
	}{
		{"simple", []string{uuid.NewString(), uuid.NewString()}, 1},
		{"multirm", []string{uuid.NewString(), uuid.NewString()}, 5},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			rms := map[string]rm.ResourceManager{}
			for i := 1; i <= tt.managers; i++ {
				ret := []*agentv1.Agent{}
				for _, n := range tt.agentNames {
					ret = append(ret, &agentv1.Agent{ResourcePools: []string{n}})
				}

				mockRM := mocks.ResourceManager{}
				mockRM.On("GetAgents").Return(&apiv1.GetAgentsResponse{Agents: ret}, nil)

				rms[uuid.NewString()] = &mockRM
			}

			m := &MultiRMRouter{rms: rms}

			rps, err := m.GetAgents()
			require.NoError(t, err)
			require.Equal(t, tt.managers*len(tt.agentNames), len(rps.Agents))

			for _, rp := range rps.Agents {
				require.Subset(t, tt.agentNames, rp.ResourcePools)
			}
		})
	}
}

func TestGetAgent(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := &apiv1.GetAgentRequest{AgentId: rp, ResourceManager: manager}
	res := &apiv1.GetAgentResponse{}

	mockRM.On("GetAgent", manager, req).Return(res, nil)

	ret, err := m.GetAgent(manager, req)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.GetAgent("bogus", req)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

func TestEnableAgent(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := &apiv1.EnableAgentRequest{AgentId: rp, ResourceManager: manager}
	res := &apiv1.EnableAgentResponse{}

	mockRM.On("EnableAgent", manager, req).Return(res, nil)

	ret, err := m.EnableAgent(manager, req)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.EnableAgent("bogus", req)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

func TestDisableAgent(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := &apiv1.DisableAgentRequest{AgentId: rp, ResourceManager: manager}
	res := &apiv1.DisableAgentResponse{}

	mockRM.On("DisableAgent", manager, req).Return(res, nil)

	ret, err := m.DisableAgent(manager, req)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.DisableAgent("bogus", req)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

func TestGetSlots(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := &apiv1.GetSlotsRequest{AgentId: rp, ResourceManager: manager}
	res := &apiv1.GetSlotsResponse{}

	mockRM.On("GetSlots", manager, req).Return(res, nil)

	ret, err := m.GetSlots(manager, req)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.GetSlots("bogus", req)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

func TestGetSlot(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := &apiv1.GetSlotRequest{AgentId: rp, ResourceManager: manager}
	res := &apiv1.GetSlotResponse{}

	mockRM.On("GetSlot", manager, req).Return(res, nil)

	ret, err := m.GetSlot(manager, req)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.GetSlot("bogus", req)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

func TestEnableSlot(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := &apiv1.EnableSlotRequest{AgentId: rp, ResourceManager: manager}
	res := &apiv1.EnableSlotResponse{}

	mockRM.On("EnableSlot", manager, req).Return(res, nil)

	ret, err := m.EnableSlot(manager, req)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.EnableSlot("bogus", req)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

func TestDisableSlot(t *testing.T) {
	mockRM := mocks.ResourceManager{}
	manager := uuid.NewString()
	log := logrus.WithField("component", "resource-router")
	m := &MultiRMRouter{
		defaultRMName: manager,
		rms:           map[string]rm.ResourceManager{manager: &mockRM},
		syslog:        log,
	}

	req := &apiv1.DisableSlotRequest{AgentId: rp, ResourceManager: manager}
	res := &apiv1.DisableSlotResponse{}

	mockRM.On("DisableSlot", manager, req).Return(res, nil)

	ret, err := m.DisableSlot(manager, req)
	require.NoError(t, err)
	require.Equal(t, ret, res)

	// Check that bogus RM calls get routed to default RM.
	ret, err = m.DisableSlot("bogus", req)
	require.NoError(t, err)
	require.Equal(t, ret, res)
}

// Only returns AgentRM for testing purposes.
func mockConfig(rmName string) *config.ResourceManagerWithPoolsConfig {
	return &config.ResourceManagerWithPoolsConfig{
		ResourceManager: &config.ResourceManagerConfig{
			AgentRM: &config.AgentResourceManagerConfig{
				Name: rmName,
			},
		},
	}
}
