// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// DbManagementPrivateEndpoint A Database Management private endpoint that allows Database Management services to connect to databases in a customer's virtual cloud network (VCN).
type DbManagementPrivateEndpoint struct {

	// The OCID of the Database Management private endpoint.
	Id *string `mandatory:"true" json:"id"`

	// The display name of the private endpoint.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID of the subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The private IP addresses assigned to the private endpoint.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The description of the private endpoint.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the private endpoint was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current state of the private endpoint.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The OCIDs of the network security groups that the private endpoint belongs to.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

func (m DbManagementPrivateEndpoint) String() string {
	return common.PointerString(m)
}
