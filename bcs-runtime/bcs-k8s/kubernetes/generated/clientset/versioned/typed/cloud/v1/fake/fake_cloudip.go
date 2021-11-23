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

// FakeCloudIPs implements CloudIPInterface
type FakeCloudIPs struct {
	Fake *FakeCloudV1
	ns   string
}

var cloudipsResource = schema.GroupVersionResource{Group: "cloud", Version: "v1", Resource: "cloudips"}

var cloudipsKind = schema.GroupVersionKind{Group: "cloud", Version: "v1", Kind: "CloudIP"}

// Get takes name of the cloudIP, and returns the corresponding cloudIP object, and an error if there is any.
func (c *FakeCloudIPs) Get(ctx context.Context, name string, options v1.GetOptions) (result *cloudv1.CloudIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudipsResource, c.ns, name), &cloudv1.CloudIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIP), err
}

// List takes label and field selectors, and returns the list of CloudIPs that match those selectors.
func (c *FakeCloudIPs) List(ctx context.Context, opts v1.ListOptions) (result *cloudv1.CloudIPList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudipsResource, cloudipsKind, c.ns, opts), &cloudv1.CloudIPList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cloudv1.CloudIPList{ListMeta: obj.(*cloudv1.CloudIPList).ListMeta}
	for _, item := range obj.(*cloudv1.CloudIPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudIPs.
func (c *FakeCloudIPs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudipsResource, c.ns, opts))

}

// Create takes the representation of a cloudIP and creates it.  Returns the server's representation of the cloudIP, and an error, if there is any.
func (c *FakeCloudIPs) Create(ctx context.Context, cloudIP *cloudv1.CloudIP, opts v1.CreateOptions) (result *cloudv1.CloudIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudipsResource, c.ns, cloudIP), &cloudv1.CloudIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIP), err
}

// Update takes the representation of a cloudIP and updates it. Returns the server's representation of the cloudIP, and an error, if there is any.
func (c *FakeCloudIPs) Update(ctx context.Context, cloudIP *cloudv1.CloudIP, opts v1.UpdateOptions) (result *cloudv1.CloudIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudipsResource, c.ns, cloudIP), &cloudv1.CloudIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudIPs) UpdateStatus(ctx context.Context, cloudIP *cloudv1.CloudIP, opts v1.UpdateOptions) (*cloudv1.CloudIP, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudipsResource, "status", c.ns, cloudIP), &cloudv1.CloudIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIP), err
}

// Delete takes name of the cloudIP and deletes it. Returns an error if one occurs.
func (c *FakeCloudIPs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudipsResource, c.ns, name), &cloudv1.CloudIP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudIPs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudipsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &cloudv1.CloudIPList{})
	return err
}

// Patch applies the patch and returns the patched cloudIP.
func (c *FakeCloudIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *cloudv1.CloudIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudipsResource, c.ns, name, pt, data, subresources...), &cloudv1.CloudIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.CloudIP), err
}