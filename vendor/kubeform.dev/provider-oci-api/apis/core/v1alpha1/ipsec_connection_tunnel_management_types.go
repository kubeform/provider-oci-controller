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

type IpsecConnectionTunnelManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpsecConnectionTunnelManagementSpec   `json:"spec,omitempty"`
	Status            IpsecConnectionTunnelManagementStatus `json:"status,omitempty"`
}

type IpsecConnectionTunnelManagementSpecBgpSessionInfo struct {
	// +optional
	BgpState *string `json:"bgpState,omitempty" tf:"bgp_state"`
	// +optional
	CustomerBGPAsn *string `json:"customerBGPAsn,omitempty" tf:"customer_bgp_asn"`
	// +optional
	CustomerInterfaceIP *string `json:"customerInterfaceIP,omitempty" tf:"customer_interface_ip"`
	// +optional
	OracleBGPAsn *string `json:"oracleBGPAsn,omitempty" tf:"oracle_bgp_asn"`
	// +optional
	OracleInterfaceIP *string `json:"oracleInterfaceIP,omitempty" tf:"oracle_interface_ip"`
}

type IpsecConnectionTunnelManagementSpecEncryptionDomainConfig struct {
	// +optional
	CpeTrafficSelector []string `json:"cpeTrafficSelector,omitempty" tf:"cpe_traffic_selector"`
	// +optional
	OracleTrafficSelector []string `json:"oracleTrafficSelector,omitempty" tf:"oracle_traffic_selector"`
}

type IpsecConnectionTunnelManagementSpec struct {
	State *IpsecConnectionTunnelManagementSpecResource `json:"state,omitempty" tf:"-"`

	Resource IpsecConnectionTunnelManagementSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type IpsecConnectionTunnelManagementSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BgpSessionInfo []IpsecConnectionTunnelManagementSpecBgpSessionInfo `json:"bgpSessionInfo,omitempty" tf:"bgp_session_info"`
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	CpeIP *string `json:"cpeIP,omitempty" tf:"cpe_ip"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	EncryptionDomainConfig *IpsecConnectionTunnelManagementSpecEncryptionDomainConfig `json:"encryptionDomainConfig,omitempty" tf:"encryption_domain_config"`
	// +optional
	IkeVersion *string `json:"ikeVersion,omitempty" tf:"ike_version"`
	IpsecID    *string `json:"ipsecID" tf:"ipsec_id"`
	Routing    *string `json:"routing" tf:"routing"`
	// +optional
	SharedSecret *string `json:"sharedSecret,omitempty" tf:"shared_secret"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeStatusUpdated *string `json:"timeStatusUpdated,omitempty" tf:"time_status_updated"`
	TunnelID          *string `json:"tunnelID" tf:"tunnel_id"`
	// +optional
	VpnIP *string `json:"vpnIP,omitempty" tf:"vpn_ip"`
}

type IpsecConnectionTunnelManagementStatus struct {
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

// IpsecConnectionTunnelManagementList is a list of IpsecConnectionTunnelManagements
type IpsecConnectionTunnelManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IpsecConnectionTunnelManagement CRD objects
	Items []IpsecConnectionTunnelManagement `json:"items,omitempty"`
}
