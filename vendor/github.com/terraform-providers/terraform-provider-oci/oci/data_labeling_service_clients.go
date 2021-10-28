// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_data_labeling_service "github.com/oracle/oci-go-sdk/v50/datalabelingservice"

	oci_common "github.com/oracle/oci-go-sdk/v50/common"
)

func init() {
	RegisterOracleClient("oci_data_labeling_service.DataLabelingManagementClient", &OracleClient{InitClientFn: initDatalabelingserviceDataLabelingManagementClient})
}

func initDatalabelingserviceDataLabelingManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_data_labeling_service.NewDataLabelingManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) dataLabelingManagementClient() *oci_data_labeling_service.DataLabelingManagementClient {
	return m.GetClient("oci_data_labeling_service.DataLabelingManagementClient").(*oci_data_labeling_service.DataLabelingManagementClient)
}
