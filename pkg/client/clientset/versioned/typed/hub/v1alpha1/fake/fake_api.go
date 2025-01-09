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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAPIs implements APIInterface
type FakeAPIs struct {
	Fake *FakeHubV1alpha1
	ns   string
}

var apisResource = v1alpha1.SchemeGroupVersion.WithResource("apis")

var apisKind = v1alpha1.SchemeGroupVersion.WithKind("API")

// Get takes name of the aPI, and returns the corresponding aPI object, and an error if there is any.
func (c *FakeAPIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.API, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apisResource, c.ns, name), &v1alpha1.API{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.API), err
}

// List takes label and field selectors, and returns the list of APIs that match those selectors.
func (c *FakeAPIs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apisResource, apisKind, c.ns, opts), &v1alpha1.APIList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIList{ListMeta: obj.(*v1alpha1.APIList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIs.
func (c *FakeAPIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apisResource, c.ns, opts))

}

// Create takes the representation of a aPI and creates it.  Returns the server's representation of the aPI, and an error, if there is any.
func (c *FakeAPIs) Create(ctx context.Context, aPI *v1alpha1.API, opts v1.CreateOptions) (result *v1alpha1.API, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apisResource, c.ns, aPI), &v1alpha1.API{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.API), err
}

// Update takes the representation of a aPI and updates it. Returns the server's representation of the aPI, and an error, if there is any.
func (c *FakeAPIs) Update(ctx context.Context, aPI *v1alpha1.API, opts v1.UpdateOptions) (result *v1alpha1.API, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apisResource, c.ns, aPI), &v1alpha1.API{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.API), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIs) UpdateStatus(ctx context.Context, aPI *v1alpha1.API, opts v1.UpdateOptions) (*v1alpha1.API, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apisResource, "status", c.ns, aPI), &v1alpha1.API{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.API), err
}

// Delete takes name of the aPI and deletes it. Returns an error if one occurs.
func (c *FakeAPIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apisResource, c.ns, name, opts), &v1alpha1.API{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIList{})
	return err
}

// Patch applies the patch and returns the patched aPI.
func (c *FakeAPIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.API, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apisResource, c.ns, name, pt, data, subresources...), &v1alpha1.API{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.API), err
}
