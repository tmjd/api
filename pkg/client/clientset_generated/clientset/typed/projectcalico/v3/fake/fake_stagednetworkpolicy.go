// Copyright (c) 2022 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStagedNetworkPolicies implements StagedNetworkPolicyInterface
type FakeStagedNetworkPolicies struct {
	Fake *FakeProjectcalicoV3
	ns   string
}

var stagednetworkpoliciesResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "stagednetworkpolicies"}

var stagednetworkpoliciesKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "StagedNetworkPolicy"}

// Get takes name of the stagedNetworkPolicy, and returns the corresponding stagedNetworkPolicy object, and an error if there is any.
func (c *FakeStagedNetworkPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.StagedNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(stagednetworkpoliciesResource, c.ns, name), &v3.StagedNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedNetworkPolicy), err
}

// List takes label and field selectors, and returns the list of StagedNetworkPolicies that match those selectors.
func (c *FakeStagedNetworkPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v3.StagedNetworkPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(stagednetworkpoliciesResource, stagednetworkpoliciesKind, c.ns, opts), &v3.StagedNetworkPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.StagedNetworkPolicyList{ListMeta: obj.(*v3.StagedNetworkPolicyList).ListMeta}
	for _, item := range obj.(*v3.StagedNetworkPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested stagedNetworkPolicies.
func (c *FakeStagedNetworkPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(stagednetworkpoliciesResource, c.ns, opts))

}

// Create takes the representation of a stagedNetworkPolicy and creates it.  Returns the server's representation of the stagedNetworkPolicy, and an error, if there is any.
func (c *FakeStagedNetworkPolicies) Create(ctx context.Context, stagedNetworkPolicy *v3.StagedNetworkPolicy, opts v1.CreateOptions) (result *v3.StagedNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(stagednetworkpoliciesResource, c.ns, stagedNetworkPolicy), &v3.StagedNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedNetworkPolicy), err
}

// Update takes the representation of a stagedNetworkPolicy and updates it. Returns the server's representation of the stagedNetworkPolicy, and an error, if there is any.
func (c *FakeStagedNetworkPolicies) Update(ctx context.Context, stagedNetworkPolicy *v3.StagedNetworkPolicy, opts v1.UpdateOptions) (result *v3.StagedNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(stagednetworkpoliciesResource, c.ns, stagedNetworkPolicy), &v3.StagedNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedNetworkPolicy), err
}

// Delete takes name of the stagedNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *FakeStagedNetworkPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(stagednetworkpoliciesResource, c.ns, name, opts), &v3.StagedNetworkPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStagedNetworkPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(stagednetworkpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.StagedNetworkPolicyList{})
	return err
}

// Patch applies the patch and returns the patched stagedNetworkPolicy.
func (c *FakeStagedNetworkPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.StagedNetworkPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(stagednetworkpoliciesResource, c.ns, name, pt, data, subresources...), &v3.StagedNetworkPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.StagedNetworkPolicy), err
}
