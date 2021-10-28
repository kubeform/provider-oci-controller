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

type GateDatabaseRegistration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GateDatabaseRegistrationSpec   `json:"spec,omitempty"`
	Status            GateDatabaseRegistrationStatus `json:"status,omitempty"`
}

type GateDatabaseRegistrationSpec struct {
	State *GateDatabaseRegistrationSpecResource `json:"state,omitempty" tf:"-"`

	Resource GateDatabaseRegistrationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GateDatabaseRegistrationSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AliasName     *string `json:"aliasName" tf:"alias_name"`
	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	ConnectionString *string `json:"connectionString,omitempty" tf:"connection_string"`
	// +optional
	DatabaseID *string `json:"databaseID,omitempty" tf:"database_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	DisplayName *string `json:"displayName" tf:"display_name"`
	Fqdn        *string `json:"fqdn" tf:"fqdn"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	KeyID *string `json:"keyID,omitempty" tf:"key_id"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	Password         *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	RcePrivateIP *string `json:"rcePrivateIP,omitempty" tf:"rce_private_ip"`
	// +optional
	SecretCompartmentID *string `json:"secretCompartmentID,omitempty" tf:"secret_compartment_id"`
	// +optional
	SecretID *string `json:"secretID,omitempty" tf:"secret_id"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	Username    *string `json:"username" tf:"username"`
	// +optional
	VaultID *string `json:"vaultID,omitempty" tf:"vault_id"`
	// +optional
	Wallet *string `json:"wallet,omitempty" tf:"wallet"`
}

type GateDatabaseRegistrationStatus struct {
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

// GateDatabaseRegistrationList is a list of GateDatabaseRegistrations
type GateDatabaseRegistrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GateDatabaseRegistration CRD objects
	Items []GateDatabaseRegistration `json:"items,omitempty"`
}
