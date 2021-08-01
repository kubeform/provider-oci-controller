// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v45/marketplace"
)

func init() {
	RegisterDatasource("oci_marketplace_listing_taxes", MarketplaceListingTaxesDataSource())
}

func MarketplaceListingTaxesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplaceListingTaxes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"listing_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"taxes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMarketplaceListingTaxes(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceListingTaxesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).marketplaceClient()

	return ReadResource(sync)
}

type MarketplaceListingTaxesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListTaxesResponse
}

func (s *MarketplaceListingTaxesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceListingTaxesDataSourceCrud) Get() error {
	request := oci_marketplace.ListTaxesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if listingId, ok := s.D.GetOkExists("listing_id"); ok {
		tmp := listingId.(string)
		request.ListingId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "marketplace")

	response, err := s.Client.ListTaxes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MarketplaceListingTaxesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("MarketplaceListingTaxesDataSource-", MarketplaceListingTaxesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		listingTax := map[string]interface{}{}

		if r.Code != nil {
			listingTax["code"] = *r.Code
		}

		if r.Country != nil {
			listingTax["country"] = *r.Country
		}

		if r.Name != nil {
			listingTax["name"] = *r.Name
		}

		if r.Url != nil {
			listingTax["url"] = *r.Url
		}

		resources = append(resources, listingTax)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, MarketplaceListingTaxesDataSource().Schema["taxes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("taxes", resources); err != nil {
		return err
	}

	return nil
}
