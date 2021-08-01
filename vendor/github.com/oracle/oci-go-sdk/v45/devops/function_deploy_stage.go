// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// FunctionDeployStage Specifies the Function stage.
type FunctionDeployStage struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Function environment OCID.
	FunctionDeployEnvironmentId *string `mandatory:"true" json:"functionDeployEnvironmentId"`

	// A Docker image artifact OCID.
	DockerImageDeployArtifactId *string `mandatory:"true" json:"dockerImageDeployArtifactId"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// User provided key and value pair configuration, which is assigned through constants or parameter.
	Config map[string]string `mandatory:"false" json:"config"`

	// Maximum usable memory for the Function (in MB).
	MaxMemoryInMBs *int64 `mandatory:"false" json:"maxMemoryInMBs"`

	// Timeout for execution of the Function. Value in seconds.
	FunctionTimeoutInSeconds *int `mandatory:"false" json:"functionTimeoutInSeconds"`

	// The current state of the deployment stage.
	LifecycleState DeployStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m FunctionDeployStage) GetId() *string {
	return m.Id
}

//GetDescription returns Description
func (m FunctionDeployStage) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m FunctionDeployStage) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m FunctionDeployStage) GetProjectId() *string {
	return m.ProjectId
}

//GetDeployPipelineId returns DeployPipelineId
func (m FunctionDeployStage) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetCompartmentId returns CompartmentId
func (m FunctionDeployStage) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m FunctionDeployStage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m FunctionDeployStage) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m FunctionDeployStage) GetLifecycleState() DeployStageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m FunctionDeployStage) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m FunctionDeployStage) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m FunctionDeployStage) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m FunctionDeployStage) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m FunctionDeployStage) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m FunctionDeployStage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FunctionDeployStage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFunctionDeployStage FunctionDeployStage
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeFunctionDeployStage
	}{
		"DEPLOY_FUNCTION",
		(MarshalTypeFunctionDeployStage)(m),
	}

	return json.Marshal(&s)
}
