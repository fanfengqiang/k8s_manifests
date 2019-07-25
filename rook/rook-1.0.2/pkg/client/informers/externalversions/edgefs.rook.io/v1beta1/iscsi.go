/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	time "time"

	edgefsrookiov1beta1 "github.com/rook/rook/pkg/apis/edgefs.rook.io/v1beta1"
	versioned "github.com/rook/rook/pkg/client/clientset/versioned"
	internalinterfaces "github.com/rook/rook/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/rook/rook/pkg/client/listers/edgefs.rook.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ISCSIInformer provides access to a shared informer and lister for
// ISCSIs.
type ISCSIInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ISCSILister
}

type iSCSIInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewISCSIInformer constructs a new informer for ISCSI type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewISCSIInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredISCSIInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredISCSIInformer constructs a new informer for ISCSI type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredISCSIInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgefsV1beta1().ISCSIs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EdgefsV1beta1().ISCSIs(namespace).Watch(options)
			},
		},
		&edgefsrookiov1beta1.ISCSI{},
		resyncPeriod,
		indexers,
	)
}

func (f *iSCSIInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredISCSIInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *iSCSIInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&edgefsrookiov1beta1.ISCSI{}, f.defaultInformer)
}

func (f *iSCSIInformer) Lister() v1beta1.ISCSILister {
	return v1beta1.NewISCSILister(f.Informer().GetIndexer())
}
