// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// GetPatternRequest wrapper for the GetPattern operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetPattern.go.html to see an example of how to use GetPatternRequest.
type GetPatternRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique pattern key.
	PatternKey *string `mandatory:"true" contributesTo:"path" name:"patternKey"`

	// Specifies the fields to return in a pattern response.
	Fields []GetPatternFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPatternRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPatternRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetPatternRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPatternRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetPatternResponse wrapper for the GetPattern operation
type GetPatternResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Pattern instance
	Pattern `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPatternResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPatternResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetPatternFieldsEnum Enum with underlying type: string
type GetPatternFieldsEnum string

// Set of constants representing the allowable values for GetPatternFieldsEnum
const (
	GetPatternFieldsKey            GetPatternFieldsEnum = "key"
	GetPatternFieldsDisplayname    GetPatternFieldsEnum = "displayName"
	GetPatternFieldsDescription    GetPatternFieldsEnum = "description"
	GetPatternFieldsCatalogid      GetPatternFieldsEnum = "catalogId"
	GetPatternFieldsExpression     GetPatternFieldsEnum = "expression"
	GetPatternFieldsLifecyclestate GetPatternFieldsEnum = "lifecycleState"
	GetPatternFieldsTimecreated    GetPatternFieldsEnum = "timeCreated"
	GetPatternFieldsTimeupdated    GetPatternFieldsEnum = "timeUpdated"
	GetPatternFieldsCreatedbyid    GetPatternFieldsEnum = "createdById"
	GetPatternFieldsUpdatedbyid    GetPatternFieldsEnum = "updatedById"
	GetPatternFieldsProperties     GetPatternFieldsEnum = "properties"
)

var mappingGetPatternFields = map[string]GetPatternFieldsEnum{
	"key":            GetPatternFieldsKey,
	"displayName":    GetPatternFieldsDisplayname,
	"description":    GetPatternFieldsDescription,
	"catalogId":      GetPatternFieldsCatalogid,
	"expression":     GetPatternFieldsExpression,
	"lifecycleState": GetPatternFieldsLifecyclestate,
	"timeCreated":    GetPatternFieldsTimecreated,
	"timeUpdated":    GetPatternFieldsTimeupdated,
	"createdById":    GetPatternFieldsCreatedbyid,
	"updatedById":    GetPatternFieldsUpdatedbyid,
	"properties":     GetPatternFieldsProperties,
}

// GetGetPatternFieldsEnumValues Enumerates the set of values for GetPatternFieldsEnum
func GetGetPatternFieldsEnumValues() []GetPatternFieldsEnum {
	values := make([]GetPatternFieldsEnum, 0)
	for _, v := range mappingGetPatternFields {
		values = append(values, v)
	}
	return values
}
