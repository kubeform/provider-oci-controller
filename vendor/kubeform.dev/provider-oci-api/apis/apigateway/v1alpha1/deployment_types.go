/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Deployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DeploymentSpec   `json:"spec,omitempty"`
	Status            DeploymentStatus `json:"status,omitempty"`
}

type DeploymentSpecSpecificationLoggingPoliciesAccessLog struct {
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
}

type DeploymentSpecSpecificationLoggingPoliciesExecutionLog struct {
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
	// +optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
}

type DeploymentSpecSpecificationLoggingPolicies struct {
	// +optional
	AccessLog *DeploymentSpecSpecificationLoggingPoliciesAccessLog `json:"accessLog,omitempty" tf:"access_log"`
	// +optional
	ExecutionLog *DeploymentSpecSpecificationLoggingPoliciesExecutionLog `json:"executionLog,omitempty" tf:"execution_log"`
}

type DeploymentSpecSpecificationRequestPoliciesAuthenticationPublicKeysKeys struct {
	// +optional
	Alg *string `json:"alg,omitempty" tf:"alg"`
	// +optional
	E      *string `json:"e,omitempty" tf:"e"`
	Format *string `json:"format" tf:"format"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	KeyOps []string `json:"keyOps,omitempty" tf:"key_ops"`
	// +optional
	Kid *string `json:"kid,omitempty" tf:"kid"`
	// +optional
	Kty *string `json:"kty,omitempty" tf:"kty"`
	// +optional
	N *string `json:"n,omitempty" tf:"n"`
	// +optional
	Use *string `json:"use,omitempty" tf:"use"`
}

type DeploymentSpecSpecificationRequestPoliciesAuthenticationPublicKeys struct {
	// +optional
	IsSslVerifyDisabled *bool `json:"isSslVerifyDisabled,omitempty" tf:"is_ssl_verify_disabled"`
	// +optional
	Keys []DeploymentSpecSpecificationRequestPoliciesAuthenticationPublicKeysKeys `json:"keys,omitempty" tf:"keys"`
	// +optional
	MaxCacheDurationInHours *int64  `json:"maxCacheDurationInHours,omitempty" tf:"max_cache_duration_in_hours"`
	Type                    *string `json:"type" tf:"type"`
	// +optional
	Uri *string `json:"uri,omitempty" tf:"uri"`
}

type DeploymentSpecSpecificationRequestPoliciesAuthenticationVerifyClaims struct {
	// +optional
	IsRequired *bool `json:"isRequired,omitempty" tf:"is_required"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type DeploymentSpecSpecificationRequestPoliciesAuthentication struct {
	// +optional
	Audiences []string `json:"audiences,omitempty" tf:"audiences"`
	// +optional
	FunctionID *string `json:"functionID,omitempty" tf:"function_id"`
	// +optional
	IsAnonymousAccessAllowed *bool `json:"isAnonymousAccessAllowed,omitempty" tf:"is_anonymous_access_allowed"`
	// +optional
	Issuers []string `json:"issuers,omitempty" tf:"issuers"`
	// +optional
	MaxClockSkewInSeconds *float64 `json:"maxClockSkewInSeconds,omitempty" tf:"max_clock_skew_in_seconds"`
	// +optional
	PublicKeys *DeploymentSpecSpecificationRequestPoliciesAuthenticationPublicKeys `json:"publicKeys,omitempty" tf:"public_keys"`
	// +optional
	TokenAuthScheme *string `json:"tokenAuthScheme,omitempty" tf:"token_auth_scheme"`
	// +optional
	TokenHeader *string `json:"tokenHeader,omitempty" tf:"token_header"`
	// +optional
	TokenQueryParam *string `json:"tokenQueryParam,omitempty" tf:"token_query_param"`
	Type            *string `json:"type" tf:"type"`
	// +optional
	VerifyClaims []DeploymentSpecSpecificationRequestPoliciesAuthenticationVerifyClaims `json:"verifyClaims,omitempty" tf:"verify_claims"`
}

type DeploymentSpecSpecificationRequestPoliciesCors struct {
	// +optional
	AllowedHeaders []string `json:"allowedHeaders,omitempty" tf:"allowed_headers"`
	// +optional
	AllowedMethods []string `json:"allowedMethods,omitempty" tf:"allowed_methods"`
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	ExposedHeaders []string `json:"exposedHeaders,omitempty" tf:"exposed_headers"`
	// +optional
	IsAllowCredentialsEnabled *bool `json:"isAllowCredentialsEnabled,omitempty" tf:"is_allow_credentials_enabled"`
	// +optional
	MaxAgeInSeconds *int64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds"`
}

type DeploymentSpecSpecificationRequestPoliciesRateLimiting struct {
	RateInRequestsPerSecond *int64  `json:"rateInRequestsPerSecond" tf:"rate_in_requests_per_second"`
	RateKey                 *string `json:"rateKey" tf:"rate_key"`
}

type DeploymentSpecSpecificationRequestPolicies struct {
	// +optional
	Authentication *DeploymentSpecSpecificationRequestPoliciesAuthentication `json:"authentication,omitempty" tf:"authentication"`
	// +optional
	Cors *DeploymentSpecSpecificationRequestPoliciesCors `json:"cors,omitempty" tf:"cors"`
	// +optional
	RateLimiting *DeploymentSpecSpecificationRequestPoliciesRateLimiting `json:"rateLimiting,omitempty" tf:"rate_limiting"`
}

type DeploymentSpecSpecificationRoutesBackendHeaders struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type DeploymentSpecSpecificationRoutesBackend struct {
	// +optional
	Body *string `json:"body,omitempty" tf:"body"`
	// +optional
	ConnectTimeoutInSeconds *float64 `json:"connectTimeoutInSeconds,omitempty" tf:"connect_timeout_in_seconds"`
	// +optional
	FunctionID *string `json:"functionID,omitempty" tf:"function_id"`
	// +optional
	Headers []DeploymentSpecSpecificationRoutesBackendHeaders `json:"headers,omitempty" tf:"headers"`
	// +optional
	IsSslVerifyDisabled *bool `json:"isSslVerifyDisabled,omitempty" tf:"is_ssl_verify_disabled"`
	// +optional
	ReadTimeoutInSeconds *float64 `json:"readTimeoutInSeconds,omitempty" tf:"read_timeout_in_seconds"`
	// +optional
	SendTimeoutInSeconds *float64 `json:"sendTimeoutInSeconds,omitempty" tf:"send_timeout_in_seconds"`
	// +optional
	Status *int64  `json:"status,omitempty" tf:"status"`
	Type   *string `json:"type" tf:"type"`
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type DeploymentSpecSpecificationRoutesLoggingPoliciesAccessLog struct {
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
}

type DeploymentSpecSpecificationRoutesLoggingPoliciesExecutionLog struct {
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
	// +optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
}

type DeploymentSpecSpecificationRoutesLoggingPolicies struct {
	// +optional
	AccessLog *DeploymentSpecSpecificationRoutesLoggingPoliciesAccessLog `json:"accessLog,omitempty" tf:"access_log"`
	// +optional
	ExecutionLog *DeploymentSpecSpecificationRoutesLoggingPoliciesExecutionLog `json:"executionLog,omitempty" tf:"execution_log"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesAuthorization struct {
	// +optional
	AllowedScope []string `json:"allowedScope,omitempty" tf:"allowed_scope"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesBodyValidationContent struct {
	MediaType      *string `json:"mediaType" tf:"media_type"`
	ValidationType *string `json:"validationType" tf:"validation_type"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesBodyValidation struct {
	// +optional
	Content []DeploymentSpecSpecificationRoutesRequestPoliciesBodyValidationContent `json:"content,omitempty" tf:"content"`
	// +optional
	Required *bool `json:"required,omitempty" tf:"required"`
	// +optional
	ValidationMode *string `json:"validationMode,omitempty" tf:"validation_mode"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesCors struct {
	// +optional
	AllowedHeaders []string `json:"allowedHeaders,omitempty" tf:"allowed_headers"`
	// +optional
	AllowedMethods []string `json:"allowedMethods,omitempty" tf:"allowed_methods"`
	AllowedOrigins []string `json:"allowedOrigins" tf:"allowed_origins"`
	// +optional
	ExposedHeaders []string `json:"exposedHeaders,omitempty" tf:"exposed_headers"`
	// +optional
	IsAllowCredentialsEnabled *bool `json:"isAllowCredentialsEnabled,omitempty" tf:"is_allow_credentials_enabled"`
	// +optional
	MaxAgeInSeconds *int64 `json:"maxAgeInSeconds,omitempty" tf:"max_age_in_seconds"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersItems struct {
	Name *string `json:"name" tf:"name"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeaders struct {
	Items []DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeadersItems `json:"items" tf:"items"`
	Type  *string                                                                                   `json:"type" tf:"type"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsRenameHeadersItems struct {
	From *string `json:"from" tf:"from"`
	To   *string `json:"to" tf:"to"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsRenameHeaders struct {
	Items []DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsRenameHeadersItems `json:"items" tf:"items"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersItems struct {
	// +optional
	IfExists *string  `json:"ifExists,omitempty" tf:"if_exists"`
	Name     *string  `json:"name" tf:"name"`
	Values   []string `json:"values" tf:"values"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeaders struct {
	Items []DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeadersItems `json:"items" tf:"items"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformations struct {
	// +optional
	FilterHeaders *DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsFilterHeaders `json:"filterHeaders,omitempty" tf:"filter_headers"`
	// +optional
	RenameHeaders *DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsRenameHeaders `json:"renameHeaders,omitempty" tf:"rename_headers"`
	// +optional
	SetHeaders *DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformationsSetHeaders `json:"setHeaders,omitempty" tf:"set_headers"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesHeaderValidationsHeaders struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	Required *bool `json:"required,omitempty" tf:"required"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesHeaderValidations struct {
	// +optional
	Headers []DeploymentSpecSpecificationRoutesRequestPoliciesHeaderValidationsHeaders `json:"headers,omitempty" tf:"headers"`
	// +optional
	ValidationMode *string `json:"validationMode,omitempty" tf:"validation_mode"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersItems struct {
	Name *string `json:"name" tf:"name"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParameters struct {
	Items []DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParametersItems `json:"items" tf:"items"`
	Type  *string                                                                                                   `json:"type" tf:"type"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsRenameQueryParametersItems struct {
	From *string `json:"from" tf:"from"`
	To   *string `json:"to" tf:"to"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsRenameQueryParameters struct {
	Items []DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsRenameQueryParametersItems `json:"items" tf:"items"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersItems struct {
	// +optional
	IfExists *string  `json:"ifExists,omitempty" tf:"if_exists"`
	Name     *string  `json:"name" tf:"name"`
	Values   []string `json:"values" tf:"values"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParameters struct {
	Items []DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParametersItems `json:"items" tf:"items"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformations struct {
	// +optional
	FilterQueryParameters *DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsFilterQueryParameters `json:"filterQueryParameters,omitempty" tf:"filter_query_parameters"`
	// +optional
	RenameQueryParameters *DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsRenameQueryParameters `json:"renameQueryParameters,omitempty" tf:"rename_query_parameters"`
	// +optional
	SetQueryParameters *DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformationsSetQueryParameters `json:"setQueryParameters,omitempty" tf:"set_query_parameters"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterValidationsParameters struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	Required *bool `json:"required,omitempty" tf:"required"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterValidations struct {
	// +optional
	Parameters []DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterValidationsParameters `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	ValidationMode *string `json:"validationMode,omitempty" tf:"validation_mode"`
}

type DeploymentSpecSpecificationRoutesRequestPoliciesResponseCacheLookup struct {
	// +optional
	CacheKeyAdditions []string `json:"cacheKeyAdditions,omitempty" tf:"cache_key_additions"`
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
	// +optional
	IsPrivateCachingEnabled *bool   `json:"isPrivateCachingEnabled,omitempty" tf:"is_private_caching_enabled"`
	Type                    *string `json:"type" tf:"type"`
}

type DeploymentSpecSpecificationRoutesRequestPolicies struct {
	// +optional
	Authorization *DeploymentSpecSpecificationRoutesRequestPoliciesAuthorization `json:"authorization,omitempty" tf:"authorization"`
	// +optional
	BodyValidation *DeploymentSpecSpecificationRoutesRequestPoliciesBodyValidation `json:"bodyValidation,omitempty" tf:"body_validation"`
	// +optional
	Cors *DeploymentSpecSpecificationRoutesRequestPoliciesCors `json:"cors,omitempty" tf:"cors"`
	// +optional
	HeaderTransformations *DeploymentSpecSpecificationRoutesRequestPoliciesHeaderTransformations `json:"headerTransformations,omitempty" tf:"header_transformations"`
	// +optional
	HeaderValidations *DeploymentSpecSpecificationRoutesRequestPoliciesHeaderValidations `json:"headerValidations,omitempty" tf:"header_validations"`
	// +optional
	QueryParameterTransformations *DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterTransformations `json:"queryParameterTransformations,omitempty" tf:"query_parameter_transformations"`
	// +optional
	QueryParameterValidations *DeploymentSpecSpecificationRoutesRequestPoliciesQueryParameterValidations `json:"queryParameterValidations,omitempty" tf:"query_parameter_validations"`
	// +optional
	ResponseCacheLookup *DeploymentSpecSpecificationRoutesRequestPoliciesResponseCacheLookup `json:"responseCacheLookup,omitempty" tf:"response_cache_lookup"`
}

type DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersItems struct {
	Name *string `json:"name" tf:"name"`
}

type DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeaders struct {
	Items []DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeadersItems `json:"items" tf:"items"`
	Type  *string                                                                                    `json:"type" tf:"type"`
}

type DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsRenameHeadersItems struct {
	From *string `json:"from" tf:"from"`
	To   *string `json:"to" tf:"to"`
}

type DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsRenameHeaders struct {
	Items []DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsRenameHeadersItems `json:"items" tf:"items"`
}

type DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersItems struct {
	// +optional
	IfExists *string  `json:"ifExists,omitempty" tf:"if_exists"`
	Name     *string  `json:"name" tf:"name"`
	Values   []string `json:"values" tf:"values"`
}

type DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeaders struct {
	Items []DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeadersItems `json:"items" tf:"items"`
}

type DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformations struct {
	// +optional
	FilterHeaders *DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsFilterHeaders `json:"filterHeaders,omitempty" tf:"filter_headers"`
	// +optional
	RenameHeaders *DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsRenameHeaders `json:"renameHeaders,omitempty" tf:"rename_headers"`
	// +optional
	SetHeaders *DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformationsSetHeaders `json:"setHeaders,omitempty" tf:"set_headers"`
}

type DeploymentSpecSpecificationRoutesResponsePoliciesResponseCacheStore struct {
	TimeToLiveInSeconds *int64  `json:"timeToLiveInSeconds" tf:"time_to_live_in_seconds"`
	Type                *string `json:"type" tf:"type"`
}

type DeploymentSpecSpecificationRoutesResponsePolicies struct {
	// +optional
	HeaderTransformations *DeploymentSpecSpecificationRoutesResponsePoliciesHeaderTransformations `json:"headerTransformations,omitempty" tf:"header_transformations"`
	// +optional
	ResponseCacheStore *DeploymentSpecSpecificationRoutesResponsePoliciesResponseCacheStore `json:"responseCacheStore,omitempty" tf:"response_cache_store"`
}

type DeploymentSpecSpecificationRoutes struct {
	Backend *DeploymentSpecSpecificationRoutesBackend `json:"backend" tf:"backend"`
	// +optional
	LoggingPolicies *DeploymentSpecSpecificationRoutesLoggingPolicies `json:"loggingPolicies,omitempty" tf:"logging_policies"`
	// +optional
	Methods []string `json:"methods,omitempty" tf:"methods"`
	Path    *string  `json:"path" tf:"path"`
	// +optional
	RequestPolicies *DeploymentSpecSpecificationRoutesRequestPolicies `json:"requestPolicies,omitempty" tf:"request_policies"`
	// +optional
	ResponsePolicies *DeploymentSpecSpecificationRoutesResponsePolicies `json:"responsePolicies,omitempty" tf:"response_policies"`
}

type DeploymentSpecSpecification struct {
	// +optional
	LoggingPolicies *DeploymentSpecSpecificationLoggingPolicies `json:"loggingPolicies,omitempty" tf:"logging_policies"`
	// +optional
	RequestPolicies *DeploymentSpecSpecificationRequestPolicies `json:"requestPolicies,omitempty" tf:"request_policies"`
	Routes          []DeploymentSpecSpecificationRoutes         `json:"routes" tf:"routes"`
}

type DeploymentSpec struct {
	State *DeploymentSpecResource `json:"state,omitempty" tf:"-"`

	Resource DeploymentSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DeploymentSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	GatewayID    *string           `json:"gatewayID" tf:"gateway_id"`
	// +optional
	LifecycleDetails *string `json:"lifecycleDetails,omitempty" tf:"lifecycle_details"`
	PathPrefix       *string `json:"pathPrefix" tf:"path_prefix"`
	// +optional
	Specification *DeploymentSpecSpecification `json:"specification,omitempty" tf:"specification"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
}

type DeploymentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DeploymentList is a list of Deployments
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Deployment CRD objects
	Items []Deployment `json:"items,omitempty"`
}
