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

// DataSelectorPatternDetails List of pattern Ids.  Used in the addition and removal of patterns in data assets.
type DataSelectorPatternDetails struct {

	// Collection of pattern Ids.
	Items []string `mandatory:"true" json:"items"`
}

func (m DataSelectorPatternDetails) String() string {
	return common.PointerString(m)
}
