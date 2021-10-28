// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// DeploymentBackupSummary The summary of the Backup.
type DeploymentBackupSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the backup being referenced.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the deployment being referenced.
	DeploymentId *string `mandatory:"true" json:"deploymentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Possible lifecycle states.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Version of OGG
	OggVersion *string `mandatory:"true" json:"oggVersion"`

	// An object's Display Name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// True if this object is automatically created
	IsAutomatic *bool `mandatory:"false" json:"isAutomatic"`

	// Describes the object's current state in detail. For example, it can be used to provide actionable information for a resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time of the resource backup. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeOfBackup *common.SDKTime `mandatory:"false" json:"timeOfBackup"`

	// Possible Deployment backup types.
	BackupType DeploymentBackupTypeEnum `mandatory:"false" json:"backupType,omitempty"`

	// Name of namespace that serves as a container for all of your buckets
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// Name of the bucket where the object is to be uploaded in the object storage
	BucketName *string `mandatory:"false" json:"bucketName"`

	// Name of the object to be uploaded to object storage
	ObjectName *string `mandatory:"false" json:"objectName"`

	// The time the resource was created. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the resource was last updated. The format is defined by RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The system tags associated with this resource, if any. The system tags are set by Oracle Cloud Infrastructure services. Each key is predefined and scoped to namespaces.  For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{orcl-cloud: {free-tier-retain: true}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DeploymentBackupSummary) String() string {
	return common.PointerString(m)
}
