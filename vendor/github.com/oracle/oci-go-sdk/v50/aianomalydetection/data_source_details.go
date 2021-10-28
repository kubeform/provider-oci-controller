// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// DataSourceDetails Possible data sources
type DataSourceDetails interface {
}

type datasourcedetails struct {
	JsonData       []byte
	DataSourceType string `json:"dataSourceType"`
}

// UnmarshalJSON unmarshals json
func (m *datasourcedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatasourcedetails datasourcedetails
	s := struct {
		Model Unmarshalerdatasourcedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DataSourceType = s.Model.DataSourceType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *datasourcedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DataSourceType {
	case "ORACLE_OBJECT_STORAGE":
		mm := DataSourceDetailsObjectStorage{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INFLUX":
		mm := DataSourceDetailsInflux{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_ATP":
		mm := DataSourceDetailsAtp{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m datasourcedetails) String() string {
	return common.PointerString(m)
}
