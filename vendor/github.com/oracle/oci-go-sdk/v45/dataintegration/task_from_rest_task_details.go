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

// TaskFromRestTaskDetails The information about the Generic REST task.
type TaskFromRestTaskDetails struct {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// An array of parameters.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	AuthDetails *AuthDetails `mandatory:"false" json:"authDetails"`

	Endpoint *Expression `mandatory:"false" json:"endpoint"`

	// The headers for the REST call.
	Headers *interface{} `mandatory:"false" json:"headers"`

	// JSON data for payload body.
	JsonData *string `mandatory:"false" json:"jsonData"`

	CancelEndpoint *Expression `mandatory:"false" json:"cancelEndpoint"`

	// The REST method to use.
	MethodType TaskFromRestTaskDetailsMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The invocation type to be used for Generic REST invocation.
	ApiCallMode TaskFromRestTaskDetailsApiCallModeEnum `mandatory:"false" json:"apiCallMode,omitempty"`

	// The REST method to use for canceling the original request.
	CancelMethodType TaskFromRestTaskDetailsCancelMethodTypeEnum `mandatory:"false" json:"cancelMethodType,omitempty"`
}

//GetKey returns Key
func (m TaskFromRestTaskDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m TaskFromRestTaskDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m TaskFromRestTaskDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m TaskFromRestTaskDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m TaskFromRestTaskDetails) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m TaskFromRestTaskDetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m TaskFromRestTaskDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m TaskFromRestTaskDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m TaskFromRestTaskDetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m TaskFromRestTaskDetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m TaskFromRestTaskDetails) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m TaskFromRestTaskDetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m TaskFromRestTaskDetails) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetMetadata returns Metadata
func (m TaskFromRestTaskDetails) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

//GetKeyMap returns KeyMap
func (m TaskFromRestTaskDetails) GetKeyMap() map[string]string {
	return m.KeyMap
}

//GetRegistryMetadata returns RegistryMetadata
func (m TaskFromRestTaskDetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m TaskFromRestTaskDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TaskFromRestTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTaskFromRestTaskDetails TaskFromRestTaskDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTaskFromRestTaskDetails
	}{
		"REST_TASK",
		(MarshalTypeTaskFromRestTaskDetails)(m),
	}

	return json.Marshal(&s)
}

// TaskFromRestTaskDetailsMethodTypeEnum Enum with underlying type: string
type TaskFromRestTaskDetailsMethodTypeEnum string

// Set of constants representing the allowable values for TaskFromRestTaskDetailsMethodTypeEnum
const (
	TaskFromRestTaskDetailsMethodTypeGet    TaskFromRestTaskDetailsMethodTypeEnum = "GET"
	TaskFromRestTaskDetailsMethodTypePost   TaskFromRestTaskDetailsMethodTypeEnum = "POST"
	TaskFromRestTaskDetailsMethodTypePatch  TaskFromRestTaskDetailsMethodTypeEnum = "PATCH"
	TaskFromRestTaskDetailsMethodTypeDelete TaskFromRestTaskDetailsMethodTypeEnum = "DELETE"
	TaskFromRestTaskDetailsMethodTypePut    TaskFromRestTaskDetailsMethodTypeEnum = "PUT"
)

var mappingTaskFromRestTaskDetailsMethodType = map[string]TaskFromRestTaskDetailsMethodTypeEnum{
	"GET":    TaskFromRestTaskDetailsMethodTypeGet,
	"POST":   TaskFromRestTaskDetailsMethodTypePost,
	"PATCH":  TaskFromRestTaskDetailsMethodTypePatch,
	"DELETE": TaskFromRestTaskDetailsMethodTypeDelete,
	"PUT":    TaskFromRestTaskDetailsMethodTypePut,
}

// GetTaskFromRestTaskDetailsMethodTypeEnumValues Enumerates the set of values for TaskFromRestTaskDetailsMethodTypeEnum
func GetTaskFromRestTaskDetailsMethodTypeEnumValues() []TaskFromRestTaskDetailsMethodTypeEnum {
	values := make([]TaskFromRestTaskDetailsMethodTypeEnum, 0)
	for _, v := range mappingTaskFromRestTaskDetailsMethodType {
		values = append(values, v)
	}
	return values
}

// TaskFromRestTaskDetailsApiCallModeEnum Enum with underlying type: string
type TaskFromRestTaskDetailsApiCallModeEnum string

// Set of constants representing the allowable values for TaskFromRestTaskDetailsApiCallModeEnum
const (
	TaskFromRestTaskDetailsApiCallModeSynchronous         TaskFromRestTaskDetailsApiCallModeEnum = "SYNCHRONOUS"
	TaskFromRestTaskDetailsApiCallModeAsyncOciWorkrequest TaskFromRestTaskDetailsApiCallModeEnum = "ASYNC_OCI_WORKREQUEST"
)

var mappingTaskFromRestTaskDetailsApiCallMode = map[string]TaskFromRestTaskDetailsApiCallModeEnum{
	"SYNCHRONOUS":           TaskFromRestTaskDetailsApiCallModeSynchronous,
	"ASYNC_OCI_WORKREQUEST": TaskFromRestTaskDetailsApiCallModeAsyncOciWorkrequest,
}

// GetTaskFromRestTaskDetailsApiCallModeEnumValues Enumerates the set of values for TaskFromRestTaskDetailsApiCallModeEnum
func GetTaskFromRestTaskDetailsApiCallModeEnumValues() []TaskFromRestTaskDetailsApiCallModeEnum {
	values := make([]TaskFromRestTaskDetailsApiCallModeEnum, 0)
	for _, v := range mappingTaskFromRestTaskDetailsApiCallMode {
		values = append(values, v)
	}
	return values
}

// TaskFromRestTaskDetailsCancelMethodTypeEnum Enum with underlying type: string
type TaskFromRestTaskDetailsCancelMethodTypeEnum string

// Set of constants representing the allowable values for TaskFromRestTaskDetailsCancelMethodTypeEnum
const (
	TaskFromRestTaskDetailsCancelMethodTypeGet    TaskFromRestTaskDetailsCancelMethodTypeEnum = "GET"
	TaskFromRestTaskDetailsCancelMethodTypePost   TaskFromRestTaskDetailsCancelMethodTypeEnum = "POST"
	TaskFromRestTaskDetailsCancelMethodTypePatch  TaskFromRestTaskDetailsCancelMethodTypeEnum = "PATCH"
	TaskFromRestTaskDetailsCancelMethodTypeDelete TaskFromRestTaskDetailsCancelMethodTypeEnum = "DELETE"
	TaskFromRestTaskDetailsCancelMethodTypePut    TaskFromRestTaskDetailsCancelMethodTypeEnum = "PUT"
)

var mappingTaskFromRestTaskDetailsCancelMethodType = map[string]TaskFromRestTaskDetailsCancelMethodTypeEnum{
	"GET":    TaskFromRestTaskDetailsCancelMethodTypeGet,
	"POST":   TaskFromRestTaskDetailsCancelMethodTypePost,
	"PATCH":  TaskFromRestTaskDetailsCancelMethodTypePatch,
	"DELETE": TaskFromRestTaskDetailsCancelMethodTypeDelete,
	"PUT":    TaskFromRestTaskDetailsCancelMethodTypePut,
}

// GetTaskFromRestTaskDetailsCancelMethodTypeEnumValues Enumerates the set of values for TaskFromRestTaskDetailsCancelMethodTypeEnum
func GetTaskFromRestTaskDetailsCancelMethodTypeEnumValues() []TaskFromRestTaskDetailsCancelMethodTypeEnum {
	values := make([]TaskFromRestTaskDetailsCancelMethodTypeEnum, 0)
	for _, v := range mappingTaskFromRestTaskDetailsCancelMethodType {
		values = append(values, v)
	}
	return values
}
