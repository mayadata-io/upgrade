/*
Copyright 2019 The OpenEBS Authors

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	scheme "github.com/openebs/maya/pkg/client/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CStorCompletedBackupsGetter has a method to return a CStorCompletedBackupInterface.
// A group's client should implement this interface.
type CStorCompletedBackupsGetter interface {
	CStorCompletedBackups(namespace string) CStorCompletedBackupInterface
}

// CStorCompletedBackupInterface has methods to work with CStorCompletedBackup resources.
type CStorCompletedBackupInterface interface {
	Create(*v1alpha1.CStorCompletedBackup) (*v1alpha1.CStorCompletedBackup, error)
	Update(*v1alpha1.CStorCompletedBackup) (*v1alpha1.CStorCompletedBackup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CStorCompletedBackup, error)
	List(opts v1.ListOptions) (*v1alpha1.CStorCompletedBackupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CStorCompletedBackup, err error)
	CStorCompletedBackupExpansion
}

// cStorCompletedBackups implements CStorCompletedBackupInterface
type cStorCompletedBackups struct {
	client rest.Interface
	ns     string
}

// newCStorCompletedBackups returns a CStorCompletedBackups
func newCStorCompletedBackups(c *OpenebsV1alpha1Client, namespace string) *cStorCompletedBackups {
	return &cStorCompletedBackups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cStorCompletedBackup, and returns the corresponding cStorCompletedBackup object, and an error if there is any.
func (c *cStorCompletedBackups) Get(name string, options v1.GetOptions) (result *v1alpha1.CStorCompletedBackup, err error) {
	result = &v1alpha1.CStorCompletedBackup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cstorcompletedbackups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CStorCompletedBackups that match those selectors.
func (c *cStorCompletedBackups) List(opts v1.ListOptions) (result *v1alpha1.CStorCompletedBackupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CStorCompletedBackupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cstorcompletedbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cStorCompletedBackups.
func (c *cStorCompletedBackups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cstorcompletedbackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cStorCompletedBackup and creates it.  Returns the server's representation of the cStorCompletedBackup, and an error, if there is any.
func (c *cStorCompletedBackups) Create(cStorCompletedBackup *v1alpha1.CStorCompletedBackup) (result *v1alpha1.CStorCompletedBackup, err error) {
	result = &v1alpha1.CStorCompletedBackup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cstorcompletedbackups").
		Body(cStorCompletedBackup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cStorCompletedBackup and updates it. Returns the server's representation of the cStorCompletedBackup, and an error, if there is any.
func (c *cStorCompletedBackups) Update(cStorCompletedBackup *v1alpha1.CStorCompletedBackup) (result *v1alpha1.CStorCompletedBackup, err error) {
	result = &v1alpha1.CStorCompletedBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cstorcompletedbackups").
		Name(cStorCompletedBackup.Name).
		Body(cStorCompletedBackup).
		Do().
		Into(result)
	return
}

// Delete takes name of the cStorCompletedBackup and deletes it. Returns an error if one occurs.
func (c *cStorCompletedBackups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cstorcompletedbackups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cStorCompletedBackups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cstorcompletedbackups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cStorCompletedBackup.
func (c *cStorCompletedBackups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CStorCompletedBackup, err error) {
	result = &v1alpha1.CStorCompletedBackup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cstorcompletedbackups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
