// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListExternalPublicationValidationsRequest wrapper for the ListExternalPublicationValidations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListExternalPublicationValidations.go.html to see an example of how to use ListExternalPublicationValidationsRequest.
type ListExternalPublicationValidationsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The task key.
	TaskKey *string `mandatory:"true" contributesTo:"path" name:"taskKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

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
	SortOrder ListExternalPublicationValidationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListExternalPublicationValidationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalPublicationValidationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalPublicationValidationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalPublicationValidationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalPublicationValidationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListExternalPublicationValidationsResponse wrapper for the ListExternalPublicationValidations operation
type ListExternalPublicationValidationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalPublicationValidationSummaryCollection instances
	ExternalPublicationValidationSummaryCollection `presentIn:"body"`

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

func (response ListExternalPublicationValidationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalPublicationValidationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalPublicationValidationsSortOrderEnum Enum with underlying type: string
type ListExternalPublicationValidationsSortOrderEnum string

// Set of constants representing the allowable values for ListExternalPublicationValidationsSortOrderEnum
const (
	ListExternalPublicationValidationsSortOrderAsc  ListExternalPublicationValidationsSortOrderEnum = "ASC"
	ListExternalPublicationValidationsSortOrderDesc ListExternalPublicationValidationsSortOrderEnum = "DESC"
)

var mappingListExternalPublicationValidationsSortOrder = map[string]ListExternalPublicationValidationsSortOrderEnum{
	"ASC":  ListExternalPublicationValidationsSortOrderAsc,
	"DESC": ListExternalPublicationValidationsSortOrderDesc,
}

// GetListExternalPublicationValidationsSortOrderEnumValues Enumerates the set of values for ListExternalPublicationValidationsSortOrderEnum
func GetListExternalPublicationValidationsSortOrderEnumValues() []ListExternalPublicationValidationsSortOrderEnum {
	values := make([]ListExternalPublicationValidationsSortOrderEnum, 0)
	for _, v := range mappingListExternalPublicationValidationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListExternalPublicationValidationsSortByEnum Enum with underlying type: string
type ListExternalPublicationValidationsSortByEnum string

// Set of constants representing the allowable values for ListExternalPublicationValidationsSortByEnum
const (
	ListExternalPublicationValidationsSortByTimeCreated ListExternalPublicationValidationsSortByEnum = "TIME_CREATED"
	ListExternalPublicationValidationsSortByDisplayName ListExternalPublicationValidationsSortByEnum = "DISPLAY_NAME"
)

var mappingListExternalPublicationValidationsSortBy = map[string]ListExternalPublicationValidationsSortByEnum{
	"TIME_CREATED": ListExternalPublicationValidationsSortByTimeCreated,
	"DISPLAY_NAME": ListExternalPublicationValidationsSortByDisplayName,
}

// GetListExternalPublicationValidationsSortByEnumValues Enumerates the set of values for ListExternalPublicationValidationsSortByEnum
func GetListExternalPublicationValidationsSortByEnumValues() []ListExternalPublicationValidationsSortByEnum {
	values := make([]ListExternalPublicationValidationsSortByEnum, 0)
	for _, v := range mappingListExternalPublicationValidationsSortBy {
		values = append(values, v)
	}
	return values
}
