// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v45/common"
)

// MigrationPhaseSummary Migration Phase Summary of details.
type MigrationPhaseSummary struct {

	// ODMS Job phase name
	Name OdmsJobPhasesEnum `mandatory:"true" json:"name"`

	// Array of actions for the corresponding phase. Empty array would indicate there is no supported action for the phase.
	SupportedActions []OdmsPhaseActionsEnum `mandatory:"true" json:"supportedActions"`

	// Action recommended for this phase. If not included in the response, there is no recommended action for the phase.
	RecommendedAction OdmsPhaseActionsEnum `mandatory:"false" json:"recommendedAction,omitempty"`
}

func (m MigrationPhaseSummary) String() string {
	return common.PointerString(m)
}
