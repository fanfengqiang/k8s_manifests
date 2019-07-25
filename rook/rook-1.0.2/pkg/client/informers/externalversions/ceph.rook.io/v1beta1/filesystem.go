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

	cephrookiov1beta1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1beta1"
	versioned "github.com/rook/rook/pkg/client/clientset/versioned"
	internalinterfaces "github.com/rook/rook/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/rook/rook/pkg/client/listers/ceph.rook.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FilesystemInformer provides access to a shared informer and lister for
// Filesystems.
type FilesystemInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.FilesystemLister
}

type filesystemInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFilesystemInformer constructs a new informer for Filesystem type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilesystemInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFilesystemInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFilesystemInformer constructs a new informer for Filesystem type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFilesystemInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CephV1beta1().Filesystems(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CephV1beta1().Filesystems(namespace).Watch(options)
			},
		},
		&cephrookiov1beta1.Filesystem{},
		resyncPeriod,
		indexers,
	)
}

func (f *filesystemInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFilesystemInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *filesystemInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cephrookiov1beta1.Filesystem{}, f.defaultInformer)
}

func (f *filesystemInformer) Lister() v1beta1.FilesystemLister {
	return v1beta1.NewFilesystemLister(f.Informer().GetIndexer())
}
