// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// ManagedInstance Detail information for an OCI Compute instance that is being managed
type ManagedInstance struct {

	// Managed Instance identifier
	DisplayName *string `mandatory:"true" json:"displayName"`

	// OCID for the managed instance
	Id *string `mandatory:"true" json:"id"`

	// OCID for the Compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Information specified by the user about the managed instance
	Description *string `mandatory:"false" json:"description"`

	// Time at which the instance last checked in
	LastCheckin *string `mandatory:"false" json:"lastCheckin"`

	// Time at which the instance last booted
	LastBoot *string `mandatory:"false" json:"lastBoot"`

	// Number of updates available to be installed
	UpdatesAvailable *int `mandatory:"false" json:"updatesAvailable"`

	// Operating System Name
	OsName *string `mandatory:"false" json:"osName"`

	// Operating System Version
	OsVersion *string `mandatory:"false" json:"osVersion"`

	// Operating System Kernel Version
	OsKernelVersion *string `mandatory:"false" json:"osKernelVersion"`

	// status of the managed instance.
	Status ManagedInstanceStatusEnum `mandatory:"false" json:"status,omitempty"`

	// the parent (base) Software Source attached to the Managed Instance
	ParentSoftwareSource *SoftwareSourceId `mandatory:"false" json:"parentSoftwareSource"`

	// list of child Software Sources attached to the Managed Instance
	ChildSoftwareSources []SoftwareSourceId `mandatory:"false" json:"childSoftwareSources"`

	// The ids of the managed instance groups of which this instance is a
	// member.
	ManagedInstanceGroups []Id `mandatory:"false" json:"managedInstanceGroups"`

	// The Operating System type of the managed instance.
	OsFamily OsFamiliesEnum `mandatory:"false" json:"osFamily,omitempty"`

	// Indicates whether a reboot is required to complete installation of updates.
	IsRebootRequired *bool `mandatory:"false" json:"isRebootRequired"`

	// Number of security type updates available to be installed
	SecurityUpdatesAvailable *int `mandatory:"false" json:"securityUpdatesAvailable"`

	// Number of bug fix type updates available to be installed
	BugUpdatesAvailable *int `mandatory:"false" json:"bugUpdatesAvailable"`

	// Number of enhancement type updates available to be installed
	EnhancementUpdatesAvailable *int `mandatory:"false" json:"enhancementUpdatesAvailable"`

	// Number of non-classified updates available to be installed
	OtherUpdatesAvailable *int `mandatory:"false" json:"otherUpdatesAvailable"`

	// Number of scheduled jobs associated with this instance
	ScheduledJobCount *int `mandatory:"false" json:"scheduledJobCount"`

	// Number of work requests associated with this instance
	WorkRequestCount *int `mandatory:"false" json:"workRequestCount"`
}

func (m ManagedInstance) String() string {
	return common.PointerString(m)
}

// ManagedInstanceStatusEnum Enum with underlying type: string
type ManagedInstanceStatusEnum string

// Set of constants representing the allowable values for ManagedInstanceStatusEnum
const (
	ManagedInstanceStatusNormal      ManagedInstanceStatusEnum = "NORMAL"
	ManagedInstanceStatusUnreachable ManagedInstanceStatusEnum = "UNREACHABLE"
	ManagedInstanceStatusError       ManagedInstanceStatusEnum = "ERROR"
	ManagedInstanceStatusWarning     ManagedInstanceStatusEnum = "WARNING"
)

var mappingManagedInstanceStatus = map[string]ManagedInstanceStatusEnum{
	"NORMAL":      ManagedInstanceStatusNormal,
	"UNREACHABLE": ManagedInstanceStatusUnreachable,
	"ERROR":       ManagedInstanceStatusError,
	"WARNING":     ManagedInstanceStatusWarning,
}

// GetManagedInstanceStatusEnumValues Enumerates the set of values for ManagedInstanceStatusEnum
func GetManagedInstanceStatusEnumValues() []ManagedInstanceStatusEnum {
	values := make([]ManagedInstanceStatusEnum, 0)
	for _, v := range mappingManagedInstanceStatus {
		values = append(values, v)
	}
	return values
}
