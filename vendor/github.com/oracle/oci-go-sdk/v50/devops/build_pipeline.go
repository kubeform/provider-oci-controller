// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// BuildPipeline A set of stages whose predecessor relation forms a directed acyclic graph.
type BuildPipeline struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Project Identifier
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Optional description about the BuildPipeline
	Description *string `mandatory:"false" json:"description"`

	// BuildPipeline identifier which can be renamed and is not necessarily unique
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the BuildPipeline was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time at which the BuildPipeline was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the BuildPipeline.
	LifecycleState BuildPipelineLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	BuildPipelineParameters *BuildPipelineParameterCollection `mandatory:"false" json:"buildPipelineParameters"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m BuildPipeline) String() string {
	return common.PointerString(m)
}

// BuildPipelineLifecycleStateEnum Enum with underlying type: string
type BuildPipelineLifecycleStateEnum string

// Set of constants representing the allowable values for BuildPipelineLifecycleStateEnum
const (
	BuildPipelineLifecycleStateCreating BuildPipelineLifecycleStateEnum = "CREATING"
	BuildPipelineLifecycleStateUpdating BuildPipelineLifecycleStateEnum = "UPDATING"
	BuildPipelineLifecycleStateActive   BuildPipelineLifecycleStateEnum = "ACTIVE"
	BuildPipelineLifecycleStateInactive BuildPipelineLifecycleStateEnum = "INACTIVE"
	BuildPipelineLifecycleStateDeleting BuildPipelineLifecycleStateEnum = "DELETING"
	BuildPipelineLifecycleStateDeleted  BuildPipelineLifecycleStateEnum = "DELETED"
	BuildPipelineLifecycleStateFailed   BuildPipelineLifecycleStateEnum = "FAILED"
)

var mappingBuildPipelineLifecycleState = map[string]BuildPipelineLifecycleStateEnum{
	"CREATING": BuildPipelineLifecycleStateCreating,
	"UPDATING": BuildPipelineLifecycleStateUpdating,
	"ACTIVE":   BuildPipelineLifecycleStateActive,
	"INACTIVE": BuildPipelineLifecycleStateInactive,
	"DELETING": BuildPipelineLifecycleStateDeleting,
	"DELETED":  BuildPipelineLifecycleStateDeleted,
	"FAILED":   BuildPipelineLifecycleStateFailed,
}

// GetBuildPipelineLifecycleStateEnumValues Enumerates the set of values for BuildPipelineLifecycleStateEnum
func GetBuildPipelineLifecycleStateEnumValues() []BuildPipelineLifecycleStateEnum {
	values := make([]BuildPipelineLifecycleStateEnum, 0)
	for _, v := range mappingBuildPipelineLifecycleState {
		values = append(values, v)
	}
	return values
}
