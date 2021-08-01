// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Query API
//
// API for the Java Management Service. Use this API to view and manage Fleets.
//

package jms

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateFleet OperationTypeEnum = "CREATE_FLEET"
	OperationTypeDeleteFleet OperationTypeEnum = "DELETE_FLEET"
	OperationTypeMoveFleet   OperationTypeEnum = "MOVE_FLEET"
	OperationTypeUpdateFleet OperationTypeEnum = "UPDATE_FLEET"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"CREATE_FLEET": OperationTypeCreateFleet,
	"DELETE_FLEET": OperationTypeDeleteFleet,
	"MOVE_FLEET":   OperationTypeMoveFleet,
	"UPDATE_FLEET": OperationTypeUpdateFleet,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
