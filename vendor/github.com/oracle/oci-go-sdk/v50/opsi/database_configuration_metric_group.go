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

// DatabaseConfigurationMetricGroup Supported configuration metric groups for database capacity planning service.
type DatabaseConfigurationMetricGroup interface {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	GetTimeCollected() *common.SDKTime
}

type databaseconfigurationmetricgroup struct {
	JsonData      []byte
	TimeCollected *common.SDKTime `mandatory:"false" json:"timeCollected"`
	MetricName    string          `json:"metricName"`
}

// UnmarshalJSON unmarshals json
func (m *databaseconfigurationmetricgroup) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabaseconfigurationmetricgroup databaseconfigurationmetricgroup
	s := struct {
		Model Unmarshalerdatabaseconfigurationmetricgroup
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
func (m *databaseconfigurationmetricgroup) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricName {
	case "DB_OS_CONFIG_INSTANCE":
		mm := DbosConfigInstance{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_EXTERNAL_INSTANCE":
		mm := DbExternalInstance{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DB_EXTERNAL_PROPERTIES":
		mm := DbExternalProperties{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetTimeCollected returns TimeCollected
func (m databaseconfigurationmetricgroup) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m databaseconfigurationmetricgroup) String() string {
	return common.PointerString(m)
}

// DatabaseConfigurationMetricGroupMetricNameEnum Enum with underlying type: string
type DatabaseConfigurationMetricGroupMetricNameEnum string

// Set of constants representing the allowable values for DatabaseConfigurationMetricGroupMetricNameEnum
const (
	DatabaseConfigurationMetricGroupMetricNameExternalProperties DatabaseConfigurationMetricGroupMetricNameEnum = "DB_EXTERNAL_PROPERTIES"
	DatabaseConfigurationMetricGroupMetricNameExternalInstance   DatabaseConfigurationMetricGroupMetricNameEnum = "DB_EXTERNAL_INSTANCE"
	DatabaseConfigurationMetricGroupMetricNameOsConfigInstance   DatabaseConfigurationMetricGroupMetricNameEnum = "DB_OS_CONFIG_INSTANCE"
)

var mappingDatabaseConfigurationMetricGroupMetricName = map[string]DatabaseConfigurationMetricGroupMetricNameEnum{
	"DB_EXTERNAL_PROPERTIES": DatabaseConfigurationMetricGroupMetricNameExternalProperties,
	"DB_EXTERNAL_INSTANCE":   DatabaseConfigurationMetricGroupMetricNameExternalInstance,
	"DB_OS_CONFIG_INSTANCE":  DatabaseConfigurationMetricGroupMetricNameOsConfigInstance,
}

// GetDatabaseConfigurationMetricGroupMetricNameEnumValues Enumerates the set of values for DatabaseConfigurationMetricGroupMetricNameEnum
func GetDatabaseConfigurationMetricGroupMetricNameEnumValues() []DatabaseConfigurationMetricGroupMetricNameEnum {
	values := make([]DatabaseConfigurationMetricGroupMetricNameEnum, 0)
	for _, v := range mappingDatabaseConfigurationMetricGroupMetricName {
		values = append(values, v)
	}
	return values
}
