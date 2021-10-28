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

// TriggerAction The action to be performed
type TriggerAction interface {
	GetFilter() Filter
}

type triggeraction struct {
	JsonData []byte
	Filter   Filter `mandatory:"false" json:"filter"`
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *triggeraction) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertriggeraction triggeraction
	s := struct {
		Model Unmarshalertriggeraction
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Filter = s.Model.Filter
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *triggeraction) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "TRIGGER_BUILD_PIPELINE":
		mm := TriggerBuildPipelineAction{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetFilter returns Filter
func (m triggeraction) GetFilter() Filter {
	return m.Filter
}

func (m triggeraction) String() string {
	return common.PointerString(m)
}

// TriggerActionTypeEnum Enum with underlying type: string
type TriggerActionTypeEnum string

// Set of constants representing the allowable values for TriggerActionTypeEnum
const (
	TriggerActionTypeTriggerBuildPipeline TriggerActionTypeEnum = "TRIGGER_BUILD_PIPELINE"
)

var mappingTriggerActionType = map[string]TriggerActionTypeEnum{
	"TRIGGER_BUILD_PIPELINE": TriggerActionTypeTriggerBuildPipeline,
}

// GetTriggerActionTypeEnumValues Enumerates the set of values for TriggerActionTypeEnum
func GetTriggerActionTypeEnumValues() []TriggerActionTypeEnum {
	values := make([]TriggerActionTypeEnum, 0)
	for _, v := range mappingTriggerActionType {
		values = append(values, v)
	}
	return values
}
