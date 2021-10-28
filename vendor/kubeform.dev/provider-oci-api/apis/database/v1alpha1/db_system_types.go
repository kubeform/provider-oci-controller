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

type DbSystem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbSystemSpec   `json:"spec,omitempty"`
	Status            DbSystemStatus `json:"status,omitempty"`
}

type DbSystemSpecDbHomeDatabaseConnectionStrings struct {
	// +optional
	AllConnectionStrings map[string]string `json:"allConnectionStrings,omitempty" tf:"all_connection_strings"`
	// +optional
	CdbDefault *string `json:"cdbDefault,omitempty" tf:"cdb_default"`
	// +optional
	CdbIPDefault *string `json:"cdbIPDefault,omitempty" tf:"cdb_ip_default"`
}

type DbSystemSpecDbHomeDatabaseDbBackupConfigBackupDestinationDetails struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type DbSystemSpecDbHomeDatabaseDbBackupConfig struct {
	// +optional
	AutoBackupEnabled *bool `json:"autoBackupEnabled,omitempty" tf:"auto_backup_enabled"`
	// +optional
	AutoBackupWindow *string `json:"autoBackupWindow,omitempty" tf:"auto_backup_window"`
	// +optional
	BackupDestinationDetails []DbSystemSpecDbHomeDatabaseDbBackupConfigBackupDestinationDetails `json:"backupDestinationDetails,omitempty" tf:"backup_destination_details"`
	// +optional
	RecoveryWindowInDays *int64 `json:"recoveryWindowInDays,omitempty" tf:"recovery_window_in_days"`
}

type DbSystemSpecDbHomeDatabase struct {
	AdminPassword *string `json:"-" sensitive:"true" tf:"admin_password"`
	// +optional
	BackupID *string `json:"backupID,omitempty" tf:"backup_id"`
	// +optional
	BackupTdePassword *string `json:"-" sensitive:"true" tf:"backup_tde_password"`
	// +optional
	CharacterSet *string `json:"characterSet,omitempty" tf:"character_set"`
	// +optional
	ConnectionStrings []DbSystemSpecDbHomeDatabaseConnectionStrings `json:"connectionStrings,omitempty" tf:"connection_strings"`
	// +optional
	DatabaseID *string `json:"databaseID,omitempty" tf:"database_id"`
	// +optional
	DatabaseSoftwareImageID *string `json:"databaseSoftwareImageID,omitempty" tf:"database_software_image_id"`
	// +optional
	DbBackupConfig *DbSystemSpecDbHomeDatabaseDbBackupConfig `json:"dbBackupConfig,omitempty" tf:"db_backup_config"`
	// +optional
	DbDomain *string `json:"dbDomain,omitempty" tf:"db_domain"`
	// +optional
	DbName *string `json:"dbName,omitempty" tf:"db_name"`
	// +optional
	DbUniqueName *string `json:"dbUniqueName,omitempty" tf:"db_unique_name"`
	// +optional
	DbWorkload *string `json:"dbWorkload,omitempty" tf:"db_workload"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	NcharacterSet *string `json:"ncharacterSet,omitempty" tf:"ncharacter_set"`
	// +optional
	PdbName *string `json:"pdbName,omitempty" tf:"pdb_name"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TdeWalletPassword *string `json:"-" sensitive:"true" tf:"tde_wallet_password"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeStampForPointInTimeRecovery *string `json:"timeStampForPointInTimeRecovery,omitempty" tf:"time_stamp_for_point_in_time_recovery"`
}

type DbSystemSpecDbHome struct {
	// +optional
	CreateAsync *bool                       `json:"createAsync,omitempty" tf:"create_async"`
	Database    *DbSystemSpecDbHomeDatabase `json:"database" tf:"database"`
	// +optional
	DatabaseSoftwareImageID *string `json:"databaseSoftwareImageID,omitempty" tf:"database_software_image_id"`
	// +optional
	DbHomeLocation *string `json:"dbHomeLocation,omitempty" tf:"db_home_location"`
	// +optional
	DbVersion *string `json:"dbVersion,omitempty" tf:"db_version"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	LastPatchHistoryEntryID *string `json:"lastPatchHistoryEntryID,omitempty" tf:"last_patch_history_entry_id"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
}

type DbSystemSpecDbSystemOptions struct {
	// +optional
	StorageManagement *string `json:"storageManagement,omitempty" tf:"storage_management"`
}

type DbSystemSpecIormConfigCacheDbPlans struct {
	// +optional
	DbName *string `json:"dbName,omitempty" tf:"db_name"`
	// +optional
	FlashCacheLimit *string `json:"flashCacheLimit,omitempty" tf:"flash_cache_limit"`
	// +optional
	Share *int64 `json:"share,omitempty" tf:"share"`
}

type DbSystemSpecIormConfigCache struct {
	// +optional
	DbPlans []DbSystemSpecIormConfigCacheDbPlans `json:"dbPlans,omitempty" tf:"db_plans"`
	// +optional
	DbSystemID *string `json:"dbSystemID,omitempty" tf:"db_system_id"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	Objective *string `json:"objective,omitempty" tf:"objective"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type DbSystemSpecMaintenanceWindowDaysOfWeek struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type DbSystemSpecMaintenanceWindowMonths struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type DbSystemSpecMaintenanceWindow struct {
	// +optional
	DaysOfWeek []DbSystemSpecMaintenanceWindowDaysOfWeek `json:"daysOfWeek,omitempty" tf:"days_of_week"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	HoursOfDay []int64 `json:"hoursOfDay,omitempty" tf:"hours_of_day"`
	// +optional
	LeadTimeInWeeks *int64 `json:"leadTimeInWeeks,omitempty" tf:"lead_time_in_weeks"`
	// +optional
	Months []DbSystemSpecMaintenanceWindowMonths `json:"months,omitempty" tf:"months"`
	// +optional
	Preference *string `json:"preference,omitempty" tf:"preference"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	// +kubebuilder:validation:MinItems=1
	WeeksOfMonth []int64 `json:"weeksOfMonth,omitempty" tf:"weeks_of_month"`
}

type DbSystemSpecMaintenanceWindowDetailsDaysOfWeek struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type DbSystemSpecMaintenanceWindowDetailsMonths struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type DbSystemSpecMaintenanceWindowDetails struct {
	// +optional
	DaysOfWeek []DbSystemSpecMaintenanceWindowDetailsDaysOfWeek `json:"daysOfWeek,omitempty" tf:"days_of_week"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	HoursOfDay []int64 `json:"hoursOfDay,omitempty" tf:"hours_of_day"`
	// +optional
	LeadTimeInWeeks *int64 `json:"leadTimeInWeeks,omitempty" tf:"lead_time_in_weeks"`
	// +optional
	Months []DbSystemSpecMaintenanceWindowDetailsMonths `json:"months,omitempty" tf:"months"`
	// +optional
	Preference *string `json:"preference,omitempty" tf:"preference"`
	// +optional
	// +kubebuilder:validation:MaxItems=4
	// +kubebuilder:validation:MinItems=1
	WeeksOfMonth []int64 `json:"weeksOfMonth,omitempty" tf:"weeks_of_month"`
}

type DbSystemSpec struct {
	State *DbSystemSpecResource `json:"state,omitempty" tf:"-"`

	Resource DbSystemSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DbSystemSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AvailabilityDomain *string `json:"availabilityDomain" tf:"availability_domain"`
	// +optional
	BackupNetworkNsgIDS []string `json:"backupNetworkNsgIDS,omitempty" tf:"backup_network_nsg_ids"`
	// +optional
	BackupSubnetID *string `json:"backupSubnetID,omitempty" tf:"backup_subnet_id"`
	// +optional
	ClusterName   *string `json:"clusterName,omitempty" tf:"cluster_name"`
	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	CpuCoreCount *int64 `json:"cpuCoreCount,omitempty" tf:"cpu_core_count"`
	// +optional
	DataStoragePercentage *int64 `json:"dataStoragePercentage,omitempty" tf:"data_storage_percentage"`
	// +optional
	DataStorageSizeInGb *int64 `json:"dataStorageSizeInGb,omitempty" tf:"data_storage_size_in_gb"`
	// +optional
	DatabaseEdition *string             `json:"databaseEdition,omitempty" tf:"database_edition"`
	DbHome          *DbSystemSpecDbHome `json:"dbHome" tf:"db_home"`
	// +optional
	DbSystemOptions *DbSystemSpecDbSystemOptions `json:"dbSystemOptions,omitempty" tf:"db_system_options"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DiskRedundancy *string `json:"diskRedundancy,omitempty" tf:"disk_redundancy"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// +optional
	FaultDomains []string `json:"faultDomains,omitempty" tf:"fault_domains"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	Hostname     *string           `json:"hostname" tf:"hostname"`
	// +optional
	IormConfigCache *DbSystemSpecIormConfigCache `json:"iormConfigCache,omitempty" tf:"iorm_config_cache"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	KmsKeyVersionID *string `json:"kmsKeyVersionID,omitempty" tf:"kms_key_version_id"`
	// +optional
	LastMaintenanceRunID *string `json:"lastMaintenanceRunID,omitempty" tf:"last_maintenance_run_id"`
	// +optional
	LastPatchHistoryEntryID *string `json:"lastPatchHistoryEntryID,omitempty" tf:"last_patch_history_entry_id"`
	// +optional
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	ListenerPort *int64 `json:"listenerPort,omitempty" tf:"listener_port"`
	// +optional
	MaintenanceWindow *DbSystemSpecMaintenanceWindow `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`
	// +optional
	MaintenanceWindowDetails *DbSystemSpecMaintenanceWindowDetails `json:"maintenanceWindowDetails,omitempty" tf:"maintenance_window_details"`
	// +optional
	NextMaintenanceRunID *string `json:"nextMaintenanceRunID,omitempty" tf:"next_maintenance_run_id"`
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
	// +optional
	NsgIDS []string `json:"nsgIDS,omitempty" tf:"nsg_ids"`
	// +optional
	PointInTimeDataDiskCloneTimestamp *string `json:"pointInTimeDataDiskCloneTimestamp,omitempty" tf:"point_in_time_data_disk_clone_timestamp"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	RecoStorageSizeInGb *int64 `json:"recoStorageSizeInGb,omitempty" tf:"reco_storage_size_in_gb"`
	// +optional
	ScanDNSName *string `json:"scanDNSName,omitempty" tf:"scan_dns_name"`
	// +optional
	ScanDNSRecordID *string `json:"scanDNSRecordID,omitempty" tf:"scan_dns_record_id"`
	// +optional
	ScanIPIDS []string `json:"scanIPIDS,omitempty" tf:"scan_ip_ids"`
	Shape     *string  `json:"shape" tf:"shape"`
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
	// +optional
	SourceDbSystemID *string `json:"sourceDbSystemID,omitempty" tf:"source_db_system_id"`
	// +optional
	SparseDiskgroup *bool    `json:"sparseDiskgroup,omitempty" tf:"sparse_diskgroup"`
	SshPublicKeys   []string `json:"sshPublicKeys" tf:"ssh_public_keys"`
	// +optional
	State    *string `json:"state,omitempty" tf:"state"`
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	VipIDS []string `json:"vipIDS,omitempty" tf:"vip_ids"`
	// +optional
	ZoneID *string `json:"zoneID,omitempty" tf:"zone_id"`
}

type DbSystemStatus struct {
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

// DbSystemList is a list of DbSystems
type DbSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbSystem CRD objects
	Items []DbSystem `json:"items,omitempty"`
}
