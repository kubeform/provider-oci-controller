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

type ConnectivityRegistryConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectivityRegistryConnectionSpec   `json:"spec,omitempty"`
	Status            ConnectivityRegistryConnectionStatus `json:"status,omitempty"`
}

type ConnectivityRegistryConnectionSpecConnectionProperties struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ConnectivityRegistryConnectionSpecMetadataAggregator struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ConnectivityRegistryConnectionSpecMetadata struct {
	// +optional
	Aggregator *ConnectivityRegistryConnectionSpecMetadataAggregator `json:"aggregator,omitempty" tf:"aggregator"`
	// +optional
	AggregatorKey *string `json:"aggregatorKey,omitempty" tf:"aggregator_key"`
	// +optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by"`
	// +optional
	CreatedByName *string `json:"createdByName,omitempty" tf:"created_by_name"`
	// +optional
	IdentifierPath *string `json:"identifierPath,omitempty" tf:"identifier_path"`
	// +optional
	InfoFields map[string]string `json:"infoFields,omitempty" tf:"info_fields"`
	// +optional
	IsFavorite *bool `json:"isFavorite,omitempty" tf:"is_favorite"`
	// +optional
	Labels []string `json:"labels,omitempty" tf:"labels"`
	// +optional
	RegistryVersion *int64 `json:"registryVersion,omitempty" tf:"registry_version"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by"`
	// +optional
	UpdatedByName *string `json:"updatedByName,omitempty" tf:"updated_by_name"`
}

type ConnectivityRegistryConnectionSpecPrimarySchemaMetadataAggregator struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ConnectivityRegistryConnectionSpecPrimarySchemaMetadata struct {
	// +optional
	Aggregator *ConnectivityRegistryConnectionSpecPrimarySchemaMetadataAggregator `json:"aggregator,omitempty" tf:"aggregator"`
	// +optional
	AggregatorKey *string `json:"aggregatorKey,omitempty" tf:"aggregator_key"`
	// +optional
	CreatedBy *string `json:"createdBy,omitempty" tf:"created_by"`
	// +optional
	CreatedByName *string `json:"createdByName,omitempty" tf:"created_by_name"`
	// +optional
	IdentifierPath *string `json:"identifierPath,omitempty" tf:"identifier_path"`
	// +optional
	InfoFields map[string]string `json:"infoFields,omitempty" tf:"info_fields"`
	// +optional
	IsFavorite *bool `json:"isFavorite,omitempty" tf:"is_favorite"`
	// +optional
	Labels []string `json:"labels,omitempty" tf:"labels"`
	// +optional
	RegistryVersion *int64 `json:"registryVersion,omitempty" tf:"registry_version"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	UpdatedBy *string `json:"updatedBy,omitempty" tf:"updated_by"`
	// +optional
	UpdatedByName *string `json:"updatedByName,omitempty" tf:"updated_by_name"`
}

type ConnectivityRegistryConnectionSpecPrimarySchemaParentRef struct {
	// +optional
	Parent *string `json:"parent,omitempty" tf:"parent"`
}

type ConnectivityRegistryConnectionSpecPrimarySchema struct {
	// +optional
	DefaultConnection *string `json:"defaultConnection,omitempty" tf:"default_connection"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	ExternalKey *string `json:"externalKey,omitempty" tf:"external_key"`
	Identifier  *string `json:"identifier" tf:"identifier"`
	// +optional
	IsHasContainers *bool   `json:"isHasContainers,omitempty" tf:"is_has_containers"`
	Key             *string `json:"key" tf:"key"`
	// +optional
	Metadata  *ConnectivityRegistryConnectionSpecPrimarySchemaMetadata `json:"metadata,omitempty" tf:"metadata"`
	ModelType *string                                                  `json:"modelType" tf:"model_type"`
	// +optional
	ModelVersion *string `json:"modelVersion,omitempty" tf:"model_version"`
	Name         *string `json:"name" tf:"name"`
	// +optional
	ObjectStatus *int64 `json:"objectStatus,omitempty" tf:"object_status"`
	// +optional
	ObjectVersion *int64 `json:"objectVersion,omitempty" tf:"object_version"`
	// +optional
	ParentRef *ConnectivityRegistryConnectionSpecPrimarySchemaParentRef `json:"parentRef,omitempty" tf:"parent_ref"`
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
}

type ConnectivityRegistryConnectionSpecRegistryMetadata struct {
	// +optional
	AggregatorKey *string `json:"aggregatorKey,omitempty" tf:"aggregator_key"`
	// +optional
	CreatedByUserID *string `json:"createdByUserID,omitempty" tf:"created_by_user_id"`
	// +optional
	CreatedByUserName *string `json:"createdByUserName,omitempty" tf:"created_by_user_name"`
	// +optional
	IsFavorite *bool `json:"isFavorite,omitempty" tf:"is_favorite"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Labels []string `json:"labels,omitempty" tf:"labels"`
	// +optional
	RegistryVersion *int64 `json:"registryVersion,omitempty" tf:"registry_version"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	UpdatedByUserID *string `json:"updatedByUserID,omitempty" tf:"updated_by_user_id"`
	// +optional
	UpdatedByUserName *string `json:"updatedByUserName,omitempty" tf:"updated_by_user_name"`
}

type ConnectivityRegistryConnectionSpec struct {
	State *ConnectivityRegistryConnectionSpecResource `json:"state,omitempty" tf:"-"`

	Resource ConnectivityRegistryConnectionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ConnectivityRegistryConnectionSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ConnectionProperties []ConnectivityRegistryConnectionSpecConnectionProperties `json:"connectionProperties,omitempty" tf:"connection_properties"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Identifier  *string `json:"identifier" tf:"identifier"`
	// +optional
	IsDefault *bool `json:"isDefault,omitempty" tf:"is_default"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Metadata *ConnectivityRegistryConnectionSpecMetadata `json:"metadata,omitempty" tf:"metadata"`
	// +optional
	ModelType *string `json:"modelType,omitempty" tf:"model_type"`
	// +optional
	ModelVersion *string `json:"modelVersion,omitempty" tf:"model_version"`
	Name         *string `json:"name" tf:"name"`
	// +optional
	ObjectStatus *int64 `json:"objectStatus,omitempty" tf:"object_status"`
	// +optional
	ObjectVersion *int64 `json:"objectVersion,omitempty" tf:"object_version"`
	// +optional
	PrimarySchema *ConnectivityRegistryConnectionSpecPrimarySchema `json:"primarySchema,omitempty" tf:"primary_schema"`
	Properties    map[string]string                                `json:"properties" tf:"properties"`
	RegistryID    *string                                          `json:"registryID" tf:"registry_id"`
	// +optional
	RegistryMetadata *ConnectivityRegistryConnectionSpecRegistryMetadata `json:"registryMetadata,omitempty" tf:"registry_metadata"`
	Type             *string                                             `json:"type" tf:"type"`
}

type ConnectivityRegistryConnectionStatus struct {
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

// ConnectivityRegistryConnectionList is a list of ConnectivityRegistryConnections
type ConnectivityRegistryConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConnectivityRegistryConnection CRD objects
	Items []ConnectivityRegistryConnection `json:"items,omitempty"`
}
