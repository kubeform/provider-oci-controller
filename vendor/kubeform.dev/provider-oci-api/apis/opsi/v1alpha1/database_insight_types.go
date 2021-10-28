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

type DatabaseInsight struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseInsightSpec   `json:"spec,omitempty"`
	Status            DatabaseInsightStatus `json:"status,omitempty"`
}

type DatabaseInsightSpec struct {
	State *DatabaseInsightSpecResource `json:"state,omitempty" tf:"-"`

	Resource DatabaseInsightSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DatabaseInsightSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DatabaseDisplayName *string `json:"databaseDisplayName,omitempty" tf:"database_display_name"`
	// +optional
	DatabaseID *string `json:"databaseID,omitempty" tf:"database_id"`
	// +optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name"`
	// +optional
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type"`
	// +optional
	DatabaseVersion *string `json:"databaseVersion,omitempty" tf:"database_version"`
	// +optional
	DefinedTags               map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	EnterpriseManagerBridgeID *string           `json:"enterpriseManagerBridgeID" tf:"enterprise_manager_bridge_id"`
	// +optional
	EnterpriseManagerEntityDisplayName *string `json:"enterpriseManagerEntityDisplayName,omitempty" tf:"enterprise_manager_entity_display_name"`
	EnterpriseManagerEntityIdentifier  *string `json:"enterpriseManagerEntityIdentifier" tf:"enterprise_manager_entity_identifier"`
	// +optional
	EnterpriseManagerEntityName *string `json:"enterpriseManagerEntityName,omitempty" tf:"enterprise_manager_entity_name"`
	// +optional
	EnterpriseManagerEntityType *string `json:"enterpriseManagerEntityType,omitempty" tf:"enterprise_manager_entity_type"`
	EnterpriseManagerIdentifier *string `json:"enterpriseManagerIdentifier" tf:"enterprise_manager_identifier"`
	EntitySource                *string `json:"entitySource" tf:"entity_source"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	ProcessorCount *int64 `json:"processorCount,omitempty" tf:"processor_count"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type DatabaseInsightStatus struct {
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

// DatabaseInsightList is a list of DatabaseInsights
type DatabaseInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatabaseInsight CRD objects
	Items []DatabaseInsight `json:"items,omitempty"`
}
