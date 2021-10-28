// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// PatternCollection Results of a patterns listing.  A Pattern is defined using an expression and can be used as data selectors or filters
// to provide a singular view of an entity across multiple physical data artifacts.
type PatternCollection struct {

	// Collection of pattern summaries.
	Items []PatternSummary `mandatory:"true" json:"items"`

	// Total number of items returned.
	Count *int `mandatory:"false" json:"count"`
}

func (m PatternCollection) String() string {
	return common.PointerString(m)
}
