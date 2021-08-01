// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

// RuleTypeEnum Enum with underlying type: string
type RuleTypeEnum string

// Set of constants representing the allowable values for RuleTypeEnum
const (
	RuleTypePrimarykey RuleTypeEnum = "PRIMARYKEY"
	RuleTypeForeignkey RuleTypeEnum = "FOREIGNKEY"
	RuleTypeUniquekey  RuleTypeEnum = "UNIQUEKEY"
)

var mappingRuleType = map[string]RuleTypeEnum{
	"PRIMARYKEY": RuleTypePrimarykey,
	"FOREIGNKEY": RuleTypeForeignkey,
	"UNIQUEKEY":  RuleTypeUniquekey,
}

// GetRuleTypeEnumValues Enumerates the set of values for RuleTypeEnum
func GetRuleTypeEnumValues() []RuleTypeEnum {
	values := make([]RuleTypeEnum, 0)
	for _, v := range mappingRuleType {
		values = append(values, v)
	}
	return values
}
