// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// WorkRequestResourceSubTypeDetail SubType information for a work request resource.
type WorkRequestResourceSubTypeDetail struct {

	// Subtype of the work request resource like osn or peer.
	SubType *string `mandatory:"true" json:"subType"`

	// The identifier of the resource subType.
	SubTypeKey *string `mandatory:"true" json:"subTypeKey"`

	// Status of the resource subType, as a result of the work tracked in this work request.
	// A resource subType would be CREATED, UPDATED or DELETED, after the work request is completed.
	SubTypeStatus WorkRequestResourceSubTypeDetailSubTypeStatusEnum `mandatory:"true" json:"subTypeStatus"`
}

func (m WorkRequestResourceSubTypeDetail) String() string {
	return common.PointerString(m)
}

// WorkRequestResourceSubTypeDetailSubTypeStatusEnum Enum with underlying type: string
type WorkRequestResourceSubTypeDetailSubTypeStatusEnum string

// Set of constants representing the allowable values for WorkRequestResourceSubTypeDetailSubTypeStatusEnum
const (
	WorkRequestResourceSubTypeDetailSubTypeStatusCreated WorkRequestResourceSubTypeDetailSubTypeStatusEnum = "CREATED"
	WorkRequestResourceSubTypeDetailSubTypeStatusUpdated WorkRequestResourceSubTypeDetailSubTypeStatusEnum = "UPDATED"
	WorkRequestResourceSubTypeDetailSubTypeStatusDeleted WorkRequestResourceSubTypeDetailSubTypeStatusEnum = "DELETED"
)

var mappingWorkRequestResourceSubTypeDetailSubTypeStatus = map[string]WorkRequestResourceSubTypeDetailSubTypeStatusEnum{
	"CREATED": WorkRequestResourceSubTypeDetailSubTypeStatusCreated,
	"UPDATED": WorkRequestResourceSubTypeDetailSubTypeStatusUpdated,
	"DELETED": WorkRequestResourceSubTypeDetailSubTypeStatusDeleted,
}

// GetWorkRequestResourceSubTypeDetailSubTypeStatusEnumValues Enumerates the set of values for WorkRequestResourceSubTypeDetailSubTypeStatusEnum
func GetWorkRequestResourceSubTypeDetailSubTypeStatusEnumValues() []WorkRequestResourceSubTypeDetailSubTypeStatusEnum {
	values := make([]WorkRequestResourceSubTypeDetailSubTypeStatusEnum, 0)
	for _, v := range mappingWorkRequestResourceSubTypeDetailSubTypeStatus {
		values = append(values, v)
	}
	return values
}
