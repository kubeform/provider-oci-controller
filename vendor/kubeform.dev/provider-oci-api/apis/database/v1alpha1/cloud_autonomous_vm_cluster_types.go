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

type CloudAutonomousVmCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudAutonomousVmClusterSpec   `json:"spec,omitempty"`
	Status            CloudAutonomousVmClusterStatus `json:"status,omitempty"`
}

type CloudAutonomousVmClusterSpec struct {
	State *CloudAutonomousVmClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource CloudAutonomousVmClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CloudAutonomousVmClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AvailabilityDomain           *string `json:"availabilityDomain,omitempty" tf:"availability_domain"`
	CloudExadataInfrastructureID *string `json:"cloudExadataInfrastructureID" tf:"cloud_exadata_infrastructure_id"`
	CompartmentID                *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	CpuCoreCount *int64 `json:"cpuCoreCount,omitempty" tf:"cpu_core_count"`
	// +optional
	DataStorageSizeInGb *float64 `json:"dataStorageSizeInGb,omitempty" tf:"data_storage_size_in_gb"`
	// +optional
	DataStorageSizeInTbs *float64 `json:"dataStorageSizeInTbs,omitempty" tf:"data_storage_size_in_tbs"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// +optional
	LastMaintenanceRunID *string `json:"lastMaintenanceRunID,omitempty" tf:"last_maintenance_run_id"`
	// +optional
	LastUpdateHistoryEntryID *string `json:"lastUpdateHistoryEntryID,omitempty" tf:"last_update_history_entry_id"`
	// +optional
	LicenseModel *string `json:"licenseModel,omitempty" tf:"license_model"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	MemorySizeInGbs *int64 `json:"memorySizeInGbs,omitempty" tf:"memory_size_in_gbs"`
	// +optional
	NextMaintenanceRunID *string `json:"nextMaintenanceRunID,omitempty" tf:"next_maintenance_run_id"`
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
	// +optional
	NsgIDS []string `json:"nsgIDS,omitempty" tf:"nsg_ids"`
	// +optional
	OcpuCount *float64 `json:"ocpuCount,omitempty" tf:"ocpu_count"`
	// +optional
	RotateOrdsCertsTrigger *bool `json:"rotateOrdsCertsTrigger,omitempty" tf:"rotate_ords_certs_trigger"`
	// +optional
	RotateSslCertsTrigger *bool `json:"rotateSslCertsTrigger,omitempty" tf:"rotate_ssl_certs_trigger"`
	// +optional
	Shape *string `json:"shape,omitempty" tf:"shape"`
	// +optional
	State    *string `json:"state,omitempty" tf:"state"`
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type CloudAutonomousVmClusterStatus struct {
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

// CloudAutonomousVmClusterList is a list of CloudAutonomousVmClusters
type CloudAutonomousVmClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudAutonomousVmCluster CRD objects
	Items []CloudAutonomousVmCluster `json:"items,omitempty"`
}
