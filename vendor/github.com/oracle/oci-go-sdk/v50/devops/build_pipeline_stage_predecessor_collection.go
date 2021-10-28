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

// BuildPipelineStagePredecessorCollection The containing collection for the predecessors of a Stage.
type BuildPipelineStagePredecessorCollection struct {

	// A list of BuildPipelineStagePredecessors for a stage.
	Items []BuildPipelineStagePredecessor `mandatory:"true" json:"items"`
}

func (m BuildPipelineStagePredecessorCollection) String() string {
	return common.PointerString(m)
}
