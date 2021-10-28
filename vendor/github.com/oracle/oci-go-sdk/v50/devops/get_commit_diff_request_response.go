// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// GetCommitDiffRequest wrapper for the GetCommitDiff operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetCommitDiff.go.html to see an example of how to use GetCommitDiffRequest.
type GetCommitDiffRequest struct {

	// unique Repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// The commit or ref name where changes are coming from
	TargetVersion *string `mandatory:"true" contributesTo:"query" name:"targetVersion"`

	// The commit or ref name to compare changes against. If baseVersion is not provided, the diff will be gone against an empty tree.
	BaseVersion *string `mandatory:"false" contributesTo:"query" name:"baseVersion"`

	// boolean for whether to use merge base or most recent revision
	IsComparisonFromMergeBase *bool `mandatory:"false" contributesTo:"query" name:"isComparisonFromMergeBase"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCommitDiffRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCommitDiffRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetCommitDiffRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCommitDiffRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetCommitDiffResponse wrapper for the GetCommitDiff operation
type GetCommitDiffResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DiffResponse instance
	DiffResponse `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCommitDiffResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCommitDiffResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
