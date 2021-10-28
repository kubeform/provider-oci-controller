// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring (APM) Control Plane API
//
// Provide a set of APIs for tenant to perform operations like create, update, delete and list APM domains, and also
// work request APIs to monitor progress of these operations.
//

package apmcontrolplane

// DataKeyTypesEnum Enum with underlying type: string
type DataKeyTypesEnum string

// Set of constants representing the allowable values for DataKeyTypesEnum
const (
	DataKeyTypesPrivate DataKeyTypesEnum = "PRIVATE"
	DataKeyTypesPublic  DataKeyTypesEnum = "PUBLIC"
)

var mappingDataKeyTypes = map[string]DataKeyTypesEnum{
	"PRIVATE": DataKeyTypesPrivate,
	"PUBLIC":  DataKeyTypesPublic,
}

// GetDataKeyTypesEnumValues Enumerates the set of values for DataKeyTypesEnum
func GetDataKeyTypesEnumValues() []DataKeyTypesEnum {
	values := make([]DataKeyTypesEnum, 0)
	for _, v := range mappingDataKeyTypes {
		values = append(values, v)
	}
	return values
}
