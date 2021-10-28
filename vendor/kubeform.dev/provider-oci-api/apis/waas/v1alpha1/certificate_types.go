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

type Certificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec,omitempty"`
	Status            CertificateStatus `json:"status,omitempty"`
}

type CertificateSpecExtensions struct {
	// +optional
	IsCritical *bool `json:"isCritical,omitempty" tf:"is_critical"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type CertificateSpecIssuerName struct {
	// +optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name"`
	// +optional
	Country *string `json:"country,omitempty" tf:"country"`
	// +optional
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address"`
	// +optional
	Locality *string `json:"locality,omitempty" tf:"locality"`
	// +optional
	Organization *string `json:"organization,omitempty" tf:"organization"`
	// +optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit"`
	// +optional
	StateProvince *string `json:"stateProvince,omitempty" tf:"state_province"`
}

type CertificateSpecPublicKeyInfo struct {
	// +optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm"`
	// +optional
	Exponent *int64 `json:"exponent,omitempty" tf:"exponent"`
	// +optional
	KeySize *int64 `json:"keySize,omitempty" tf:"key_size"`
}

type CertificateSpecSubjectName struct {
	// +optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name"`
	// +optional
	Country *string `json:"country,omitempty" tf:"country"`
	// +optional
	EmailAddress *string `json:"emailAddress,omitempty" tf:"email_address"`
	// +optional
	Locality *string `json:"locality,omitempty" tf:"locality"`
	// +optional
	Organization *string `json:"organization,omitempty" tf:"organization"`
	// +optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit"`
	// +optional
	StateProvince *string `json:"stateProvince,omitempty" tf:"state_province"`
}

type CertificateSpec struct {
	State *CertificateSpecResource `json:"state,omitempty" tf:"-"`

	Resource CertificateSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CertificateSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CertificateData *string `json:"certificateData" tf:"certificate_data"`
	CompartmentID   *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	Extensions []CertificateSpecExtensions `json:"extensions,omitempty" tf:"extensions"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	IsTrustVerificationDisabled *bool `json:"isTrustVerificationDisabled,omitempty" tf:"is_trust_verification_disabled"`
	// +optional
	IssuedBy *string `json:"issuedBy,omitempty" tf:"issued_by"`
	// +optional
	IssuerName     *CertificateSpecIssuerName `json:"issuerName,omitempty" tf:"issuer_name"`
	PrivateKeyData *string                    `json:"-" sensitive:"true" tf:"private_key_data"`
	// +optional
	PublicKeyInfo *CertificateSpecPublicKeyInfo `json:"publicKeyInfo,omitempty" tf:"public_key_info"`
	// +optional
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number"`
	// +optional
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty" tf:"signature_algorithm"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SubjectName *CertificateSpecSubjectName `json:"subjectName,omitempty" tf:"subject_name"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeNotValidAfter *string `json:"timeNotValidAfter,omitempty" tf:"time_not_valid_after"`
	// +optional
	TimeNotValidBefore *string `json:"timeNotValidBefore,omitempty" tf:"time_not_valid_before"`
	// +optional
	Version *int64 `json:"version,omitempty" tf:"version"`
}

type CertificateStatus struct {
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

// CertificateList is a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Certificate CRD objects
	Items []Certificate `json:"items,omitempty"`
}
