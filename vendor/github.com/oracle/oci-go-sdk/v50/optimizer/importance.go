// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

// ImportanceEnum Enum with underlying type: string
type ImportanceEnum string

// Set of constants representing the allowable values for ImportanceEnum
const (
	ImportanceCritical ImportanceEnum = "CRITICAL"
	ImportanceHigh     ImportanceEnum = "HIGH"
	ImportanceModerate ImportanceEnum = "MODERATE"
	ImportanceLow      ImportanceEnum = "LOW"
	ImportanceMinor    ImportanceEnum = "MINOR"
)

var mappingImportance = map[string]ImportanceEnum{
	"CRITICAL": ImportanceCritical,
	"HIGH":     ImportanceHigh,
	"MODERATE": ImportanceModerate,
	"LOW":      ImportanceLow,
	"MINOR":    ImportanceMinor,
}

// GetImportanceEnumValues Enumerates the set of values for ImportanceEnum
func GetImportanceEnumValues() []ImportanceEnum {
	values := make([]ImportanceEnum, 0)
	for _, v := range mappingImportance {
		values = append(values, v)
	}
	return values
}
