// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// SensitiveAttribute The sensitive attribute to be used for sensitive content (for password/wallet).
type SensitiveAttribute struct {
	SecretConfig SecretConfig `mandatory:"false" json:"secretConfig"`

	// Attribute to provide sensitive content.
	Value *string `mandatory:"false" json:"value"`
}

func (m SensitiveAttribute) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *SensitiveAttribute) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		SecretConfig secretconfig `json:"secretConfig"`
		Value        *string      `json:"value"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.SecretConfig.UnmarshalPolymorphicJSON(model.SecretConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecretConfig = nn.(SecretConfig)
	} else {
		m.SecretConfig = nil
	}

	m.Value = model.Value

	return
}
