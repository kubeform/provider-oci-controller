// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataLabelingService API
//
// A description of the DataLabelingService API
//

package datalabelingservice

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// Dataset A dataset is a logical collection of records. The dataset contains all the information necessary to describe a record's source, format, type of annotations allowed on these records, and labels allowed on annotations.
type Dataset struct {

	// The OCID of the Dataset.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment of the resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the resource was created, in the timestamp format defined by RFC3339.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was last updated, in the timestamp format defined by RFC3339.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The state of a dataset.
	// CREATING - The dataset is being created.  It will transition to ACTIVE when it is ready for labeling.
	// ACTIVE   - The dataset is ready for labeling.
	// UPDATING - The dataset is being updated.  It and its related resources may be unavailable for other updates until it returns to ACTIVE.
	// NEEDS_ATTENTION - A dataset updation operation has failed due to validation or other errors and needs attention.
	// DELETING - The dataset and its related resources are being deleted.
	// DELETED  - The dataset has been deleted and is no longer available.
	// FAILED   - The dataset has failed due to validation or other errors.
	LifecycleState DatasetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The annotation format name required for labeling records.
	AnnotationFormat *string `mandatory:"true" json:"annotationFormat"`

	DatasetSourceDetails DatasetSourceDetails `mandatory:"true" json:"datasetSourceDetails"`

	DatasetFormatDetails DatasetFormatDetails `mandatory:"true" json:"datasetFormatDetails"`

	LabelSet *LabelSet `mandatory:"true" json:"labelSet"`

	// A user-friendly display name for the resource.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A user provided description of the dataset
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in FAILED or NEEDS_ATTENTION state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	InitialRecordGenerationConfiguration *InitialRecordGenerationConfiguration `mandatory:"false" json:"initialRecordGenerationConfiguration"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Dataset) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *Dataset) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName                          *string                               `json:"displayName"`
		Description                          *string                               `json:"description"`
		LifecycleDetails                     *string                               `json:"lifecycleDetails"`
		InitialRecordGenerationConfiguration *InitialRecordGenerationConfiguration `json:"initialRecordGenerationConfiguration"`
		FreeformTags                         map[string]string                     `json:"freeformTags"`
		DefinedTags                          map[string]map[string]interface{}     `json:"definedTags"`
		SystemTags                           map[string]map[string]interface{}     `json:"systemTags"`
		Id                                   *string                               `json:"id"`
		CompartmentId                        *string                               `json:"compartmentId"`
		TimeCreated                          *common.SDKTime                       `json:"timeCreated"`
		TimeUpdated                          *common.SDKTime                       `json:"timeUpdated"`
		LifecycleState                       DatasetLifecycleStateEnum             `json:"lifecycleState"`
		AnnotationFormat                     *string                               `json:"annotationFormat"`
		DatasetSourceDetails                 datasetsourcedetails                  `json:"datasetSourceDetails"`
		DatasetFormatDetails                 datasetformatdetails                  `json:"datasetFormatDetails"`
		LabelSet                             *LabelSet                             `json:"labelSet"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.LifecycleDetails = model.LifecycleDetails

	m.InitialRecordGenerationConfiguration = model.InitialRecordGenerationConfiguration

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.AnnotationFormat = model.AnnotationFormat

	nn, e = model.DatasetSourceDetails.UnmarshalPolymorphicJSON(model.DatasetSourceDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatasetSourceDetails = nn.(DatasetSourceDetails)
	} else {
		m.DatasetSourceDetails = nil
	}

	nn, e = model.DatasetFormatDetails.UnmarshalPolymorphicJSON(model.DatasetFormatDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.DatasetFormatDetails = nn.(DatasetFormatDetails)
	} else {
		m.DatasetFormatDetails = nil
	}

	m.LabelSet = model.LabelSet

	return
}

// DatasetLifecycleStateEnum Enum with underlying type: string
type DatasetLifecycleStateEnum string

// Set of constants representing the allowable values for DatasetLifecycleStateEnum
const (
	DatasetLifecycleStateCreating       DatasetLifecycleStateEnum = "CREATING"
	DatasetLifecycleStateUpdating       DatasetLifecycleStateEnum = "UPDATING"
	DatasetLifecycleStateActive         DatasetLifecycleStateEnum = "ACTIVE"
	DatasetLifecycleStateNeedsAttention DatasetLifecycleStateEnum = "NEEDS_ATTENTION"
	DatasetLifecycleStateDeleting       DatasetLifecycleStateEnum = "DELETING"
	DatasetLifecycleStateDeleted        DatasetLifecycleStateEnum = "DELETED"
	DatasetLifecycleStateFailed         DatasetLifecycleStateEnum = "FAILED"
)

var mappingDatasetLifecycleState = map[string]DatasetLifecycleStateEnum{
	"CREATING":        DatasetLifecycleStateCreating,
	"UPDATING":        DatasetLifecycleStateUpdating,
	"ACTIVE":          DatasetLifecycleStateActive,
	"NEEDS_ATTENTION": DatasetLifecycleStateNeedsAttention,
	"DELETING":        DatasetLifecycleStateDeleting,
	"DELETED":         DatasetLifecycleStateDeleted,
	"FAILED":          DatasetLifecycleStateFailed,
}

// GetDatasetLifecycleStateEnumValues Enumerates the set of values for DatasetLifecycleStateEnum
func GetDatasetLifecycleStateEnumValues() []DatasetLifecycleStateEnum {
	values := make([]DatasetLifecycleStateEnum, 0)
	for _, v := range mappingDatasetLifecycleState {
		values = append(values, v)
	}
	return values
}
