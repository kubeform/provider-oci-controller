// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Query API
//
// API for the Java Management Service. Use this API to view and manage Fleets.
//

package jms

// OperationStatusEnum Enum with underlying type: string
type OperationStatusEnum string

// Set of constants representing the allowable values for OperationStatusEnum
const (
	OperationStatusAccepted   OperationStatusEnum = "ACCEPTED"
	OperationStatusCanceled   OperationStatusEnum = "CANCELED"
	OperationStatusCanceling  OperationStatusEnum = "CANCELING"
	OperationStatusFailed     OperationStatusEnum = "FAILED"
	OperationStatusInProgress OperationStatusEnum = "IN_PROGRESS"
	OperationStatusSucceeded  OperationStatusEnum = "SUCCEEDED"
)

var mappingOperationStatus = map[string]OperationStatusEnum{
	"ACCEPTED":    OperationStatusAccepted,
	"CANCELED":    OperationStatusCanceled,
	"CANCELING":   OperationStatusCanceling,
	"FAILED":      OperationStatusFailed,
	"IN_PROGRESS": OperationStatusInProgress,
	"SUCCEEDED":   OperationStatusSucceeded,
}

// GetOperationStatusEnumValues Enumerates the set of values for OperationStatusEnum
func GetOperationStatusEnumValues() []OperationStatusEnum {
	values := make([]OperationStatusEnum, 0)
	for _, v := range mappingOperationStatus {
		values = append(values, v)
	}
	return values
}
