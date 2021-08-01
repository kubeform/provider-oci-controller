// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_apm "github.com/oracle/oci-go-sdk/v45/apmcontrolplane"

	oci_common "github.com/oracle/oci-go-sdk/v45/common"
)

func init() {
	RegisterOracleClient("oci_apm.ApmDomainClient", &OracleClient{initClientFn: initApmcontrolplaneApmDomainClient})
}

func initApmcontrolplaneApmDomainClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apm.NewApmDomainClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) apmDomainClient() *oci_apm.ApmDomainClient {
	return m.GetClient("oci_apm.ApmDomainClient").(*oci_apm.ApmDomainClient)
}
