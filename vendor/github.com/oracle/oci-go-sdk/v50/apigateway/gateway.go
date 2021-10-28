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

// Gateway A gateway is a virtual network appliance in a regional subnet. A gateway routes inbound traffic to back-end services including public, private, and partner HTTP APIs, as well as Oracle Functions. Avoid entering confidential information. For more information, see
// API Gateway Concepts (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayconcepts.htm).
type Gateway struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Gateway endpoint type. `PUBLIC` will have a public ip address assigned to it, while `PRIVATE` will only be
	// accessible on a private IP address on the subnet.
	// Example: `PUBLIC` or `PRIVATE`
	EndpointType GatewayEndpointTypeEnum `mandatory:"true" json:"endpointType"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet in which
	// related resources are created.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// An array of Network Security Groups OCIDs associated with this API Gateway.
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the gateway.
	LifecycleState GatewayLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a
	// resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The hostname for APIs deployed on the gateway.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	CertificateId *string `mandatory:"false" json:"certificateId"`

	// An array of IP addresses associated with the gateway.
	IpAddresses []IpAddress `mandatory:"false" json:"ipAddresses"`

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

func (m Gateway) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Gateway) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName             *string                           `json:"displayName"`
		SubnetId                *string                           `json:"subnetId"`
		NetworkSecurityGroupIds []string                          `json:"networkSecurityGroupIds"`
		TimeCreated             *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated             *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState          GatewayLifecycleStateEnum         `json:"lifecycleState"`
		LifecycleDetails        *string                           `json:"lifecycleDetails"`
		Hostname                *string                           `json:"hostname"`
		CertificateId           *string                           `json:"certificateId"`
		IpAddresses             []IpAddress                       `json:"ipAddresses"`
		ResponseCacheDetails    responsecachedetails              `json:"responseCacheDetails"`
		FreeformTags            map[string]string                 `json:"freeformTags"`
		DefinedTags             map[string]map[string]interface{} `json:"definedTags"`
		Id                      *string                           `json:"id"`
		CompartmentId           *string                           `json:"compartmentId"`
		EndpointType            GatewayEndpointTypeEnum           `json:"endpointType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.SubnetId = model.SubnetId

	m.NetworkSecurityGroupIds = make([]string, len(model.NetworkSecurityGroupIds))
	for i, n := range model.NetworkSecurityGroupIds {
		m.NetworkSecurityGroupIds[i] = n
	}

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.Hostname = model.Hostname

	m.CertificateId = model.CertificateId

	m.IpAddresses = make([]IpAddress, len(model.IpAddresses))
	for i, n := range model.IpAddresses {
		m.IpAddresses[i] = n
	}

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

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.EndpointType = model.EndpointType

	return
}

// GatewayEndpointTypeEnum Enum with underlying type: string
type GatewayEndpointTypeEnum string

// Set of constants representing the allowable values for GatewayEndpointTypeEnum
const (
	GatewayEndpointTypePublic  GatewayEndpointTypeEnum = "PUBLIC"
	GatewayEndpointTypePrivate GatewayEndpointTypeEnum = "PRIVATE"
)

var mappingGatewayEndpointType = map[string]GatewayEndpointTypeEnum{
	"PUBLIC":  GatewayEndpointTypePublic,
	"PRIVATE": GatewayEndpointTypePrivate,
}

// GetGatewayEndpointTypeEnumValues Enumerates the set of values for GatewayEndpointTypeEnum
func GetGatewayEndpointTypeEnumValues() []GatewayEndpointTypeEnum {
	values := make([]GatewayEndpointTypeEnum, 0)
	for _, v := range mappingGatewayEndpointType {
		values = append(values, v)
	}
	return values
}

// GatewayLifecycleStateEnum Enum with underlying type: string
type GatewayLifecycleStateEnum string

// Set of constants representing the allowable values for GatewayLifecycleStateEnum
const (
	GatewayLifecycleStateCreating GatewayLifecycleStateEnum = "CREATING"
	GatewayLifecycleStateActive   GatewayLifecycleStateEnum = "ACTIVE"
	GatewayLifecycleStateUpdating GatewayLifecycleStateEnum = "UPDATING"
	GatewayLifecycleStateDeleting GatewayLifecycleStateEnum = "DELETING"
	GatewayLifecycleStateDeleted  GatewayLifecycleStateEnum = "DELETED"
	GatewayLifecycleStateFailed   GatewayLifecycleStateEnum = "FAILED"
)

var mappingGatewayLifecycleState = map[string]GatewayLifecycleStateEnum{
	"CREATING": GatewayLifecycleStateCreating,
	"ACTIVE":   GatewayLifecycleStateActive,
	"UPDATING": GatewayLifecycleStateUpdating,
	"DELETING": GatewayLifecycleStateDeleting,
	"DELETED":  GatewayLifecycleStateDeleted,
	"FAILED":   GatewayLifecycleStateFailed,
}

// GetGatewayLifecycleStateEnumValues Enumerates the set of values for GatewayLifecycleStateEnum
func GetGatewayLifecycleStateEnumValues() []GatewayLifecycleStateEnum {
	values := make([]GatewayLifecycleStateEnum, 0)
	for _, v := range mappingGatewayLifecycleState {
		values = append(values, v)
	}
	return values
}
