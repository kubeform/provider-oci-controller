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

type ProxySubscriptionRedeemableUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxySubscriptionRedeemableUserSpec   `json:"spec,omitempty"`
	Status            ProxySubscriptionRedeemableUserStatus `json:"status,omitempty"`
}

type ProxySubscriptionRedeemableUserSpecItems struct {
	EmailID *string `json:"emailID" tf:"email_id"`
}

type ProxySubscriptionRedeemableUserSpec struct {
	State *ProxySubscriptionRedeemableUserSpecResource `json:"state,omitempty" tf:"-"`

	Resource ProxySubscriptionRedeemableUserSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ProxySubscriptionRedeemableUserSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Items          []ProxySubscriptionRedeemableUserSpecItems `json:"items" tf:"items"`
	SubscriptionID *string                                    `json:"subscriptionID" tf:"subscription_id"`
	TenancyID      *string                                    `json:"tenancyID" tf:"tenancy_id"`
	// +optional
	UserID *string `json:"userID,omitempty" tf:"user_id"`
}

type ProxySubscriptionRedeemableUserStatus struct {
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

// ProxySubscriptionRedeemableUserList is a list of ProxySubscriptionRedeemableUsers
type ProxySubscriptionRedeemableUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProxySubscriptionRedeemableUser CRD objects
	Items []ProxySubscriptionRedeemableUser `json:"items,omitempty"`
}
