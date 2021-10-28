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

// FindingSummary The particular finding reported by the security assessment.
type FindingSummary struct {

	// The severity of the finding.
	Severity FindingSummarySeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// The OCID of the assessment that generated this finding.
	AssessmentId *string `mandatory:"false" json:"assessmentId"`

	// The OCID of the target database.
	TargetId *string `mandatory:"false" json:"targetId"`

	// The unique finding key. This is a system-generated identifier. To get the finding key for a finding, use ListFindings.
	Key *string `mandatory:"false" json:"key"`

	// The short title for the finding.
	Title *string `mandatory:"false" json:"title"`

	// The explanation of the issue in this finding. It explains the reason for the rule and, if a risk is reported, it may also explain the recommended actions for remediation.
	Remarks *string `mandatory:"false" json:"remarks"`

	// The details of the finding. Provides detailed information to explain the finding summary, typically results from the assessed database, followed by any recommendations for changes.
	Details *interface{} `mandatory:"false" json:"details"`

	// The brief summary of the finding. When the finding is informational, the summary typically reports only the number of data elements that were examined.
	Summary *string `mandatory:"false" json:"summary"`

	// Provides information on whether the finding is related to a CIS Oracle Database Benchmark recommendation, a STIG rule, or a GDPR Article/Recital.
	References *References `mandatory:"false" json:"references"`
}

func (m FindingSummary) String() string {
	return common.PointerString(m)
}

// FindingSummarySeverityEnum Enum with underlying type: string
type FindingSummarySeverityEnum string

// Set of constants representing the allowable values for FindingSummarySeverityEnum
const (
	FindingSummarySeverityHigh     FindingSummarySeverityEnum = "HIGH"
	FindingSummarySeverityMedium   FindingSummarySeverityEnum = "MEDIUM"
	FindingSummarySeverityLow      FindingSummarySeverityEnum = "LOW"
	FindingSummarySeverityEvaluate FindingSummarySeverityEnum = "EVALUATE"
	FindingSummarySeverityAdvisory FindingSummarySeverityEnum = "ADVISORY"
	FindingSummarySeverityPass     FindingSummarySeverityEnum = "PASS"
)

var mappingFindingSummarySeverity = map[string]FindingSummarySeverityEnum{
	"HIGH":     FindingSummarySeverityHigh,
	"MEDIUM":   FindingSummarySeverityMedium,
	"LOW":      FindingSummarySeverityLow,
	"EVALUATE": FindingSummarySeverityEvaluate,
	"ADVISORY": FindingSummarySeverityAdvisory,
	"PASS":     FindingSummarySeverityPass,
}

// GetFindingSummarySeverityEnumValues Enumerates the set of values for FindingSummarySeverityEnum
func GetFindingSummarySeverityEnumValues() []FindingSummarySeverityEnum {
	values := make([]FindingSummarySeverityEnum, 0)
	for _, v := range mappingFindingSummarySeverity {
		values = append(values, v)
	}
	return values
}
