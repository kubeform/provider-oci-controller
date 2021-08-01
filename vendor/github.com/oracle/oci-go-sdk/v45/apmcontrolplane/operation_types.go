// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring (APM) Control Plane API
//
// Provide a set of APIs for tenant to perform operations like create, update, delete and list APM domains, and also
// work request APIs to monitor progress of these operations.
//

package apmcontrolplane

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateApmDomain  OperationTypesEnum = "CREATE_APM_DOMAIN"
	OperationTypesUpdateApmDomain  OperationTypesEnum = "UPDATE_APM_DOMAIN"
	OperationTypesDeleteApmDomain  OperationTypesEnum = "DELETE_APM_DOMAIN"
	OperationTypesGenerateDataKeys OperationTypesEnum = "GENERATE_DATA_KEYS"
	OperationTypesRemoveDataKeys   OperationTypesEnum = "REMOVE_DATA_KEYS"
)

var mappingOperationTypes = map[string]OperationTypesEnum{
	"CREATE_APM_DOMAIN":  OperationTypesCreateApmDomain,
	"UPDATE_APM_DOMAIN":  OperationTypesUpdateApmDomain,
	"DELETE_APM_DOMAIN":  OperationTypesDeleteApmDomain,
	"GENERATE_DATA_KEYS": OperationTypesGenerateDataKeys,
	"REMOVE_DATA_KEYS":   OperationTypesRemoveDataKeys,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypes {
		values = append(values, v)
	}
	return values
}
