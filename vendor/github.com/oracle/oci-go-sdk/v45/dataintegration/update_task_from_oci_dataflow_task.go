// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// UpdateTaskFromOciDataflowTask The information about the OCI Dataflow task.
type UpdateTaskFromOciDataflowTask struct {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// An array of parameters.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	DataflowApplication *DataflowApplication `mandatory:"false" json:"dataflowApplication"`
}

//GetKey returns Key
func (m UpdateTaskFromOciDataflowTask) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateTaskFromOciDataflowTask) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m UpdateTaskFromOciDataflowTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m UpdateTaskFromOciDataflowTask) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateTaskFromOciDataflowTask) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateTaskFromOciDataflowTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateTaskFromOciDataflowTask) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateTaskFromOciDataflowTask) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m UpdateTaskFromOciDataflowTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m UpdateTaskFromOciDataflowTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m UpdateTaskFromOciDataflowTask) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m UpdateTaskFromOciDataflowTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m UpdateTaskFromOciDataflowTask) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateTaskFromOciDataflowTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateTaskFromOciDataflowTask) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateTaskFromOciDataflowTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateTaskFromOciDataflowTask UpdateTaskFromOciDataflowTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateTaskFromOciDataflowTask
	}{
		"OCI_DATAFLOW_TASK",
		(MarshalTypeUpdateTaskFromOciDataflowTask)(m),
	}

	return json.Marshal(&s)
}
