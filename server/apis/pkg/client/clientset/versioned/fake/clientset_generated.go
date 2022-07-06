/*
Copyright 2021 The Volcano Authors.

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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "server/apis/pkg/client/clientset/versioned"
	batchv1alpha1 "server/apis/pkg/client/clientset/versioned/typed/batch/v1alpha1"
	fakebatchv1alpha1 "server/apis/pkg/client/clientset/versioned/typed/batch/v1alpha1/fake"
	busv1alpha1 "server/apis/pkg/client/clientset/versioned/typed/bus/v1alpha1"
	fakebusv1alpha1 "server/apis/pkg/client/clientset/versioned/typed/bus/v1alpha1/fake"
	flowv1alpha1 "server/apis/pkg/client/clientset/versioned/typed/flow/v1alpha1"
	fakeflowv1alpha1 "server/apis/pkg/client/clientset/versioned/typed/flow/v1alpha1/fake"
	nodeinfov1alpha1 "server/apis/pkg/client/clientset/versioned/typed/nodeinfo/v1alpha1"
	fakenodeinfov1alpha1 "server/apis/pkg/client/clientset/versioned/typed/nodeinfo/v1alpha1/fake"
	schedulingv1beta1 "server/apis/pkg/client/clientset/versioned/typed/scheduling/v1beta1"
	fakeschedulingv1beta1 "server/apis/pkg/client/clientset/versioned/typed/scheduling/v1beta1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// BatchV1alpha1 retrieves the BatchV1alpha1Client
func (c *Clientset) BatchV1alpha1() batchv1alpha1.BatchV1alpha1Interface {
	return &fakebatchv1alpha1.FakeBatchV1alpha1{Fake: &c.Fake}
}

// BusV1alpha1 retrieves the BusV1alpha1Client
func (c *Clientset) BusV1alpha1() busv1alpha1.BusV1alpha1Interface {
	return &fakebusv1alpha1.FakeBusV1alpha1{Fake: &c.Fake}
}

// FlowV1alpha1 retrieves the FlowV1alpha1Client
func (c *Clientset) FlowV1alpha1() flowv1alpha1.FlowV1alpha1Interface {
	return &fakeflowv1alpha1.FakeFlowV1alpha1{Fake: &c.Fake}
}

// NodeinfoV1alpha1 retrieves the NodeinfoV1alpha1Client
func (c *Clientset) NodeinfoV1alpha1() nodeinfov1alpha1.NodeinfoV1alpha1Interface {
	return &fakenodeinfov1alpha1.FakeNodeinfoV1alpha1{Fake: &c.Fake}
}

// SchedulingV1beta1 retrieves the SchedulingV1beta1Client
func (c *Clientset) SchedulingV1beta1() schedulingv1beta1.SchedulingV1beta1Interface {
	return &fakeschedulingv1beta1.FakeSchedulingV1beta1{Fake: &c.Fake}
}
