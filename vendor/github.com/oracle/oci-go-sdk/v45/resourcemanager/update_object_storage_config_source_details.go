// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// UpdateObjectStorageConfigSourceDetails Updates property details for the Object Storage bucket that contains Terraform configuration files.
type UpdateObjectStorageConfigSourceDetails struct {

	// The path of the directory from which to run terraform. If not specified, the the root will be used. This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// The name of the bucket's region.
	// Example: `PHX`
	Region *string `mandatory:"false" json:"region"`

	// The Object Storage namespace that contains the bucket.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The name of the bucket that contains the Terraform configuration files.
	BucketName *string `mandatory:"false" json:"bucketName"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m UpdateObjectStorageConfigSourceDetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m UpdateObjectStorageConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateObjectStorageConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateObjectStorageConfigSourceDetails UpdateObjectStorageConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeUpdateObjectStorageConfigSourceDetails
	}{
		"OBJECT_STORAGE_CONFIG_SOURCE",
		(MarshalTypeUpdateObjectStorageConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
