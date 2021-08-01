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
func (in *LoadBalancerBackend) DeepCopyInto(out *LoadBalancerBackend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackend.
func (in *LoadBalancerBackend) DeepCopy() *LoadBalancerBackend {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerBackend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendList) DeepCopyInto(out *LoadBalancerBackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerBackend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendList.
func (in *LoadBalancerBackendList) DeepCopy() *LoadBalancerBackendList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerBackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendSet) DeepCopyInto(out *LoadBalancerBackendSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendSet.
func (in *LoadBalancerBackendSet) DeepCopy() *LoadBalancerBackendSet {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerBackendSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendSetList) DeepCopyInto(out *LoadBalancerBackendSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerBackendSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendSetList.
func (in *LoadBalancerBackendSetList) DeepCopy() *LoadBalancerBackendSetList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerBackendSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendSetSpec) DeepCopyInto(out *LoadBalancerBackendSetSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LoadBalancerBackendSetSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendSetSpec.
func (in *LoadBalancerBackendSetSpec) DeepCopy() *LoadBalancerBackendSetSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendSetSpecBackends) DeepCopyInto(out *LoadBalancerBackendSetSpecBackends) {
	*out = *in
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.IsBackup != nil {
		in, out := &in.IsBackup, &out.IsBackup
		*out = new(bool)
		**out = **in
	}
	if in.IsDrain != nil {
		in, out := &in.IsDrain, &out.IsDrain
		*out = new(bool)
		**out = **in
	}
	if in.IsOffline != nil {
		in, out := &in.IsOffline, &out.IsOffline
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.TargetID != nil {
		in, out := &in.TargetID, &out.TargetID
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendSetSpecBackends.
func (in *LoadBalancerBackendSetSpecBackends) DeepCopy() *LoadBalancerBackendSetSpecBackends {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendSetSpecBackends)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendSetSpecHealthChecker) DeepCopyInto(out *LoadBalancerBackendSetSpecHealthChecker) {
	*out = *in
	if in.IntervalInMillis != nil {
		in, out := &in.IntervalInMillis, &out.IntervalInMillis
		*out = new(int64)
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
	if in.RequestData != nil {
		in, out := &in.RequestData, &out.RequestData
		*out = new(string)
		**out = **in
	}
	if in.ResponseBodyRegex != nil {
		in, out := &in.ResponseBodyRegex, &out.ResponseBodyRegex
		*out = new(string)
		**out = **in
	}
	if in.ResponseData != nil {
		in, out := &in.ResponseData, &out.ResponseData
		*out = new(string)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(int64)
		**out = **in
	}
	if in.ReturnCode != nil {
		in, out := &in.ReturnCode, &out.ReturnCode
		*out = new(int64)
		**out = **in
	}
	if in.TimeoutInMillis != nil {
		in, out := &in.TimeoutInMillis, &out.TimeoutInMillis
		*out = new(int64)
		**out = **in
	}
	if in.UrlPath != nil {
		in, out := &in.UrlPath, &out.UrlPath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendSetSpecHealthChecker.
func (in *LoadBalancerBackendSetSpecHealthChecker) DeepCopy() *LoadBalancerBackendSetSpecHealthChecker {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendSetSpecHealthChecker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendSetSpecResource) DeepCopyInto(out *LoadBalancerBackendSetSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]LoadBalancerBackendSetSpecBackends, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HealthChecker != nil {
		in, out := &in.HealthChecker, &out.HealthChecker
		*out = new(LoadBalancerBackendSetSpecHealthChecker)
		(*in).DeepCopyInto(*out)
	}
	if in.IsPreserveSource != nil {
		in, out := &in.IsPreserveSource, &out.IsPreserveSource
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkLoadBalancerID != nil {
		in, out := &in.NetworkLoadBalancerID, &out.NetworkLoadBalancerID
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendSetSpecResource.
func (in *LoadBalancerBackendSetSpecResource) DeepCopy() *LoadBalancerBackendSetSpecResource {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendSetSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendSetStatus) DeepCopyInto(out *LoadBalancerBackendSetStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendSetStatus.
func (in *LoadBalancerBackendSetStatus) DeepCopy() *LoadBalancerBackendSetStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendSpec) DeepCopyInto(out *LoadBalancerBackendSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LoadBalancerBackendSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendSpec.
func (in *LoadBalancerBackendSpec) DeepCopy() *LoadBalancerBackendSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendSpecResource) DeepCopyInto(out *LoadBalancerBackendSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.BackendSetName != nil {
		in, out := &in.BackendSetName, &out.BackendSetName
		*out = new(string)
		**out = **in
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.IsBackup != nil {
		in, out := &in.IsBackup, &out.IsBackup
		*out = new(bool)
		**out = **in
	}
	if in.IsDrain != nil {
		in, out := &in.IsDrain, &out.IsDrain
		*out = new(bool)
		**out = **in
	}
	if in.IsOffline != nil {
		in, out := &in.IsOffline, &out.IsOffline
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkLoadBalancerID != nil {
		in, out := &in.NetworkLoadBalancerID, &out.NetworkLoadBalancerID
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.TargetID != nil {
		in, out := &in.TargetID, &out.TargetID
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendSpecResource.
func (in *LoadBalancerBackendSpecResource) DeepCopy() *LoadBalancerBackendSpecResource {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerBackendStatus) DeepCopyInto(out *LoadBalancerBackendStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerBackendStatus.
func (in *LoadBalancerBackendStatus) DeepCopy() *LoadBalancerBackendStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerBackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerListener) DeepCopyInto(out *LoadBalancerListener) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerListener.
func (in *LoadBalancerListener) DeepCopy() *LoadBalancerListener {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerListener) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerListenerList) DeepCopyInto(out *LoadBalancerListenerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerListener, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerListenerList.
func (in *LoadBalancerListenerList) DeepCopy() *LoadBalancerListenerList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerListenerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerListenerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerListenerSpec) DeepCopyInto(out *LoadBalancerListenerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LoadBalancerListenerSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerListenerSpec.
func (in *LoadBalancerListenerSpec) DeepCopy() *LoadBalancerListenerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerListenerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerListenerSpecResource) DeepCopyInto(out *LoadBalancerListenerSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultBackendSetName != nil {
		in, out := &in.DefaultBackendSetName, &out.DefaultBackendSetName
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkLoadBalancerID != nil {
		in, out := &in.NetworkLoadBalancerID, &out.NetworkLoadBalancerID
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerListenerSpecResource.
func (in *LoadBalancerListenerSpecResource) DeepCopy() *LoadBalancerListenerSpecResource {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerListenerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerListenerStatus) DeepCopyInto(out *LoadBalancerListenerStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerListenerStatus.
func (in *LoadBalancerListenerStatus) DeepCopy() *LoadBalancerListenerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerListenerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerNetworkLoadBalancer) DeepCopyInto(out *LoadBalancerNetworkLoadBalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerNetworkLoadBalancer.
func (in *LoadBalancerNetworkLoadBalancer) DeepCopy() *LoadBalancerNetworkLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerNetworkLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerNetworkLoadBalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerNetworkLoadBalancerList) DeepCopyInto(out *LoadBalancerNetworkLoadBalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LoadBalancerNetworkLoadBalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerNetworkLoadBalancerList.
func (in *LoadBalancerNetworkLoadBalancerList) DeepCopy() *LoadBalancerNetworkLoadBalancerList {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerNetworkLoadBalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LoadBalancerNetworkLoadBalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerNetworkLoadBalancerSpec) DeepCopyInto(out *LoadBalancerNetworkLoadBalancerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LoadBalancerNetworkLoadBalancerSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerNetworkLoadBalancerSpec.
func (in *LoadBalancerNetworkLoadBalancerSpec) DeepCopy() *LoadBalancerNetworkLoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerNetworkLoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerNetworkLoadBalancerSpecIpAddresses) DeepCopyInto(out *LoadBalancerNetworkLoadBalancerSpecIpAddresses) {
	*out = *in
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		*out = new(bool)
		**out = **in
	}
	if in.ReservedIP != nil {
		in, out := &in.ReservedIP, &out.ReservedIP
		*out = new(LoadBalancerNetworkLoadBalancerSpecIpAddressesReservedIP)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerNetworkLoadBalancerSpecIpAddresses.
func (in *LoadBalancerNetworkLoadBalancerSpecIpAddresses) DeepCopy() *LoadBalancerNetworkLoadBalancerSpecIpAddresses {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerNetworkLoadBalancerSpecIpAddresses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerNetworkLoadBalancerSpecIpAddressesReservedIP) DeepCopyInto(out *LoadBalancerNetworkLoadBalancerSpecIpAddressesReservedIP) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerNetworkLoadBalancerSpecIpAddressesReservedIP.
func (in *LoadBalancerNetworkLoadBalancerSpecIpAddressesReservedIP) DeepCopy() *LoadBalancerNetworkLoadBalancerSpecIpAddressesReservedIP {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerNetworkLoadBalancerSpecIpAddressesReservedIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerNetworkLoadBalancerSpecReservedIPS) DeepCopyInto(out *LoadBalancerNetworkLoadBalancerSpecReservedIPS) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerNetworkLoadBalancerSpecReservedIPS.
func (in *LoadBalancerNetworkLoadBalancerSpecReservedIPS) DeepCopy() *LoadBalancerNetworkLoadBalancerSpecReservedIPS {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerNetworkLoadBalancerSpecReservedIPS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerNetworkLoadBalancerSpecResource) DeepCopyInto(out *LoadBalancerNetworkLoadBalancerSpecResource) {
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
	if in.IpAddresses != nil {
		in, out := &in.IpAddresses, &out.IpAddresses
		*out = make([]LoadBalancerNetworkLoadBalancerSpecIpAddresses, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsPreserveSourceDestination != nil {
		in, out := &in.IsPreserveSourceDestination, &out.IsPreserveSourceDestination
		*out = new(bool)
		**out = **in
	}
	if in.IsPrivate != nil {
		in, out := &in.IsPrivate, &out.IsPrivate
		*out = new(bool)
		**out = **in
	}
	if in.LifecycleDetails != nil {
		in, out := &in.LifecycleDetails, &out.LifecycleDetails
		*out = new(string)
		**out = **in
	}
	if in.NetworkSecurityGroupIDS != nil {
		in, out := &in.NetworkSecurityGroupIDS, &out.NetworkSecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReservedIPS != nil {
		in, out := &in.ReservedIPS, &out.ReservedIPS
		*out = make([]LoadBalancerNetworkLoadBalancerSpecReservedIPS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.SystemTags != nil {
		in, out := &in.SystemTags, &out.SystemTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeUpdated != nil {
		in, out := &in.TimeUpdated, &out.TimeUpdated
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerNetworkLoadBalancerSpecResource.
func (in *LoadBalancerNetworkLoadBalancerSpecResource) DeepCopy() *LoadBalancerNetworkLoadBalancerSpecResource {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerNetworkLoadBalancerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerNetworkLoadBalancerStatus) DeepCopyInto(out *LoadBalancerNetworkLoadBalancerStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerNetworkLoadBalancerStatus.
func (in *LoadBalancerNetworkLoadBalancerStatus) DeepCopy() *LoadBalancerNetworkLoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerNetworkLoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}
