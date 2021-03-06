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

type GeneratedKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GeneratedKeySpec   `json:"spec,omitempty"`
	Status            GeneratedKeyStatus `json:"status,omitempty"`
}

type GeneratedKeySpecKeyShape struct {
	Algorithm *string `json:"algorithm" tf:"algorithm"`
	// +optional
	CurveID *string `json:"curveID,omitempty" tf:"curve_id"`
	Length  *int64  `json:"length" tf:"length"`
}

type GeneratedKeySpec struct {
	State *GeneratedKeySpecResource `json:"state,omitempty" tf:"-"`

	Resource GeneratedKeySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GeneratedKeySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AssociatedData map[string]string `json:"associatedData,omitempty" tf:"associated_data"`
	// +optional
	Ciphertext          *string                   `json:"ciphertext,omitempty" tf:"ciphertext"`
	CryptoEndpoint      *string                   `json:"cryptoEndpoint" tf:"crypto_endpoint"`
	IncludePlaintextKey *bool                     `json:"includePlaintextKey" tf:"include_plaintext_key"`
	KeyID               *string                   `json:"keyID" tf:"key_id"`
	KeyShape            *GeneratedKeySpecKeyShape `json:"keyShape" tf:"key_shape"`
	// +optional
	LoggingContext map[string]string `json:"loggingContext,omitempty" tf:"logging_context"`
	// +optional
	Plaintext *string `json:"plaintext,omitempty" tf:"plaintext"`
	// +optional
	PlaintextChecksum *string `json:"plaintextChecksum,omitempty" tf:"plaintext_checksum"`
}

type GeneratedKeyStatus struct {
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

// GeneratedKeyList is a list of GeneratedKeys
type GeneratedKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GeneratedKey CRD objects
	Items []GeneratedKey `json:"items,omitempty"`
}
