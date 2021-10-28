// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// SecurityAssessmentStatistics Statistics showing the number of findings for each category grouped by risk levels for all
// the targets in the specified security assessment.
// The categories include Auditing, Authorization Control, Data Encryption, Database Configuration,
// Fine-Grained Access Control, Privileges and Roles, and User Accounts.
// The risk levels include High Risk, Medium Risk, Low Risk, Advisory, Evaluate, and Pass.
type SecurityAssessmentStatistics struct {

	// The total number of targets in this security assessment.
	TargetsCount *int `mandatory:"false" json:"targetsCount"`

	HighRisk *SectionStatistics `mandatory:"false" json:"highRisk"`

	MediumRisk *SectionStatistics `mandatory:"false" json:"mediumRisk"`

	LowRisk *SectionStatistics `mandatory:"false" json:"lowRisk"`

	Advisory *SectionStatistics `mandatory:"false" json:"advisory"`

	Evaluate *SectionStatistics `mandatory:"false" json:"evaluate"`

	Pass *SectionStatistics `mandatory:"false" json:"pass"`
}

func (m SecurityAssessmentStatistics) String() string {
	return common.PointerString(m)
}
