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
	"github.com/oracle/oci-go-sdk/v50/common"
)

// HostConfigurationMetricGroup Base Metric Group for Host configuration metrics
type HostConfigurationMetricGroup interface {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	GetTimeCollected() *common.SDKTime
}

type hostconfigurationmetricgroup struct {
	JsonData      []byte
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`
	MetricName    string          `json:"metricName"`
}

// UnmarshalJSON unmarshals json
func (m *hostconfigurationmetricgroup) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerhostconfigurationmetricgroup hostconfigurationmetricgroup
	s := struct {
		Model Unmarshalerhostconfigurationmetricgroup
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
func (m *hostconfigurationmetricgroup) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricName {
	case "HOST_RESOURCE_ALLOCATION":
		mm := HostResourceAllocation{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_PRODUCT":
		mm := HostProduct{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_NETWORK_CONFIGURATION":
		mm := HostNetworkConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_ENTITIES":
		mm := HostEntities{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_MEMORY_CONFIGURATION":
		mm := HostMemoryConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_CPU_HARDWARE_CONFIGURATION":
		mm := HostCpuHardwareConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HOST_HARDWARE_CONFIGURATION":
		mm := HostHardwareConfiguration{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetTimeCollected returns TimeCollected
func (m hostconfigurationmetricgroup) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m hostconfigurationmetricgroup) String() string {
	return common.PointerString(m)
}

// HostConfigurationMetricGroupMetricNameEnum Enum with underlying type: string
type HostConfigurationMetricGroupMetricNameEnum string

// Set of constants representing the allowable values for HostConfigurationMetricGroupMetricNameEnum
const (
	HostConfigurationMetricGroupMetricNameProduct                  HostConfigurationMetricGroupMetricNameEnum = "HOST_PRODUCT"
	HostConfigurationMetricGroupMetricNameResourceAllocation       HostConfigurationMetricGroupMetricNameEnum = "HOST_RESOURCE_ALLOCATION"
	HostConfigurationMetricGroupMetricNameMemoryConfiguration      HostConfigurationMetricGroupMetricNameEnum = "HOST_MEMORY_CONFIGURATION"
	HostConfigurationMetricGroupMetricNameHardwareConfiguration    HostConfigurationMetricGroupMetricNameEnum = "HOST_HARDWARE_CONFIGURATION"
	HostConfigurationMetricGroupMetricNameCpuHardwareConfiguration HostConfigurationMetricGroupMetricNameEnum = "HOST_CPU_HARDWARE_CONFIGURATION"
	HostConfigurationMetricGroupMetricNameNetworkConfiguration     HostConfigurationMetricGroupMetricNameEnum = "HOST_NETWORK_CONFIGURATION"
	HostConfigurationMetricGroupMetricNameEntites                  HostConfigurationMetricGroupMetricNameEnum = "HOST_ENTITES"
)

var mappingHostConfigurationMetricGroupMetricName = map[string]HostConfigurationMetricGroupMetricNameEnum{
	"HOST_PRODUCT":                    HostConfigurationMetricGroupMetricNameProduct,
	"HOST_RESOURCE_ALLOCATION":        HostConfigurationMetricGroupMetricNameResourceAllocation,
	"HOST_MEMORY_CONFIGURATION":       HostConfigurationMetricGroupMetricNameMemoryConfiguration,
	"HOST_HARDWARE_CONFIGURATION":     HostConfigurationMetricGroupMetricNameHardwareConfiguration,
	"HOST_CPU_HARDWARE_CONFIGURATION": HostConfigurationMetricGroupMetricNameCpuHardwareConfiguration,
	"HOST_NETWORK_CONFIGURATION":      HostConfigurationMetricGroupMetricNameNetworkConfiguration,
	"HOST_ENTITES":                    HostConfigurationMetricGroupMetricNameEntites,
}

// GetHostConfigurationMetricGroupMetricNameEnumValues Enumerates the set of values for HostConfigurationMetricGroupMetricNameEnum
func GetHostConfigurationMetricGroupMetricNameEnumValues() []HostConfigurationMetricGroupMetricNameEnum {
	values := make([]HostConfigurationMetricGroupMetricNameEnum, 0)
	for _, v := range mappingHostConfigurationMetricGroupMetricName {
		values = append(values, v)
	}
	return values
}
