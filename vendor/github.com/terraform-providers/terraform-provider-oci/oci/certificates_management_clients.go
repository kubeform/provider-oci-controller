// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_certificates_management "github.com/oracle/oci-go-sdk/v50/certificatesmanagement"

	oci_common "github.com/oracle/oci-go-sdk/v50/common"
)

func init() {
	RegisterOracleClient("oci_certificates_management.CertificatesManagementClient", &OracleClient{InitClientFn: initCertificatesmanagementCertificatesManagementClient})
}

func initCertificatesmanagementCertificatesManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_certificates_management.NewCertificatesManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) certificatesManagementClient() *oci_certificates_management.CertificatesManagementClient {
	return m.GetClient("oci_certificates_management.CertificatesManagementClient").(*oci_certificates_management.CertificatesManagementClient)
}
