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
func (in *ProxySubscriptionRedeemableUser) DeepCopyInto(out *ProxySubscriptionRedeemableUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySubscriptionRedeemableUser.
func (in *ProxySubscriptionRedeemableUser) DeepCopy() *ProxySubscriptionRedeemableUser {
	if in == nil {
		return nil
	}
	out := new(ProxySubscriptionRedeemableUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxySubscriptionRedeemableUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySubscriptionRedeemableUserList) DeepCopyInto(out *ProxySubscriptionRedeemableUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxySubscriptionRedeemableUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySubscriptionRedeemableUserList.
func (in *ProxySubscriptionRedeemableUserList) DeepCopy() *ProxySubscriptionRedeemableUserList {
	if in == nil {
		return nil
	}
	out := new(ProxySubscriptionRedeemableUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProxySubscriptionRedeemableUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySubscriptionRedeemableUserSpec) DeepCopyInto(out *ProxySubscriptionRedeemableUserSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ProxySubscriptionRedeemableUserSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySubscriptionRedeemableUserSpec.
func (in *ProxySubscriptionRedeemableUserSpec) DeepCopy() *ProxySubscriptionRedeemableUserSpec {
	if in == nil {
		return nil
	}
	out := new(ProxySubscriptionRedeemableUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySubscriptionRedeemableUserSpecItems) DeepCopyInto(out *ProxySubscriptionRedeemableUserSpecItems) {
	*out = *in
	if in.EmailID != nil {
		in, out := &in.EmailID, &out.EmailID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySubscriptionRedeemableUserSpecItems.
func (in *ProxySubscriptionRedeemableUserSpecItems) DeepCopy() *ProxySubscriptionRedeemableUserSpecItems {
	if in == nil {
		return nil
	}
	out := new(ProxySubscriptionRedeemableUserSpecItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySubscriptionRedeemableUserSpecResource) DeepCopyInto(out *ProxySubscriptionRedeemableUserSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProxySubscriptionRedeemableUserSpecItems, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubscriptionID != nil {
		in, out := &in.SubscriptionID, &out.SubscriptionID
		*out = new(string)
		**out = **in
	}
	if in.TenancyID != nil {
		in, out := &in.TenancyID, &out.TenancyID
		*out = new(string)
		**out = **in
	}
	if in.UserID != nil {
		in, out := &in.UserID, &out.UserID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySubscriptionRedeemableUserSpecResource.
func (in *ProxySubscriptionRedeemableUserSpecResource) DeepCopy() *ProxySubscriptionRedeemableUserSpecResource {
	if in == nil {
		return nil
	}
	out := new(ProxySubscriptionRedeemableUserSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxySubscriptionRedeemableUserStatus) DeepCopyInto(out *ProxySubscriptionRedeemableUserStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxySubscriptionRedeemableUserStatus.
func (in *ProxySubscriptionRedeemableUserStatus) DeepCopy() *ProxySubscriptionRedeemableUserStatus {
	if in == nil {
		return nil
	}
	out := new(ProxySubscriptionRedeemableUserStatus)
	in.DeepCopyInto(out)
	return out
}
