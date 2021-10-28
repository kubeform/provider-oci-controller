// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v50/common"
)

// LogAnalyticsSource LogAnalyticsSource
type LogAnalyticsSource struct {

	// The label alert conditions.
	LabelConditions []LogAnalyticsSourceLabelCondition `mandatory:"false" json:"labelConditions"`

	// The association count.
	AssociationCount *int `mandatory:"false" json:"associationCount"`

	// The association entities.
	AssociationEntity []LogAnalyticsAssociation `mandatory:"false" json:"associationEntity"`

	// The data filter definitions.
	DataFilterDefinitions []LogAnalyticsSourceDataFilter `mandatory:"false" json:"dataFilterDefinitions"`

	// The database credential.
	DatabaseCredential *string `mandatory:"false" json:"databaseCredential"`

	// The extended field definitions.
	ExtendedFieldDefinitions []LogAnalyticsSourceExtendedFieldDefinition `mandatory:"false" json:"extendedFieldDefinitions"`

	// A flag indicating whether or not this is a cloud source.
	IsForCloud *bool `mandatory:"false" json:"isForCloud"`

	// The labels associated with the source.
	Labels []LogAnalyticsLabelView `mandatory:"false" json:"labels"`

	// The metric definitions.
	MetricDefinitions []LogAnalyticsMetric `mandatory:"false" json:"metricDefinitions"`

	// The metric source map.
	Metrics []LogAnalyticsSourceMetric `mandatory:"false" json:"metrics"`

	// The built in parsers associated with source.
	OobParsers []LogAnalyticsParser `mandatory:"false" json:"oobParsers"`

	// The source parameters.
	Parameters []LogAnalyticsParameter `mandatory:"false" json:"parameters"`

	// The pattern count.
	PatternCount *int `mandatory:"false" json:"patternCount"`

	// The source patterns.
	Patterns []LogAnalyticsSourcePattern `mandatory:"false" json:"patterns"`

	// The source description.
	Description *string `mandatory:"false" json:"description"`

	// The source display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The source edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The source functions.
	Functions []LogAnalyticsSourceFunction `mandatory:"false" json:"functions"`

	// The source unique identifier.
	SourceId *int64 `mandatory:"false" json:"sourceId"`

	// The source internal name.
	Name *string `mandatory:"false" json:"name"`

	// A flag indicating whether or not the source content is secure.
	IsSecureContent *bool `mandatory:"false" json:"isSecureContent"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The list of parsers used by the source.
	Parsers []LogAnalyticsParser `mandatory:"false" json:"parsers"`

	// A flag indicating whether or not the source is marked for auto-association.
	IsAutoAssociationEnabled *bool `mandatory:"false" json:"isAutoAssociationEnabled"`

	// A flag indicating whether or not the auto-association state should be overriden.
	IsAutoAssociationOverride *bool `mandatory:"false" json:"isAutoAssociationOverride"`

	// The rule unique identifier.
	RuleId *int64 `mandatory:"false" json:"ruleId"`

	// The source type internal name.
	TypeName *string `mandatory:"false" json:"typeName"`

	// The source type name.
	TypeDisplayName *string `mandatory:"false" json:"typeDisplayName"`

	// The source warning configuration.
	WarningConfig *int64 `mandatory:"false" json:"warningConfig"`

	// The source metadata fields.
	MetadataFields []LogAnalyticsSourceMetadataField `mandatory:"false" json:"metadataFields"`

	// The labls used by the source.
	LabelDefinitions []LogAnalyticsLabelDefinition `mandatory:"false" json:"labelDefinitions"`

	// The entity types.
	EntityTypes []LogAnalyticsSourceEntityType `mandatory:"false" json:"entityTypes"`

	// A flag indicating whether or not the source has a time zone override.
	IsTimezoneOverride *bool `mandatory:"false" json:"isTimezoneOverride"`

	// An array of custom parsers.
	UserParsers []LogAnalyticsParser `mandatory:"false" json:"userParsers"`

	// The last updated date.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An array of event types.
	EventTypes []EventType `mandatory:"false" json:"eventTypes"`
}

func (m LogAnalyticsSource) String() string {
	return common.PointerString(m)
}
