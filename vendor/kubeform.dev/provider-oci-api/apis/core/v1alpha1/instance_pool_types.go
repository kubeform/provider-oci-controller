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

type InstancePool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstancePoolSpec   `json:"spec,omitempty"`
	Status            InstancePoolStatus `json:"status,omitempty"`
}

type InstancePoolSpecLoadBalancers struct {
	BackendSetName *string `json:"backendSetName" tf:"backend_set_name"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	InstancePoolID *string `json:"instancePoolID,omitempty" tf:"instance_pool_id"`
	LoadBalancerID *string `json:"loadBalancerID" tf:"load_balancer_id"`
	Port           *int64  `json:"port" tf:"port"`
	// +optional
	State         *string `json:"state,omitempty" tf:"state"`
	VnicSelection *string `json:"vnicSelection" tf:"vnic_selection"`
}

type InstancePoolSpecPlacementConfigurationsSecondaryVnicSubnets struct {
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	SubnetID    *string `json:"subnetID" tf:"subnet_id"`
}

type InstancePoolSpecPlacementConfigurations struct {
	AvailabilityDomain *string `json:"availabilityDomain" tf:"availability_domain"`
	// +optional
	FaultDomains    []string `json:"faultDomains,omitempty" tf:"fault_domains"`
	PrimarySubnetID *string  `json:"primarySubnetID" tf:"primary_subnet_id"`
	// +optional
	SecondaryVnicSubnets []InstancePoolSpecPlacementConfigurationsSecondaryVnicSubnets `json:"secondaryVnicSubnets,omitempty" tf:"secondary_vnic_subnets"`
}

type InstancePoolSpec struct {
	State *InstancePoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource InstancePoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InstancePoolSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActualSize    *int64  `json:"actualSize,omitempty" tf:"actual_size"`
	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags            map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	InstanceConfigurationID *string           `json:"instanceConfigurationID" tf:"instance_configuration_id"`
	// +optional
	LoadBalancers           []InstancePoolSpecLoadBalancers           `json:"loadBalancers,omitempty" tf:"load_balancers"`
	PlacementConfigurations []InstancePoolSpecPlacementConfigurations `json:"placementConfigurations" tf:"placement_configurations"`
	Size                    *int64                                    `json:"size" tf:"size"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
}

type InstancePoolStatus struct {
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

// InstancePoolList is a list of InstancePools
type InstancePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of InstancePool CRD objects
	Items []InstancePool `json:"items,omitempty"`
}