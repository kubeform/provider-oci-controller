// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v45/networkloadbalancer"
)

func init() {
	RegisterDatasource("oci_network_load_balancer_listeners", NetworkLoadBalancerListenersDataSource())
}

func NetworkLoadBalancerListenersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readNetworkLoadBalancerListeners,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"network_load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"listener_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     GetDataSourceItemSchema(NetworkLoadBalancerListenerResource()),
						},
					},
				},
			},
		},
	}
}

func readNetworkLoadBalancerListeners(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerListenersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).networkLoadBalancerClient()

	return ReadResource(sync)
}

type NetworkLoadBalancerListenersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.ListListenersResponse
}

func (s *NetworkLoadBalancerListenersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerListenersDataSourceCrud) Get() error {
	request := oci_network_load_balancer.ListListenersRequest{}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.ListListeners(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListListeners(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *NetworkLoadBalancerListenersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("NetworkLoadBalancerListenersDataSource-", NetworkLoadBalancerListenersDataSource(), s.D))
	resources := []map[string]interface{}{}
	listener := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NlbListenerSummaryToMap(item))
	}
	listener["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = ApplyFiltersInCollection(f.(*schema.Set), items, NetworkLoadBalancerListenersDataSource().Schema["listener_collection"].Elem.(*schema.Resource).Schema)
		listener["items"] = items
	}

	resources = append(resources, listener)
	if err := s.D.Set("listener_collection", resources); err != nil {
		return err
	}

	return nil
}
