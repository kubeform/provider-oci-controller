// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// UpdateConnectDescriptor Connect Descriptor details. If a Private Endpoint was specified in the Connection, the host entry should be a valid IP address.
type UpdateConnectDescriptor struct {

	// Host or IP address of the connect descriptor.
	Host *string `mandatory:"false" json:"host"`

	// Port of the connect descriptor.
	Port *int `mandatory:"false" json:"port"`

	// Database service name.
	DatabaseServiceName *string `mandatory:"false" json:"databaseServiceName"`

	// Connect String. If specified, this will override the stored connect descriptor details.
	// If a Private Endpoint was specified in the Connection, the host entry should be a valid IP address.
	// Supported formats:
	// Easy connect: <host>:<port>/<db_service_name>
	// Long format: (description= (address=(port=<port>)(host=<host>))(connect_data=(service_name=<db_service_name>)))
	ConnectString *string `mandatory:"false" json:"connectString"`
}

func (m UpdateConnectDescriptor) String() string {
	return common.PointerString(m)
}
