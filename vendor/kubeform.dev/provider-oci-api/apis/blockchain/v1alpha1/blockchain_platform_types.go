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

type BlockchainPlatform struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BlockchainPlatformSpec   `json:"spec,omitempty"`
	Status            BlockchainPlatformStatus `json:"status,omitempty"`
}

type BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam struct {
	// +optional
	OcpuAllocationNumber *float64 `json:"ocpuAllocationNumber,omitempty" tf:"ocpu_allocation_number"`
}

type BlockchainPlatformSpecComponentDetailsOsns struct {
	// +optional
	Ad *string `json:"ad,omitempty" tf:"ad"`
	// +optional
	OcpuAllocationParam *BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam `json:"ocpuAllocationParam,omitempty" tf:"ocpu_allocation_param"`
	// +optional
	OsnKey *string `json:"osnKey,omitempty" tf:"osn_key"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam struct {
	// +optional
	OcpuAllocationNumber *float64 `json:"ocpuAllocationNumber,omitempty" tf:"ocpu_allocation_number"`
}

type BlockchainPlatformSpecComponentDetailsPeers struct {
	// +optional
	Ad *string `json:"ad,omitempty" tf:"ad"`
	// +optional
	Alias *string `json:"alias,omitempty" tf:"alias"`
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// +optional
	OcpuAllocationParam *BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam `json:"ocpuAllocationParam,omitempty" tf:"ocpu_allocation_param"`
	// +optional
	PeerKey *string `json:"peerKey,omitempty" tf:"peer_key"`
	// +optional
	Role *string `json:"role,omitempty" tf:"role"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type BlockchainPlatformSpecComponentDetails struct {
	// +optional
	Osns []BlockchainPlatformSpecComponentDetailsOsns `json:"osns,omitempty" tf:"osns"`
	// +optional
	Peers []BlockchainPlatformSpecComponentDetailsPeers `json:"peers,omitempty" tf:"peers"`
}

type BlockchainPlatformSpecHostOcpuUtilizationInfo struct {
	// +optional
	Host *string `json:"host,omitempty" tf:"host"`
	// +optional
	OcpuCapacityNumber *float64 `json:"ocpuCapacityNumber,omitempty" tf:"ocpu_capacity_number"`
	// +optional
	OcpuUtilizationNumber *float64 `json:"ocpuUtilizationNumber,omitempty" tf:"ocpu_utilization_number"`
}

type BlockchainPlatformSpecReplicas struct {
	// +optional
	CaCount *int64 `json:"caCount,omitempty" tf:"ca_count"`
	// +optional
	ConsoleCount *int64 `json:"consoleCount,omitempty" tf:"console_count"`
	// +optional
	ProxyCount *int64 `json:"proxyCount,omitempty" tf:"proxy_count"`
}

type BlockchainPlatformSpec struct {
	State *BlockchainPlatformSpecResource `json:"state,omitempty" tf:"-"`

	Resource BlockchainPlatformSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BlockchainPlatformSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CaCertArchiveText *string `json:"caCertArchiveText,omitempty" tf:"ca_cert_archive_text"`
	CompartmentID     *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	ComponentDetails *BlockchainPlatformSpecComponentDetails `json:"componentDetails,omitempty" tf:"component_details"`
	ComputeShape     *string                                 `json:"computeShape" tf:"compute_shape"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	// +optional
	FederatedUserID *string `json:"federatedUserID,omitempty" tf:"federated_user_id"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	HostOcpuUtilizationInfo []BlockchainPlatformSpecHostOcpuUtilizationInfo `json:"hostOcpuUtilizationInfo,omitempty" tf:"host_ocpu_utilization_info"`
	IdcsAccessToken         *string                                         `json:"-" sensitive:"true" tf:"idcs_access_token"`
	// +optional
	IsByol *bool `json:"isByol,omitempty" tf:"is_byol"`
	// +optional
	IsMultiAd *bool `json:"isMultiAd,omitempty" tf:"is_multi_ad"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	LoadBalancerShape *string `json:"loadBalancerShape,omitempty" tf:"load_balancer_shape"`
	PlatformRole      *string `json:"platformRole" tf:"platform_role"`
	// +optional
	PlatformShapeType *string `json:"platformShapeType,omitempty" tf:"platform_shape_type"`
	// +optional
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version"`
	// +optional
	Replicas *BlockchainPlatformSpecReplicas `json:"replicas,omitempty" tf:"replicas"`
	// +optional
	ServiceEndpoint *string `json:"serviceEndpoint,omitempty" tf:"service_endpoint"`
	// +optional
	ServiceVersion *string `json:"serviceVersion,omitempty" tf:"service_version"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	StorageSizeInTbs *float64 `json:"storageSizeInTbs,omitempty" tf:"storage_size_in_tbs"`
	// +optional
	StorageUsedInTbs *float64 `json:"storageUsedInTbs,omitempty" tf:"storage_used_in_tbs"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	TotalOcpuCapacity *int64 `json:"totalOcpuCapacity,omitempty" tf:"total_ocpu_capacity"`
}

type BlockchainPlatformStatus struct {
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

// BlockchainPlatformList is a list of BlockchainPlatforms
type BlockchainPlatformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BlockchainPlatform CRD objects
	Items []BlockchainPlatform `json:"items,omitempty"`
}
