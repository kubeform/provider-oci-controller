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

type ClusterNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterNetworkSpec   `json:"spec,omitempty"`
	Status            ClusterNetworkStatus `json:"status,omitempty"`
}

type ClusterNetworkSpecInstancePoolsLoadBalancers struct {
	// +optional
	BackendSetName *string `json:"backendSetName,omitempty" tf:"backend_set_name"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InstancePoolID *string `json:"instancePoolID,omitempty" tf:"instance_pool_id"`
	// +optional
	LoadBalancerID *string `json:"loadBalancerID,omitempty" tf:"load_balancer_id"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	VnicSelection *string `json:"vnicSelection,omitempty" tf:"vnic_selection"`
}

type ClusterNetworkSpecInstancePoolsPlacementConfigurationsSecondaryVnicSubnets struct {
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
}

type ClusterNetworkSpecInstancePoolsPlacementConfigurations struct {
	// +optional
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain"`
	// +optional
	FaultDomains []string `json:"faultDomains,omitempty" tf:"fault_domains"`
	// +optional
	PrimarySubnetID *string `json:"primarySubnetID,omitempty" tf:"primary_subnet_id"`
	// +optional
	SecondaryVnicSubnets []ClusterNetworkSpecInstancePoolsPlacementConfigurationsSecondaryVnicSubnets `json:"secondaryVnicSubnets,omitempty" tf:"secondary_vnic_subnets"`
}

type ClusterNetworkSpecInstancePools struct {
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	ID                      *string `json:"ID,omitempty" tf:"id"`
	InstanceConfigurationID *string `json:"instanceConfigurationID" tf:"instance_configuration_id"`
	// +optional
	LoadBalancers []ClusterNetworkSpecInstancePoolsLoadBalancers `json:"loadBalancers,omitempty" tf:"load_balancers"`
	// +optional
	PlacementConfigurations []ClusterNetworkSpecInstancePoolsPlacementConfigurations `json:"placementConfigurations,omitempty" tf:"placement_configurations"`
	Size                    *int64                                                   `json:"size" tf:"size"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
}

type ClusterNetworkSpecPlacementConfigurationSecondaryVnicSubnets struct {
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	SubnetID    *string `json:"subnetID" tf:"subnet_id"`
}

type ClusterNetworkSpecPlacementConfiguration struct {
	AvailabilityDomain *string `json:"availabilityDomain" tf:"availability_domain"`
	PrimarySubnetID    *string `json:"primarySubnetID" tf:"primary_subnet_id"`
	// +optional
	SecondaryVnicSubnets []ClusterNetworkSpecPlacementConfigurationSecondaryVnicSubnets `json:"secondaryVnicSubnets,omitempty" tf:"secondary_vnic_subnets"`
}

type ClusterNetworkSpec struct {
	State *ClusterNetworkSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterNetworkSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ClusterNetworkSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags           map[string]string                         `json:"freeformTags,omitempty" tf:"freeform_tags"`
	InstancePools          []ClusterNetworkSpecInstancePools         `json:"instancePools" tf:"instance_pools"`
	PlacementConfiguration *ClusterNetworkSpecPlacementConfiguration `json:"placementConfiguration" tf:"placement_configuration"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type ClusterNetworkStatus struct {
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

// ClusterNetworkList is a list of ClusterNetworks
type ClusterNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterNetwork CRD objects
	Items []ClusterNetwork `json:"items,omitempty"`
}