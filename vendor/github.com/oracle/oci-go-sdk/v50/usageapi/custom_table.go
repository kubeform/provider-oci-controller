// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CustomTable The saved custom table.
type CustomTable struct {

	// The custom table OCID.
	Id *string `mandatory:"true" json:"id"`

	// The custom table associated saved report OCID.
	SavedReportId *string `mandatory:"false" json:"savedReportId"`

	// The custom table compartment OCID.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	SavedCustomTable *SavedCustomTable `mandatory:"false" json:"savedCustomTable"`
}

func (m CustomTable) String() string {
	return common.PointerString(m)
}
