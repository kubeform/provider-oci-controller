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

type AutonomousContainerDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutonomousContainerDatabaseSpec   `json:"spec,omitempty"`
	Status            AutonomousContainerDatabaseStatus `json:"status,omitempty"`
}

type AutonomousContainerDatabaseSpecBackupConfigBackupDestinationDetails struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InternetProxy *string `json:"internetProxy,omitempty" tf:"internet_proxy"`
	Type          *string `json:"type" tf:"type"`
	// +optional
	VpcPassword *string `json:"-" sensitive:"true" tf:"vpc_password"`
	// +optional
	VpcUser *string `json:"vpcUser,omitempty" tf:"vpc_user"`
}

type AutonomousContainerDatabaseSpecBackupConfig struct {
	// +optional
	BackupDestinationDetails *AutonomousContainerDatabaseSpecBackupConfigBackupDestinationDetails `json:"backupDestinationDetails,omitempty" tf:"backup_destination_details"`
	// +optional
	RecoveryWindowInDays *int64 `json:"recoveryWindowInDays,omitempty" tf:"recovery_window_in_days"`
}

type AutonomousContainerDatabaseSpecMaintenanceWindowDaysOfWeek struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type AutonomousContainerDatabaseSpecMaintenanceWindowMonths struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type AutonomousContainerDatabaseSpecMaintenanceWindow struct {
	// +optional
	DaysOfWeek []AutonomousContainerDatabaseSpecMaintenanceWindowDaysOfWeek `json:"daysOfWeek,omitempty" tf:"days_of_week"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	HoursOfDay []int64 `json:"hoursOfDay,omitempty" tf:"hours_of_day"`
	// +optional
	LeadTimeInWeeks *int64 `json:"leadTimeInWeeks,omitempty" tf:"lead_time_in_weeks"`
	// +optional
	Months []AutonomousContainerDatabaseSpecMaintenanceWindowMonths `json:"months,omitempty" tf:"months"`
	// +optional
	Preference *string `json:"preference,omitempty" tf:"preference"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	// +kubebuilder:validation:MinItems=1
	WeeksOfMonth []int64 `json:"weeksOfMonth,omitempty" tf:"weeks_of_month"`
}

type AutonomousContainerDatabaseSpecMaintenanceWindowDetailsDaysOfWeek struct {
	Name *string `json:"name" tf:"name"`
}

type AutonomousContainerDatabaseSpecMaintenanceWindowDetailsMonths struct {
	Name *string `json:"name" tf:"name"`
}

type AutonomousContainerDatabaseSpecMaintenanceWindowDetails struct {
	// +optional
	DaysOfWeek []AutonomousContainerDatabaseSpecMaintenanceWindowDetailsDaysOfWeek `json:"daysOfWeek,omitempty" tf:"days_of_week"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	HoursOfDay []int64 `json:"hoursOfDay,omitempty" tf:"hours_of_day"`
	// +optional
	LeadTimeInWeeks *int64 `json:"leadTimeInWeeks,omitempty" tf:"lead_time_in_weeks"`
	// +optional
	Months     []AutonomousContainerDatabaseSpecMaintenanceWindowDetailsMonths `json:"months,omitempty" tf:"months"`
	Preference *string                                                         `json:"preference" tf:"preference"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	// +kubebuilder:validation:MinItems=1
	WeeksOfMonth []int64 `json:"weeksOfMonth,omitempty" tf:"weeks_of_month"`
}

type AutonomousContainerDatabaseSpecPeerAutonomousContainerDatabaseBackupConfigBackupDestinationDetails struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InternetProxy *string `json:"internetProxy,omitempty" tf:"internet_proxy"`
	Type          *string `json:"type" tf:"type"`
	// +optional
	VpcPassword *string `json:"-" sensitive:"true" tf:"vpc_password"`
	// +optional
	VpcUser *string `json:"vpcUser,omitempty" tf:"vpc_user"`
}

type AutonomousContainerDatabaseSpecPeerAutonomousContainerDatabaseBackupConfig struct {
	// +optional
	BackupDestinationDetails []AutonomousContainerDatabaseSpecPeerAutonomousContainerDatabaseBackupConfigBackupDestinationDetails `json:"backupDestinationDetails,omitempty" tf:"backup_destination_details"`
	// +optional
	RecoveryWindowInDays *int64 `json:"recoveryWindowInDays,omitempty" tf:"recovery_window_in_days"`
}

type AutonomousContainerDatabaseSpec struct {
	State *AutonomousContainerDatabaseSpecResource `json:"state,omitempty" tf:"-"`

	Resource AutonomousContainerDatabaseSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AutonomousContainerDatabaseSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutonomousExadataInfrastructureID *string `json:"autonomousExadataInfrastructureID,omitempty" tf:"autonomous_exadata_infrastructure_id"`
	// +optional
	AutonomousVmClusterID *string `json:"autonomousVmClusterID,omitempty" tf:"autonomous_vm_cluster_id"`
	// +optional
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain"`
	// +optional
	BackupConfig *AutonomousContainerDatabaseSpecBackupConfig `json:"backupConfig,omitempty" tf:"backup_config"`
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	DbUniqueName *string `json:"dbUniqueName,omitempty" tf:"db_unique_name"`
	// +optional
	DbVersion *string `json:"dbVersion,omitempty" tf:"db_version"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	DisplayName *string           `json:"displayName" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	InfrastructureType *string `json:"infrastructureType,omitempty" tf:"infrastructure_type"`
	// +optional
	KeyStoreID *string `json:"keyStoreID,omitempty" tf:"key_store_id"`
	// +optional
	KeyStoreWalletName *string `json:"keyStoreWalletName,omitempty" tf:"key_store_wallet_name"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	LastMaintenanceRunID *string `json:"lastMaintenanceRunID,omitempty" tf:"last_maintenance_run_id"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	MaintenanceWindow *AutonomousContainerDatabaseSpecMaintenanceWindow `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`
	// +optional
	MaintenanceWindowDetails *AutonomousContainerDatabaseSpecMaintenanceWindowDetails `json:"maintenanceWindowDetails,omitempty" tf:"maintenance_window_details"`
	// +optional
	NextMaintenanceRunID *string `json:"nextMaintenanceRunID,omitempty" tf:"next_maintenance_run_id"`
	// +optional
	PatchID    *string `json:"patchID,omitempty" tf:"patch_id"`
	PatchModel *string `json:"patchModel" tf:"patch_model"`
	// +optional
	PeerAutonomousContainerDatabaseBackupConfig *AutonomousContainerDatabaseSpecPeerAutonomousContainerDatabaseBackupConfig `json:"peerAutonomousContainerDatabaseBackupConfig,omitempty" tf:"peer_autonomous_container_database_backup_config"`
	// +optional
	PeerAutonomousContainerDatabaseCompartmentID *string `json:"peerAutonomousContainerDatabaseCompartmentID,omitempty" tf:"peer_autonomous_container_database_compartment_id"`
	// +optional
	PeerAutonomousContainerDatabaseDisplayName *string `json:"peerAutonomousContainerDatabaseDisplayName,omitempty" tf:"peer_autonomous_container_database_display_name"`
	// +optional
	PeerAutonomousExadataInfrastructureID *string `json:"peerAutonomousExadataInfrastructureID,omitempty" tf:"peer_autonomous_exadata_infrastructure_id"`
	// +optional
	PeerAutonomousVmClusterID *string `json:"peerAutonomousVmClusterID,omitempty" tf:"peer_autonomous_vm_cluster_id"`
	// +optional
	PeerDbUniqueName *string `json:"peerDbUniqueName,omitempty" tf:"peer_db_unique_name"`
	// +optional
	ProtectionMode *string `json:"protectionMode,omitempty" tf:"protection_mode"`
	// +optional
	Role *string `json:"role,omitempty" tf:"role"`
	// +optional
	RotateKeyTrigger *bool `json:"rotateKeyTrigger,omitempty" tf:"rotate_key_trigger"`
	// +optional
	ServiceLevelAgreementType *string `json:"serviceLevelAgreementType,omitempty" tf:"service_level_agreement_type"`
	// +optional
	StandbyMaintenanceBufferInDays *int64 `json:"standbyMaintenanceBufferInDays,omitempty" tf:"standby_maintenance_buffer_in_days"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	VaultID *string `json:"vaultID,omitempty" tf:"vault_id"`
}

type AutonomousContainerDatabaseStatus struct {
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

// AutonomousContainerDatabaseList is a list of AutonomousContainerDatabases
type AutonomousContainerDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutonomousContainerDatabase CRD objects
	Items []AutonomousContainerDatabase `json:"items,omitempty"`
}
