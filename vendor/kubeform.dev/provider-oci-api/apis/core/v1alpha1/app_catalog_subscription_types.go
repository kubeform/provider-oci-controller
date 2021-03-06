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

type AppCatalogSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppCatalogSubscriptionSpec   `json:"spec,omitempty"`
	Status            AppCatalogSubscriptionStatus `json:"status,omitempty"`
}

type AppCatalogSubscriptionSpec struct {
	State *AppCatalogSubscriptionSpecResource `json:"state,omitempty" tf:"-"`

	Resource AppCatalogSubscriptionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AppCatalogSubscriptionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	EulaLink  *string `json:"eulaLink,omitempty" tf:"eula_link"`
	ListingID *string `json:"listingID" tf:"listing_id"`
	// +optional
	ListingResourceID      *string `json:"listingResourceID,omitempty" tf:"listing_resource_id"`
	ListingResourceVersion *string `json:"listingResourceVersion" tf:"listing_resource_version"`
	OracleTermsOfUseLink   *string `json:"oracleTermsOfUseLink" tf:"oracle_terms_of_use_link"`
	// +optional
	PublisherName *string `json:"publisherName,omitempty" tf:"publisher_name"`
	Signature     *string `json:"signature" tf:"signature"`
	// +optional
	Summary *string `json:"summary,omitempty" tf:"summary"`
	// +optional
	TimeCreated   *string `json:"timeCreated,omitempty" tf:"time_created"`
	TimeRetrieved *string `json:"timeRetrieved" tf:"time_retrieved"`
}

type AppCatalogSubscriptionStatus struct {
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

// AppCatalogSubscriptionList is a list of AppCatalogSubscriptions
type AppCatalogSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppCatalogSubscription CRD objects
	Items []AppCatalogSubscription `json:"items,omitempty"`
}
