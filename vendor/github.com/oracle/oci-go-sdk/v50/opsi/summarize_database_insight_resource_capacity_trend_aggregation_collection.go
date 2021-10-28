// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection Collection of resource capacity trend.
type SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (example: CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Identifies the units of the current resource metric (CORES, GB).
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Capacity Data with time interval
	CapacityData []ResourceCapacityTrendAggregation `mandatory:"true" json:"capacityData"`
}

func (m SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu       SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricStorage   SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricIo        SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemoryPga SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "MEMORY_PGA"
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemorySga SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "MEMORY_SGA"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetric = map[string]SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum{
	"CPU":        SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu,
	"STORAGE":    SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricStorage,
	"IO":         SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricIo,
	"MEMORY":     SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory,
	"MEMORY_PGA": SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemoryPga,
	"MEMORY_SGA": SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricMemorySga,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendAggregationCollectionResourceMetric {
		values = append(values, v)
	}
	return values
}
