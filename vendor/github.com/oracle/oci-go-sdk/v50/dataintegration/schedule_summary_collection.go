// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ScheduleSummaryCollection A collection of schedule summaries. The collection can be lightweight details or full definitions.
type ScheduleSummaryCollection struct {

	// An array of Schedule summaries
	Items []ScheduleSummary `mandatory:"true" json:"items"`
}

func (m ScheduleSummaryCollection) String() string {
	return common.PointerString(m)
}
