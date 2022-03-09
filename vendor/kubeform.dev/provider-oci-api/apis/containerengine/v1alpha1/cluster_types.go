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

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpecEndpointConfig struct {
	// +optional
	IsPublicIPEnabled *bool `json:"isPublicIPEnabled,omitempty" tf:"is_public_ip_enabled"`
	// +optional
	NsgIDS   []string `json:"nsgIDS,omitempty" tf:"nsg_ids"`
	SubnetID *string  `json:"subnetID" tf:"subnet_id"`
}

type ClusterSpecEndpoints struct {
	// +optional
	Kubernetes *string `json:"kubernetes,omitempty" tf:"kubernetes"`
	// +optional
	PrivateEndpoint *string `json:"privateEndpoint,omitempty" tf:"private_endpoint"`
	// +optional
	PublicEndpoint *string `json:"publicEndpoint,omitempty" tf:"public_endpoint"`
}

type ClusterSpecImagePolicyConfigKeyDetails struct {
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
}

type ClusterSpecImagePolicyConfig struct {
	// +optional
	IsPolicyEnabled *bool `json:"isPolicyEnabled,omitempty" tf:"is_policy_enabled"`
	// +optional
	KeyDetails []ClusterSpecImagePolicyConfigKeyDetails `json:"keyDetails,omitempty" tf:"key_details"`
}

type ClusterSpecMetadata struct {
	// +optional
	CreatedByUserID *string `json:"createdByUserID,omitempty" tf:"created_by_user_id"`
	// +optional
	CreatedByWorkRequestID *string `json:"createdByWorkRequestID,omitempty" tf:"created_by_work_request_id"`
	// +optional
	DeletedByUserID *string `json:"deletedByUserID,omitempty" tf:"deleted_by_user_id"`
	// +optional
	DeletedByWorkRequestID *string `json:"deletedByWorkRequestID,omitempty" tf:"deleted_by_work_request_id"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeDeleted *string `json:"timeDeleted,omitempty" tf:"time_deleted"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	UpdatedByUserID *string `json:"updatedByUserID,omitempty" tf:"updated_by_user_id"`
	// +optional
	UpdatedByWorkRequestID *string `json:"updatedByWorkRequestID,omitempty" tf:"updated_by_work_request_id"`
}

type ClusterSpecOptionsAddOns struct {
	// +optional
	IsKubernetesDashboardEnabled *bool `json:"isKubernetesDashboardEnabled,omitempty" tf:"is_kubernetes_dashboard_enabled"`
	// +optional
	IsTillerEnabled *bool `json:"isTillerEnabled,omitempty" tf:"is_tiller_enabled"`
}

type ClusterSpecOptionsAdmissionControllerOptions struct {
	// +optional
	IsPodSecurityPolicyEnabled *bool `json:"isPodSecurityPolicyEnabled,omitempty" tf:"is_pod_security_policy_enabled"`
}

type ClusterSpecOptionsKubernetesNetworkConfig struct {
	// +optional
	PodsCIDR *string `json:"podsCIDR,omitempty" tf:"pods_cidr"`
	// +optional
	ServicesCIDR *string `json:"servicesCIDR,omitempty" tf:"services_cidr"`
}

type ClusterSpecOptionsPersistentVolumeConfig struct {
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
}

type ClusterSpecOptionsServiceLbConfig struct {
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
}

type ClusterSpecOptions struct {
	// +optional
	AddOns *ClusterSpecOptionsAddOns `json:"addOns,omitempty" tf:"add_ons"`
	// +optional
	AdmissionControllerOptions *ClusterSpecOptionsAdmissionControllerOptions `json:"admissionControllerOptions,omitempty" tf:"admission_controller_options"`
	// +optional
	KubernetesNetworkConfig *ClusterSpecOptionsKubernetesNetworkConfig `json:"kubernetesNetworkConfig,omitempty" tf:"kubernetes_network_config"`
	// +optional
	PersistentVolumeConfig *ClusterSpecOptionsPersistentVolumeConfig `json:"persistentVolumeConfig,omitempty" tf:"persistent_volume_config"`
	// +optional
	ServiceLbConfig *ClusterSpecOptionsServiceLbConfig `json:"serviceLbConfig,omitempty" tf:"service_lb_config"`
	// +optional
	ServiceLbSubnetIDS []string `json:"serviceLbSubnetIDS,omitempty" tf:"service_lb_subnet_ids"`
}

type ClusterSpec struct {
	State *ClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AvailableKubernetesUpgrades []string `json:"availableKubernetesUpgrades,omitempty" tf:"available_kubernetes_upgrades"`
	CompartmentID               *string  `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	EndpointConfig *ClusterSpecEndpointConfig `json:"endpointConfig,omitempty" tf:"endpoint_config"`
	// +optional
	Endpoints *ClusterSpecEndpoints `json:"endpoints,omitempty" tf:"endpoints"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	ImagePolicyConfig *ClusterSpecImagePolicyConfig `json:"imagePolicyConfig,omitempty" tf:"image_policy_config"`
	// +optional
	KmsKeyID          *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	KubernetesVersion *string `json:"kubernetesVersion" tf:"kubernetes_version"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	Metadata *ClusterSpecMetadata `json:"metadata,omitempty" tf:"metadata"`
	Name     *string              `json:"name" tf:"name"`
	// +optional
	Options *ClusterSpecOptions `json:"options,omitempty" tf:"options"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	VcnID      *string           `json:"vcnID" tf:"vcn_id"`
}

type ClusterStatus struct {
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

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
