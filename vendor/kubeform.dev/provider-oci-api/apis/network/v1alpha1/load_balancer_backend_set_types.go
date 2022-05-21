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

type LoadBalancerBackendSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerBackendSetSpec   `json:"spec,omitempty"`
	Status            LoadBalancerBackendSetStatus `json:"status,omitempty"`
}

type LoadBalancerBackendSetSpecBackends struct {
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	IsBackup *bool `json:"isBackup,omitempty" tf:"is_backup"`
	// +optional
	IsDrain *bool `json:"isDrain,omitempty" tf:"is_drain"`
	// +optional
	IsOffline *bool `json:"isOffline,omitempty" tf:"is_offline"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	Port *int64  `json:"port" tf:"port"`
	// +optional
	TargetID *string `json:"targetID,omitempty" tf:"target_id"`
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type LoadBalancerBackendSetSpecHealthChecker struct {
	// +optional
	IntervalInMillis *int64 `json:"intervalInMillis,omitempty" tf:"interval_in_millis"`
	// +optional
	Port     *int64  `json:"port,omitempty" tf:"port"`
	Protocol *string `json:"protocol" tf:"protocol"`
	// +optional
	RequestData *string `json:"requestData,omitempty" tf:"request_data"`
	// +optional
	ResponseBodyRegex *string `json:"responseBodyRegex,omitempty" tf:"response_body_regex"`
	// +optional
	ResponseData *string `json:"responseData,omitempty" tf:"response_data"`
	// +optional
	Retries *int64 `json:"retries,omitempty" tf:"retries"`
	// +optional
	ReturnCode *int64 `json:"returnCode,omitempty" tf:"return_code"`
	// +optional
	TimeoutInMillis *int64 `json:"timeoutInMillis,omitempty" tf:"timeout_in_millis"`
	// +optional
	UrlPath *string `json:"urlPath,omitempty" tf:"url_path"`
}

type LoadBalancerBackendSetSpec struct {
	State *LoadBalancerBackendSetSpecResource `json:"state,omitempty" tf:"-"`

	Resource LoadBalancerBackendSetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type LoadBalancerBackendSetSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Backends      []LoadBalancerBackendSetSpecBackends     `json:"backends,omitempty" tf:"backends"`
	HealthChecker *LoadBalancerBackendSetSpecHealthChecker `json:"healthChecker" tf:"health_checker"`
	// +optional
	IpVersion *string `json:"ipVersion,omitempty" tf:"ip_version"`
	// +optional
	IsPreserveSource      *bool   `json:"isPreserveSource,omitempty" tf:"is_preserve_source"`
	Name                  *string `json:"name" tf:"name"`
	NetworkLoadBalancerID *string `json:"networkLoadBalancerID" tf:"network_load_balancer_id"`
	Policy                *string `json:"policy" tf:"policy"`
}

type LoadBalancerBackendSetStatus struct {
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

// LoadBalancerBackendSetList is a list of LoadBalancerBackendSets
type LoadBalancerBackendSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoadBalancerBackendSet CRD objects
	Items []LoadBalancerBackendSet `json:"items,omitempty"`
}