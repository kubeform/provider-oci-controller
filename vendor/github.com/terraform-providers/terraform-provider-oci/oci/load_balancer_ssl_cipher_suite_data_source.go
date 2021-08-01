// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v45/loadbalancer"
)

func init() {
	RegisterDatasource("oci_load_balancer_ssl_cipher_suite", LoadBalancerSslCipherSuiteDataSource())
}

func LoadBalancerSslCipherSuiteDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["load_balancer_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(LoadBalancerSslCipherSuiteResource(), fieldMap, readSingularLoadBalancerSslCipherSuite)
}

func readSingularLoadBalancerSslCipherSuite(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerSslCipherSuiteDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return ReadResource(sync)
}

type LoadBalancerSslCipherSuiteDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetSSLCipherSuiteResponse
}

func (s *LoadBalancerSslCipherSuiteDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerSslCipherSuiteDataSourceCrud) Get() error {
	request := oci_load_balancer.GetSSLCipherSuiteRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetSSLCipherSuite(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerSslCipherSuiteDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LoadBalancerSslCipherSuiteDataSource-", LoadBalancerSslCipherSuiteDataSource(), s.D))

	s.D.Set("ciphers", s.Res.Ciphers)

	return nil
}
