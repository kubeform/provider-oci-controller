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

type EncryptedData struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EncryptedDataSpec   `json:"spec,omitempty"`
	Status            EncryptedDataStatus `json:"status,omitempty"`
}

type EncryptedDataSpec struct {
	State *EncryptedDataSpecResource `json:"state,omitempty" tf:"-"`

	Resource EncryptedDataSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EncryptedDataSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AssociatedData map[string]string `json:"associatedData,omitempty" tf:"associated_data"`
	// +optional
	Ciphertext     *string `json:"ciphertext,omitempty" tf:"ciphertext"`
	CryptoEndpoint *string `json:"cryptoEndpoint" tf:"crypto_endpoint"`
	// +optional
	EncryptionAlgorithm *string `json:"encryptionAlgorithm,omitempty" tf:"encryption_algorithm"`
	KeyID               *string `json:"keyID" tf:"key_id"`
	// +optional
	KeyVersionID *string `json:"keyVersionID,omitempty" tf:"key_version_id"`
	// +optional
	LoggingContext map[string]string `json:"loggingContext,omitempty" tf:"logging_context"`
	Plaintext      *string           `json:"plaintext" tf:"plaintext"`
}

type EncryptedDataStatus struct {
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

// EncryptedDataList is a list of EncryptedDatas
type EncryptedDataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EncryptedData CRD objects
	Items []EncryptedData `json:"items,omitempty"`
}
