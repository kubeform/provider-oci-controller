// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

// ModelDeploymentTypeEnum Enum with underlying type: string
type ModelDeploymentTypeEnum string

// Set of constants representing the allowable values for ModelDeploymentTypeEnum
const (
	ModelDeploymentTypeSingleModel ModelDeploymentTypeEnum = "SINGLE_MODEL"
)

var mappingModelDeploymentType = map[string]ModelDeploymentTypeEnum{
	"SINGLE_MODEL": ModelDeploymentTypeSingleModel,
}

// GetModelDeploymentTypeEnumValues Enumerates the set of values for ModelDeploymentTypeEnum
func GetModelDeploymentTypeEnumValues() []ModelDeploymentTypeEnum {
	values := make([]ModelDeploymentTypeEnum, 0)
	for _, v := range mappingModelDeploymentType {
		values = append(values, v)
	}
	return values
}
