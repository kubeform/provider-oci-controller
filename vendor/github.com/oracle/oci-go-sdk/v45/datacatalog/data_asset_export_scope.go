// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// DataAssetExportScope Scope of asset export, which consists of a container object (bucket, folder, schema, etc) within the asset,
// and types of child objects contained by that object to be included.
//    objectKey - Key of the container object to be exported. For example, key of schema_1.
//    exportTypeIds - Type key(s) of objects within the container object to be exported. For example, type key of table or view.
type DataAssetExportScope struct {

	// Unique key of the object selected for export.
	ObjectKey *string `mandatory:"false" json:"objectKey"`

	// Array of type keys selected for export.
	ExportTypeIds []string `mandatory:"false" json:"exportTypeIds"`
}

func (m DataAssetExportScope) String() string {
	return common.PointerString(m)
}
