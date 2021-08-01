// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v45/common"
	"net/http"
)

// ListWorkRequestLogsRequest wrapper for the ListWorkRequestLogs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogsRequest.
type ListWorkRequestLogsRequest struct {

	// The ID of the asynchronous work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListWorkRequestLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default sort order is descending and is based on the timeAccepted field.
	SortBy ListWorkRequestLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkRequestLogsResponse wrapper for the ListWorkRequestLogs operation
type ListWorkRequestLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestLogEntryCollection instances
	WorkRequestLogEntryCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListWorkRequestLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestLogsSortOrderEnum Enum with underlying type: string
type ListWorkRequestLogsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestLogsSortOrderEnum
const (
	ListWorkRequestLogsSortOrderAsc  ListWorkRequestLogsSortOrderEnum = "ASC"
	ListWorkRequestLogsSortOrderDesc ListWorkRequestLogsSortOrderEnum = "DESC"
)

var mappingListWorkRequestLogsSortOrder = map[string]ListWorkRequestLogsSortOrderEnum{
	"ASC":  ListWorkRequestLogsSortOrderAsc,
	"DESC": ListWorkRequestLogsSortOrderDesc,
}

// GetListWorkRequestLogsSortOrderEnumValues Enumerates the set of values for ListWorkRequestLogsSortOrderEnum
func GetListWorkRequestLogsSortOrderEnumValues() []ListWorkRequestLogsSortOrderEnum {
	values := make([]ListWorkRequestLogsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestLogsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestLogsSortByEnum Enum with underlying type: string
type ListWorkRequestLogsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestLogsSortByEnum
const (
	ListWorkRequestLogsSortByTimeaccepted ListWorkRequestLogsSortByEnum = "timeAccepted"
)

var mappingListWorkRequestLogsSortBy = map[string]ListWorkRequestLogsSortByEnum{
	"timeAccepted": ListWorkRequestLogsSortByTimeaccepted,
}

// GetListWorkRequestLogsSortByEnumValues Enumerates the set of values for ListWorkRequestLogsSortByEnum
func GetListWorkRequestLogsSortByEnumValues() []ListWorkRequestLogsSortByEnum {
	values := make([]ListWorkRequestLogsSortByEnum, 0)
	for _, v := range mappingListWorkRequestLogsSortBy {
		values = append(values, v)
	}
	return values
}
