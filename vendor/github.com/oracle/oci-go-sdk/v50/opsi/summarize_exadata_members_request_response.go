// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// SummarizeExadataMembersRequest wrapper for the SummarizeExadataMembers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataMembers.go.html to see an example of how to use SummarizeExadataMembersRequest.
type SummarizeExadataMembersRequest struct {

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource.
	ExadataInsightId *string `mandatory:"true" contributesTo:"query" name:"exadataInsightId"`

	// Filter by one or more Exadata types.
	// Possible value are DBMACHINE, EXACS, and EXACC.
	ExadataType []string `contributesTo:"query" name:"exadataType" collectionFormat:"multi"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeExadataMembersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The order in which exadata member records are listed
	SortBy SummarizeExadataMembersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeExadataMembersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeExadataMembersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeExadataMembersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeExadataMembersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeExadataMembersResponse wrapper for the SummarizeExadataMembers operation
type SummarizeExadataMembersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExadataMemberCollection instances
	ExadataMemberCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeExadataMembersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeExadataMembersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeExadataMembersSortOrderEnum Enum with underlying type: string
type SummarizeExadataMembersSortOrderEnum string

// Set of constants representing the allowable values for SummarizeExadataMembersSortOrderEnum
const (
	SummarizeExadataMembersSortOrderAsc  SummarizeExadataMembersSortOrderEnum = "ASC"
	SummarizeExadataMembersSortOrderDesc SummarizeExadataMembersSortOrderEnum = "DESC"
)

var mappingSummarizeExadataMembersSortOrder = map[string]SummarizeExadataMembersSortOrderEnum{
	"ASC":  SummarizeExadataMembersSortOrderAsc,
	"DESC": SummarizeExadataMembersSortOrderDesc,
}

// GetSummarizeExadataMembersSortOrderEnumValues Enumerates the set of values for SummarizeExadataMembersSortOrderEnum
func GetSummarizeExadataMembersSortOrderEnumValues() []SummarizeExadataMembersSortOrderEnum {
	values := make([]SummarizeExadataMembersSortOrderEnum, 0)
	for _, v := range mappingSummarizeExadataMembersSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeExadataMembersSortByEnum Enum with underlying type: string
type SummarizeExadataMembersSortByEnum string

// Set of constants representing the allowable values for SummarizeExadataMembersSortByEnum
const (
	SummarizeExadataMembersSortByName        SummarizeExadataMembersSortByEnum = "name"
	SummarizeExadataMembersSortByDisplayname SummarizeExadataMembersSortByEnum = "displayName"
	SummarizeExadataMembersSortByEntitytype  SummarizeExadataMembersSortByEnum = "entityType"
)

var mappingSummarizeExadataMembersSortBy = map[string]SummarizeExadataMembersSortByEnum{
	"name":        SummarizeExadataMembersSortByName,
	"displayName": SummarizeExadataMembersSortByDisplayname,
	"entityType":  SummarizeExadataMembersSortByEntitytype,
}

// GetSummarizeExadataMembersSortByEnumValues Enumerates the set of values for SummarizeExadataMembersSortByEnum
func GetSummarizeExadataMembersSortByEnumValues() []SummarizeExadataMembersSortByEnum {
	values := make([]SummarizeExadataMembersSortByEnum, 0)
	for _, v := range mappingSummarizeExadataMembersSortBy {
		values = append(values, v)
	}
	return values
}
