// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListUsersRequest wrapper for the ListUsers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUsers.go.html to see an example of how to use ListUsersRequest.
type ListUsersRequest struct {

	// The OCID of the user assessment.
	UserAssessmentId *string `mandatory:"true" contributesTo:"path" name:"userAssessmentId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListUsersAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only items that match the specified user category.
	UserCategory *string `mandatory:"false" contributesTo:"query" name:"userCategory"`

	// A filter to return only items that match the specified user key.
	UserKey *string `mandatory:"false" contributesTo:"query" name:"userKey"`

	// A filter to return only items that match the specified account status.
	AccountStatus *string `mandatory:"false" contributesTo:"query" name:"accountStatus"`

	// A filter to return only items that match the specified authentication type.
	AuthenticationType *string `mandatory:"false" contributesTo:"query" name:"authenticationType"`

	// A filter to return only items that match the specified user name.
	UserName *string `mandatory:"false" contributesTo:"query" name:"userName"`

	// A filter to return only items that match the specified target.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return users whose last login time in the database is greater than or equal to the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeLastLoginGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastLoginGreaterThanOrEqualTo"`

	// A filter to return users whose last login time in the database is less than the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeLastLoginLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLastLoginLessThan"`

	// A filter to return users whose creation time in the database is greater than or equal to the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeUserCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUserCreatedGreaterThanOrEqualTo"`

	// A filter to return users whose creation time in the database is less than the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeUserCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUserCreatedLessThan"`

	// A filter to return users whose last password change in the database is greater than or equal to the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimePasswordLastChangedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timePasswordLastChangedGreaterThanOrEqualTo"`

	// A filter to return users whose last password change in the database is less than the date and time specified, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// **Example:** 2016-12-19T16:39:57.600Z
	TimePasswordLastChangedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timePasswordLastChangedLessThan"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order for userName is ascending.
	SortBy ListUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListUsersResponse wrapper for the ListUsers operation
type ListUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []UserSummary instances
	Items []UserSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUsersAccessLevelEnum Enum with underlying type: string
type ListUsersAccessLevelEnum string

// Set of constants representing the allowable values for ListUsersAccessLevelEnum
const (
	ListUsersAccessLevelRestricted ListUsersAccessLevelEnum = "RESTRICTED"
	ListUsersAccessLevelAccessible ListUsersAccessLevelEnum = "ACCESSIBLE"
)

var mappingListUsersAccessLevel = map[string]ListUsersAccessLevelEnum{
	"RESTRICTED": ListUsersAccessLevelRestricted,
	"ACCESSIBLE": ListUsersAccessLevelAccessible,
}

// GetListUsersAccessLevelEnumValues Enumerates the set of values for ListUsersAccessLevelEnum
func GetListUsersAccessLevelEnumValues() []ListUsersAccessLevelEnum {
	values := make([]ListUsersAccessLevelEnum, 0)
	for _, v := range mappingListUsersAccessLevel {
		values = append(values, v)
	}
	return values
}

// ListUsersSortOrderEnum Enum with underlying type: string
type ListUsersSortOrderEnum string

// Set of constants representing the allowable values for ListUsersSortOrderEnum
const (
	ListUsersSortOrderAsc  ListUsersSortOrderEnum = "ASC"
	ListUsersSortOrderDesc ListUsersSortOrderEnum = "DESC"
)

var mappingListUsersSortOrder = map[string]ListUsersSortOrderEnum{
	"ASC":  ListUsersSortOrderAsc,
	"DESC": ListUsersSortOrderDesc,
}

// GetListUsersSortOrderEnumValues Enumerates the set of values for ListUsersSortOrderEnum
func GetListUsersSortOrderEnumValues() []ListUsersSortOrderEnum {
	values := make([]ListUsersSortOrderEnum, 0)
	for _, v := range mappingListUsersSortOrder {
		values = append(values, v)
	}
	return values
}

// ListUsersSortByEnum Enum with underlying type: string
type ListUsersSortByEnum string

// Set of constants representing the allowable values for ListUsersSortByEnum
const (
	ListUsersSortByUsername            ListUsersSortByEnum = "userName"
	ListUsersSortByUsercategory        ListUsersSortByEnum = "userCategory"
	ListUsersSortByAccountstatus       ListUsersSortByEnum = "accountStatus"
	ListUsersSortByTimelastlogin       ListUsersSortByEnum = "timeLastLogin"
	ListUsersSortByTargetid            ListUsersSortByEnum = "targetId"
	ListUsersSortByTimeusercreated     ListUsersSortByEnum = "timeUserCreated"
	ListUsersSortByAuthenticationtype  ListUsersSortByEnum = "authenticationType"
	ListUsersSortByTimepasswordchanged ListUsersSortByEnum = "timePasswordChanged"
)

var mappingListUsersSortBy = map[string]ListUsersSortByEnum{
	"userName":            ListUsersSortByUsername,
	"userCategory":        ListUsersSortByUsercategory,
	"accountStatus":       ListUsersSortByAccountstatus,
	"timeLastLogin":       ListUsersSortByTimelastlogin,
	"targetId":            ListUsersSortByTargetid,
	"timeUserCreated":     ListUsersSortByTimeusercreated,
	"authenticationType":  ListUsersSortByAuthenticationtype,
	"timePasswordChanged": ListUsersSortByTimepasswordchanged,
}

// GetListUsersSortByEnumValues Enumerates the set of values for ListUsersSortByEnum
func GetListUsersSortByEnumValues() []ListUsersSortByEnum {
	values := make([]ListUsersSortByEnum, 0)
	for _, v := range mappingListUsersSortBy {
		values = append(values, v)
	}
	return values
}
