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

type VolumeBackupPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeBackupPolicySpec   `json:"spec,omitempty"`
	Status            VolumeBackupPolicyStatus `json:"status,omitempty"`
}

type VolumeBackupPolicySpecSchedules struct {
	BackupType *string `json:"backupType" tf:"backup_type"`
	// +optional
	DayOfMonth *int64 `json:"dayOfMonth,omitempty" tf:"day_of_month"`
	// +optional
	DayOfWeek *string `json:"dayOfWeek,omitempty" tf:"day_of_week"`
	// +optional
	HourOfDay *int64 `json:"hourOfDay,omitempty" tf:"hour_of_day"`
	// +optional
	Month *string `json:"month,omitempty" tf:"month"`
	// +optional
	OffsetSeconds *int64 `json:"offsetSeconds,omitempty" tf:"offset_seconds"`
	// +optional
	OffsetType       *string `json:"offsetType,omitempty" tf:"offset_type"`
	Period           *string `json:"period" tf:"period"`
	RetentionSeconds *int64  `json:"retentionSeconds" tf:"retention_seconds"`
	// +optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone"`
}

type VolumeBackupPolicySpec struct {
	State *VolumeBackupPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource VolumeBackupPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VolumeBackupPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DestinationRegion *string `json:"destinationRegion,omitempty" tf:"destination_region"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	Schedules []VolumeBackupPolicySpecSchedules `json:"schedules,omitempty" tf:"schedules"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
}

type VolumeBackupPolicyStatus struct {
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

// VolumeBackupPolicyList is a list of VolumeBackupPolicys
type VolumeBackupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VolumeBackupPolicy CRD objects
	Items []VolumeBackupPolicy `json:"items,omitempty"`
}
