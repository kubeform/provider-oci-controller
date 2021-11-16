// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package certificatesmanagement

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// ListAssociationsRequest wrapper for the ListAssociations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/certificatesmanagement/ListAssociations.go.html to see an example of how to use ListAssociationsRequest.
type ListAssociationsRequest struct {

	// Unique Oracle-assigned identifier for the request. If provided, the returned request ID
	// will include this value. Otherwise, a random request ID will be
	// generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter that returns only resources that match the given compartment OCID.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter that returns only resources that match the given OCID of a certificate-related resource.
	CertificatesResourceId *string `mandatory:"false" contributesTo:"query" name:"certificatesResourceId"`

	// A filter that returns only resources that match the given OCID of an associated Oracle Cloud Infrastructure resource.
	AssociatedResourceId *string `mandatory:"false" contributesTo:"query" name:"associatedResourceId"`

	// The OCID of the association. If the parameter is set to null, the service lists all associations.
	AssociationId *string `mandatory:"false" contributesTo:"query" name:"associationId"`

	// A filter that returns only resources that match the specified name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can specify only one sort order. The default order for `TIMECREATED` is descending.
	// The default order for `NAME` is ascending.
	SortBy ListAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Type of associations to list. If the parameter is set to null, the service lists all types of associations.
	AssociationType ListAssociationsAssociationTypeEnum `mandatory:"false" contributesTo:"query" name:"associationType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAssociationsResponse wrapper for the ListAssociations operation
type ListAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssociationCollection instances
	AssociationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssociationsSortByEnum Enum with underlying type: string
type ListAssociationsSortByEnum string

// Set of constants representing the allowable values for ListAssociationsSortByEnum
const (
	ListAssociationsSortByName        ListAssociationsSortByEnum = "NAME"
	ListAssociationsSortByTimecreated ListAssociationsSortByEnum = "TIMECREATED"
)

var mappingListAssociationsSortBy = map[string]ListAssociationsSortByEnum{
	"NAME":        ListAssociationsSortByName,
	"TIMECREATED": ListAssociationsSortByTimecreated,
}

// GetListAssociationsSortByEnumValues Enumerates the set of values for ListAssociationsSortByEnum
func GetListAssociationsSortByEnumValues() []ListAssociationsSortByEnum {
	values := make([]ListAssociationsSortByEnum, 0)
	for _, v := range mappingListAssociationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListAssociationsSortOrderEnum Enum with underlying type: string
type ListAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListAssociationsSortOrderEnum
const (
	ListAssociationsSortOrderAsc  ListAssociationsSortOrderEnum = "ASC"
	ListAssociationsSortOrderDesc ListAssociationsSortOrderEnum = "DESC"
)

var mappingListAssociationsSortOrder = map[string]ListAssociationsSortOrderEnum{
	"ASC":  ListAssociationsSortOrderAsc,
	"DESC": ListAssociationsSortOrderDesc,
}

// GetListAssociationsSortOrderEnumValues Enumerates the set of values for ListAssociationsSortOrderEnum
func GetListAssociationsSortOrderEnumValues() []ListAssociationsSortOrderEnum {
	values := make([]ListAssociationsSortOrderEnum, 0)
	for _, v := range mappingListAssociationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAssociationsAssociationTypeEnum Enum with underlying type: string
type ListAssociationsAssociationTypeEnum string

// Set of constants representing the allowable values for ListAssociationsAssociationTypeEnum
const (
	ListAssociationsAssociationTypeCertificate          ListAssociationsAssociationTypeEnum = "CERTIFICATE"
	ListAssociationsAssociationTypeCertificateAuthority ListAssociationsAssociationTypeEnum = "CERTIFICATE_AUTHORITY"
	ListAssociationsAssociationTypeCaBundle             ListAssociationsAssociationTypeEnum = "CA_BUNDLE"
)

var mappingListAssociationsAssociationType = map[string]ListAssociationsAssociationTypeEnum{
	"CERTIFICATE":           ListAssociationsAssociationTypeCertificate,
	"CERTIFICATE_AUTHORITY": ListAssociationsAssociationTypeCertificateAuthority,
	"CA_BUNDLE":             ListAssociationsAssociationTypeCaBundle,
}

// GetListAssociationsAssociationTypeEnumValues Enumerates the set of values for ListAssociationsAssociationTypeEnum
func GetListAssociationsAssociationTypeEnumValues() []ListAssociationsAssociationTypeEnum {
	values := make([]ListAssociationsAssociationTypeEnum, 0)
	for _, v := range mappingListAssociationsAssociationType {
		values = append(values, v)
	}
	return values
}