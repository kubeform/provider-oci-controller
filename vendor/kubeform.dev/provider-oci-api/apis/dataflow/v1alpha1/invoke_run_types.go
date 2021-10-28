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

type InvokeRun struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InvokeRunSpec   `json:"spec,omitempty"`
	Status            InvokeRunStatus `json:"status,omitempty"`
}

type InvokeRunSpecParameters struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type InvokeRunSpec struct {
	State *InvokeRunSpecResource `json:"state,omitempty" tf:"-"`

	Resource InvokeRunSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InvokeRunSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApplicationID *string `json:"applicationID,omitempty" tf:"application_id"`
	// +optional
	ArchiveURI *string `json:"archiveURI,omitempty" tf:"archive_uri"`
	// +optional
	Arguments []string `json:"arguments,omitempty" tf:"arguments"`
	// +optional
	Asynchronous *bool `json:"asynchronous,omitempty" tf:"asynchronous"`
	// +optional
	ClassName     *string `json:"className,omitempty" tf:"class_name"`
	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	Configuration map[string]string `json:"configuration,omitempty" tf:"configuration"`
	// +optional
	DataReadInBytes *string `json:"dataReadInBytes,omitempty" tf:"data_read_in_bytes"`
	// +optional
	DataWrittenInBytes *string `json:"dataWrittenInBytes,omitempty" tf:"data_written_in_bytes"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	DriverShape *string `json:"driverShape,omitempty" tf:"driver_shape"`
	// +optional
	Execute *string `json:"execute,omitempty" tf:"execute"`
	// +optional
	ExecutorShape *string `json:"executorShape,omitempty" tf:"executor_shape"`
	// +optional
	FileURI *string `json:"fileURI,omitempty" tf:"file_uri"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	Language *string `json:"language,omitempty" tf:"language"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	LogsBucketURI *string `json:"logsBucketURI,omitempty" tf:"logs_bucket_uri"`
	// +optional
	MetastoreID *string `json:"metastoreID,omitempty" tf:"metastore_id"`
	// +optional
	NumExecutors *int64 `json:"numExecutors,omitempty" tf:"num_executors"`
	// +optional
	OpcRequestID *string `json:"opcRequestID,omitempty" tf:"opc_request_id"`
	// +optional
	OwnerPrincipalID *string `json:"ownerPrincipalID,omitempty" tf:"owner_principal_id"`
	// +optional
	OwnerUserName *string `json:"ownerUserName,omitempty" tf:"owner_user_name"`
	// +optional
	Parameters []InvokeRunSpecParameters `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	PrivateEndpointDNSZones []string `json:"privateEndpointDNSZones,omitempty" tf:"private_endpoint_dns_zones"`
	// +optional
	PrivateEndpointID *string `json:"privateEndpointID,omitempty" tf:"private_endpoint_id"`
	// +optional
	PrivateEndpointMaxHostCount *int64 `json:"privateEndpointMaxHostCount,omitempty" tf:"private_endpoint_max_host_count"`
	// +optional
	PrivateEndpointNsgIDS []string `json:"privateEndpointNsgIDS,omitempty" tf:"private_endpoint_nsg_ids"`
	// +optional
	PrivateEndpointSubnetID *string `json:"privateEndpointSubnetID,omitempty" tf:"private_endpoint_subnet_id"`
	// +optional
	RunDurationInMilliseconds *string `json:"runDurationInMilliseconds,omitempty" tf:"run_duration_in_milliseconds"`
	// +optional
	SparkVersion *string `json:"sparkVersion,omitempty" tf:"spark_version"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	TotalOcpu *int64 `json:"totalOcpu,omitempty" tf:"total_ocpu"`
	// +optional
	WarehouseBucketURI *string `json:"warehouseBucketURI,omitempty" tf:"warehouse_bucket_uri"`
}

type InvokeRunStatus struct {
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

// InvokeRunList is a list of InvokeRuns
type InvokeRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InvokeRun CRD objects
	Items []InvokeRun `json:"items,omitempty"`
}
