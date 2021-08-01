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

type LoadBalancerNetworkLoadBalancer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerNetworkLoadBalancerSpec   `json:"spec,omitempty"`
	Status            LoadBalancerNetworkLoadBalancerStatus `json:"status,omitempty"`
}

type LoadBalancerNetworkLoadBalancerSpecIpAddressesReservedIP struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
}

type LoadBalancerNetworkLoadBalancerSpecIpAddresses struct {
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public"`
	// +optional
	ReservedIP *LoadBalancerNetworkLoadBalancerSpecIpAddressesReservedIP `json:"reservedIP,omitempty" tf:"reserved_ip"`
}

type LoadBalancerNetworkLoadBalancerSpecReservedIPS struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
}

type LoadBalancerNetworkLoadBalancerSpec struct {
	State *LoadBalancerNetworkLoadBalancerSpecResource `json:"state,omitempty" tf:"-"`

	Resource LoadBalancerNetworkLoadBalancerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LoadBalancerNetworkLoadBalancerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	DisplayName *string           `json:"displayName" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	IpAddresses []LoadBalancerNetworkLoadBalancerSpecIpAddresses `json:"ipAddresses,omitempty" tf:"ip_addresses"`
	// +optional
	IsPreserveSourceDestination *bool `json:"isPreserveSourceDestination,omitempty" tf:"is_preserve_source_destination"`
	// +optional
	IsPrivate *bool `json:"isPrivate,omitempty" tf:"is_private"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	NetworkSecurityGroupIDS []string `json:"networkSecurityGroupIDS,omitempty" tf:"network_security_group_ids"`
	// +optional
	ReservedIPS []LoadBalancerNetworkLoadBalancerSpecReservedIPS `json:"reservedIPS,omitempty" tf:"reserved_ips"`
	// +optional
	State    *string `json:"state,omitempty" tf:"state"`
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type LoadBalancerNetworkLoadBalancerStatus struct {
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

// LoadBalancerNetworkLoadBalancerList is a list of LoadBalancerNetworkLoadBalancers
type LoadBalancerNetworkLoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoadBalancerNetworkLoadBalancer CRD objects
	Items []LoadBalancerNetworkLoadBalancer `json:"items,omitempty"`
}
