// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// ApplyJobOperationDetails Job details that are specific to apply operations.
type ApplyJobOperationDetails struct {
	TerraformAdvancedOptions *TerraformAdvancedOptions `mandatory:"false" json:"terraformAdvancedOptions"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the plan job that contains the execution plan used for this job,
	// or `null` if no execution plan was used.
	ExecutionPlanJobId *string `mandatory:"false" json:"executionPlanJobId"`

	// Specifies the source of the execution plan to apply.
	// Use `AUTO_APPROVED` to run the job without an execution plan.
	ExecutionPlanStrategy ApplyJobOperationDetailsExecutionPlanStrategyEnum `mandatory:"true" json:"executionPlanStrategy"`
}

func (m ApplyJobOperationDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ApplyJobOperationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeApplyJobOperationDetails ApplyJobOperationDetails
	s := struct {
		DiscriminatorParam string `json:"operation"`
		MarshalTypeApplyJobOperationDetails
	}{
		"APPLY",
		(MarshalTypeApplyJobOperationDetails)(m),
	}

	return json.Marshal(&s)
}

// ApplyJobOperationDetailsExecutionPlanStrategyEnum Enum with underlying type: string
type ApplyJobOperationDetailsExecutionPlanStrategyEnum string

// Set of constants representing the allowable values for ApplyJobOperationDetailsExecutionPlanStrategyEnum
const (
	ApplyJobOperationDetailsExecutionPlanStrategyFromPlanJobId     ApplyJobOperationDetailsExecutionPlanStrategyEnum = "FROM_PLAN_JOB_ID"
	ApplyJobOperationDetailsExecutionPlanStrategyFromLatestPlanJob ApplyJobOperationDetailsExecutionPlanStrategyEnum = "FROM_LATEST_PLAN_JOB"
	ApplyJobOperationDetailsExecutionPlanStrategyAutoApproved      ApplyJobOperationDetailsExecutionPlanStrategyEnum = "AUTO_APPROVED"
)

var mappingApplyJobOperationDetailsExecutionPlanStrategy = map[string]ApplyJobOperationDetailsExecutionPlanStrategyEnum{
	"FROM_PLAN_JOB_ID":     ApplyJobOperationDetailsExecutionPlanStrategyFromPlanJobId,
	"FROM_LATEST_PLAN_JOB": ApplyJobOperationDetailsExecutionPlanStrategyFromLatestPlanJob,
	"AUTO_APPROVED":        ApplyJobOperationDetailsExecutionPlanStrategyAutoApproved,
}

// GetApplyJobOperationDetailsExecutionPlanStrategyEnumValues Enumerates the set of values for ApplyJobOperationDetailsExecutionPlanStrategyEnum
func GetApplyJobOperationDetailsExecutionPlanStrategyEnumValues() []ApplyJobOperationDetailsExecutionPlanStrategyEnum {
	values := make([]ApplyJobOperationDetailsExecutionPlanStrategyEnum, 0)
	for _, v := range mappingApplyJobOperationDetailsExecutionPlanStrategy {
		values = append(values, v)
	}
	return values
}
