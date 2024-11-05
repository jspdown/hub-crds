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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"net/http"

	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	"github.com/traefik/hub-crds/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type HubV1alpha1Interface interface {
	RESTClient() rest.Interface
	APIsGetter
	APIAccessesGetter
	APIBundlesGetter
	APIPlansGetter
	APIPortalsGetter
	APIRateLimitsGetter
	APIVersionsGetter
	AccessControlPoliciesGetter
	ManagedSubscriptionsGetter
}

// HubV1alpha1Client is used to interact with features provided by the hub.traefik.io group.
type HubV1alpha1Client struct {
	restClient rest.Interface
}

func (c *HubV1alpha1Client) APIs(namespace string) APIInterface {
	return newAPIs(c, namespace)
}

func (c *HubV1alpha1Client) APIAccesses(namespace string) APIAccessInterface {
	return newAPIAccesses(c, namespace)
}

func (c *HubV1alpha1Client) APIBundles(namespace string) APIBundleInterface {
	return newAPIBundles(c, namespace)
}

func (c *HubV1alpha1Client) APIPlans(namespace string) APIPlanInterface {
	return newAPIPlans(c, namespace)
}

func (c *HubV1alpha1Client) APIPortals(namespace string) APIPortalInterface {
	return newAPIPortals(c, namespace)
}

func (c *HubV1alpha1Client) APIRateLimits(namespace string) APIRateLimitInterface {
	return newAPIRateLimits(c, namespace)
}

func (c *HubV1alpha1Client) APIVersions(namespace string) APIVersionInterface {
	return newAPIVersions(c, namespace)
}

func (c *HubV1alpha1Client) AccessControlPolicies() AccessControlPolicyInterface {
	return newAccessControlPolicies(c)
}

func (c *HubV1alpha1Client) ManagedSubscriptions(namespace string) ManagedSubscriptionInterface {
	return newManagedSubscriptions(c, namespace)
}

// NewForConfig creates a new HubV1alpha1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*HubV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new HubV1alpha1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*HubV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &HubV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new HubV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *HubV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new HubV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *HubV1alpha1Client {
	return &HubV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *HubV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
