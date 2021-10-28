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
	"github.com/oracle/oci-go-sdk/v50/common"
)

// KernelVmCoreInformation VMcore information.
type KernelVmCoreInformation struct {

	// Kernel module responsible of the crash.
	Component *string `mandatory:"false" json:"component"`

	// Crash backtrace.
	Backtrace *string `mandatory:"false" json:"backtrace"`
}

func (m KernelVmCoreInformation) String() string {
	return common.PointerString(m)
}
