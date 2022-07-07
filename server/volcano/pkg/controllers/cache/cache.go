/*
Copyright 2019 The Volcano Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cache

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"golang.org/x/time/rate"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"

	typeJob "server/apis/pkg/apis/batch/v1alpha1"
	"server/volcano/pkg/controllers/job/state"

	typeApis "server/volcano/pkg/controllers/apis"

	"volcano.sh/apis/pkg/apis/batch/v1alpha1"
	"volcano.sh/volcano/pkg/controllers/apis"
)

type jobCache struct {
	sync.Mutex

	jobs        map[string]*typeApis.JobInfo
	deletedJobs workqueue.RateLimitingInterface
}

func keyFn(ns, name string) string {
	return fmt.Sprintf("%s/%s", ns, name)
}

//JobKeyByName gets the key for the job name.
func JobKeyByName(namespace string, name string) string {
	return keyFn(namespace, name)
}

//JobKeyByReq gets the key for the job request.
func JobKeyByReq(req *apis.Request) string {
	return keyFn(req.Namespace, req.JobName)
}

//JobKey gets the "ns"/"name" format of the given job.
func JobKey(job *typeJob.Job) string {
	return keyFn(job.Namespace, job.Name)
}

func jobTerminated(job *typeApis.JobInfo) bool {
	return job.Job == nil && len(job.Pods) == 0
}

func jobKeyOfPod(pod *v1.Pod) (string, error) {
	jobName, found := pod.Annotations[v1alpha1.JobNameKey]
	if !found {
		return "", fmt.Errorf("failed to find job name of pod <%s/%s>",
			pod.Namespace, pod.Name)
	}

	return keyFn(pod.Namespace, jobName), nil
}

// New gets the job Cache.
func New() Cache {
	queue := workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(5*time.Millisecond, 180*time.Second),
		// 10 qps, 100 bucket size.  This is only for retry speed and its only the overall factor (not per item)
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
	)

	return &jobCache{
		jobs:        map[string]*typeApis.JobInfo{},
		deletedJobs: workqueue.NewRateLimitingQueue(queue),
	}
}

func (jc *jobCache) Get(key string) (*typeApis.JobInfo, error) {
	jc.Lock()
	defer jc.Unlock()

	job, found := jc.jobs[key]
	if !found {
		return nil, fmt.Errorf("failed to find job <%s>", key)
	}

	if job.Job == nil {
		return nil, fmt.Errorf("job <%s> is not ready", key)
	}

	return job.Clone(), nil
}

func (jc *jobCache) GetStatus(key string) (*typeJob.JobStatus, error) {
	jc.Lock()
	defer jc.Unlock()

	job, found := jc.jobs[key]
	if !found {
		return nil, fmt.Errorf("failed to find job <%s>", key)
	}

	if job.Job == nil {
		return nil, fmt.Errorf("job <%s> is not ready", key)
	}

	status := job.Job.Status

	return &status, nil
}

func (jc *jobCache) Add(job *typeJob.Job) error {
	jc.Lock()
	defer jc.Unlock()
	key := JobKey(job)
	if jobInfo, found := jc.jobs[key]; found {
		if jobInfo.Job == nil {
			jobInfo.SetJob(job)

			return nil
		}
		return fmt.Errorf("duplicated jobInfo <%v>", key)
	}

	jc.jobs[key] = &typeApis.JobInfo{}
	jc.jobs[key].Name = job.Name
	jc.jobs[key].Namespace = job.Namespace
	jc.jobs[key].Job = job
	jc.jobs[key].Pods = make(map[string]map[string]*v1.Pod)

	return nil
}

func (jc *jobCache) Update(obj *typeJob.Job) error {
	jc.Lock()
	defer jc.Unlock()

	key := JobKey(obj)
	job, found := jc.jobs[key]
	if !found {
		return fmt.Errorf("failed to find job <%v>", key)
	}

	var oldResourceversion, newResourceversion uint64
	var err error
	if oldResourceversion, err = strconv.ParseUint(job.Job.ResourceVersion, 10, 64); err != nil {
		return fmt.Errorf("failed to parase job <%v> resource version <%s>", key, job.Job.ResourceVersion)
	}

	if newResourceversion, err = strconv.ParseUint(obj.ResourceVersion, 10, 64); err != nil {
		return fmt.Errorf("failed to parase job <%v> resource version <%s>", key, obj.ResourceVersion)
	}
	if newResourceversion < oldResourceversion {
		return fmt.Errorf("job <%v> has too old resource version: %d (%d)", key, newResourceversion, oldResourceversion)
	}
	job.Job = obj
	return nil
}

func (jc *jobCache) Delete(obj *typeJob.Job) error {
	jc.Lock()
	defer jc.Unlock()

	key := JobKey(obj)
	jobInfo, found := jc.jobs[key]
	if !found {
		return fmt.Errorf("failed to find job <%v>", key)
	}
	jobInfo.Job = nil
	jc.deleteJob(jobInfo)

	return nil
}

func (jc *jobCache) AddPod(pod *v1.Pod) error {
	jc.Lock()
	defer jc.Unlock()

	key, err := jobKeyOfPod(pod)
	if err != nil {
		return err
	}

	job, found := jc.jobs[key]
	if !found {
		job = &typeApis.JobInfo{}
		job.Pods = make(map[string]map[string]*v1.Pod)
		jc.jobs[key] = job
	}

	return job.AddPod(pod)
}

func (jc *jobCache) UpdatePod(pod *v1.Pod) error {
	jc.Lock()
	defer jc.Unlock()

	key, err := jobKeyOfPod(pod)
	if err != nil {
		return err
	}

	job, found := jc.jobs[key]
	if !found {
		job = &typeApis.JobInfo{}
		job.Pods = make(map[string]map[string]*v1.Pod)
		jc.jobs[key] = job
	}

	return job.UpdatePod(pod)
}

func (jc *jobCache) DeletePod(pod *v1.Pod) error {
	jc.Lock()
	defer jc.Unlock()

	key, err := jobKeyOfPod(pod)
	if err != nil {
		return err
	}

	job, found := jc.jobs[key]
	if !found {
		job = &typeApis.JobInfo{}
		job.Pods = make(map[string]map[string]*v1.Pod)
		jc.jobs[key] = job
	}

	if err := job.DeletePod(pod); err != nil {
		return err
	}

	if jc.jobs[key].Job == nil {
		jc.deleteJob(job)
	}

	return nil
}

func (jc *jobCache) Run(stopCh <-chan struct{}) {
	wait.Until(jc.worker, 0, stopCh)
}

func (jc *jobCache) TaskCompleted(jobKey, taskName string) bool {
	jc.Lock()
	defer jc.Unlock()

	var isCompleted, isSuccess bool

	jobInfo, found := jc.jobs[jobKey]
	if !found {
		return false
	}

	taskPods, found := jobInfo.Pods[taskName]

	if !found {
		return false
	}

	if jobInfo.Job == nil {
		return false
	}

	completionPolicy := state.GetTaskCompletionPolicy(jobInfo.Job, taskName)

	taskStatus := state.GetTaskStatus(jobInfo.Job, taskName)

	var (
		failed    int32 = 0
		succeeded int32 = 0
	)

	for _, pod := range taskPods {
		if pod.Status.Phase == v1.PodSucceeded {
			succeeded++
		}
		if pod.Status.Phase == v1.PodFailed {
			failed++
		}
	}

	if completionPolicy.MinSucceeded <= succeeded {
		isCompleted, isSuccess = true, true
	}

	if completionPolicy.MaxFailed <= failed {
		isCompleted, isSuccess = true, false
	}

	if isCompleted {
		if isSuccess {
			taskStatus.Phase = string(typeJob.Completed)
		} else {
			taskStatus.Phase = string(typeJob.Failed)
		}
		state.StopReplicas(taskStatus)
		return true
	}

	return false
}

func (jc *jobCache) TaskFailed(jobKey, taskName string) bool {
	jc.Lock()
	defer jc.Unlock()

	var taskReplicas, retried, maxRetry int32

	jobInfo, found := jc.jobs[jobKey]
	if !found {
		return false
	}

	taskPods, found := jobInfo.Pods[taskName]

	if !found || jobInfo.Job == nil {
		return false
	}

	for _, task := range jobInfo.Job.Spec.Tasks {
		if task.Name == taskName {
			maxRetry = task.MaxRetry
			taskReplicas = task.Replicas
			break
		}
	}

	// maxRetry == -1 means no limit
	if taskReplicas == 0 || maxRetry == -1 {
		return false
	}

	// Compatible with existing job
	if maxRetry == 0 {
		maxRetry = 3
	}

	for _, pod := range taskPods {
		if pod.Status.Phase == v1.PodRunning || pod.Status.Phase == v1.PodPending {
			for j := range pod.Status.InitContainerStatuses {
				stat := pod.Status.InitContainerStatuses[j]
				retried += stat.RestartCount
			}
			for j := range pod.Status.ContainerStatuses {
				stat := pod.Status.ContainerStatuses[j]
				retried += stat.RestartCount
			}
		}
	}
	return retried >= maxRetry
}

func (jc *jobCache) worker() {
	for jc.processCleanupJob() {
	}
}

func (jc *jobCache) processCleanupJob() bool {
	obj, shutdown := jc.deletedJobs.Get()
	if shutdown {
		return false
	}
	defer jc.deletedJobs.Done(obj)

	job, ok := obj.(*typeApis.JobInfo)
	if !ok {
		klog.Errorf("failed to convert %v to *apis.JobInfo", obj)
		return true
	}

	jc.Mutex.Lock()
	defer jc.Mutex.Unlock()

	if jobTerminated(job) {
		jc.deletedJobs.Forget(obj)
		key := keyFn(job.Namespace, job.Name)
		delete(jc.jobs, key)
		klog.V(3).Infof("Job <%s> was deleted.", key)
	} else {
		// Retry
		jc.deleteJob(job)
	}
	return true
}

func (jc *jobCache) deleteJob(job *typeApis.JobInfo) {
	klog.V(3).Infof("Try to delete Job <%v/%v>",
		job.Namespace, job.Name)

	jc.deletedJobs.AddRateLimited(job)
}
