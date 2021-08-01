// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// CreateAutonomousDatabaseBase Details to create an Oracle Autonomous Database.
// **Notes:**
// - To specify OCPU core count, you must use either `ocpuCount` or `cpuCoreCount`. You cannot use both parameters at the same time.
// - To specify a storage allocation, you must use  either `dataStorageSizeInGBs` or `dataStorageSizeInTBs`.
// - See the individual parameter discriptions for more information on the OCPU and storage value parameters.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateAutonomousDatabaseBase interface {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment of the Autonomous Database.
	GetCompartmentId() *string

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 14 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	GetDbName() *string

	// The number of OCPU cores to be made available to the database. For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `ocpuCount` parameter.
	GetCpuCoreCount() *int

	// The number of OCPU cores to be made available to the database.
	// The following points apply:
	// - For Autonomous Databases on dedicated Exadata infrastructure, to provision less than 1 core, enter a fractional value in an increment of 0.1. For example, you can provision 0.3 or 0.4 cores, but not 0.35 cores. (Note that fractional OCPU values are not supported for Autonomous Databasese on shared Exadata infrastructure.)
	// - To provision 1 or more cores, you must enter an integer between 1 and the maximum number of cores available for the infrastructure shape. For example, you can provision 2 cores or 3 cores, but not 2.5 cores. This applies to Autonomous Databases on both shared and dedicated Exadata infrastructure.
	// For Autonomous Databases on dedicated Exadata infrastructure, the maximum number of cores is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `cpuCoreCount` parameter.
	GetOcpuCount() *float32

	// The Autonomous Database workload type. The following values are valid:
	// - OLTP - indicates an Autonomous Transaction Processing database
	// - DW - indicates an Autonomous Data Warehouse database
	// - AJD - indicates an Autonomous JSON Database
	// - APEX - indicates an Autonomous Database with the Oracle APEX Application Development workload type.
	GetDbWorkload() CreateAutonomousDatabaseBaseDbWorkloadEnum

	// The size, in terabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. For Autonomous Databases on dedicated Exadata infrastructure, the maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Note:** This parameter cannot be used with the `dataStorageSizeInGBs` parameter.
	GetDataStorageSizeInTBs() *int

	// The size, in gigabytes, of the data volume that will be created and attached to the database. This storage can later be scaled up if needed. The maximum storage value is determined by the infrastructure shape. See Characteristics of Infrastructure Shapes (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/autonomous-database&id=ATPFG-GUID-B0F033C1-CC5A-42F0-B2E7-3CECFEDA1FD1) for shape details.
	// **Notes**
	// - This parameter is only supported for dedicated Exadata infrastructure.
	// - This parameter cannot be used with the `dataStorageSizeInTBs` parameter.
	GetDataStorageSizeInGBs() *int

	// Indicates if this is an Always Free resource. The default value is false. Note that Always Free Autonomous Databases have 1 CPU and 20GB of memory. For Always Free databases, memory and CPU cannot be scaled.
	GetIsFreeTier() *bool

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	GetKmsKeyId() *string

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.cloud.oracle.com/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	GetVaultId() *string

	// The password must be between 12 and 30 characters long, and must contain at least 1 uppercase, 1 lowercase, and 1 numeric character. It cannot contain the double quote symbol (") or the username "admin", regardless of casing.
	GetAdminPassword() *string

	// The user-friendly name for the Autonomous Database. The name does not have to be unique.
	GetDisplayName() *string

	// The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle PaaS and IaaS services in the cloud.
	// License Included allows you to subscribe to new Oracle Database software licenses and the Database service.
	// Note that when provisioning an Autonomous Database on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm), this attribute must be null because the attribute is already set at the
	// Autonomous Exadata Infrastructure level. When using shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI), if a value is not specified, the system will supply the value of `BRING_YOUR_OWN_LICENSE`.
	GetLicenseModel() CreateAutonomousDatabaseBaseLicenseModelEnum

	// If set to `TRUE`, indicates that an Autonomous Database preview version is being provisioned, and that the preview version's terms of service have been accepted. Note that preview version software is only available for databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI).
	GetIsPreviewVersionWithServiceTermsAccepted() *bool

	// Indicates if auto scaling is enabled for the Autonomous Database OCPU core count. The default value is `FALSE`.
	GetIsAutoScalingEnabled() *bool

	// True if the database is on dedicated Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adbddoverview.htm).
	GetIsDedicated() *bool

	// The Autonomous Container Database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	GetAutonomousContainerDatabaseId() *string

	// Indicates if the database-level access control is enabled.
	// If disabled, database access is defined by the network security rules.
	// If enabled, database access is restricted to the IP addresses defined by the rules specified with the `whitelistedIps` property. While specifying `whitelistedIps` rules is optional,
	//  if database-level access control is enabled and no rules are specified, the database will become inaccessible. The rules can be added later using the `UpdateAutonomousDatabase` API operation or edit option in console.
	// When creating a database clone, the desired access control setting should be specified. By default, database-level access control will be disabled for the clone.
	// This property is applicable only to Autonomous Databases on the Exadata Cloud@Customer platform.
	GetIsAccessControlEnabled() *bool

	// The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	GetWhitelistedIps() []string

	// This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled.
	// It's value would be `TRUE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses primary IP access control list (ACL) for standby.
	// It's value would be `FALSE` if Autonomous Database is Data Guard enabled and Access Control is enabled and if the Autonomous Database uses different IP access control list (ACL) for standby compared to primary.
	GetArePrimaryWhitelistedIpsUsed() *bool

	// The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure (https://docs.cloud.oracle.com/Content/Database/Concepts/adboverview.htm#AEI) and on Exadata Cloud@Customer.
	// Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance.
	// For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.
	// Use a semicolon (;) as a deliminator between the VCN-specific subnets or IPs.
	// Example: `["1.1.1.1","1.1.1.0/24","ocid1.vcn.oc1.sea.<unique_id>","ocid1.vcn.oc1.sea.<unique_id1>;1.1.1.1","ocid1.vcn.oc1.sea.<unique_id2>;1.1.0.0/16"]`
	// For Exadata Cloud@Customer, this is an array of IP addresses or CIDR (Classless Inter-Domain Routing) notations.
	// Example: `["1.1.1.1","1.1.1.0/24","1.1.2.25"]`
	// For an update operation, if you want to delete all the IPs in the ACL, use an array with a single empty string entry.
	GetStandbyWhitelistedIps() []string

	// Indicates whether the Autonomous Database has Data Guard enabled.
	GetIsDataGuardEnabled() *bool

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the subnet the resource is associated with.
	// **Subnet Restrictions:**
	// - For bare metal DB systems and for single node virtual machine DB systems, do not use a subnet that overlaps with 192.168.16.16/28.
	// - For Exadata and virtual machine 2-node RAC systems, do not use a subnet that overlaps with 192.168.128.0/20.
	// - For Autonomous Database, setting this will disable public secure access to the database.
	// These subnets are used by the Oracle Clusterware private interconnect on the database instance.
	// Specifying an overlapping subnet will cause the private interconnect to malfunction.
	// This restriction applies to both the client subnet and the backup subnet.
	GetSubnetId() *string

	// A list of the OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the network security groups (NSGs) that this resource belongs to. Setting this to an empty array after the list is created removes the resource from all NSGs. For more information about NSGs, see Security Rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securityrules.htm).
	// **NsgIds restrictions:**
	// - Autonomous Databases with private access require at least 1 Network Security Group (NSG). The nsgIds array cannot be empty.
	GetNsgIds() []string

	// The private endpoint label for the resource. Setting this to an empty string, after the private endpoint database gets created, will change the same private endpoint database to the public endpoint database.
	GetPrivateEndpointLabel() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	GetDefinedTags() map[string]map[string]interface{}

	// A valid Oracle Database version for Autonomous Database.
	GetDbVersion() *string

	// Customer Contacts.
	GetCustomerContacts() []CustomerContact
}

type createautonomousdatabasebase struct {
	JsonData                                 []byte
	CompartmentId                            *string                                      `mandatory:"true" json:"compartmentId"`
	DbName                                   *string                                      `mandatory:"true" json:"dbName"`
	CpuCoreCount                             *int                                         `mandatory:"false" json:"cpuCoreCount"`
	OcpuCount                                *float32                                     `mandatory:"false" json:"ocpuCount"`
	DbWorkload                               CreateAutonomousDatabaseBaseDbWorkloadEnum   `mandatory:"false" json:"dbWorkload,omitempty"`
	DataStorageSizeInTBs                     *int                                         `mandatory:"false" json:"dataStorageSizeInTBs"`
	DataStorageSizeInGBs                     *int                                         `mandatory:"false" json:"dataStorageSizeInGBs"`
	IsFreeTier                               *bool                                        `mandatory:"false" json:"isFreeTier"`
	KmsKeyId                                 *string                                      `mandatory:"false" json:"kmsKeyId"`
	VaultId                                  *string                                      `mandatory:"false" json:"vaultId"`
	AdminPassword                            *string                                      `mandatory:"false" json:"adminPassword"`
	DisplayName                              *string                                      `mandatory:"false" json:"displayName"`
	LicenseModel                             CreateAutonomousDatabaseBaseLicenseModelEnum `mandatory:"false" json:"licenseModel,omitempty"`
	IsPreviewVersionWithServiceTermsAccepted *bool                                        `mandatory:"false" json:"isPreviewVersionWithServiceTermsAccepted"`
	IsAutoScalingEnabled                     *bool                                        `mandatory:"false" json:"isAutoScalingEnabled"`
	IsDedicated                              *bool                                        `mandatory:"false" json:"isDedicated"`
	AutonomousContainerDatabaseId            *string                                      `mandatory:"false" json:"autonomousContainerDatabaseId"`
	IsAccessControlEnabled                   *bool                                        `mandatory:"false" json:"isAccessControlEnabled"`
	WhitelistedIps                           []string                                     `mandatory:"false" json:"whitelistedIps"`
	ArePrimaryWhitelistedIpsUsed             *bool                                        `mandatory:"false" json:"arePrimaryWhitelistedIpsUsed"`
	StandbyWhitelistedIps                    []string                                     `mandatory:"false" json:"standbyWhitelistedIps"`
	IsDataGuardEnabled                       *bool                                        `mandatory:"false" json:"isDataGuardEnabled"`
	SubnetId                                 *string                                      `mandatory:"false" json:"subnetId"`
	NsgIds                                   []string                                     `mandatory:"false" json:"nsgIds"`
	PrivateEndpointLabel                     *string                                      `mandatory:"false" json:"privateEndpointLabel"`
	FreeformTags                             map[string]string                            `mandatory:"false" json:"freeformTags"`
	DefinedTags                              map[string]map[string]interface{}            `mandatory:"false" json:"definedTags"`
	DbVersion                                *string                                      `mandatory:"false" json:"dbVersion"`
	CustomerContacts                         []CustomerContact                            `mandatory:"false" json:"customerContacts"`
	Source                                   string                                       `json:"source"`
}

// UnmarshalJSON unmarshals json
func (m *createautonomousdatabasebase) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateautonomousdatabasebase createautonomousdatabasebase
	s := struct {
		Model Unmarshalercreateautonomousdatabasebase
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.DbName = s.Model.DbName
	m.CpuCoreCount = s.Model.CpuCoreCount
	m.OcpuCount = s.Model.OcpuCount
	m.DbWorkload = s.Model.DbWorkload
	m.DataStorageSizeInTBs = s.Model.DataStorageSizeInTBs
	m.DataStorageSizeInGBs = s.Model.DataStorageSizeInGBs
	m.IsFreeTier = s.Model.IsFreeTier
	m.KmsKeyId = s.Model.KmsKeyId
	m.VaultId = s.Model.VaultId
	m.AdminPassword = s.Model.AdminPassword
	m.DisplayName = s.Model.DisplayName
	m.LicenseModel = s.Model.LicenseModel
	m.IsPreviewVersionWithServiceTermsAccepted = s.Model.IsPreviewVersionWithServiceTermsAccepted
	m.IsAutoScalingEnabled = s.Model.IsAutoScalingEnabled
	m.IsDedicated = s.Model.IsDedicated
	m.AutonomousContainerDatabaseId = s.Model.AutonomousContainerDatabaseId
	m.IsAccessControlEnabled = s.Model.IsAccessControlEnabled
	m.WhitelistedIps = s.Model.WhitelistedIps
	m.ArePrimaryWhitelistedIpsUsed = s.Model.ArePrimaryWhitelistedIpsUsed
	m.StandbyWhitelistedIps = s.Model.StandbyWhitelistedIps
	m.IsDataGuardEnabled = s.Model.IsDataGuardEnabled
	m.SubnetId = s.Model.SubnetId
	m.NsgIds = s.Model.NsgIds
	m.PrivateEndpointLabel = s.Model.PrivateEndpointLabel
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DbVersion = s.Model.DbVersion
	m.CustomerContacts = s.Model.CustomerContacts
	m.Source = s.Model.Source

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createautonomousdatabasebase) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Source {
	case "DATABASE":
		mm := CreateAutonomousDatabaseCloneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CLONE_TO_REFRESHABLE":
		mm := CreateRefreshableAutonomousDatabaseCloneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BACKUP_FROM_ID":
		mm := CreateAutonomousDatabaseFromBackupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BACKUP_FROM_TIMESTAMP":
		mm := CreateAutonomousDatabaseFromBackupTimestampDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CROSS_REGION_DATAGUARD":
		mm := CreateCrossRegionAutonomousDatabaseDataGuardDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NONE":
		mm := CreateAutonomousDatabaseDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCompartmentId returns CompartmentId
func (m createautonomousdatabasebase) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetDbName returns DbName
func (m createautonomousdatabasebase) GetDbName() *string {
	return m.DbName
}

//GetCpuCoreCount returns CpuCoreCount
func (m createautonomousdatabasebase) GetCpuCoreCount() *int {
	return m.CpuCoreCount
}

//GetOcpuCount returns OcpuCount
func (m createautonomousdatabasebase) GetOcpuCount() *float32 {
	return m.OcpuCount
}

//GetDbWorkload returns DbWorkload
func (m createautonomousdatabasebase) GetDbWorkload() CreateAutonomousDatabaseBaseDbWorkloadEnum {
	return m.DbWorkload
}

//GetDataStorageSizeInTBs returns DataStorageSizeInTBs
func (m createautonomousdatabasebase) GetDataStorageSizeInTBs() *int {
	return m.DataStorageSizeInTBs
}

//GetDataStorageSizeInGBs returns DataStorageSizeInGBs
func (m createautonomousdatabasebase) GetDataStorageSizeInGBs() *int {
	return m.DataStorageSizeInGBs
}

//GetIsFreeTier returns IsFreeTier
func (m createautonomousdatabasebase) GetIsFreeTier() *bool {
	return m.IsFreeTier
}

//GetKmsKeyId returns KmsKeyId
func (m createautonomousdatabasebase) GetKmsKeyId() *string {
	return m.KmsKeyId
}

//GetVaultId returns VaultId
func (m createautonomousdatabasebase) GetVaultId() *string {
	return m.VaultId
}

//GetAdminPassword returns AdminPassword
func (m createautonomousdatabasebase) GetAdminPassword() *string {
	return m.AdminPassword
}

//GetDisplayName returns DisplayName
func (m createautonomousdatabasebase) GetDisplayName() *string {
	return m.DisplayName
}

//GetLicenseModel returns LicenseModel
func (m createautonomousdatabasebase) GetLicenseModel() CreateAutonomousDatabaseBaseLicenseModelEnum {
	return m.LicenseModel
}

//GetIsPreviewVersionWithServiceTermsAccepted returns IsPreviewVersionWithServiceTermsAccepted
func (m createautonomousdatabasebase) GetIsPreviewVersionWithServiceTermsAccepted() *bool {
	return m.IsPreviewVersionWithServiceTermsAccepted
}

//GetIsAutoScalingEnabled returns IsAutoScalingEnabled
func (m createautonomousdatabasebase) GetIsAutoScalingEnabled() *bool {
	return m.IsAutoScalingEnabled
}

//GetIsDedicated returns IsDedicated
func (m createautonomousdatabasebase) GetIsDedicated() *bool {
	return m.IsDedicated
}

//GetAutonomousContainerDatabaseId returns AutonomousContainerDatabaseId
func (m createautonomousdatabasebase) GetAutonomousContainerDatabaseId() *string {
	return m.AutonomousContainerDatabaseId
}

//GetIsAccessControlEnabled returns IsAccessControlEnabled
func (m createautonomousdatabasebase) GetIsAccessControlEnabled() *bool {
	return m.IsAccessControlEnabled
}

//GetWhitelistedIps returns WhitelistedIps
func (m createautonomousdatabasebase) GetWhitelistedIps() []string {
	return m.WhitelistedIps
}

//GetArePrimaryWhitelistedIpsUsed returns ArePrimaryWhitelistedIpsUsed
func (m createautonomousdatabasebase) GetArePrimaryWhitelistedIpsUsed() *bool {
	return m.ArePrimaryWhitelistedIpsUsed
}

//GetStandbyWhitelistedIps returns StandbyWhitelistedIps
func (m createautonomousdatabasebase) GetStandbyWhitelistedIps() []string {
	return m.StandbyWhitelistedIps
}

//GetIsDataGuardEnabled returns IsDataGuardEnabled
func (m createautonomousdatabasebase) GetIsDataGuardEnabled() *bool {
	return m.IsDataGuardEnabled
}

//GetSubnetId returns SubnetId
func (m createautonomousdatabasebase) GetSubnetId() *string {
	return m.SubnetId
}

//GetNsgIds returns NsgIds
func (m createautonomousdatabasebase) GetNsgIds() []string {
	return m.NsgIds
}

//GetPrivateEndpointLabel returns PrivateEndpointLabel
func (m createautonomousdatabasebase) GetPrivateEndpointLabel() *string {
	return m.PrivateEndpointLabel
}

//GetFreeformTags returns FreeformTags
func (m createautonomousdatabasebase) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m createautonomousdatabasebase) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetDbVersion returns DbVersion
func (m createautonomousdatabasebase) GetDbVersion() *string {
	return m.DbVersion
}

//GetCustomerContacts returns CustomerContacts
func (m createautonomousdatabasebase) GetCustomerContacts() []CustomerContact {
	return m.CustomerContacts
}

func (m createautonomousdatabasebase) String() string {
	return common.PointerString(m)
}

// CreateAutonomousDatabaseBaseDbWorkloadEnum Enum with underlying type: string
type CreateAutonomousDatabaseBaseDbWorkloadEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseBaseDbWorkloadEnum
const (
	CreateAutonomousDatabaseBaseDbWorkloadOltp CreateAutonomousDatabaseBaseDbWorkloadEnum = "OLTP"
	CreateAutonomousDatabaseBaseDbWorkloadDw   CreateAutonomousDatabaseBaseDbWorkloadEnum = "DW"
	CreateAutonomousDatabaseBaseDbWorkloadAjd  CreateAutonomousDatabaseBaseDbWorkloadEnum = "AJD"
	CreateAutonomousDatabaseBaseDbWorkloadApex CreateAutonomousDatabaseBaseDbWorkloadEnum = "APEX"
)

var mappingCreateAutonomousDatabaseBaseDbWorkload = map[string]CreateAutonomousDatabaseBaseDbWorkloadEnum{
	"OLTP": CreateAutonomousDatabaseBaseDbWorkloadOltp,
	"DW":   CreateAutonomousDatabaseBaseDbWorkloadDw,
	"AJD":  CreateAutonomousDatabaseBaseDbWorkloadAjd,
	"APEX": CreateAutonomousDatabaseBaseDbWorkloadApex,
}

// GetCreateAutonomousDatabaseBaseDbWorkloadEnumValues Enumerates the set of values for CreateAutonomousDatabaseBaseDbWorkloadEnum
func GetCreateAutonomousDatabaseBaseDbWorkloadEnumValues() []CreateAutonomousDatabaseBaseDbWorkloadEnum {
	values := make([]CreateAutonomousDatabaseBaseDbWorkloadEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseBaseDbWorkload {
		values = append(values, v)
	}
	return values
}

// CreateAutonomousDatabaseBaseLicenseModelEnum Enum with underlying type: string
type CreateAutonomousDatabaseBaseLicenseModelEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseBaseLicenseModelEnum
const (
	CreateAutonomousDatabaseBaseLicenseModelLicenseIncluded     CreateAutonomousDatabaseBaseLicenseModelEnum = "LICENSE_INCLUDED"
	CreateAutonomousDatabaseBaseLicenseModelBringYourOwnLicense CreateAutonomousDatabaseBaseLicenseModelEnum = "BRING_YOUR_OWN_LICENSE"
)

var mappingCreateAutonomousDatabaseBaseLicenseModel = map[string]CreateAutonomousDatabaseBaseLicenseModelEnum{
	"LICENSE_INCLUDED":       CreateAutonomousDatabaseBaseLicenseModelLicenseIncluded,
	"BRING_YOUR_OWN_LICENSE": CreateAutonomousDatabaseBaseLicenseModelBringYourOwnLicense,
}

// GetCreateAutonomousDatabaseBaseLicenseModelEnumValues Enumerates the set of values for CreateAutonomousDatabaseBaseLicenseModelEnum
func GetCreateAutonomousDatabaseBaseLicenseModelEnumValues() []CreateAutonomousDatabaseBaseLicenseModelEnum {
	values := make([]CreateAutonomousDatabaseBaseLicenseModelEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseBaseLicenseModel {
		values = append(values, v)
	}
	return values
}

// CreateAutonomousDatabaseBaseSourceEnum Enum with underlying type: string
type CreateAutonomousDatabaseBaseSourceEnum string

// Set of constants representing the allowable values for CreateAutonomousDatabaseBaseSourceEnum
const (
	CreateAutonomousDatabaseBaseSourceNone                 CreateAutonomousDatabaseBaseSourceEnum = "NONE"
	CreateAutonomousDatabaseBaseSourceDatabase             CreateAutonomousDatabaseBaseSourceEnum = "DATABASE"
	CreateAutonomousDatabaseBaseSourceBackupFromId         CreateAutonomousDatabaseBaseSourceEnum = "BACKUP_FROM_ID"
	CreateAutonomousDatabaseBaseSourceBackupFromTimestamp  CreateAutonomousDatabaseBaseSourceEnum = "BACKUP_FROM_TIMESTAMP"
	CreateAutonomousDatabaseBaseSourceCloneToRefreshable   CreateAutonomousDatabaseBaseSourceEnum = "CLONE_TO_REFRESHABLE"
	CreateAutonomousDatabaseBaseSourceCrossRegionDataguard CreateAutonomousDatabaseBaseSourceEnum = "CROSS_REGION_DATAGUARD"
)

var mappingCreateAutonomousDatabaseBaseSource = map[string]CreateAutonomousDatabaseBaseSourceEnum{
	"NONE":                   CreateAutonomousDatabaseBaseSourceNone,
	"DATABASE":               CreateAutonomousDatabaseBaseSourceDatabase,
	"BACKUP_FROM_ID":         CreateAutonomousDatabaseBaseSourceBackupFromId,
	"BACKUP_FROM_TIMESTAMP":  CreateAutonomousDatabaseBaseSourceBackupFromTimestamp,
	"CLONE_TO_REFRESHABLE":   CreateAutonomousDatabaseBaseSourceCloneToRefreshable,
	"CROSS_REGION_DATAGUARD": CreateAutonomousDatabaseBaseSourceCrossRegionDataguard,
}

// GetCreateAutonomousDatabaseBaseSourceEnumValues Enumerates the set of values for CreateAutonomousDatabaseBaseSourceEnum
func GetCreateAutonomousDatabaseBaseSourceEnumValues() []CreateAutonomousDatabaseBaseSourceEnum {
	values := make([]CreateAutonomousDatabaseBaseSourceEnum, 0)
	for _, v := range mappingCreateAutonomousDatabaseBaseSource {
		values = append(values, v)
	}
	return values
}
