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

type VolumeAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeAttachmentSpec   `json:"spec,omitempty"`
	Status            VolumeAttachmentStatus `json:"status,omitempty"`
}

type VolumeAttachmentSpecMultipathDevices struct {
	// +optional
	Ipv4 *string `json:"ipv4,omitempty" tf:"ipv4"`
	// +optional
	Iqn *string `json:"iqn,omitempty" tf:"iqn"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
}

type VolumeAttachmentSpec struct {
	State *VolumeAttachmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource VolumeAttachmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VolumeAttachmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AttachmentType *string `json:"attachmentType" tf:"attachment_type"`
	// +optional
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain"`
	// +optional
	ChapSecret *string `json:"chapSecret,omitempty" tf:"chap_secret"`
	// +optional
	ChapUsername *string `json:"chapUsername,omitempty" tf:"chap_username"`
	// +optional
	// Deprecated
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	Device *string `json:"device,omitempty" tf:"device"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	EncryptionInTransitType *string `json:"encryptionInTransitType,omitempty" tf:"encryption_in_transit_type"`
	InstanceID              *string `json:"instanceID" tf:"instance_id"`
	// +optional
	Ipv4 *string `json:"ipv4,omitempty" tf:"ipv4"`
	// +optional
	Iqn *string `json:"iqn,omitempty" tf:"iqn"`
	// +optional
	IsMultipath *bool `json:"isMultipath,omitempty" tf:"is_multipath"`
	// +optional
	IsPvEncryptionInTransitEnabled *bool `json:"isPvEncryptionInTransitEnabled,omitempty" tf:"is_pv_encryption_in_transit_enabled"`
	// +optional
	IsReadOnly *bool `json:"isReadOnly,omitempty" tf:"is_read_only"`
	// +optional
	IsShareable *bool `json:"isShareable,omitempty" tf:"is_shareable"`
	// +optional
	IscsiLoginState *string `json:"iscsiLoginState,omitempty" tf:"iscsi_login_state"`
	// +optional
	MultipathDevices []VolumeAttachmentSpecMultipathDevices `json:"multipathDevices,omitempty" tf:"multipath_devices"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	UseCHAP  *bool   `json:"useCHAP,omitempty" tf:"use_chap"`
	VolumeID *string `json:"volumeID" tf:"volume_id"`
}

type VolumeAttachmentStatus struct {
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

// VolumeAttachmentList is a list of VolumeAttachments
type VolumeAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VolumeAttachment CRD objects
	Items []VolumeAttachment `json:"items,omitempty"`
}