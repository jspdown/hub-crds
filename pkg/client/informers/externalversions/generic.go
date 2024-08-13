/*
The GNU AFFERO GENERAL PUBLIC LICENSE

Copyright (c) 2020-2024 Traefik Labs

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

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=hub.traefik.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("apis"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hub().V1alpha1().APIs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("apiaccesses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hub().V1alpha1().APIAccesses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("apiplans"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hub().V1alpha1().APIPlans().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("apiportals"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hub().V1alpha1().APIPortals().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("apiratelimits"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hub().V1alpha1().APIRateLimits().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("apiversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hub().V1alpha1().APIVersions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("accesscontrolpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Hub().V1alpha1().AccessControlPolicies().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
