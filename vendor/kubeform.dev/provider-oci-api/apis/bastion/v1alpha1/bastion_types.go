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

type Bastion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BastionSpec   `json:"spec,omitempty"`
	Status            BastionStatus `json:"status,omitempty"`
}

type BastionSpec struct {
	State *BastionSpecResource `json:"state,omitempty" tf:"-"`

	Resource BastionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BastionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BastionType *string `json:"bastionType" tf:"bastion_type"`
	// +optional
	ClientCIDRBlockAllowList []string `json:"clientCIDRBlockAllowList,omitempty" tf:"client_cidr_block_allow_list"`
	CompartmentID            *string  `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	MaxSessionTtlInSeconds *int64 `json:"maxSessionTtlInSeconds,omitempty" tf:"max_session_ttl_in_seconds"`
	// +optional
	MaxSessionsAllowed *int64 `json:"maxSessionsAllowed,omitempty" tf:"max_sessions_allowed"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PhoneBookEntry *string `json:"phoneBookEntry,omitempty" tf:"phone_book_entry"`
	// +optional
	PrivateEndpointIPAddress *string `json:"privateEndpointIPAddress,omitempty" tf:"private_endpoint_ip_address"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	StaticJumpHostIPAddresses []string `json:"staticJumpHostIPAddresses,omitempty" tf:"static_jump_host_ip_addresses"`
	// +optional
	SystemTags     map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	TargetSubnetID *string           `json:"targetSubnetID" tf:"target_subnet_id"`
	// +optional
	TargetVcnID *string `json:"targetVcnID,omitempty" tf:"target_vcn_id"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type BastionStatus struct {
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

// BastionList is a list of Bastions
type BastionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Bastion CRD objects
	Items []Bastion `json:"items,omitempty"`
}
