// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v50/loganalytics"
)

func init() {
	RegisterDatasource("oci_log_analytics_namespace_scheduled_tasks", LogAnalyticsNamespaceScheduledTasksDataSource())
}

func LogAnalyticsNamespaceScheduledTasksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLogAnalyticsNamespaceScheduledTasks,
		Schema: map[string]*schema.Schema{
			"filter": DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scheduled_task_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(LogAnalyticsNamespaceScheduledTaskResource()),
						},
					},
				},
			},
		},
	}
}

func readLogAnalyticsNamespaceScheduledTasks(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceScheduledTasksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).logAnalyticsClient()

	return ReadResource(sync)
}

type LogAnalyticsNamespaceScheduledTasksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.ListScheduledTasksResponse
}

func (s *LogAnalyticsNamespaceScheduledTasksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceScheduledTasksDataSourceCrud) Get() error {
	request := oci_log_analytics.ListScheduledTasksRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	if taskType, ok := s.D.GetOkExists("task_type"); ok {
		request.TaskType = oci_log_analytics.ListScheduledTasksTaskTypeEnum(taskType.(string))
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.ListScheduledTasks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListScheduledTasks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LogAnalyticsNamespaceScheduledTasksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LogAnalyticsNamespaceScheduledTasksDataSource-", LogAnalyticsNamespaceScheduledTasksDataSource(), s.D))
	resources := []map[string]interface{}{}
	namespaceScheduledTask := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ScheduledTaskSummaryToMap(item))
	}
	namespaceScheduledTask["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, LogAnalyticsNamespaceScheduledTasksDataSource().Schema["scheduled_task_collection"].Elem.(*schema.Resource).Schema)
		namespaceScheduledTask["items"] = items
	}

	resources = append(resources, namespaceScheduledTask)
	if err := s.D.Set("scheduled_task_collection", resources); err != nil {
		return err
	}

	return nil
}
