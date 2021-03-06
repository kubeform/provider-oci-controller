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

type DataGuardAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataGuardAssociationSpec   `json:"spec,omitempty"`
	Status            DataGuardAssociationStatus `json:"status,omitempty"`
}

type DataGuardAssociationSpec struct {
	State *DataGuardAssociationSpecResource `json:"state,omitempty" tf:"-"`

	Resource DataGuardAssociationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DataGuardAssociationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApplyLag *string `json:"applyLag,omitempty" tf:"apply_lag"`
	// +optional
	ApplyRate *string `json:"applyRate,omitempty" tf:"apply_rate"`
	// +optional
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain"`
	// +optional
	BackupNetworkNsgIDS []string `json:"backupNetworkNsgIDS,omitempty" tf:"backup_network_nsg_ids"`
	// +optional
	CreateAsync           *bool   `json:"createAsync,omitempty" tf:"create_async"`
	CreationType          *string `json:"creationType" tf:"creation_type"`
	DatabaseAdminPassword *string `json:"-" sensitive:"true" tf:"database_admin_password"`
	DatabaseID            *string `json:"databaseID" tf:"database_id"`
	// +optional
	DatabaseSoftwareImageID     *string `json:"databaseSoftwareImageID,omitempty" tf:"database_software_image_id"`
	DeleteStandbyDbHomeOnDelete *string `json:"deleteStandbyDbHomeOnDelete" tf:"delete_standby_db_home_on_delete"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname"`
	// +optional
	IsActiveDataGuardEnabled *bool `json:"isActiveDataGuardEnabled,omitempty" tf:"is_active_data_guard_enabled"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	NsgIDS []string `json:"nsgIDS,omitempty" tf:"nsg_ids"`
	// +optional
	PeerDataGuardAssociationID *string `json:"peerDataGuardAssociationID,omitempty" tf:"peer_data_guard_association_id"`
	// +optional
	PeerDatabaseID *string `json:"peerDatabaseID,omitempty" tf:"peer_database_id"`
	// +optional
	PeerDbHomeID *string `json:"peerDbHomeID,omitempty" tf:"peer_db_home_id"`
	// +optional
	PeerDbSystemID *string `json:"peerDbSystemID,omitempty" tf:"peer_db_system_id"`
	// +optional
	PeerDbUniqueName *string `json:"peerDbUniqueName,omitempty" tf:"peer_db_unique_name"`
	// +optional
	PeerRole *string `json:"peerRole,omitempty" tf:"peer_role"`
	// +optional
	PeerSidPrefix *string `json:"peerSidPrefix,omitempty" tf:"peer_sid_prefix"`
	// +optional
	PeerVmClusterID *string `json:"peerVmClusterID,omitempty" tf:"peer_vm_cluster_id"`
	ProtectionMode  *string `json:"protectionMode" tf:"protection_mode"`
	// +optional
	Role *string `json:"role,omitempty" tf:"role"`
	// +optional
	Shape *string `json:"shape,omitempty" tf:"shape"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	TimeCreated   *string `json:"timeCreated,omitempty" tf:"time_created"`
	TransportType *string `json:"transportType" tf:"transport_type"`
}

type DataGuardAssociationStatus struct {
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

// DataGuardAssociationList is a list of DataGuardAssociations
type DataGuardAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataGuardAssociation CRD objects
	Items []DataGuardAssociation `json:"items,omitempty"`
}
