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

// PutRepositoryTagDetails The information needed to create a lightweight Tag
type PutRepositoryTagDetails struct {

	// SHA-1 hash value of the object pointed to by the tag.
	ObjectId *string `mandatory:"true" json:"objectId"`
}

func (m PutRepositoryTagDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PutRepositoryTagDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePutRepositoryTagDetails PutRepositoryTagDetails
	s := struct {
		DiscriminatorParam string `json:"refType"`
		MarshalTypePutRepositoryTagDetails
	}{
		"TAG",
		(MarshalTypePutRepositoryTagDetails)(m),
	}

	return json.Marshal(&s)
}
