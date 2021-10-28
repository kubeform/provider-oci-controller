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

type DatabaseUpgrade struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseUpgradeSpec   `json:"spec,omitempty"`
	Status            DatabaseUpgradeStatus `json:"status,omitempty"`
}

type DatabaseUpgradeSpecConnectionStrings struct {
	// +optional
	AllConnectionStrings map[string]string `json:"allConnectionStrings,omitempty" tf:"all_connection_strings"`
	// +optional
	CdbDefault *string `json:"cdbDefault,omitempty" tf:"cdb_default"`
	// +optional
	CdbIPDefault *string `json:"cdbIPDefault,omitempty" tf:"cdb_ip_default"`
}

type DatabaseUpgradeSpecDatabaseUpgradeSourceDetails struct {
	// +optional
	DatabaseSoftwareImageID *string `json:"databaseSoftwareImageID,omitempty" tf:"database_software_image_id"`
	// +optional
	DbVersion *string `json:"dbVersion,omitempty" tf:"db_version"`
	// +optional
	Options *string `json:"options,omitempty" tf:"options"`
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
}

type DatabaseUpgradeSpecDbBackupConfigBackupDestinationDetails struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InternetProxy *string `json:"internetProxy,omitempty" tf:"internet_proxy"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	VpcPassword *string `json:"vpcPassword,omitempty" tf:"vpc_password"`
	// +optional
	VpcUser *string `json:"vpcUser,omitempty" tf:"vpc_user"`
}

type DatabaseUpgradeSpecDbBackupConfig struct {
	// +optional
	AutoBackupEnabled *bool `json:"autoBackupEnabled,omitempty" tf:"auto_backup_enabled"`
	// +optional
	AutoBackupWindow *string `json:"autoBackupWindow,omitempty" tf:"auto_backup_window"`
	// +optional
	BackupDestinationDetails []DatabaseUpgradeSpecDbBackupConfigBackupDestinationDetails `json:"backupDestinationDetails,omitempty" tf:"backup_destination_details"`
	// +optional
	RecoveryWindowInDays *int64 `json:"recoveryWindowInDays,omitempty" tf:"recovery_window_in_days"`
}

type DatabaseUpgradeSpec struct {
	State *DatabaseUpgradeSpecResource `json:"state,omitempty" tf:"-"`

	Resource DatabaseUpgradeSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DatabaseUpgradeSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Action *string `json:"action" tf:"action"`
	// +optional
	CharacterSet *string `json:"characterSet,omitempty" tf:"character_set"`
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	ConnectionStrings *DatabaseUpgradeSpecConnectionStrings `json:"connectionStrings,omitempty" tf:"connection_strings"`
	DatabaseID        *string                               `json:"databaseID" tf:"database_id"`
	// +optional
	DatabaseSoftwareImageID *string `json:"databaseSoftwareImageID,omitempty" tf:"database_software_image_id"`
	// +optional
	DatabaseUpgradeSourceDetails *DatabaseUpgradeSpecDatabaseUpgradeSourceDetails `json:"databaseUpgradeSourceDetails,omitempty" tf:"database_upgrade_source_details"`
	// +optional
	DbBackupConfig *DatabaseUpgradeSpecDbBackupConfig `json:"dbBackupConfig,omitempty" tf:"db_backup_config"`
	// +optional
	DbHomeID *string `json:"dbHomeID,omitempty" tf:"db_home_id"`
	// +optional
	DbName *string `json:"dbName,omitempty" tf:"db_name"`
	// +optional
	DbSystemID *string `json:"dbSystemID,omitempty" tf:"db_system_id"`
	// +optional
	DbUniqueName *string `json:"dbUniqueName,omitempty" tf:"db_unique_name"`
	// +optional
	DbWorkload *string `json:"dbWorkload,omitempty" tf:"db_workload"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	LastBackupTimestamp *string `json:"lastBackupTimestamp,omitempty" tf:"last_backup_timestamp"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	NcharacterSet *string `json:"ncharacterSet,omitempty" tf:"ncharacter_set"`
	// +optional
	PdbName *string `json:"pdbName,omitempty" tf:"pdb_name"`
	// +optional
	SourceDatabasePointInTimeRecoveryTimestamp *string `json:"sourceDatabasePointInTimeRecoveryTimestamp,omitempty" tf:"source_database_point_in_time_recovery_timestamp"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	VmClusterID *string `json:"vmClusterID,omitempty" tf:"vm_cluster_id"`
}

type DatabaseUpgradeStatus struct {
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

// DatabaseUpgradeList is a list of DatabaseUpgrades
type DatabaseUpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatabaseUpgrade CRD objects
	Items []DatabaseUpgrade `json:"items,omitempty"`
}
