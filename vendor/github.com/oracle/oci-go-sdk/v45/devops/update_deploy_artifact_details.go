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
	"github.com/oracle/oci-go-sdk/v45/common"
)

// UpdateDeployArtifactDetails The information to be updated for the artifact.
type UpdateDeployArtifactDetails struct {

	// Optional description about the deployment artifact.
	Description *string `mandatory:"false" json:"description"`

	// Deployment artifact display name. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Type of the deployment artifact.
	DeployArtifactType DeployArtifactDeployArtifactTypeEnum `mandatory:"false" json:"deployArtifactType,omitempty"`

	DeployArtifactSource DeployArtifactSource `mandatory:"false" json:"deployArtifactSource"`

	// Mode for artifact parameter substitution.
	ArgumentSubstitutionMode DeployArtifactArgumentSubstitutionModeEnum `mandatory:"false" json:"argumentSubstitutionMode,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateDeployArtifactDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateDeployArtifactDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description              *string                                    `json:"description"`
		DisplayName              *string                                    `json:"displayName"`
		DeployArtifactType       DeployArtifactDeployArtifactTypeEnum       `json:"deployArtifactType"`
		DeployArtifactSource     deployartifactsource                       `json:"deployArtifactSource"`
		ArgumentSubstitutionMode DeployArtifactArgumentSubstitutionModeEnum `json:"argumentSubstitutionMode"`
		FreeformTags             map[string]string                          `json:"freeformTags"`
		DefinedTags              map[string]map[string]interface{}          `json:"definedTags"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.DeployArtifactType = model.DeployArtifactType

	nn, e = model.DeployArtifactSource.UnmarshalPolymorphicJSON(model.DeployArtifactSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DeployArtifactSource = nn.(DeployArtifactSource)
	} else {
		m.DeployArtifactSource = nil
	}

	m.ArgumentSubstitutionMode = model.ArgumentSubstitutionMode

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	return
}
