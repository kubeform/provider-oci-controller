// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v50/bds"
)

func init() {
	RegisterDatasource("oci_bds_auto_scaling_configuration", BdsAutoScalingConfigurationDataSource())
}

func BdsAutoScalingConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["auto_scaling_configuration_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["bds_instance_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(BdsAutoScalingConfigurationResource(), fieldMap, readSingularBdsAutoScalingConfiguration)
}

func readSingularBdsAutoScalingConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &BdsAutoScalingConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).bdsClient()

	return ReadResource(sync)
}

type BdsAutoScalingConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bds.BdsClient
	Res    *oci_bds.GetAutoScalingConfigurationResponse
}

func (s *BdsAutoScalingConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BdsAutoScalingConfigurationDataSourceCrud) Get() error {
	request := oci_bds.GetAutoScalingConfigurationRequest{}

	if autoScalingConfigurationId, ok := s.D.GetOkExists("auto_scaling_configuration_id"); ok {
		tmp := autoScalingConfigurationId.(string)
		request.AutoScalingConfigurationId = &tmp
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "bds")

	response, err := s.Client.GetAutoScalingConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BdsAutoScalingConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("node_type", s.Res.NodeType)

	if s.Res.Policy != nil {
		s.D.Set("policy", []interface{}{AutoScalePolicyToMap(s.Res.Policy)})
	} else {
		s.D.Set("policy", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}