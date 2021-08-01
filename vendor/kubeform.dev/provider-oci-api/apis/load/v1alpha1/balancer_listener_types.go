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

type BalancerListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BalancerListenerSpec   `json:"spec,omitempty"`
	Status            BalancerListenerStatus `json:"status,omitempty"`
}

type BalancerListenerSpecConnectionConfiguration struct {
	// +optional
	BackendTcpProxyProtocolVersion *int64  `json:"backendTcpProxyProtocolVersion,omitempty" tf:"backend_tcp_proxy_protocol_version"`
	IdleTimeoutInSeconds           *string `json:"idleTimeoutInSeconds" tf:"idle_timeout_in_seconds"`
}

type BalancerListenerSpecSslConfiguration struct {
	CertificateName *string `json:"certificateName" tf:"certificate_name"`
	// +optional
	CipherSuiteName *string `json:"cipherSuiteName,omitempty" tf:"cipher_suite_name"`
	// +optional
	Protocols []string `json:"protocols,omitempty" tf:"protocols"`
	// +optional
	ServerOrderPreference *string `json:"serverOrderPreference,omitempty" tf:"server_order_preference"`
	// +optional
	VerifyDepth *int64 `json:"verifyDepth,omitempty" tf:"verify_depth"`
	// +optional
	VerifyPeerCertificate *bool `json:"verifyPeerCertificate,omitempty" tf:"verify_peer_certificate"`
}

type BalancerListenerSpec struct {
	State *BalancerListenerSpecResource `json:"state,omitempty" tf:"-"`

	Resource BalancerListenerSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BalancerListenerSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ConnectionConfiguration *BalancerListenerSpecConnectionConfiguration `json:"connectionConfiguration,omitempty" tf:"connection_configuration"`
	DefaultBackendSetName   *string                                      `json:"defaultBackendSetName" tf:"default_backend_set_name"`
	// +optional
	HostnameNames  []string `json:"hostnameNames,omitempty" tf:"hostname_names"`
	LoadBalancerID *string  `json:"loadBalancerID" tf:"load_balancer_id"`
	Name           *string  `json:"name" tf:"name"`
	// +optional
	PathRouteSetName *string `json:"pathRouteSetName,omitempty" tf:"path_route_set_name"`
	Port             *int64  `json:"port" tf:"port"`
	Protocol         *string `json:"protocol" tf:"protocol"`
	// +optional
	RoutingPolicyName *string `json:"routingPolicyName,omitempty" tf:"routing_policy_name"`
	// +optional
	RuleSetNames []string `json:"ruleSetNames,omitempty" tf:"rule_set_names"`
	// +optional
	SslConfiguration *BalancerListenerSpecSslConfiguration `json:"sslConfiguration,omitempty" tf:"ssl_configuration"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type BalancerListenerStatus struct {
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

// BalancerListenerList is a list of BalancerListeners
type BalancerListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BalancerListener CRD objects
	Items []BalancerListener `json:"items,omitempty"`
}
