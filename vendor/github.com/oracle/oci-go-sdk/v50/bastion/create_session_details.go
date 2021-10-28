// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Oracle Cloud Infrastructure Bastion provides restricted and time-limited access to target resources that don't have public endpoints. Through the configuration of a bastion, you can let authorized users connect from specific IP addresses to target resources by way of Secure Shell (SSH) sessions hosted on the bastion.
//

package bastion

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CreateSessionDetails The configuration details for a new bastion session. A session lets authorized users connect to a target resource for a predetermined amount of time. The Bastion service recognizes two types of sessions, managed SSH sessions and SSH port forwarding sessions. Managed SSH sessions require that the target resource has an OpenSSH server and the Oracle Cloud Agent both running.
type CreateSessionDetails struct {

	// The unique identifier (OCID) of the bastion on which to create this session.
	BastionId *string `mandatory:"true" json:"bastionId"`

	TargetResourceDetails CreateSessionTargetResourceDetails `mandatory:"true" json:"targetResourceDetails"`

	KeyDetails *PublicKeyDetails `mandatory:"true" json:"keyDetails"`

	// The name of the session.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of the key used to connect to the session. PUB is a standard public key in OpenSSH format.
	KeyType CreateSessionDetailsKeyTypeEnum `mandatory:"false" json:"keyType,omitempty"`

	// The amount of time the session can remain active.
	SessionTtlInSeconds *int `mandatory:"false" json:"sessionTtlInSeconds"`
}

func (m CreateSessionDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateSessionDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName           *string                            `json:"displayName"`
		KeyType               CreateSessionDetailsKeyTypeEnum    `json:"keyType"`
		SessionTtlInSeconds   *int                               `json:"sessionTtlInSeconds"`
		BastionId             *string                            `json:"bastionId"`
		TargetResourceDetails createsessiontargetresourcedetails `json:"targetResourceDetails"`
		KeyDetails            *PublicKeyDetails                  `json:"keyDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.KeyType = model.KeyType

	m.SessionTtlInSeconds = model.SessionTtlInSeconds

	m.BastionId = model.BastionId

	nn, e = model.TargetResourceDetails.UnmarshalPolymorphicJSON(model.TargetResourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TargetResourceDetails = nn.(CreateSessionTargetResourceDetails)
	} else {
		m.TargetResourceDetails = nil
	}

	m.KeyDetails = model.KeyDetails

	return
}

// CreateSessionDetailsKeyTypeEnum Enum with underlying type: string
type CreateSessionDetailsKeyTypeEnum string

// Set of constants representing the allowable values for CreateSessionDetailsKeyTypeEnum
const (
	CreateSessionDetailsKeyTypePub CreateSessionDetailsKeyTypeEnum = "PUB"
)

var mappingCreateSessionDetailsKeyType = map[string]CreateSessionDetailsKeyTypeEnum{
	"PUB": CreateSessionDetailsKeyTypePub,
}

// GetCreateSessionDetailsKeyTypeEnumValues Enumerates the set of values for CreateSessionDetailsKeyTypeEnum
func GetCreateSessionDetailsKeyTypeEnumValues() []CreateSessionDetailsKeyTypeEnum {
	values := make([]CreateSessionDetailsKeyTypeEnum, 0)
	for _, v := range mappingCreateSessionDetailsKeyType {
		values = append(values, v)
	}
	return values
}
