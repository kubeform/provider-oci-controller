// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// CreateCustomEndpointDetails Details for a custom endpoint for the integration instance (update).
type CreateCustomEndpointDetails struct {

	// A custom hostname to be used for the integration instance URL, in FQDN format.
	Hostname *string `mandatory:"true" json:"hostname"`

	// Optional OCID of a vault/secret containing a private SSL certificate bundle to be used for the custom hostname.
	// All certificates should be stored in a single base64 encoded secret
	// Note the update will fail if this is not a valid certificate.
	CertificateSecretId *string `mandatory:"false" json:"certificateSecretId"`
}

func (m CreateCustomEndpointDetails) String() string {
	return common.PointerString(m)
}
