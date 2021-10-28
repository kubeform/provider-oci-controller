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

// DiffResponseEntry Entry for description of change on a file.
type DiffResponseEntry struct {

	// Type of change made to file.
	ChangeType *string `mandatory:"true" json:"changeType"`

	// The type of the changed object.
	ObjectType *string `mandatory:"false" json:"objectType"`

	// The ID of the commit where the change is coming from.
	CommitId *string `mandatory:"false" json:"commitId"`

	// The path on the target to the changed object.
	OldPath *string `mandatory:"false" json:"oldPath"`

	// The path on the source to the changed object.
	NewPath *string `mandatory:"false" json:"newPath"`

	// The ID of the changed object on the target.
	OldId *string `mandatory:"false" json:"oldId"`

	// The ID of the changed object on the source.
	NewId *string `mandatory:"false" json:"newId"`

	// The URL of the changed object.
	Url *string `mandatory:"false" json:"url"`

	// The number of lines added in whole diff.
	AddedLinesCount *int `mandatory:"false" json:"addedLinesCount"`

	// The number of lines deleted in whole diff.
	DeletedLinesCount *int `mandatory:"false" json:"deletedLinesCount"`

	// Indicates whether the changed file contains conflicts.
	AreConflictsInFile *bool `mandatory:"false" json:"areConflictsInFile"`
}

func (m DiffResponseEntry) String() string {
	return common.PointerString(m)
}
