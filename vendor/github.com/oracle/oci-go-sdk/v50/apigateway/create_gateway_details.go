// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CreateGatewayDetails Information about the new gateway.
type CreateGatewayDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Gateway endpoint type. `PUBLIC` will have a public ip address assigned to it, while `PRIVATE` will only be
	// accessible on a private IP address on the subnet.
	// Example: `PUBLIC` or `PRIVATE`
	EndpointType GatewayEndpointTypeEnum `mandatory:"true" json:"endpointType"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet in which
	// related resources are created.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An array of Network Security Groups OCIDs associated with this API Gateway.
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	CertificateId *string `mandatory:"false" json:"certificateId"`

	ResponseCacheDetails ResponseCacheDetails `mandatory:"false" json:"responseCacheDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	// with no predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateGatewayDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *CreateGatewayDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                           `json:"displayName"`
		NetworkSecurityGroupIds []string                          `json:"networkSecurityGroupIds"`
		CertificateId           *string                           `json:"certificateId"`
		ResponseCacheDetails    responsecachedetails              `json:"responseCacheDetails"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId           *string                           `json:"compartmentId"`
		EndpointType            GatewayEndpointTypeEnum           `json:"endpointType"`
		SubnetId                *string                           `json:"subnetId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.NetworkSecurityGroupIds = make([]string, len(model.NetworkSecurityGroupIds))
	for i, n := range model.NetworkSecurityGroupIds {
		m.NetworkSecurityGroupIds[i] = n
	}

	m.CertificateId = model.CertificateId

	nn, e = model.ResponseCacheDetails.UnmarshalPolymorphicJSON(model.ResponseCacheDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResponseCacheDetails = nn.(ResponseCacheDetails)
	} else {
		m.ResponseCacheDetails = nil
	}

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.EndpointType = model.EndpointType

	m.SubnetId = model.SubnetId

	return
}
