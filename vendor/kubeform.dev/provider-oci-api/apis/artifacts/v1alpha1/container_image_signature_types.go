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

type ContainerImageSignature struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerImageSignatureSpec   `json:"spec,omitempty"`
	Status            ContainerImageSignatureStatus `json:"status,omitempty"`
}

type ContainerImageSignatureSpec struct {
	State *ContainerImageSignatureSpecResource `json:"state,omitempty" tf:"-"`

	Resource ContainerImageSignatureSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ContainerImageSignatureSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by"`
	// +optional
	DisplayName      *string `json:"displayName,omitempty" tf:"display_name"`
	ImageID          *string `json:"imageID" tf:"image_id"`
	KmsKeyID         *string `json:"kmsKeyID" tf:"kms_key_id"`
	KmsKeyVersionID  *string `json:"kmsKeyVersionID" tf:"kms_key_version_id"`
	Message          *string `json:"message" tf:"message"`
	Signature        *string `json:"signature" tf:"signature"`
	SigningAlgorithm *string `json:"signingAlgorithm" tf:"signing_algorithm"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
}

type ContainerImageSignatureStatus struct {
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

// ContainerImageSignatureList is a list of ContainerImageSignatures
type ContainerImageSignatureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ContainerImageSignature CRD objects
	Items []ContainerImageSignature `json:"items,omitempty"`
}