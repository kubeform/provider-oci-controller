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

// AwrDbSnapshotRangeSummary The summary data for a range of AWR snapshots.
type AwrDbSnapshotRangeSummary struct {

	// The internal ID of the database. The internal ID of the database is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs
	AwrDbId *string `mandatory:"true" json:"awrDbId"`

	// The name of the database.
	DbName *string `mandatory:"true" json:"dbName"`

	// The database instance numbers.
	InstanceList []int `mandatory:"false" json:"instanceList"`

	// The timestamp of the database startup.
	TimeDbStartup *common.SDKTime `mandatory:"false" json:"timeDbStartup"`

	// The start time of the earliest snapshot.
	TimeFirstSnapshotBegin *common.SDKTime `mandatory:"false" json:"timeFirstSnapshotBegin"`

	// The end time of the latest snapshot.
	TimeLatestSnapshotEnd *common.SDKTime `mandatory:"false" json:"timeLatestSnapshotEnd"`

	// The ID of the earliest snapshot. The snapshot ID is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbSnapshots
	FirstSnapshotId *int `mandatory:"false" json:"firstSnapshotId"`

	// The ID of the latest snapshot. The snapshot ID is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs/{awrDbId}/awrDbSnapshots
	LatestSnapshotId *int `mandatory:"false" json:"latestSnapshotId"`

	// The total number of snapshots.
	SnapshotCount *int64 `mandatory:"false" json:"snapshotCount"`

	// The interval time between snapshots (in minutes).
	SnapshotIntervalInMin *int `mandatory:"false" json:"snapshotIntervalInMin"`

	// ID of the database container. The database container ID is not the OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	// It can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" json:"containerId"`

	// The version of the database.
	DbVersion *string `mandatory:"false" json:"dbVersion"`

	// The time zone of the snapshot.
	SnapshotTimezone *string `mandatory:"false" json:"snapshotTimezone"`
}

func (m AwrDbSnapshotRangeSummary) String() string {
	return common.PointerString(m)
}
