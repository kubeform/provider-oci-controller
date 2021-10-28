// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v50/filestorage"
)

func init() {
	RegisterDatasource("oci_file_storage_export_sets", FileStorageExportSetsDataSource())
}

func FileStorageExportSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageExportSets,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"export_sets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(FileStorageExportSetResource()),
			},
		},
	}
}

func readFileStorageExportSets(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageExportSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).fileStorageClient()

	return ReadResource(sync)
}

type FileStorageExportSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListExportSetsResponse
}

func (s *FileStorageExportSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageExportSetsDataSourceCrud) Get() error {
	request := oci_file_storage.ListExportSetsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_file_storage.ListExportSetsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "file_storage")

	response, err := s.Client.ListExportSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExportSets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FileStorageExportSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("FileStorageExportSetsDataSource-", FileStorageExportSetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		exportSet := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DisplayName != nil {
			exportSet["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			exportSet["id"] = *r.Id
		}

		exportSet["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			exportSet["time_created"] = r.TimeCreated.String()
		}

		if r.VcnId != nil {
			exportSet["vcn_id"] = *r.VcnId
		}

		resources = append(resources, exportSet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, FileStorageExportSetsDataSource().Schema["export_sets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("export_sets", resources); err != nil {
		return err
	}

	return nil
}
