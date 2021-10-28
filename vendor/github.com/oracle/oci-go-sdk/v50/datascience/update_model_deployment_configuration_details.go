// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// UpdateModelDeploymentConfigurationDetails The model deployment configuration details for update.
type UpdateModelDeploymentConfigurationDetails interface {
}

type updatemodeldeploymentconfigurationdetails struct {
	JsonData       []byte
	DeploymentType string `json:"deploymentType"`
}

// UnmarshalJSON unmarshals json
func (m *updatemodeldeploymentconfigurationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatemodeldeploymentconfigurationdetails updatemodeldeploymentconfigurationdetails
	s := struct {
		Model Unmarshalerupdatemodeldeploymentconfigurationdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DeploymentType = s.Model.DeploymentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatemodeldeploymentconfigurationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeploymentType {
	case "SINGLE_MODEL":
		mm := UpdateSingleModelDeploymentConfigurationDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m updatemodeldeploymentconfigurationdetails) String() string {
	return common.PointerString(m)
}
