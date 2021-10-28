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

type SafeTargetDatabase struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SafeTargetDatabaseSpec   `json:"spec,omitempty"`
	Status            SafeTargetDatabaseStatus `json:"status,omitempty"`
}

type SafeTargetDatabaseSpecConnectionOption struct {
	ConnectionType *string `json:"connectionType" tf:"connection_type"`
	// +optional
	DatasafePrivateEndpointID *string `json:"datasafePrivateEndpointID,omitempty" tf:"datasafe_private_endpoint_id"`
	// +optional
	OnPremConnectorID *string `json:"onPremConnectorID,omitempty" tf:"on_prem_connector_id"`
}

type SafeTargetDatabaseSpecCredentials struct {
	Password *string `json:"-" sensitive:"true" tf:"password"`
	UserName *string `json:"userName" tf:"user_name"`
}

type SafeTargetDatabaseSpecDatabaseDetails struct {
	// +optional
	AutonomousDatabaseID *string `json:"autonomousDatabaseID,omitempty" tf:"autonomous_database_id"`
	DatabaseType         *string `json:"databaseType" tf:"database_type"`
	// +optional
	DbSystemID         *string `json:"dbSystemID,omitempty" tf:"db_system_id"`
	InfrastructureType *string `json:"infrastructureType" tf:"infrastructure_type"`
	// +optional
	InstanceID *string `json:"instanceID,omitempty" tf:"instance_id"`
	// +optional
	IpAddresses []string `json:"ipAddresses,omitempty" tf:"ip_addresses"`
	// +optional
	ListenerPort *int64 `json:"listenerPort,omitempty" tf:"listener_port"`
	// +optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
	// +optional
	VmClusterID *string `json:"vmClusterID,omitempty" tf:"vm_cluster_id"`
}

type SafeTargetDatabaseSpecTlsConfig struct {
	// +optional
	CertificateStoreType *string `json:"certificateStoreType,omitempty" tf:"certificate_store_type"`
	// +optional
	KeyStoreContent *string `json:"keyStoreContent,omitempty" tf:"key_store_content"`
	Status          *string `json:"status" tf:"status"`
	// +optional
	StorePassword *string `json:"-" sensitive:"true" tf:"store_password"`
	// +optional
	TrustStoreContent *string `json:"trustStoreContent,omitempty" tf:"trust_store_content"`
}

type SafeTargetDatabaseSpec struct {
	State *SafeTargetDatabaseSpecResource `json:"state,omitempty" tf:"-"`

	Resource SafeTargetDatabaseSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SafeTargetDatabaseSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	ConnectionOption *SafeTargetDatabaseSpecConnectionOption `json:"connectionOption,omitempty" tf:"connection_option"`
	// +optional
	Credentials     *SafeTargetDatabaseSpecCredentials     `json:"credentials,omitempty" tf:"credentials"`
	DatabaseDetails *SafeTargetDatabaseSpecDatabaseDetails `json:"databaseDetails" tf:"database_details"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	TlsConfig *SafeTargetDatabaseSpecTlsConfig `json:"tlsConfig,omitempty" tf:"tls_config"`
}

type SafeTargetDatabaseStatus struct {
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

// SafeTargetDatabaseList is a list of SafeTargetDatabases
type SafeTargetDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SafeTargetDatabase CRD objects
	Items []SafeTargetDatabase `json:"items,omitempty"`
}
