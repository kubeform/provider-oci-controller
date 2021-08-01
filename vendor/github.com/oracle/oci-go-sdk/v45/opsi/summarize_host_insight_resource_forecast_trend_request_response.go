// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v45/common"
	"net/http"
)

// SummarizeHostInsightResourceForecastTrendRequest wrapper for the SummarizeHostInsightResourceForecastTrend operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceForecastTrend.go.html to see an example of how to use SummarizeHostInsightResourceForecastTrendRequest.
type SummarizeHostInsightResourceForecastTrendRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by host resource metric.
	// Supported values are CPU, MEMORY, and LOGICAL_MEMORY.
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

	// Filter by one or more platform types.
	// Possible value is LINUX.
	PlatformType []SummarizeHostInsightResourceForecastTrendPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Choose the type of statistic metric data to be used for forecasting.
	Statistic SummarizeHostInsightResourceForecastTrendStatisticEnum `mandatory:"false" contributesTo:"query" name:"statistic" omitEmpty:"true"`

	// Number of days used for utilization forecast analysis.
	ForecastDays *int `mandatory:"false" contributesTo:"query" name:"forecastDays"`

	// Choose algorithm model for the forecasting.
	// Possible values:
	//   - LINEAR: Uses linear regression algorithm for forecasting.
	//   - ML_AUTO: Automatically detects best algorithm to use for forecasting.
	//   - ML_NO_AUTO: Automatically detects seasonality of the data for forecasting using linear or seasonal algorithm.
	ForecastModel SummarizeHostInsightResourceForecastTrendForecastModelEnum `mandatory:"false" contributesTo:"query" name:"forecastModel" omitEmpty:"true"`

	// Filter by utilization level by the following buckets:
	//   - HIGH_UTILIZATION: DBs with utilization greater or equal than 75.
	//   - LOW_UTILIZATION: DBs with utilization lower than 25.
	//   - MEDIUM_HIGH_UTILIZATION: DBs with utilization greater or equal than 50 but lower than 75.
	//   - MEDIUM_LOW_UTILIZATION: DBs with utilization greater or equal than 25 but lower than 50.
	UtilizationLevel SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum `mandatory:"false" contributesTo:"query" name:"utilizationLevel" omitEmpty:"true"`

	// This parameter is used to change data's confidence level, this data is ingested by the
	// forecast algorithm.
	// Confidence is the probability of an interval to contain the expected population parameter.
	// Manipulation of this value will lead to different results.
	// If not set, default confidence value is 95%.
	Confidence *int `mandatory:"false" contributesTo:"query" name:"confidence"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

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

func (request SummarizeHostInsightResourceForecastTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightResourceForecastTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightResourceForecastTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightResourceForecastTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeHostInsightResourceForecastTrendResponse wrapper for the SummarizeHostInsightResourceForecastTrend operation
type SummarizeHostInsightResourceForecastTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeHostInsightResourceForecastTrendAggregation instances
	SummarizeHostInsightResourceForecastTrendAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SummarizeHostInsightResourceForecastTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightResourceForecastTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightResourceForecastTrendPlatformTypeEnum Enum with underlying type: string
type SummarizeHostInsightResourceForecastTrendPlatformTypeEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceForecastTrendPlatformTypeEnum
const (
	SummarizeHostInsightResourceForecastTrendPlatformTypeLinux SummarizeHostInsightResourceForecastTrendPlatformTypeEnum = "LINUX"
)

var mappingSummarizeHostInsightResourceForecastTrendPlatformType = map[string]SummarizeHostInsightResourceForecastTrendPlatformTypeEnum{
	"LINUX": SummarizeHostInsightResourceForecastTrendPlatformTypeLinux,
}

// GetSummarizeHostInsightResourceForecastTrendPlatformTypeEnumValues Enumerates the set of values for SummarizeHostInsightResourceForecastTrendPlatformTypeEnum
func GetSummarizeHostInsightResourceForecastTrendPlatformTypeEnumValues() []SummarizeHostInsightResourceForecastTrendPlatformTypeEnum {
	values := make([]SummarizeHostInsightResourceForecastTrendPlatformTypeEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceForecastTrendPlatformType {
		values = append(values, v)
	}
	return values
}

// SummarizeHostInsightResourceForecastTrendStatisticEnum Enum with underlying type: string
type SummarizeHostInsightResourceForecastTrendStatisticEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceForecastTrendStatisticEnum
const (
	SummarizeHostInsightResourceForecastTrendStatisticAvg SummarizeHostInsightResourceForecastTrendStatisticEnum = "AVG"
	SummarizeHostInsightResourceForecastTrendStatisticMax SummarizeHostInsightResourceForecastTrendStatisticEnum = "MAX"
)

var mappingSummarizeHostInsightResourceForecastTrendStatistic = map[string]SummarizeHostInsightResourceForecastTrendStatisticEnum{
	"AVG": SummarizeHostInsightResourceForecastTrendStatisticAvg,
	"MAX": SummarizeHostInsightResourceForecastTrendStatisticMax,
}

// GetSummarizeHostInsightResourceForecastTrendStatisticEnumValues Enumerates the set of values for SummarizeHostInsightResourceForecastTrendStatisticEnum
func GetSummarizeHostInsightResourceForecastTrendStatisticEnumValues() []SummarizeHostInsightResourceForecastTrendStatisticEnum {
	values := make([]SummarizeHostInsightResourceForecastTrendStatisticEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceForecastTrendStatistic {
		values = append(values, v)
	}
	return values
}

// SummarizeHostInsightResourceForecastTrendForecastModelEnum Enum with underlying type: string
type SummarizeHostInsightResourceForecastTrendForecastModelEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceForecastTrendForecastModelEnum
const (
	SummarizeHostInsightResourceForecastTrendForecastModelLinear   SummarizeHostInsightResourceForecastTrendForecastModelEnum = "LINEAR"
	SummarizeHostInsightResourceForecastTrendForecastModelMlAuto   SummarizeHostInsightResourceForecastTrendForecastModelEnum = "ML_AUTO"
	SummarizeHostInsightResourceForecastTrendForecastModelMlNoAuto SummarizeHostInsightResourceForecastTrendForecastModelEnum = "ML_NO_AUTO"
)

var mappingSummarizeHostInsightResourceForecastTrendForecastModel = map[string]SummarizeHostInsightResourceForecastTrendForecastModelEnum{
	"LINEAR":     SummarizeHostInsightResourceForecastTrendForecastModelLinear,
	"ML_AUTO":    SummarizeHostInsightResourceForecastTrendForecastModelMlAuto,
	"ML_NO_AUTO": SummarizeHostInsightResourceForecastTrendForecastModelMlNoAuto,
}

// GetSummarizeHostInsightResourceForecastTrendForecastModelEnumValues Enumerates the set of values for SummarizeHostInsightResourceForecastTrendForecastModelEnum
func GetSummarizeHostInsightResourceForecastTrendForecastModelEnumValues() []SummarizeHostInsightResourceForecastTrendForecastModelEnum {
	values := make([]SummarizeHostInsightResourceForecastTrendForecastModelEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceForecastTrendForecastModel {
		values = append(values, v)
	}
	return values
}

// SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum Enum with underlying type: string
type SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum
const (
	SummarizeHostInsightResourceForecastTrendUtilizationLevelHighUtilization       SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum = "HIGH_UTILIZATION"
	SummarizeHostInsightResourceForecastTrendUtilizationLevelLowUtilization        SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum = "LOW_UTILIZATION"
	SummarizeHostInsightResourceForecastTrendUtilizationLevelMediumHighUtilization SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum = "MEDIUM_HIGH_UTILIZATION"
	SummarizeHostInsightResourceForecastTrendUtilizationLevelMediumLowUtilization  SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum = "MEDIUM_LOW_UTILIZATION"
)

var mappingSummarizeHostInsightResourceForecastTrendUtilizationLevel = map[string]SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum{
	"HIGH_UTILIZATION":        SummarizeHostInsightResourceForecastTrendUtilizationLevelHighUtilization,
	"LOW_UTILIZATION":         SummarizeHostInsightResourceForecastTrendUtilizationLevelLowUtilization,
	"MEDIUM_HIGH_UTILIZATION": SummarizeHostInsightResourceForecastTrendUtilizationLevelMediumHighUtilization,
	"MEDIUM_LOW_UTILIZATION":  SummarizeHostInsightResourceForecastTrendUtilizationLevelMediumLowUtilization,
}

// GetSummarizeHostInsightResourceForecastTrendUtilizationLevelEnumValues Enumerates the set of values for SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum
func GetSummarizeHostInsightResourceForecastTrendUtilizationLevelEnumValues() []SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum {
	values := make([]SummarizeHostInsightResourceForecastTrendUtilizationLevelEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceForecastTrendUtilizationLevel {
		values = append(values, v)
	}
	return values
}
