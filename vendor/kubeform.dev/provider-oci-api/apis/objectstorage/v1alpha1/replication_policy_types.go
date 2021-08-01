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

type ReplicationPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationPolicySpec   `json:"spec,omitempty"`
	Status            ReplicationPolicyStatus `json:"status,omitempty"`
}

type ReplicationPolicySpec struct {
	State *ReplicationPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource ReplicationPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ReplicationPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Bucket *string `json:"bucket" tf:"bucket"`
	// +optional
	// Deprecated
	DeleteObjectInDestinationBucket *string `json:"deleteObjectInDestinationBucket,omitempty" tf:"delete_object_in_destination_bucket"`
	DestinationBucketName           *string `json:"destinationBucketName" tf:"destination_bucket_name"`
	DestinationRegionName           *string `json:"destinationRegionName" tf:"destination_region_name"`
	Name                            *string `json:"name" tf:"name"`
	Namespace                       *string `json:"namespace" tf:"namespace"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StatusMessage *string `json:"statusMessage,omitempty" tf:"status_message"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeLastSync *string `json:"timeLastSync,omitempty" tf:"time_last_sync"`
}

type ReplicationPolicyStatus struct {
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

// ReplicationPolicyList is a list of ReplicationPolicys
type ReplicationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ReplicationPolicy CRD objects
	Items []ReplicationPolicy `json:"items,omitempty"`
}
