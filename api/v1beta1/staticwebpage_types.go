/*


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

// StaticWebpageSpec defines the desired state of StaticWebpage
type StaticWebpageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of StaticWebpage. Edit StaticWebpage_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// StaticWebpageStatus defines the observed state of StaticWebpage
type StaticWebpageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// StaticWebpage is the Schema for the staticwebpages API
type StaticWebpage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StaticWebpageSpec   `json:"spec,omitempty"`
	Status StaticWebpageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StaticWebpageList contains a list of StaticWebpage
type StaticWebpageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StaticWebpage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StaticWebpage{}, &StaticWebpageList{})
}
