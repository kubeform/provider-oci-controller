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

type SyntheticsMonitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SyntheticsMonitorSpec   `json:"spec,omitempty"`
	Status            SyntheticsMonitorStatus `json:"status,omitempty"`
}

type SyntheticsMonitorSpecConfigurationReqAuthenticationDetailsAuthHeaders struct {
	// +optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name"`
	// +optional
	HeaderValue *string `json:"headerValue,omitempty" tf:"header_value"`
}

type SyntheticsMonitorSpecConfigurationReqAuthenticationDetails struct {
	// +optional
	AuthHeaders []SyntheticsMonitorSpecConfigurationReqAuthenticationDetailsAuthHeaders `json:"authHeaders,omitempty" tf:"auth_headers"`
	// +optional
	AuthRequestMethod *string `json:"authRequestMethod,omitempty" tf:"auth_request_method"`
	// +optional
	AuthRequestPostBody *string `json:"authRequestPostBody,omitempty" tf:"auth_request_post_body"`
	// +optional
	AuthToken *string `json:"authToken,omitempty" tf:"auth_token"`
	// +optional
	AuthURL *string `json:"authURL,omitempty" tf:"auth_url"`
	// +optional
	AuthUserName *string `json:"authUserName,omitempty" tf:"auth_user_name"`
	// +optional
	AuthUserPassword *string `json:"-" sensitive:"true" tf:"auth_user_password"`
	// +optional
	OauthScheme *string `json:"oauthScheme,omitempty" tf:"oauth_scheme"`
}

type SyntheticsMonitorSpecConfigurationRequestHeaders struct {
	// +optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name"`
	// +optional
	HeaderValue *string `json:"headerValue,omitempty" tf:"header_value"`
}

type SyntheticsMonitorSpecConfigurationRequestQueryParams struct {
	// +optional
	ParamName *string `json:"paramName,omitempty" tf:"param_name"`
	// +optional
	ParamValue *string `json:"paramValue,omitempty" tf:"param_value"`
}

type SyntheticsMonitorSpecConfigurationVerifyTexts struct {
	// +optional
	Text *string `json:"text,omitempty" tf:"text"`
}

type SyntheticsMonitorSpecConfiguration struct {
	// +optional
	ConfigType *string `json:"configType,omitempty" tf:"config_type"`
	// +optional
	IsCertificateValidationEnabled *bool `json:"isCertificateValidationEnabled,omitempty" tf:"is_certificate_validation_enabled"`
	// +optional
	IsFailureRetried *bool `json:"isFailureRetried,omitempty" tf:"is_failure_retried"`
	// +optional
	IsRedirectionEnabled *bool `json:"isRedirectionEnabled,omitempty" tf:"is_redirection_enabled"`
	// +optional
	ReqAuthenticationDetails *SyntheticsMonitorSpecConfigurationReqAuthenticationDetails `json:"reqAuthenticationDetails,omitempty" tf:"req_authentication_details"`
	// +optional
	ReqAuthenticationScheme *string `json:"reqAuthenticationScheme,omitempty" tf:"req_authentication_scheme"`
	// +optional
	RequestHeaders []SyntheticsMonitorSpecConfigurationRequestHeaders `json:"requestHeaders,omitempty" tf:"request_headers"`
	// +optional
	RequestMethod *string `json:"requestMethod,omitempty" tf:"request_method"`
	// +optional
	RequestPostBody *string `json:"requestPostBody,omitempty" tf:"request_post_body"`
	// +optional
	RequestQueryParams []SyntheticsMonitorSpecConfigurationRequestQueryParams `json:"requestQueryParams,omitempty" tf:"request_query_params"`
	// +optional
	// +kubebuilder:validation:MaxItems=3
	VerifyResponseCodes []string `json:"verifyResponseCodes,omitempty" tf:"verify_response_codes"`
	// +optional
	VerifyResponseContent *string `json:"verifyResponseContent,omitempty" tf:"verify_response_content"`
	// +optional
	VerifyTexts []SyntheticsMonitorSpecConfigurationVerifyTexts `json:"verifyTexts,omitempty" tf:"verify_texts"`
}

type SyntheticsMonitorSpecScriptParametersMonitorScriptParameter struct {
	// +optional
	ParamName *string `json:"paramName,omitempty" tf:"param_name"`
	// +optional
	ParamValue *string `json:"paramValue,omitempty" tf:"param_value"`
}

type SyntheticsMonitorSpecScriptParameters struct {
	// +optional
	IsOverwritten *bool `json:"isOverwritten,omitempty" tf:"is_overwritten"`
	// +optional
	IsSecret *bool `json:"isSecret,omitempty" tf:"is_secret"`
	// +optional
	MonitorScriptParameter *SyntheticsMonitorSpecScriptParametersMonitorScriptParameter `json:"monitorScriptParameter,omitempty" tf:"monitor_script_parameter"`
	ParamName              *string                                                      `json:"paramName" tf:"param_name"`
	ParamValue             *string                                                      `json:"paramValue" tf:"param_value"`
}

type SyntheticsMonitorSpec struct {
	State *SyntheticsMonitorSpecResource `json:"state,omitempty" tf:"-"`

	Resource SyntheticsMonitorSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type SyntheticsMonitorSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApmDomainID *string `json:"apmDomainID" tf:"apm_domain_id"`
	// +optional
	Configuration *SyntheticsMonitorSpecConfiguration `json:"configuration,omitempty" tf:"configuration"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	DisplayName *string           `json:"displayName" tf:"display_name"`
	// +optional
	FreeformTags            map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	MonitorType             *string           `json:"monitorType" tf:"monitor_type"`
	RepeatIntervalInSeconds *int64            `json:"repeatIntervalInSeconds" tf:"repeat_interval_in_seconds"`
	// +optional
	ScriptID *string `json:"scriptID,omitempty" tf:"script_id"`
	// +optional
	ScriptName *string `json:"scriptName,omitempty" tf:"script_name"`
	// +optional
	ScriptParameters []SyntheticsMonitorSpecScriptParameters `json:"scriptParameters,omitempty" tf:"script_parameters"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Target *string `json:"target,omitempty" tf:"target"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	TimeUpdated *string `json:"timeUpdated,omitempty" tf:"time_updated"`
	// +optional
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds"`
	// +optional
	VantagePointCount *int64 `json:"vantagePointCount,omitempty" tf:"vantage_point_count"`
	// +kubebuilder:validation:MaxItems=50
	// +kubebuilder:validation:MinItems=1
	VantagePoints []string `json:"vantagePoints" tf:"vantage_points"`
}

type SyntheticsMonitorStatus struct {
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

// SyntheticsMonitorList is a list of SyntheticsMonitors
type SyntheticsMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SyntheticsMonitor CRD objects
	Items []SyntheticsMonitor `json:"items,omitempty"`
}