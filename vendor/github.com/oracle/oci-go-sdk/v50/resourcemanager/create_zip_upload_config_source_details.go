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

// CreateZipUploadConfigSourceDetails Property details for uploading the configuration zip file.
type CreateZipUploadConfigSourceDetails struct {
	ZipFileBase64Encoded *string `mandatory:"true" json:"zipFileBase64Encoded"`

	// File path to the directory from which Terraform runs.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m CreateZipUploadConfigSourceDetails) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CreateZipUploadConfigSourceDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateZipUploadConfigSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateZipUploadConfigSourceDetails CreateZipUploadConfigSourceDetails
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCreateZipUploadConfigSourceDetails
	}{
		"ZIP_UPLOAD",
		(MarshalTypeCreateZipUploadConfigSourceDetails)(m),
	}

	return json.Marshal(&s)
}
