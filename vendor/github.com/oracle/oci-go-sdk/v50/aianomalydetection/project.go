// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// Project Project enable users to organize their resources.
type Project struct {

	// The OCID of the project that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID for the project's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the resource was created in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The lifecycle state of the Project.
	LifecycleState ProjectLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A short description of the project.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the resource was updated in the timestamp format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Project) String() string {
	return common.PointerString(m)
}

// ProjectLifecycleStateEnum Enum with underlying type: string
type ProjectLifecycleStateEnum string

// Set of constants representing the allowable values for ProjectLifecycleStateEnum
const (
	ProjectLifecycleStateActive   ProjectLifecycleStateEnum = "ACTIVE"
	ProjectLifecycleStateCreating ProjectLifecycleStateEnum = "CREATING"
	ProjectLifecycleStateFailed   ProjectLifecycleStateEnum = "FAILED"
	ProjectLifecycleStateDeleting ProjectLifecycleStateEnum = "DELETING"
	ProjectLifecycleStateDeleted  ProjectLifecycleStateEnum = "DELETED"
	ProjectLifecycleStateUpdating ProjectLifecycleStateEnum = "UPDATING"
)

var mappingProjectLifecycleState = map[string]ProjectLifecycleStateEnum{
	"ACTIVE":   ProjectLifecycleStateActive,
	"CREATING": ProjectLifecycleStateCreating,
	"FAILED":   ProjectLifecycleStateFailed,
	"DELETING": ProjectLifecycleStateDeleting,
	"DELETED":  ProjectLifecycleStateDeleted,
	"UPDATING": ProjectLifecycleStateUpdating,
}

// GetProjectLifecycleStateEnumValues Enumerates the set of values for ProjectLifecycleStateEnum
func GetProjectLifecycleStateEnumValues() []ProjectLifecycleStateEnum {
	values := make([]ProjectLifecycleStateEnum, 0)
	for _, v := range mappingProjectLifecycleState {
		values = append(values, v)
	}
	return values
}
