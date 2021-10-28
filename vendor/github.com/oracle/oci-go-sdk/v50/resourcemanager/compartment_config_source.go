// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CompartmentConfigSource Compartment to use for creating the stack.
// The new stack will include definitions for supported resource types in this compartment.
type CompartmentConfigSource struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to use
	// for creating the stack. The new stack will include definitions for supported
	// resource types in this compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The region to use for creating the stack. The new stack will include definitions for
	// supported resource types in this region.
	Region *string `mandatory:"true" json:"region"`

	// File path to the directory to use for running Terraform.
	// If not specified, the root directory is used.
	// This parameter is ignored for the `configSourceType` value of `COMPARTMENT_CONFIG_SOURCE`.
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	// Filter for services to use with Resource Discovery (https://www.terraform.io/docs/providers/oci/guides/resource_discovery.html#services).
	// For example, "database" limits resource discovery to resource types within the Database service.
	// The specified services must be in scope of the given compartment OCID (tenancy level for root compartment, compartment level otherwise).
	// If not specified, then all services at the scope of the given compartment OCID are used.
	ServicesToDiscover []string `mandatory:"false" json:"servicesToDiscover"`
}

//GetWorkingDirectory returns WorkingDirectory
func (m CompartmentConfigSource) GetWorkingDirectory() *string {
	return m.WorkingDirectory
}

func (m CompartmentConfigSource) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CompartmentConfigSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCompartmentConfigSource CompartmentConfigSource
	s := struct {
		DiscriminatorParam string `json:"configSourceType"`
		MarshalTypeCompartmentConfigSource
	}{
		"COMPARTMENT_CONFIG_SOURCE",
		(MarshalTypeCompartmentConfigSource)(m),
	}

	return json.Marshal(&s)
}
