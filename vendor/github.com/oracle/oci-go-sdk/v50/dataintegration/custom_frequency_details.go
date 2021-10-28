// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CustomFrequencyDetails Frequency details model to set cron-based frequency
type CustomFrequencyDetails struct {

	// This holds the complete cron expression for this schedule, for example, 10 0/5 * * * ? that fires every 5 minutes, at 10 seconds after the minute (i.e. 10:00:10 am, 10:05:10 am, etc.)
	CustomExpression *string `mandatory:"false" json:"customExpression"`

	// the frequency of the schedule.
	Frequency AbstractFrequencyDetailsFrequencyEnum `mandatory:"false" json:"frequency,omitempty"`
}

//GetFrequency returns Frequency
func (m CustomFrequencyDetails) GetFrequency() AbstractFrequencyDetailsFrequencyEnum {
	return m.Frequency
}

func (m CustomFrequencyDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CustomFrequencyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCustomFrequencyDetails CustomFrequencyDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCustomFrequencyDetails
	}{
		"CUSTOM",
		(MarshalTypeCustomFrequencyDetails)(m),
	}

	return json.Marshal(&s)
}
