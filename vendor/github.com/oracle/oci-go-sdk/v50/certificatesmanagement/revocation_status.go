// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// RevocationStatus The current revocation status of the entity.
type RevocationStatus struct {

	// The time when the entity was revoked, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfRevocation *common.SDKTime `mandatory:"true" json:"timeOfRevocation"`

	// The reason the certificate or certificate authority (CA) was revoked.
	RevocationReason RevocationReasonEnum `mandatory:"true" json:"revocationReason"`
}

func (m RevocationStatus) String() string {
	return common.PointerString(m)
}
