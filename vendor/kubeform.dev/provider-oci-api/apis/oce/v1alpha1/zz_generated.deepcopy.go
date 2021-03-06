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
func (in *OceInstance) DeepCopyInto(out *OceInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OceInstance.
func (in *OceInstance) DeepCopy() *OceInstance {
	if in == nil {
		return nil
	}
	out := new(OceInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OceInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OceInstanceList) DeepCopyInto(out *OceInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OceInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OceInstanceList.
func (in *OceInstanceList) DeepCopy() *OceInstanceList {
	if in == nil {
		return nil
	}
	out := new(OceInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OceInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OceInstanceSpec) DeepCopyInto(out *OceInstanceSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(OceInstanceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OceInstanceSpec.
func (in *OceInstanceSpec) DeepCopy() *OceInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(OceInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OceInstanceSpecResource) DeepCopyInto(out *OceInstanceSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AdminEmail != nil {
		in, out := &in.AdminEmail, &out.AdminEmail
		*out = new(string)
		**out = **in
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
	if in.Description != nil {
		in, out := &in.Description, &out.Description
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
	if in.Guid != nil {
		in, out := &in.Guid, &out.Guid
		*out = new(string)
		**out = **in
	}
	if in.IdcsAccessToken != nil {
		in, out := &in.IdcsAccessToken, &out.IdcsAccessToken
		*out = new(string)
		**out = **in
	}
	if in.IdcsTenancy != nil {
		in, out := &in.IdcsTenancy, &out.IdcsTenancy
		*out = new(string)
		**out = **in
	}
	if in.InstanceAccessType != nil {
		in, out := &in.InstanceAccessType, &out.InstanceAccessType
		*out = new(string)
		**out = **in
	}
	if in.InstanceLicenseType != nil {
		in, out := &in.InstanceLicenseType, &out.InstanceLicenseType
		*out = new(string)
		**out = **in
	}
	if in.InstanceUsageType != nil {
		in, out := &in.InstanceUsageType, &out.InstanceUsageType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObjectStorageNamespace != nil {
		in, out := &in.ObjectStorageNamespace, &out.ObjectStorageNamespace
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateMessage != nil {
		in, out := &in.StateMessage, &out.StateMessage
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
	if in.TenancyID != nil {
		in, out := &in.TenancyID, &out.TenancyID
		*out = new(string)
		**out = **in
	}
	if in.TenancyName != nil {
		in, out := &in.TenancyName, &out.TenancyName
		*out = new(string)
		**out = **in
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
	if in.UpgradeSchedule != nil {
		in, out := &in.UpgradeSchedule, &out.UpgradeSchedule
		*out = new(string)
		**out = **in
	}
	if in.WafPrimaryDomain != nil {
		in, out := &in.WafPrimaryDomain, &out.WafPrimaryDomain
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OceInstanceSpecResource.
func (in *OceInstanceSpecResource) DeepCopy() *OceInstanceSpecResource {
	if in == nil {
		return nil
	}
	out := new(OceInstanceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OceInstanceStatus) DeepCopyInto(out *OceInstanceStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OceInstanceStatus.
func (in *OceInstanceStatus) DeepCopy() *OceInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(OceInstanceStatus)
	in.DeepCopyInto(out)
	return out
}
