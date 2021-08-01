// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// CreateOkeDeployStageDetails Specifies the Kubernetes cluster deployment stage.
type CreateOkeDeployStageDetails struct {

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"true" json:"deployStagePredecessorCollection"`

	// Kubernetes cluster environment OCID for deployment.
	OkeClusterDeployEnvironmentId *string `mandatory:"true" json:"okeClusterDeployEnvironmentId"`

	// List of Kubernetes manifest artifact OCIDs, the manifests should not include any job resource.
	KubernetesManifestDeployArtifactIds []string `mandatory:"true" json:"kubernetesManifestDeployArtifactIds"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Default namespace to be used for Kubernetes deployment when not specified in the manifest.
	Namespace *string `mandatory:"false" json:"namespace"`

	RollbackPolicy DeployStageRollbackPolicy `mandatory:"false" json:"rollbackPolicy"`
}

//GetDescription returns Description
func (m CreateOkeDeployStageDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m CreateOkeDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDeployPipelineId returns DeployPipelineId
func (m CreateOkeDeployStageDetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m CreateOkeDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateOkeDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateOkeDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateOkeDeployStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateOkeDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateOkeDeployStageDetails CreateOkeDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeCreateOkeDeployStageDetails
	}{
		"OKE_DEPLOYMENT",
		(MarshalTypeCreateOkeDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateOkeDeployStageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                         *string                           `json:"description"`
		DisplayName                         *string                           `json:"displayName"`
		FreeformTags                        map[string]string                 `json:"freeformTags"`
		DefinedTags                         map[string]map[string]interface{} `json:"definedTags"`
		Namespace                           *string                           `json:"namespace"`
		RollbackPolicy                      deploystagerollbackpolicy         `json:"rollbackPolicy"`
		DeployPipelineId                    *string                           `json:"deployPipelineId"`
		DeployStagePredecessorCollection    *DeployStagePredecessorCollection `json:"deployStagePredecessorCollection"`
		OkeClusterDeployEnvironmentId       *string                           `json:"okeClusterDeployEnvironmentId"`
		KubernetesManifestDeployArtifactIds []string                          `json:"kubernetesManifestDeployArtifactIds"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Namespace = model.Namespace

	nn, e = model.RollbackPolicy.UnmarshalPolymorphicJSON(model.RollbackPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RollbackPolicy = nn.(DeployStageRollbackPolicy)
	} else {
		m.RollbackPolicy = nil
	}

	m.DeployPipelineId = model.DeployPipelineId

	m.DeployStagePredecessorCollection = model.DeployStagePredecessorCollection

	m.OkeClusterDeployEnvironmentId = model.OkeClusterDeployEnvironmentId

	m.KubernetesManifestDeployArtifactIds = make([]string, len(model.KubernetesManifestDeployArtifactIds))
	for i, n := range model.KubernetesManifestDeployArtifactIds {
		m.KubernetesManifestDeployArtifactIds[i] = n
	}

	return
}
