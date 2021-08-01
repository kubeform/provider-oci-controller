// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// ExternalStorage BICC Connector Attribute.Object Storage as External storage where the BICC extracted files are written
type ExternalStorage struct {

	// Id of the external stoarge configured in BICC console. Usually its numeric.
	StorageId *string `mandatory:"false" json:"storageId"`

	// Name of the external storage configured in BICC console
	StorageName *string `mandatory:"false" json:"storageName"`

	// Object Storage host Url. DO not give http/https.
	Host *string `mandatory:"false" json:"host"`

	// Tenancy OCID for the OOS bucket
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// Namespace for the OOS bucket
	Namespace *string `mandatory:"false" json:"namespace"`

	// Bucket Name where BICC extracts stores the files
	Bucket *string `mandatory:"false" json:"bucket"`
}

func (m ExternalStorage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ExternalStorage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExternalStorage ExternalStorage
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeExternalStorage
	}{
		"EXTERNAL_STORAGE",
		(MarshalTypeExternalStorage)(m),
	}

	return json.Marshal(&s)
}
