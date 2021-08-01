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

// CreateDeployPipelineDeploymentDetails Details of the new deployment to be created that will run all the stages in the pipeline.
type CreateDeployPipelineDeploymentDetails struct {

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// Deployment display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DeploymentArguments *DeploymentArgumentCollection `mandatory:"false" json:"deploymentArguments"`

	DeployArtifactOverrideArguments *DeployArtifactOverrideArgumentCollection `mandatory:"false" json:"deployArtifactOverrideArguments"`
}

//GetDeployPipelineId returns DeployPipelineId
func (m CreateDeployPipelineDeploymentDetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetDisplayName returns DisplayName
func (m CreateDeployPipelineDeploymentDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m CreateDeployPipelineDeploymentDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateDeployPipelineDeploymentDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateDeployPipelineDeploymentDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateDeployPipelineDeploymentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateDeployPipelineDeploymentDetails CreateDeployPipelineDeploymentDetails
	s := struct {
		DiscriminatorParam string `json:"deploymentType"`
		MarshalTypeCreateDeployPipelineDeploymentDetails
	}{
		"PIPELINE_DEPLOYMENT",
		(MarshalTypeCreateDeployPipelineDeploymentDetails)(m),
	}

	return json.Marshal(&s)
}
