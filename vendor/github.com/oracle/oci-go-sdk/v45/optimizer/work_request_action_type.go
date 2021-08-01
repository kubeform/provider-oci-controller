// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

// WorkRequestActionTypeEnum Enum with underlying type: string
type WorkRequestActionTypeEnum string

// Set of constants representing the allowable values for WorkRequestActionTypeEnum
const (
	WorkRequestActionTypeCreated    WorkRequestActionTypeEnum = "CREATED"
	WorkRequestActionTypeUpdated    WorkRequestActionTypeEnum = "UPDATED"
	WorkRequestActionTypeDeleted    WorkRequestActionTypeEnum = "DELETED"
	WorkRequestActionTypeInProgress WorkRequestActionTypeEnum = "IN_PROGRESS"
	WorkRequestActionTypeRelated    WorkRequestActionTypeEnum = "RELATED"
)

var mappingWorkRequestActionType = map[string]WorkRequestActionTypeEnum{
	"CREATED":     WorkRequestActionTypeCreated,
	"UPDATED":     WorkRequestActionTypeUpdated,
	"DELETED":     WorkRequestActionTypeDeleted,
	"IN_PROGRESS": WorkRequestActionTypeInProgress,
	"RELATED":     WorkRequestActionTypeRelated,
}

// GetWorkRequestActionTypeEnumValues Enumerates the set of values for WorkRequestActionTypeEnum
func GetWorkRequestActionTypeEnumValues() []WorkRequestActionTypeEnum {
	values := make([]WorkRequestActionTypeEnum, 0)
	for _, v := range mappingWorkRequestActionType {
		values = append(values, v)
	}
	return values
}
