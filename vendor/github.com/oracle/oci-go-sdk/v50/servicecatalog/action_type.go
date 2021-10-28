// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeCreated ActionTypeEnum = "CREATED"
	ActionTypeUpdated ActionTypeEnum = "UPDATED"
	ActionTypeDeleted ActionTypeEnum = "DELETED"
	ActionTypeFailed  ActionTypeEnum = "FAILED"
)

var mappingActionType = map[string]ActionTypeEnum{
	"CREATED": ActionTypeCreated,
	"UPDATED": ActionTypeUpdated,
	"DELETED": ActionTypeDeleted,
	"FAILED":  ActionTypeFailed,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionType {
		values = append(values, v)
	}
	return values
}
