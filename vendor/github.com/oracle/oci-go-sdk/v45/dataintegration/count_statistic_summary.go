// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// CountStatisticSummary Details of the count statistic summary object.
type CountStatisticSummary struct {

	// The type of object for the count statistic object.
	ObjectType CountStatisticSummaryObjectTypeEnum `mandatory:"false" json:"objectType,omitempty"`

	// The value for the count statistic object.
	ObjectCount *int64 `mandatory:"false" json:"objectCount"`
}

func (m CountStatisticSummary) String() string {
	return common.PointerString(m)
}

// CountStatisticSummaryObjectTypeEnum Enum with underlying type: string
type CountStatisticSummaryObjectTypeEnum string

// Set of constants representing the allowable values for CountStatisticSummaryObjectTypeEnum
const (
	CountStatisticSummaryObjectTypeProject     CountStatisticSummaryObjectTypeEnum = "PROJECT"
	CountStatisticSummaryObjectTypeFolder      CountStatisticSummaryObjectTypeEnum = "FOLDER"
	CountStatisticSummaryObjectTypeDataFlow    CountStatisticSummaryObjectTypeEnum = "DATA_FLOW"
	CountStatisticSummaryObjectTypeDataAsset   CountStatisticSummaryObjectTypeEnum = "DATA_ASSET"
	CountStatisticSummaryObjectTypeConnection  CountStatisticSummaryObjectTypeEnum = "CONNECTION"
	CountStatisticSummaryObjectTypeTask        CountStatisticSummaryObjectTypeEnum = "TASK"
	CountStatisticSummaryObjectTypeApplication CountStatisticSummaryObjectTypeEnum = "APPLICATION"
)

var mappingCountStatisticSummaryObjectType = map[string]CountStatisticSummaryObjectTypeEnum{
	"PROJECT":     CountStatisticSummaryObjectTypeProject,
	"FOLDER":      CountStatisticSummaryObjectTypeFolder,
	"DATA_FLOW":   CountStatisticSummaryObjectTypeDataFlow,
	"DATA_ASSET":  CountStatisticSummaryObjectTypeDataAsset,
	"CONNECTION":  CountStatisticSummaryObjectTypeConnection,
	"TASK":        CountStatisticSummaryObjectTypeTask,
	"APPLICATION": CountStatisticSummaryObjectTypeApplication,
}

// GetCountStatisticSummaryObjectTypeEnumValues Enumerates the set of values for CountStatisticSummaryObjectTypeEnum
func GetCountStatisticSummaryObjectTypeEnumValues() []CountStatisticSummaryObjectTypeEnum {
	values := make([]CountStatisticSummaryObjectTypeEnum, 0)
	for _, v := range mappingCountStatisticSummaryObjectType {
		values = append(values, v)
	}
	return values
}
