// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// DiffResponse Response object for obtaining list of changed files.
type DiffResponse struct {

	// List of changes in the diff.
	Changes []DiffResponseEntry `mandatory:"true" json:"changes"`

	// Boolean for whether all changes are included in the response.
	AreAllChangesIncluded *bool `mandatory:"false" json:"areAllChangesIncluded"`

	// Count of each type of change in diff.
	ChangeTypeCount map[string]int `mandatory:"false" json:"changeTypeCount"`

	// The ID of the common commit between source and target.
	CommonCommit *string `mandatory:"false" json:"commonCommit"`

	// The number of commits source is ahead of target by.
	CommitsAheadCount *int `mandatory:"false" json:"commitsAheadCount"`

	// The number of commits source is behind target by.
	CommitsBehindCount *int `mandatory:"false" json:"commitsBehindCount"`

	// The number of lines added in whole diff.
	AddedLinesCount *int `mandatory:"false" json:"addedLinesCount"`

	// The number of lines deleted in whole diff.
	DeletedLinesCount *int `mandatory:"false" json:"deletedLinesCount"`
}

func (m DiffResponse) String() string {
	return common.PointerString(m)
}
