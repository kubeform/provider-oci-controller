// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

// StrategyParameterTypeEnum Enum with underlying type: string
type StrategyParameterTypeEnum string

// Set of constants representing the allowable values for StrategyParameterTypeEnum
const (
	StrategyParameterTypeString   StrategyParameterTypeEnum = "STRING"
	StrategyParameterTypeBoolean  StrategyParameterTypeEnum = "BOOLEAN"
	StrategyParameterTypeNumber   StrategyParameterTypeEnum = "NUMBER"
	StrategyParameterTypeDatetime StrategyParameterTypeEnum = "DATETIME"
)

var mappingStrategyParameterType = map[string]StrategyParameterTypeEnum{
	"STRING":   StrategyParameterTypeString,
	"BOOLEAN":  StrategyParameterTypeBoolean,
	"NUMBER":   StrategyParameterTypeNumber,
	"DATETIME": StrategyParameterTypeDatetime,
}

// GetStrategyParameterTypeEnumValues Enumerates the set of values for StrategyParameterTypeEnum
func GetStrategyParameterTypeEnumValues() []StrategyParameterTypeEnum {
	values := make([]StrategyParameterTypeEnum, 0)
	for _, v := range mappingStrategyParameterType {
		values = append(values, v)
	}
	return values
}
