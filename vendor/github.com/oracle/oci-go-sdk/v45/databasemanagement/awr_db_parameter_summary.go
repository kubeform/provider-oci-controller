// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// AwrDbParameterSummary The summary of the AWR change history data for a single database parameter.
type AwrDbParameterSummary struct {

	// The name of the parameter.
	Name *string `mandatory:"true" json:"name"`

	// The database instance number.
	InstanceNumber *int `mandatory:"false" json:"instanceNumber"`

	// The parameter value when the period began.
	BeginValue *string `mandatory:"false" json:"beginValue"`

	// The parameter value when the period ended.
	EndValue *string `mandatory:"false" json:"endValue"`

	// Indicates whether the parameter value changed within the period.
	IsChanged *bool `mandatory:"false" json:"isChanged"`

	// Indicates whether the parameter has been modified after instance startup:
	//  - MODIFIED - Parameter has been modified with ALTER SESSION
	//  - SYSTEM_MOD - Parameter has been modified with ALTER SYSTEM (which causes all the currently logged in sessions’ values to be modified)
	//  - FALSE - Parameter has not been modified after instance startup
	ValueModified *string `mandatory:"false" json:"valueModified"`

	// Indicates whether the parameter value in the end snapshot is the default.
	IsDefault *bool `mandatory:"false" json:"isDefault"`
}

func (m AwrDbParameterSummary) String() string {
	return common.PointerString(m)
}
