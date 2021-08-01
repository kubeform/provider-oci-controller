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

type GuardDataMaskRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuardDataMaskRuleSpec   `json:"spec,omitempty"`
	Status            GuardDataMaskRuleStatus `json:"status,omitempty"`
}

type GuardDataMaskRuleSpecTargetSelected struct {
	Kind *string `json:"kind" tf:"kind"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type GuardDataMaskRuleSpec struct {
	State *GuardDataMaskRuleSpecResource `json:"state,omitempty" tf:"-"`

	Resource GuardDataMaskRuleSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GuardDataMaskRuleSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID      *string  `json:"compartmentID" tf:"compartment_id"`
	DataMaskCategories []string `json:"dataMaskCategories" tf:"data_mask_categories"`
	// +optional
	DataMaskRuleStatus *string `json:"dataMaskRuleStatus,omitempty" tf:"data_mask_rule_status"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	IamGroupID   *string           `json:"iamGroupID" tf:"iam_group_id"`
	// +optional
	LifecyleDetails *string `json:"lifecyleDetails,omitempty" tf:"lifecyle_details"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SystemTags     map[string]string                    `json:"systemTags,omitempty" tf:"system_tags"`
	TargetSelected *GuardDataMaskRuleSpecTargetSelected `json:"targetSelected" tf:"target_selected"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type GuardDataMaskRuleStatus struct {
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

// GuardDataMaskRuleList is a list of GuardDataMaskRules
type GuardDataMaskRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GuardDataMaskRule CRD objects
	Items []GuardDataMaskRule `json:"items,omitempty"`
}
