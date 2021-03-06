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

type WaasPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WaasPolicySpec   `json:"spec,omitempty"`
	Status            WaasPolicyStatus `json:"status,omitempty"`
}

type WaasPolicySpecOriginGroupsOriginGroup struct {
	Origin *string `json:"origin" tf:"origin"`
	// +optional
	Weight *int64 `json:"weight,omitempty" tf:"weight"`
}

type WaasPolicySpecOriginGroups struct {
	Label       *string                                 `json:"label" tf:"label"`
	OriginGroup []WaasPolicySpecOriginGroupsOriginGroup `json:"originGroup" tf:"origin_group"`
}

type WaasPolicySpecOriginsCustomHeaders struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type WaasPolicySpecOrigins struct {
	// +optional
	CustomHeaders []WaasPolicySpecOriginsCustomHeaders `json:"customHeaders,omitempty" tf:"custom_headers"`
	// +optional
	HttpPort *int64 `json:"httpPort,omitempty" tf:"http_port"`
	// +optional
	HttpsPort *int64  `json:"httpsPort,omitempty" tf:"https_port"`
	Label     *string `json:"label" tf:"label"`
	Uri       *string `json:"uri" tf:"uri"`
}

type WaasPolicySpecPolicyConfigHealthChecks struct {
	// +optional
	ExpectedResponseCodeGroup []string `json:"expectedResponseCodeGroup,omitempty" tf:"expected_response_code_group"`
	// +optional
	ExpectedResponseText *string `json:"expectedResponseText,omitempty" tf:"expected_response_text"`
	// +optional
	Headers map[string]string `json:"headers,omitempty" tf:"headers"`
	// +optional
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold"`
	// +optional
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds"`
	// +optional
	IsEnabled *bool `json:"isEnabled,omitempty" tf:"is_enabled"`
	// +optional
	IsResponseTextCheckEnabled *bool `json:"isResponseTextCheckEnabled,omitempty" tf:"is_response_text_check_enabled"`
	// +optional
	Method *string `json:"method,omitempty" tf:"method"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds"`
	// +optional
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold"`
}

type WaasPolicySpecPolicyConfigLoadBalancingMethod struct {
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// +optional
	ExpirationTimeInSeconds *int64  `json:"expirationTimeInSeconds,omitempty" tf:"expiration_time_in_seconds"`
	Method                  *string `json:"method" tf:"method"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type WaasPolicySpecPolicyConfig struct {
	// +optional
	CertificateID *string `json:"certificateID,omitempty" tf:"certificate_id"`
	// +optional
	CipherGroup *string `json:"cipherGroup,omitempty" tf:"cipher_group"`
	// +optional
	ClientAddressHeader *string `json:"clientAddressHeader,omitempty" tf:"client_address_header"`
	// +optional
	HealthChecks *WaasPolicySpecPolicyConfigHealthChecks `json:"healthChecks,omitempty" tf:"health_checks"`
	// +optional
	IsBehindCdn *bool `json:"isBehindCdn,omitempty" tf:"is_behind_cdn"`
	// +optional
	IsCacheControlRespected *bool `json:"isCacheControlRespected,omitempty" tf:"is_cache_control_respected"`
	// +optional
	IsHTTPSEnabled *bool `json:"isHTTPSEnabled,omitempty" tf:"is_https_enabled"`
	// +optional
	IsHTTPSForced *bool `json:"isHTTPSForced,omitempty" tf:"is_https_forced"`
	// +optional
	IsOriginCompressionEnabled *bool `json:"isOriginCompressionEnabled,omitempty" tf:"is_origin_compression_enabled"`
	// +optional
	IsResponseBufferingEnabled *bool `json:"isResponseBufferingEnabled,omitempty" tf:"is_response_buffering_enabled"`
	// +optional
	IsSniEnabled *bool `json:"isSniEnabled,omitempty" tf:"is_sni_enabled"`
	// +optional
	LoadBalancingMethod *WaasPolicySpecPolicyConfigLoadBalancingMethod `json:"loadBalancingMethod,omitempty" tf:"load_balancing_method"`
	// +optional
	TlsProtocols []string `json:"tlsProtocols,omitempty" tf:"tls_protocols"`
	// +optional
	WebsocketPathPrefixes []string `json:"websocketPathPrefixes,omitempty" tf:"websocket_path_prefixes"`
}

type WaasPolicySpecWafConfigAccessRulesCriteria struct {
	Condition *string `json:"condition" tf:"condition"`
	// +optional
	IsCaseSensitive *bool   `json:"isCaseSensitive,omitempty" tf:"is_case_sensitive"`
	Value           *string `json:"value" tf:"value"`
}

type WaasPolicySpecWafConfigAccessRulesResponseHeaderManipulation struct {
	Action *string `json:"action" tf:"action"`
	Header *string `json:"header" tf:"header"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type WaasPolicySpecWafConfigAccessRules struct {
	Action *string `json:"action" tf:"action"`
	// +optional
	BlockAction *string `json:"blockAction,omitempty" tf:"block_action"`
	// +optional
	BlockErrorPageCode *string `json:"blockErrorPageCode,omitempty" tf:"block_error_page_code"`
	// +optional
	BlockErrorPageDescription *string `json:"blockErrorPageDescription,omitempty" tf:"block_error_page_description"`
	// +optional
	BlockErrorPageMessage *string `json:"blockErrorPageMessage,omitempty" tf:"block_error_page_message"`
	// +optional
	BlockResponseCode *int64 `json:"blockResponseCode,omitempty" tf:"block_response_code"`
	// +optional
	BypassChallenges []string `json:"bypassChallenges,omitempty" tf:"bypass_challenges"`
	// +optional
	CaptchaFooter *string `json:"captchaFooter,omitempty" tf:"captcha_footer"`
	// +optional
	CaptchaHeader *string `json:"captchaHeader,omitempty" tf:"captcha_header"`
	// +optional
	CaptchaSubmitLabel *string `json:"captchaSubmitLabel,omitempty" tf:"captcha_submit_label"`
	// +optional
	CaptchaTitle *string                                      `json:"captchaTitle,omitempty" tf:"captcha_title"`
	Criteria     []WaasPolicySpecWafConfigAccessRulesCriteria `json:"criteria" tf:"criteria"`
	Name         *string                                      `json:"name" tf:"name"`
	// +optional
	RedirectResponseCode *string `json:"redirectResponseCode,omitempty" tf:"redirect_response_code"`
	// +optional
	RedirectURL *string `json:"redirectURL,omitempty" tf:"redirect_url"`
	// +optional
	ResponseHeaderManipulation []WaasPolicySpecWafConfigAccessRulesResponseHeaderManipulation `json:"responseHeaderManipulation,omitempty" tf:"response_header_manipulation"`
}

type WaasPolicySpecWafConfigAddressRateLimiting struct {
	// +optional
	AllowedRatePerAddress *int64 `json:"allowedRatePerAddress,omitempty" tf:"allowed_rate_per_address"`
	// +optional
	BlockResponseCode *int64 `json:"blockResponseCode,omitempty" tf:"block_response_code"`
	IsEnabled         *bool  `json:"isEnabled" tf:"is_enabled"`
	// +optional
	MaxDelayedCountPerAddress *int64 `json:"maxDelayedCountPerAddress,omitempty" tf:"max_delayed_count_per_address"`
}

type WaasPolicySpecWafConfigCachingRulesCriteria struct {
	Condition *string `json:"condition" tf:"condition"`
	Value     *string `json:"value" tf:"value"`
}

type WaasPolicySpecWafConfigCachingRules struct {
	Action *string `json:"action" tf:"action"`
	// +optional
	CachingDuration *string `json:"cachingDuration,omitempty" tf:"caching_duration"`
	// +optional
	ClientCachingDuration *string                                       `json:"clientCachingDuration,omitempty" tf:"client_caching_duration"`
	Criteria              []WaasPolicySpecWafConfigCachingRulesCriteria `json:"criteria" tf:"criteria"`
	// +optional
	IsClientCachingEnabled *bool `json:"isClientCachingEnabled,omitempty" tf:"is_client_caching_enabled"`
	// +optional
	Key  *string `json:"key,omitempty" tf:"key"`
	Name *string `json:"name" tf:"name"`
}

type WaasPolicySpecWafConfigCaptchas struct {
	FailureMessage *string `json:"failureMessage" tf:"failure_message"`
	// +optional
	FooterText *string `json:"footerText,omitempty" tf:"footer_text"`
	// +optional
	HeaderText                 *string `json:"headerText,omitempty" tf:"header_text"`
	SessionExpirationInSeconds *int64  `json:"sessionExpirationInSeconds" tf:"session_expiration_in_seconds"`
	SubmitLabel                *string `json:"submitLabel" tf:"submit_label"`
	Title                      *string `json:"title" tf:"title"`
	Url                        *string `json:"url" tf:"url"`
}

type WaasPolicySpecWafConfigCustomProtectionRulesExclusions struct {
	// +optional
	Exclusions []string `json:"exclusions,omitempty" tf:"exclusions"`
	// +optional
	Target *string `json:"target,omitempty" tf:"target"`
}

type WaasPolicySpecWafConfigCustomProtectionRules struct {
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// +optional
	Exclusions []WaasPolicySpecWafConfigCustomProtectionRulesExclusions `json:"exclusions,omitempty" tf:"exclusions"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
}

type WaasPolicySpecWafConfigDeviceFingerprintChallengeChallengeSettings struct {
	// +optional
	BlockAction *string `json:"blockAction,omitempty" tf:"block_action"`
	// +optional
	BlockErrorPageCode *string `json:"blockErrorPageCode,omitempty" tf:"block_error_page_code"`
	// +optional
	BlockErrorPageDescription *string `json:"blockErrorPageDescription,omitempty" tf:"block_error_page_description"`
	// +optional
	BlockErrorPageMessage *string `json:"blockErrorPageMessage,omitempty" tf:"block_error_page_message"`
	// +optional
	BlockResponseCode *int64 `json:"blockResponseCode,omitempty" tf:"block_response_code"`
	// +optional
	CaptchaFooter *string `json:"captchaFooter,omitempty" tf:"captcha_footer"`
	// +optional
	CaptchaHeader *string `json:"captchaHeader,omitempty" tf:"captcha_header"`
	// +optional
	CaptchaSubmitLabel *string `json:"captchaSubmitLabel,omitempty" tf:"captcha_submit_label"`
	// +optional
	CaptchaTitle *string `json:"captchaTitle,omitempty" tf:"captcha_title"`
}

type WaasPolicySpecWafConfigDeviceFingerprintChallenge struct {
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// +optional
	ActionExpirationInSeconds *int64 `json:"actionExpirationInSeconds,omitempty" tf:"action_expiration_in_seconds"`
	// +optional
	ChallengeSettings *WaasPolicySpecWafConfigDeviceFingerprintChallengeChallengeSettings `json:"challengeSettings,omitempty" tf:"challenge_settings"`
	// +optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`
	// +optional
	FailureThresholdExpirationInSeconds *int64 `json:"failureThresholdExpirationInSeconds,omitempty" tf:"failure_threshold_expiration_in_seconds"`
	IsEnabled                           *bool  `json:"isEnabled" tf:"is_enabled"`
	// +optional
	MaxAddressCount *int64 `json:"maxAddressCount,omitempty" tf:"max_address_count"`
	// +optional
	MaxAddressCountExpirationInSeconds *int64 `json:"maxAddressCountExpirationInSeconds,omitempty" tf:"max_address_count_expiration_in_seconds"`
}

type WaasPolicySpecWafConfigHumanInteractionChallengeChallengeSettings struct {
	// +optional
	BlockAction *string `json:"blockAction,omitempty" tf:"block_action"`
	// +optional
	BlockErrorPageCode *string `json:"blockErrorPageCode,omitempty" tf:"block_error_page_code"`
	// +optional
	BlockErrorPageDescription *string `json:"blockErrorPageDescription,omitempty" tf:"block_error_page_description"`
	// +optional
	BlockErrorPageMessage *string `json:"blockErrorPageMessage,omitempty" tf:"block_error_page_message"`
	// +optional
	BlockResponseCode *int64 `json:"blockResponseCode,omitempty" tf:"block_response_code"`
	// +optional
	CaptchaFooter *string `json:"captchaFooter,omitempty" tf:"captcha_footer"`
	// +optional
	CaptchaHeader *string `json:"captchaHeader,omitempty" tf:"captcha_header"`
	// +optional
	CaptchaSubmitLabel *string `json:"captchaSubmitLabel,omitempty" tf:"captcha_submit_label"`
	// +optional
	CaptchaTitle *string `json:"captchaTitle,omitempty" tf:"captcha_title"`
}

type WaasPolicySpecWafConfigHumanInteractionChallengeSetHTTPHeader struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type WaasPolicySpecWafConfigHumanInteractionChallenge struct {
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// +optional
	ActionExpirationInSeconds *int64 `json:"actionExpirationInSeconds,omitempty" tf:"action_expiration_in_seconds"`
	// +optional
	ChallengeSettings *WaasPolicySpecWafConfigHumanInteractionChallengeChallengeSettings `json:"challengeSettings,omitempty" tf:"challenge_settings"`
	// +optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`
	// +optional
	FailureThresholdExpirationInSeconds *int64 `json:"failureThresholdExpirationInSeconds,omitempty" tf:"failure_threshold_expiration_in_seconds"`
	// +optional
	InteractionThreshold *int64 `json:"interactionThreshold,omitempty" tf:"interaction_threshold"`
	IsEnabled            *bool  `json:"isEnabled" tf:"is_enabled"`
	// +optional
	IsNATEnabled *bool `json:"isNATEnabled,omitempty" tf:"is_nat_enabled"`
	// +optional
	RecordingPeriodInSeconds *int64 `json:"recordingPeriodInSeconds,omitempty" tf:"recording_period_in_seconds"`
	// +optional
	SetHTTPHeader *WaasPolicySpecWafConfigHumanInteractionChallengeSetHTTPHeader `json:"setHTTPHeader,omitempty" tf:"set_http_header"`
}

type WaasPolicySpecWafConfigJsChallengeChallengeSettings struct {
	// +optional
	BlockAction *string `json:"blockAction,omitempty" tf:"block_action"`
	// +optional
	BlockErrorPageCode *string `json:"blockErrorPageCode,omitempty" tf:"block_error_page_code"`
	// +optional
	BlockErrorPageDescription *string `json:"blockErrorPageDescription,omitempty" tf:"block_error_page_description"`
	// +optional
	BlockErrorPageMessage *string `json:"blockErrorPageMessage,omitempty" tf:"block_error_page_message"`
	// +optional
	BlockResponseCode *int64 `json:"blockResponseCode,omitempty" tf:"block_response_code"`
	// +optional
	CaptchaFooter *string `json:"captchaFooter,omitempty" tf:"captcha_footer"`
	// +optional
	CaptchaHeader *string `json:"captchaHeader,omitempty" tf:"captcha_header"`
	// +optional
	CaptchaSubmitLabel *string `json:"captchaSubmitLabel,omitempty" tf:"captcha_submit_label"`
	// +optional
	CaptchaTitle *string `json:"captchaTitle,omitempty" tf:"captcha_title"`
}

type WaasPolicySpecWafConfigJsChallengeCriteria struct {
	Condition *string `json:"condition" tf:"condition"`
	// +optional
	IsCaseSensitive *bool   `json:"isCaseSensitive,omitempty" tf:"is_case_sensitive"`
	Value           *string `json:"value" tf:"value"`
}

type WaasPolicySpecWafConfigJsChallengeSetHTTPHeader struct {
	Name  *string `json:"name" tf:"name"`
	Value *string `json:"value" tf:"value"`
}

type WaasPolicySpecWafConfigJsChallenge struct {
	// +optional
	Action *string `json:"action,omitempty" tf:"action"`
	// +optional
	ActionExpirationInSeconds *int64 `json:"actionExpirationInSeconds,omitempty" tf:"action_expiration_in_seconds"`
	// +optional
	AreRedirectsChallenged *bool `json:"areRedirectsChallenged,omitempty" tf:"are_redirects_challenged"`
	// +optional
	ChallengeSettings *WaasPolicySpecWafConfigJsChallengeChallengeSettings `json:"challengeSettings,omitempty" tf:"challenge_settings"`
	// +optional
	Criteria []WaasPolicySpecWafConfigJsChallengeCriteria `json:"criteria,omitempty" tf:"criteria"`
	// +optional
	FailureThreshold *int64 `json:"failureThreshold,omitempty" tf:"failure_threshold"`
	IsEnabled        *bool  `json:"isEnabled" tf:"is_enabled"`
	// +optional
	IsNATEnabled *bool `json:"isNATEnabled,omitempty" tf:"is_nat_enabled"`
	// +optional
	SetHTTPHeader *WaasPolicySpecWafConfigJsChallengeSetHTTPHeader `json:"setHTTPHeader,omitempty" tf:"set_http_header"`
}

type WaasPolicySpecWafConfigProtectionSettings struct {
	// +optional
	AllowedHTTPMethods []string `json:"allowedHTTPMethods,omitempty" tf:"allowed_http_methods"`
	// +optional
	BlockAction *string `json:"blockAction,omitempty" tf:"block_action"`
	// +optional
	BlockErrorPageCode *string `json:"blockErrorPageCode,omitempty" tf:"block_error_page_code"`
	// +optional
	BlockErrorPageDescription *string `json:"blockErrorPageDescription,omitempty" tf:"block_error_page_description"`
	// +optional
	BlockErrorPageMessage *string `json:"blockErrorPageMessage,omitempty" tf:"block_error_page_message"`
	// +optional
	BlockResponseCode *int64 `json:"blockResponseCode,omitempty" tf:"block_response_code"`
	// +optional
	IsResponseInspected *bool `json:"isResponseInspected,omitempty" tf:"is_response_inspected"`
	// +optional
	MaxArgumentCount *int64 `json:"maxArgumentCount,omitempty" tf:"max_argument_count"`
	// +optional
	MaxNameLengthPerArgument *int64 `json:"maxNameLengthPerArgument,omitempty" tf:"max_name_length_per_argument"`
	// +optional
	MaxResponseSizeInKiB *int64 `json:"maxResponseSizeInKiB,omitempty" tf:"max_response_size_in_ki_b"`
	// +optional
	MaxTotalNameLengthOfArguments *int64 `json:"maxTotalNameLengthOfArguments,omitempty" tf:"max_total_name_length_of_arguments"`
	// +optional
	MediaTypes []string `json:"mediaTypes,omitempty" tf:"media_types"`
	// +optional
	RecommendationsPeriodInDays *int64 `json:"recommendationsPeriodInDays,omitempty" tf:"recommendations_period_in_days"`
}

type WaasPolicySpecWafConfigWhitelists struct {
	// +optional
	AddressLists []string `json:"addressLists,omitempty" tf:"address_lists"`
	// +optional
	Addresses []string `json:"addresses,omitempty" tf:"addresses"`
	Name      *string  `json:"name" tf:"name"`
}

type WaasPolicySpecWafConfig struct {
	// +optional
	AccessRules []WaasPolicySpecWafConfigAccessRules `json:"accessRules,omitempty" tf:"access_rules"`
	// +optional
	AddressRateLimiting *WaasPolicySpecWafConfigAddressRateLimiting `json:"addressRateLimiting,omitempty" tf:"address_rate_limiting"`
	// +optional
	CachingRules []WaasPolicySpecWafConfigCachingRules `json:"cachingRules,omitempty" tf:"caching_rules"`
	// +optional
	Captchas []WaasPolicySpecWafConfigCaptchas `json:"captchas,omitempty" tf:"captchas"`
	// +optional
	CustomProtectionRules []WaasPolicySpecWafConfigCustomProtectionRules `json:"customProtectionRules,omitempty" tf:"custom_protection_rules"`
	// +optional
	DeviceFingerprintChallenge *WaasPolicySpecWafConfigDeviceFingerprintChallenge `json:"deviceFingerprintChallenge,omitempty" tf:"device_fingerprint_challenge"`
	// +optional
	HumanInteractionChallenge *WaasPolicySpecWafConfigHumanInteractionChallenge `json:"humanInteractionChallenge,omitempty" tf:"human_interaction_challenge"`
	// +optional
	JsChallenge *WaasPolicySpecWafConfigJsChallenge `json:"jsChallenge,omitempty" tf:"js_challenge"`
	// +optional
	Origin *string `json:"origin,omitempty" tf:"origin"`
	// +optional
	OriginGroups []string `json:"originGroups,omitempty" tf:"origin_groups"`
	// +optional
	ProtectionSettings *WaasPolicySpecWafConfigProtectionSettings `json:"protectionSettings,omitempty" tf:"protection_settings"`
	// +optional
	Whitelists []WaasPolicySpecWafConfigWhitelists `json:"whitelists,omitempty" tf:"whitelists"`
}

type WaasPolicySpec struct {
	State *WaasPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource WaasPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type WaasPolicySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AdditionalDomains []string `json:"additionalDomains,omitempty" tf:"additional_domains"`
	// +optional
	Cname         *string `json:"cname,omitempty" tf:"cname"`
	CompartmentID *string `json:"compartmentID" tf:"compartment_id"`
	// +optional
	DefinedTags map[string]string `json:"definedTags,omitempty" tf:"defined_tags"`
	// +optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`
	Domain      *string `json:"domain" tf:"domain"`
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty" tf:"freeform_tags"`
	// +optional
	OriginGroups []WaasPolicySpecOriginGroups `json:"originGroups,omitempty" tf:"origin_groups"`
	// +optional
	Origins []WaasPolicySpecOrigins `json:"origins,omitempty" tf:"origins"`
	// +optional
	PolicyConfig *WaasPolicySpecPolicyConfig `json:"policyConfig,omitempty" tf:"policy_config"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created"`
	// +optional
	WafConfig *WaasPolicySpecWafConfig `json:"wafConfig,omitempty" tf:"waf_config"`
}

type WaasPolicyStatus struct {
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

// WaasPolicyList is a list of WaasPolicys
type WaasPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WaasPolicy CRD objects
	Items []WaasPolicy `json:"items,omitempty"`
}
