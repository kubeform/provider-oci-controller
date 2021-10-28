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

// GitlabFilter The filter for Gitlab events
type GitlabFilter struct {
	Include *GitlabFilterAttributes `mandatory:"false" json:"include"`

	// The events, example PUSH, PULL_REQUEST_MERGE etc.
	Events []GitlabFilterEventsEnum `mandatory:"false" json:"events,omitempty"`
}

func (m GitlabFilter) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m GitlabFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGitlabFilter GitlabFilter
	s := struct {
		DiscriminatorParam string `json:"triggerSource"`
		MarshalTypeGitlabFilter
	}{
		"GITLAB",
		(MarshalTypeGitlabFilter)(m),
	}

	return json.Marshal(&s)
}

// GitlabFilterEventsEnum Enum with underlying type: string
type GitlabFilterEventsEnum string

// Set of constants representing the allowable values for GitlabFilterEventsEnum
const (
	GitlabFilterEventsPush                GitlabFilterEventsEnum = "PUSH"
	GitlabFilterEventsPullRequestCreated  GitlabFilterEventsEnum = "PULL_REQUEST_CREATED"
	GitlabFilterEventsPullRequestUpdated  GitlabFilterEventsEnum = "PULL_REQUEST_UPDATED"
	GitlabFilterEventsPullRequestReopened GitlabFilterEventsEnum = "PULL_REQUEST_REOPENED"
	GitlabFilterEventsPullRequestMerged   GitlabFilterEventsEnum = "PULL_REQUEST_MERGED"
)

var mappingGitlabFilterEvents = map[string]GitlabFilterEventsEnum{
	"PUSH":                  GitlabFilterEventsPush,
	"PULL_REQUEST_CREATED":  GitlabFilterEventsPullRequestCreated,
	"PULL_REQUEST_UPDATED":  GitlabFilterEventsPullRequestUpdated,
	"PULL_REQUEST_REOPENED": GitlabFilterEventsPullRequestReopened,
	"PULL_REQUEST_MERGED":   GitlabFilterEventsPullRequestMerged,
}

// GetGitlabFilterEventsEnumValues Enumerates the set of values for GitlabFilterEventsEnum
func GetGitlabFilterEventsEnumValues() []GitlabFilterEventsEnum {
	values := make([]GitlabFilterEventsEnum, 0)
	for _, v := range mappingGitlabFilterEvents {
		values = append(values, v)
	}
	return values
}
