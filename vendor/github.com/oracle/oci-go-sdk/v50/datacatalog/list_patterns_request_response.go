// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListPatternsRequest wrapper for the ListPatterns operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListPatterns.go.html to see an example of how to use ListPatternsRequest.
type ListPatternsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListPatternsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Specifies the fields to return in a pattern summary response.
	Fields []ListPatternsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListPatternsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListPatternsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPatternsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPatternsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPatternsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPatternsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPatternsResponse wrapper for the ListPatterns operation
type ListPatternsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PatternCollection instances
	PatternCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPatternsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPatternsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPatternsLifecycleStateEnum Enum with underlying type: string
type ListPatternsLifecycleStateEnum string

// Set of constants representing the allowable values for ListPatternsLifecycleStateEnum
const (
	ListPatternsLifecycleStateCreating ListPatternsLifecycleStateEnum = "CREATING"
	ListPatternsLifecycleStateActive   ListPatternsLifecycleStateEnum = "ACTIVE"
	ListPatternsLifecycleStateInactive ListPatternsLifecycleStateEnum = "INACTIVE"
	ListPatternsLifecycleStateUpdating ListPatternsLifecycleStateEnum = "UPDATING"
	ListPatternsLifecycleStateDeleting ListPatternsLifecycleStateEnum = "DELETING"
	ListPatternsLifecycleStateDeleted  ListPatternsLifecycleStateEnum = "DELETED"
	ListPatternsLifecycleStateFailed   ListPatternsLifecycleStateEnum = "FAILED"
	ListPatternsLifecycleStateMoving   ListPatternsLifecycleStateEnum = "MOVING"
)

var mappingListPatternsLifecycleState = map[string]ListPatternsLifecycleStateEnum{
	"CREATING": ListPatternsLifecycleStateCreating,
	"ACTIVE":   ListPatternsLifecycleStateActive,
	"INACTIVE": ListPatternsLifecycleStateInactive,
	"UPDATING": ListPatternsLifecycleStateUpdating,
	"DELETING": ListPatternsLifecycleStateDeleting,
	"DELETED":  ListPatternsLifecycleStateDeleted,
	"FAILED":   ListPatternsLifecycleStateFailed,
	"MOVING":   ListPatternsLifecycleStateMoving,
}

// GetListPatternsLifecycleStateEnumValues Enumerates the set of values for ListPatternsLifecycleStateEnum
func GetListPatternsLifecycleStateEnumValues() []ListPatternsLifecycleStateEnum {
	values := make([]ListPatternsLifecycleStateEnum, 0)
	for _, v := range mappingListPatternsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListPatternsFieldsEnum Enum with underlying type: string
type ListPatternsFieldsEnum string

// Set of constants representing the allowable values for ListPatternsFieldsEnum
const (
	ListPatternsFieldsKey            ListPatternsFieldsEnum = "key"
	ListPatternsFieldsDisplayname    ListPatternsFieldsEnum = "displayName"
	ListPatternsFieldsDescription    ListPatternsFieldsEnum = "description"
	ListPatternsFieldsCatalogid      ListPatternsFieldsEnum = "catalogId"
	ListPatternsFieldsExpression     ListPatternsFieldsEnum = "expression"
	ListPatternsFieldsLifecyclestate ListPatternsFieldsEnum = "lifecycleState"
	ListPatternsFieldsTimecreated    ListPatternsFieldsEnum = "timeCreated"
)

var mappingListPatternsFields = map[string]ListPatternsFieldsEnum{
	"key":            ListPatternsFieldsKey,
	"displayName":    ListPatternsFieldsDisplayname,
	"description":    ListPatternsFieldsDescription,
	"catalogId":      ListPatternsFieldsCatalogid,
	"expression":     ListPatternsFieldsExpression,
	"lifecycleState": ListPatternsFieldsLifecyclestate,
	"timeCreated":    ListPatternsFieldsTimecreated,
}

// GetListPatternsFieldsEnumValues Enumerates the set of values for ListPatternsFieldsEnum
func GetListPatternsFieldsEnumValues() []ListPatternsFieldsEnum {
	values := make([]ListPatternsFieldsEnum, 0)
	for _, v := range mappingListPatternsFields {
		values = append(values, v)
	}
	return values
}

// ListPatternsSortByEnum Enum with underlying type: string
type ListPatternsSortByEnum string

// Set of constants representing the allowable values for ListPatternsSortByEnum
const (
	ListPatternsSortByTimecreated ListPatternsSortByEnum = "TIMECREATED"
	ListPatternsSortByDisplayname ListPatternsSortByEnum = "DISPLAYNAME"
)

var mappingListPatternsSortBy = map[string]ListPatternsSortByEnum{
	"TIMECREATED": ListPatternsSortByTimecreated,
	"DISPLAYNAME": ListPatternsSortByDisplayname,
}

// GetListPatternsSortByEnumValues Enumerates the set of values for ListPatternsSortByEnum
func GetListPatternsSortByEnumValues() []ListPatternsSortByEnum {
	values := make([]ListPatternsSortByEnum, 0)
	for _, v := range mappingListPatternsSortBy {
		values = append(values, v)
	}
	return values
}

// ListPatternsSortOrderEnum Enum with underlying type: string
type ListPatternsSortOrderEnum string

// Set of constants representing the allowable values for ListPatternsSortOrderEnum
const (
	ListPatternsSortOrderAsc  ListPatternsSortOrderEnum = "ASC"
	ListPatternsSortOrderDesc ListPatternsSortOrderEnum = "DESC"
)

var mappingListPatternsSortOrder = map[string]ListPatternsSortOrderEnum{
	"ASC":  ListPatternsSortOrderAsc,
	"DESC": ListPatternsSortOrderDesc,
}

// GetListPatternsSortOrderEnumValues Enumerates the set of values for ListPatternsSortOrderEnum
func GetListPatternsSortOrderEnumValues() []ListPatternsSortOrderEnum {
	values := make([]ListPatternsSortOrderEnum, 0)
	for _, v := range mappingListPatternsSortOrder {
		values = append(values, v)
	}
	return values
}
