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

type Zone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZoneSpec   `json:"spec,omitempty"`
	Status            ZoneStatus `json:"status,omitempty"`
}

type ZoneSpecExternalMasters struct {
	Address *string `json:"address" tf:"address"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	TsigKeyID *string `json:"tsigKeyID,omitempty" tf:"tsig_key_id"`
}

type ZoneSpecNameservers struct {
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
}

type ZoneSpec struct {
	State *ZoneSpecResource `json:"state,omitempty" tf:"-"`

	Resource ZoneSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ZoneSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	ExternalMasters []ZoneSpecExternalMasters `json:"externalMasters,omitempty" tf:"external_masters"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	IsProtected *bool   `json:"isProtected,omitempty" tf:"is_protected"`
	Name        *string `json:"name" tf:"name"`
	// +optional
	Nameservers []ZoneSpecNameservers `json:"nameservers,omitempty" tf:"nameservers"`
	// +optional
	Scope *string `json:"scope,omitempty" tf:"scope"`
	// +optional
	Self *string `json:"self,omitempty" tf:"self"`
	// +optional
	Serial *int64 `json:"serial,omitempty" tf:"serial"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	ViewID   *string `json:"viewID,omitempty" tf:"view_id"`
	ZoneType *string `json:"zoneType" tf:"zone_type"`
}

type ZoneStatus struct {
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

// ZoneList is a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Zone CRD objects
	Items []Zone `json:"items,omitempty"`
}
