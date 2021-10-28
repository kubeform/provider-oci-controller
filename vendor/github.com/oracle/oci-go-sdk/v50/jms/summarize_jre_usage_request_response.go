// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// SummarizeJreUsageRequest wrapper for the SummarizeJreUsage operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/SummarizeJreUsage.go.html to see an example of how to use SummarizeJreUsageRequest.
type SummarizeJreUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// The vendor of the Java Runtime.
	JreVendor *string `mandatory:"false" contributesTo:"query" name:"jreVendor"`

	// The distribution of the Java Runtime.
	JreDistribution *string `mandatory:"false" contributesTo:"query" name:"jreDistribution"`

	// The version of the Java Runtime.
	JreVersion *string `mandatory:"false" contributesTo:"query" name:"jreVersion"`

	// The Fleet-unique identifier of the related application.
	ApplicationId *string `mandatory:"false" contributesTo:"query" name:"applicationId"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// Additional fields to include into the returned model on top of the required ones.
	// This parameter can also include 'approximateApplicationCount', 'approximateInstallationCount' and 'approximateManagedInstanceCount'.
	// For example 'approximateApplicationCount,approximateManagedInstanceCount'.
	Fields []SummarizeJreUsageFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The start of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end of the time period during which resources are searched (formatted according to RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339)).
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder SummarizeJreUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort JRE usages. Only one sort order may be provided.
	// Default order for _timeFirstSeen_, _timeLastSeen_, and _version_ is **descending**.
	// Default order for _timeFirstSeen_, _timeLastSeen_, _version_, _approximateInstallationCount_,
	// _approximateApplicationCount_ and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _distribution_, _vendor_, and _osName_ is **ascending**.
	// If no value is specified _timeLastSeen_ is default.
	SortBy SummarizeJreUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The operating system type.
	OsFamily []OsFamilyEnum `contributesTo:"query" name:"osFamily" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeJreUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeJreUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeJreUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeJreUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeJreUsageResponse wrapper for the SummarizeJreUsage operation
type SummarizeJreUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JreUsageCollection instances
	JreUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeJreUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeJreUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeJreUsageSortOrderEnum Enum with underlying type: string
type SummarizeJreUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeJreUsageSortOrderEnum
const (
	SummarizeJreUsageSortOrderAsc  SummarizeJreUsageSortOrderEnum = "ASC"
	SummarizeJreUsageSortOrderDesc SummarizeJreUsageSortOrderEnum = "DESC"
)

var mappingSummarizeJreUsageSortOrder = map[string]SummarizeJreUsageSortOrderEnum{
	"ASC":  SummarizeJreUsageSortOrderAsc,
	"DESC": SummarizeJreUsageSortOrderDesc,
}

// GetSummarizeJreUsageSortOrderEnumValues Enumerates the set of values for SummarizeJreUsageSortOrderEnum
func GetSummarizeJreUsageSortOrderEnumValues() []SummarizeJreUsageSortOrderEnum {
	values := make([]SummarizeJreUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeJreUsageSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeJreUsageSortByEnum Enum with underlying type: string
type SummarizeJreUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeJreUsageSortByEnum
const (
	SummarizeJreUsageSortByDistribution                    SummarizeJreUsageSortByEnum = "distribution"
	SummarizeJreUsageSortByTimefirstseen                   SummarizeJreUsageSortByEnum = "timeFirstSeen"
	SummarizeJreUsageSortByTimelastseen                    SummarizeJreUsageSortByEnum = "timeLastSeen"
	SummarizeJreUsageSortByVendor                          SummarizeJreUsageSortByEnum = "vendor"
	SummarizeJreUsageSortByVersion                         SummarizeJreUsageSortByEnum = "version"
	SummarizeJreUsageSortByApproximateinstallationcount    SummarizeJreUsageSortByEnum = "approximateInstallationCount"
	SummarizeJreUsageSortByApproximateapplicationcount     SummarizeJreUsageSortByEnum = "approximateApplicationCount"
	SummarizeJreUsageSortByApproximatemanagedinstancecount SummarizeJreUsageSortByEnum = "approximateManagedInstanceCount"
	SummarizeJreUsageSortByOsname                          SummarizeJreUsageSortByEnum = "osName"
)

var mappingSummarizeJreUsageSortBy = map[string]SummarizeJreUsageSortByEnum{
	"distribution":                    SummarizeJreUsageSortByDistribution,
	"timeFirstSeen":                   SummarizeJreUsageSortByTimefirstseen,
	"timeLastSeen":                    SummarizeJreUsageSortByTimelastseen,
	"vendor":                          SummarizeJreUsageSortByVendor,
	"version":                         SummarizeJreUsageSortByVersion,
	"approximateInstallationCount":    SummarizeJreUsageSortByApproximateinstallationcount,
	"approximateApplicationCount":     SummarizeJreUsageSortByApproximateapplicationcount,
	"approximateManagedInstanceCount": SummarizeJreUsageSortByApproximatemanagedinstancecount,
	"osName":                          SummarizeJreUsageSortByOsname,
}

// GetSummarizeJreUsageSortByEnumValues Enumerates the set of values for SummarizeJreUsageSortByEnum
func GetSummarizeJreUsageSortByEnumValues() []SummarizeJreUsageSortByEnum {
	values := make([]SummarizeJreUsageSortByEnum, 0)
	for _, v := range mappingSummarizeJreUsageSortBy {
		values = append(values, v)
	}
	return values
}
