// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_service_catalog "github.com/oracle/oci-go-sdk/v50/servicecatalog"
)

func init() {
	RegisterDatasource("oci_service_catalog_service_catalog_associations", ServiceCatalogServiceCatalogAssociationsDataSource())
}

func ServiceCatalogServiceCatalogAssociationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServiceCatalogServiceCatalogAssociations,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_catalog_association_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_catalog_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_catalog_association_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(ServiceCatalogServiceCatalogAssociationResource()),
						},
					},
				},
			},
		},
	}
}

func readServiceCatalogServiceCatalogAssociations(d *schema.ResourceData, m interface{}) error {
	sync := &ServiceCatalogServiceCatalogAssociationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).serviceCatalogClient()

	return ReadResource(sync)
}

type ServiceCatalogServiceCatalogAssociationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_service_catalog.ServiceCatalogClient
	Res    *oci_service_catalog.ListServiceCatalogAssociationsResponse
}

func (s *ServiceCatalogServiceCatalogAssociationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServiceCatalogServiceCatalogAssociationsDataSourceCrud) Get() error {
	request := oci_service_catalog.ListServiceCatalogAssociationsRequest{}

	if entityId, ok := s.D.GetOkExists("entity_id"); ok {
		tmp := entityId.(string)
		request.EntityId = &tmp
	}

	if entityType, ok := s.D.GetOkExists("entity_type"); ok {
		tmp := entityType.(string)
		request.EntityType = &tmp
	}

	if serviceCatalogAssociationId, ok := s.D.GetOkExists("id"); ok {
		tmp := serviceCatalogAssociationId.(string)
		request.ServiceCatalogAssociationId = &tmp
	}

	if serviceCatalogId, ok := s.D.GetOkExists("service_catalog_id"); ok {
		tmp := serviceCatalogId.(string)
		request.ServiceCatalogId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "service_catalog")

	response, err := s.Client.ListServiceCatalogAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServiceCatalogAssociations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServiceCatalogServiceCatalogAssociationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("ServiceCatalogServiceCatalogAssociationsDataSource-", ServiceCatalogServiceCatalogAssociationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	serviceCatalogAssociation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ServiceCatalogAssociationSummaryToMap(item))
	}
	serviceCatalogAssociation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, ServiceCatalogServiceCatalogAssociationsDataSource().Schema["service_catalog_association_collection"].Elem.(*schema.Resource).Schema)
		serviceCatalogAssociation["items"] = items
	}

	resources = append(resources, serviceCatalogAssociation)
	if err := s.D.Set("service_catalog_association_collection", resources); err != nil {
		return err
	}

	return nil
}
