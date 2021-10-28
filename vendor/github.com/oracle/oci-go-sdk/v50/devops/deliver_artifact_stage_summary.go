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
	"github.com/oracle/oci-go-sdk/v50/common"
)

// DeliverArtifactStageSummary Specifies the DeliverArtifact Stage.
type DeliverArtifactStageSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Build Project Identifier
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Build Pipeline Identifier
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	DeliverArtifactCollection *DeliverArtifactCollection `mandatory:"true" json:"deliverArtifactCollection"`

	// Stage identifier which can be renamed and is not necessarily unique
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time at which the Stage was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time at which the Stage was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Optional description about the BuildStage
	Description *string `mandatory:"false" json:"description"`

	BuildPipelineStagePredecessorCollection *BuildPipelineStagePredecessorCollection `mandatory:"false" json:"buildPipelineStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the Stage.
	LifecycleState BuildPipelineStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m DeliverArtifactStageSummary) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m DeliverArtifactStageSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m DeliverArtifactStageSummary) GetProjectId() *string {
	return m.ProjectId
}

//GetBuildPipelineId returns BuildPipelineId
func (m DeliverArtifactStageSummary) GetBuildPipelineId() *string {
	return m.BuildPipelineId
}

//GetCompartmentId returns CompartmentId
func (m DeliverArtifactStageSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m DeliverArtifactStageSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m DeliverArtifactStageSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m DeliverArtifactStageSummary) GetLifecycleState() BuildPipelineStageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m DeliverArtifactStageSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetDescription returns Description
func (m DeliverArtifactStageSummary) GetDescription() *string {
	return m.Description
}

//GetBuildPipelineStagePredecessorCollection returns BuildPipelineStagePredecessorCollection
func (m DeliverArtifactStageSummary) GetBuildPipelineStagePredecessorCollection() *BuildPipelineStagePredecessorCollection {
	return m.BuildPipelineStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m DeliverArtifactStageSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m DeliverArtifactStageSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m DeliverArtifactStageSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m DeliverArtifactStageSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DeliverArtifactStageSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeliverArtifactStageSummary DeliverArtifactStageSummary
	s := struct {
		DiscriminatorParam string `json:"buildPipelineStageType"`
		MarshalTypeDeliverArtifactStageSummary
	}{
		"DELIVER_ARTIFACT",
		(MarshalTypeDeliverArtifactStageSummary)(m),
	}

	return json.Marshal(&s)
}
