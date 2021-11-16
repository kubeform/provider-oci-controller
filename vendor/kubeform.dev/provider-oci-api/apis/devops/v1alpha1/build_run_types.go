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

type BuildRun struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BuildRunSpec   `json:"spec,omitempty"`
	Status            BuildRunStatus `json:"status,omitempty"`
}

type BuildRunSpecBuildOutputsArtifactOverrideParametersItems struct {
	// +optional
	DeployArtifactID *string `json:"deployArtifactID,omitempty" tf:"deploy_artifact_id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type BuildRunSpecBuildOutputsArtifactOverrideParameters struct {
	// +optional
	Items []BuildRunSpecBuildOutputsArtifactOverrideParametersItems `json:"items,omitempty" tf:"items"`
}

type BuildRunSpecBuildOutputsDeliveredArtifactsItems struct {
	// +optional
	ArtifactRepositoryID *string `json:"artifactRepositoryID,omitempty" tf:"artifact_repository_id"`
	// +optional
	ArtifactType *string `json:"artifactType,omitempty" tf:"artifact_type"`
	// +optional
	DeliveredArtifactHash *string `json:"deliveredArtifactHash,omitempty" tf:"delivered_artifact_hash"`
	// +optional
	DeliveredArtifactID *string `json:"deliveredArtifactID,omitempty" tf:"delivered_artifact_id"`
	// +optional
	DeployArtifactID *string `json:"deployArtifactID,omitempty" tf:"deploy_artifact_id"`
	// +optional
	ImageURI *string `json:"imageURI,omitempty" tf:"image_uri"`
	// +optional
	OutputArtifactName *string `json:"outputArtifactName,omitempty" tf:"output_artifact_name"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
}

type BuildRunSpecBuildOutputsDeliveredArtifacts struct {
	// +optional
	Items []BuildRunSpecBuildOutputsDeliveredArtifactsItems `json:"items,omitempty" tf:"items"`
}

type BuildRunSpecBuildOutputsExportedVariablesItems struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type BuildRunSpecBuildOutputsExportedVariables struct {
	// +optional
	Items []BuildRunSpecBuildOutputsExportedVariablesItems `json:"items,omitempty" tf:"items"`
}

type BuildRunSpecBuildOutputs struct {
	// +optional
	ArtifactOverrideParameters *BuildRunSpecBuildOutputsArtifactOverrideParameters `json:"artifactOverrideParameters,omitempty" tf:"artifact_override_parameters"`
	// +optional
	DeliveredArtifacts *BuildRunSpecBuildOutputsDeliveredArtifacts `json:"deliveredArtifacts,omitempty" tf:"delivered_artifacts"`
	// +optional
	ExportedVariables *BuildRunSpecBuildOutputsExportedVariables `json:"exportedVariables,omitempty" tf:"exported_variables"`
}

type BuildRunSpecBuildRunArgumentsItems struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type BuildRunSpecBuildRunArguments struct {
	Items []BuildRunSpecBuildRunArgumentsItems `json:"items" tf:"items"`
}

type BuildRunSpecBuildRunProgress struct {
	// +optional
	BuildPipelineStageRunProgress map[string]string `json:"buildPipelineStageRunProgress,omitempty" tf:"build_pipeline_stage_run_progress"`
	// +optional
	TimeFinished *string `json:"timeFinished,omitempty" tf:"time_finished"`
	// +optional
	TimeStarted *string `json:"timeStarted,omitempty" tf:"time_started"`
}

type BuildRunSpecBuildRunSourceTriggerInfoActionsFilterInclude struct {
	// +optional
	BaseRef *string `json:"baseRef,omitempty" tf:"base_ref"`
	// +optional
	HeadRef *string `json:"headRef,omitempty" tf:"head_ref"`
}

type BuildRunSpecBuildRunSourceTriggerInfoActionsFilter struct {
	// +optional
	Events []string `json:"events,omitempty" tf:"events"`
	// +optional
	Include *BuildRunSpecBuildRunSourceTriggerInfoActionsFilterInclude `json:"include,omitempty" tf:"include"`
	// +optional
	TriggerSource *string `json:"triggerSource,omitempty" tf:"trigger_source"`
}

type BuildRunSpecBuildRunSourceTriggerInfoActions struct {
	// +optional
	BuildPipelineID *string `json:"buildPipelineID,omitempty" tf:"build_pipeline_id"`
	// +optional
	Filter *BuildRunSpecBuildRunSourceTriggerInfoActionsFilter `json:"filter,omitempty" tf:"filter"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type BuildRunSpecBuildRunSourceTriggerInfo struct {
	// +optional
	Actions []BuildRunSpecBuildRunSourceTriggerInfoActions `json:"actions,omitempty" tf:"actions"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
}

type BuildRunSpecBuildRunSource struct {
	// +optional
	RepositoryID *string `json:"repositoryID,omitempty" tf:"repository_id"`
	// +optional
	SourceType *string `json:"sourceType,omitempty" tf:"source_type"`
	// +optional
	TriggerID *string `json:"triggerID,omitempty" tf:"trigger_id"`
	// +optional
	TriggerInfo *BuildRunSpecBuildRunSourceTriggerInfo `json:"triggerInfo,omitempty" tf:"trigger_info"`
}

type BuildRunSpecCommitInfo struct {
	CommitHash       *string `json:"commitHash" tf:"commit_hash"`
	RepositoryBranch *string `json:"repositoryBranch" tf:"repository_branch"`
	RepositoryURL    *string `json:"repositoryURL" tf:"repository_url"`
}

type BuildRunSpec struct {
	State *BuildRunSpecResource `json:"state,omitempty" tf:"-"`

	Resource BuildRunSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type BuildRunSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BuildOutputs    *BuildRunSpecBuildOutputs `json:"buildOutputs,omitempty" tf:"build_outputs"`
	BuildPipelineID *string                   `json:"buildPipelineID" tf:"build_pipeline_id"`
	// +optional
	BuildRunArguments *BuildRunSpecBuildRunArguments `json:"buildRunArguments,omitempty" tf:"build_run_arguments"`
	// +optional
	BuildRunProgress *BuildRunSpecBuildRunProgress `json:"buildRunProgress,omitempty" tf:"build_run_progress"`
	// +optional
	BuildRunSource *BuildRunSpecBuildRunSource `json:"buildRunSource,omitempty" tf:"build_run_source"`
	// +optional
	CommitInfo *BuildRunSpecCommitInfo `json:"commitInfo,omitempty" tf:"commit_info"`
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	ProjectID *string `json:"projectID,omitempty" tf:"project_id"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type BuildRunStatus struct {
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

// BuildRunList is a list of BuildRuns
type BuildRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BuildRun CRD objects
	Items []BuildRun `json:"items,omitempty"`
}