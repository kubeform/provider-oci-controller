// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// API for the Email Delivery service. Use this API to send high-volume, application-generated
// emails. For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//
// **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
// If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CreateEmailDomainDetails The configuration details for creating a new email domain.
type CreateEmailDomainDetails struct {

	// The name of the email domain in the Internet Domain Name System (DNS).
	// The email domain name must be unique in the region for this tenancy.
	// Domain names limited to ASCII characters use alphanumeric, dash ("-"), and dot (".") characters.
	// The dash and dot are only allowed between alphanumeric characters.
	// For details, please see: https://tools.ietf.org/html/rfc5321#section-4.1.2
	// Non-ASCII domain names should adopt IDNA2008 normalization (RFC 5891-5892).
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for this email domain.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A string that describes the details about the domain. It does not have to be unique,
	// and you can change it. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateEmailDomainDetails) String() string {
	return common.PointerString(m)
}
