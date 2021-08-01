// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// DeployPipeline A set of stages whose predecessor relation forms a directed acyclic graph.
type DeployPipeline struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of the compartment where the pipeline is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	DeployPipelineParameters *DeployPipelineParameterCollection `mandatory:"true" json:"deployPipelineParameters"`

	// Optional description about the deployment pipeline.
	Description *string `mandatory:"false" json:"description"`

	// Deployment pipeline display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	DeployPipelineArtifacts *DeployPipelineArtifactCollection `mandatory:"false" json:"deployPipelineArtifacts"`

	DeployPipelineEnvironments *DeployPipelineEnvironmentCollection `mandatory:"false" json:"deployPipelineEnvironments"`

	// Time the deployment pipeline was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment pipeline was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the deployment pipeline.
	LifecycleState DeployPipelineLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DeployPipeline) String() string {
	return common.PointerString(m)
}

// DeployPipelineLifecycleStateEnum Enum with underlying type: string
type DeployPipelineLifecycleStateEnum string

// Set of constants representing the allowable values for DeployPipelineLifecycleStateEnum
const (
	DeployPipelineLifecycleStateCreating DeployPipelineLifecycleStateEnum = "CREATING"
	DeployPipelineLifecycleStateUpdating DeployPipelineLifecycleStateEnum = "UPDATING"
	DeployPipelineLifecycleStateActive   DeployPipelineLifecycleStateEnum = "ACTIVE"
	DeployPipelineLifecycleStateInactive DeployPipelineLifecycleStateEnum = "INACTIVE"
	DeployPipelineLifecycleStateDeleting DeployPipelineLifecycleStateEnum = "DELETING"
	DeployPipelineLifecycleStateDeleted  DeployPipelineLifecycleStateEnum = "DELETED"
	DeployPipelineLifecycleStateFailed   DeployPipelineLifecycleStateEnum = "FAILED"
)

var mappingDeployPipelineLifecycleState = map[string]DeployPipelineLifecycleStateEnum{
	"CREATING": DeployPipelineLifecycleStateCreating,
	"UPDATING": DeployPipelineLifecycleStateUpdating,
	"ACTIVE":   DeployPipelineLifecycleStateActive,
	"INACTIVE": DeployPipelineLifecycleStateInactive,
	"DELETING": DeployPipelineLifecycleStateDeleting,
	"DELETED":  DeployPipelineLifecycleStateDeleted,
	"FAILED":   DeployPipelineLifecycleStateFailed,
}

// GetDeployPipelineLifecycleStateEnumValues Enumerates the set of values for DeployPipelineLifecycleStateEnum
func GetDeployPipelineLifecycleStateEnumValues() []DeployPipelineLifecycleStateEnum {
	values := make([]DeployPipelineLifecycleStateEnum, 0)
	for _, v := range mappingDeployPipelineLifecycleState {
		values = append(values, v)
	}
	return values
}
