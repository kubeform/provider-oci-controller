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

// GithubTriggerCreateResult Trigger Create response specific to Github
type GithubTriggerCreateResult struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Project to which the Trigger belongs
	ProjectId *string `mandatory:"true" json:"projectId"`

	// Compartment to which the Trigger belongs
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of actions that are to be performed for this Trigger
	Actions []TriggerAction `mandatory:"true" json:"actions"`

	// The secret used to validate the incoming Trigger call (this is visible only once after the resource is created)
	Secret *string `mandatory:"true" json:"secret"`

	// The endpoint which listens to Trigger events
	TriggerUrl *string `mandatory:"true" json:"triggerUrl"`

	// Name for Trigger.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description about the Trigger
	Description *string `mandatory:"false" json:"description"`

	// The time the the Trigger was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Trigger was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the Trigger.
	LifecycleState TriggerLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m GithubTriggerCreateResult) GetId() *string {
	return m.Id
}

//GetDisplayName returns DisplayName
func (m GithubTriggerCreateResult) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m GithubTriggerCreateResult) GetDescription() *string {
	return m.Description
}

//GetProjectId returns ProjectId
func (m GithubTriggerCreateResult) GetProjectId() *string {
	return m.ProjectId
}

//GetCompartmentId returns CompartmentId
func (m GithubTriggerCreateResult) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m GithubTriggerCreateResult) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m GithubTriggerCreateResult) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m GithubTriggerCreateResult) GetLifecycleState() TriggerLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m GithubTriggerCreateResult) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetActions returns Actions
func (m GithubTriggerCreateResult) GetActions() []TriggerAction {
	return m.Actions
}

//GetFreeformTags returns FreeformTags
func (m GithubTriggerCreateResult) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m GithubTriggerCreateResult) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m GithubTriggerCreateResult) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m GithubTriggerCreateResult) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GithubTriggerCreateResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGithubTriggerCreateResult GithubTriggerCreateResult
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeGithubTriggerCreateResult
	}{
		"GITHUB",
		(MarshalTypeGithubTriggerCreateResult)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *GithubTriggerCreateResult) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName      *string                           `json:"displayName"`
		Description      *string                           `json:"description"`
		TimeCreated      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated      *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState   TriggerLifecycleStateEnum         `json:"lifecycleState"`
		LifecycleDetails *string                           `json:"lifecycleDetails"`
		FreeformTags     map[string]string                 `json:"freeformTags"`
		DefinedTags      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags       map[string]map[string]interface{} `json:"systemTags"`
		Id               *string                           `json:"id"`
		ProjectId        *string                           `json:"projectId"`
		CompartmentId    *string                           `json:"compartmentId"`
		Actions          []triggeraction                   `json:"actions"`
		Secret           *string                           `json:"secret"`
		TriggerUrl       *string                           `json:"triggerUrl"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.ProjectId = model.ProjectId

	m.CompartmentId = model.CompartmentId

	m.Actions = make([]TriggerAction, len(model.Actions))
	for i, n := range model.Actions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Actions[i] = nn.(TriggerAction)
		} else {
			m.Actions[i] = nil
		}
	}

	m.Secret = model.Secret

	m.TriggerUrl = model.TriggerUrl

	return
}
