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
	examplev1 "github.com/YaoZengzeng/practice/kubernetes-controller-example/pkg/apis/example/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDogs implements DogInterface
type FakeDogs struct {
	Fake *FakeExampleV1
	ns   string
}

var dogsResource = schema.GroupVersionResource{Group: "example", Version: "v1", Resource: "dogs"}

var dogsKind = schema.GroupVersionKind{Group: "example", Version: "v1", Kind: "Dog"}

// Get takes name of the dog, and returns the corresponding dog object, and an error if there is any.
func (c *FakeDogs) Get(name string, options v1.GetOptions) (result *examplev1.Dog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dogsResource, c.ns, name), &examplev1.Dog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*examplev1.Dog), err
}

// List takes label and field selectors, and returns the list of Dogs that match those selectors.
func (c *FakeDogs) List(opts v1.ListOptions) (result *examplev1.DogList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dogsResource, dogsKind, c.ns, opts), &examplev1.DogList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &examplev1.DogList{ListMeta: obj.(*examplev1.DogList).ListMeta}
	for _, item := range obj.(*examplev1.DogList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dogs.
func (c *FakeDogs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dogsResource, c.ns, opts))

}

// Create takes the representation of a dog and creates it.  Returns the server's representation of the dog, and an error, if there is any.
func (c *FakeDogs) Create(dog *examplev1.Dog) (result *examplev1.Dog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dogsResource, c.ns, dog), &examplev1.Dog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*examplev1.Dog), err
}

// Update takes the representation of a dog and updates it. Returns the server's representation of the dog, and an error, if there is any.
func (c *FakeDogs) Update(dog *examplev1.Dog) (result *examplev1.Dog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dogsResource, c.ns, dog), &examplev1.Dog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*examplev1.Dog), err
}

// Delete takes name of the dog and deletes it. Returns an error if one occurs.
func (c *FakeDogs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dogsResource, c.ns, name), &examplev1.Dog{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDogs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dogsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &examplev1.DogList{})
	return err
}

// Patch applies the patch and returns the patched dog.
func (c *FakeDogs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *examplev1.Dog, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dogsResource, c.ns, name, pt, data, subresources...), &examplev1.Dog{})

	if obj == nil {
		return nil, err
	}
	return obj.(*examplev1.Dog), err
}
