//go:build !ignore_autogenerated
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

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPMonitor) DeepCopyInto(out *ChecksHTTPMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPMonitor.
func (in *ChecksHTTPMonitor) DeepCopy() *ChecksHTTPMonitor {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChecksHTTPMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPMonitorList) DeepCopyInto(out *ChecksHTTPMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChecksHTTPMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPMonitorList.
func (in *ChecksHTTPMonitorList) DeepCopy() *ChecksHTTPMonitorList {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChecksHTTPMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPMonitorSpec) DeepCopyInto(out *ChecksHTTPMonitorSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ChecksHTTPMonitorSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPMonitorSpec.
func (in *ChecksHTTPMonitorSpec) DeepCopy() *ChecksHTTPMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPMonitorSpecResource) DeepCopyInto(out *ChecksHTTPMonitorSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
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
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HomeRegion != nil {
		in, out := &in.HomeRegion, &out.HomeRegion
		*out = new(string)
		**out = **in
	}
	if in.IntervalInSeconds != nil {
		in, out := &in.IntervalInSeconds, &out.IntervalInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ResultsURL != nil {
		in, out := &in.ResultsURL, &out.ResultsURL
		*out = new(string)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeoutInSeconds != nil {
		in, out := &in.TimeoutInSeconds, &out.TimeoutInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.VantagePointNames != nil {
		in, out := &in.VantagePointNames, &out.VantagePointNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPMonitorSpecResource.
func (in *ChecksHTTPMonitorSpecResource) DeepCopy() *ChecksHTTPMonitorSpecResource {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPMonitorSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPMonitorStatus) DeepCopyInto(out *ChecksHTTPMonitorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPMonitorStatus.
func (in *ChecksHTTPMonitorStatus) DeepCopy() *ChecksHTTPMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPProbe) DeepCopyInto(out *ChecksHTTPProbe) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPProbe.
func (in *ChecksHTTPProbe) DeepCopy() *ChecksHTTPProbe {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChecksHTTPProbe) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPProbeList) DeepCopyInto(out *ChecksHTTPProbeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChecksHTTPProbe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPProbeList.
func (in *ChecksHTTPProbeList) DeepCopy() *ChecksHTTPProbeList {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPProbeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChecksHTTPProbeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPProbeSpec) DeepCopyInto(out *ChecksHTTPProbeSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ChecksHTTPProbeSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPProbeSpec.
func (in *ChecksHTTPProbeSpec) DeepCopy() *ChecksHTTPProbeSpec {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPProbeSpecResource) DeepCopyInto(out *ChecksHTTPProbeSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HomeRegion != nil {
		in, out := &in.HomeRegion, &out.HomeRegion
		*out = new(string)
		**out = **in
	}
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ResultsURL != nil {
		in, out := &in.ResultsURL, &out.ResultsURL
		*out = new(string)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeoutInSeconds != nil {
		in, out := &in.TimeoutInSeconds, &out.TimeoutInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.VantagePointNames != nil {
		in, out := &in.VantagePointNames, &out.VantagePointNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPProbeSpecResource.
func (in *ChecksHTTPProbeSpecResource) DeepCopy() *ChecksHTTPProbeSpecResource {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPProbeSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksHTTPProbeStatus) DeepCopyInto(out *ChecksHTTPProbeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksHTTPProbeStatus.
func (in *ChecksHTTPProbeStatus) DeepCopy() *ChecksHTTPProbeStatus {
	if in == nil {
		return nil
	}
	out := new(ChecksHTTPProbeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingMonitor) DeepCopyInto(out *ChecksPingMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingMonitor.
func (in *ChecksPingMonitor) DeepCopy() *ChecksPingMonitor {
	if in == nil {
		return nil
	}
	out := new(ChecksPingMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChecksPingMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingMonitorList) DeepCopyInto(out *ChecksPingMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChecksPingMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingMonitorList.
func (in *ChecksPingMonitorList) DeepCopy() *ChecksPingMonitorList {
	if in == nil {
		return nil
	}
	out := new(ChecksPingMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChecksPingMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingMonitorSpec) DeepCopyInto(out *ChecksPingMonitorSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ChecksPingMonitorSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingMonitorSpec.
func (in *ChecksPingMonitorSpec) DeepCopy() *ChecksPingMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(ChecksPingMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingMonitorSpecResource) DeepCopyInto(out *ChecksPingMonitorSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
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
	if in.HomeRegion != nil {
		in, out := &in.HomeRegion, &out.HomeRegion
		*out = new(string)
		**out = **in
	}
	if in.IntervalInSeconds != nil {
		in, out := &in.IntervalInSeconds, &out.IntervalInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.IsEnabled != nil {
		in, out := &in.IsEnabled, &out.IsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ResultsURL != nil {
		in, out := &in.ResultsURL, &out.ResultsURL
		*out = new(string)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeoutInSeconds != nil {
		in, out := &in.TimeoutInSeconds, &out.TimeoutInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.VantagePointNames != nil {
		in, out := &in.VantagePointNames, &out.VantagePointNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingMonitorSpecResource.
func (in *ChecksPingMonitorSpecResource) DeepCopy() *ChecksPingMonitorSpecResource {
	if in == nil {
		return nil
	}
	out := new(ChecksPingMonitorSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingMonitorStatus) DeepCopyInto(out *ChecksPingMonitorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingMonitorStatus.
func (in *ChecksPingMonitorStatus) DeepCopy() *ChecksPingMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(ChecksPingMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingProbe) DeepCopyInto(out *ChecksPingProbe) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingProbe.
func (in *ChecksPingProbe) DeepCopy() *ChecksPingProbe {
	if in == nil {
		return nil
	}
	out := new(ChecksPingProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChecksPingProbe) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingProbeList) DeepCopyInto(out *ChecksPingProbeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChecksPingProbe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingProbeList.
func (in *ChecksPingProbeList) DeepCopy() *ChecksPingProbeList {
	if in == nil {
		return nil
	}
	out := new(ChecksPingProbeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChecksPingProbeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingProbeSpec) DeepCopyInto(out *ChecksPingProbeSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ChecksPingProbeSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingProbeSpec.
func (in *ChecksPingProbeSpec) DeepCopy() *ChecksPingProbeSpec {
	if in == nil {
		return nil
	}
	out := new(ChecksPingProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingProbeSpecResource) DeepCopyInto(out *ChecksPingProbeSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.HomeRegion != nil {
		in, out := &in.HomeRegion, &out.HomeRegion
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ResultsURL != nil {
		in, out := &in.ResultsURL, &out.ResultsURL
		*out = new(string)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeoutInSeconds != nil {
		in, out := &in.TimeoutInSeconds, &out.TimeoutInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.VantagePointNames != nil {
		in, out := &in.VantagePointNames, &out.VantagePointNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingProbeSpecResource.
func (in *ChecksPingProbeSpecResource) DeepCopy() *ChecksPingProbeSpecResource {
	if in == nil {
		return nil
	}
	out := new(ChecksPingProbeSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChecksPingProbeStatus) DeepCopyInto(out *ChecksPingProbeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChecksPingProbeStatus.
func (in *ChecksPingProbeStatus) DeepCopy() *ChecksPingProbeStatus {
	if in == nil {
		return nil
	}
	out := new(ChecksPingProbeStatus)
	in.DeepCopyInto(out)
	return out
}
