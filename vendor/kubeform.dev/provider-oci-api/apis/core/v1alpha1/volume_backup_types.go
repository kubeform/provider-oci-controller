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

type VolumeBackup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeBackupSpec   `json:"spec,omitempty"`
	Status            VolumeBackupStatus `json:"status,omitempty"`
}

type VolumeBackupSpecSourceDetails struct {
	// +optional
	KmsKeyID       *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	Region         *string `json:"region" tf:"region"`
	VolumeBackupID *string `json:"volumeBackupID" tf:"volume_backup_id"`
}

type VolumeBackupSpec struct {
	State *VolumeBackupSpecResource `json:"state,omitempty" tf:"-"`

	Resource VolumeBackupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VolumeBackupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	SizeInGbs *string `json:"sizeInGbs,omitempty" tf:"size_in_gbs"`
	// +optional
	// Deprecated
	SizeInMbs *string `json:"sizeInMbs,omitempty" tf:"size_in_mbs"`
	// +optional
	SourceDetails *VolumeBackupSpecSourceDetails `json:"sourceDetails,omitempty" tf:"source_details"`
	// +optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type"`
	// +optional
	SourceVolumeBackupID *string `json:"sourceVolumeBackupID,omitempty" tf:"source_volume_backup_id"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeRequestReceived *string `json:"timeRequestReceived,omitempty" tf:"time_request_received"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	UniqueSizeInGbs *string `json:"uniqueSizeInGbs,omitempty" tf:"unique_size_in_gbs"`
	// +optional
	// Deprecated
	UniqueSizeInMbs *string `json:"uniqueSizeInMbs,omitempty" tf:"unique_size_in_mbs"`
	// +optional
	VolumeID *string `json:"volumeID,omitempty" tf:"volume_id"`
}

type VolumeBackupStatus struct {
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

// VolumeBackupList is a list of VolumeBackups
type VolumeBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VolumeBackup CRD objects
	Items []VolumeBackup `json:"items,omitempty"`
}
