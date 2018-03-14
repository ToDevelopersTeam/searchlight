/*
Copyright 2018 The Searchlight Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	time "time"

	monitoring "github.com/appscode/searchlight/apis/monitoring"
	clientset_internalversion "github.com/appscode/searchlight/client/clientset/internalversion"
	internalinterfaces "github.com/appscode/searchlight/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/appscode/searchlight/client/listers/monitoring/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IncidentInformer provides access to a shared informer and lister for
// Incidents.
type IncidentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.IncidentLister
}

type incidentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIncidentInformer constructs a new informer for Incident type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIncidentInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIncidentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIncidentInformer constructs a new informer for Incident type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIncidentInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Monitoring().Incidents(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Monitoring().Incidents(namespace).Watch(options)
			},
		},
		&monitoring.Incident{},
		resyncPeriod,
		indexers,
	)
}

func (f *incidentInformer) defaultInformer(client clientset_internalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIncidentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *incidentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoring.Incident{}, f.defaultInformer)
}

func (f *incidentInformer) Lister() internalversion.IncidentLister {
	return internalversion.NewIncidentLister(f.Informer().GetIndexer())
}