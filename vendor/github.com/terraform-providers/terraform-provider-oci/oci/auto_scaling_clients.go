// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_auto_scaling "github.com/oracle/oci-go-sdk/v45/autoscaling"

	oci_common "github.com/oracle/oci-go-sdk/v45/common"
)

func init() {
	RegisterOracleClient("oci_auto_scaling.AutoScalingClient", &OracleClient{initClientFn: initAutoscalingAutoScalingClient})
}

func initAutoscalingAutoScalingClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_auto_scaling.NewAutoScalingClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) autoScalingClient() *oci_auto_scaling.AutoScalingClient {
	return m.GetClient("oci_auto_scaling.AutoScalingClient").(*oci_auto_scaling.AutoScalingClient)
}
