// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// GetSecurityAssessmentComparisonRequest wrapper for the GetSecurityAssessmentComparison operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetSecurityAssessmentComparison.go.html to see an example of how to use GetSecurityAssessmentComparisonRequest.
type GetSecurityAssessmentComparisonRequest struct {

	// The OCID of the security assessment.
	SecurityAssessmentId *string `mandatory:"true" contributesTo:"path" name:"securityAssessmentId"`

	// The OCID of the baseline security assessment.
	ComparisonSecurityAssessmentId *string `mandatory:"true" contributesTo:"path" name:"comparisonSecurityAssessmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetSecurityAssessmentComparisonRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetSecurityAssessmentComparisonRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetSecurityAssessmentComparisonRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetSecurityAssessmentComparisonRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetSecurityAssessmentComparisonResponse wrapper for the GetSecurityAssessmentComparison operation
type GetSecurityAssessmentComparisonResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SecurityAssessmentComparison instance
	SecurityAssessmentComparison `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. For more information, see ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven)
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetSecurityAssessmentComparisonResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetSecurityAssessmentComparisonResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
