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

// ConfigSource Information about the Terraform configuration.
type ConfigSource interface {

	// File path to the directory to use for running Terraform.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	GetWorkingDirectory() *string
}

type configsource struct {
	JsonData         []byte
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`
	ConfigSourceType string  `json:"configSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *configsource) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigsource configsource
	s := struct {
		Model Unmarshalerconfigsource
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.WorkingDirectory = s.Model.WorkingDirectory
	m.ConfigSourceType = s.Model.ConfigSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configsource) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConfigSourceType {
	case "GIT_CONFIG_SOURCE":
		mm := GitConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OBJECT_STORAGE_CONFIG_SOURCE":
		mm := ObjectStorageConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPARTMENT_CONFIG_SOURCE":
		mm := CompartmentConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ZIP_UPLOAD":
		mm := ZipUploadConfigSource{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetWorkingDirectory returns WorkingDirectory
func (m configsource) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m configsource) String() string {
	return common.PointerString(m)
}

// ConfigSourceConfigSourceTypeEnum Enum with underlying type: string
type ConfigSourceConfigSourceTypeEnum string

// Set of constants representing the allowable values for ConfigSourceConfigSourceTypeEnum
const (
	ConfigSourceConfigSourceTypeZipUpload                 ConfigSourceConfigSourceTypeEnum = "ZIP_UPLOAD"
	ConfigSourceConfigSourceTypeGitConfigSource           ConfigSourceConfigSourceTypeEnum = "GIT_CONFIG_SOURCE"
	ConfigSourceConfigSourceTypeCompartmentConfigSource   ConfigSourceConfigSourceTypeEnum = "COMPARTMENT_CONFIG_SOURCE"
	ConfigSourceConfigSourceTypeObjectStorageConfigSource ConfigSourceConfigSourceTypeEnum = "OBJECT_STORAGE_CONFIG_SOURCE"
)

var mappingConfigSourceConfigSourceType = map[string]ConfigSourceConfigSourceTypeEnum{
	"ZIP_UPLOAD":                   ConfigSourceConfigSourceTypeZipUpload,
	"GIT_CONFIG_SOURCE":            ConfigSourceConfigSourceTypeGitConfigSource,
	"COMPARTMENT_CONFIG_SOURCE":    ConfigSourceConfigSourceTypeCompartmentConfigSource,
	"OBJECT_STORAGE_CONFIG_SOURCE": ConfigSourceConfigSourceTypeObjectStorageConfigSource,
}

// GetConfigSourceConfigSourceTypeEnumValues Enumerates the set of values for ConfigSourceConfigSourceTypeEnum
func GetConfigSourceConfigSourceTypeEnumValues() []ConfigSourceConfigSourceTypeEnum {
	values := make([]ConfigSourceConfigSourceTypeEnum, 0)
	for _, v := range mappingConfigSourceConfigSourceType {
		values = append(values, v)
	}
	return values
}
