// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// BulkApplyResourceAction The resource action that a recommendation will be applied to.
type BulkApplyResourceAction struct {

	// The unique OCIDs of the resource actions that recommendations are applied to.
	ResourceActionId *string `mandatory:"true" json:"resourceActionId"`

	// The current status of the recommendation.
	Status StatusEnum `mandatory:"false" json:"status,omitempty"`

	// The date and time the current status will change. The format is defined by RFC3339.
	// For example, "The current `postponed` status of the resource action will end and change to `pending` on this
	// date and time."
	TimeStatusEnd *common.SDKTime `mandatory:"false" json:"timeStatusEnd"`

	// Additional parameter key-value pairs defining the resource action.
	// For example:
	// `{"timeAmount": 15, "timeUnit": "seconds"}`
	Parameters map[string]interface{} `mandatory:"false" json:"parameters"`

	// The name of the strategy.
	StrategyName *string `mandatory:"false" json:"strategyName"`
}

func (m BulkApplyResourceAction) String() string {
	return common.PointerString(m)
}
