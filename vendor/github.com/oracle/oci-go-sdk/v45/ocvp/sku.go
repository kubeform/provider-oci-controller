// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud VMware Solution API
//
// Use this API to manage your Oracle Cloud VMware Solution (https://docs.cloud.oracle.com/iaas/Content/VMware/Concepts/ocvsoverview.htm).
//

package ocvp

// SkuEnum Enum with underlying type: string
type SkuEnum string

// Set of constants representing the allowable values for SkuEnum
const (
	SkuHour       SkuEnum = "HOUR"
	SkuMonth      SkuEnum = "MONTH"
	SkuOneYear    SkuEnum = "ONE_YEAR"
	SkuThreeYears SkuEnum = "THREE_YEARS"
)

var mappingSku = map[string]SkuEnum{
	"HOUR":        SkuHour,
	"MONTH":       SkuMonth,
	"ONE_YEAR":    SkuOneYear,
	"THREE_YEARS": SkuThreeYears,
}

// GetSkuEnumValues Enumerates the set of values for SkuEnum
func GetSkuEnumValues() []SkuEnum {
	values := make([]SkuEnum, 0)
	for _, v := range mappingSku {
		values = append(values, v)
	}
	return values
}
