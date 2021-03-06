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

type Session struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SessionSpec   `json:"spec,omitempty"`
	Status            SessionStatus `json:"status,omitempty"`
}

type SessionSpecKeyDetails struct {
	PublicKeyContent *string `json:"publicKeyContent" tf:"public_key_content"`
}

type SessionSpecTargetResourceDetails struct {
	SessionType *string `json:"sessionType" tf:"session_type"`
	// +optional
	TargetResourceDisplayName *string `json:"targetResourceDisplayName,omitempty" tf:"target_resource_display_name"`
	// +optional
	TargetResourceID *string `json:"targetResourceID,omitempty" tf:"target_resource_id"`
	// +optional
	TargetResourceOperatingSystemUserName *string `json:"targetResourceOperatingSystemUserName,omitempty" tf:"target_resource_operating_system_user_name"`
	// +optional
	TargetResourcePort *int64 `json:"targetResourcePort,omitempty" tf:"target_resource_port"`
	// +optional
	TargetResourcePrivateIPAddress *string `json:"targetResourcePrivateIPAddress,omitempty" tf:"target_resource_private_ip_address"`
}

type SessionSpec struct {
	State *SessionSpecResource `json:"state,omitempty" tf:"-"`

	Resource SessionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SessionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BastionID *string `json:"bastionID" tf:"bastion_id"`
	// +optional
	BastionName *string `json:"bastionName,omitempty" tf:"bastion_name"`
	// +optional
	BastionPublicHostKeyInfo *string `json:"bastionPublicHostKeyInfo,omitempty" tf:"bastion_public_host_key_info"`
	// +optional
	BastionUserName *string `json:"bastionUserName,omitempty" tf:"bastion_user_name"`
	// +optional
	DisplayName *string                `json:"displayName,omitempty" tf:"display_name"`
	KeyDetails  *SessionSpecKeyDetails `json:"keyDetails" tf:"key_details"`
	// +optional
	KeyType *string `json:"keyType,omitempty" tf:"key_type"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	SessionTtlInSeconds *int64 `json:"sessionTtlInSeconds,omitempty" tf:"session_ttl_in_seconds"`
	// +optional
	SshMetadata map[string]string `json:"sshMetadata,omitempty" tf:"ssh_metadata"`
	// +optional
	State                 *string                           `json:"state,omitempty" tf:"state"`
	TargetResourceDetails *SessionSpecTargetResourceDetails `json:"targetResourceDetails" tf:"target_resource_details"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type SessionStatus struct {
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

// SessionList is a list of Sessions
type SessionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Session CRD objects
	Items []Session `json:"items,omitempty"`
}
