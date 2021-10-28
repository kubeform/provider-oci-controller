// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// Intersect The information about a intersect object.
type Intersect struct {

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

	// The information about the intersect all.
	IsAll *bool `mandatory:"false" json:"isAll"`

	// intersectType
	IntersectType IntersectIntersectTypeEnum `mandatory:"false" json:"intersectType,omitempty"`
}

//GetKey returns Key
func (m Intersect) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m Intersect) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m Intersect) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m Intersect) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m Intersect) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m Intersect) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m Intersect) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m Intersect) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m Intersect) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m Intersect) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m Intersect) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m Intersect) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m Intersect) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Intersect) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeIntersect Intersect
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeIntersect
	}{
		"INTERSECT_OPERATOR",
		(MarshalTypeIntersect)(m),
	}

	return json.Marshal(&s)
}

// IntersectIntersectTypeEnum Enum with underlying type: string
type IntersectIntersectTypeEnum string

// Set of constants representing the allowable values for IntersectIntersectTypeEnum
const (
	IntersectIntersectTypeName     IntersectIntersectTypeEnum = "NAME"
	IntersectIntersectTypePosition IntersectIntersectTypeEnum = "POSITION"
)

var mappingIntersectIntersectType = map[string]IntersectIntersectTypeEnum{
	"NAME":     IntersectIntersectTypeName,
	"POSITION": IntersectIntersectTypePosition,
}

// GetIntersectIntersectTypeEnumValues Enumerates the set of values for IntersectIntersectTypeEnum
func GetIntersectIntersectTypeEnumValues() []IntersectIntersectTypeEnum {
	values := make([]IntersectIntersectTypeEnum, 0)
	for _, v := range mappingIntersectIntersectType {
		values = append(values, v)
	}
	return values
}
