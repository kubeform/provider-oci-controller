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

type ManagementManagedDatabaseGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementManagedDatabaseGroupSpec   `json:"spec,omitempty"`
	Status            ManagementManagedDatabaseGroupStatus `json:"status,omitempty"`
}

type ManagementManagedDatabaseGroupSpecManagedDatabases struct {
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	DatabaseSubType *string `json:"databaseSubType,omitempty" tf:"database_sub_type"`
	// +optional
	DatabaseType *string `json:"databaseType,omitempty" tf:"database_type"`
	// +optional
	DeploymentType *string `json:"deploymentType,omitempty" tf:"deployment_type"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	TimeAdded *string `json:"timeAdded,omitempty" tf:"time_added"`
}

type ManagementManagedDatabaseGroupSpec struct {
	State *ManagementManagedDatabaseGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagementManagedDatabaseGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagementManagedDatabaseGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	ManagedDatabases []ManagementManagedDatabaseGroupSpecManagedDatabases `json:"managedDatabases,omitempty" tf:"managed_databases"`
	Name             *string                                              `json:"name" tf:"name"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type ManagementManagedDatabaseGroupStatus struct {
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

// ManagementManagedDatabaseGroupList is a list of ManagementManagedDatabaseGroups
type ManagementManagedDatabaseGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagementManagedDatabaseGroup CRD objects
	Items []ManagementManagedDatabaseGroup `json:"items,omitempty"`
}
