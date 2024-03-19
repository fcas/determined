// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/run/v1/run.proto

package runv1

import (
	_ "github.com/determined-ai/determined/proto/pkg/commonv1"
	trialv1 "github.com/determined-ai/determined/proto/pkg/trialv1"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Minimal experiment object
type FlatRunExperiment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the experiment linked to the run.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The type of searcher for the experiment.
	SearcherType string `protobuf:"bytes,2,opt,name=searcher_type,json=searcherType,proto3" json:"searcher_type,omitempty"`
	// The searcher metric name for the experiment.
	SearcherMetric string `protobuf:"bytes,3,opt,name=searcher_metric,json=searcherMetric,proto3" json:"searcher_metric,omitempty"`
	// Original id of a forked or continued experiment.
	ForkedFrom *int32 `protobuf:"varint,4,opt,name=forked_from,json=forkedFrom,proto3,oneof" json:"forked_from,omitempty"`
	// The id of external experiment
	ExternalExperimentId *string `protobuf:"bytes,5,opt,name=external_experiment_id,json=externalExperimentId,proto3,oneof" json:"external_experiment_id,omitempty"`
	// The resource pool the experiment was created in.
	ResourcePool string `protobuf:"bytes,6,opt,name=resource_pool,json=resourcePool,proto3" json:"resource_pool,omitempty"`
	// The current progress of a running experiment.
	Progress float32 `protobuf:"fixed32,7,opt,name=progress,proto3" json:"progress,omitempty"`
	// The description of the experiment.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// The experiment name.
	Name string `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	// Unmanaged experiments are detached.
	Unmanaged bool `protobuf:"varint,10,opt,name=unmanaged,proto3" json:"unmanaged,omitempty"`
	// True if the associated experiment is a multitrial experiment
	IsMultitrial bool `protobuf:"varint,11,opt,name=is_multitrial,json=isMultitrial,proto3" json:"is_multitrial,omitempty"`
}

func (x *FlatRunExperiment) Reset() {
	*x = FlatRunExperiment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_run_v1_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlatRunExperiment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlatRunExperiment) ProtoMessage() {}

func (x *FlatRunExperiment) ProtoReflect() protoreflect.Message {
	mi := &file_determined_run_v1_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlatRunExperiment.ProtoReflect.Descriptor instead.
func (*FlatRunExperiment) Descriptor() ([]byte, []int) {
	return file_determined_run_v1_run_proto_rawDescGZIP(), []int{0}
}

func (x *FlatRunExperiment) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FlatRunExperiment) GetSearcherType() string {
	if x != nil {
		return x.SearcherType
	}
	return ""
}

func (x *FlatRunExperiment) GetSearcherMetric() string {
	if x != nil {
		return x.SearcherMetric
	}
	return ""
}

func (x *FlatRunExperiment) GetForkedFrom() int32 {
	if x != nil && x.ForkedFrom != nil {
		return *x.ForkedFrom
	}
	return 0
}

func (x *FlatRunExperiment) GetExternalExperimentId() string {
	if x != nil && x.ExternalExperimentId != nil {
		return *x.ExternalExperimentId
	}
	return ""
}

func (x *FlatRunExperiment) GetResourcePool() string {
	if x != nil {
		return x.ResourcePool
	}
	return ""
}

func (x *FlatRunExperiment) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *FlatRunExperiment) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *FlatRunExperiment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlatRunExperiment) GetUnmanaged() bool {
	if x != nil {
		return x.Unmanaged
	}
	return false
}

func (x *FlatRunExperiment) GetIsMultitrial() bool {
	if x != nil {
		return x.IsMultitrial
	}
	return false
}

// Flat run respresentation.
type FlatRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the run.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The time the run was started.
	StartTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time the run ended.
	EndTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// The current state of the run(trial).
	State trialv1.State `protobuf:"varint,4,opt,name=state,proto3,enum=determined.trial.v1.State" json:"state,omitempty"`
	// The tags of the associated experiment.
	// TODO(aaron.amanuel): Create add/remove tags for runs.
	Labels []string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty"`
	// The total size of checkpoints.
	CheckpointSize int64 `protobuf:"varint,6,opt,name=checkpoint_size,json=checkpointSize,proto3" json:"checkpoint_size,omitempty"`
	// The count of checkpoints.
	CheckpointCount int32 `protobuf:"varint,7,opt,name=checkpoint_count,json=checkpointCount,proto3" json:"checkpoint_count,omitempty"`
	// Signed searcher metrics value.
	SearcherMetricValue float64 `protobuf:"fixed64,8,opt,name=searcher_metric_value,json=searcherMetricValue,proto3" json:"searcher_metric_value,omitempty"`
	// The id of external run
	ExternalRunId *int32 `protobuf:"varint,9,opt,name=external_run_id,json=externalRunId,proto3,oneof" json:"external_run_id,omitempty"`
	// Trial hyperparameters.
	Hyperparameters *_struct.Struct `protobuf:"bytes,10,opt,name=hyperparameters,proto3,oneof" json:"hyperparameters,omitempty"`
	// summary metrics.
	SummaryMetrics *_struct.Struct `protobuf:"bytes,11,opt,name=summary_metrics,json=summaryMetrics,proto3,oneof" json:"summary_metrics,omitempty"`
	// The id of the user who created the run.
	UserId *int32 `protobuf:"varint,12,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	// Time in seconds which the run ran or has been running.
	Duration *int32 `protobuf:"varint,13,opt,name=duration,proto3,oneof" json:"duration,omitempty"`
	// The id of the project associated with this run.
	ProjectId int32 `protobuf:"varint,14,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The name of the project associated with this run.
	ProjectName string `protobuf:"bytes,15,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	// The id of the workspace associated with this run.
	WorkspaceId int32 `protobuf:"varint,16,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// The name of the workspace associated with this run.
	WorkspaceName string `protobuf:"bytes,17,opt,name=workspace_name,json=workspaceName,proto3" json:"workspace_name,omitempty"`
	// The archived status of the parent project (can be inherited from
	// workspace).
	ParentArchived bool `protobuf:"varint,18,opt,name=parent_archived,json=parentArchived,proto3" json:"parent_archived,omitempty"`
	// Data related the the experiment associated with this run.
	Experiment *FlatRunExperiment `protobuf:"bytes,19,opt,name=experiment,proto3,oneof" json:"experiment,omitempty"`
}

func (x *FlatRun) Reset() {
	*x = FlatRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_run_v1_run_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlatRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlatRun) ProtoMessage() {}

func (x *FlatRun) ProtoReflect() protoreflect.Message {
	mi := &file_determined_run_v1_run_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlatRun.ProtoReflect.Descriptor instead.
func (*FlatRun) Descriptor() ([]byte, []int) {
	return file_determined_run_v1_run_proto_rawDescGZIP(), []int{1}
}

func (x *FlatRun) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FlatRun) GetStartTime() *timestamp.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *FlatRun) GetEndTime() *timestamp.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *FlatRun) GetState() trialv1.State {
	if x != nil {
		return x.State
	}
	return trialv1.State_STATE_UNSPECIFIED
}

func (x *FlatRun) GetLabels() []string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *FlatRun) GetCheckpointSize() int64 {
	if x != nil {
		return x.CheckpointSize
	}
	return 0
}

func (x *FlatRun) GetCheckpointCount() int32 {
	if x != nil {
		return x.CheckpointCount
	}
	return 0
}

func (x *FlatRun) GetSearcherMetricValue() float64 {
	if x != nil {
		return x.SearcherMetricValue
	}
	return 0
}

func (x *FlatRun) GetExternalRunId() int32 {
	if x != nil && x.ExternalRunId != nil {
		return *x.ExternalRunId
	}
	return 0
}

func (x *FlatRun) GetHyperparameters() *_struct.Struct {
	if x != nil {
		return x.Hyperparameters
	}
	return nil
}

func (x *FlatRun) GetSummaryMetrics() *_struct.Struct {
	if x != nil {
		return x.SummaryMetrics
	}
	return nil
}

func (x *FlatRun) GetUserId() int32 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *FlatRun) GetDuration() int32 {
	if x != nil && x.Duration != nil {
		return *x.Duration
	}
	return 0
}

func (x *FlatRun) GetProjectId() int32 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *FlatRun) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *FlatRun) GetWorkspaceId() int32 {
	if x != nil {
		return x.WorkspaceId
	}
	return 0
}

func (x *FlatRun) GetWorkspaceName() string {
	if x != nil {
		return x.WorkspaceName
	}
	return ""
}

func (x *FlatRun) GetParentArchived() bool {
	if x != nil {
		return x.ParentArchived
	}
	return false
}

func (x *FlatRun) GetExperiment() *FlatRunExperiment {
	if x != nil {
		return x.Experiment
	}
	return nil
}

var File_determined_run_v1_run_proto protoreflect.FileDescriptor

var file_determined_run_v1_run_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x72, 0x75, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x74,
	0x72, 0x69, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb1, 0x04, 0x0a, 0x11, 0x46, 0x6c, 0x61, 0x74, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x24, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6f,
	0x72, 0x6b, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x16, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x14, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69,
	0x73, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x74, 0x72, 0x69, 0x61, 0x6c,
	0x3a, 0x78, 0x92, 0x41, 0x75, 0x0a, 0x73, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2, 0x01, 0x0d, 0x69,
	0x73, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x74, 0x72, 0x69, 0x61, 0x6c, 0xd2, 0x01, 0x09, 0x75,
	0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0xd2, 0x01, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0xd2, 0x01, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0xd2, 0x01, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0xd2, 0x01, 0x0f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x66,
	0x6f, 0x72, 0x6b, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0xda, 0x08, 0x0a, 0x07, 0x46, 0x6c, 0x61, 0x74, 0x52, 0x75,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x13, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2b, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x75, 0x6e, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x46, 0x0a, 0x0f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x48, 0x01, 0x52, 0x0f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x0f, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x02, 0x52, 0x0e,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1f, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x64, 0x12, 0x49, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x64, 0x65, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c,
	0x61, 0x74, 0x52, 0x75, 0x6e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x48,
	0x05, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01,
	0x3a, 0x9b, 0x01, 0x92, 0x41, 0x97, 0x01, 0x0a, 0x94, 0x01, 0xd2, 0x01, 0x02, 0x69, 0x64, 0xd2,
	0x01, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0xd2, 0x01, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0xd2, 0x01, 0x04, 0x74, 0x61, 0x67, 0x73, 0xd2, 0x01, 0x0f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0xd2, 0x01, 0x10,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0xd2, 0x01, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0xd2, 0x01, 0x0c,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x0c, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0xd2, 0x01, 0x0e, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x0f, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x75, 0x6e, 0x5f,
	0x69, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2d, 0x61, 0x69, 0x2f, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x72, 0x75, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_determined_run_v1_run_proto_rawDescOnce sync.Once
	file_determined_run_v1_run_proto_rawDescData = file_determined_run_v1_run_proto_rawDesc
)

func file_determined_run_v1_run_proto_rawDescGZIP() []byte {
	file_determined_run_v1_run_proto_rawDescOnce.Do(func() {
		file_determined_run_v1_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_determined_run_v1_run_proto_rawDescData)
	})
	return file_determined_run_v1_run_proto_rawDescData
}

var file_determined_run_v1_run_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_determined_run_v1_run_proto_goTypes = []interface{}{
	(*FlatRunExperiment)(nil),   // 0: determined.run.v1.FlatRunExperiment
	(*FlatRun)(nil),             // 1: determined.run.v1.FlatRun
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(trialv1.State)(0),          // 3: determined.trial.v1.State
	(*_struct.Struct)(nil),      // 4: google.protobuf.Struct
}
var file_determined_run_v1_run_proto_depIdxs = []int32{
	2, // 0: determined.run.v1.FlatRun.start_time:type_name -> google.protobuf.Timestamp
	2, // 1: determined.run.v1.FlatRun.end_time:type_name -> google.protobuf.Timestamp
	3, // 2: determined.run.v1.FlatRun.state:type_name -> determined.trial.v1.State
	4, // 3: determined.run.v1.FlatRun.hyperparameters:type_name -> google.protobuf.Struct
	4, // 4: determined.run.v1.FlatRun.summary_metrics:type_name -> google.protobuf.Struct
	0, // 5: determined.run.v1.FlatRun.experiment:type_name -> determined.run.v1.FlatRunExperiment
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_determined_run_v1_run_proto_init() }
func file_determined_run_v1_run_proto_init() {
	if File_determined_run_v1_run_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_determined_run_v1_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlatRunExperiment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_determined_run_v1_run_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlatRun); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_determined_run_v1_run_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_determined_run_v1_run_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_determined_run_v1_run_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_determined_run_v1_run_proto_goTypes,
		DependencyIndexes: file_determined_run_v1_run_proto_depIdxs,
		MessageInfos:      file_determined_run_v1_run_proto_msgTypes,
	}.Build()
	File_determined_run_v1_run_proto = out.File
	file_determined_run_v1_run_proto_rawDesc = nil
	file_determined_run_v1_run_proto_goTypes = nil
	file_determined_run_v1_run_proto_depIdxs = nil
}
