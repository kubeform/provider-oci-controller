// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v50/devops"
)

func init() {
	RegisterDatasource("oci_devops_deploy_environments", DevopsDeployEnvironmentsDataSource())
}

func DevopsDeployEnvironmentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsDeployEnvironments,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deploy_environment_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(DevopsDeployEnvironmentResource()),
						},
					},
				},
			},
		},
	}
}

func readDevopsDeployEnvironments(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployEnvironmentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).devopsClient()

	return ReadResource(sync)
}

type DevopsDeployEnvironmentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListDeployEnvironmentsResponse
}

func (s *DevopsDeployEnvironmentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsDeployEnvironmentsDataSourceCrud) Get() error {
	request := oci_devops.ListDeployEnvironmentsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_devops.DeployEnvironmentLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "devops")

	response, err := s.Client.ListDeployEnvironments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDeployEnvironments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsDeployEnvironmentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DevopsDeployEnvironmentsDataSource-", DevopsDeployEnvironmentsDataSource(), s.D))
	resources := []map[string]interface{}{}
	deployEnvironment := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DeployEnvironmentSummaryToMap(item))
	}
	deployEnvironment["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, DevopsDeployEnvironmentsDataSource().Schema["deploy_environment_collection"].Elem.(*schema.Resource).Schema)
		deployEnvironment["items"] = items
	}

	resources = append(resources, deployEnvironment)
	if err := s.D.Set("deploy_environment_collection", resources); err != nil {
		return err
	}

	return nil
}
