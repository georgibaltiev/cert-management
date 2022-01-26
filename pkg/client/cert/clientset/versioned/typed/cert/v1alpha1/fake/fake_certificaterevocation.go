/*
Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gardener/cert-management/pkg/apis/cert/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificateRevocations implements CertificateRevocationInterface
type FakeCertificateRevocations struct {
	Fake *FakeCertV1alpha1
	ns   string
}

var certificaterevocationsResource = schema.GroupVersionResource{Group: "cert.gardener.cloud", Version: "v1alpha1", Resource: "certificaterevocations"}

var certificaterevocationsKind = schema.GroupVersionKind{Group: "cert.gardener.cloud", Version: "v1alpha1", Kind: "CertificateRevocation"}

// Get takes name of the certificateRevocation, and returns the corresponding certificateRevocation object, and an error if there is any.
func (c *FakeCertificateRevocations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CertificateRevocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificaterevocationsResource, c.ns, name), &v1alpha1.CertificateRevocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRevocation), err
}

// List takes label and field selectors, and returns the list of CertificateRevocations that match those selectors.
func (c *FakeCertificateRevocations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CertificateRevocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificaterevocationsResource, certificaterevocationsKind, c.ns, opts), &v1alpha1.CertificateRevocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CertificateRevocationList{ListMeta: obj.(*v1alpha1.CertificateRevocationList).ListMeta}
	for _, item := range obj.(*v1alpha1.CertificateRevocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificateRevocations.
func (c *FakeCertificateRevocations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificaterevocationsResource, c.ns, opts))

}

// Create takes the representation of a certificateRevocation and creates it.  Returns the server's representation of the certificateRevocation, and an error, if there is any.
func (c *FakeCertificateRevocations) Create(ctx context.Context, certificateRevocation *v1alpha1.CertificateRevocation, opts v1.CreateOptions) (result *v1alpha1.CertificateRevocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificaterevocationsResource, c.ns, certificateRevocation), &v1alpha1.CertificateRevocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRevocation), err
}

// Update takes the representation of a certificateRevocation and updates it. Returns the server's representation of the certificateRevocation, and an error, if there is any.
func (c *FakeCertificateRevocations) Update(ctx context.Context, certificateRevocation *v1alpha1.CertificateRevocation, opts v1.UpdateOptions) (result *v1alpha1.CertificateRevocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificaterevocationsResource, c.ns, certificateRevocation), &v1alpha1.CertificateRevocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRevocation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificateRevocations) UpdateStatus(ctx context.Context, certificateRevocation *v1alpha1.CertificateRevocation, opts v1.UpdateOptions) (*v1alpha1.CertificateRevocation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(certificaterevocationsResource, "status", c.ns, certificateRevocation), &v1alpha1.CertificateRevocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRevocation), err
}

// Delete takes name of the certificateRevocation and deletes it. Returns an error if one occurs.
func (c *FakeCertificateRevocations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(certificaterevocationsResource, c.ns, name), &v1alpha1.CertificateRevocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificateRevocations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificaterevocationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CertificateRevocationList{})
	return err
}

// Patch applies the patch and returns the patched certificateRevocation.
func (c *FakeCertificateRevocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CertificateRevocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificaterevocationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CertificateRevocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertificateRevocation), err
}
