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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// APIBundleLister helps list APIBundles.
// All objects returned here must be treated as read-only.
type APIBundleLister interface {
	// List lists all APIBundles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.APIBundle, err error)
	// APIBundles returns an object that can list and get APIBundles.
	APIBundles(namespace string) APIBundleNamespaceLister
	APIBundleListerExpansion
}

// aPIBundleLister implements the APIBundleLister interface.
type aPIBundleLister struct {
	indexer cache.Indexer
}

// NewAPIBundleLister returns a new APIBundleLister.
func NewAPIBundleLister(indexer cache.Indexer) APIBundleLister {
	return &aPIBundleLister{indexer: indexer}
}

// List lists all APIBundles in the indexer.
func (s *aPIBundleLister) List(selector labels.Selector) (ret []*v1alpha1.APIBundle, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.APIBundle))
	})
	return ret, err
}

// APIBundles returns an object that can list and get APIBundles.
func (s *aPIBundleLister) APIBundles(namespace string) APIBundleNamespaceLister {
	return aPIBundleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// APIBundleNamespaceLister helps list and get APIBundles.
// All objects returned here must be treated as read-only.
type APIBundleNamespaceLister interface {
	// List lists all APIBundles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.APIBundle, err error)
	// Get retrieves the APIBundle from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.APIBundle, error)
	APIBundleNamespaceListerExpansion
}

// aPIBundleNamespaceLister implements the APIBundleNamespaceLister
// interface.
type aPIBundleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all APIBundles in the indexer for a given namespace.
func (s aPIBundleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.APIBundle, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.APIBundle))
	})
	return ret, err
}

// Get retrieves the APIBundle from the indexer for a given namespace and name.
func (s aPIBundleNamespaceLister) Get(name string) (*v1alpha1.APIBundle, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apibundle"), name)
	}
	return obj.(*v1alpha1.APIBundle), nil
}
