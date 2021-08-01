// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Java Management Service Query API
//
// API for the Java Management Service. Use this API to view and manage Fleets.
//

package jms

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// ChangeFleetCompartmentDetails Attributes to change the compartment of a Fleet.
type ChangeFleetCompartmentDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment into which the Fleet should be moved.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`
}

func (m ChangeFleetCompartmentDetails) String() string {
	return common.PointerString(m)
}
