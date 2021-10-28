// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// AddExadataInsightMembersDetails The information about the members of Exadata system to be added.
type AddExadataInsightMembersDetails interface {
}

type addexadatainsightmembersdetails struct {
	JsonData     []byte
	EntitySource string `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *addexadatainsightmembersdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleraddexadatainsightmembersdetails addexadatainsightmembersdetails
	s := struct {
		Model Unmarshaleraddexadatainsightmembersdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *addexadatainsightmembersdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "EM_MANAGED_EXTERNAL_EXADATA":
		mm := AddEmManagedExternalExadataInsightMembersDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m addexadatainsightmembersdetails) String() string {
	return common.PointerString(m)
}
