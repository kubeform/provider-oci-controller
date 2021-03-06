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

type VolumeGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeGroupSpec   `json:"spec,omitempty"`
	Status            VolumeGroupStatus `json:"status,omitempty"`
}

type VolumeGroupSpecSourceDetails struct {
	Type *string `json:"type" tf:"type"`
	// +optional
	VolumeGroupBackupID *string `json:"volumeGroupBackupID,omitempty" tf:"volume_group_backup_id"`
	// +optional
	VolumeGroupID *string `json:"volumeGroupID,omitempty" tf:"volume_group_id"`
	// +optional
	VolumeGroupReplicaID *string `json:"volumeGroupReplicaID,omitempty" tf:"volume_group_replica_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=64
	VolumeIDS []string `json:"volumeIDS,omitempty" tf:"volume_ids"`
}

type VolumeGroupSpecVolumeGroupReplicas struct {
	AvailabilityDomain *string `json:"availabilityDomain" tf:"availability_domain"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	VolumeGroupReplicaID *string `json:"volumeGroupReplicaID,omitempty" tf:"volume_group_replica_id"`
}

type VolumeGroupSpec struct {
	State *VolumeGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource VolumeGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VolumeGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AvailabilityDomain *string `json:"availabilityDomain" tf:"availability_domain"`
	// +optional
	BackupPolicyID *string `json:"backupPolicyID,omitempty" tf:"backup_policy_id"`
	CompartmentID  *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	IsHydrated *bool `json:"isHydrated,omitempty" tf:"is_hydrated"`
	// +optional
	PreserveVolumeReplica *bool `json:"preserveVolumeReplica,omitempty" tf:"preserve_volume_replica"`
	// +optional
	SizeInGbs *string `json:"sizeInGbs,omitempty" tf:"size_in_gbs"`
	// +optional
	SizeInMbs     *string                       `json:"sizeInMbs,omitempty" tf:"size_in_mbs"`
	SourceDetails *VolumeGroupSpecSourceDetails `json:"sourceDetails" tf:"source_details"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	VolumeGroupReplicas []VolumeGroupSpecVolumeGroupReplicas `json:"volumeGroupReplicas,omitempty" tf:"volume_group_replicas"`
	// +optional
	VolumeGroupReplicasDeletion *bool `json:"volumeGroupReplicasDeletion,omitempty" tf:"volume_group_replicas_deletion"`
	// +optional
	VolumeIDS []string `json:"volumeIDS,omitempty" tf:"volume_ids"`
}

type VolumeGroupStatus struct {
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

// VolumeGroupList is a list of VolumeGroups
type VolumeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VolumeGroup CRD objects
	Items []VolumeGroup `json:"items,omitempty"`
}
