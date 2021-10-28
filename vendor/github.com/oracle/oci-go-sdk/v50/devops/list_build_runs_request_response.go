// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListBuildRunsRequest wrapper for the ListBuildRuns operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListBuildRuns.go.html to see an example of how to use ListBuildRunsRequest.
type ListBuildRunsRequest struct {

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Unique build pipeline identifier.
	BuildPipelineId *string `mandatory:"false" contributesTo:"query" name:"buildPipelineId"`

	// unique project identifier
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only Build Runs that matches the given lifecycleState.
	LifecycleState BuildRunLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListBuildRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListBuildRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBuildRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBuildRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBuildRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBuildRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBuildRunsResponse wrapper for the ListBuildRuns operation
type ListBuildRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BuildRunSummaryCollection instances
	BuildRunSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBuildRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBuildRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBuildRunsSortOrderEnum Enum with underlying type: string
type ListBuildRunsSortOrderEnum string

// Set of constants representing the allowable values for ListBuildRunsSortOrderEnum
const (
	ListBuildRunsSortOrderAsc  ListBuildRunsSortOrderEnum = "ASC"
	ListBuildRunsSortOrderDesc ListBuildRunsSortOrderEnum = "DESC"
)

var mappingListBuildRunsSortOrder = map[string]ListBuildRunsSortOrderEnum{
	"ASC":  ListBuildRunsSortOrderAsc,
	"DESC": ListBuildRunsSortOrderDesc,
}

// GetListBuildRunsSortOrderEnumValues Enumerates the set of values for ListBuildRunsSortOrderEnum
func GetListBuildRunsSortOrderEnumValues() []ListBuildRunsSortOrderEnum {
	values := make([]ListBuildRunsSortOrderEnum, 0)
	for _, v := range mappingListBuildRunsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListBuildRunsSortByEnum Enum with underlying type: string
type ListBuildRunsSortByEnum string

// Set of constants representing the allowable values for ListBuildRunsSortByEnum
const (
	ListBuildRunsSortByTimecreated ListBuildRunsSortByEnum = "timeCreated"
	ListBuildRunsSortByDisplayname ListBuildRunsSortByEnum = "displayName"
)

var mappingListBuildRunsSortBy = map[string]ListBuildRunsSortByEnum{
	"timeCreated": ListBuildRunsSortByTimecreated,
	"displayName": ListBuildRunsSortByDisplayname,
}

// GetListBuildRunsSortByEnumValues Enumerates the set of values for ListBuildRunsSortByEnum
func GetListBuildRunsSortByEnumValues() []ListBuildRunsSortByEnum {
	values := make([]ListBuildRunsSortByEnum, 0)
	for _, v := range mappingListBuildRunsSortBy {
		values = append(values, v)
	}
	return values
}
