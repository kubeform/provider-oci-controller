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

type ManagedInstanceManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedInstanceManagementSpec   `json:"spec,omitempty"`
	Status            ManagedInstanceManagementStatus `json:"status,omitempty"`
}

type ManagedInstanceManagementSpecChildSoftwareSources struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type ManagedInstanceManagementSpecManagedInstanceGroups struct {
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
}

type ManagedInstanceManagementSpecParentSoftwareSource struct {
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type ManagedInstanceManagementSpec struct {
	State *ManagedInstanceManagementSpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagedInstanceManagementSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagedInstanceManagementSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ChildSoftwareSources []ManagedInstanceManagementSpecChildSoftwareSources `json:"childSoftwareSources,omitempty" tf:"child_software_sources"`
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	LastBoot *string `json:"lastBoot,omitempty" tf:"last_boot"`
	// +optional
	LastCheckin *string `json:"lastCheckin,omitempty" tf:"last_checkin"`
	// +optional
	ManagedInstanceGroups []ManagedInstanceManagementSpecManagedInstanceGroups `json:"managedInstanceGroups,omitempty" tf:"managed_instance_groups"`
	ManagedInstanceID     *string                                              `json:"managedInstanceID" tf:"managed_instance_id"`
	// +optional
	OsKernelVersion *string `json:"osKernelVersion,omitempty" tf:"os_kernel_version"`
	// +optional
	OsName *string `json:"osName,omitempty" tf:"os_name"`
	// +optional
	OsVersion *string `json:"osVersion,omitempty" tf:"os_version"`
	// +optional
	ParentSoftwareSource *ManagedInstanceManagementSpecParentSoftwareSource `json:"parentSoftwareSource,omitempty" tf:"parent_software_source"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	UpdatesAvailable *int64 `json:"updatesAvailable,omitempty" tf:"updates_available"`
}

type ManagedInstanceManagementStatus struct {
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

// ManagedInstanceManagementList is a list of ManagedInstanceManagements
type ManagedInstanceManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedInstanceManagement CRD objects
	Items []ManagedInstanceManagement `json:"items,omitempty"`
}
