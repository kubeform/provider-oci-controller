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

type SteeringPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SteeringPolicyAttachmentSpec   `json:"spec,omitempty"`
	Status            SteeringPolicyAttachmentStatus `json:"status,omitempty"`
}

type SteeringPolicyAttachmentSpec struct {
	State *SteeringPolicyAttachmentSpecResource `json:"state,omitempty" tf:"-"`

	Resource SteeringPolicyAttachmentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SteeringPolicyAttachmentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	DomainName  *string `json:"domainName" tf:"domain_name"`
	// +optional
	Rtypes []string `json:"rtypes,omitempty" tf:"rtypes"`
	// +optional
	Self *string `json:"self,omitempty" tf:"self"`
	// +optional
	State            *string `json:"state,omitempty" tf:"state"`
	SteeringPolicyID *string `json:"steeringPolicyID" tf:"steering_policy_id"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	ZoneID      *string `json:"zoneID" tf:"zone_id"`
}

type SteeringPolicyAttachmentStatus struct {
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

// SteeringPolicyAttachmentList is a list of SteeringPolicyAttachments
type SteeringPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SteeringPolicyAttachment CRD objects
	Items []SteeringPolicyAttachment `json:"items,omitempty"`
}
