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

type ServiceConnector struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceConnectorSpec   `json:"spec,omitempty"`
	Status            ServiceConnectorStatus `json:"status,omitempty"`
}

type ServiceConnectorSpecSourceCursor struct {
	// +optional
	Kind *string `json:"kind,omitempty" tf:"kind"`
}

type ServiceConnectorSpecSourceLogSources struct {
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	LogGroupID *string `json:"logGroupID,omitempty" tf:"log_group_id"`
	// +optional
	LogID *string `json:"logID,omitempty" tf:"log_id"`
}

type ServiceConnectorSpecSource struct {
	// +optional
	Cursor *ServiceConnectorSpecSourceCursor `json:"cursor,omitempty" tf:"cursor"`
	Kind   *string                           `json:"kind" tf:"kind"`
	// +optional
	LogSources []ServiceConnectorSpecSourceLogSources `json:"logSources,omitempty" tf:"log_sources"`
	// +optional
	StreamID *string `json:"streamID,omitempty" tf:"stream_id"`
}

type ServiceConnectorSpecTargetDimensionsDimensionValue struct {
	Kind *string `json:"kind" tf:"kind"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ServiceConnectorSpecTargetDimensions struct {
	// +optional
	DimensionValue *ServiceConnectorSpecTargetDimensionsDimensionValue `json:"dimensionValue,omitempty" tf:"dimension_value"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type ServiceConnectorSpecTarget struct {
	// +optional
	BatchRolloverSizeInMbs *int64 `json:"batchRolloverSizeInMbs,omitempty" tf:"batch_rollover_size_in_mbs"`
	// +optional
	BatchRolloverTimeInMs *int64 `json:"batchRolloverTimeInMs,omitempty" tf:"batch_rollover_time_in_ms"`
	// +optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket"`
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	Dimensions []ServiceConnectorSpecTargetDimensions `json:"dimensions,omitempty" tf:"dimensions"`
	// +optional
	EnableFormattedMessaging *bool `json:"enableFormattedMessaging,omitempty" tf:"enable_formatted_messaging"`
	// +optional
	FunctionID *string `json:"functionID,omitempty" tf:"function_id"`
	Kind       *string `json:"kind" tf:"kind"`
	// +optional
	LogGroupID *string `json:"logGroupID,omitempty" tf:"log_group_id"`
	// +optional
	Metric *string `json:"metric,omitempty" tf:"metric"`
	// +optional
	MetricNamespace *string `json:"metricNamespace,omitempty" tf:"metric_namespace"`
	// +optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
	// +optional
	ObjectNamePrefix *string `json:"objectNamePrefix,omitempty" tf:"object_name_prefix"`
	// +optional
	StreamID *string `json:"streamID,omitempty" tf:"stream_id"`
	// +optional
	TopicID *string `json:"topicID,omitempty" tf:"topic_id"`
}

type ServiceConnectorSpecTasks struct {
	// +optional
	BatchSizeInKbs *int64 `json:"batchSizeInKbs,omitempty" tf:"batch_size_in_kbs"`
	// +optional
	BatchTimeInSec *int64 `json:"batchTimeInSec,omitempty" tf:"batch_time_in_sec"`
	// +optional
	Condition *string `json:"condition,omitempty" tf:"condition"`
	// +optional
	FunctionID *string `json:"functionID,omitempty" tf:"function_id"`
	Kind       *string `json:"kind" tf:"kind"`
}

type ServiceConnectorSpec struct {
	State *ServiceConnectorSpecResource `json:"state,omitempty" tf:"-"`

	Resource ServiceConnectorSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ServiceConnectorSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	LifecyleDetails *string                     `json:"lifecyleDetails,omitempty" tf:"lifecyle_details"`
	Source          *ServiceConnectorSpecSource `json:"source" tf:"source"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SystemTags map[string]string           `json:"systemTags,omitempty" tf:"system_tags"`
	Target     *ServiceConnectorSpecTarget `json:"target" tf:"target"`
	// +optional
	Tasks []ServiceConnectorSpecTasks `json:"tasks,omitempty" tf:"tasks"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type ServiceConnectorStatus struct {
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

// ServiceConnectorList is a list of ServiceConnectors
type ServiceConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceConnector CRD objects
	Items []ServiceConnector `json:"items,omitempty"`
}