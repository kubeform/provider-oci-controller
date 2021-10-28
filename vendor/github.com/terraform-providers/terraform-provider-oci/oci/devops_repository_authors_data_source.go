// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v50/devops"
)

func init() {
	RegisterDatasource("oci_devops_repository_authors", DevopsRepositoryAuthorsDataSource())
}

func DevopsRepositoryAuthorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsRepositoryAuthors,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"ref_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository_author_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed

									"author_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDevopsRepositoryAuthors(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryAuthorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).devopsClient()

	return ReadResource(sync)
}

type DevopsRepositoryAuthorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListAuthorsResponse
}

func (s *DevopsRepositoryAuthorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryAuthorsDataSourceCrud) Get() error {
	request := oci_devops.ListAuthorsRequest{}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "devops")

	response, err := s.Client.ListAuthors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAuthors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsRepositoryAuthorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DevopsRepositoryAuthorsDataSource-", DevopsRepositoryAuthorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	repositoryAuthor := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RepositoryAuthorSummaryToMap(item))
	}
	repositoryAuthor["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, DevopsRepositoryAuthorsDataSource().Schema["repository_author_collection"].Elem.(*schema.Resource).Schema)
		repositoryAuthor["items"] = items
	}

	resources = append(resources, repositoryAuthor)
	if err := s.D.Set("repository_author_collection", resources); err != nil {
		return err
	}

	return nil
}
