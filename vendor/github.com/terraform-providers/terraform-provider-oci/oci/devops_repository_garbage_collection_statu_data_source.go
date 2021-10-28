//// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
//// Licensed under the Mozilla Public License v2.0
//
package oci

//
//import (
//	"context"
//
//	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
//	oci_devops "github.com/oracle/oci-go-sdk/v50/devops"
//)
//
//func init() {
//	RegisterDatasource("oci_devops_repository_garbage_collection_statu", DevopsRepositoryGarbageCollectionStatuDataSource())
//}
//
//func DevopsRepositoryGarbageCollectionStatuDataSource() *schema.Resource {
//	return &schema.Resource{
//		Read: readSingularDevopsRepositoryGarbageCollectionStatu,
//		Schema: map[string]*schema.Schema{
//			"repository_id": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			// Computed
//			"status": {
//				Type:     schema.TypeString,
//				Computed: true,
//			},
//		},
//	}
//}
//
//func readSingularDevopsRepositoryGarbageCollectionStatu(d *schema.ResourceData, m interface{}) error {
//	sync := &DevopsRepositoryGarbageCollectionStatuDataSourceCrud{}
//	sync.D = d
//	sync.Client = m.(*OracleClients).devopsClient()
//
//	return ReadResource(sync)
//}
//
//type DevopsRepositoryGarbageCollectionStatuDataSourceCrud struct {
//	D      *schema.ResourceData
//	Client *oci_devops.DevopsClient
//	Res    *oci_devops.GetGarbageCollectionStatusResponse
//}
//
//func (s *DevopsRepositoryGarbageCollectionStatuDataSourceCrud) VoidState() {
//	s.D.SetId("")
//}
//
//func (s *DevopsRepositoryGarbageCollectionStatuDataSourceCrud) Get() error {
//	request := oci_devops.GetGarbageCollectionStatusRequest{}
//
//	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
//		tmp := repositoryId.(string)
//		request.RepositoryId = &tmp
//	}
//
//	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "devops")
//
//	response, err := s.Client.GetGarbageCollectionStatus(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response
//	return nil
//}
//
//func (s *DevopsRepositoryGarbageCollectionStatuDataSourceCrud) SetData() error {
//	if s.Res == nil {
//		return nil
//	}
//
//	s.D.SetId(GenerateDataSourceHashID("DevopsRepositoryGarbageCollectionStatuDataSource-", DevopsRepositoryGarbageCollectionStatuDataSource(), s.D))
//
//	s.D.Set("status", s.Res.Status)
//
//	return nil
//}
