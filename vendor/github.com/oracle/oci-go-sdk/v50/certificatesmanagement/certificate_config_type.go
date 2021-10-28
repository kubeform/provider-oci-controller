// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Certificates Service Management API
//
// API for managing certificates.
//

package certificatesmanagement

// CertificateConfigTypeEnum Enum with underlying type: string
type CertificateConfigTypeEnum string

// Set of constants representing the allowable values for CertificateConfigTypeEnum
const (
	CertificateConfigTypeIssuedByInternalCa                  CertificateConfigTypeEnum = "ISSUED_BY_INTERNAL_CA"
	CertificateConfigTypeManagedExternallyIssuedByInternalCa CertificateConfigTypeEnum = "MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA"
	CertificateConfigTypeImported                            CertificateConfigTypeEnum = "IMPORTED"
)

var mappingCertificateConfigType = map[string]CertificateConfigTypeEnum{
	"ISSUED_BY_INTERNAL_CA":                    CertificateConfigTypeIssuedByInternalCa,
	"MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA": CertificateConfigTypeManagedExternallyIssuedByInternalCa,
	"IMPORTED": CertificateConfigTypeImported,
}

// GetCertificateConfigTypeEnumValues Enumerates the set of values for CertificateConfigTypeEnum
func GetCertificateConfigTypeEnumValues() []CertificateConfigTypeEnum {
	values := make([]CertificateConfigTypeEnum, 0)
	for _, v := range mappingCertificateConfigType {
		values = append(values, v)
	}
	return values
}
