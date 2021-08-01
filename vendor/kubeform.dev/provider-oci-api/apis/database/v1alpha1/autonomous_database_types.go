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

type AutonomousDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutonomousDatabaseSpec   `json:"spec,omitempty"`
	Status            AutonomousDatabaseStatus `json:"status,omitempty"`
}

type AutonomousDatabaseSpecApexDetails struct {
	// +optional
	ApexVersion *string `json:"apexVersion,omitempty" tf:"apex_version"`
	// +optional
	OrdsVersion *string `json:"ordsVersion,omitempty" tf:"ords_version"`
}

type AutonomousDatabaseSpecBackupConfig struct {
	// +optional
	ManualBackupBucketName *string `json:"manualBackupBucketName,omitempty" tf:"manual_backup_bucket_name"`
	// +optional
	ManualBackupType *string `json:"manualBackupType,omitempty" tf:"manual_backup_type"`
}

type AutonomousDatabaseSpecConnectionStrings struct {
	// +optional
	AllConnectionStrings map[string]string `json:"allConnectionStrings,omitempty" tf:"all_connection_strings"`
	// +optional
	Dedicated *string `json:"dedicated,omitempty" tf:"dedicated"`
	// +optional
	High *string `json:"high,omitempty" tf:"high"`
	// +optional
	Low *string `json:"low,omitempty" tf:"low"`
	// +optional
	Medium *string `json:"medium,omitempty" tf:"medium"`
}

type AutonomousDatabaseSpecConnectionUrls struct {
	// +optional
	ApexURL *string `json:"apexURL,omitempty" tf:"apex_url"`
	// +optional
	GraphStudioURL *string `json:"graphStudioURL,omitempty" tf:"graph_studio_url"`
	// +optional
	MachineLearningUserManagementURL *string `json:"machineLearningUserManagementURL,omitempty" tf:"machine_learning_user_management_url"`
	// +optional
	SqlDevWebURL *string `json:"sqlDevWebURL,omitempty" tf:"sql_dev_web_url"`
}

type AutonomousDatabaseSpecCustomerContacts struct {
	// +optional
	Email *string `json:"email,omitempty" tf:"email"`
}

type AutonomousDatabaseSpecKeyHistoryEntry struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	TimeActivated *string `json:"timeActivated,omitempty" tf:"time_activated"`
	// +optional
	VaultID *string `json:"vaultID,omitempty" tf:"vault_id"`
}

type AutonomousDatabaseSpecStandbyDb struct {
	// +optional
	LagTimeInSeconds *int64 `json:"lagTimeInSeconds,omitempty" tf:"lag_time_in_seconds"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type AutonomousDatabaseSpec struct {
	State *AutonomousDatabaseSpecResource `json:"state,omitempty" tf:"-"`

	Resource AutonomousDatabaseSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type AutonomousDatabaseSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdminPassword *string `json:"-" sensitive:"true" tf:"admin_password"`
	// +optional
	ApexDetails *AutonomousDatabaseSpecApexDetails `json:"apexDetails,omitempty" tf:"apex_details"`
	// +optional
	ArePrimaryWhitelistedIPSUsed *bool `json:"arePrimaryWhitelistedIPSUsed,omitempty" tf:"are_primary_whitelisted_ips_used"`
	// +optional
	AutonomousContainerDatabaseID *string `json:"autonomousContainerDatabaseID,omitempty" tf:"autonomous_container_database_id"`
	// +optional
	AutonomousDatabaseBackupID *string `json:"autonomousDatabaseBackupID,omitempty" tf:"autonomous_database_backup_id"`
	// +optional
	AutonomousDatabaseID *string `json:"autonomousDatabaseID,omitempty" tf:"autonomous_database_id"`
	// +optional
	AvailableUpgradeVersions []string `json:"availableUpgradeVersions,omitempty" tf:"available_upgrade_versions"`
	// +optional
	BackupConfig *AutonomousDatabaseSpecBackupConfig `json:"backupConfig,omitempty" tf:"backup_config"`
	// +optional
	CloneType     *string `json:"cloneType,omitempty" tf:"clone_type"`
	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	ConnectionStrings *AutonomousDatabaseSpecConnectionStrings `json:"connectionStrings,omitempty" tf:"connection_strings"`
	// +optional
	ConnectionUrls *AutonomousDatabaseSpecConnectionUrls `json:"connectionUrls,omitempty" tf:"connection_urls"`
	// +optional
	CpuCoreCount *int64 `json:"cpuCoreCount,omitempty" tf:"cpu_core_count"`
	// +optional
	CustomerContacts []AutonomousDatabaseSpecCustomerContacts `json:"customerContacts,omitempty" tf:"customer_contacts"`
	// +optional
	DataSafeStatus *string `json:"dataSafeStatus,omitempty" tf:"data_safe_status"`
	// +optional
	DataStorageSizeInGb *int64 `json:"dataStorageSizeInGb,omitempty" tf:"data_storage_size_in_gb"`
	// +optional
	DataStorageSizeInTbs *int64  `json:"dataStorageSizeInTbs,omitempty" tf:"data_storage_size_in_tbs"`
	DbName               *string `json:"dbName" tf:"db_name"`
	// +optional
	DbVersion *string `json:"dbVersion,omitempty" tf:"db_version"`
	// +optional
	DbWorkload *string `json:"dbWorkload,omitempty" tf:"db_workload"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FailedDataRecoveryInSeconds *int64 `json:"failedDataRecoveryInSeconds,omitempty" tf:"failed_data_recovery_in_seconds"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	InfrastructureType *string `json:"infrastructureType,omitempty" tf:"infrastructure_type"`
	// +optional
	IsAccessControlEnabled *bool `json:"isAccessControlEnabled,omitempty" tf:"is_access_control_enabled"`
	// +optional
	IsAutoScalingEnabled *bool `json:"isAutoScalingEnabled,omitempty" tf:"is_auto_scaling_enabled"`
	// +optional
	IsDataGuardEnabled *bool `json:"isDataGuardEnabled,omitempty" tf:"is_data_guard_enabled"`
	// +optional
	IsDedicated *bool `json:"isDedicated,omitempty" tf:"is_dedicated"`
	// +optional
	IsFreeTier *bool `json:"isFreeTier,omitempty" tf:"is_free_tier"`
	// +optional
	IsPreview *bool `json:"isPreview,omitempty" tf:"is_preview"`
	// +optional
	IsPreviewVersionWithServiceTermsAccepted *bool `json:"isPreviewVersionWithServiceTermsAccepted,omitempty" tf:"is_preview_version_with_service_terms_accepted"`
	// +optional
	IsRefreshableClone *bool `json:"isRefreshableClone,omitempty" tf:"is_refreshable_clone"`
	// +optional
	KeyHistoryEntry []AutonomousDatabaseSpecKeyHistoryEntry `json:"keyHistoryEntry,omitempty" tf:"key_history_entry"`
	// +optional
	KeyStoreID *string `json:"keyStoreID,omitempty" tf:"key_store_id"`
	// +optional
	KeyStoreWalletName *string `json:"keyStoreWalletName,omitempty" tf:"key_store_wallet_name"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	KmsKeyLifecycleDetails *string `json:"kmsKeyLifecycleDetails,omitempty" tf:"kms_key_lifecycle_details"`
	// +optional
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	NsgIDS []string `json:"nsgIDS,omitempty" tf:"nsg_ids"`
	// +optional
	OcpuCount *float64 `json:"ocpuCount,omitempty" tf:"ocpu_count"`
	// +optional
	OpenMode *string `json:"openMode,omitempty" tf:"open_mode"`
	// +optional
	OperationsInsightsStatus *string `json:"operationsInsightsStatus,omitempty" tf:"operations_insights_status"`
	// +optional
	PermissionLevel *string `json:"permissionLevel,omitempty" tf:"permission_level"`
	// +optional
	PrivateEndpoint *string `json:"privateEndpoint,omitempty" tf:"private_endpoint"`
	// +optional
	PrivateEndpointIP *string `json:"privateEndpointIP,omitempty" tf:"private_endpoint_ip"`
	// +optional
	PrivateEndpointLabel *string `json:"privateEndpointLabel,omitempty" tf:"private_endpoint_label"`
	// +optional
	RefreshableMode *string `json:"refreshableMode,omitempty" tf:"refreshable_mode"`
	// +optional
	RefreshableStatus *string `json:"refreshableStatus,omitempty" tf:"refreshable_status"`
	// +optional
	Role *string `json:"role,omitempty" tf:"role"`
	// +optional
	RotateKeyTrigger *bool `json:"rotateKeyTrigger,omitempty" tf:"rotate_key_trigger"`
	// +optional
	ServiceConsoleURL *string `json:"serviceConsoleURL,omitempty" tf:"service_console_url"`
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
	// +optional
	SourceID *string `json:"sourceID,omitempty" tf:"source_id"`
	// +optional
	StandbyDb *AutonomousDatabaseSpecStandbyDb `json:"standbyDb,omitempty" tf:"standby_db"`
	// +optional
	StandbyWhitelistedIPS []string `json:"standbyWhitelistedIPS,omitempty" tf:"standby_whitelisted_ips"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	SwitchoverTo *string `json:"switchoverTo,omitempty" tf:"switchover_to"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeDeletionOfFreeAutonomousDatabase *string `json:"timeDeletionOfFreeAutonomousDatabase,omitempty" tf:"time_deletion_of_free_autonomous_database"`
	// +optional
	TimeMaintenanceBegin *string `json:"timeMaintenanceBegin,omitempty" tf:"time_maintenance_begin"`
	// +optional
	TimeMaintenanceEnd *string `json:"timeMaintenanceEnd,omitempty" tf:"time_maintenance_end"`
	// +optional
	TimeOfLastFailover *string `json:"timeOfLastFailover,omitempty" tf:"time_of_last_failover"`
	// +optional
	TimeOfLastRefresh *string `json:"timeOfLastRefresh,omitempty" tf:"time_of_last_refresh"`
	// +optional
	TimeOfLastRefreshPoint *string `json:"timeOfLastRefreshPoint,omitempty" tf:"time_of_last_refresh_point"`
	// +optional
	TimeOfLastSwitchover *string `json:"timeOfLastSwitchover,omitempty" tf:"time_of_last_switchover"`
	// +optional
	TimeOfNextRefresh *string `json:"timeOfNextRefresh,omitempty" tf:"time_of_next_refresh"`
	// +optional
	TimeReclamationOfFreeAutonomousDatabase *string `json:"timeReclamationOfFreeAutonomousDatabase,omitempty" tf:"time_reclamation_of_free_autonomous_database"`
	// +optional
	Timestamp *string `json:"timestamp,omitempty" tf:"timestamp"`
	// +optional
	UsedDataStorageSizeInTbs *int64 `json:"usedDataStorageSizeInTbs,omitempty" tf:"used_data_storage_size_in_tbs"`
	// +optional
	VaultID *string `json:"vaultID,omitempty" tf:"vault_id"`
	// +optional
	WhitelistedIPS []string `json:"whitelistedIPS,omitempty" tf:"whitelisted_ips"`
}

type AutonomousDatabaseStatus struct {
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

// AutonomousDatabaseList is a list of AutonomousDatabases
type AutonomousDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutonomousDatabase CRD objects
	Items []AutonomousDatabase `json:"items,omitempty"`
}
