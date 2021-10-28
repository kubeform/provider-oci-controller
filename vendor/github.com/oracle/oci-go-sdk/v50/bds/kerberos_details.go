// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// KerberosDetails Details about the Kerberos principals.
type KerberosDetails struct {

	// Name of the Kerberos principal.
	PrincipalName *string `mandatory:"true" json:"principalName"`

	// Location of the keytab file
	KeytabFile *string `mandatory:"true" json:"keytabFile"`
}

func (m KerberosDetails) String() string {
	return common.PointerString(m)
}
