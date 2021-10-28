// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Manage solutions in Oracle Cloud Infrastructure Service Catalog.
//

package servicecatalog

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// PrivateApplicationStackPackage A stack package for private applications.
type PrivateApplicationStackPackage struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private application package.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the private application where the package is hosted.
	PrivateApplicationId *string `mandatory:"true" json:"privateApplicationId"`

	// The package version.
	Version *string `mandatory:"true" json:"version"`

	// The date and time the private application package was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339)
	// timestamp format.
	// Example: `2021-05-27T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The display name of the package.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The content URL of the terraform configuration.
	ContentUrl *string `mandatory:"false" json:"contentUrl"`

	// The MIME type of the terraform configuration.
	MimeType *string `mandatory:"false" json:"mimeType"`
}

//GetId returns Id
func (m PrivateApplicationStackPackage) GetId() *string {
	return m.Id
}

//GetPrivateApplicationId returns PrivateApplicationId
func (m PrivateApplicationStackPackage) GetPrivateApplicationId() *string {
	return m.PrivateApplicationId
}

//GetDisplayName returns DisplayName
func (m PrivateApplicationStackPackage) GetDisplayName() *string {
	return m.DisplayName
}

//GetVersion returns Version
func (m PrivateApplicationStackPackage) GetVersion() *string {
	return m.Version
}

//GetTimeCreated returns TimeCreated
func (m PrivateApplicationStackPackage) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

func (m PrivateApplicationStackPackage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m PrivateApplicationStackPackage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePrivateApplicationStackPackage PrivateApplicationStackPackage
	s := struct {
		DiscriminatorParam string `json:"packageType"`
		MarshalTypePrivateApplicationStackPackage
	}{
		"STACK",
		(MarshalTypePrivateApplicationStackPackage)(m),
	}

	return json.Marshal(&s)
}
