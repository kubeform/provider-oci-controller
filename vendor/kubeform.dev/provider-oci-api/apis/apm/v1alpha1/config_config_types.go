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

type ConfigConfig struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigSpec   `json:"spec,omitempty"`
	Status            ConfigConfigStatus `json:"status,omitempty"`
}

type ConfigConfigSpecDimensions struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	ValueSource *string `json:"valueSource,omitempty" tf:"value_source"`
}

type ConfigConfigSpecMetrics struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Unit *string `json:"unit,omitempty" tf:"unit"`
	// +optional
	ValueSource *string `json:"valueSource,omitempty" tf:"value_source"`
}

type ConfigConfigSpecRules struct {
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FilterText *string `json:"filterText,omitempty" tf:"filter_text"`
	// +optional
	IsApplyToErrorSpans *bool `json:"isApplyToErrorSpans,omitempty" tf:"is_apply_to_error_spans"`
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	SatisfiedResponseTime *int64 `json:"satisfiedResponseTime,omitempty" tf:"satisfied_response_time"`
	// +optional
	ToleratingResponseTime *int64 `json:"toleratingResponseTime,omitempty" tf:"tolerating_response_time"`
}

type ConfigConfigSpec struct {
	State *ConfigConfigSpecResource `json:"state,omitempty" tf:"-"`

	Resource ConfigConfigSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ConfigConfigSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApmDomainID *string `json:"apmDomainID" tf:"apm_domain_id"`
	ConfigType  *string `json:"configType" tf:"config_type"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Dimensions  []ConfigConfigSpecDimensions `json:"dimensions,omitempty" tf:"dimensions"`
	DisplayName *string                      `json:"displayName" tf:"display_name"`
	// +optional
	FilterID *string `json:"filterID,omitempty" tf:"filter_id"`
	// +optional
	FilterText *string `json:"filterText,omitempty" tf:"filter_text"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	Metrics []ConfigConfigSpecMetrics `json:"metrics,omitempty" tf:"metrics"`
	// +optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
	// +optional
	OpcDryRun *string `json:"opcDryRun,omitempty" tf:"opc_dry_run"`
	// +optional
	Rules []ConfigConfigSpecRules `json:"rules,omitempty" tf:"rules"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type ConfigConfigStatus struct {
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

// ConfigConfigList is a list of ConfigConfigs
type ConfigConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfig CRD objects
	Items []ConfigConfig `json:"items,omitempty"`
}
