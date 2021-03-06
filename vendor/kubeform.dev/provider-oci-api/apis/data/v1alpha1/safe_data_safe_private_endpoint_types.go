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

type SafeDataSafePrivateEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SafeDataSafePrivateEndpointSpec   `json:"spec,omitempty"`
	Status            SafeDataSafePrivateEndpointStatus `json:"status,omitempty"`
}

type SafeDataSafePrivateEndpointSpec struct {
	State *SafeDataSafePrivateEndpointSpecResource `json:"state,omitempty" tf:"-"`

	Resource SafeDataSafePrivateEndpointSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SafeDataSafePrivateEndpointSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	// +optional
	EndpointFqdn *string `json:"endpointFqdn,omitempty" tf:"endpoint_fqdn"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	NsgIDS []string `json:"nsgIDS,omitempty" tf:"nsg_ids"`
	// +optional
	PrivateEndpointID *string `json:"privateEndpointID,omitempty" tf:"private_endpoint_id"`
	// +optional
	PrivateEndpointIP *string `json:"privateEndpointIP,omitempty" tf:"private_endpoint_ip"`
	// +optional
	State    *string `json:"state,omitempty" tf:"state"`
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	VcnID       *string `json:"vcnID" tf:"vcn_id"`
}

type SafeDataSafePrivateEndpointStatus struct {
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

// SafeDataSafePrivateEndpointList is a list of SafeDataSafePrivateEndpoints
type SafeDataSafePrivateEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SafeDataSafePrivateEndpoint CRD objects
	Items []SafeDataSafePrivateEndpoint `json:"items,omitempty"`
}
