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

// DiffSummary Response object for showing diffs for a file between two revisions.
type DiffSummary struct {

	// List of changed section in the file.
	Changes []DiffChunk `mandatory:"true" json:"changes"`

	// The path on the baseVersion to the changed object.
	OldPath *string `mandatory:"false" json:"oldPath"`

	// The path on the targetVersion to the changed object.
	NewPath *string `mandatory:"false" json:"newPath"`

	// The ID of the changed object on the baseVersion.
	OldId *string `mandatory:"false" json:"oldId"`

	// The ID of the changed object on the targetVersion.
	NewId *string `mandatory:"false" json:"newId"`

	// Indicates whether the changed file contains conflicts.
	AreConflictsInFile *bool `mandatory:"false" json:"areConflictsInFile"`

	// Indicates whether the file is large.
	IsLarge *bool `mandatory:"false" json:"isLarge"`

	// Indicates whether the file is binary.
	IsBinary *bool `mandatory:"false" json:"isBinary"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m DiffSummary) String() string {
	return common.PointerString(m)
}
