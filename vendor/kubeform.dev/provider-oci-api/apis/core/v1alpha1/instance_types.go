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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecAgentConfigPluginsConfig struct {
	DesiredState *string `json:"desiredState" tf:"desired_state"`
	Name         *string `json:"name" tf:"name"`
}

type InstanceSpecAgentConfig struct {
	// +optional
	AreAllPluginsDisabled *bool `json:"areAllPluginsDisabled,omitempty" tf:"are_all_plugins_disabled"`
	// +optional
	IsManagementDisabled *bool `json:"isManagementDisabled,omitempty" tf:"is_management_disabled"`
	// +optional
	IsMonitoringDisabled *bool `json:"isMonitoringDisabled,omitempty" tf:"is_monitoring_disabled"`
	// +optional
	PluginsConfig []InstanceSpecAgentConfigPluginsConfig `json:"pluginsConfig,omitempty" tf:"plugins_config"`
}

type InstanceSpecAvailabilityConfig struct {
	// +optional
	IsLiveMigrationPreferred *bool `json:"isLiveMigrationPreferred,omitempty" tf:"is_live_migration_preferred"`
	// +optional
	RecoveryAction *string `json:"recoveryAction,omitempty" tf:"recovery_action"`
}

type InstanceSpecCreateVnicDetails struct {
	// +optional
	AssignPrivateDNSRecord *bool `json:"assignPrivateDNSRecord,omitempty" tf:"assign_private_dns_record"`
	// +optional
	AssignPublicIP *string `json:"assignPublicIP,omitempty" tf:"assign_public_ip"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	HostnameLabel *string `json:"hostnameLabel,omitempty" tf:"hostname_label"`
	// +optional
	NsgIDS []string `json:"nsgIDS,omitempty" tf:"nsg_ids"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	SkipSourceDestCheck *bool `json:"skipSourceDestCheck,omitempty" tf:"skip_source_dest_check"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	VlanID *string `json:"vlanID,omitempty" tf:"vlan_id"`
}

type InstanceSpecInstanceOptions struct {
	// +optional
	AreLegacyImdsEndpointsDisabled *bool `json:"areLegacyImdsEndpointsDisabled,omitempty" tf:"are_legacy_imds_endpoints_disabled"`
}

type InstanceSpecLaunchOptions struct {
	// +optional
	BootVolumeType *string `json:"bootVolumeType,omitempty" tf:"boot_volume_type"`
	// +optional
	Firmware *string `json:"firmware,omitempty" tf:"firmware"`
	// +optional
	IsConsistentVolumeNamingEnabled *bool `json:"isConsistentVolumeNamingEnabled,omitempty" tf:"is_consistent_volume_naming_enabled"`
	// +optional
	IsPvEncryptionInTransitEnabled *bool `json:"isPvEncryptionInTransitEnabled,omitempty" tf:"is_pv_encryption_in_transit_enabled"`
	// +optional
	NetworkType *string `json:"networkType,omitempty" tf:"network_type"`
	// +optional
	RemoteDataVolumeType *string `json:"remoteDataVolumeType,omitempty" tf:"remote_data_volume_type"`
}

type InstanceSpecPlatformConfig struct {
	// +optional
	IsMeasuredBootEnabled *bool `json:"isMeasuredBootEnabled,omitempty" tf:"is_measured_boot_enabled"`
	// +optional
	IsSecureBootEnabled *bool `json:"isSecureBootEnabled,omitempty" tf:"is_secure_boot_enabled"`
	// +optional
	IsTrustedPlatformModuleEnabled *bool `json:"isTrustedPlatformModuleEnabled,omitempty" tf:"is_trusted_platform_module_enabled"`
	// +optional
	NumaNodesPerSocket *string `json:"numaNodesPerSocket,omitempty" tf:"numa_nodes_per_socket"`
	Type               *string `json:"type" tf:"type"`
}

type InstanceSpecPreemptibleInstanceConfigPreemptionAction struct {
	// +optional
	PreserveBootVolume *bool   `json:"preserveBootVolume,omitempty" tf:"preserve_boot_volume"`
	Type               *string `json:"type" tf:"type"`
}

type InstanceSpecPreemptibleInstanceConfig struct {
	PreemptionAction *InstanceSpecPreemptibleInstanceConfigPreemptionAction `json:"preemptionAction" tf:"preemption_action"`
}

type InstanceSpecShapeConfig struct {
	// +optional
	BaselineOcpuUtilization *string `json:"baselineOcpuUtilization,omitempty" tf:"baseline_ocpu_utilization"`
	// +optional
	GpuDescription *string `json:"gpuDescription,omitempty" tf:"gpu_description"`
	// +optional
	Gpus *int64 `json:"gpus,omitempty" tf:"gpus"`
	// +optional
	LocalDiskDescription *string `json:"localDiskDescription,omitempty" tf:"local_disk_description"`
	// +optional
	LocalDisks *int64 `json:"localDisks,omitempty" tf:"local_disks"`
	// +optional
	LocalDisksTotalSizeInGbs *float64 `json:"localDisksTotalSizeInGbs,omitempty" tf:"local_disks_total_size_in_gbs"`
	// +optional
	MaxVnicAttachments *int64 `json:"maxVnicAttachments,omitempty" tf:"max_vnic_attachments"`
	// +optional
	MemoryInGbs *float64 `json:"memoryInGbs,omitempty" tf:"memory_in_gbs"`
	// +optional
	NetworkingBandwidthInGbps *float64 `json:"networkingBandwidthInGbps,omitempty" tf:"networking_bandwidth_in_gbps"`
	// +optional
	Ocpus *float64 `json:"ocpus,omitempty" tf:"ocpus"`
	// +optional
	ProcessorDescription *string `json:"processorDescription,omitempty" tf:"processor_description"`
}

type InstanceSpecSourceDetails struct {
	// +optional
	BootVolumeSizeInGbs *string `json:"bootVolumeSizeInGbs,omitempty" tf:"boot_volume_size_in_gbs"`
	// +optional
	KmsKeyID   *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	SourceID   *string `json:"sourceID" tf:"source_id"`
	SourceType *string `json:"sourceType" tf:"source_type"`
}

type InstanceSpec struct {
	State *InstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AgentConfig *InstanceSpecAgentConfig `json:"agentConfig,omitempty" tf:"agent_config"`
	// +optional
	Async *bool `json:"async,omitempty" tf:"async"`
	// +optional
	AvailabilityConfig *InstanceSpecAvailabilityConfig `json:"availabilityConfig,omitempty" tf:"availability_config"`
	AvailabilityDomain *string                         `json:"availabilityDomain" tf:"availability_domain"`
	// +optional
	BootVolumeID *string `json:"bootVolumeID,omitempty" tf:"boot_volume_id"`
	// +optional
	CapacityReservationID *string `json:"capacityReservationID,omitempty" tf:"capacity_reservation_id"`
	CompartmentID         *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	CreateVnicDetails *InstanceSpecCreateVnicDetails `json:"createVnicDetails,omitempty" tf:"create_vnic_details"`
	// +optional
	DedicatedVmHostID *string `json:"dedicatedVmHostID,omitempty" tf:"dedicated_vm_host_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	ExtendedMetadata map[string]string `json:"extendedMetadata,omitempty" tf:"extended_metadata"`
	// +optional
	FaultDomain *string `json:"faultDomain,omitempty" tf:"fault_domain"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	// Deprecated
	HostnameLabel *string `json:"hostnameLabel,omitempty" tf:"hostname_label"`
	// +optional
	// Deprecated
	Image *string `json:"image,omitempty" tf:"image"`
	// +optional
	InstanceOptions *InstanceSpecInstanceOptions `json:"instanceOptions,omitempty" tf:"instance_options"`
	// +optional
	IpxeScript *string `json:"ipxeScript,omitempty" tf:"ipxe_script"`
	// +optional
	IsPvEncryptionInTransitEnabled *bool `json:"isPvEncryptionInTransitEnabled,omitempty" tf:"is_pv_encryption_in_transit_enabled"`
	// +optional
	LaunchMode *string `json:"launchMode,omitempty" tf:"launch_mode"`
	// +optional
	LaunchOptions *InstanceSpecLaunchOptions `json:"launchOptions,omitempty" tf:"launch_options"`
	// +optional
	Metadata map[string]string `json:"metadata,omitempty" tf:"metadata"`
	// +optional
	PlatformConfig *InstanceSpecPlatformConfig `json:"platformConfig,omitempty" tf:"platform_config"`
	// +optional
	PreemptibleInstanceConfig *InstanceSpecPreemptibleInstanceConfig `json:"preemptibleInstanceConfig,omitempty" tf:"preemptible_instance_config"`
	// +optional
	PreserveBootVolume *bool `json:"preserveBootVolume,omitempty" tf:"preserve_boot_volume"`
	// +optional
	PrivateIP *string `json:"privateIP,omitempty" tf:"private_ip"`
	// +optional
	PublicIP *string `json:"publicIP,omitempty" tf:"public_ip"`
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	Shape  *string `json:"shape" tf:"shape"`
	// +optional
	ShapeConfig *InstanceSpecShapeConfig `json:"shapeConfig,omitempty" tf:"shape_config"`
	// +optional
	SourceDetails *InstanceSpecSourceDetails `json:"sourceDetails,omitempty" tf:"source_details"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	// Deprecated
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeMaintenanceRebootDue *string `json:"timeMaintenanceRebootDue,omitempty" tf:"time_maintenance_reboot_due"`
}

type InstanceStatus struct {
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

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
