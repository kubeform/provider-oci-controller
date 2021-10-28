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

// InfluxDetailsV1v8 Influx details for V_1_8.
type InfluxDetailsV1v8 struct {

	// DB Name for influx connection
	DatabaseName *string `mandatory:"true" json:"databaseName"`

	// retention policy is how long the bucket would last
	RetentionPolicyName *string `mandatory:"false" json:"retentionPolicyName"`
}

func (m InfluxDetailsV1v8) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InfluxDetailsV1v8) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInfluxDetailsV1v8 InfluxDetailsV1v8
	s := struct {
		DiscriminatorParam string `json:"influxVersion"`
		MarshalTypeInfluxDetailsV1v8
	}{
		"V_1_8",
		(MarshalTypeInfluxDetailsV1v8)(m),
	}

	return json.Marshal(&s)
}
