/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	cloudv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/cloud/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudSubnets implements CloudSubnetInterface
type FakeCloudSubnets struct {
	Fake *FakeCloudV1
	ns   string
}

var cloudsubnetsResource = schema.GroupVersionResource{Group: "cloud", Version: "v1", Resource: "cloudsubnets"}

var cloudsubnetsKind = schema.GroupVersionKind{Group: "cloud", Version: "v1", Kind: "CloudSubnet"}

// Get takes name of the cloudSubnet, and returns the corresponding cloudSubnet object, and an error if there is any.
func (c *FakeCloudSubnets) Get(ctx context.Context, name string, options v1.GetOptions) (result *cloudv1.CloudSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudsubnetsResource, c.ns, name), &cloudv1.CloudSubnet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudSubnet), err
}

// List takes label and field selectors, and returns the list of CloudSubnets that match those selectors.
func (c *FakeCloudSubnets) List(ctx context.Context, opts v1.ListOptions) (result *cloudv1.CloudSubnetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudsubnetsResource, cloudsubnetsKind, c.ns, opts), &cloudv1.CloudSubnetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cloudv1.CloudSubnetList{ListMeta: obj.(*cloudv1.CloudSubnetList).ListMeta}
	for _, item := range obj.(*cloudv1.CloudSubnetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudSubnets.
func (c *FakeCloudSubnets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudsubnetsResource, c.ns, opts))

}

// Create takes the representation of a cloudSubnet and creates it.  Returns the server's representation of the cloudSubnet, and an error, if there is any.
func (c *FakeCloudSubnets) Create(ctx context.Context, cloudSubnet *cloudv1.CloudSubnet, opts v1.CreateOptions) (result *cloudv1.CloudSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudsubnetsResource, c.ns, cloudSubnet), &cloudv1.CloudSubnet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudSubnet), err
}

// Update takes the representation of a cloudSubnet and updates it. Returns the server's representation of the cloudSubnet, and an error, if there is any.
func (c *FakeCloudSubnets) Update(ctx context.Context, cloudSubnet *cloudv1.CloudSubnet, opts v1.UpdateOptions) (result *cloudv1.CloudSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudsubnetsResource, c.ns, cloudSubnet), &cloudv1.CloudSubnet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudSubnet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudSubnets) UpdateStatus(ctx context.Context, cloudSubnet *cloudv1.CloudSubnet, opts v1.UpdateOptions) (*cloudv1.CloudSubnet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudsubnetsResource, "status", c.ns, cloudSubnet), &cloudv1.CloudSubnet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudSubnet), err
}

// Delete takes name of the cloudSubnet and deletes it. Returns an error if one occurs.
func (c *FakeCloudSubnets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudsubnetsResource, c.ns, name), &cloudv1.CloudSubnet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudSubnets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudsubnetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &cloudv1.CloudSubnetList{})
	return err
}

// Patch applies the patch and returns the patched cloudSubnet.
func (c *FakeCloudSubnets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cloudv1.CloudSubnet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudsubnetsResource, c.ns, name, pt, data, subresources...), &cloudv1.CloudSubnet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudSubnet), err
}