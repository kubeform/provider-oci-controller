// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// RequestProtection Module that allows to enable OCI-managed protection capabilities for incoming HTTP requests.
type RequestProtection struct {

	// Ordered list of ProtectionRules. Rules are executed in order of appearance in this array.
	// ProtectionRules in this array can only use protection cCapabilities of REQUEST_PROTECTION_CAPABILITY type.
	Rules []ProtectionRule `mandatory:"false" json:"rules"`
}

func (m RequestProtection) String() string {
	return common.PointerString(m)
}
