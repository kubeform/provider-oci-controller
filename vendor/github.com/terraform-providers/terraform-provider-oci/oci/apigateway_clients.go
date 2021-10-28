// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_apigateway "github.com/oracle/oci-go-sdk/v50/apigateway"

	oci_common "github.com/oracle/oci-go-sdk/v50/common"
)

func init() {
	RegisterOracleClient("oci_apigateway.ApiGatewayClient", &OracleClient{InitClientFn: initApigatewayApiGatewayClient})
	RegisterOracleClient("oci_apigateway.WorkRequestsClient", &OracleClient{InitClientFn: initApigatewayWorkRequestsClient})
	RegisterOracleClient("oci_apigateway.DeploymentClient", &OracleClient{InitClientFn: initApigatewayDeploymentClient})
	RegisterOracleClient("oci_apigateway.GatewayClient", &OracleClient{InitClientFn: initApigatewayGatewayClient})
}

func initApigatewayApiGatewayClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewApiGatewayClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) apiGatewayClient() *oci_apigateway.ApiGatewayClient {
	return m.GetClient("oci_apigateway.ApiGatewayClient").(*oci_apigateway.ApiGatewayClient)
}

func initApigatewayWorkRequestsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewWorkRequestsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) apigatewayWorkRequestsClient() *oci_apigateway.WorkRequestsClient {
	return m.GetClient("oci_apigateway.WorkRequestsClient").(*oci_apigateway.WorkRequestsClient)
}

func initApigatewayDeploymentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewDeploymentClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) deploymentClient() *oci_apigateway.DeploymentClient {
	return m.GetClient("oci_apigateway.DeploymentClient").(*oci_apigateway.DeploymentClient)
}

func initApigatewayGatewayClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewGatewayClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) gatewayClient() *oci_apigateway.GatewayClient {
	return m.GetClient("oci_apigateway.GatewayClient").(*oci_apigateway.GatewayClient)
}
