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

type ComputationCustomTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputationCustomTableSpec   `json:"spec,omitempty"`
	Status            ComputationCustomTableStatus `json:"status,omitempty"`
}

type ComputationCustomTableSpecSavedCustomTableGroupByTag struct {
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ComputationCustomTableSpecSavedCustomTable struct {
	// +optional
	ColumnGroupBy []string `json:"columnGroupBy,omitempty" tf:"column_group_by"`
	// +optional
	CompartmentDepth *float64 `json:"compartmentDepth,omitempty" tf:"compartment_depth"`
	DisplayName      *string  `json:"displayName" tf:"display_name"`
	// +optional
	GroupByTag []ComputationCustomTableSpecSavedCustomTableGroupByTag `json:"groupByTag,omitempty" tf:"group_by_tag"`
	// +optional
	RowGroupBy []string `json:"rowGroupBy,omitempty" tf:"row_group_by"`
	// +optional
	Version *float64 `json:"version,omitempty" tf:"version"`
}

type ComputationCustomTableSpec struct {
	State *ComputationCustomTableSpecResource `json:"state,omitempty" tf:"-"`

	Resource ComputationCustomTableSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ComputationCustomTableSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID    *string                                     `json:"compartmentID" tf:"compartment_id"`
	SavedCustomTable *ComputationCustomTableSpecSavedCustomTable `json:"savedCustomTable" tf:"saved_custom_table"`
	SavedReportID    *string                                     `json:"savedReportID" tf:"saved_report_id"`
}

type ComputationCustomTableStatus struct {
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

// ComputationCustomTableList is a list of ComputationCustomTables
type ComputationCustomTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputationCustomTable CRD objects
	Items []ComputationCustomTable `json:"items,omitempty"`
}
