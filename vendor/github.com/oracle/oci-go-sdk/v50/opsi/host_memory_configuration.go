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

// HostMemoryConfiguration Memory Configuration metric for the host
type HostMemoryConfiguration struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Page size in kilobytes
	PageSizeInKB *float64 `mandatory:"false" json:"pageSizeInKB"`

	// Amount of memory used for page tables in kilobytes
	PageTablesInKB *float64 `mandatory:"false" json:"pageTablesInKB"`

	// Amount of total swap space in kilobytes
	SwapTotalInKB *float64 `mandatory:"false" json:"swapTotalInKB"`

	// Size of huge pages in kilobytes
	HugePageSizeInKB *float64 `mandatory:"false" json:"hugePageSizeInKB"`

	// Total number of huge pages
	HugePagesTotal *int `mandatory:"false" json:"hugePagesTotal"`
}

//GetTimeCollected returns TimeCollected
func (m HostMemoryConfiguration) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostMemoryConfiguration) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HostMemoryConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostMemoryConfiguration HostMemoryConfiguration
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostMemoryConfiguration
	}{
		"HOST_MEMORY_CONFIGURATION",
		(MarshalTypeHostMemoryConfiguration)(m),
	}

	return json.Marshal(&s)
}
