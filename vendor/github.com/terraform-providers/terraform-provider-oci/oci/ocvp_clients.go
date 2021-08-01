// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_ocvp "github.com/oracle/oci-go-sdk/v45/ocvp"

	oci_common "github.com/oracle/oci-go-sdk/v45/common"
)

func init() {
	RegisterOracleClient("oci_ocvp.EsxiHostClient", &OracleClient{initClientFn: initOcvpEsxiHostClient})
	RegisterOracleClient("oci_ocvp.WorkRequestClient", &OracleClient{initClientFn: initOcvpWorkRequestClient})
	RegisterOracleClient("oci_ocvp.SddcClient", &OracleClient{initClientFn: initOcvpSddcClient})
}

func initOcvpEsxiHostClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ocvp.NewEsxiHostClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) esxiHostClient() *oci_ocvp.EsxiHostClient {
	return m.GetClient("oci_ocvp.EsxiHostClient").(*oci_ocvp.EsxiHostClient)
}

func initOcvpWorkRequestClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ocvp.NewWorkRequestClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ocvpWorkRequestClient() *oci_ocvp.WorkRequestClient {
	return m.GetClient("oci_ocvp.WorkRequestClient").(*oci_ocvp.WorkRequestClient)
}

func initOcvpSddcClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_ocvp.NewSddcClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) sddcClient() *oci_ocvp.SddcClient {
	return m.GetClient("oci_ocvp.SddcClient").(*oci_ocvp.SddcClient)
}
