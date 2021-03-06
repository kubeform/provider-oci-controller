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

type ManagementCertificateAuthority struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagementCertificateAuthoritySpec   `json:"spec,omitempty"`
	Status            ManagementCertificateAuthorityStatus `json:"status,omitempty"`
}

type ManagementCertificateAuthoritySpecCertificateAuthorityConfigSubject struct {
	CommonName *string `json:"commonName" tf:"common_name"`
	// +optional
	Country *string `json:"country,omitempty" tf:"country"`
	// +optional
	DistinguishedNameQualifier *string `json:"distinguishedNameQualifier,omitempty" tf:"distinguished_name_qualifier"`
	// +optional
	DomainComponent *string `json:"domainComponent,omitempty" tf:"domain_component"`
	// +optional
	GenerationQualifier *string `json:"generationQualifier,omitempty" tf:"generation_qualifier"`
	// +optional
	GivenName *string `json:"givenName,omitempty" tf:"given_name"`
	// +optional
	Initials *string `json:"initials,omitempty" tf:"initials"`
	// +optional
	LocalityName *string `json:"localityName,omitempty" tf:"locality_name"`
	// +optional
	Organization *string `json:"organization,omitempty" tf:"organization"`
	// +optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit"`
	// +optional
	Pseudonym *string `json:"pseudonym,omitempty" tf:"pseudonym"`
	// +optional
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number"`
	// +optional
	StateOrProvinceName *string `json:"stateOrProvinceName,omitempty" tf:"state_or_province_name"`
	// +optional
	Street *string `json:"street,omitempty" tf:"street"`
	// +optional
	Surname *string `json:"surname,omitempty" tf:"surname"`
	// +optional
	Title *string `json:"title,omitempty" tf:"title"`
	// +optional
	UserID *string `json:"userID,omitempty" tf:"user_id"`
}

type ManagementCertificateAuthoritySpecCertificateAuthorityConfigValidity struct {
	TimeOfValidityNotAfter *string `json:"timeOfValidityNotAfter" tf:"time_of_validity_not_after"`
	// +optional
	TimeOfValidityNotBefore *string `json:"timeOfValidityNotBefore,omitempty" tf:"time_of_validity_not_before"`
}

type ManagementCertificateAuthoritySpecCertificateAuthorityConfig struct {
	ConfigType *string `json:"configType" tf:"config_type"`
	// +optional
	IssuerCertificateAuthorityID *string `json:"issuerCertificateAuthorityID,omitempty" tf:"issuer_certificate_authority_id"`
	// +optional
	SigningAlgorithm *string                                                              `json:"signingAlgorithm,omitempty" tf:"signing_algorithm"`
	Subject          *ManagementCertificateAuthoritySpecCertificateAuthorityConfigSubject `json:"subject" tf:"subject"`
	// +optional
	Validity *ManagementCertificateAuthoritySpecCertificateAuthorityConfigValidity `json:"validity,omitempty" tf:"validity"`
	// +optional
	VersionName *string `json:"versionName,omitempty" tf:"version_name"`
}

type ManagementCertificateAuthoritySpecCertificateAuthorityRules struct {
	// +optional
	CertificateAuthorityMaxValidityDuration *string `json:"certificateAuthorityMaxValidityDuration,omitempty" tf:"certificate_authority_max_validity_duration"`
	// +optional
	LeafCertificateMaxValidityDuration *string `json:"leafCertificateMaxValidityDuration,omitempty" tf:"leaf_certificate_max_validity_duration"`
	RuleType                           *string `json:"ruleType" tf:"rule_type"`
}

type ManagementCertificateAuthoritySpecCertificateRevocationListDetailsObjectStorageConfig struct {
	ObjectStorageBucketName *string `json:"objectStorageBucketName" tf:"object_storage_bucket_name"`
	// +optional
	ObjectStorageNamespace        *string `json:"objectStorageNamespace,omitempty" tf:"object_storage_namespace"`
	ObjectStorageObjectNameFormat *string `json:"objectStorageObjectNameFormat" tf:"object_storage_object_name_format"`
}

type ManagementCertificateAuthoritySpecCertificateRevocationListDetails struct {
	// +optional
	CustomFormattedUrls []string                                                                               `json:"customFormattedUrls,omitempty" tf:"custom_formatted_urls"`
	ObjectStorageConfig *ManagementCertificateAuthoritySpecCertificateRevocationListDetailsObjectStorageConfig `json:"objectStorageConfig" tf:"object_storage_config"`
}

type ManagementCertificateAuthoritySpecCurrentVersionRevocationStatus struct {
	// +optional
	RevocationReason *string `json:"revocationReason,omitempty" tf:"revocation_reason"`
	// +optional
	TimeOfRevocation *string `json:"timeOfRevocation,omitempty" tf:"time_of_revocation"`
}

type ManagementCertificateAuthoritySpecCurrentVersionValidity struct {
	// +optional
	TimeOfValidityNotAfter *string `json:"timeOfValidityNotAfter,omitempty" tf:"time_of_validity_not_after"`
	// +optional
	TimeOfValidityNotBefore *string `json:"timeOfValidityNotBefore,omitempty" tf:"time_of_validity_not_before"`
}

type ManagementCertificateAuthoritySpecCurrentVersion struct {
	// +optional
	CertificateAuthorityID *string `json:"certificateAuthorityID,omitempty" tf:"certificate_authority_id"`
	// +optional
	IssuerCaVersionNumber *string `json:"issuerCaVersionNumber,omitempty" tf:"issuer_ca_version_number"`
	// +optional
	RevocationStatus *ManagementCertificateAuthoritySpecCurrentVersionRevocationStatus `json:"revocationStatus,omitempty" tf:"revocation_status"`
	// +optional
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number"`
	// +optional
	Stages []string `json:"stages,omitempty" tf:"stages"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeOfDeletion *string `json:"timeOfDeletion,omitempty" tf:"time_of_deletion"`
	// +optional
	Validity *ManagementCertificateAuthoritySpecCurrentVersionValidity `json:"validity,omitempty" tf:"validity"`
	// +optional
	VersionName *string `json:"versionName,omitempty" tf:"version_name"`
	// +optional
	VersionNumber *string `json:"versionNumber,omitempty" tf:"version_number"`
}

type ManagementCertificateAuthoritySpecSubject struct {
	// +optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name"`
	// +optional
	Country *string `json:"country,omitempty" tf:"country"`
	// +optional
	DistinguishedNameQualifier *string `json:"distinguishedNameQualifier,omitempty" tf:"distinguished_name_qualifier"`
	// +optional
	DomainComponent *string `json:"domainComponent,omitempty" tf:"domain_component"`
	// +optional
	GenerationQualifier *string `json:"generationQualifier,omitempty" tf:"generation_qualifier"`
	// +optional
	GivenName *string `json:"givenName,omitempty" tf:"given_name"`
	// +optional
	Initials *string `json:"initials,omitempty" tf:"initials"`
	// +optional
	LocalityName *string `json:"localityName,omitempty" tf:"locality_name"`
	// +optional
	Organization *string `json:"organization,omitempty" tf:"organization"`
	// +optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit"`
	// +optional
	Pseudonym *string `json:"pseudonym,omitempty" tf:"pseudonym"`
	// +optional
	SerialNumber *string `json:"serialNumber,omitempty" tf:"serial_number"`
	// +optional
	StateOrProvinceName *string `json:"stateOrProvinceName,omitempty" tf:"state_or_province_name"`
	// +optional
	Street *string `json:"street,omitempty" tf:"street"`
	// +optional
	Surname *string `json:"surname,omitempty" tf:"surname"`
	// +optional
	Title *string `json:"title,omitempty" tf:"title"`
	// +optional
	UserID *string `json:"userID,omitempty" tf:"user_id"`
}

type ManagementCertificateAuthoritySpec struct {
	State *ManagementCertificateAuthoritySpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagementCertificateAuthoritySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagementCertificateAuthoritySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CertificateAuthorityConfig *ManagementCertificateAuthoritySpecCertificateAuthorityConfig `json:"certificateAuthorityConfig" tf:"certificate_authority_config"`
	// +optional
	CertificateAuthorityRules []ManagementCertificateAuthoritySpecCertificateAuthorityRules `json:"certificateAuthorityRules,omitempty" tf:"certificate_authority_rules"`
	// +optional
	CertificateRevocationListDetails *ManagementCertificateAuthoritySpecCertificateRevocationListDetails `json:"certificateRevocationListDetails,omitempty" tf:"certificate_revocation_list_details"`
	CompartmentID                    *string                                                             `json:"compartmentID" tf:"compartment_id"`
	// +optional
	ConfigType *string `json:"configType,omitempty" tf:"config_type"`
	// +optional
	CurrentVersion *ManagementCertificateAuthoritySpecCurrentVersion `json:"currentVersion,omitempty" tf:"current_version"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	IssuerCertificateAuthorityID *string `json:"issuerCertificateAuthorityID,omitempty" tf:"issuer_certificate_authority_id"`
	KmsKeyID                     *string `json:"kmsKeyID" tf:"kms_key_id"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	Name             *string `json:"name" tf:"name"`
	// +optional
	SigningAlgorithm *string `json:"signingAlgorithm,omitempty" tf:"signing_algorithm"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	Subject *ManagementCertificateAuthoritySpecSubject `json:"subject,omitempty" tf:"subject"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeOfDeletion *string `json:"timeOfDeletion,omitempty" tf:"time_of_deletion"`
}

type ManagementCertificateAuthorityStatus struct {
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

// ManagementCertificateAuthorityList is a list of ManagementCertificateAuthoritys
type ManagementCertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagementCertificateAuthority CRD objects
	Items []ManagementCertificateAuthority `json:"items,omitempty"`
}
