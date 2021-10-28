// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataLabelingService API
//
// A description of the DataLabelingService API
//

package datalabelingservice

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// InitialRecordGenerationConfiguration Initial Generate Records configuration, generates records from the Dataset's source.
type InitialRecordGenerationConfiguration struct {

	// the maximum number of records to generate.
	Limit *float32 `mandatory:"false" json:"limit"`
}

func (m InitialRecordGenerationConfiguration) String() string {
	return common.PointerString(m)
}
