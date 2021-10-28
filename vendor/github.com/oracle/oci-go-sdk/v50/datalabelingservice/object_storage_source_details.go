// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataLabelingService API
//
// A description of the DataLabelingService API
//

package datalabelingservice

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ObjectStorageSourceDetails Specifies the dataset location in object storage. This requires that all records are in this bucket, and under this prefix. We do not support a dataset with objects in arbitrary locations across buckets or prefixes.
type ObjectStorageSourceDetails struct {

	// Namespace of the bucket that contains the dataset data source
	Namespace *string `mandatory:"true" json:"namespace"`

	// The object storage bucket that contains the dataset data source
	Bucket *string `mandatory:"true" json:"bucket"`

	// A common path prefix shared by the objects that make up the dataset.
	Prefix *string `mandatory:"false" json:"prefix"`
}

func (m ObjectStorageSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageSourceDetails ObjectStorageSourceDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeObjectStorageSourceDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeObjectStorageSourceDetails)(m),
	}

	return json.Marshal(&s)
}
