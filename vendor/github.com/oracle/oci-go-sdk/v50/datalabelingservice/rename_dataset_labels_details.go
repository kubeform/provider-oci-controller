// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DataLabelingService API
//
// A description of the DataLabelingService API
//

package datalabelingservice

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// RenameDatasetLabelsDetails Renames a subset of Labels in the Dataset's LabelSet.  The Labels in the source LabelSet will be replaced with the Labels in the target LabelSet.
// Labels are correlated by index, i.e. the first Label in the source LabelSet will be replaced by the first Label in the target LabelSet.
// If the size of the source and target LabelSets are not equal, the request will be rejected.
type RenameDatasetLabelsDetails struct {
	SourceLabelSet *LabelSet `mandatory:"false" json:"sourceLabelSet"`

	TargetLabelSet *LabelSet `mandatory:"false" json:"targetLabelSet"`
}

func (m RenameDatasetLabelsDetails) String() string {
	return common.PointerString(m)
}
