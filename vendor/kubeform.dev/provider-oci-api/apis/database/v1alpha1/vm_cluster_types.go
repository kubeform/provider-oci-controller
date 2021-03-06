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

type VmCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VmClusterSpec   `json:"spec,omitempty"`
	Status            VmClusterStatus `json:"status,omitempty"`
}

type VmClusterSpec struct {
	State *VmClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource VmClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VmClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	CpuCoreCount  *int64  `json:"cpuCoreCount" tf:"cpu_core_count"`
	// +optional
	CpusEnabled *int64 `json:"cpusEnabled,omitempty" tf:"cpus_enabled"`
	// +optional
	DataStorageSizeInGb *float64 `json:"dataStorageSizeInGb,omitempty" tf:"data_storage_size_in_gb"`
	// +optional
	DataStorageSizeInTbs *float64 `json:"dataStorageSizeInTbs,omitempty" tf:"data_storage_size_in_tbs"`
	// +optional
	DbNodeStorageSizeInGbs *int64 `json:"dbNodeStorageSizeInGbs,omitempty" tf:"db_node_storage_size_in_gbs"`
	// +optional
	DbServers []string `json:"dbServers,omitempty" tf:"db_servers"`
	// +optional
	DefinedTags             map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	DisplayName             *string           `json:"displayName" tf:"display_name"`
	ExadataInfrastructureID *string           `json:"exadataInfrastructureID" tf:"exadata_infrastructure_id"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	GiVersion    *string           `json:"giVersion" tf:"gi_version"`
	// +optional
	IsLocalBackupEnabled *bool `json:"isLocalBackupEnabled,omitempty" tf:"is_local_backup_enabled"`
	// +optional
	IsSparseDiskgroupEnabled *bool `json:"isSparseDiskgroupEnabled,omitempty" tf:"is_sparse_diskgroup_enabled"`
	// +optional
	LastPatchHistoryEntryID *string `json:"lastPatchHistoryEntryID,omitempty" tf:"last_patch_history_entry_id"`
	// +optional
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	MemorySizeInGbs *int64 `json:"memorySizeInGbs,omitempty" tf:"memory_size_in_gbs"`
	// +optional
	OcpuCount *float64 `json:"ocpuCount,omitempty" tf:"ocpu_count"`
	// +optional
	OcpusEnabled *float64 `json:"ocpusEnabled,omitempty" tf:"ocpus_enabled"`
	// +optional
	Shape         *string  `json:"shape,omitempty" tf:"shape"`
	SshPublicKeys []string `json:"sshPublicKeys" tf:"ssh_public_keys"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SystemVersion *string `json:"systemVersion,omitempty" tf:"system_version"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeZone           *string `json:"timeZone,omitempty" tf:"time_zone"`
	VmClusterNetworkID *string `json:"vmClusterNetworkID" tf:"vm_cluster_network_id"`
}

type VmClusterStatus struct {
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

// VmClusterList is a list of VmClusters
type VmClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VmCluster CRD objects
	Items []VmCluster `json:"items,omitempty"`
}
