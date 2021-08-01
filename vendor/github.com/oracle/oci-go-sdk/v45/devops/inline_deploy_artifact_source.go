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

// InlineDeployArtifactSource Specifies the inline deployment artifact source details.
type InlineDeployArtifactSource struct {

	// base64 Encoded String
	Base64EncodedContent []byte `mandatory:"true" json:"base64EncodedContent"`
}

func (m InlineDeployArtifactSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InlineDeployArtifactSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInlineDeployArtifactSource InlineDeployArtifactSource
	s := struct {
		DiscriminatorParam string `json:"deployArtifactSourceType"`
		MarshalTypeInlineDeployArtifactSource
	}{
		"INLINE",
		(MarshalTypeInlineDeployArtifactSource)(m),
	}

	return json.Marshal(&s)
}
