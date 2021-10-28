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

type VirtualCircuit struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualCircuitSpec   `json:"spec,omitempty"`
	Status            VirtualCircuitStatus `json:"status,omitempty"`
}

type VirtualCircuitSpecCrossConnectMappings struct {
	// +optional
	BgpMd5authKey *string `json:"bgpMd5authKey,omitempty" tf:"bgp_md5auth_key"`
	// +optional
	CrossConnectOrCrossConnectGroupID *string `json:"crossConnectOrCrossConnectGroupID,omitempty" tf:"cross_connect_or_cross_connect_group_id"`
	// +optional
	CustomerBGPPeeringIP *string `json:"customerBGPPeeringIP,omitempty" tf:"customer_bgp_peering_ip"`
	// +optional
	CustomerBGPPeeringIpv6 *string `json:"customerBGPPeeringIpv6,omitempty" tf:"customer_bgp_peering_ipv6"`
	// +optional
	OracleBGPPeeringIP *string `json:"oracleBGPPeeringIP,omitempty" tf:"oracle_bgp_peering_ip"`
	// +optional
	OracleBGPPeeringIpv6 *string `json:"oracleBGPPeeringIpv6,omitempty" tf:"oracle_bgp_peering_ipv6"`
	// +optional
	Vlan *int64 `json:"vlan,omitempty" tf:"vlan"`
}

type VirtualCircuitSpecPublicPrefixes struct {
	CidrBlock *string `json:"cidrBlock" tf:"cidr_block"`
}

type VirtualCircuitSpec struct {
	State *VirtualCircuitSpecResource `json:"state,omitempty" tf:"-"`

	Resource VirtualCircuitSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VirtualCircuitSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BandwidthShapeName *string `json:"bandwidthShapeName,omitempty" tf:"bandwidth_shape_name"`
	// +optional
	BgpIpv6sessionState *string `json:"bgpIpv6sessionState,omitempty" tf:"bgp_ipv6session_state"`
	// +optional
	// Deprecated
	BgpManagement *string `json:"bgpManagement,omitempty" tf:"bgp_management"`
	// +optional
	BgpSessionState *string `json:"bgpSessionState,omitempty" tf:"bgp_session_state"`
	CompartmentID   *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	CrossConnectMappings []VirtualCircuitSpecCrossConnectMappings `json:"crossConnectMappings,omitempty" tf:"cross_connect_mappings"`
	// +optional
	CustomerAsn *string `json:"customerAsn,omitempty" tf:"customer_asn"`
	// +optional
	// Deprecated
	CustomerBGPAsn *int64 `json:"customerBGPAsn,omitempty" tf:"customer_bgp_asn"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	GatewayID *string `json:"gatewayID,omitempty" tf:"gateway_id"`
	// +optional
	OracleBGPAsn *int64 `json:"oracleBGPAsn,omitempty" tf:"oracle_bgp_asn"`
	// +optional
	ProviderServiceID *string `json:"providerServiceID,omitempty" tf:"provider_service_id"`
	// +optional
	ProviderServiceKeyName *string `json:"providerServiceKeyName,omitempty" tf:"provider_service_key_name"`
	// +optional
	ProviderState *string `json:"providerState,omitempty" tf:"provider_state"`
	// +optional
	PublicPrefixes []VirtualCircuitSpecPublicPrefixes `json:"publicPrefixes,omitempty" tf:"public_prefixes"`
	// +optional
	ReferenceComment *string `json:"referenceComment,omitempty" tf:"reference_comment"`
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// +optional
	RoutingPolicy []string `json:"routingPolicy,omitempty" tf:"routing_policy"`
	// +optional
	ServiceType *string `json:"serviceType,omitempty" tf:"service_type"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	Type        *string `json:"type" tf:"type"`
}

type VirtualCircuitStatus struct {
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

// VirtualCircuitList is a list of VirtualCircuits
type VirtualCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualCircuit CRD objects
	Items []VirtualCircuit `json:"items,omitempty"`
}
