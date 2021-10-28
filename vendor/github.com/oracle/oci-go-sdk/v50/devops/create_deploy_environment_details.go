// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CreateDeployEnvironmentDetails The information about new deployment environment.
type CreateDeployEnvironmentDetails interface {

	// The OCID of a project.
	GetProjectId() *string

	// Optional description about the deployment environment.
	GetDescription() *string

	// Deployment environment display name. Avoid entering confidential information.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createdeployenvironmentdetails struct {
	JsonData              []byte
	ProjectId             *string                           `mandatory:"true" json:"projectId"`
	Description           *string                           `mandatory:"false" json:"description"`
	DisplayName           *string                           `mandatory:"false" json:"displayName"`
	FreeformTags          map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags           map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	DeployEnvironmentType string                            `json:"deployEnvironmentType"`
}

// UnmarshalJSON unmarshals json
func (m *createdeployenvironmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatedeployenvironmentdetails createdeployenvironmentdetails
	s := struct {
		Model Unmarshalercreatedeployenvironmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ProjectId = s.Model.ProjectId
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DeployEnvironmentType = s.Model.DeployEnvironmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createdeployenvironmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployEnvironmentType {
	case "COMPUTE_INSTANCE_GROUP":
		mm := CreateComputeInstanceGroupDeployEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CLUSTER":
		mm := CreateOkeClusterDeployEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FUNCTION":
		mm := CreateFunctionDeployEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetProjectId returns ProjectId
func (m createdeployenvironmentdetails) GetProjectId() *string {
	return m.ProjectId
}

//GetDescription returns Description
func (m createdeployenvironmentdetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m createdeployenvironmentdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m createdeployenvironmentdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createdeployenvironmentdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createdeployenvironmentdetails) String() string {
	return common.PointerString(m)
}
