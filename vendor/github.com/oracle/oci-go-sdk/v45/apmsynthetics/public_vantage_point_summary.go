// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// PublicVantagePointSummary Information about public vantage points.
type PublicVantagePointSummary struct {

	// Unique name that can be edited. The name should not contain any confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique permanent name of the vantage point.
	Name *string `mandatory:"true" json:"name"`

	Geo *GeoSummary `mandatory:"false" json:"geo"`
}

func (m PublicVantagePointSummary) String() string {
	return common.PointerString(m)
}
