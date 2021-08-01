// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v45/common"
	"net/http"
)

// SummarizeDatabaseInsightResourceCapacityTrendRequest wrapper for the SummarizeDatabaseInsightResourceCapacityTrend operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceCapacityTrendRequest.
type SummarizeDatabaseInsightResourceCapacityTrendRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by resource metric.
	// Supported values are CPU , STORAGE, MEMORY and IO.
	ResourceMetric *string `mandatory:"true" contributesTo:"query" name:"resourceMetric"`

	// Specify time period in ISO 8601 format with respect to current time.
	// Default is last 30 days represented by P30D.
	// If timeInterval is specified, then timeIntervalStart and timeIntervalEnd will be ignored.
	// Examples  P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months), . Maximum value allowed is 25 months prior to current time (P25M).
	AnalysisTimeInterval *string `mandatory:"false" contributesTo:"query" name:"analysisTimeInterval"`

	// Analysis start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// The minimum allowed value is 2 years prior to the current day.
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	// If analysisTimeInterval is specified, this parameter is ignored.
	TimeIntervalStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalStart"`

	// Analysis end time in UTC in ISO 8601 format(exclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd are used together.
	// If timeIntervalEnd is not specified, current time is used as timeIntervalEnd.
	TimeIntervalEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalEnd"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	DatabaseType []SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Filter by utilization level by the following buckets:
	//   - HIGH_UTILIZATION: DBs with utilization greater or equal than 75.
	//   - LOW_UTILIZATION: DBs with utilization lower than 25.
	//   - MEDIUM_HIGH_UTILIZATION: DBs with utilization greater or equal than 50 but lower than 75.
	//   - MEDIUM_LOW_UTILIZATION: DBs with utilization greater or equal than 25 but lower than 50.
	UtilizationLevel SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum `mandatory:"false" contributesTo:"query" name:"utilizationLevel" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sorts using end timestamp , capacity or baseCapacity
	SortBy SummarizeDatabaseInsightResourceCapacityTrendSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Tablespace name for a database
	TablespaceName *string `mandatory:"false" contributesTo:"query" name:"tablespaceName"`

	// Filter by one or more hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Flag to indicate if database instance level metrics should be returned. The flag is ignored when a host name filter is not applied.
	// When a hostname filter is applied this flag will determine whether to return metrics for the instances located on the specified host or for the
	// whole database which contains an instance on this host.
	IsDatabaseInstanceLevelMetrics *bool `mandatory:"false" contributesTo:"query" name:"isDatabaseInstanceLevelMetrics"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeDatabaseInsightResourceCapacityTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceCapacityTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceCapacityTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeDatabaseInsightResourceCapacityTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeDatabaseInsightResourceCapacityTrendResponse wrapper for the SummarizeDatabaseInsightResourceCapacityTrend operation
type SummarizeDatabaseInsightResourceCapacityTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection instances
	SummarizeDatabaseInsightResourceCapacityTrendAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeDatabaseInsightResourceCapacityTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeDatabaseInsightResourceCapacityTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwS           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "ADW-S"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpS           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "ATP-S"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwD           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "ADW-D"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpD           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "ATP-D"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalPdb    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "EXTERNAL-PDB"
	SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalNoncdb SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum = "EXTERNAL-NONCDB"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendDatabaseType = map[string]SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum{
	"ADW-S":           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwS,
	"ATP-S":           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpS,
	"ADW-D":           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAdwD,
	"ATP-D":           SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeAtpD,
	"EXTERNAL-PDB":    SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB": SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeExternalNoncdb,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendDatabaseType {
		values = append(values, v)
	}
	return values
}

// SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelHighUtilization       SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum = "HIGH_UTILIZATION"
	SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelLowUtilization        SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum = "LOW_UTILIZATION"
	SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum = "MEDIUM_HIGH_UTILIZATION"
	SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization  SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum = "MEDIUM_LOW_UTILIZATION"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevel = map[string]SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum{
	"HIGH_UTILIZATION":        SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelHighUtilization,
	"LOW_UTILIZATION":         SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelLowUtilization,
	"MEDIUM_HIGH_UTILIZATION": SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization,
	"MEDIUM_LOW_UTILIZATION":  SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendUtilizationLevelEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendUtilizationLevel {
		values = append(values, v)
	}
	return values
}

// SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendSortOrderAsc  SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum = "ASC"
	SummarizeDatabaseInsightResourceCapacityTrendSortOrderDesc SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum = "DESC"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendSortOrder = map[string]SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum{
	"ASC":  SummarizeDatabaseInsightResourceCapacityTrendSortOrderAsc,
	"DESC": SummarizeDatabaseInsightResourceCapacityTrendSortOrderDesc,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendSortOrderEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendSortOrderEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeDatabaseInsightResourceCapacityTrendSortByEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceCapacityTrendSortByEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceCapacityTrendSortByEnum
const (
	SummarizeDatabaseInsightResourceCapacityTrendSortByEndtimestamp SummarizeDatabaseInsightResourceCapacityTrendSortByEnum = "endTimestamp"
	SummarizeDatabaseInsightResourceCapacityTrendSortByCapacity     SummarizeDatabaseInsightResourceCapacityTrendSortByEnum = "capacity"
	SummarizeDatabaseInsightResourceCapacityTrendSortByBasecapacity SummarizeDatabaseInsightResourceCapacityTrendSortByEnum = "baseCapacity"
)

var mappingSummarizeDatabaseInsightResourceCapacityTrendSortBy = map[string]SummarizeDatabaseInsightResourceCapacityTrendSortByEnum{
	"endTimestamp": SummarizeDatabaseInsightResourceCapacityTrendSortByEndtimestamp,
	"capacity":     SummarizeDatabaseInsightResourceCapacityTrendSortByCapacity,
	"baseCapacity": SummarizeDatabaseInsightResourceCapacityTrendSortByBasecapacity,
}

// GetSummarizeDatabaseInsightResourceCapacityTrendSortByEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceCapacityTrendSortByEnum
func GetSummarizeDatabaseInsightResourceCapacityTrendSortByEnumValues() []SummarizeDatabaseInsightResourceCapacityTrendSortByEnum {
	values := make([]SummarizeDatabaseInsightResourceCapacityTrendSortByEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceCapacityTrendSortBy {
		values = append(values, v)
	}
	return values
}
