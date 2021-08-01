// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v45/common"
	"net/http"
)

// ListAssociatedEntitiesRequest wrapper for the ListAssociatedEntities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListAssociatedEntities.go.html to see an example of how to use ListAssociatedEntitiesRequest.
type ListAssociatedEntitiesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The entity OCID.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// The entity type used for filtering.  Only associations on an entity with the
	// specified type will be returned.
	EntityType *string `mandatory:"false" contributesTo:"query" name:"entityType"`

	// The entity type display name used for filtering.  Only items associated with the entity
	// with the specified type display name will be returned.
	EntityTypeDisplayName *string `mandatory:"false" contributesTo:"query" name:"entityTypeDisplayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAssociatedEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned entities
	SortBy ListAssociatedEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssociatedEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssociatedEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssociatedEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssociatedEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAssociatedEntitiesResponse wrapper for the ListAssociatedEntities operation
type ListAssociatedEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsAssociatedEntityCollection instances
	LogAnalyticsAssociatedEntityCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAssociatedEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssociatedEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssociatedEntitiesSortOrderEnum Enum with underlying type: string
type ListAssociatedEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListAssociatedEntitiesSortOrderEnum
const (
	ListAssociatedEntitiesSortOrderAsc  ListAssociatedEntitiesSortOrderEnum = "ASC"
	ListAssociatedEntitiesSortOrderDesc ListAssociatedEntitiesSortOrderEnum = "DESC"
)

var mappingListAssociatedEntitiesSortOrder = map[string]ListAssociatedEntitiesSortOrderEnum{
	"ASC":  ListAssociatedEntitiesSortOrderAsc,
	"DESC": ListAssociatedEntitiesSortOrderDesc,
}

// GetListAssociatedEntitiesSortOrderEnumValues Enumerates the set of values for ListAssociatedEntitiesSortOrderEnum
func GetListAssociatedEntitiesSortOrderEnumValues() []ListAssociatedEntitiesSortOrderEnum {
	values := make([]ListAssociatedEntitiesSortOrderEnum, 0)
	for _, v := range mappingListAssociatedEntitiesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAssociatedEntitiesSortByEnum Enum with underlying type: string
type ListAssociatedEntitiesSortByEnum string

// Set of constants representing the allowable values for ListAssociatedEntitiesSortByEnum
const (
	ListAssociatedEntitiesSortByEntityname            ListAssociatedEntitiesSortByEnum = "entityName"
	ListAssociatedEntitiesSortByEntitytypedisplayname ListAssociatedEntitiesSortByEnum = "entityTypeDisplayName"
	ListAssociatedEntitiesSortByAssociationcount      ListAssociatedEntitiesSortByEnum = "associationCount"
)

var mappingListAssociatedEntitiesSortBy = map[string]ListAssociatedEntitiesSortByEnum{
	"entityName":            ListAssociatedEntitiesSortByEntityname,
	"entityTypeDisplayName": ListAssociatedEntitiesSortByEntitytypedisplayname,
	"associationCount":      ListAssociatedEntitiesSortByAssociationcount,
}

// GetListAssociatedEntitiesSortByEnumValues Enumerates the set of values for ListAssociatedEntitiesSortByEnum
func GetListAssociatedEntitiesSortByEnumValues() []ListAssociatedEntitiesSortByEnum {
	values := make([]ListAssociatedEntitiesSortByEnum, 0)
	for _, v := range mappingListAssociatedEntitiesSortBy {
		values = append(values, v)
	}
	return values
}
