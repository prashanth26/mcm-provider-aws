/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package v1alpha2

import (
	"time"

	v1alpha2 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha2"
	scheme "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineClassesGetter has a method to return a MachineClassInterface.
// A group's client should implement this interface.
type MachineClassesGetter interface {
	MachineClasses(namespace string) MachineClassInterface
}

// MachineClassInterface has methods to work with MachineClass resources.
type MachineClassInterface interface {
	Create(*v1alpha2.MachineClass) (*v1alpha2.MachineClass, error)
	Update(*v1alpha2.MachineClass) (*v1alpha2.MachineClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.MachineClass, error)
	List(opts v1.ListOptions) (*v1alpha2.MachineClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.MachineClass, err error)
	MachineClassExpansion
}

// machineClasses implements MachineClassInterface
type machineClasses struct {
	client rest.Interface
	ns     string
}

// newMachineClasses returns a MachineClasses
func newMachineClasses(c *MachineV1alpha2Client, namespace string) *machineClasses {
	return &machineClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the machineClass, and returns the corresponding machineClass object, and an error if there is any.
func (c *machineClasses) Get(name string, options v1.GetOptions) (result *v1alpha2.MachineClass, err error) {
	result = &v1alpha2.MachineClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machineclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineClasses that match those selectors.
func (c *machineClasses) List(opts v1.ListOptions) (result *v1alpha2.MachineClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.MachineClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineClasses.
func (c *machineClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("machineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a machineClass and creates it.  Returns the server's representation of the machineClass, and an error, if there is any.
func (c *machineClasses) Create(machineClass *v1alpha2.MachineClass) (result *v1alpha2.MachineClass, err error) {
	result = &v1alpha2.MachineClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("machineclasses").
		Body(machineClass).
		Do().
		Into(result)
	return
}

// Update takes the representation of a machineClass and updates it. Returns the server's representation of the machineClass, and an error, if there is any.
func (c *machineClasses) Update(machineClass *v1alpha2.MachineClass) (result *v1alpha2.MachineClass, err error) {
	result = &v1alpha2.MachineClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machineclasses").
		Name(machineClass.Name).
		Body(machineClass).
		Do().
		Into(result)
	return
}

// Delete takes name of the machineClass and deletes it. Returns an error if one occurs.
func (c *machineClasses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machineclasses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machineclasses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched machineClass.
func (c *machineClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.MachineClass, err error) {
	result = &v1alpha2.MachineClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("machineclasses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
