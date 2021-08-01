// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// APIs for managing Cloud Advisor. Cloud Advisor provides recommendations that help you maximize cost savings and improve the security posture of your tenancy.
//

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// CreateProfileDetails Details for creating a profile.
type CreateProfileDetails struct {

	// The OCID of the tenancy. The tenancy is the root compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The name assigned to the profile. Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// Text describing the profile. Avoid entering confidential information.
	Description *string `mandatory:"true" json:"description"`

	LevelsConfiguration *LevelsConfiguration `mandatory:"true" json:"levelsConfiguration"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair applied without any predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	TargetCompartments *TargetCompartments `mandatory:"false" json:"targetCompartments"`

	TargetTags *TargetTags `mandatory:"false" json:"targetTags"`
}

func (m CreateProfileDetails) String() string {
	return common.PointerString(m)
}
