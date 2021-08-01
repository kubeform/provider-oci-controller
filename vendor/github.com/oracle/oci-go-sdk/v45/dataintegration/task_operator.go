// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// TaskOperator An operator for task
type TaskOperator struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Details about the operator.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// An array of parameters used in the data flow.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	// The number of retry attempts.
	RetryAttempts *int `mandatory:"false" json:"retryAttempts"`

	// The retry delay, the unit for measurement is in the property retry delay unit.
	RetryDelay *float64 `mandatory:"false" json:"retryDelay"`

	// The expected duration for the task run.
	ExpectedDuration *float64 `mandatory:"false" json:"expectedDuration"`

	Task Task `mandatory:"false" json:"task"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	// The unit for the retry delay.
	RetryDelayUnit TaskOperatorRetryDelayUnitEnum `mandatory:"false" json:"retryDelayUnit,omitempty"`

	// The expected duration unit of measure.
	ExpectedDurationUnit TaskOperatorExpectedDurationUnitEnum `mandatory:"false" json:"expectedDurationUnit,omitempty"`

	// The type of the task referenced in the task property.
	TaskType TaskOperatorTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	// The merge condition. The conditions are
	// ALL_SUCCESS - All the preceeding operators need to be successful.
	// ALL_FAILED - All the preceeding operators should have failed.
	// ALL_COMPLETE - All the preceeding operators should have completed. It could have executed successfully or failed.
	TriggerRule TaskOperatorTriggerRuleEnum `mandatory:"false" json:"triggerRule,omitempty"`
}

//GetKey returns Key
func (m TaskOperator) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m TaskOperator) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m TaskOperator) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m TaskOperator) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m TaskOperator) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m TaskOperator) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m TaskOperator) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m TaskOperator) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m TaskOperator) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m TaskOperator) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m TaskOperator) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m TaskOperator) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m TaskOperator) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TaskOperator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTaskOperator TaskOperator
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTaskOperator
	}{
		"TASK_OPERATOR",
		(MarshalTypeTaskOperator)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TaskOperator) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                    *string                              `json:"key"`
		ModelVersion           *string                              `json:"modelVersion"`
		ParentRef              *ParentReference                     `json:"parentRef"`
		Name                   *string                              `json:"name"`
		Description            *string                              `json:"description"`
		ObjectVersion          *int                                 `json:"objectVersion"`
		InputPorts             []InputPort                          `json:"inputPorts"`
		OutputPorts            []OutputPort                         `json:"outputPorts"`
		ObjectStatus           *int                                 `json:"objectStatus"`
		Identifier             *string                              `json:"identifier"`
		Parameters             []Parameter                          `json:"parameters"`
		OpConfigValues         *ConfigValues                        `json:"opConfigValues"`
		RetryAttempts          *int                                 `json:"retryAttempts"`
		RetryDelayUnit         TaskOperatorRetryDelayUnitEnum       `json:"retryDelayUnit"`
		RetryDelay             *float64                             `json:"retryDelay"`
		ExpectedDuration       *float64                             `json:"expectedDuration"`
		ExpectedDurationUnit   TaskOperatorExpectedDurationUnitEnum `json:"expectedDurationUnit"`
		TaskType               TaskOperatorTaskTypeEnum             `json:"taskType"`
		Task                   task                                 `json:"task"`
		TriggerRule            TaskOperatorTriggerRuleEnum          `json:"triggerRule"`
		ConfigProviderDelegate *ConfigProvider                      `json:"configProviderDelegate"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.InputPorts = make([]InputPort, len(model.InputPorts))
	for i, n := range model.InputPorts {
		m.InputPorts[i] = n
	}

	m.OutputPorts = make([]OutputPort, len(model.OutputPorts))
	for i, n := range model.OutputPorts {
		m.OutputPorts[i] = n
	}

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	m.Parameters = make([]Parameter, len(model.Parameters))
	for i, n := range model.Parameters {
		m.Parameters[i] = n
	}

	m.OpConfigValues = model.OpConfigValues

	m.RetryAttempts = model.RetryAttempts

	m.RetryDelayUnit = model.RetryDelayUnit

	m.RetryDelay = model.RetryDelay

	m.ExpectedDuration = model.ExpectedDuration

	m.ExpectedDurationUnit = model.ExpectedDurationUnit

	m.TaskType = model.TaskType

	nn, e = model.Task.UnmarshalPolymorphicJSON(model.Task.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Task = nn.(Task)
	} else {
		m.Task = nil
	}

	m.TriggerRule = model.TriggerRule

	m.ConfigProviderDelegate = model.ConfigProviderDelegate

	return
}

// TaskOperatorRetryDelayUnitEnum Enum with underlying type: string
type TaskOperatorRetryDelayUnitEnum string

// Set of constants representing the allowable values for TaskOperatorRetryDelayUnitEnum
const (
	TaskOperatorRetryDelayUnitSeconds TaskOperatorRetryDelayUnitEnum = "SECONDS"
	TaskOperatorRetryDelayUnitMinutes TaskOperatorRetryDelayUnitEnum = "MINUTES"
	TaskOperatorRetryDelayUnitHours   TaskOperatorRetryDelayUnitEnum = "HOURS"
	TaskOperatorRetryDelayUnitDays    TaskOperatorRetryDelayUnitEnum = "DAYS"
)

var mappingTaskOperatorRetryDelayUnit = map[string]TaskOperatorRetryDelayUnitEnum{
	"SECONDS": TaskOperatorRetryDelayUnitSeconds,
	"MINUTES": TaskOperatorRetryDelayUnitMinutes,
	"HOURS":   TaskOperatorRetryDelayUnitHours,
	"DAYS":    TaskOperatorRetryDelayUnitDays,
}

// GetTaskOperatorRetryDelayUnitEnumValues Enumerates the set of values for TaskOperatorRetryDelayUnitEnum
func GetTaskOperatorRetryDelayUnitEnumValues() []TaskOperatorRetryDelayUnitEnum {
	values := make([]TaskOperatorRetryDelayUnitEnum, 0)
	for _, v := range mappingTaskOperatorRetryDelayUnit {
		values = append(values, v)
	}
	return values
}

// TaskOperatorExpectedDurationUnitEnum Enum with underlying type: string
type TaskOperatorExpectedDurationUnitEnum string

// Set of constants representing the allowable values for TaskOperatorExpectedDurationUnitEnum
const (
	TaskOperatorExpectedDurationUnitSeconds TaskOperatorExpectedDurationUnitEnum = "SECONDS"
	TaskOperatorExpectedDurationUnitMinutes TaskOperatorExpectedDurationUnitEnum = "MINUTES"
	TaskOperatorExpectedDurationUnitHours   TaskOperatorExpectedDurationUnitEnum = "HOURS"
	TaskOperatorExpectedDurationUnitDays    TaskOperatorExpectedDurationUnitEnum = "DAYS"
)

var mappingTaskOperatorExpectedDurationUnit = map[string]TaskOperatorExpectedDurationUnitEnum{
	"SECONDS": TaskOperatorExpectedDurationUnitSeconds,
	"MINUTES": TaskOperatorExpectedDurationUnitMinutes,
	"HOURS":   TaskOperatorExpectedDurationUnitHours,
	"DAYS":    TaskOperatorExpectedDurationUnitDays,
}

// GetTaskOperatorExpectedDurationUnitEnumValues Enumerates the set of values for TaskOperatorExpectedDurationUnitEnum
func GetTaskOperatorExpectedDurationUnitEnumValues() []TaskOperatorExpectedDurationUnitEnum {
	values := make([]TaskOperatorExpectedDurationUnitEnum, 0)
	for _, v := range mappingTaskOperatorExpectedDurationUnit {
		values = append(values, v)
	}
	return values
}

// TaskOperatorTaskTypeEnum Enum with underlying type: string
type TaskOperatorTaskTypeEnum string

// Set of constants representing the allowable values for TaskOperatorTaskTypeEnum
const (
	TaskOperatorTaskTypePipelineTask    TaskOperatorTaskTypeEnum = "PIPELINE_TASK"
	TaskOperatorTaskTypeIntegrationTask TaskOperatorTaskTypeEnum = "INTEGRATION_TASK"
	TaskOperatorTaskTypeDataLoaderTask  TaskOperatorTaskTypeEnum = "DATA_LOADER_TASK"
	TaskOperatorTaskTypeSqlTask         TaskOperatorTaskTypeEnum = "SQL_TASK"
	TaskOperatorTaskTypeOciDataflowTask TaskOperatorTaskTypeEnum = "OCI_DATAFLOW_TASK"
	TaskOperatorTaskTypeRestTask        TaskOperatorTaskTypeEnum = "REST_TASK"
)

var mappingTaskOperatorTaskType = map[string]TaskOperatorTaskTypeEnum{
	"PIPELINE_TASK":     TaskOperatorTaskTypePipelineTask,
	"INTEGRATION_TASK":  TaskOperatorTaskTypeIntegrationTask,
	"DATA_LOADER_TASK":  TaskOperatorTaskTypeDataLoaderTask,
	"SQL_TASK":          TaskOperatorTaskTypeSqlTask,
	"OCI_DATAFLOW_TASK": TaskOperatorTaskTypeOciDataflowTask,
	"REST_TASK":         TaskOperatorTaskTypeRestTask,
}

// GetTaskOperatorTaskTypeEnumValues Enumerates the set of values for TaskOperatorTaskTypeEnum
func GetTaskOperatorTaskTypeEnumValues() []TaskOperatorTaskTypeEnum {
	values := make([]TaskOperatorTaskTypeEnum, 0)
	for _, v := range mappingTaskOperatorTaskType {
		values = append(values, v)
	}
	return values
}

// TaskOperatorTriggerRuleEnum Enum with underlying type: string
type TaskOperatorTriggerRuleEnum string

// Set of constants representing the allowable values for TaskOperatorTriggerRuleEnum
const (
	TaskOperatorTriggerRuleSuccess  TaskOperatorTriggerRuleEnum = "ALL_SUCCESS"
	TaskOperatorTriggerRuleFailed   TaskOperatorTriggerRuleEnum = "ALL_FAILED"
	TaskOperatorTriggerRuleComplete TaskOperatorTriggerRuleEnum = "ALL_COMPLETE"
)

var mappingTaskOperatorTriggerRule = map[string]TaskOperatorTriggerRuleEnum{
	"ALL_SUCCESS":  TaskOperatorTriggerRuleSuccess,
	"ALL_FAILED":   TaskOperatorTriggerRuleFailed,
	"ALL_COMPLETE": TaskOperatorTriggerRuleComplete,
}

// GetTaskOperatorTriggerRuleEnumValues Enumerates the set of values for TaskOperatorTriggerRuleEnum
func GetTaskOperatorTriggerRuleEnumValues() []TaskOperatorTriggerRuleEnum {
	values := make([]TaskOperatorTriggerRuleEnum, 0)
	for _, v := range mappingTaskOperatorTriggerRule {
		values = append(values, v)
	}
	return values
}
