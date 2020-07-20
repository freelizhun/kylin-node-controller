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

package v1

import (
	"time"

	v1 "github.com/kylin/kylin-node-controller/pkg/apis/crd/v1"
	scheme "github.com/kylin/kylin-node-controller/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KylinNodesGetter has a method to return a KylinNodeInterface.
// A group's client should implement this interface.
type KylinNodesGetter interface {
	KylinNodes(namespace string) KylinNodeInterface
}

// KylinNodeInterface has methods to work with KylinNode resources.
type KylinNodeInterface interface {
	Create(*v1.KylinNode) (*v1.KylinNode, error)
	Update(*v1.KylinNode) (*v1.KylinNode, error)
	UpdateStatus(*v1.KylinNode) (*v1.KylinNode, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.KylinNode, error)
	List(opts metav1.ListOptions) (*v1.KylinNodeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KylinNode, err error)
	KylinNodeExpansion
}

// kylinNodes implements KylinNodeInterface
type kylinNodes struct {
	client rest.Interface
	ns     string
}

// newKylinNodes returns a KylinNodes
func newKylinNodes(c *CrdV1Client, namespace string) *kylinNodes {
	return &kylinNodes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kylinNode, and returns the corresponding kylinNode object, and an error if there is any.
func (c *kylinNodes) Get(name string, options metav1.GetOptions) (result *v1.KylinNode, err error) {
	result = &v1.KylinNode{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kylinnodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KylinNodes that match those selectors.
func (c *kylinNodes) List(opts metav1.ListOptions) (result *v1.KylinNodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KylinNodeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kylinnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kylinNodes.
func (c *kylinNodes) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kylinnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kylinNode and creates it.  Returns the server's representation of the kylinNode, and an error, if there is any.
func (c *kylinNodes) Create(kylinNode *v1.KylinNode) (result *v1.KylinNode, err error) {
	result = &v1.KylinNode{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kylinnodes").
		Body(kylinNode).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kylinNode and updates it. Returns the server's representation of the kylinNode, and an error, if there is any.
func (c *kylinNodes) Update(kylinNode *v1.KylinNode) (result *v1.KylinNode, err error) {
	result = &v1.KylinNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kylinnodes").
		Name(kylinNode.Name).
		Body(kylinNode).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *kylinNodes) UpdateStatus(kylinNode *v1.KylinNode) (result *v1.KylinNode, err error) {
	result = &v1.KylinNode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kylinnodes").
		Name(kylinNode.Name).
		SubResource("status").
		Body(kylinNode).
		Do().
		Into(result)
	return
}

// Delete takes name of the kylinNode and deletes it. Returns an error if one occurs.
func (c *kylinNodes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kylinnodes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kylinNodes) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kylinnodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kylinNode.
func (c *kylinNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KylinNode, err error) {
	result = &v1.KylinNode{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kylinnodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}