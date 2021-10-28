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

// DeployStageExecutionProgress Details about the execution progress of a stage in a deployment.
type DeployStageExecutionProgress interface {

	// Stage display name. Avoid entering confidential information.
	GetDeployStageDisplayName() *string

	// The OCID of the stage.
	GetDeployStageId() *string

	// Time the stage started executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeStarted() *common.SDKTime

	// Time the stage finished executing. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	GetTimeFinished() *common.SDKTime

	// The current state of the stage.
	GetStatus() DeployStageExecutionProgressStatusEnum

	GetDeployStagePredecessors() *DeployStagePredecessorCollection

	// Details about stage execution for all the target environments.
	GetDeployStageExecutionProgressDetails() []DeployStageExecutionProgressDetails
}

type deploystageexecutionprogress struct {
	JsonData                            []byte
	DeployStageDisplayName              *string                                `mandatory:"false" json:"deployStageDisplayName"`
	DeployStageId                       *string                                `mandatory:"false" json:"deployStageId"`
	TimeStarted                         *common.SDKTime                        `mandatory:"false" json:"timeStarted"`
	TimeFinished                        *common.SDKTime                        `mandatory:"false" json:"timeFinished"`
	Status                              DeployStageExecutionProgressStatusEnum `mandatory:"false" json:"status,omitempty"`
	DeployStagePredecessors             *DeployStagePredecessorCollection      `mandatory:"false" json:"deployStagePredecessors"`
	DeployStageExecutionProgressDetails []DeployStageExecutionProgressDetails  `mandatory:"false" json:"deployStageExecutionProgressDetails"`
	DeployStageType                     string                                 `json:"deployStageType"`
}

// UnmarshalJSON unmarshals json
func (m *deploystageexecutionprogress) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeploystageexecutionprogress deploystageexecutionprogress
	s := struct {
		Model Unmarshalerdeploystageexecutionprogress
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DeployStageDisplayName = s.Model.DeployStageDisplayName
	m.DeployStageId = s.Model.DeployStageId
	m.TimeStarted = s.Model.TimeStarted
	m.TimeFinished = s.Model.TimeFinished
	m.Status = s.Model.Status
	m.DeployStagePredecessors = s.Model.DeployStagePredecessors
	m.DeployStageExecutionProgressDetails = s.Model.DeployStageExecutionProgressDetails
	m.DeployStageType = s.Model.DeployStageType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *deploystageexecutionprogress) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployStageType {
	case "MANUAL_APPROVAL":
		mm := ManualApprovalDeployStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT":
		mm := ComputeInstanceGroupDeployStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_DEPLOYMENT":
		mm := OkeDeployStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LOAD_BALANCER_TRAFFIC_SHIFT":
		mm := LoadBalancerTrafficShiftDeployStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEPLOY_FUNCTION":
		mm := FunctionDeployStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INVOKE_FUNCTION":
		mm := InvokeFunctionDeployStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "WAIT":
		mm := WaitDeployStageExecutionProgress{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDeployStageDisplayName returns DeployStageDisplayName
func (m deploystageexecutionprogress) GetDeployStageDisplayName() *string {
	return m.DeployStageDisplayName
}

//GetDeployStageId returns DeployStageId
func (m deploystageexecutionprogress) GetDeployStageId() *string {
	return m.DeployStageId
}

//GetTimeStarted returns TimeStarted
func (m deploystageexecutionprogress) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

//GetTimeFinished returns TimeFinished
func (m deploystageexecutionprogress) GetTimeFinished() *common.SDKTime {
	return m.TimeFinished
}

//GetStatus returns Status
func (m deploystageexecutionprogress) GetStatus() DeployStageExecutionProgressStatusEnum {
	return m.Status
}

//GetDeployStagePredecessors returns DeployStagePredecessors
func (m deploystageexecutionprogress) GetDeployStagePredecessors() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessors
}

//GetDeployStageExecutionProgressDetails returns DeployStageExecutionProgressDetails
func (m deploystageexecutionprogress) GetDeployStageExecutionProgressDetails() []DeployStageExecutionProgressDetails {
	return m.DeployStageExecutionProgressDetails
}

func (m deploystageexecutionprogress) String() string {
	return common.PointerString(m)
}

// DeployStageExecutionProgressStatusEnum Enum with underlying type: string
type DeployStageExecutionProgressStatusEnum string

// Set of constants representing the allowable values for DeployStageExecutionProgressStatusEnum
const (
	DeployStageExecutionProgressStatusAccepted           DeployStageExecutionProgressStatusEnum = "ACCEPTED"
	DeployStageExecutionProgressStatusInProgress         DeployStageExecutionProgressStatusEnum = "IN_PROGRESS"
	DeployStageExecutionProgressStatusFailed             DeployStageExecutionProgressStatusEnum = "FAILED"
	DeployStageExecutionProgressStatusSucceeded          DeployStageExecutionProgressStatusEnum = "SUCCEEDED"
	DeployStageExecutionProgressStatusCanceling          DeployStageExecutionProgressStatusEnum = "CANCELING"
	DeployStageExecutionProgressStatusCanceled           DeployStageExecutionProgressStatusEnum = "CANCELED"
	DeployStageExecutionProgressStatusRollbackInProgress DeployStageExecutionProgressStatusEnum = "ROLLBACK_IN_PROGRESS"
	DeployStageExecutionProgressStatusRollbackSucceeded  DeployStageExecutionProgressStatusEnum = "ROLLBACK_SUCCEEDED"
	DeployStageExecutionProgressStatusRollbackFailed     DeployStageExecutionProgressStatusEnum = "ROLLBACK_FAILED"
)

var mappingDeployStageExecutionProgressStatus = map[string]DeployStageExecutionProgressStatusEnum{
	"ACCEPTED":             DeployStageExecutionProgressStatusAccepted,
	"IN_PROGRESS":          DeployStageExecutionProgressStatusInProgress,
	"FAILED":               DeployStageExecutionProgressStatusFailed,
	"SUCCEEDED":            DeployStageExecutionProgressStatusSucceeded,
	"CANCELING":            DeployStageExecutionProgressStatusCanceling,
	"CANCELED":             DeployStageExecutionProgressStatusCanceled,
	"ROLLBACK_IN_PROGRESS": DeployStageExecutionProgressStatusRollbackInProgress,
	"ROLLBACK_SUCCEEDED":   DeployStageExecutionProgressStatusRollbackSucceeded,
	"ROLLBACK_FAILED":      DeployStageExecutionProgressStatusRollbackFailed,
}

// GetDeployStageExecutionProgressStatusEnumValues Enumerates the set of values for DeployStageExecutionProgressStatusEnum
func GetDeployStageExecutionProgressStatusEnumValues() []DeployStageExecutionProgressStatusEnum {
	values := make([]DeployStageExecutionProgressStatusEnum, 0)
	for _, v := range mappingDeployStageExecutionProgressStatus {
		values = append(values, v)
	}
	return values
}
