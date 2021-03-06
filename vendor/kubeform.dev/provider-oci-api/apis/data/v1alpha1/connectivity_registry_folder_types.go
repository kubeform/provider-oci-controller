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

type ConnectivityRegistryFolder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectivityRegistryFolderSpec   `json:"spec,omitempty"`
	Status            ConnectivityRegistryFolderStatus `json:"status,omitempty"`
}

type ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionConnectionProperties struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionMetadataAggregator struct {
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

type ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionMetadata struct {
	// +optional
	Aggregator *ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionMetadataAggregator `json:"aggregator,omitempty" tf:"aggregator"`
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

type ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionPrimarySchemaMetadataAggregator struct {
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

type ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionPrimarySchemaMetadata struct {
	// +optional
	Aggregator *ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionPrimarySchemaMetadataAggregator `json:"aggregator,omitempty" tf:"aggregator"`
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

type ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionPrimarySchemaParentRef struct {
	// +optional
	Parent *string `json:"parent,omitempty" tf:"parent"`
}

type ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionPrimarySchema struct {
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
	Metadata  *ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionPrimarySchemaMetadata `json:"metadata,omitempty" tf:"metadata"`
	ModelType *string                                                                         `json:"modelType" tf:"model_type"`
	// +optional
	ModelVersion *string `json:"modelVersion,omitempty" tf:"model_version"`
	Name         *string `json:"name" tf:"name"`
	// +optional
	ObjectStatus *int64 `json:"objectStatus,omitempty" tf:"object_status"`
	// +optional
	ObjectVersion *int64 `json:"objectVersion,omitempty" tf:"object_version"`
	// +optional
	ParentRef *ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionPrimarySchemaParentRef `json:"parentRef,omitempty" tf:"parent_ref"`
	// +optional
	ResourceName *string `json:"resourceName,omitempty" tf:"resource_name"`
}

type ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionRegistryMetadata struct {
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

type ConnectivityRegistryFolderSpecDataAssetsDefaultConnection struct {
	// +optional
	ConnectionProperties []ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionConnectionProperties `json:"connectionProperties,omitempty" tf:"connection_properties"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Identifier  *string `json:"identifier" tf:"identifier"`
	// +optional
	IsDefault *bool   `json:"isDefault,omitempty" tf:"is_default"`
	Key       *string `json:"key" tf:"key"`
	// +optional
	Metadata *ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionMetadata `json:"metadata,omitempty" tf:"metadata"`
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
	PrimarySchema *ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionPrimarySchema `json:"primarySchema,omitempty" tf:"primary_schema"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
	// +optional
	RegistryMetadata *ConnectivityRegistryFolderSpecDataAssetsDefaultConnectionRegistryMetadata `json:"registryMetadata,omitempty" tf:"registry_metadata"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ConnectivityRegistryFolderSpecDataAssetsMetadataAggregator struct {
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

type ConnectivityRegistryFolderSpecDataAssetsMetadata struct {
	// +optional
	Aggregator *ConnectivityRegistryFolderSpecDataAssetsMetadataAggregator `json:"aggregator,omitempty" tf:"aggregator"`
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

type ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemParentRef struct {
	// +optional
	Parent *string `json:"parent,omitempty" tf:"parent"`
}

type ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemTypesConfigDefinitionParentRef struct {
	// +optional
	Parent *string `json:"parent,omitempty" tf:"parent"`
}

type ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemTypesConfigDefinition struct {
	// +optional
	ConfigParameterDefinitions map[string]string `json:"configParameterDefinitions,omitempty" tf:"config_parameter_definitions"`
	// +optional
	IsContained *bool `json:"isContained,omitempty" tf:"is_contained"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	ModelType *string `json:"modelType,omitempty" tf:"model_type"`
	// +optional
	ModelVersion *string `json:"modelVersion,omitempty" tf:"model_version"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	ObjectStatus *int64 `json:"objectStatus,omitempty" tf:"object_status"`
	// +optional
	ParentRef *ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemTypesConfigDefinitionParentRef `json:"parentRef,omitempty" tf:"parent_ref"`
}

type ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemTypesParentRef struct {
	// +optional
	Parent *string `json:"parent,omitempty" tf:"parent"`
}

type ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemTypes struct {
	// +optional
	ConfigDefinition *ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemTypesConfigDefinition `json:"configDefinition,omitempty" tf:"config_definition"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DtType *string `json:"dtType,omitempty" tf:"dt_type"`
	// +optional
	Key       *string `json:"key,omitempty" tf:"key"`
	ModelType *string `json:"modelType" tf:"model_type"`
	// +optional
	ModelVersion *string `json:"modelVersion,omitempty" tf:"model_version"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	ObjectStatus *int64 `json:"objectStatus,omitempty" tf:"object_status"`
	// +optional
	ParentRef *ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemTypesParentRef `json:"parentRef,omitempty" tf:"parent_ref"`
	// +optional
	TypeSystemName *string `json:"typeSystemName,omitempty" tf:"type_system_name"`
}

type ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystem struct {
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	ModelType *string `json:"modelType,omitempty" tf:"model_type"`
	// +optional
	ModelVersion *string `json:"modelVersion,omitempty" tf:"model_version"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	ObjectStatus *int64 `json:"objectStatus,omitempty" tf:"object_status"`
	// +optional
	ObjectVersion *int64 `json:"objectVersion,omitempty" tf:"object_version"`
	// +optional
	ParentRef *ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemParentRef `json:"parentRef,omitempty" tf:"parent_ref"`
	// +optional
	TypeMappingFrom map[string]string `json:"typeMappingFrom,omitempty" tf:"type_mapping_from"`
	// +optional
	TypeMappingTo map[string]string `json:"typeMappingTo,omitempty" tf:"type_mapping_to"`
	// +optional
	Types []ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystemTypes `json:"types,omitempty" tf:"types"`
}

type ConnectivityRegistryFolderSpecDataAssetsRegistryMetadata struct {
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

type ConnectivityRegistryFolderSpecDataAssets struct {
	// +optional
	AssetProperties map[string]string `json:"assetProperties,omitempty" tf:"asset_properties"`
	// +optional
	DefaultConnection *ConnectivityRegistryFolderSpecDataAssetsDefaultConnection `json:"defaultConnection,omitempty" tf:"default_connection"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	ExternalKey *string `json:"externalKey,omitempty" tf:"external_key"`
	Identifier  *string `json:"identifier" tf:"identifier"`
	Key         *string `json:"key" tf:"key"`
	// +optional
	Metadata *ConnectivityRegistryFolderSpecDataAssetsMetadata `json:"metadata,omitempty" tf:"metadata"`
	// +optional
	ModelType *string `json:"modelType,omitempty" tf:"model_type"`
	// +optional
	ModelVersion *string `json:"modelVersion,omitempty" tf:"model_version"`
	Name         *string `json:"name" tf:"name"`
	// +optional
	NativeTypeSystem *ConnectivityRegistryFolderSpecDataAssetsNativeTypeSystem `json:"nativeTypeSystem,omitempty" tf:"native_type_system"`
	// +optional
	ObjectStatus *int64 `json:"objectStatus,omitempty" tf:"object_status"`
	// +optional
	ObjectVersion *int64 `json:"objectVersion,omitempty" tf:"object_version"`
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
	// +optional
	RegistryMetadata *ConnectivityRegistryFolderSpecDataAssetsRegistryMetadata `json:"registryMetadata,omitempty" tf:"registry_metadata"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type ConnectivityRegistryFolderSpecParentRef struct {
	// +optional
	Parent *string `json:"parent,omitempty" tf:"parent"`
}

type ConnectivityRegistryFolderSpec struct {
	State *ConnectivityRegistryFolderSpecResource `json:"state,omitempty" tf:"-"`

	Resource ConnectivityRegistryFolderSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ConnectivityRegistryFolderSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	DataAssets []ConnectivityRegistryFolderSpecDataAssets `json:"dataAssets,omitempty" tf:"data_assets"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	Identifier  *string `json:"identifier" tf:"identifier"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
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
	ParentRef  *ConnectivityRegistryFolderSpecParentRef `json:"parentRef,omitempty" tf:"parent_ref"`
	RegistryID *string                                  `json:"registryID" tf:"registry_id"`
}

type ConnectivityRegistryFolderStatus struct {
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

// ConnectivityRegistryFolderList is a list of ConnectivityRegistryFolders
type ConnectivityRegistryFolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConnectivityRegistryFolder CRD objects
	Items []ConnectivityRegistryFolder `json:"items,omitempty"`
}
