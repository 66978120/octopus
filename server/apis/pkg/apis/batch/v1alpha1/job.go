package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

// JobPhase defines the phase of the job.
type JobPhase string

const (
	// Pending is the phase that job is pending in the queue, waiting for scheduling decision
	Pending JobPhase = "Pending"
	// Aborting is the phase that job is aborted, waiting for releasing pods
	Aborting JobPhase = "Aborting"
	// Aborted is the phase that job is aborted by user or error handling
	Aborted JobPhase = "Aborted"
	// Running is the phase that minimal available tasks of Job are running
	Running JobPhase = "Running"
	// Restarting is the phase that the Job is restarted, waiting for pod releasing and recreating
	Restarting JobPhase = "Restarting"
	// Completing is the phase that required tasks of job are completed, job starts to clean up
	Completing JobPhase = "Completing"
	// Completed is the phase that all tasks of Job are completed
	Completed JobPhase = "Completed"
	// Terminating is the phase that the Job is terminated, waiting for releasing pods
	Terminating JobPhase = "Terminating"
	// Terminated is the phase that the job is finished unexpected, e.g. events
	Terminated JobPhase = "Terminated"
	// Failed is the phase that the job is restarted failed reached the maximum number of retries.
	Failed JobPhase = "Failed"
	// Succeeded is ..
	Succeeded JobPhase = "Succeeded"
	// Unknown is ..
	Unknown JobPhase = "Unknown"
)

type Job struct {
	v1alpha1.Job

	Spec   JobSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status JobStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// JobStatus represents the current status of a Job.
type JobStatus struct {
	v1alpha1.JobStatus

	TaskRoleStatus []TaskRoleStatus `json:"roleStatus,omitempty" protobuf:"bytes,12,opt,name=roleStatus"`

	CreatedAt metav1.Time `json:"createdAt,omitempty" protobuf:"bytes,13,opt,name=createdAt"` //任务创建时间

	StartAt *metav1.Time `json:"startAt,omitempty" protobuf:"bytes,14,opt,name=startAt"` //任务开始时间

	FinishAt *metav1.Time `json:"finishAt,omitempty" protobuf:"bytes,15,opt,name=finishAt"` //任务结束时
}

// JobSpec describes how the job execution will look like and when it will actually run.
type JobSpec struct {
	v1alpha1.JobSpec

	Tasks []TaskSpec `json:"tasks,omitempty" protobuf:"bytes,4,opt,name=tasks"`
}

// TaskSpec specifies the task specification of Job.
type TaskSpec struct {
	v1alpha1.TaskSpec

	CompletionPolicy CompletionPolicy `json:"completionPolicy,omitempty" protobuf:"bytes,4,opt,name=completionpolicy"`
}

// TaskRoleStatus record the status of a task role
type TaskRoleStatus struct {
	Name            string          `json:"name"`
	Phase           string          `json:"phase"`
	PhaseMessage    string          `json:"phaseMessage"`
	TransitionTime  metav1.Time     `json:"transitionTime"`
	State           string          `json:"state"`
	ReplicaStatuses []ReplicaStatus `json:"replicaStatus"`
}

// ReplicaStatus record the  status of a replica
type ReplicaStatus struct {
	Index             uint                     `json:"index"`
	Name              string                   `json:"name"`
	Phase             string                   `json:"phase"`
	PhaseMessage      string                   `json:"phaseMessage"`
	Stopped           bool                     `json:"stopped"`
	TransitionTime    metav1.Time              `json:"transitionTime"`
	StartAt           *metav1.Time             `json:"startAt"`
	FinishAt          *metav1.Time             `json:"finishAt"`
	TotalRetriedCount uint                     `json:"totalRetriedCount"`
	PodName           string                   `json:"podName"`
	PodReason         string                   `json:"podReason"`
	PodUID            *types.UID               `json:"podUID"`
	PodIP             string                   `json:"podIP"`
	PodHostIP         string                   `json:"podHostIP"`
	ContainerName     string                   `json:"containerName"`
	ContainerID       string                   `json:"containerID"`
	TerminatedInfo    *ContainerTerminatedInfo `json:"terminatedInfo"`
}

//ContainerTerminatedInfo contains the terminated information
type ContainerTerminatedInfo struct {
	ExitCode    int32  `json:"exitCode"`
	ExitMessage string `json:"exitMessage"`
	Signal      int32  `json:"signal"`
	Reason      string `json:"reason"`
}

// CompletionPolicy declare the condition of completion
type CompletionPolicy struct {
	MaxFailed    int32 `json:"maxFailed"`
	MinSucceeded int32 `json:"minSucceeded"`
}
