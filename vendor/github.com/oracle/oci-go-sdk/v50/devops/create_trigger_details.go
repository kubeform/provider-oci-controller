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

// CreateTriggerDetails The information about new Trigger.
type CreateTriggerDetails interface {

	// Project to which the Trigger will belong
	GetProjectId() *string

	// The list of actions that are to be performed for this Trigger
	GetActions() []TriggerAction

	// Name of the Trigger
	GetDisplayName() *string

	// Optional description about the Trigger
	GetDescription() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createtriggerdetails struct {
	JsonData      []byte
	ProjectId     *string                           `mandatory:"true" json:"projectId"`
	Actions       json.RawMessage                   `mandatory:"true" json:"actions"`
	DisplayName   *string                           `mandatory:"false" json:"displayName"`
	Description   *string                           `mandatory:"false" json:"description"`
	FreeformTags  map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags   map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	TriggerSource string                            `json:"triggerSource"`
}

// UnmarshalJSON unmarshals json
func (m *createtriggerdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatetriggerdetails createtriggerdetails
	s := struct {
		Model Unmarshalercreatetriggerdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ProjectId = s.Model.ProjectId
	m.Actions = s.Model.Actions
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.TriggerSource = s.Model.TriggerSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createtriggerdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.TriggerSource {
	case "GITHUB":
		mm := CreateGithubTriggerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DEVOPS_CODE_REPOSITORY":
		mm := CreateDevopsCodeRepositoryTriggerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GITLAB":
		mm := CreateGitlabTriggerDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetProjectId returns ProjectId
func (m createtriggerdetails) GetProjectId() *string {
	return m.ProjectId
}

//GetActions returns Actions
func (m createtriggerdetails) GetActions() json.RawMessage {
	return m.Actions
}

//GetDisplayName returns DisplayName
func (m createtriggerdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m createtriggerdetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m createtriggerdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createtriggerdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m createtriggerdetails) String() string {
	return common.PointerString(m)
}
