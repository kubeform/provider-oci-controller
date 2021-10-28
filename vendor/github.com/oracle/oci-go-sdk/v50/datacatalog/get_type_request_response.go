// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// GetTypeRequest wrapper for the GetType operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetType.go.html to see an example of how to use GetTypeRequest.
type GetTypeRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique type key.
	TypeKey *string `mandatory:"true" contributesTo:"path" name:"typeKey"`

	// Specifies the fields to return in a type response.
	Fields []GetTypeFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetTypeRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetTypeRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetTypeRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetTypeRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetTypeResponse wrapper for the GetType operation
type GetTypeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ModelType instance
	ModelType `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTypeResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetTypeResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetTypeFieldsEnum Enum with underlying type: string
type GetTypeFieldsEnum string

// Set of constants representing the allowable values for GetTypeFieldsEnum
const (
	GetTypeFieldsKey              GetTypeFieldsEnum = "key"
	GetTypeFieldsDescription      GetTypeFieldsEnum = "description"
	GetTypeFieldsName             GetTypeFieldsEnum = "name"
	GetTypeFieldsCatalogid        GetTypeFieldsEnum = "catalogId"
	GetTypeFieldsProperties       GetTypeFieldsEnum = "properties"
	GetTypeFieldsIsinternal       GetTypeFieldsEnum = "isInternal"
	GetTypeFieldsIstag            GetTypeFieldsEnum = "isTag"
	GetTypeFieldsIsapproved       GetTypeFieldsEnum = "isApproved"
	GetTypeFieldsTypecategory     GetTypeFieldsEnum = "typeCategory"
	GetTypeFieldsExternaltypename GetTypeFieldsEnum = "externalTypeName"
	GetTypeFieldsLifecyclestate   GetTypeFieldsEnum = "lifecycleState"
	GetTypeFieldsUri              GetTypeFieldsEnum = "uri"
)

var mappingGetTypeFields = map[string]GetTypeFieldsEnum{
	"key":              GetTypeFieldsKey,
	"description":      GetTypeFieldsDescription,
	"name":             GetTypeFieldsName,
	"catalogId":        GetTypeFieldsCatalogid,
	"properties":       GetTypeFieldsProperties,
	"isInternal":       GetTypeFieldsIsinternal,
	"isTag":            GetTypeFieldsIstag,
	"isApproved":       GetTypeFieldsIsapproved,
	"typeCategory":     GetTypeFieldsTypecategory,
	"externalTypeName": GetTypeFieldsExternaltypename,
	"lifecycleState":   GetTypeFieldsLifecyclestate,
	"uri":              GetTypeFieldsUri,
}

// GetGetTypeFieldsEnumValues Enumerates the set of values for GetTypeFieldsEnum
func GetGetTypeFieldsEnumValues() []GetTypeFieldsEnum {
	values := make([]GetTypeFieldsEnum, 0)
	for _, v := range mappingGetTypeFields {
		values = append(values, v)
	}
	return values
}
