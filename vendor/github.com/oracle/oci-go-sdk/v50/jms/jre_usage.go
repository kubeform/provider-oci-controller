// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service API
//
// API for the Java Management Service. Use this API to view, create, and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// JreUsage Java Runtime usage during a specified time period. A Java Runtime is identified by its vendor and version.
type JreUsage struct {

	// The vendor of the Java Runtime.
	Vendor *string `mandatory:"true" json:"vendor"`

	// The distribution of a Java Runtime is the name of the lineage of product to which it belongs, for example _Java(TM) SE Runtime Environment_.
	Distribution *string `mandatory:"true" json:"distribution"`

	// The version of the Java Runtime.
	Version *string `mandatory:"true" json:"version"`

	// The operating systems that have this Java Runtime installed.
	OperatingSystems []OperatingSystem `mandatory:"false" json:"operatingSystems"`

	// The approximate count of installations that are installations of this Java Runtime.
	ApproximateInstallationCount *int `mandatory:"false" json:"approximateInstallationCount"`

	// The approximate count of the applications running on this Java Runtime.
	ApproximateApplicationCount *int `mandatory:"false" json:"approximateApplicationCount"`

	// The approximate count of the managed instances that report this Java Runtime.
	ApproximateManagedInstanceCount *int `mandatory:"false" json:"approximateManagedInstanceCount"`

	// Lower bound of the specified time period filter.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Upper bound of the specified time period filter.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The date and time the resource was _first_ reported to JMS.
	// This is potentially _before_ the specified time period provided by the filters.
	// For example, a resource can be first reported to JMS before the start of a specified time period,
	// if it is also reported during the time period.
	TimeFirstSeen *common.SDKTime `mandatory:"false" json:"timeFirstSeen"`

	// The date and time the resource was _last_ reported to JMS.
	// This is potentially _after_ the specified time period provided by the filters.
	// For example, a resource can be last reported to JMS before the start of a specified time period,
	// if it is also reported during the time period.
	TimeLastSeen *common.SDKTime `mandatory:"false" json:"timeLastSeen"`
}

func (m JreUsage) String() string {
	return common.PointerString(m)
}
