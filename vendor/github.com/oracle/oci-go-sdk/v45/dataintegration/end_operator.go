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

// EndOperator Represents end of a pipeline
type EndOperator struct {

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

	// The merge condition. The conditions are
	// ALL_SUCCESS - All the preceeding operators need to be successful.
	// ALL_FAILED - All the preceeding operators should have failed.
	// ALL_COMPLETE - All the preceeding operators should have completed. It could have executed successfully or failed.
	TriggerRule EndOperatorTriggerRuleEnum `mandatory:"false" json:"triggerRule,omitempty"`
}

//GetKey returns Key
func (m EndOperator) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m EndOperator) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m EndOperator) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m EndOperator) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m EndOperator) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m EndOperator) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m EndOperator) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m EndOperator) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m EndOperator) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m EndOperator) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m EndOperator) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m EndOperator) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m EndOperator) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m EndOperator) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEndOperator EndOperator
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeEndOperator
	}{
		"END_OPERATOR",
		(MarshalTypeEndOperator)(m),
	}

	return json.Marshal(&s)
}

// EndOperatorTriggerRuleEnum Enum with underlying type: string
type EndOperatorTriggerRuleEnum string

// Set of constants representing the allowable values for EndOperatorTriggerRuleEnum
const (
	EndOperatorTriggerRuleSuccess  EndOperatorTriggerRuleEnum = "ALL_SUCCESS"
	EndOperatorTriggerRuleFailed   EndOperatorTriggerRuleEnum = "ALL_FAILED"
	EndOperatorTriggerRuleComplete EndOperatorTriggerRuleEnum = "ALL_COMPLETE"
)

var mappingEndOperatorTriggerRule = map[string]EndOperatorTriggerRuleEnum{
	"ALL_SUCCESS":  EndOperatorTriggerRuleSuccess,
	"ALL_FAILED":   EndOperatorTriggerRuleFailed,
	"ALL_COMPLETE": EndOperatorTriggerRuleComplete,
}

// GetEndOperatorTriggerRuleEnumValues Enumerates the set of values for EndOperatorTriggerRuleEnum
func GetEndOperatorTriggerRuleEnumValues() []EndOperatorTriggerRuleEnum {
	values := make([]EndOperatorTriggerRuleEnum, 0)
	for _, v := range mappingEndOperatorTriggerRule {
		values = append(values, v)
	}
	return values
}
