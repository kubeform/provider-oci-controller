// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// AiPrivateEndpoint A private network reverse connection creates a connection from service to customer subnet over a private network.
type AiPrivateEndpoint struct {

	// Unique identifier that is immutable.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Subnet Identifier
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// List of DNS zones to be used by the data assets.
	// Example: custpvtsubnet.oraclevcn.com for data asset: db.custpvtsubnet.oraclevcn.com
	DnsZones []string `mandatory:"true" json:"dnsZones"`

	// Private Reverse Connection Endpoint display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time the private endpoint was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the private endpoint was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the private endpoint resource.
	LifecycleState AiPrivateEndpointLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in 'Failed' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The list of dataAssets using the private reverse connection endpoint.
	AttachedDataAssets []string `mandatory:"false" json:"attachedDataAssets"`
}

func (m AiPrivateEndpoint) String() string {
	return common.PointerString(m)
}

// AiPrivateEndpointLifecycleStateEnum Enum with underlying type: string
type AiPrivateEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for AiPrivateEndpointLifecycleStateEnum
const (
	AiPrivateEndpointLifecycleStateCreating AiPrivateEndpointLifecycleStateEnum = "CREATING"
	AiPrivateEndpointLifecycleStateUpdating AiPrivateEndpointLifecycleStateEnum = "UPDATING"
	AiPrivateEndpointLifecycleStateActive   AiPrivateEndpointLifecycleStateEnum = "ACTIVE"
	AiPrivateEndpointLifecycleStateDeleting AiPrivateEndpointLifecycleStateEnum = "DELETING"
	AiPrivateEndpointLifecycleStateDeleted  AiPrivateEndpointLifecycleStateEnum = "DELETED"
	AiPrivateEndpointLifecycleStateFailed   AiPrivateEndpointLifecycleStateEnum = "FAILED"
)

var mappingAiPrivateEndpointLifecycleState = map[string]AiPrivateEndpointLifecycleStateEnum{
	"CREATING": AiPrivateEndpointLifecycleStateCreating,
	"UPDATING": AiPrivateEndpointLifecycleStateUpdating,
	"ACTIVE":   AiPrivateEndpointLifecycleStateActive,
	"DELETING": AiPrivateEndpointLifecycleStateDeleting,
	"DELETED":  AiPrivateEndpointLifecycleStateDeleted,
	"FAILED":   AiPrivateEndpointLifecycleStateFailed,
}

// GetAiPrivateEndpointLifecycleStateEnumValues Enumerates the set of values for AiPrivateEndpointLifecycleStateEnum
func GetAiPrivateEndpointLifecycleStateEnumValues() []AiPrivateEndpointLifecycleStateEnum {
	values := make([]AiPrivateEndpointLifecycleStateEnum, 0)
	for _, v := range mappingAiPrivateEndpointLifecycleState {
		values = append(values, v)
	}
	return values
}
