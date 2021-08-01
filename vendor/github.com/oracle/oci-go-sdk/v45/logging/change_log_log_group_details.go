// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// ChangeLogLogGroupDetails Contains details indicating which log group the log should move to.
type ChangeLogLogGroupDetails struct {

	// Log group OCID.
	TargetLogGroupId *string `mandatory:"false" json:"targetLogGroupId"`
}

func (m ChangeLogLogGroupDetails) String() string {
	return common.PointerString(m)
}
