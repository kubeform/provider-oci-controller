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

// DeployStageExecutionStep Details about each steps in stage execution for a target environment.
type DeployStageExecutionStep struct {

	// Name of the step.
	Name *string `mandatory:"false" json:"name"`

	// State of the step.
	State DeployStageExecutionStepStateEnum `mandatory:"false" json:"state,omitempty"`

	// Time when the step started.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Time when the step finished.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m DeployStageExecutionStep) String() string {
	return common.PointerString(m)
}

// DeployStageExecutionStepStateEnum Enum with underlying type: string
type DeployStageExecutionStepStateEnum string

// Set of constants representing the allowable values for DeployStageExecutionStepStateEnum
const (
	DeployStageExecutionStepStateWaiting    DeployStageExecutionStepStateEnum = "WAITING"
	DeployStageExecutionStepStateInProgress DeployStageExecutionStepStateEnum = "IN_PROGRESS"
	DeployStageExecutionStepStateFailed     DeployStageExecutionStepStateEnum = "FAILED"
	DeployStageExecutionStepStateSucceeded  DeployStageExecutionStepStateEnum = "SUCCEEDED"
	DeployStageExecutionStepStateCanceled   DeployStageExecutionStepStateEnum = "CANCELED"
)

var mappingDeployStageExecutionStepState = map[string]DeployStageExecutionStepStateEnum{
	"WAITING":     DeployStageExecutionStepStateWaiting,
	"IN_PROGRESS": DeployStageExecutionStepStateInProgress,
	"FAILED":      DeployStageExecutionStepStateFailed,
	"SUCCEEDED":   DeployStageExecutionStepStateSucceeded,
	"CANCELED":    DeployStageExecutionStepStateCanceled,
}

// GetDeployStageExecutionStepStateEnumValues Enumerates the set of values for DeployStageExecutionStepStateEnum
func GetDeployStageExecutionStepStateEnumValues() []DeployStageExecutionStepStateEnum {
	values := make([]DeployStageExecutionStepStateEnum, 0)
	for _, v := range mappingDeployStageExecutionStepState {
		values = append(values, v)
	}
	return values
}
