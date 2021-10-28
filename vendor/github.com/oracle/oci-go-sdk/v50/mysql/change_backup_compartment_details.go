// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ChangeBackupCompartmentDetails OCID of the target compartment for DB System Backup change compartment request.
type ChangeBackupCompartmentDetails struct {

	// OCID of the target compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeBackupCompartmentDetails) String() string {
	return common.PointerString(m)
}
