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

type ChecksPingMonitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ChecksPingMonitorSpec   `json:"spec,omitempty"`
	Status            ChecksPingMonitorStatus `json:"status,omitempty"`
}

type ChecksPingMonitorSpec struct {
	State *ChecksPingMonitorSpecResource `json:"state,omitempty" tf:"-"`

	Resource ChecksPingMonitorSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ChecksPingMonitorSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	DisplayName *string           `json:"displayName" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	HomeRegion        *string `json:"homeRegion,omitempty" tf:"home_region"`
	IntervalInSeconds *int64  `json:"intervalInSeconds" tf:"interval_in_seconds"`
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
	// +optional
	Port     *int64  `json:"port,omitempty" tf:"port"`
	Protocol *string `json:"protocol" tf:"protocol"`
	// +optional
	ResultsURL *string  `json:"resultsURL,omitempty" tf:"results_url"`
	Targets    []string `json:"targets" tf:"targets"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds"`
	// +optional
	VantagePointNames []string `json:"vantagePointNames,omitempty" tf:"vantage_point_names"`
}

type ChecksPingMonitorStatus struct {
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

// ChecksPingMonitorList is a list of ChecksPingMonitors
type ChecksPingMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ChecksPingMonitor CRD objects
	Items []ChecksPingMonitor `json:"items,omitempty"`
}