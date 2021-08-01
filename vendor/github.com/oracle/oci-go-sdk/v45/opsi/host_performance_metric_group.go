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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// HostPerformanceMetricGroup Base Metric Group for Host performance metrics
type HostPerformanceMetricGroup interface {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	GetTimeCollected() *common.SDKTime
}

type hostperformancemetricgroup struct {
	JsonData      []byte
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`
	MetricName    string          `json:"metricName"`
}

// UnmarshalJSON unmarshals json
func (m *hostperformancemetricgroup) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhostperformancemetricgroup hostperformancemetricgroup
	s := struct {
		Model Unmarshalerhostperformancemetricgroup
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.TimeCollected = s.Model.TimeCollected
	m.MetricName = s.Model.MetricName

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *hostperformancemetricgroup) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricName {
	case "HOST_MEMORY_USAGE":
		mm := HostMemoryUsage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_CPU_USAGE":
		mm := HostCpuUsage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_NETWORK_ACTIVITY_SUMMARY":
		mm := HostNetworkActivitySummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetTimeCollected returns TimeCollected
func (m hostperformancemetricgroup) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m hostperformancemetricgroup) String() string {
	return common.PointerString(m)
}

// HostPerformanceMetricGroupMetricNameEnum Enum with underlying type: string
type HostPerformanceMetricGroupMetricNameEnum string

// Set of constants representing the allowable values for HostPerformanceMetricGroupMetricNameEnum
const (
	HostPerformanceMetricGroupMetricNameCpuUsage               HostPerformanceMetricGroupMetricNameEnum = "HOST_CPU_USAGE"
	HostPerformanceMetricGroupMetricNameMemoryUsage            HostPerformanceMetricGroupMetricNameEnum = "HOST_MEMORY_USAGE"
	HostPerformanceMetricGroupMetricNameNetworkActivitySummary HostPerformanceMetricGroupMetricNameEnum = "HOST_NETWORK_ACTIVITY_SUMMARY"
)

var mappingHostPerformanceMetricGroupMetricName = map[string]HostPerformanceMetricGroupMetricNameEnum{
	"HOST_CPU_USAGE":                HostPerformanceMetricGroupMetricNameCpuUsage,
	"HOST_MEMORY_USAGE":             HostPerformanceMetricGroupMetricNameMemoryUsage,
	"HOST_NETWORK_ACTIVITY_SUMMARY": HostPerformanceMetricGroupMetricNameNetworkActivitySummary,
}

// GetHostPerformanceMetricGroupMetricNameEnumValues Enumerates the set of values for HostPerformanceMetricGroupMetricNameEnum
func GetHostPerformanceMetricGroupMetricNameEnumValues() []HostPerformanceMetricGroupMetricNameEnum {
	values := make([]HostPerformanceMetricGroupMetricNameEnum, 0)
	for _, v := range mappingHostPerformanceMetricGroupMetricName {
		values = append(values, v)
	}
	return values
}
