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

type BuildPipelineStage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BuildPipelineStageSpec   `json:"spec,omitempty"`
	Status            BuildPipelineStageStatus `json:"status,omitempty"`
}

type BuildPipelineStageSpecBuildPipelineStagePredecessorCollectionItems struct {
	ID *string `json:"ID" tf:"id"`
}

type BuildPipelineStageSpecBuildPipelineStagePredecessorCollection struct {
	Items []BuildPipelineStageSpecBuildPipelineStagePredecessorCollectionItems `json:"items" tf:"items"`
}

type BuildPipelineStageSpecBuildSourceCollectionItems struct {
	// +optional
	Branch *string `json:"branch,omitempty" tf:"branch"`
	// +optional
	ConnectionID   *string `json:"connectionID,omitempty" tf:"connection_id"`
	ConnectionType *string `json:"connectionType" tf:"connection_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	RepositoryID *string `json:"repositoryID,omitempty" tf:"repository_id"`
	// +optional
	RepositoryURL *string `json:"repositoryURL,omitempty" tf:"repository_url"`
}

type BuildPipelineStageSpecBuildSourceCollection struct {
	// +optional
	Items []BuildPipelineStageSpecBuildSourceCollectionItems `json:"items,omitempty" tf:"items"`
}

type BuildPipelineStageSpecDeliverArtifactCollectionItems struct {
	// +optional
	ArtifactID *string `json:"artifactID,omitempty" tf:"artifact_id"`
	// +optional
	ArtifactName *string `json:"artifactName,omitempty" tf:"artifact_name"`
}

type BuildPipelineStageSpecDeliverArtifactCollection struct {
	// +optional
	Items []BuildPipelineStageSpecDeliverArtifactCollectionItems `json:"items,omitempty" tf:"items"`
}

type BuildPipelineStageSpecWaitCriteria struct {
	WaitDuration *string `json:"waitDuration" tf:"wait_duration"`
	WaitType     *string `json:"waitType" tf:"wait_type"`
}

type BuildPipelineStageSpec struct {
	State *BuildPipelineStageSpecResource `json:"state,omitempty" tf:"-"`

	Resource BuildPipelineStageSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BuildPipelineStageSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BuildPipelineID                         *string                                                        `json:"buildPipelineID" tf:"build_pipeline_id"`
	BuildPipelineStagePredecessorCollection *BuildPipelineStageSpecBuildPipelineStagePredecessorCollection `json:"buildPipelineStagePredecessorCollection" tf:"build_pipeline_stage_predecessor_collection"`
	BuildPipelineStageType                  *string                                                        `json:"buildPipelineStageType" tf:"build_pipeline_stage_type"`
	// +optional
	BuildSourceCollection *BuildPipelineStageSpecBuildSourceCollection `json:"buildSourceCollection,omitempty" tf:"build_source_collection"`
	// +optional
	BuildSpecFile *string `json:"buildSpecFile,omitempty" tf:"build_spec_file"`
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DeliverArtifactCollection *BuildPipelineStageSpecDeliverArtifactCollection `json:"deliverArtifactCollection,omitempty" tf:"deliver_artifact_collection"`
	// +optional
	DeployPipelineID *string `json:"deployPipelineID,omitempty" tf:"deploy_pipeline_id"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	Image *string `json:"image,omitempty" tf:"image"`
	// +optional
	IsPassAllParametersEnabled *bool `json:"isPassAllParametersEnabled,omitempty" tf:"is_pass_all_parameters_enabled"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	PrimaryBuildSource *string `json:"primaryBuildSource,omitempty" tf:"primary_build_source"`
	// +optional
	ProjectID *string `json:"projectID,omitempty" tf:"project_id"`
	// +optional
	StageExecutionTimeoutInSeconds *int64 `json:"stageExecutionTimeoutInSeconds,omitempty" tf:"stage_execution_timeout_in_seconds"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	WaitCriteria *BuildPipelineStageSpecWaitCriteria `json:"waitCriteria,omitempty" tf:"wait_criteria"`
}

type BuildPipelineStageStatus struct {
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

// BuildPipelineStageList is a list of BuildPipelineStages
type BuildPipelineStageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BuildPipelineStage CRD objects
	Items []BuildPipelineStage `json:"items,omitempty"`
}
