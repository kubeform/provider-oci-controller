// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v50/networkloadbalancer"
)

func init() {
	RegisterDatasource("oci_network_load_balancer_network_load_balancer", NetworkLoadBalancerNetworkLoadBalancerDataSource())
}

func NetworkLoadBalancerNetworkLoadBalancerDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["network_load_balancer_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(NetworkLoadBalancerNetworkLoadBalancerResource(), fieldMap, readSingularNetworkLoadBalancerNetworkLoadBalancer)
}

func readSingularNetworkLoadBalancerNetworkLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).networkLoadBalancerClient()

	return ReadResource(sync)
}

type NetworkLoadBalancerNetworkLoadBalancerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.GetNetworkLoadBalancerResponse
}

func (s *NetworkLoadBalancerNetworkLoadBalancerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerNetworkLoadBalancerDataSourceCrud) Get() error {
	request := oci_network_load_balancer.GetNetworkLoadBalancerRequest{}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = GetRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.GetNetworkLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkLoadBalancerNetworkLoadBalancerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	ipAddresses := []interface{}{}
	for _, item := range s.Res.IpAddresses {
		ipAddresses = append(ipAddresses, NetworkLoadBalancerIpAddressToMap(item))
	}
	s.D.Set("ip_addresses", ipAddresses)

	if s.Res.IsPreserveSourceDestination != nil {
		s.D.Set("is_preserve_source_destination", *s.Res.IsPreserveSourceDestination)
	}

	if s.Res.IsPrivate != nil {
		s.D.Set("is_private", *s.Res.IsPrivate)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("network_security_group_ids", s.Res.NetworkSecurityGroupIds)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", systemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
