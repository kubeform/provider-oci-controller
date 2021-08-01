// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v45/common"
	"net/http"
)

// GetAwrDbReportRequest wrapper for the GetAwrDbReport operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetAwrDbReport.go.html to see an example of how to use GetAwrDbReportRequest.
type GetAwrDbReportRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The parameter to filter the database by internal ID.
	// Note that the internal ID of the database can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbs:
	AwrDbId *string `mandatory:"true" contributesTo:"path" name:"awrDbId"`

	// The optional multiple value query parameter to filter the database instance numbers.
	InstNums []int `contributesTo:"query" name:"instNums" collectionFormat:"csv"`

	// The optional greater than or equal to filter on the snapshot ID.
	BeginSnIdGreaterThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"beginSnIdGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the snapshot ID.
	EndSnIdLessThanOrEqualTo *int `mandatory:"false" contributesTo:"query" name:"endSnIdLessThanOrEqualTo"`

	// The optional greater than or equal to query parameter to filter the timestamp.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The query parameter to filter the AWR report types.
	ReportType GetAwrDbReportReportTypeEnum `mandatory:"false" contributesTo:"query" name:"reportType" omitEmpty:"true"`

	// The optional query parameter to filter the database container by an exact ID value.
	// Note that the database container ID can be retrieved from the following endpoint:
	// /managedDatabases/{managedDatabaseId}/awrDbSnapshotRanges
	ContainerId *int `mandatory:"false" contributesTo:"query" name:"containerId"`

	// The format of the AWR report.
	ReportFormat GetAwrDbReportReportFormatEnum `mandatory:"false" contributesTo:"query" name:"reportFormat" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAwrDbReportRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAwrDbReportRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAwrDbReportRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAwrDbReportRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAwrDbReportResponse wrapper for the GetAwrDbReport operation
type GetAwrDbReportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AwrDbReport instance
	AwrDbReport `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAwrDbReportResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAwrDbReportResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAwrDbReportReportTypeEnum Enum with underlying type: string
type GetAwrDbReportReportTypeEnum string

// Set of constants representing the allowable values for GetAwrDbReportReportTypeEnum
const (
	GetAwrDbReportReportTypeAwr GetAwrDbReportReportTypeEnum = "AWR"
	GetAwrDbReportReportTypeAsh GetAwrDbReportReportTypeEnum = "ASH"
)

var mappingGetAwrDbReportReportType = map[string]GetAwrDbReportReportTypeEnum{
	"AWR": GetAwrDbReportReportTypeAwr,
	"ASH": GetAwrDbReportReportTypeAsh,
}

// GetGetAwrDbReportReportTypeEnumValues Enumerates the set of values for GetAwrDbReportReportTypeEnum
func GetGetAwrDbReportReportTypeEnumValues() []GetAwrDbReportReportTypeEnum {
	values := make([]GetAwrDbReportReportTypeEnum, 0)
	for _, v := range mappingGetAwrDbReportReportType {
		values = append(values, v)
	}
	return values
}

// GetAwrDbReportReportFormatEnum Enum with underlying type: string
type GetAwrDbReportReportFormatEnum string

// Set of constants representing the allowable values for GetAwrDbReportReportFormatEnum
const (
	GetAwrDbReportReportFormatHtml GetAwrDbReportReportFormatEnum = "HTML"
	GetAwrDbReportReportFormatText GetAwrDbReportReportFormatEnum = "TEXT"
)

var mappingGetAwrDbReportReportFormat = map[string]GetAwrDbReportReportFormatEnum{
	"HTML": GetAwrDbReportReportFormatHtml,
	"TEXT": GetAwrDbReportReportFormatText,
}

// GetGetAwrDbReportReportFormatEnumValues Enumerates the set of values for GetAwrDbReportReportFormatEnum
func GetGetAwrDbReportReportFormatEnumValues() []GetAwrDbReportReportFormatEnum {
	values := make([]GetAwrDbReportReportFormatEnum, 0)
	for _, v := range mappingGetAwrDbReportReportFormat {
		values = append(values, v)
	}
	return values
}
