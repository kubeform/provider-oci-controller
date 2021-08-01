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

// ManagedDatabaseGroupSummary A group of Managed Databases that will be managed together.
type ManagedDatabaseGroupSummary struct {

	// The name of the Managed Database Group.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of Managed Databases in the Managed Database Group.
	ManagedDatabaseCount *int `mandatory:"true" json:"managedDatabaseCount"`

	// The current lifecycle state of the Managed Database Group.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Managed Database Group was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The information specified by the user about the Managed Database Group.
	Description *string `mandatory:"false" json:"description"`
}

func (m ManagedDatabaseGroupSummary) String() string {
	return common.PointerString(m)
}
