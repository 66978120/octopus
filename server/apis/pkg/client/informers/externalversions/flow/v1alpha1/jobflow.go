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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	flowv1alpha1 "volcano.sh/apis/pkg/apis/flow/v1alpha1"
	versioned "volcano.sh/apis/pkg/client/clientset/versioned"
	internalinterfaces "volcano.sh/apis/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "volcano.sh/apis/pkg/client/listers/flow/v1alpha1"
)

// JobFlowInformer provides access to a shared informer and lister for
// JobFlows.
type JobFlowInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.JobFlowLister
}

type jobFlowInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewJobFlowInformer constructs a new informer for JobFlow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewJobFlowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredJobFlowInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredJobFlowInformer constructs a new informer for JobFlow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredJobFlowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowV1alpha1().JobFlows(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlowV1alpha1().JobFlows(namespace).Watch(context.TODO(), options)
			},
		},
		&flowv1alpha1.JobFlow{},
		resyncPeriod,
		indexers,
	)
}

func (f *jobFlowInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredJobFlowInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *jobFlowInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&flowv1alpha1.JobFlow{}, f.defaultInformer)
}

func (f *jobFlowInformer) Lister() v1alpha1.JobFlowLister {
	return v1alpha1.NewJobFlowLister(f.Informer().GetIndexer())
}
