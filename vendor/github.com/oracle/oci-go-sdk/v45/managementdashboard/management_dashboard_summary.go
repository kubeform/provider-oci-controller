// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// ManagementDashboard API
//
// API for the Management Dashboard micro-service. Use this API for dashboard and saved search metadata preservation and to perform  tasks such as creating a dashboard, creating a saved search, and obtaining a list of dashboards and saved searches in a compartment.
//
//

package managementdashboard

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// ManagementDashboardSummary Summary of the properties of a dashboard.
type ManagementDashboardSummary struct {

	// ID of the dashboard.  Same as id.
	DashboardId *string `mandatory:"true" json:"dashboardId"`

	// ID of the dashboard.  Same as dashboardId.
	Id *string `mandatory:"true" json:"id"`

	// Display name of the dashboard.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Description of the dashboard.
	Description *string `mandatory:"true" json:"description"`

	// OCID of the compartment in which the dashboard resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Determines whether the dashboard is an Out-of-the-Box (OOB) dashboard. Note that OOB dashboards are only provided by Oracle and cannot be modified.
	IsOobDashboard *bool `mandatory:"true" json:"isOobDashboard"`

	// User who created the dashboard.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// Date and time the dashboard was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// User who updated the dashboard.
	UpdatedBy *string `mandatory:"true" json:"updatedBy"`

	// Date and time the dashboard was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Version of the metadata.
	MetadataVersion *string `mandatory:"true" json:"metadataVersion"`

	// Screen image of the dashboard.
	ScreenImage *string `mandatory:"true" json:"screenImage"`

	// JSON that contains internationalization options.
	Nls *interface{} `mandatory:"true" json:"nls"`

	// Type of dashboard. NORMAL denotes a single dashboard and SET denotes a dashboard set.
	Type *string `mandatory:"true" json:"type"`

	// Current lifecycle state of the dashboard.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
}

func (m ManagementDashboardSummary) String() string {
	return common.PointerString(m)
}
