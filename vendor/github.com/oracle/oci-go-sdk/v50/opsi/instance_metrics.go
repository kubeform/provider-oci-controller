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
	"github.com/oracle/oci-go-sdk/v50/common"
)

// InstanceMetrics Object containing instance metrics.
type InstanceMetrics struct {

	// The hostname of the database insight resource.
	HostName *string `mandatory:"false" json:"hostName"`

	// The instance name of the database insight resource.
	InstanceName *string `mandatory:"false" json:"instanceName"`

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"false" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE).
	Capacity *float64 `mandatory:"false" json:"capacity"`

	// Resource utilization in percentage
	UtilizationPercent *float64 `mandatory:"false" json:"utilizationPercent"`

	// Change in resource utilization in percentage
	UsageChangePercent *float64 `mandatory:"false" json:"usageChangePercent"`
}

func (m InstanceMetrics) String() string {
	return common.PointerString(m)
}
