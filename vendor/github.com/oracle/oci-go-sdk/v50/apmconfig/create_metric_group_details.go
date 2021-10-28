// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Apm Configuration API
//
// An API for the APM Configuration service. Use this API to query and set APM configuration.
//

package apmconfig

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// CreateMetricGroupDetails A Metric Group.
type CreateMetricGroupDetails struct {

	// The name of this metric group
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of a Span Filter. The filterId is mandatory for the creation
	// of MetricGroups. A filterId will be generated when a Span Filter is created.
	FilterId *string `mandatory:"true" json:"filterId"`

	Metrics []Metric `mandatory:"true" json:"metrics"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The namespace to write the metrics to
	Namespace *string `mandatory:"false" json:"namespace"`

	// A list of dimensions for this metric
	Dimensions []Dimension `mandatory:"false" json:"dimensions"`
}

//GetFreeformTags returns FreeformTags
func (m CreateMetricGroupDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateMetricGroupDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateMetricGroupDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateMetricGroupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMetricGroupDetails CreateMetricGroupDetails
	s := struct {
		DiscriminatorParam string `json:"configType"`
		MarshalTypeCreateMetricGroupDetails
	}{
		"METRIC_GROUP",
		(MarshalTypeCreateMetricGroupDetails)(m),
	}

	return json.Marshal(&s)
}
