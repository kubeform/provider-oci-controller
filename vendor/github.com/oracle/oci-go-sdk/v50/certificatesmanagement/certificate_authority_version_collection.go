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

// CertificateAuthorityVersionCollection The results of a certificate authority (CA) version search. This object contains CA version summary objects and other data.
type CertificateAuthorityVersionCollection struct {

	// A list of certificate authority version summary objects.
	Items []CertificateAuthorityVersionSummary `mandatory:"true" json:"items"`
}

func (m CertificateAuthorityVersionCollection) String() string {
	return common.PointerString(m)
}
