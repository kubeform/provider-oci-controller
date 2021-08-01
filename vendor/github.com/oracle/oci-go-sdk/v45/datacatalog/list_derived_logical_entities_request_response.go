// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v45/common"
	"net/http"
)

// ListDerivedLogicalEntitiesRequest wrapper for the ListDerivedLogicalEntities operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListDerivedLogicalEntities.go.html to see an example of how to use ListDerivedLogicalEntitiesRequest.
type ListDerivedLogicalEntitiesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique pattern key.
	PatternKey *string `mandatory:"true" contributesTo:"path" name:"patternKey"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu".
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListDerivedLogicalEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDerivedLogicalEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

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

func (request ListDerivedLogicalEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDerivedLogicalEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDerivedLogicalEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDerivedLogicalEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDerivedLogicalEntitiesResponse wrapper for the ListDerivedLogicalEntities operation
type ListDerivedLogicalEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EntityCollection instances
	EntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListDerivedLogicalEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDerivedLogicalEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDerivedLogicalEntitiesSortByEnum Enum with underlying type: string
type ListDerivedLogicalEntitiesSortByEnum string

// Set of constants representing the allowable values for ListDerivedLogicalEntitiesSortByEnum
const (
	ListDerivedLogicalEntitiesSortByTimecreated ListDerivedLogicalEntitiesSortByEnum = "TIMECREATED"
	ListDerivedLogicalEntitiesSortByDisplayname ListDerivedLogicalEntitiesSortByEnum = "DISPLAYNAME"
)

var mappingListDerivedLogicalEntitiesSortBy = map[string]ListDerivedLogicalEntitiesSortByEnum{
	"TIMECREATED": ListDerivedLogicalEntitiesSortByTimecreated,
	"DISPLAYNAME": ListDerivedLogicalEntitiesSortByDisplayname,
}

// GetListDerivedLogicalEntitiesSortByEnumValues Enumerates the set of values for ListDerivedLogicalEntitiesSortByEnum
func GetListDerivedLogicalEntitiesSortByEnumValues() []ListDerivedLogicalEntitiesSortByEnum {
	values := make([]ListDerivedLogicalEntitiesSortByEnum, 0)
	for _, v := range mappingListDerivedLogicalEntitiesSortBy {
		values = append(values, v)
	}
	return values
}

// ListDerivedLogicalEntitiesSortOrderEnum Enum with underlying type: string
type ListDerivedLogicalEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListDerivedLogicalEntitiesSortOrderEnum
const (
	ListDerivedLogicalEntitiesSortOrderAsc  ListDerivedLogicalEntitiesSortOrderEnum = "ASC"
	ListDerivedLogicalEntitiesSortOrderDesc ListDerivedLogicalEntitiesSortOrderEnum = "DESC"
)

var mappingListDerivedLogicalEntitiesSortOrder = map[string]ListDerivedLogicalEntitiesSortOrderEnum{
	"ASC":  ListDerivedLogicalEntitiesSortOrderAsc,
	"DESC": ListDerivedLogicalEntitiesSortOrderDesc,
}

// GetListDerivedLogicalEntitiesSortOrderEnumValues Enumerates the set of values for ListDerivedLogicalEntitiesSortOrderEnum
func GetListDerivedLogicalEntitiesSortOrderEnumValues() []ListDerivedLogicalEntitiesSortOrderEnum {
	values := make([]ListDerivedLogicalEntitiesSortOrderEnum, 0)
	for _, v := range mappingListDerivedLogicalEntitiesSortOrder {
		values = append(values, v)
	}
	return values
}
