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

type Deployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSpec   `json:"spec,omitempty"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

type DeploymentSpecDeployArtifactOverrideArgumentsItems struct {
	// +optional
	DeployArtifactID *string `json:"deployArtifactID,omitempty" tf:"deploy_artifact_id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type DeploymentSpecDeployArtifactOverrideArguments struct {
	// +optional
	Items []DeploymentSpecDeployArtifactOverrideArgumentsItems `json:"items,omitempty" tf:"items"`
}

type DeploymentSpecDeployPipelineArtifactsItemsDeployPipelineStagesItems struct {
	// +optional
	DeployStageID *string `json:"deployStageID,omitempty" tf:"deploy_stage_id"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
}

type DeploymentSpecDeployPipelineArtifactsItemsDeployPipelineStages struct {
	// +optional
	Items []DeploymentSpecDeployPipelineArtifactsItemsDeployPipelineStagesItems `json:"items,omitempty" tf:"items"`
}

type DeploymentSpecDeployPipelineArtifactsItems struct {
	// +optional
	DeployArtifactID *string `json:"deployArtifactID,omitempty" tf:"deploy_artifact_id"`
	// +optional
	DeployPipelineStages *DeploymentSpecDeployPipelineArtifactsItemsDeployPipelineStages `json:"deployPipelineStages,omitempty" tf:"deploy_pipeline_stages"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
}

type DeploymentSpecDeployPipelineArtifacts struct {
	// +optional
	Items []DeploymentSpecDeployPipelineArtifactsItems `json:"items,omitempty" tf:"items"`
}

type DeploymentSpecDeployPipelineEnvironmentsItemsDeployPipelineStagesItems struct {
	// +optional
	DeployStageID *string `json:"deployStageID,omitempty" tf:"deploy_stage_id"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
}

type DeploymentSpecDeployPipelineEnvironmentsItemsDeployPipelineStages struct {
	// +optional
	Items []DeploymentSpecDeployPipelineEnvironmentsItemsDeployPipelineStagesItems `json:"items,omitempty" tf:"items"`
}

type DeploymentSpecDeployPipelineEnvironmentsItems struct {
	// +optional
	DeployEnvironmentID *string `json:"deployEnvironmentID,omitempty" tf:"deploy_environment_id"`
	// +optional
	DeployPipelineStages *DeploymentSpecDeployPipelineEnvironmentsItemsDeployPipelineStages `json:"deployPipelineStages,omitempty" tf:"deploy_pipeline_stages"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
}

type DeploymentSpecDeployPipelineEnvironments struct {
	// +optional
	Items []DeploymentSpecDeployPipelineEnvironmentsItems `json:"items,omitempty" tf:"items"`
}

type DeploymentSpecDeploymentArgumentsItems struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type DeploymentSpecDeploymentArguments struct {
	// +optional
	Items []DeploymentSpecDeploymentArgumentsItems `json:"items,omitempty" tf:"items"`
}

type DeploymentSpecDeploymentExecutionProgress struct {
	// +optional
	DeployStageExecutionProgress map[string]string `json:"deployStageExecutionProgress,omitempty" tf:"deploy_stage_execution_progress"`
	// +optional
	TimeFinished *string `json:"timeFinished,omitempty" tf:"time_finished"`
	// +optional
	TimeStarted *string `json:"timeStarted,omitempty" tf:"time_started"`
}

type DeploymentSpec struct {
	State *DeploymentSpecResource `json:"state,omitempty" tf:"-"`

	Resource DeploymentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DeploymentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DeployArtifactOverrideArguments *DeploymentSpecDeployArtifactOverrideArguments `json:"deployArtifactOverrideArguments,omitempty" tf:"deploy_artifact_override_arguments"`
	// +optional
	DeployPipelineArtifacts *DeploymentSpecDeployPipelineArtifacts `json:"deployPipelineArtifacts,omitempty" tf:"deploy_pipeline_artifacts"`
	// +optional
	DeployPipelineEnvironments *DeploymentSpecDeployPipelineEnvironments `json:"deployPipelineEnvironments,omitempty" tf:"deploy_pipeline_environments"`
	DeployPipelineID           *string                                   `json:"deployPipelineID" tf:"deploy_pipeline_id"`
	// +optional
	DeployStageID *string `json:"deployStageID,omitempty" tf:"deploy_stage_id"`
	// +optional
	DeploymentArguments *DeploymentSpecDeploymentArguments `json:"deploymentArguments,omitempty" tf:"deployment_arguments"`
	// +optional
	DeploymentExecutionProgress *DeploymentSpecDeploymentExecutionProgress `json:"deploymentExecutionProgress,omitempty" tf:"deployment_execution_progress"`
	DeploymentType              *string                                    `json:"deploymentType" tf:"deployment_type"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	PreviousDeploymentID *string `json:"previousDeploymentID,omitempty" tf:"previous_deployment_id"`
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

type DeploymentStatus struct {
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

// DeploymentList is a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Deployment CRD objects
	Items []Deployment `json:"items,omitempty"`
}
