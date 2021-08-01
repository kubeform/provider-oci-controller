// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v45/common"
	"io"
	"net/http"
)

// UploadLogEventsFileRequest wrapper for the UploadLogEventsFile operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UploadLogEventsFile.go.html to see an example of how to use UploadLogEventsFileRequest.
type UploadLogEventsFileRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The log group OCID that gets mapped to the uploaded logs.
	LogGroupId *string `mandatory:"true" contributesTo:"query" name:"logGroupId"`

	// Log events data to be uploaded. The data could be uploaded with or without logSet information depending on whether logSet is enabled for the tenancy or not. Supported formats include
	// 1. json file : logSet (if needed) should be sent as "logSet" query parameter
	// 2. gzip file : logSet (if needed) should be sent as "logSet" query parameter
	// 3. zip file : containing multiple json files. LogSet information (if needed) should be appended to every filename in the zip.
	//    Supported filename format with logSet detail is &lt;filename&gt;_logSet=&lt;logset&gt;.json
	UploadLogEventsFileDetails io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The log set that gets associated with the uploaded logs.
	LogSet *string `mandatory:"false" contributesTo:"query" name:"logSet"`

	// Identifies the type of request payload.
	PayloadType UploadLogEventsFilePayloadTypeEnum `mandatory:"false" contributesTo:"query" name:"payloadType" omitEmpty:"true"`

	// The content type of the log data.
	ContentType *string `mandatory:"false" contributesTo:"header" name:"content-type"`

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

func (request UploadLogEventsFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UploadLogEventsFileRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request UploadLogEventsFileRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.UploadLogEventsFileDetails)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UploadLogEventsFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UploadLogEventsFileResponse wrapper for the UploadLogEventsFile operation
type UploadLogEventsFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for log data.
	OpcObjectId *string `presentIn:"header" name:"opc-object-id"`

	// The time the upload was created, in the format defined by RFC3339
	TimeCreated *common.SDKTime `presentIn:"header" name:"timecreated"`
}

func (response UploadLogEventsFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UploadLogEventsFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UploadLogEventsFilePayloadTypeEnum Enum with underlying type: string
type UploadLogEventsFilePayloadTypeEnum string

// Set of constants representing the allowable values for UploadLogEventsFilePayloadTypeEnum
const (
	UploadLogEventsFilePayloadTypeJson UploadLogEventsFilePayloadTypeEnum = "JSON"
	UploadLogEventsFilePayloadTypeGzip UploadLogEventsFilePayloadTypeEnum = "GZIP"
	UploadLogEventsFilePayloadTypeZip  UploadLogEventsFilePayloadTypeEnum = "ZIP"
)

var mappingUploadLogEventsFilePayloadType = map[string]UploadLogEventsFilePayloadTypeEnum{
	"JSON": UploadLogEventsFilePayloadTypeJson,
	"GZIP": UploadLogEventsFilePayloadTypeGzip,
	"ZIP":  UploadLogEventsFilePayloadTypeZip,
}

// GetUploadLogEventsFilePayloadTypeEnumValues Enumerates the set of values for UploadLogEventsFilePayloadTypeEnum
func GetUploadLogEventsFilePayloadTypeEnumValues() []UploadLogEventsFilePayloadTypeEnum {
	values := make([]UploadLogEventsFilePayloadTypeEnum, 0)
	for _, v := range mappingUploadLogEventsFilePayloadType {
		values = append(values, v)
	}
	return values
}
