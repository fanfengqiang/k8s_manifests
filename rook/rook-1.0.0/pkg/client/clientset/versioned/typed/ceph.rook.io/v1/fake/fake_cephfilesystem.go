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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	cephrookiov1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCephFilesystems implements CephFilesystemInterface
type FakeCephFilesystems struct {
	Fake *FakeCephV1
	ns   string
}

var cephfilesystemsResource = schema.GroupVersionResource{Group: "ceph.rook.io", Version: "v1", Resource: "cephfilesystems"}

var cephfilesystemsKind = schema.GroupVersionKind{Group: "ceph.rook.io", Version: "v1", Kind: "CephFilesystem"}

// Get takes name of the cephFilesystem, and returns the corresponding cephFilesystem object, and an error if there is any.
func (c *FakeCephFilesystems) Get(name string, options v1.GetOptions) (result *cephrookiov1.CephFilesystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cephfilesystemsResource, c.ns, name), &cephrookiov1.CephFilesystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephFilesystem), err
}

// List takes label and field selectors, and returns the list of CephFilesystems that match those selectors.
func (c *FakeCephFilesystems) List(opts v1.ListOptions) (result *cephrookiov1.CephFilesystemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cephfilesystemsResource, cephfilesystemsKind, c.ns, opts), &cephrookiov1.CephFilesystemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cephrookiov1.CephFilesystemList{ListMeta: obj.(*cephrookiov1.CephFilesystemList).ListMeta}
	for _, item := range obj.(*cephrookiov1.CephFilesystemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephFilesystems.
func (c *FakeCephFilesystems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cephfilesystemsResource, c.ns, opts))

}

// Create takes the representation of a cephFilesystem and creates it.  Returns the server's representation of the cephFilesystem, and an error, if there is any.
func (c *FakeCephFilesystems) Create(cephFilesystem *cephrookiov1.CephFilesystem) (result *cephrookiov1.CephFilesystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cephfilesystemsResource, c.ns, cephFilesystem), &cephrookiov1.CephFilesystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephFilesystem), err
}

// Update takes the representation of a cephFilesystem and updates it. Returns the server's representation of the cephFilesystem, and an error, if there is any.
func (c *FakeCephFilesystems) Update(cephFilesystem *cephrookiov1.CephFilesystem) (result *cephrookiov1.CephFilesystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cephfilesystemsResource, c.ns, cephFilesystem), &cephrookiov1.CephFilesystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephFilesystem), err
}

// Delete takes name of the cephFilesystem and deletes it. Returns an error if one occurs.
func (c *FakeCephFilesystems) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cephfilesystemsResource, c.ns, name), &cephrookiov1.CephFilesystem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephFilesystems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cephfilesystemsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &cephrookiov1.CephFilesystemList{})
	return err
}

// Patch applies the patch and returns the patched cephFilesystem.
func (c *FakeCephFilesystems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cephrookiov1.CephFilesystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cephfilesystemsResource, c.ns, name, pt, data, subresources...), &cephrookiov1.CephFilesystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*cephrookiov1.CephFilesystem), err
}
