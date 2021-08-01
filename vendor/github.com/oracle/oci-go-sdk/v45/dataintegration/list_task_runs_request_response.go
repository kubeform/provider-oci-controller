// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v45/common"
	"net/http"
)

// ListTaskRunsRequest wrapper for the ListTaskRuns operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTaskRuns.go.html to see an example of how to use ListTaskRunsRequest.
type ListTaskRunsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used to filter by the key of the object.
	Key []string `contributesTo:"query" name:"key" collectionFormat:"multi"`

	// Used to filter by the project or the folder object.
	AggregatorKey *string `mandatory:"false" contributesTo:"query" name:"aggregatorKey"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListTaskRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListTaskRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// This filter parameter can be used to filter by model specific queryable fields of the object <br><br><B>Examples:-</B><br> <ul> <li><B>?filter=status eq Failed</B> returns all objects that have a status field with value Failed</li> </ul>
	Filter []string `contributesTo:"query" name:"filter" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTaskRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTaskRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTaskRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTaskRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTaskRunsResponse wrapper for the ListTaskRuns operation
type ListTaskRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TaskRunSummaryCollection instances
	TaskRunSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListTaskRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTaskRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTaskRunsSortOrderEnum Enum with underlying type: string
type ListTaskRunsSortOrderEnum string

// Set of constants representing the allowable values for ListTaskRunsSortOrderEnum
const (
	ListTaskRunsSortOrderAsc  ListTaskRunsSortOrderEnum = "ASC"
	ListTaskRunsSortOrderDesc ListTaskRunsSortOrderEnum = "DESC"
)

var mappingListTaskRunsSortOrder = map[string]ListTaskRunsSortOrderEnum{
	"ASC":  ListTaskRunsSortOrderAsc,
	"DESC": ListTaskRunsSortOrderDesc,
}

// GetListTaskRunsSortOrderEnumValues Enumerates the set of values for ListTaskRunsSortOrderEnum
func GetListTaskRunsSortOrderEnumValues() []ListTaskRunsSortOrderEnum {
	values := make([]ListTaskRunsSortOrderEnum, 0)
	for _, v := range mappingListTaskRunsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListTaskRunsSortByEnum Enum with underlying type: string
type ListTaskRunsSortByEnum string

// Set of constants representing the allowable values for ListTaskRunsSortByEnum
const (
	ListTaskRunsSortByTimeCreated ListTaskRunsSortByEnum = "TIME_CREATED"
	ListTaskRunsSortByDisplayName ListTaskRunsSortByEnum = "DISPLAY_NAME"
)

var mappingListTaskRunsSortBy = map[string]ListTaskRunsSortByEnum{
	"TIME_CREATED": ListTaskRunsSortByTimeCreated,
	"DISPLAY_NAME": ListTaskRunsSortByDisplayName,
}

// GetListTaskRunsSortByEnumValues Enumerates the set of values for ListTaskRunsSortByEnum
func GetListTaskRunsSortByEnumValues() []ListTaskRunsSortByEnum {
	values := make([]ListTaskRunsSortByEnum, 0)
	for _, v := range mappingListTaskRunsSortBy {
		values = append(values, v)
	}
	return values
}
