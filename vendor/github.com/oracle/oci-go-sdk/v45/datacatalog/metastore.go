// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// Metastore A Data Catalog Metastore provides a centralized metastore repository for use by other OCI services.
type Metastore struct {

	// The metastore's OCID.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the compartment which holds the metastore.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Location under which managed tables will be created by default. This references Object Storage
	// using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/
	DefaultManagedTableLocation *string `mandatory:"true" json:"defaultManagedTableLocation"`

	// Location under which external tables will be created by default. This references Object Storage
	// using an HDFS URI format. Example: oci://bucket@namespace/sub-dir/
	DefaultExternalTableLocation *string `mandatory:"true" json:"defaultExternalTableLocation"`

	// Mutable name of the metastore.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time at which the metastore was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time at which the metastore was last modified. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the metastore.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Metastore) String() string {
	return common.PointerString(m)
}
