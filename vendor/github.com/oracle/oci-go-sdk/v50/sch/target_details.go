// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Connector Hub API
//
// Use the Service Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Service Connector Hub, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
//

package sch

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// TargetDetails An object that represents the target of the flow defined by the service connector.
// An example target is a stream.
// For more information about flows defined by service connectors, see
// Service Connector Hub Overview (https://docs.cloud.oracle.com/iaas/Content/service-connector-hub/overview.htm).
type TargetDetails interface {
}

type targetdetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *targetdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertargetdetails targetdetails
	s := struct {
		Model Unmarshalertargetdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *targetdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "notifications":
		mm := NotificationsTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "objectStorage":
		mm := ObjectStorageTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "monitoring":
		mm := MonitoringTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "functions":
		mm := FunctionsTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "loggingAnalytics":
		mm := LoggingAnalyticsTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "streaming":
		mm := StreamingTargetDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m targetdetails) String() string {
	return common.PointerString(m)
}

// TargetDetailsKindEnum Enum with underlying type: string
type TargetDetailsKindEnum string

// Set of constants representing the allowable values for TargetDetailsKindEnum
const (
	TargetDetailsKindFunctions        TargetDetailsKindEnum = "functions"
	TargetDetailsKindLogginganalytics TargetDetailsKindEnum = "loggingAnalytics"
	TargetDetailsKindMonitoring       TargetDetailsKindEnum = "monitoring"
	TargetDetailsKindNotifications    TargetDetailsKindEnum = "notifications"
	TargetDetailsKindObjectstorage    TargetDetailsKindEnum = "objectStorage"
	TargetDetailsKindStreaming        TargetDetailsKindEnum = "streaming"
)

var mappingTargetDetailsKind = map[string]TargetDetailsKindEnum{
	"functions":        TargetDetailsKindFunctions,
	"loggingAnalytics": TargetDetailsKindLogginganalytics,
	"monitoring":       TargetDetailsKindMonitoring,
	"notifications":    TargetDetailsKindNotifications,
	"objectStorage":    TargetDetailsKindObjectstorage,
	"streaming":        TargetDetailsKindStreaming,
}

// GetTargetDetailsKindEnumValues Enumerates the set of values for TargetDetailsKindEnum
func GetTargetDetailsKindEnumValues() []TargetDetailsKindEnum {
	values := make([]TargetDetailsKindEnum, 0)
	for _, v := range mappingTargetDetailsKind {
		values = append(values, v)
	}
	return values
}
