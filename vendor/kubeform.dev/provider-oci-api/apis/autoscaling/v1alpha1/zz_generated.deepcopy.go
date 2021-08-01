// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfiguration) DeepCopyInto(out *AutoScalingConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfiguration.
func (in *AutoScalingConfiguration) DeepCopy() *AutoScalingConfiguration {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoScalingConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationList) DeepCopyInto(out *AutoScalingConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AutoScalingConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationList.
func (in *AutoScalingConfigurationList) DeepCopy() *AutoScalingConfigurationList {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AutoScalingConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpec) DeepCopyInto(out *AutoScalingConfigurationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AutoScalingConfigurationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpec.
func (in *AutoScalingConfigurationSpec) DeepCopy() *AutoScalingConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecAutoScalingResources) DeepCopyInto(out *AutoScalingConfigurationSpecAutoScalingResources) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecAutoScalingResources.
func (in *AutoScalingConfigurationSpecAutoScalingResources) DeepCopy() *AutoScalingConfigurationSpecAutoScalingResources {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecAutoScalingResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecPolicies) DeepCopyInto(out *AutoScalingConfigurationSpecPolicies) {
	*out = *in
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(AutoScalingConfigurationSpecPoliciesCapacity)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ExecutionSchedule != nil {
		in, out := &in.ExecutionSchedule, &out.ExecutionSchedule
		*out = new(AutoScalingConfigurationSpecPoliciesExecutionSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
	if in.ResourceAction != nil {
		in, out := &in.ResourceAction, &out.ResourceAction
		*out = new(AutoScalingConfigurationSpecPoliciesResourceAction)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]AutoScalingConfigurationSpecPoliciesRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecPolicies.
func (in *AutoScalingConfigurationSpecPolicies) DeepCopy() *AutoScalingConfigurationSpecPolicies {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecPoliciesCapacity) DeepCopyInto(out *AutoScalingConfigurationSpecPoliciesCapacity) {
	*out = *in
	if in.Initial != nil {
		in, out := &in.Initial, &out.Initial
		*out = new(int64)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(int64)
		**out = **in
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecPoliciesCapacity.
func (in *AutoScalingConfigurationSpecPoliciesCapacity) DeepCopy() *AutoScalingConfigurationSpecPoliciesCapacity {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecPoliciesCapacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecPoliciesExecutionSchedule) DeepCopyInto(out *AutoScalingConfigurationSpecPoliciesExecutionSchedule) {
	*out = *in
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Timezone != nil {
		in, out := &in.Timezone, &out.Timezone
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecPoliciesExecutionSchedule.
func (in *AutoScalingConfigurationSpecPoliciesExecutionSchedule) DeepCopy() *AutoScalingConfigurationSpecPoliciesExecutionSchedule {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecPoliciesExecutionSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecPoliciesResourceAction) DeepCopyInto(out *AutoScalingConfigurationSpecPoliciesResourceAction) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.ActionType != nil {
		in, out := &in.ActionType, &out.ActionType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecPoliciesResourceAction.
func (in *AutoScalingConfigurationSpecPoliciesResourceAction) DeepCopy() *AutoScalingConfigurationSpecPoliciesResourceAction {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecPoliciesResourceAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecPoliciesRules) DeepCopyInto(out *AutoScalingConfigurationSpecPoliciesRules) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(AutoScalingConfigurationSpecPoliciesRulesAction)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(AutoScalingConfigurationSpecPoliciesRulesMetric)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecPoliciesRules.
func (in *AutoScalingConfigurationSpecPoliciesRules) DeepCopy() *AutoScalingConfigurationSpecPoliciesRules {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecPoliciesRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecPoliciesRulesAction) DeepCopyInto(out *AutoScalingConfigurationSpecPoliciesRulesAction) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecPoliciesRulesAction.
func (in *AutoScalingConfigurationSpecPoliciesRulesAction) DeepCopy() *AutoScalingConfigurationSpecPoliciesRulesAction {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecPoliciesRulesAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecPoliciesRulesMetric) DeepCopyInto(out *AutoScalingConfigurationSpecPoliciesRulesMetric) {
	*out = *in
	if in.MetricType != nil {
		in, out := &in.MetricType, &out.MetricType
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(AutoScalingConfigurationSpecPoliciesRulesMetricThreshold)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecPoliciesRulesMetric.
func (in *AutoScalingConfigurationSpecPoliciesRulesMetric) DeepCopy() *AutoScalingConfigurationSpecPoliciesRulesMetric {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecPoliciesRulesMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecPoliciesRulesMetricThreshold) DeepCopyInto(out *AutoScalingConfigurationSpecPoliciesRulesMetricThreshold) {
	*out = *in
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecPoliciesRulesMetricThreshold.
func (in *AutoScalingConfigurationSpecPoliciesRulesMetricThreshold) DeepCopy() *AutoScalingConfigurationSpecPoliciesRulesMetricThreshold {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecPoliciesRulesMetricThreshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationSpecResource) DeepCopyInto(out *AutoScalingConfigurationSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoScalingResources != nil {
		in, out := &in.AutoScalingResources, &out.AutoScalingResources
		*out = new(AutoScalingConfigurationSpecAutoScalingResources)
		(*in).DeepCopyInto(*out)
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.CoolDownInSeconds != nil {
		in, out := &in.CoolDownInSeconds, &out.CoolDownInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.DefinedTags != nil {
		in, out := &in.DefinedTags, &out.DefinedTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MaxResourceCount != nil {
		in, out := &in.MaxResourceCount, &out.MaxResourceCount
		*out = new(int64)
		**out = **in
	}
	if in.MinResourceCount != nil {
		in, out := &in.MinResourceCount, &out.MinResourceCount
		*out = new(int64)
		**out = **in
	}
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]AutoScalingConfigurationSpecPolicies, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationSpecResource.
func (in *AutoScalingConfigurationSpecResource) DeepCopy() *AutoScalingConfigurationSpecResource {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingConfigurationStatus) DeepCopyInto(out *AutoScalingConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingConfigurationStatus.
func (in *AutoScalingConfigurationStatus) DeepCopy() *AutoScalingConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(AutoScalingConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}