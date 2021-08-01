// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// DefaultStopDeploymentDetails Definiton of the additional attributes for default deployment stop.
type DefaultStopDeploymentDetails struct {
}

func (m DefaultStopDeploymentDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m DefaultStopDeploymentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDefaultStopDeploymentDetails DefaultStopDeploymentDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDefaultStopDeploymentDetails
	}{
		"DEFAULT",
		(MarshalTypeDefaultStopDeploymentDetails)(m),
	}

	return json.Marshal(&s)
}
