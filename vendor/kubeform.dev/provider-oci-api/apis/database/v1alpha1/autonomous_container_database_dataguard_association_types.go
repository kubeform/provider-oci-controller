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

type AutonomousContainerDatabaseDataguardAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutonomousContainerDatabaseDataguardAssociationSpec   `json:"spec,omitempty"`
	Status            AutonomousContainerDatabaseDataguardAssociationStatus `json:"status,omitempty"`
}

type AutonomousContainerDatabaseDataguardAssociationSpec struct {
	State *AutonomousContainerDatabaseDataguardAssociationSpecResource `json:"state,omitempty" tf:"-"`

	Resource AutonomousContainerDatabaseDataguardAssociationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AutonomousContainerDatabaseDataguardAssociationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApplyLag *string `json:"applyLag,omitempty" tf:"apply_lag"`
	// +optional
	ApplyRate                                         *string `json:"applyRate,omitempty" tf:"apply_rate"`
	AutonomousContainerDatabaseDataguardAssociationID *string `json:"autonomousContainerDatabaseDataguardAssociationID" tf:"autonomous_container_database_dataguard_association_id"`
	AutonomousContainerDatabaseID                     *string `json:"autonomousContainerDatabaseID" tf:"autonomous_container_database_id"`
	// +optional
	IsAutomaticFailoverEnabled *bool `json:"isAutomaticFailoverEnabled,omitempty" tf:"is_automatic_failover_enabled"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	PeerAutonomousContainerDatabaseDataguardAssociationID *string `json:"peerAutonomousContainerDatabaseDataguardAssociationID,omitempty" tf:"peer_autonomous_container_database_dataguard_association_id"`
	// +optional
	PeerAutonomousContainerDatabaseID *string `json:"peerAutonomousContainerDatabaseID,omitempty" tf:"peer_autonomous_container_database_id"`
	// +optional
	PeerLifecycleState *string `json:"peerLifecycleState,omitempty" tf:"peer_lifecycle_state"`
	// +optional
	PeerRole *string `json:"peerRole,omitempty" tf:"peer_role"`
	// +optional
	ProtectionMode *string `json:"protectionMode,omitempty" tf:"protection_mode"`
	// +optional
	Role *string `json:"role,omitempty" tf:"role"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeLastRoleChanged *string `json:"timeLastRoleChanged,omitempty" tf:"time_last_role_changed"`
	// +optional
	TimeLastSynced *string `json:"timeLastSynced,omitempty" tf:"time_last_synced"`
	// +optional
	TransportLag *string `json:"transportLag,omitempty" tf:"transport_lag"`
}

type AutonomousContainerDatabaseDataguardAssociationStatus struct {
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

// AutonomousContainerDatabaseDataguardAssociationList is a list of AutonomousContainerDatabaseDataguardAssociations
type AutonomousContainerDatabaseDataguardAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutonomousContainerDatabaseDataguardAssociation CRD objects
	Items []AutonomousContainerDatabaseDataguardAssociation `json:"items,omitempty"`
}
