// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateProject            OperationTypeEnum = "CREATE_PROJECT"
	OperationTypeUpdateProject            OperationTypeEnum = "UPDATE_PROJECT"
	OperationTypeDeleteProject            OperationTypeEnum = "DELETE_PROJECT"
	OperationTypeMoveProject              OperationTypeEnum = "MOVE_PROJECT"
	OperationTypeCreateDeployPipeline     OperationTypeEnum = "CREATE_DEPLOY_PIPELINE"
	OperationTypeUpdateDeployPipeline     OperationTypeEnum = "UPDATE_DEPLOY_PIPELINE"
	OperationTypeDeleteDeployPipeline     OperationTypeEnum = "DELETE_DEPLOY_PIPELINE"
	OperationTypeCreateDeployStage        OperationTypeEnum = "CREATE_DEPLOY_STAGE"
	OperationTypeUpdateDeployStage        OperationTypeEnum = "UPDATE_DEPLOY_STAGE"
	OperationTypeDeleteDeployStage        OperationTypeEnum = "DELETE_DEPLOY_STAGE"
	OperationTypeCreateDeployArtifact     OperationTypeEnum = "CREATE_DEPLOY_ARTIFACT"
	OperationTypeUpdateDeployArtifact     OperationTypeEnum = "UPDATE_DEPLOY_ARTIFACT"
	OperationTypeDeleteDeployArtifact     OperationTypeEnum = "DELETE_DEPLOY_ARTIFACT"
	OperationTypeCreateDeployEnvironment  OperationTypeEnum = "CREATE_DEPLOY_ENVIRONMENT"
	OperationTypeUpdateDeployEnvironment  OperationTypeEnum = "UPDATE_DEPLOY_ENVIRONMENT"
	OperationTypeDeleteDeployEnvironment  OperationTypeEnum = "DELETE_DEPLOY_ENVIRONMENT"
	OperationTypeCreateDeployment         OperationTypeEnum = "CREATE_DEPLOYMENT"
	OperationTypeUpdateDeployment         OperationTypeEnum = "UPDATE_DEPLOYMENT"
	OperationTypeDeleteDeployment         OperationTypeEnum = "DELETE_DEPLOYMENT"
	OperationTypeCreateBuildPipeline      OperationTypeEnum = "CREATE_BUILD_PIPELINE"
	OperationTypeUpdateBuildPipeline      OperationTypeEnum = "UPDATE_BUILD_PIPELINE"
	OperationTypeDeleteBuildPipeline      OperationTypeEnum = "DELETE_BUILD_PIPELINE"
	OperationTypeCreateBuildPipelineStage OperationTypeEnum = "CREATE_BUILD_PIPELINE_STAGE"
	OperationTypeUpdateBuildPipelineStage OperationTypeEnum = "UPDATE_BUILD_PIPELINE_STAGE"
	OperationTypeDeleteBuildPipelineStage OperationTypeEnum = "DELETE_BUILD_PIPELINE_STAGE"
	OperationTypeCreateConnection         OperationTypeEnum = "CREATE_CONNECTION"
	OperationTypeUpdateConnection         OperationTypeEnum = "UPDATE_CONNECTION"
	OperationTypeDeleteConnection         OperationTypeEnum = "DELETE_CONNECTION"
	OperationTypeCreateTrigger            OperationTypeEnum = "CREATE_TRIGGER"
	OperationTypeUpdateTrigger            OperationTypeEnum = "UPDATE_TRIGGER"
	OperationTypeDeleteTrigger            OperationTypeEnum = "DELETE_TRIGGER"
	OperationTypeExecuteTrigger           OperationTypeEnum = "EXECUTE_TRIGGER"
	OperationTypeCreateRepository         OperationTypeEnum = "CREATE_REPOSITORY"
	OperationTypeUpdateRepository         OperationTypeEnum = "UPDATE_REPOSITORY"
	OperationTypeDeleteRepository         OperationTypeEnum = "DELETE_REPOSITORY"
	OperationTypeMirrorRepository         OperationTypeEnum = "MIRROR_REPOSITORY"
)

var mappingOperationType = map[string]OperationTypeEnum{
	"CREATE_PROJECT":              OperationTypeCreateProject,
	"UPDATE_PROJECT":              OperationTypeUpdateProject,
	"DELETE_PROJECT":              OperationTypeDeleteProject,
	"MOVE_PROJECT":                OperationTypeMoveProject,
	"CREATE_DEPLOY_PIPELINE":      OperationTypeCreateDeployPipeline,
	"UPDATE_DEPLOY_PIPELINE":      OperationTypeUpdateDeployPipeline,
	"DELETE_DEPLOY_PIPELINE":      OperationTypeDeleteDeployPipeline,
	"CREATE_DEPLOY_STAGE":         OperationTypeCreateDeployStage,
	"UPDATE_DEPLOY_STAGE":         OperationTypeUpdateDeployStage,
	"DELETE_DEPLOY_STAGE":         OperationTypeDeleteDeployStage,
	"CREATE_DEPLOY_ARTIFACT":      OperationTypeCreateDeployArtifact,
	"UPDATE_DEPLOY_ARTIFACT":      OperationTypeUpdateDeployArtifact,
	"DELETE_DEPLOY_ARTIFACT":      OperationTypeDeleteDeployArtifact,
	"CREATE_DEPLOY_ENVIRONMENT":   OperationTypeCreateDeployEnvironment,
	"UPDATE_DEPLOY_ENVIRONMENT":   OperationTypeUpdateDeployEnvironment,
	"DELETE_DEPLOY_ENVIRONMENT":   OperationTypeDeleteDeployEnvironment,
	"CREATE_DEPLOYMENT":           OperationTypeCreateDeployment,
	"UPDATE_DEPLOYMENT":           OperationTypeUpdateDeployment,
	"DELETE_DEPLOYMENT":           OperationTypeDeleteDeployment,
	"CREATE_BUILD_PIPELINE":       OperationTypeCreateBuildPipeline,
	"UPDATE_BUILD_PIPELINE":       OperationTypeUpdateBuildPipeline,
	"DELETE_BUILD_PIPELINE":       OperationTypeDeleteBuildPipeline,
	"CREATE_BUILD_PIPELINE_STAGE": OperationTypeCreateBuildPipelineStage,
	"UPDATE_BUILD_PIPELINE_STAGE": OperationTypeUpdateBuildPipelineStage,
	"DELETE_BUILD_PIPELINE_STAGE": OperationTypeDeleteBuildPipelineStage,
	"CREATE_CONNECTION":           OperationTypeCreateConnection,
	"UPDATE_CONNECTION":           OperationTypeUpdateConnection,
	"DELETE_CONNECTION":           OperationTypeDeleteConnection,
	"CREATE_TRIGGER":              OperationTypeCreateTrigger,
	"UPDATE_TRIGGER":              OperationTypeUpdateTrigger,
	"DELETE_TRIGGER":              OperationTypeDeleteTrigger,
	"EXECUTE_TRIGGER":             OperationTypeExecuteTrigger,
	"CREATE_REPOSITORY":           OperationTypeCreateRepository,
	"UPDATE_REPOSITORY":           OperationTypeUpdateRepository,
	"DELETE_REPOSITORY":           OperationTypeDeleteRepository,
	"MIRROR_REPOSITORY":           OperationTypeMirrorRepository,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationType {
		values = append(values, v)
	}
	return values
}
