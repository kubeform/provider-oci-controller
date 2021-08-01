// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Management Agent API
//
// API for Management Agent Cloud Service
//

package managementagent

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// AvailabilityHistorySummary Availability history of Management Agent.
type AvailabilityHistorySummary struct {

	// agent identifier
	ManagementAgentId *string `mandatory:"true" json:"managementAgentId"`

	// The availability status of managementAgent
	AvailabilityStatus AvailabilityStatusEnum `mandatory:"true" json:"availabilityStatus"`

	// The time at which the Management Agent moved to the availability status. An RFC3339 formatted datetime string
	TimeAvailabilityStatusStarted *common.SDKTime `mandatory:"false" json:"timeAvailabilityStatusStarted"`

	// The time till which the Management Agent was known to be in the availability status. An RFC3339 formatted datetime string
	TimeAvailabilityStatusEnded *common.SDKTime `mandatory:"false" json:"timeAvailabilityStatusEnded"`
}

func (m AvailabilityHistorySummary) String() string {
	return common.PointerString(m)
}
