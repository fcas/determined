// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// source: determined/run/v1/run.proto

package runv1

import (
	_ "github.com/determined-ai/determined/proto/pkg/commonv1"
	trialv1 "github.com/determined-ai/determined/proto/pkg/trialv1"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
	// The type of searcher for the experiment.
	SearcherType *string `protobuf:"bytes,8,opt,name=searcher_type,json=searcherType,proto3,oneof" json:"searcher_type,omitempty"`
	// The searcher metric name for the experiment.
	SearcherMetric *string `protobuf:"bytes,9,opt,name=searcher_metric,json=searcherMetric,proto3,oneof" json:"searcher_metric,omitempty"`
	// Signed searcher metrics value.
	SearcherMetricValue float64 `protobuf:"fixed64,10,opt,name=searcher_metric_value,json=searcherMetricValue,proto3" json:"searcher_metric_value,omitempty"`
	// The id of external run
	ExternalRunId *int32 `protobuf:"varint,11,opt,name=external_run_id,json=externalRunId,proto3,oneof" json:"external_run_id,omitempty"`
	// Trial hyperparameters.
	Hyperparameters string `protobuf:"bytes,12,opt,name=hyperparameters,proto3" json:"hyperparameters,omitempty"`
	// summary metrics.
	SummaryMetrics *_struct.Struct `protobuf:"bytes,13,opt,name=summary_metrics,json=summaryMetrics,proto3,oneof" json:"summary_metrics,omitempty"`
	// The id of the experiment linked to the run.
	ExperimentId *int32 `protobuf:"varint,14,opt,name=experiment_id,json=experimentId,proto3,oneof" json:"experiment_id,omitempty"`
	// The id of the user who created the run.
	UserId *int32 `protobuf:"varint,15,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	// Original id of a forked or continued experiment.
	ForkedFrom *wrappers.Int32Value `protobuf:"bytes,16,opt,name=forked_from,json=forkedFrom,proto3,oneof" json:"forked_from,omitempty"`
	// The id of external experiment
	ExternalExperimentId *string `protobuf:"bytes,17,opt,name=external_experiment_id,json=externalExperimentId,proto3,oneof" json:"external_experiment_id,omitempty"`
	// The resource pool the experiment was created in.
	ResourcePool *string `protobuf:"bytes,18,opt,name=resource_pool,json=resourcePool,proto3,oneof" json:"resource_pool,omitempty"`
	// The current progress of a running experiment.
	ExperimentProgress *float32 `protobuf:"fixed32,19,opt,name=experiment_progress,json=experimentProgress,proto3,oneof" json:"experiment_progress,omitempty"`
	// The description of the experiment.
	ExperimentDescription *string `protobuf:"bytes,20,opt,name=experiment_description,json=experimentDescription,proto3,oneof" json:"experiment_description,omitempty"`
	// The experiment name.
	ExperimentName *string `protobuf:"bytes,21,opt,name=experiment_name,json=experimentName,proto3,oneof" json:"experiment_name,omitempty"`
	// Time in seconds which experiment ran or has been running.
	Duration *int32 `protobuf:"varint,22,opt,name=duration,proto3,oneof" json:"duration,omitempty"`
	// The id of the project associated with this experiment.
	ProjectId int32 `protobuf:"varint,23,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The name of the project associated with this experiment.
	ProjectName string `protobuf:"bytes,24,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	// The id of the workspace associated with this experiment.
	WorkspaceId int32 `protobuf:"varint,25,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// The name of the workspace associated with this experiment.
	WorkspaceName string `protobuf:"bytes,26,opt,name=workspace_name,json=workspaceName,proto3" json:"workspace_name,omitempty"`
	// The archived status of the parent project (can be inherited from
	// workspace).
	ParentArchived *bool `protobuf:"varint,27,opt,name=parent_archived,json=parentArchived,proto3,oneof" json:"parent_archived,omitempty"`
	// Unmanaged experiments are detached.
	Unmanaged *bool `protobuf:"varint,28,opt,name=unmanaged,proto3,oneof" json:"unmanaged,omitempty"`
	// True if the associated experiment is a multitrial experiment
	IsExpMultitrial bool `protobuf:"varint,29,opt,name=is_exp_multitrial,json=isExpMultitrial,proto3" json:"is_exp_multitrial,omitempty"`
}

func (x *FlatRun) Reset() {
	*x = FlatRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_determined_run_v1_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlatRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlatRun) ProtoMessage() {}

func (x *FlatRun) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FlatRun.ProtoReflect.Descriptor instead.
func (*FlatRun) Descriptor() ([]byte, []int) {
	return file_determined_run_v1_run_proto_rawDescGZIP(), []int{0}
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

func (x *FlatRun) GetSearcherType() string {
	if x != nil && x.SearcherType != nil {
		return *x.SearcherType
	}
	return ""
}

func (x *FlatRun) GetSearcherMetric() string {
	if x != nil && x.SearcherMetric != nil {
		return *x.SearcherMetric
	}
	return ""
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

func (x *FlatRun) GetHyperparameters() string {
	if x != nil {
		return x.Hyperparameters
	}
	return ""
}

func (x *FlatRun) GetSummaryMetrics() *_struct.Struct {
	if x != nil {
		return x.SummaryMetrics
	}
	return nil
}

func (x *FlatRun) GetExperimentId() int32 {
	if x != nil && x.ExperimentId != nil {
		return *x.ExperimentId
	}
	return 0
}

func (x *FlatRun) GetUserId() int32 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *FlatRun) GetForkedFrom() *wrappers.Int32Value {
	if x != nil {
		return x.ForkedFrom
	}
	return nil
}

func (x *FlatRun) GetExternalExperimentId() string {
	if x != nil && x.ExternalExperimentId != nil {
		return *x.ExternalExperimentId
	}
	return ""
}

func (x *FlatRun) GetResourcePool() string {
	if x != nil && x.ResourcePool != nil {
		return *x.ResourcePool
	}
	return ""
}

func (x *FlatRun) GetExperimentProgress() float32 {
	if x != nil && x.ExperimentProgress != nil {
		return *x.ExperimentProgress
	}
	return 0
}

func (x *FlatRun) GetExperimentDescription() string {
	if x != nil && x.ExperimentDescription != nil {
		return *x.ExperimentDescription
	}
	return ""
}

func (x *FlatRun) GetExperimentName() string {
	if x != nil && x.ExperimentName != nil {
		return *x.ExperimentName
	}
	return ""
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
	if x != nil && x.ParentArchived != nil {
		return *x.ParentArchived
	}
	return false
}

func (x *FlatRun) GetUnmanaged() bool {
	if x != nil && x.Unmanaged != nil {
		return *x.Unmanaged
	}
	return false
}

func (x *FlatRun) GetIsExpMultitrial() bool {
	if x != nil {
		return x.IsExpMultitrial
	}
	return false
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
	0x6f, 0x22, 0xcc, 0x0d, 0x0a, 0x07, 0x46, 0x6c, 0x61, 0x74, 0x52, 0x75, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2e, 0x74, 0x72, 0x69, 0x61,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a,
	0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65,
	0x72, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2b, 0x0a, 0x0f, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x02, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x75,
	0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x70,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x45, 0x0a, 0x0f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x48, 0x03, 0x52, 0x0e, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x48, 0x04,
	0x52, 0x0c, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x41, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x48, 0x06, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x88,
	0x01, 0x01, 0x12, 0x39, 0x0a, 0x16, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x07, 0x52, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x09, 0x52, 0x12, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a,
	0x16, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52,
	0x15, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x0b, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x48, 0x0c, 0x52, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0d, 0x52,
	0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x18,
	0x1c, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0e, 0x52, 0x09, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x5f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x69, 0x73, 0x45, 0x78, 0x70, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x74, 0x72, 0x69, 0x61,
	0x6c, 0x3a, 0x9d, 0x01, 0x92, 0x41, 0x99, 0x01, 0x0a, 0x96, 0x01, 0xd2, 0x01, 0x02, 0x69, 0x64,
	0xd2, 0x01, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0xd2, 0x01, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0xd2, 0x01, 0x04, 0x74, 0x61, 0x67, 0x73, 0xd2, 0x01, 0x0f, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0xd2, 0x01,
	0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0xd2, 0x01, 0x11, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0xd2, 0x01, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0xd2, 0x01, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0xd2, 0x01, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0xd2, 0x01, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x65, 0x72,
	0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x5f,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x66, 0x6f, 0x72, 0x6b, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x19, 0x0a,
	0x17, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72,
	0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x75, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2d, 0x61, 0x69, 0x2f, 0x64, 0x65, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x72, 0x75, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_determined_run_v1_run_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_determined_run_v1_run_proto_goTypes = []interface{}{
	(*FlatRun)(nil),             // 0: determined.run.v1.FlatRun
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(trialv1.State)(0),          // 2: determined.trial.v1.State
	(*_struct.Struct)(nil),      // 3: google.protobuf.Struct
	(*wrappers.Int32Value)(nil), // 4: google.protobuf.Int32Value
}
var file_determined_run_v1_run_proto_depIdxs = []int32{
	1, // 0: determined.run.v1.FlatRun.start_time:type_name -> google.protobuf.Timestamp
	1, // 1: determined.run.v1.FlatRun.end_time:type_name -> google.protobuf.Timestamp
	2, // 2: determined.run.v1.FlatRun.state:type_name -> determined.trial.v1.State
	3, // 3: determined.run.v1.FlatRun.summary_metrics:type_name -> google.protobuf.Struct
	4, // 4: determined.run.v1.FlatRun.forked_from:type_name -> google.protobuf.Int32Value
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_determined_run_v1_run_proto_init() }
func file_determined_run_v1_run_proto_init() {
	if File_determined_run_v1_run_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_determined_run_v1_run_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_determined_run_v1_run_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
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
