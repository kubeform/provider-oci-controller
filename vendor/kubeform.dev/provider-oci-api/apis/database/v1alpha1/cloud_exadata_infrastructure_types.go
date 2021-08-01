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

type CloudExadataInfrastructure struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudExadataInfrastructureSpec   `json:"spec,omitempty"`
	Status            CloudExadataInfrastructureStatus `json:"status,omitempty"`
}

type CloudExadataInfrastructureSpecCustomerContacts struct {
	// +optional
	Email *string `json:"email,omitempty" tf:"email"`
}

type CloudExadataInfrastructureSpecMaintenanceWindowDaysOfWeek struct {
	Name *string `json:"name" tf:"name"`
}

type CloudExadataInfrastructureSpecMaintenanceWindowMonths struct {
	Name *string `json:"name" tf:"name"`
}

type CloudExadataInfrastructureSpecMaintenanceWindow struct {
	// +optional
	DaysOfWeek []CloudExadataInfrastructureSpecMaintenanceWindowDaysOfWeek `json:"daysOfWeek,omitempty" tf:"days_of_week"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	HoursOfDay []int64 `json:"hoursOfDay,omitempty" tf:"hours_of_day"`
	// +optional
	LeadTimeInWeeks *int64 `json:"leadTimeInWeeks,omitempty" tf:"lead_time_in_weeks"`
	// +optional
	Months     []CloudExadataInfrastructureSpecMaintenanceWindowMonths `json:"months,omitempty" tf:"months"`
	Preference *string                                                 `json:"preference" tf:"preference"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	// +kubebuilder:validation:MinItems=1
	WeeksOfMonth []int64 `json:"weeksOfMonth,omitempty" tf:"weeks_of_month"`
}

type CloudExadataInfrastructureSpec struct {
	State *CloudExadataInfrastructureSpecResource `json:"state,omitempty" tf:"-"`

	Resource CloudExadataInfrastructureSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CloudExadataInfrastructureSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AvailabilityDomain *string `json:"availabilityDomain" tf:"availability_domain"`
	// +optional
	AvailableStorageSizeInGbs *int64  `json:"availableStorageSizeInGbs,omitempty" tf:"available_storage_size_in_gbs"`
	CompartmentID             *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	ComputeCount *int64 `json:"computeCount,omitempty" tf:"compute_count"`
	// +optional
	CustomerContacts []CloudExadataInfrastructureSpecCustomerContacts `json:"customerContacts,omitempty" tf:"customer_contacts"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	DisplayName *string           `json:"displayName" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	LastMaintenanceRunID *string `json:"lastMaintenanceRunID,omitempty" tf:"last_maintenance_run_id"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	MaintenanceWindow *CloudExadataInfrastructureSpecMaintenanceWindow `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`
	// +optional
	NextMaintenanceRunID *string `json:"nextMaintenanceRunID,omitempty" tf:"next_maintenance_run_id"`
	Shape                *string `json:"shape" tf:"shape"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	StorageCount *int64 `json:"storageCount,omitempty" tf:"storage_count"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TotalStorageSizeInGbs *int64 `json:"totalStorageSizeInGbs,omitempty" tf:"total_storage_size_in_gbs"`
}

type CloudExadataInfrastructureStatus struct {
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

// CloudExadataInfrastructureList is a list of CloudExadataInfrastructures
type CloudExadataInfrastructureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudExadataInfrastructure CRD objects
	Items []CloudExadataInfrastructure `json:"items,omitempty"`
}
