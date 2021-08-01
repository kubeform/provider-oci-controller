// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

// RequestMethodsEnum Enum with underlying type: string
type RequestMethodsEnum string

// Set of constants representing the allowable values for RequestMethodsEnum
const (
	RequestMethodsGet  RequestMethodsEnum = "GET"
	RequestMethodsPost RequestMethodsEnum = "POST"
)

var mappingRequestMethods = map[string]RequestMethodsEnum{
	"GET":  RequestMethodsGet,
	"POST": RequestMethodsPost,
}

// GetRequestMethodsEnumValues Enumerates the set of values for RequestMethodsEnum
func GetRequestMethodsEnumValues() []RequestMethodsEnum {
	values := make([]RequestMethodsEnum, 0)
	for _, v := range mappingRequestMethods {
		values = append(values, v)
	}
	return values
}
