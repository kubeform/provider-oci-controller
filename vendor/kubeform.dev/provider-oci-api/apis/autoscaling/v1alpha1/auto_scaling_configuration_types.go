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

type AutoScalingConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoScalingConfigurationSpec   `json:"spec,omitempty"`
	Status            AutoScalingConfigurationStatus `json:"status,omitempty"`
}

type AutoScalingConfigurationSpecAutoScalingResources struct {
	ID   *string `json:"ID" tf:"id"`
	Type *string `json:"type" tf:"type"`
}

type AutoScalingConfigurationSpecPoliciesCapacity struct {
	// +optional
	Initial *int64 `json:"initial,omitempty" tf:"initial"`
	// +optional
	Max *int64 `json:"max,omitempty" tf:"max"`
	// +optional
	Min *int64 `json:"min,omitempty" tf:"min"`
}

type AutoScalingConfigurationSpecPoliciesExecutionSchedule struct {
	Expression *string `json:"expression" tf:"expression"`
	Timezone   *string `json:"timezone" tf:"timezone"`
	Type       *string `json:"type" tf:"type"`
}

type AutoScalingConfigurationSpecPoliciesResourceAction struct {
	Action     *string `json:"action" tf:"action"`
	ActionType *string `json:"actionType" tf:"action_type"`
}

type AutoScalingConfigurationSpecPoliciesRulesAction struct {
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// +optional
	Value *int64 `json:"value,omitempty" tf:"value"`
}

type AutoScalingConfigurationSpecPoliciesRulesMetricThreshold struct {
	// +optional
	Operator *string `json:"operator,omitempty" tf:"operator"`
	// +optional
	Value *int64 `json:"value,omitempty" tf:"value"`
}

type AutoScalingConfigurationSpecPoliciesRulesMetric struct {
	// +optional
	MetricType *string `json:"metricType,omitempty" tf:"metric_type"`
	// +optional
	Threshold *AutoScalingConfigurationSpecPoliciesRulesMetricThreshold `json:"threshold,omitempty" tf:"threshold"`
}

type AutoScalingConfigurationSpecPoliciesRules struct {
	// +optional
	Action      *AutoScalingConfigurationSpecPoliciesRulesAction `json:"action,omitempty" tf:"action"`
	DisplayName *string                                          `json:"displayName" tf:"display_name"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Metric *AutoScalingConfigurationSpecPoliciesRulesMetric `json:"metric,omitempty" tf:"metric"`
}

type AutoScalingConfigurationSpecPolicies struct {
	// +optional
	Capacity *AutoScalingConfigurationSpecPoliciesCapacity `json:"capacity,omitempty" tf:"capacity"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	ExecutionSchedule *AutoScalingConfigurationSpecPoliciesExecutionSchedule `json:"executionSchedule,omitempty" tf:"execution_schedule"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	IsEnabled  *bool   `json:"isEnabled,omitempty" tf:"is_enabled"`
	PolicyType *string `json:"policyType" tf:"policy_type"`
	// +optional
	ResourceAction *AutoScalingConfigurationSpecPoliciesResourceAction `json:"resourceAction,omitempty" tf:"resource_action"`
	// +optional
	Rules []AutoScalingConfigurationSpecPoliciesRules `json:"rules,omitempty" tf:"rules"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
}

type AutoScalingConfigurationSpec struct {
	State *AutoScalingConfigurationSpecResource `json:"state,omitempty" tf:"-"`

	Resource AutoScalingConfigurationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type AutoScalingConfigurationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AutoScalingResources *AutoScalingConfigurationSpecAutoScalingResources `json:"autoScalingResources" tf:"auto_scaling_resources"`
	CompartmentID        *string                                           `json:"compartmentID" tf:"compartment_id"`
	// +optional
	CoolDownInSeconds *int64 `json:"coolDownInSeconds,omitempty" tf:"cool_down_in_seconds"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
	// +optional
	MaxResourceCount *int64 `json:"maxResourceCount,omitempty" tf:"max_resource_count"`
	// +optional
	MinResourceCount *int64                                 `json:"minResourceCount,omitempty" tf:"min_resource_count"`
	Policies         []AutoScalingConfigurationSpecPolicies `json:"policies" tf:"policies"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
}

type AutoScalingConfigurationStatus struct {
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

// AutoScalingConfigurationList is a list of AutoScalingConfigurations
type AutoScalingConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoScalingConfiguration CRD objects
	Items []AutoScalingConfiguration `json:"items,omitempty"`
}
