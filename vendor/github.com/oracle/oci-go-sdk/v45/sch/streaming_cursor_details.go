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
	"github.com/oracle/oci-go-sdk/v45/common"
)

// StreamingCursorDetails The type of cursor (https://docs.cloud.oracle.com/iaas/Content/Streaming/Tasks/using_a_single_consumer.htm#usingcursors), which determines the starting point from which the stream will be consumed.
type StreamingCursorDetails interface {
}

type streamingcursordetails struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *streamingcursordetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerstreamingcursordetails streamingcursordetails
	s := struct {
		Model Unmarshalerstreamingcursordetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *streamingcursordetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "TRIM_HORIZON":
		mm := TrimHorizonStreamingCursor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LATEST":
		mm := LatestStreamingCursor{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m streamingcursordetails) String() string {
	return common.PointerString(m)
}

// StreamingCursorDetailsKindEnum Enum with underlying type: string
type StreamingCursorDetailsKindEnum string

// Set of constants representing the allowable values for StreamingCursorDetailsKindEnum
const (
	StreamingCursorDetailsKindLatest      StreamingCursorDetailsKindEnum = "LATEST"
	StreamingCursorDetailsKindTrimHorizon StreamingCursorDetailsKindEnum = "TRIM_HORIZON"
)

var mappingStreamingCursorDetailsKind = map[string]StreamingCursorDetailsKindEnum{
	"LATEST":       StreamingCursorDetailsKindLatest,
	"TRIM_HORIZON": StreamingCursorDetailsKindTrimHorizon,
}

// GetStreamingCursorDetailsKindEnumValues Enumerates the set of values for StreamingCursorDetailsKindEnum
func GetStreamingCursorDetailsKindEnumValues() []StreamingCursorDetailsKindEnum {
	values := make([]StreamingCursorDetailsKindEnum, 0)
	for _, v := range mappingStreamingCursorDetailsKind {
		values = append(values, v)
	}
	return values
}
