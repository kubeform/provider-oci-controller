// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ObjectStorageConfigSource Metadata about the Object Storage configuration source.
type ObjectStorageConfigSource struct {

	// The name of the bucket's region.
	// Example: `PHX`
	Region *string `mandatory:"true" json:"region"`

	// The Object Storage namespace that contains the bucket.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the bucket that contains the Terraform configuration files.
	// Maximum file size (applies to each file in the bucket): 100 MB. (In a bucket, a file is an object.)
	BucketName *string `mandatory:"true" json:"bucketName"`

	// File path to the directory to use for running Terraform.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m ObjectStorageConfigSource) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m ObjectStorageConfigSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ObjectStorageConfigSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeObjectStorageConfigSource ObjectStorageConfigSource
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeObjectStorageConfigSource
	}{
		"OBJECT_STORAGE_CONFIG_SOURCE",
		(MarshalTypeObjectStorageConfigSource)(m),
	}

	return json.Marshal(&s)
}
