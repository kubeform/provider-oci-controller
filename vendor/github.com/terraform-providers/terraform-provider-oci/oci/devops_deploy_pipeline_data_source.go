// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v45/devops"
)

func init() {
	RegisterDatasource("oci_devops_deploy_pipeline", DevopsDeployPipelineDataSource())
}

func DevopsDeployPipelineDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["deploy_pipeline_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DevopsDeployPipelineResource(), fieldMap, readSingularDevopsDeployPipeline)
}

func readSingularDevopsDeployPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployPipelineDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).devopsClient()

	return ReadResource(sync)
}

type DevopsDeployPipelineDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetDeployPipelineResponse
}

func (s *DevopsDeployPipelineDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsDeployPipelineDataSourceCrud) Get() error {
	request := oci_devops.GetDeployPipelineRequest{}

	if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
		tmp := deployPipelineId.(string)
		request.DeployPipelineId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "devops")

	response, err := s.Client.GetDeployPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsDeployPipelineDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DeployPipelineArtifacts != nil {
		s.D.Set("deploy_pipeline_artifacts", []interface{}{DeployPipelineArtifactCollectionToMap(s.Res.DeployPipelineArtifacts)})
	} else {
		s.D.Set("deploy_pipeline_artifacts", nil)
	}

	if s.Res.DeployPipelineEnvironments != nil {
		s.D.Set("deploy_pipeline_environments", []interface{}{DeployPipelineEnvironmentCollectionToMap(s.Res.DeployPipelineEnvironments)})
	} else {
		s.D.Set("deploy_pipeline_environments", nil)
	}

	if s.Res.DeployPipelineParameters != nil {
		s.D.Set("deploy_pipeline_parameters", []interface{}{DeployPipelineParameterCollectionToMap(s.Res.DeployPipelineParameters)})
	} else {
		s.D.Set("deploy_pipeline_parameters", nil)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
