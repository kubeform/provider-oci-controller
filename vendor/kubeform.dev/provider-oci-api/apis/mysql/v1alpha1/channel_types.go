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

type Channel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChannelSpec   `json:"spec,omitempty"`
	Status            ChannelStatus `json:"status,omitempty"`
}

type ChannelSpecSourceSslCaCertificate struct {
	CertificateType *string `json:"certificateType" tf:"certificate_type"`
	Contents        *string `json:"contents" tf:"contents"`
}

type ChannelSpecSource struct {
	Hostname *string `json:"hostname" tf:"hostname"`
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	Port       *int64  `json:"port,omitempty" tf:"port"`
	SourceType *string `json:"sourceType" tf:"source_type"`
	// +optional
	SslCaCertificate *ChannelSpecSourceSslCaCertificate `json:"sslCaCertificate,omitempty" tf:"ssl_ca_certificate"`
	SslMode          *string                            `json:"sslMode" tf:"ssl_mode"`
	Username         *string                            `json:"username" tf:"username"`
}

type ChannelSpecTarget struct {
	// +optional
	ApplierUsername *string `json:"applierUsername,omitempty" tf:"applier_username"`
	// +optional
	ChannelName *string `json:"channelName,omitempty" tf:"channel_name"`
	DbSystemID  *string `json:"dbSystemID" tf:"db_system_id"`
	TargetType  *string `json:"targetType" tf:"target_type"`
}

type ChannelSpec struct {
	State *ChannelSpecResource `json:"state,omitempty" tf:"-"`

	Resource ChannelSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ChannelSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
	// +optional
	LifecycleDetails *string            `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	Source           *ChannelSpecSource `json:"source" tf:"source"`
	// +optional
	State  *string            `json:"state,omitempty" tf:"state"`
	Target *ChannelSpecTarget `json:"target" tf:"target"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type ChannelStatus struct {
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

// ChannelList is a list of Channels
type ChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Channel CRD objects
	Items []Channel `json:"items,omitempty"`
}
