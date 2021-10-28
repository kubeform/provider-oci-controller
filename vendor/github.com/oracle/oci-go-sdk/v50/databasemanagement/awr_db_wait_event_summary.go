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
	"github.com/oracle/oci-go-sdk/v50/common"
)

// AwrDbWaitEventSummary The summary of the AWR wait event time series data for one event.
type AwrDbWaitEventSummary struct {

	// The name of the event.
	Name *string `mandatory:"true" json:"name"`

	// The begin time of the wait event.
	TimeBegin *common.SDKTime `mandatory:"false" json:"timeBegin"`

	// The end time of the wait event.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The wait count per second.
	WaitsPerSec *float64 `mandatory:"false" json:"waitsPerSec"`

	// The average wait time per second.
	AvgWaitTimePerSec *float64 `mandatory:"false" json:"avgWaitTimePerSec"`

	// The ID of the snapshot. The snapshot ID is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbSnapshots
	SnapshotId *int `mandatory:"false" json:"snapshotId"`
}

func (m AwrDbWaitEventSummary) String() string {
	return common.PointerString(m)
}
