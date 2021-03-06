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

type AnalyticsLogAnalyticsEntity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyticsLogAnalyticsEntitySpec   `json:"spec,omitempty"`
	Status            AnalyticsLogAnalyticsEntityStatus `json:"status,omitempty"`
}

type AnalyticsLogAnalyticsEntitySpec struct {
	State *AnalyticsLogAnalyticsEntitySpecResource `json:"state,omitempty" tf:"-"`

	Resource AnalyticsLogAnalyticsEntitySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AnalyticsLogAnalyticsEntitySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AreLogsCollected *bool `json:"areLogsCollected,omitempty" tf:"are_logs_collected"`
	// +optional
	CloudResourceID *string `json:"cloudResourceID,omitempty" tf:"cloud_resource_id"`
	CompartmentID   *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	EntityTypeInternalName *string `json:"entityTypeInternalName,omitempty" tf:"entity_type_internal_name"`
	EntityTypeName         *string `json:"entityTypeName" tf:"entity_type_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	ManagementAgentCompartmentID *string `json:"managementAgentCompartmentID,omitempty" tf:"management_agent_compartment_id"`
	// +optional
	ManagementAgentDisplayName *string `json:"managementAgentDisplayName,omitempty" tf:"management_agent_display_name"`
	// +optional
	ManagementAgentID *string `json:"managementAgentID,omitempty" tf:"management_agent_id"`
	Name              *string `json:"name" tf:"name"`
	Namespace         *string `json:"namespace" tf:"namespace"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
	// +optional
	SourceID *string `json:"sourceID,omitempty" tf:"source_id"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	TimezoneRegion *string `json:"timezoneRegion,omitempty" tf:"timezone_region"`
}

type AnalyticsLogAnalyticsEntityStatus struct {
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

// AnalyticsLogAnalyticsEntityList is a list of AnalyticsLogAnalyticsEntitys
type AnalyticsLogAnalyticsEntityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AnalyticsLogAnalyticsEntity CRD objects
	Items []AnalyticsLogAnalyticsEntity `json:"items,omitempty"`
}
