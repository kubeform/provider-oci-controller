// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v45/loganalytics"
)

func init() {
	RegisterDatasource("oci_log_analytics_namespace", LogAnalyticsNamespaceDataSource())
}

func LogAnalyticsNamespaceDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLogAnalyticsNamespace,
		Schema: map[string]*schema.Schema{
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_onboarded": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func readSingularLogAnalyticsNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).logAnalyticsClient()

	return ReadResource(sync)
}

type LogAnalyticsNamespaceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetNamespaceResponse
}

func (s *LogAnalyticsNamespaceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceDataSourceCrud) Get() error {
	request := oci_log_analytics.GetNamespaceRequest{}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LogAnalyticsNamespaceDataSource-", LogAnalyticsNamespaceDataSource(), s.D))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.IsOnboarded != nil {
		s.D.Set("is_onboarded", *s.Res.IsOnboarded)
	}

	return nil
}
