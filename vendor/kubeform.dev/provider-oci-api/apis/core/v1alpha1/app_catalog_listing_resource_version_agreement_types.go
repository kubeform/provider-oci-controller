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

type AppCatalogListingResourceVersionAgreement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppCatalogListingResourceVersionAgreementSpec   `json:"spec,omitempty"`
	Status            AppCatalogListingResourceVersionAgreementStatus `json:"status,omitempty"`
}

type AppCatalogListingResourceVersionAgreementSpec struct {
	State *AppCatalogListingResourceVersionAgreementSpecResource `json:"state,omitempty" tf:"-"`

	Resource AppCatalogListingResourceVersionAgreementSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AppCatalogListingResourceVersionAgreementSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	EulaLink               *string `json:"eulaLink,omitempty" tf:"eula_link"`
	ListingID              *string `json:"listingID" tf:"listing_id"`
	ListingResourceVersion *string `json:"listingResourceVersion" tf:"listing_resource_version"`
	// +optional
	OracleTermsOfUseLink *string `json:"oracleTermsOfUseLink,omitempty" tf:"oracle_terms_of_use_link"`
	// +optional
	Signature *string `json:"signature,omitempty" tf:"signature"`
	// +optional
	TimeRetrieved *string `json:"timeRetrieved,omitempty" tf:"time_retrieved"`
}

type AppCatalogListingResourceVersionAgreementStatus struct {
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

// AppCatalogListingResourceVersionAgreementList is a list of AppCatalogListingResourceVersionAgreements
type AppCatalogListingResourceVersionAgreementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppCatalogListingResourceVersionAgreement CRD objects
	Items []AppCatalogListingResourceVersionAgreement `json:"items,omitempty"`
}
