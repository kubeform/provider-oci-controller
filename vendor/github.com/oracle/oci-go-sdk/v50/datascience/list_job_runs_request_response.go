// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListJobRunsRequest wrapper for the ListJobRuns operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListJobRuns.go.html to see an example of how to use ListJobRunsRequest.
type ListJobRunsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	JobId *string `mandatory:"false" contributesTo:"query" name:"jobId"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListJobRunsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. When you sort by `displayName`, the results are
	// shown in ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListJobRunsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListJobRunsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJobRunsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobRunsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListJobRunsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobRunsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListJobRunsResponse wrapper for the ListJobRuns operation
type ListJobRunsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []JobRunSummary instances
	Items []JobRunSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListJobRunsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobRunsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobRunsSortOrderEnum Enum with underlying type: string
type ListJobRunsSortOrderEnum string

// Set of constants representing the allowable values for ListJobRunsSortOrderEnum
const (
	ListJobRunsSortOrderAsc  ListJobRunsSortOrderEnum = "ASC"
	ListJobRunsSortOrderDesc ListJobRunsSortOrderEnum = "DESC"
)

var mappingListJobRunsSortOrder = map[string]ListJobRunsSortOrderEnum{
	"ASC":  ListJobRunsSortOrderAsc,
	"DESC": ListJobRunsSortOrderDesc,
}

// GetListJobRunsSortOrderEnumValues Enumerates the set of values for ListJobRunsSortOrderEnum
func GetListJobRunsSortOrderEnumValues() []ListJobRunsSortOrderEnum {
	values := make([]ListJobRunsSortOrderEnum, 0)
	for _, v := range mappingListJobRunsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListJobRunsSortByEnum Enum with underlying type: string
type ListJobRunsSortByEnum string

// Set of constants representing the allowable values for ListJobRunsSortByEnum
const (
	ListJobRunsSortByTimecreated ListJobRunsSortByEnum = "timeCreated"
	ListJobRunsSortByDisplayname ListJobRunsSortByEnum = "displayName"
)

var mappingListJobRunsSortBy = map[string]ListJobRunsSortByEnum{
	"timeCreated": ListJobRunsSortByTimecreated,
	"displayName": ListJobRunsSortByDisplayname,
}

// GetListJobRunsSortByEnumValues Enumerates the set of values for ListJobRunsSortByEnum
func GetListJobRunsSortByEnumValues() []ListJobRunsSortByEnum {
	values := make([]ListJobRunsSortByEnum, 0)
	for _, v := range mappingListJobRunsSortBy {
		values = append(values, v)
	}
	return values
}

// ListJobRunsLifecycleStateEnum Enum with underlying type: string
type ListJobRunsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJobRunsLifecycleStateEnum
const (
	ListJobRunsLifecycleStateAccepted       ListJobRunsLifecycleStateEnum = "ACCEPTED"
	ListJobRunsLifecycleStateInProgress     ListJobRunsLifecycleStateEnum = "IN_PROGRESS"
	ListJobRunsLifecycleStateFailed         ListJobRunsLifecycleStateEnum = "FAILED"
	ListJobRunsLifecycleStateSucceeded      ListJobRunsLifecycleStateEnum = "SUCCEEDED"
	ListJobRunsLifecycleStateCanceling      ListJobRunsLifecycleStateEnum = "CANCELING"
	ListJobRunsLifecycleStateCanceled       ListJobRunsLifecycleStateEnum = "CANCELED"
	ListJobRunsLifecycleStateDeleted        ListJobRunsLifecycleStateEnum = "DELETED"
	ListJobRunsLifecycleStateNeedsAttention ListJobRunsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListJobRunsLifecycleState = map[string]ListJobRunsLifecycleStateEnum{
	"ACCEPTED":        ListJobRunsLifecycleStateAccepted,
	"IN_PROGRESS":     ListJobRunsLifecycleStateInProgress,
	"FAILED":          ListJobRunsLifecycleStateFailed,
	"SUCCEEDED":       ListJobRunsLifecycleStateSucceeded,
	"CANCELING":       ListJobRunsLifecycleStateCanceling,
	"CANCELED":        ListJobRunsLifecycleStateCanceled,
	"DELETED":         ListJobRunsLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListJobRunsLifecycleStateNeedsAttention,
}

// GetListJobRunsLifecycleStateEnumValues Enumerates the set of values for ListJobRunsLifecycleStateEnum
func GetListJobRunsLifecycleStateEnumValues() []ListJobRunsLifecycleStateEnum {
	values := make([]ListJobRunsLifecycleStateEnum, 0)
	for _, v := range mappingListJobRunsLifecycleState {
		values = append(values, v)
	}
	return values
}
