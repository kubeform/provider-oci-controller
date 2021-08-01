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

type StorageMountTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageMountTargetSpec   `json:"spec,omitempty"`
	Status            StorageMountTargetStatus `json:"status,omitempty"`
}

type StorageMountTargetSpec struct {
	State *StorageMountTargetSpecResource `json:"state,omitempty" tf:"-"`

	Resource StorageMountTargetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type StorageMountTargetSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AvailabilityDomain *string `json:"availabilityDomain" tf:"availability_domain"`
	CompartmentID      *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	ExportSetID *string `json:"exportSetID,omitempty" tf:"export_set_id"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	HostnameLabel *string `json:"hostnameLabel,omitempty" tf:"hostname_label"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	NsgIDS []string `json:"nsgIDS,omitempty" tf:"nsg_ids"`
	// +optional
	PrivateIPIDS []string `json:"privateIPIDS,omitempty" tf:"private_ip_ids"`
	// +optional
	State    *string `json:"state,omitempty" tf:"state"`
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
}

type StorageMountTargetStatus struct {
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

// StorageMountTargetList is a list of StorageMountTargets
type StorageMountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageMountTarget CRD objects
	Items []StorageMountTarget `json:"items,omitempty"`
}
