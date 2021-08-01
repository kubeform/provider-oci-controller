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

type Recommendation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecommendationSpec   `json:"spec,omitempty"`
	Status            RecommendationStatus `json:"status,omitempty"`
}

type RecommendationSpecResourceCounts struct {
	// +optional
	Count *int64 `json:"count,omitempty" tf:"count"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type RecommendationSpecSupportedLevelsItems struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type RecommendationSpecSupportedLevels struct {
	// +optional
	Items []RecommendationSpecSupportedLevelsItems `json:"items,omitempty" tf:"items"`
}

type RecommendationSpec struct {
	State *RecommendationSpecResource `json:"state,omitempty" tf:"-"`

	Resource RecommendationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RecommendationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CategoryID *string `json:"categoryID,omitempty" tf:"category_id"`
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	EstimatedCostSaving *float64 `json:"estimatedCostSaving,omitempty" tf:"estimated_cost_saving"`
	// +optional
	Importance *string `json:"importance,omitempty" tf:"importance"`
	// +optional
	Name             *string `json:"name,omitempty" tf:"name"`
	RecommendationID *string `json:"recommendationID" tf:"recommendation_id"`
	// +optional
	ResourceCounts []RecommendationSpecResourceCounts `json:"resourceCounts,omitempty" tf:"resource_counts"`
	// +optional
	State  *string `json:"state,omitempty" tf:"state"`
	Status *string `json:"status" tf:"status"`
	// +optional
	SupportedLevels *RecommendationSpecSupportedLevels `json:"supportedLevels,omitempty" tf:"supported_levels"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeStatusBegin *string `json:"timeStatusBegin,omitempty" tf:"time_status_begin"`
	// +optional
	TimeStatusEnd *string `json:"timeStatusEnd,omitempty" tf:"time_status_end"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type RecommendationStatus struct {
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

// RecommendationList is a list of Recommendations
type RecommendationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Recommendation CRD objects
	Items []Recommendation `json:"items,omitempty"`
}