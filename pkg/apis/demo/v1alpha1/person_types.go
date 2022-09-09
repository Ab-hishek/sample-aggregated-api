
/*
Copyright 2022.

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

package v1alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Person
// +k8s:openapi-gen=true
type Person struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PersonSpec   `json:"spec,omitempty"`
	Status PersonStatus `json:"status,omitempty"`
}

// PersonList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type PersonList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Person `json:"items"`
}

// PersonSpec defines the desired state of Person
type PersonSpec struct {
	// Name of the person
	// +required
	Name string `json:"name"`
	// Age of the person
	// +required
	Age int32 `json:"age"`
}

var _ resource.Object = &Person{}
var _ resourcestrategy.Validater = &Person{}

func (in *Person) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *Person) NamespaceScoped() bool {
	return true
}

func (in *Person) New() runtime.Object {
	return &Person{}
}

func (in *Person) NewList() runtime.Object {
	return &PersonList{}
}

func (in *Person) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "demo.sample.abhishek.io",
		Version:  "v1alpha1",
		Resource: "people",
	}
}

func (in *Person) IsStorageVersion() bool {
	return true
}

func (in *Person) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &PersonList{}

func (in *PersonList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// PersonStatus defines the observed state of Person
type PersonStatus struct {
	// Category of the person based on age
	// +optional
	Category Category `json:"category"`
}

type Category string

const (
	MinorCategory Category = "minor"
	AdultCategory Category = "adult"
	YouthCategory Category = "youth"
)

func (in PersonStatus) SubResourceName() string {
	return "status"
}

// Person implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &Person{}

func (in *Person) GetStatus() resource.StatusSubResource {
	return in.Status
}

// PersonStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &PersonStatus{}

func (in PersonStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*Person).Status = in
}
