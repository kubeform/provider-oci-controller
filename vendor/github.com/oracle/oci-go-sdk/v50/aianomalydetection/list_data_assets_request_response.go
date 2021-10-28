// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package aianomalydetection

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListDataAssetsRequest wrapper for the ListDataAssets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/aianomalydetection/ListDataAssets.go.html to see an example of how to use ListDataAssetsRequest.
type ListDataAssetsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The ID of the project for which to list the objects.
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState DataAssetLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataAssetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDataAssetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataAssetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataAssetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataAssetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataAssetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDataAssetsResponse wrapper for the ListDataAssets operation
type ListDataAssetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataAssetCollection instances
	DataAssetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataAssetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataAssetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataAssetsSortOrderEnum Enum with underlying type: string
type ListDataAssetsSortOrderEnum string

// Set of constants representing the allowable values for ListDataAssetsSortOrderEnum
const (
	ListDataAssetsSortOrderAsc  ListDataAssetsSortOrderEnum = "ASC"
	ListDataAssetsSortOrderDesc ListDataAssetsSortOrderEnum = "DESC"
)

var mappingListDataAssetsSortOrder = map[string]ListDataAssetsSortOrderEnum{
	"ASC":  ListDataAssetsSortOrderAsc,
	"DESC": ListDataAssetsSortOrderDesc,
}

// GetListDataAssetsSortOrderEnumValues Enumerates the set of values for ListDataAssetsSortOrderEnum
func GetListDataAssetsSortOrderEnumValues() []ListDataAssetsSortOrderEnum {
	values := make([]ListDataAssetsSortOrderEnum, 0)
	for _, v := range mappingListDataAssetsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDataAssetsSortByEnum Enum with underlying type: string
type ListDataAssetsSortByEnum string

// Set of constants representing the allowable values for ListDataAssetsSortByEnum
const (
	ListDataAssetsSortByTimecreated ListDataAssetsSortByEnum = "timeCreated"
	ListDataAssetsSortByDisplayname ListDataAssetsSortByEnum = "displayName"
)

var mappingListDataAssetsSortBy = map[string]ListDataAssetsSortByEnum{
	"timeCreated": ListDataAssetsSortByTimecreated,
	"displayName": ListDataAssetsSortByDisplayname,
}

// GetListDataAssetsSortByEnumValues Enumerates the set of values for ListDataAssetsSortByEnum
func GetListDataAssetsSortByEnumValues() []ListDataAssetsSortByEnum {
	values := make([]ListDataAssetsSortByEnum, 0)
	for _, v := range mappingListDataAssetsSortBy {
		values = append(values, v)
	}
	return values
}
