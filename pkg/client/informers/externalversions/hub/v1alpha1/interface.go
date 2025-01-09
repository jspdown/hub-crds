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
	internalinterfaces "github.com/traefik/hub-crds/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AIServices returns a AIServiceInformer.
	AIServices() AIServiceInformer
	// APIs returns a APIInformer.
	APIs() APIInformer
	// APIAccesses returns a APIAccessInformer.
	APIAccesses() APIAccessInformer
	// APIBundles returns a APIBundleInformer.
	APIBundles() APIBundleInformer
	// APICatalogItems returns a APICatalogItemInformer.
	APICatalogItems() APICatalogItemInformer
	// APIPlans returns a APIPlanInformer.
	APIPlans() APIPlanInformer
	// APIPortals returns a APIPortalInformer.
	APIPortals() APIPortalInformer
	// APIRateLimits returns a APIRateLimitInformer.
	APIRateLimits() APIRateLimitInformer
	// APIVersions returns a APIVersionInformer.
	APIVersions() APIVersionInformer
	// AccessControlPolicies returns a AccessControlPolicyInformer.
	AccessControlPolicies() AccessControlPolicyInformer
	// ManagedSubscriptions returns a ManagedSubscriptionInformer.
	ManagedSubscriptions() ManagedSubscriptionInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AIServices returns a AIServiceInformer.
func (v *version) AIServices() AIServiceInformer {
	return &aIServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIs returns a APIInformer.
func (v *version) APIs() APIInformer {
	return &aPIInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIAccesses returns a APIAccessInformer.
func (v *version) APIAccesses() APIAccessInformer {
	return &aPIAccessInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIBundles returns a APIBundleInformer.
func (v *version) APIBundles() APIBundleInformer {
	return &aPIBundleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APICatalogItems returns a APICatalogItemInformer.
func (v *version) APICatalogItems() APICatalogItemInformer {
	return &aPICatalogItemInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIPlans returns a APIPlanInformer.
func (v *version) APIPlans() APIPlanInformer {
	return &aPIPlanInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIPortals returns a APIPortalInformer.
func (v *version) APIPortals() APIPortalInformer {
	return &aPIPortalInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIRateLimits returns a APIRateLimitInformer.
func (v *version) APIRateLimits() APIRateLimitInformer {
	return &aPIRateLimitInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// APIVersions returns a APIVersionInformer.
func (v *version) APIVersions() APIVersionInformer {
	return &aPIVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AccessControlPolicies returns a AccessControlPolicyInformer.
func (v *version) AccessControlPolicies() AccessControlPolicyInformer {
	return &accessControlPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ManagedSubscriptions returns a ManagedSubscriptionInformer.
func (v *version) ManagedSubscriptions() ManagedSubscriptionInformer {
	return &managedSubscriptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
