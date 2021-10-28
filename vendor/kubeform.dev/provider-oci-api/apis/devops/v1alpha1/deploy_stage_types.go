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

type DeployStage struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeployStageSpec   `json:"spec,omitempty"`
	Status            DeployStageStatus `json:"status,omitempty"`
}

type DeployStageSpecApprovalPolicy struct {
	ApprovalPolicyType        *string `json:"approvalPolicyType" tf:"approval_policy_type"`
	NumberOfApprovalsRequired *int64  `json:"numberOfApprovalsRequired" tf:"number_of_approvals_required"`
}

type DeployStageSpecBlueBackendIPS struct {
	// +optional
	Items []string `json:"items,omitempty" tf:"items"`
}

type DeployStageSpecDeployStagePredecessorCollectionItems struct {
	ID *string `json:"ID" tf:"id"`
}

type DeployStageSpecDeployStagePredecessorCollection struct {
	Items []DeployStageSpecDeployStagePredecessorCollectionItems `json:"items" tf:"items"`
}

type DeployStageSpecFailurePolicy struct {
	// +optional
	FailureCount *int64 `json:"failureCount,omitempty" tf:"failure_count"`
	// +optional
	FailurePercentage *int64  `json:"failurePercentage,omitempty" tf:"failure_percentage"`
	PolicyType        *string `json:"policyType" tf:"policy_type"`
}

type DeployStageSpecGreenBackendIPS struct {
	// +optional
	Items []string `json:"items,omitempty" tf:"items"`
}

type DeployStageSpecLoadBalancerConfig struct {
	// +optional
	BackendPort *int64 `json:"backendPort,omitempty" tf:"backend_port"`
	// +optional
	ListenerName *string `json:"listenerName,omitempty" tf:"listener_name"`
	// +optional
	LoadBalancerID *string `json:"loadBalancerID,omitempty" tf:"load_balancer_id"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type DeployStageSpecRollbackPolicy struct {
	// +optional
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type"`
}

type DeployStageSpecRolloutPolicy struct {
	// +optional
	BatchCount *int64 `json:"batchCount,omitempty" tf:"batch_count"`
	// +optional
	BatchDelayInSeconds *int64 `json:"batchDelayInSeconds,omitempty" tf:"batch_delay_in_seconds"`
	// +optional
	BatchPercentage *int64  `json:"batchPercentage,omitempty" tf:"batch_percentage"`
	PolicyType      *string `json:"policyType" tf:"policy_type"`
	// +optional
	RampLimitPercent *float64 `json:"rampLimitPercent,omitempty" tf:"ramp_limit_percent"`
}

type DeployStageSpecWaitCriteria struct {
	WaitDuration *string `json:"waitDuration" tf:"wait_duration"`
	WaitType     *string `json:"waitType" tf:"wait_type"`
}

type DeployStageSpec struct {
	State *DeployStageSpecResource `json:"state,omitempty" tf:"-"`

	Resource DeployStageSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type DeployStageSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApprovalPolicy *DeployStageSpecApprovalPolicy `json:"approvalPolicy,omitempty" tf:"approval_policy"`
	// +optional
	BlueBackendIPS *DeployStageSpecBlueBackendIPS `json:"blueBackendIPS,omitempty" tf:"blue_backend_ips"`
	// +optional
	CompartmentID *string `json:"compartmentID,omitempty" tf:"compartment_id"`
	// +optional
	ComputeInstanceGroupDeployEnvironmentID *string `json:"computeInstanceGroupDeployEnvironmentID,omitempty" tf:"compute_instance_group_deploy_environment_id"`
	// +optional
	Config map[string]string `json:"config,omitempty" tf:"config"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DeployArtifactID *string `json:"deployArtifactID,omitempty" tf:"deploy_artifact_id"`
	// +optional
	DeployArtifactIDS                []string                                         `json:"deployArtifactIDS,omitempty" tf:"deploy_artifact_ids"`
	DeployPipelineID                 *string                                          `json:"deployPipelineID" tf:"deploy_pipeline_id"`
	DeployStagePredecessorCollection *DeployStageSpecDeployStagePredecessorCollection `json:"deployStagePredecessorCollection" tf:"deploy_stage_predecessor_collection"`
	DeployStageType                  *string                                          `json:"deployStageType" tf:"deploy_stage_type"`
	// +optional
	DeploymentSpecDeployArtifactID *string `json:"deploymentSpecDeployArtifactID,omitempty" tf:"deployment_spec_deploy_artifact_id"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	DockerImageDeployArtifactID *string `json:"dockerImageDeployArtifactID,omitempty" tf:"docker_image_deploy_artifact_id"`
	// +optional
	FailurePolicy *DeployStageSpecFailurePolicy `json:"failurePolicy,omitempty" tf:"failure_policy"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	FunctionDeployEnvironmentID *string `json:"functionDeployEnvironmentID,omitempty" tf:"function_deploy_environment_id"`
	// +optional
	FunctionTimeoutInSeconds *int64 `json:"functionTimeoutInSeconds,omitempty" tf:"function_timeout_in_seconds"`
	// +optional
	GreenBackendIPS *DeployStageSpecGreenBackendIPS `json:"greenBackendIPS,omitempty" tf:"green_backend_ips"`
	// +optional
	IsAsync *bool `json:"isAsync,omitempty" tf:"is_async"`
	// +optional
	IsValidationEnabled *bool `json:"isValidationEnabled,omitempty" tf:"is_validation_enabled"`
	// +optional
	KubernetesManifestDeployArtifactIDS []string `json:"kubernetesManifestDeployArtifactIDS,omitempty" tf:"kubernetes_manifest_deploy_artifact_ids"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	// +optional
	LoadBalancerConfig *DeployStageSpecLoadBalancerConfig `json:"loadBalancerConfig,omitempty" tf:"load_balancer_config"`
	// +optional
	MaxMemoryInMbs *string `json:"maxMemoryInMbs,omitempty" tf:"max_memory_in_mbs"`
	// +optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace"`
	// +optional
	OkeClusterDeployEnvironmentID *string `json:"okeClusterDeployEnvironmentID,omitempty" tf:"oke_cluster_deploy_environment_id"`
	// +optional
	ProjectID *string `json:"projectID,omitempty" tf:"project_id"`
	// +optional
	RollbackPolicy *DeployStageSpecRollbackPolicy `json:"rollbackPolicy,omitempty" tf:"rollback_policy"`
	// +optional
	RolloutPolicy *DeployStageSpecRolloutPolicy `json:"rolloutPolicy,omitempty" tf:"rollout_policy"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	SystemTags map[string]string `json:"systemTags,omitempty" tf:"system_tags"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	TrafficShiftTarget *string `json:"trafficShiftTarget,omitempty" tf:"traffic_shift_target"`
	// +optional
	WaitCriteria *DeployStageSpecWaitCriteria `json:"waitCriteria,omitempty" tf:"wait_criteria"`
}

type DeployStageStatus struct {
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

// DeployStageList is a list of DeployStages
type DeployStageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DeployStage CRD objects
	Items []DeployStage `json:"items,omitempty"`
}
