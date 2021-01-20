/*
Copyright 2021.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StoreFrntSpec defines the desired state of StoreFrnt
type StoreFrntSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Size is an example field of StoreFrnt. Edit StoreFrnt_types.go to remove/update
	Size     int32  `json:"size,omitempty"`
	Location string `json:"location,omitempty"`
}

// StoreFrntStatus defines the observed state of StoreFrnt
type StoreFrntStatus struct {
	Phase string `json:"phase,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=sf
// StoreFrnt is the Schema for the storefrnts API
type StoreFrnt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StoreFrntSpec   `json:"spec,omitempty"`
	Status StoreFrntStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoreFrntList contains a list of StoreFrnt
type StoreFrntList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoreFrnt `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StoreFrnt{}, &StoreFrntList{})
}
