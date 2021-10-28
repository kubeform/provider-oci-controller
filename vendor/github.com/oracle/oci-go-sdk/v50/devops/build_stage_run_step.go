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

// BuildStageRunStep The details about each step in Build Stage.
type BuildStageRunStep struct {

	// Name of the step.
	Name *string `mandatory:"false" json:"name"`

	// State of the step.
	State BuildStageRunStepStateEnum `mandatory:"false" json:"state,omitempty"`

	// Time when the step started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time when the step finished.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m BuildStageRunStep) String() string {
	return common.PointerString(m)
}

// BuildStageRunStepStateEnum Enum with underlying type: string
type BuildStageRunStepStateEnum string

// Set of constants representing the allowable values for BuildStageRunStepStateEnum
const (
	BuildStageRunStepStateWaiting    BuildStageRunStepStateEnum = "WAITING"
	BuildStageRunStepStateInProgress BuildStageRunStepStateEnum = "IN_PROGRESS"
	BuildStageRunStepStateFailed     BuildStageRunStepStateEnum = "FAILED"
	BuildStageRunStepStateSucceeded  BuildStageRunStepStateEnum = "SUCCEEDED"
)

var mappingBuildStageRunStepState = map[string]BuildStageRunStepStateEnum{
	"WAITING":     BuildStageRunStepStateWaiting,
	"IN_PROGRESS": BuildStageRunStepStateInProgress,
	"FAILED":      BuildStageRunStepStateFailed,
	"SUCCEEDED":   BuildStageRunStepStateSucceeded,
}

// GetBuildStageRunStepStateEnumValues Enumerates the set of values for BuildStageRunStepStateEnum
func GetBuildStageRunStepStateEnumValues() []BuildStageRunStepStateEnum {
	values := make([]BuildStageRunStepStateEnum, 0)
	for _, v := range mappingBuildStageRunStepState {
		values = append(values, v)
	}
	return values
}
