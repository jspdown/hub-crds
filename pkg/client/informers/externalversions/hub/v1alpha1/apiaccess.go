/*
The GNU AFFERO GENERAL PUBLIC LICENSE

Copyright (c) 2020-2025 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	hubv1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	versioned "github.com/traefik/hub-crds/pkg/client/clientset/versioned"
	internalinterfaces "github.com/traefik/hub-crds/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/traefik/hub-crds/pkg/client/listers/hub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// APIAccessInformer provides access to a shared informer and lister for
// APIAccesses.
type APIAccessInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.APIAccessLister
}

type aPIAccessInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAPIAccessInformer constructs a new informer for APIAccess type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIAccessInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIAccessInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAPIAccessInformer constructs a new informer for APIAccess type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIAccessInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HubV1alpha1().APIAccesses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HubV1alpha1().APIAccesses(namespace).Watch(context.TODO(), options)
			},
		},
		&hubv1alpha1.APIAccess{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIAccessInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIAccessInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *aPIAccessInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hubv1alpha1.APIAccess{}, f.defaultInformer)
}

func (f *aPIAccessInformer) Lister() v1alpha1.APIAccessLister {
	return v1alpha1.NewAPIAccessLister(f.Informer().GetIndexer())
}
