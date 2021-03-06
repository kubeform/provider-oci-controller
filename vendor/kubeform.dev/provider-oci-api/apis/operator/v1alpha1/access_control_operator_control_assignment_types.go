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

type AccessControlOperatorControlAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessControlOperatorControlAssignmentSpec   `json:"spec,omitempty"`
	Status            AccessControlOperatorControlAssignmentStatus `json:"status,omitempty"`
}

type AccessControlOperatorControlAssignmentSpec struct {
	State *AccessControlOperatorControlAssignmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource AccessControlOperatorControlAssignmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AccessControlOperatorControlAssignmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AssignerID *string `json:"assignerID,omitempty" tf:"assigner_id"`
	// +optional
	Comment       *string `json:"comment,omitempty" tf:"comment"`
	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DetachmentDescription *string `json:"detachmentDescription,omitempty" tf:"detachment_description"`
	// +optional
	ErrorCode *int64 `json:"errorCode,omitempty" tf:"error_code"`
	// +optional
	ErrorMessage *string `json:"errorMessage,omitempty" tf:"error_message"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	IsAutoApproveDuringMaintenance *bool `json:"isAutoApproveDuringMaintenance,omitempty" tf:"is_auto_approve_during_maintenance"`
	IsEnforcedAlways               *bool `json:"isEnforcedAlways" tf:"is_enforced_always"`
	// +optional
	IsLogForwarded    *bool   `json:"isLogForwarded,omitempty" tf:"is_log_forwarded"`
	OperatorControlID *string `json:"operatorControlID" tf:"operator_control_id"`
	// +optional
	RemoteSyslogServerAddress *string `json:"remoteSyslogServerAddress,omitempty" tf:"remote_syslog_server_address"`
	// +optional
	RemoteSyslogServerCaCert *string `json:"remoteSyslogServerCaCert,omitempty" tf:"remote_syslog_server_ca_cert"`
	// +optional
	RemoteSyslogServerPort *int64  `json:"remoteSyslogServerPort,omitempty" tf:"remote_syslog_server_port"`
	ResourceCompartmentID  *string `json:"resourceCompartmentID" tf:"resource_compartment_id"`
	ResourceID             *string `json:"resourceID" tf:"resource_id"`
	ResourceName           *string `json:"resourceName" tf:"resource_name"`
	ResourceType           *string `json:"resourceType" tf:"resource_type"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeAssignmentFrom *string `json:"timeAssignmentFrom,omitempty" tf:"time_assignment_from"`
	// +optional
	TimeAssignmentTo *string `json:"timeAssignmentTo,omitempty" tf:"time_assignment_to"`
	// +optional
	TimeOfAssignment *string `json:"timeOfAssignment,omitempty" tf:"time_of_assignment"`
	// +optional
	TimeOfDeletion *string `json:"timeOfDeletion,omitempty" tf:"time_of_deletion"`
	// +optional
	UnassignerID *string `json:"unassignerID,omitempty" tf:"unassigner_id"`
}

type AccessControlOperatorControlAssignmentStatus struct {
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

// AccessControlOperatorControlAssignmentList is a list of AccessControlOperatorControlAssignments
type AccessControlOperatorControlAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AccessControlOperatorControlAssignment CRD objects
	Items []AccessControlOperatorControlAssignment `json:"items,omitempty"`
}
