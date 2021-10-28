// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// FunctionLibraryDetails The details including name, description for the function library, which is a container for user defined functions.
type FunctionLibraryDetails struct {

	// Generated key that can be used in API calls to identify FunctionLibrary.
	Key *string `mandatory:"true" json:"key"`

	// The type of the object.
	ModelType FunctionLibraryDetailsModelTypeEnum `mandatory:"true" json:"modelType"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// A user defined description for the FunctionLibrary.
	Description *string `mandatory:"false" json:"description"`

	// The category name.
	CategoryName *string `mandatory:"false" json:"categoryName"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m FunctionLibraryDetails) String() string {
	return common.PointerString(m)
}

// FunctionLibraryDetailsModelTypeEnum Enum with underlying type: string
type FunctionLibraryDetailsModelTypeEnum string

// Set of constants representing the allowable values for FunctionLibraryDetailsModelTypeEnum
const (
	FunctionLibraryDetailsModelTypeFunctionLibrary FunctionLibraryDetailsModelTypeEnum = "FUNCTION_LIBRARY"
)

var mappingFunctionLibraryDetailsModelType = map[string]FunctionLibraryDetailsModelTypeEnum{
	"FUNCTION_LIBRARY": FunctionLibraryDetailsModelTypeFunctionLibrary,
}

// GetFunctionLibraryDetailsModelTypeEnumValues Enumerates the set of values for FunctionLibraryDetailsModelTypeEnum
func GetFunctionLibraryDetailsModelTypeEnumValues() []FunctionLibraryDetailsModelTypeEnum {
	values := make([]FunctionLibraryDetailsModelTypeEnum, 0)
	for _, v := range mappingFunctionLibraryDetailsModelType {
		values = append(values, v)
	}
	return values
}
