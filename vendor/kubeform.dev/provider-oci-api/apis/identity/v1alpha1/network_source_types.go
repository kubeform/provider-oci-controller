/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type NetworkSource struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSourceSpec   `json:"spec,omitempty"`
	Status            NetworkSourceStatus `json:"status,omitempty"`
}

type NetworkSourceSpecVirtualSourceList struct {
	IpRanges []string `json:"ipRanges" tf:"ip_ranges"`
	VcnID    *string  `json:"vcnID" tf:"vcn_id"`
}

type NetworkSourceSpec struct {
	State *NetworkSourceSpecResource `json:"state,omitempty" tf:"-"`

	Resource NetworkSourceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type NetworkSourceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	Description *string           `json:"description" tf:"description"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	InactiveState *string `json:"inactiveState,omitempty" tf:"inactive_state"`
	Name          *string `json:"name" tf:"name"`
	// +optional
	PublicSourceList []string `json:"publicSourceList,omitempty" tf:"public_source_list"`
	// +optional
	Services []string `json:"services,omitempty" tf:"services"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	VirtualSourceList []NetworkSourceSpecVirtualSourceList `json:"virtualSourceList,omitempty" tf:"virtual_source_list"`
}

type NetworkSourceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkSourceList is a list of NetworkSources
type NetworkSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkSource CRD objects
	Items []NetworkSource `json:"items,omitempty"`
}
