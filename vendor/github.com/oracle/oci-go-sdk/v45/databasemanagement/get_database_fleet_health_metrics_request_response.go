// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v45/common"
	"net/http"
)

// GetDatabaseFleetHealthMetricsRequest wrapper for the GetDatabaseFleetHealthMetrics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetDatabaseFleetHealthMetrics.go.html to see an example of how to use GetDatabaseFleetHealthMetricsRequest.
type GetDatabaseFleetHealthMetricsRequest struct {

	// The baseline time for metrics comparison.
	CompareBaselineTime *string `mandatory:"true" contributesTo:"query" name:"compareBaselineTime"`

	// The target time for metrics comparison.
	CompareTargetTime *string `mandatory:"true" contributesTo:"query" name:"compareTargetTime"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	ManagedDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"managedDatabaseGroupId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The time window used for metrics comparison.
	CompareType GetDatabaseFleetHealthMetricsCompareTypeEnum `mandatory:"false" contributesTo:"query" name:"compareType" omitEmpty:"true"`

	// The filter used to retrieve a specific set of metrics by passing the desired metric names with a comma separator. Note that, by default, the service returns all supported metrics.
	FilterByMetricNames *string `mandatory:"false" contributesTo:"query" name:"filterByMetricNames"`

	// The filter used to filter the databases in the fleet by a specific Oracle Database type.
	FilterByDatabaseType *string `mandatory:"false" contributesTo:"query" name:"filterByDatabaseType"`

	// The filter used to filter the databases in the fleet by a specific Oracle Database subtype.
	FilterByDatabaseSubType *string `mandatory:"false" contributesTo:"query" name:"filterByDatabaseSubType"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDatabaseFleetHealthMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDatabaseFleetHealthMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDatabaseFleetHealthMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDatabaseFleetHealthMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDatabaseFleetHealthMetricsResponse wrapper for the GetDatabaseFleetHealthMetrics operation
type GetDatabaseFleetHealthMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DatabaseFleetHealthMetrics instance
	DatabaseFleetHealthMetrics `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDatabaseFleetHealthMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDatabaseFleetHealthMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDatabaseFleetHealthMetricsCompareTypeEnum Enum with underlying type: string
type GetDatabaseFleetHealthMetricsCompareTypeEnum string

// Set of constants representing the allowable values for GetDatabaseFleetHealthMetricsCompareTypeEnum
const (
	GetDatabaseFleetHealthMetricsCompareTypeHour GetDatabaseFleetHealthMetricsCompareTypeEnum = "HOUR"
	GetDatabaseFleetHealthMetricsCompareTypeDay  GetDatabaseFleetHealthMetricsCompareTypeEnum = "DAY"
)

var mappingGetDatabaseFleetHealthMetricsCompareType = map[string]GetDatabaseFleetHealthMetricsCompareTypeEnum{
	"HOUR": GetDatabaseFleetHealthMetricsCompareTypeHour,
	"DAY":  GetDatabaseFleetHealthMetricsCompareTypeDay,
}

// GetGetDatabaseFleetHealthMetricsCompareTypeEnumValues Enumerates the set of values for GetDatabaseFleetHealthMetricsCompareTypeEnum
func GetGetDatabaseFleetHealthMetricsCompareTypeEnumValues() []GetDatabaseFleetHealthMetricsCompareTypeEnum {
	values := make([]GetDatabaseFleetHealthMetricsCompareTypeEnum, 0)
	for _, v := range mappingGetDatabaseFleetHealthMetricsCompareType {
		values = append(values, v)
	}
	return values
}
