// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// ExternalPublicationValidation The information about external published task validation.
type ExternalPublicationValidation struct {

	// Total number of validation messages.
	TotalMessageCount *int `mandatory:"false" json:"totalMessageCount"`

	// Total number of validation error messages.
	ErrorMessageCount *int `mandatory:"false" json:"errorMessageCount"`

	// Total number of validation warning messages.
	WarnMessageCount *int `mandatory:"false" json:"warnMessageCount"`

	// Total number of validation information messages.
	InfoMessageCount *int `mandatory:"false" json:"infoMessageCount"`

	// Detailed information of the data flow object validation.
	ValidationMessages map[string][]ValidationMessage `mandatory:"false" json:"validationMessages"`

	// Objects use a 36 character key as unique ID. It is system generated and cannot be modified.
	Key *string `mandatory:"false" json:"key"`
}

func (m ExternalPublicationValidation) String() string {
	return common.PointerString(m)
}
