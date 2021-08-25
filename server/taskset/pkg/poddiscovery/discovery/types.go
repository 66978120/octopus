// MIT License
//
// Copyright (c) PCL. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE
//

package discovery

import (
	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
	vcclientset "volcano.sh/volcano/pkg/client/clientset/versioned"
	electionClient "scheduler/pkg/poddiscovery/leaderelection"
	leaselock "scheduler/pkg/poddiscovery/leaderelection/leaselock"

	"k8s.io/client-go/rest"
)

type Discovery struct {
	config       *rest.Config
	coClientBeta *electionClient.CoordinationClient
	coClientV1   *electionClient.CoordinationClient
	tsClient     *vcclientset.Clientset
	leaseLock    *leaselock.LeaseLock
	identity     string
	namespace    string
	taskset      string
	role         string
	localIP      string
	replicaIndex string
	cache        *typeJob.Job
	ipPairs      map[string]string
	completed    bool
}

const (
	//HOSTS_JSON_FILE is the temp file for all pods's address
	HOSTS_JSON_FILE = "/etc/hosts_json.json"
	//HOSTS_FILE is the hosts file
	HOSTS_FILE = "/etc/hosts"

	//EnvNameNamespace define which namespace the taskset is running
	EnvNameNamespace = "TASKSET_NAMESPACE"
	// EnvNameTaskSetName define the key of TaskSet name which will be injected in  the pod
	EnvNameTaskSetName = "TASKSET_NAME"
	// EnvNameTaskRoleName define the key of TaskRole name which will be injected in the pod
	EnvNameTaskRoleName = "TASKROLE_NAME"
	// EnvNameTaskRoleReplicaIndex define the key of TaskRole replica index
	EnvNameTaskRoleReplicaIndex = "TASKROLE_REPLICA_INDEX"
)
