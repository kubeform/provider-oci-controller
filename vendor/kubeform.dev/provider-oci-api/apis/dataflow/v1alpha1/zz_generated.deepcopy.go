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
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ApplicationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpecParameters) DeepCopyInto(out *ApplicationSpecParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpecParameters.
func (in *ApplicationSpecParameters) DeepCopy() *ApplicationSpecParameters {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpecResource) DeepCopyInto(out *ApplicationSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ArchiveURI != nil {
		in, out := &in.ArchiveURI, &out.ArchiveURI
		*out = new(string)
		**out = **in
	}
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClassName != nil {
		in, out := &in.ClassName, &out.ClassName
		*out = new(string)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
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
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.DriverShape != nil {
		in, out := &in.DriverShape, &out.DriverShape
		*out = new(string)
		**out = **in
	}
	if in.Execute != nil {
		in, out := &in.Execute, &out.Execute
		*out = new(string)
		**out = **in
	}
	if in.ExecutorShape != nil {
		in, out := &in.ExecutorShape, &out.ExecutorShape
		*out = new(string)
		**out = **in
	}
	if in.FileURI != nil {
		in, out := &in.FileURI, &out.FileURI
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
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.LogsBucketURI != nil {
		in, out := &in.LogsBucketURI, &out.LogsBucketURI
		*out = new(string)
		**out = **in
	}
	if in.NumExecutors != nil {
		in, out := &in.NumExecutors, &out.NumExecutors
		*out = new(int64)
		**out = **in
	}
	if in.OwnerPrincipalID != nil {
		in, out := &in.OwnerPrincipalID, &out.OwnerPrincipalID
		*out = new(string)
		**out = **in
	}
	if in.OwnerUserName != nil {
		in, out := &in.OwnerUserName, &out.OwnerUserName
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ApplicationSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateEndpointID != nil {
		in, out := &in.PrivateEndpointID, &out.PrivateEndpointID
		*out = new(string)
		**out = **in
	}
	if in.SparkVersion != nil {
		in, out := &in.SparkVersion, &out.SparkVersion
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
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
	if in.WarehouseBucketURI != nil {
		in, out := &in.WarehouseBucketURI, &out.WarehouseBucketURI
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpecResource.
func (in *ApplicationSpecResource) DeepCopy() *ApplicationSpecResource {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeRun) DeepCopyInto(out *InvokeRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeRun.
func (in *InvokeRun) DeepCopy() *InvokeRun {
	if in == nil {
		return nil
	}
	out := new(InvokeRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InvokeRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeRunList) DeepCopyInto(out *InvokeRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InvokeRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeRunList.
func (in *InvokeRunList) DeepCopy() *InvokeRunList {
	if in == nil {
		return nil
	}
	out := new(InvokeRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InvokeRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeRunSpec) DeepCopyInto(out *InvokeRunSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(InvokeRunSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeRunSpec.
func (in *InvokeRunSpec) DeepCopy() *InvokeRunSpec {
	if in == nil {
		return nil
	}
	out := new(InvokeRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeRunSpecParameters) DeepCopyInto(out *InvokeRunSpecParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeRunSpecParameters.
func (in *InvokeRunSpecParameters) DeepCopy() *InvokeRunSpecParameters {
	if in == nil {
		return nil
	}
	out := new(InvokeRunSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeRunSpecResource) DeepCopyInto(out *InvokeRunSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ApplicationID != nil {
		in, out := &in.ApplicationID, &out.ApplicationID
		*out = new(string)
		**out = **in
	}
	if in.ArchiveURI != nil {
		in, out := &in.ArchiveURI, &out.ArchiveURI
		*out = new(string)
		**out = **in
	}
	if in.Arguments != nil {
		in, out := &in.Arguments, &out.Arguments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Asynchronous != nil {
		in, out := &in.Asynchronous, &out.Asynchronous
		*out = new(bool)
		**out = **in
	}
	if in.ClassName != nil {
		in, out := &in.ClassName, &out.ClassName
		*out = new(string)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DataReadInBytes != nil {
		in, out := &in.DataReadInBytes, &out.DataReadInBytes
		*out = new(string)
		**out = **in
	}
	if in.DataWrittenInBytes != nil {
		in, out := &in.DataWrittenInBytes, &out.DataWrittenInBytes
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
	if in.DriverShape != nil {
		in, out := &in.DriverShape, &out.DriverShape
		*out = new(string)
		**out = **in
	}
	if in.Execute != nil {
		in, out := &in.Execute, &out.Execute
		*out = new(string)
		**out = **in
	}
	if in.ExecutorShape != nil {
		in, out := &in.ExecutorShape, &out.ExecutorShape
		*out = new(string)
		**out = **in
	}
	if in.FileURI != nil {
		in, out := &in.FileURI, &out.FileURI
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
	if in.Language != nil {
		in, out := &in.Language, &out.Language
		*out = new(string)
		**out = **in
	}
	if in.LifecycleDetails != nil {
		in, out := &in.LifecycleDetails, &out.LifecycleDetails
		*out = new(string)
		**out = **in
	}
	if in.LogsBucketURI != nil {
		in, out := &in.LogsBucketURI, &out.LogsBucketURI
		*out = new(string)
		**out = **in
	}
	if in.NumExecutors != nil {
		in, out := &in.NumExecutors, &out.NumExecutors
		*out = new(int64)
		**out = **in
	}
	if in.OpcRequestID != nil {
		in, out := &in.OpcRequestID, &out.OpcRequestID
		*out = new(string)
		**out = **in
	}
	if in.OwnerPrincipalID != nil {
		in, out := &in.OwnerPrincipalID, &out.OwnerPrincipalID
		*out = new(string)
		**out = **in
	}
	if in.OwnerUserName != nil {
		in, out := &in.OwnerUserName, &out.OwnerUserName
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]InvokeRunSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateEndpointDNSZones != nil {
		in, out := &in.PrivateEndpointDNSZones, &out.PrivateEndpointDNSZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrivateEndpointID != nil {
		in, out := &in.PrivateEndpointID, &out.PrivateEndpointID
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpointMaxHostCount != nil {
		in, out := &in.PrivateEndpointMaxHostCount, &out.PrivateEndpointMaxHostCount
		*out = new(int64)
		**out = **in
	}
	if in.PrivateEndpointNsgIDS != nil {
		in, out := &in.PrivateEndpointNsgIDS, &out.PrivateEndpointNsgIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrivateEndpointSubnetID != nil {
		in, out := &in.PrivateEndpointSubnetID, &out.PrivateEndpointSubnetID
		*out = new(string)
		**out = **in
	}
	if in.RunDurationInMilliseconds != nil {
		in, out := &in.RunDurationInMilliseconds, &out.RunDurationInMilliseconds
		*out = new(string)
		**out = **in
	}
	if in.SparkVersion != nil {
		in, out := &in.SparkVersion, &out.SparkVersion
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
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
	if in.TotalOcpu != nil {
		in, out := &in.TotalOcpu, &out.TotalOcpu
		*out = new(int64)
		**out = **in
	}
	if in.WarehouseBucketURI != nil {
		in, out := &in.WarehouseBucketURI, &out.WarehouseBucketURI
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeRunSpecResource.
func (in *InvokeRunSpecResource) DeepCopy() *InvokeRunSpecResource {
	if in == nil {
		return nil
	}
	out := new(InvokeRunSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InvokeRunStatus) DeepCopyInto(out *InvokeRunStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InvokeRunStatus.
func (in *InvokeRunStatus) DeepCopy() *InvokeRunStatus {
	if in == nil {
		return nil
	}
	out := new(InvokeRunStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpoint) DeepCopyInto(out *PrivateEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpoint.
func (in *PrivateEndpoint) DeepCopy() *PrivateEndpoint {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointList) DeepCopyInto(out *PrivateEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrivateEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointList.
func (in *PrivateEndpointList) DeepCopy() *PrivateEndpointList {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointSpec) DeepCopyInto(out *PrivateEndpointSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PrivateEndpointSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointSpec.
func (in *PrivateEndpointSpec) DeepCopy() *PrivateEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointSpecResource) DeepCopyInto(out *PrivateEndpointSpecResource) {
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
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.DnsZones != nil {
		in, out := &in.DnsZones, &out.DnsZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LifecycleDetails != nil {
		in, out := &in.LifecycleDetails, &out.LifecycleDetails
		*out = new(string)
		**out = **in
	}
	if in.MaxHostCount != nil {
		in, out := &in.MaxHostCount, &out.MaxHostCount
		*out = new(int64)
		**out = **in
	}
	if in.NsgIDS != nil {
		in, out := &in.NsgIDS, &out.NsgIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OwnerPrincipalID != nil {
		in, out := &in.OwnerPrincipalID, &out.OwnerPrincipalID
		*out = new(string)
		**out = **in
	}
	if in.OwnerUserName != nil {
		in, out := &in.OwnerUserName, &out.OwnerUserName
		*out = new(string)
		**out = **in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointSpecResource.
func (in *PrivateEndpointSpecResource) DeepCopy() *PrivateEndpointSpecResource {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateEndpointStatus) DeepCopyInto(out *PrivateEndpointStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateEndpointStatus.
func (in *PrivateEndpointStatus) DeepCopy() *PrivateEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(PrivateEndpointStatus)
	in.DeepCopyInto(out)
	return out
}
