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
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ManagementSavedSearchSummary Summary of the properties of a saved search.
type ManagementSavedSearchSummary struct {

	// ID of the saved search.
	Id *string `mandatory:"true" json:"id"`

	// Display name of the saved search.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Determines whether the saved search is an Out-of-the-Box (OOB) saved search. Note that OOB saved searches are only provided by Oracle and cannot be modified.
	IsOobSavedSearch *bool `mandatory:"true" json:"isOobSavedSearch"`

	// OCID of the compartment in which the saved search resides.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// ID of the service (for example log-analytics) that owns the saved search. Each service has a unique ID.
	ProviderId *string `mandatory:"true" json:"providerId"`

	// Version of the service that owns this saved search.
	ProviderVersion *string `mandatory:"true" json:"providerVersion"`

	// Name of the service (for example, Logging Analytics) that owns the saved search.
	ProviderName *string `mandatory:"true" json:"providerName"`

	// Description of the saved search.
	Description *string `mandatory:"true" json:"description"`

	// JSON that contains internationalization options.
	Nls *interface{} `mandatory:"true" json:"nls"`

	// Determines how the saved search is displayed in a dashboard.
	Type SavedSearchTypesEnum `mandatory:"true" json:"type"`

	// JSON that contains user interface options.
	UiConfig *interface{} `mandatory:"true" json:"uiConfig"`

	// Array of JSON that contain data source options.
	DataConfig []interface{} `mandatory:"true" json:"dataConfig"`

	// User who created the saved search.
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// User who updated the saved search.
	UpdatedBy *string `mandatory:"true" json:"updatedBy"`

	// Date and time the saved search was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Date and time the saved search was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Screen image of the saved search.
	ScreenImage *string `mandatory:"true" json:"screenImage"`

	// Version of the metadata.
	MetadataVersion *string `mandatory:"true" json:"metadataVersion"`

	// Reference to the HTML file of the widget.
	WidgetTemplate *string `mandatory:"true" json:"widgetTemplate"`

	// Reference to the view model of the widget.
	WidgetVM *string `mandatory:"true" json:"widgetVM"`

	// Current lifecycle state of the saved search.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// Defines parameters for the saved search.
	ParametersConfig []interface{} `mandatory:"false" json:"parametersConfig"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ManagementSavedSearchSummary) String() string {
	return common.PointerString(m)
}
