// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// UpdateDatabaseRegistrationDetails The information to update for a DatabaseRegistration.
type UpdateDatabaseRegistrationDetails struct {

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Metadata about this specific object.
	Description *string `mandatory:"false" json:"description"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A three-label Fully Qualified Domain Name (FQDN) for a resource.
	Fqdn *string `mandatory:"false" json:"fqdn"`

	// The username Oracle GoldenGate uses to connect the associated RDBMS.  This username must already exist and be available for use by the database.  It must conform to the security requirements implemented by the database including length, case sensitivity, and so on.
	Username *string `mandatory:"false" json:"username"`

	// The password Oracle GoldenGate uses to connect the associated RDBMS.  It must conform to the specific security requirements implemented by the database including length, case sensitivity, and so on.
	Password *string `mandatory:"false" json:"password"`

	// Connect descriptor or Easy Connect Naming method that Oracle GoldenGate uses to connect to a database.
	ConnectionString *string `mandatory:"false" json:"connectionString"`

	// The wallet contents Oracle GoldenGate uses to make connections to a database.  This attribute is expected to be base64 encoded.
	Wallet *string `mandatory:"false" json:"wallet"`

	// Credential store alias.
	AliasName *string `mandatory:"false" json:"aliasName"`
}

func (m UpdateDatabaseRegistrationDetails) String() string {
	return common.PointerString(m)
}
