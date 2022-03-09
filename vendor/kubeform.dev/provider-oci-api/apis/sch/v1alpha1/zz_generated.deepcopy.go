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
func (in *ServiceConnector) DeepCopyInto(out *ServiceConnector) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnector.
func (in *ServiceConnector) DeepCopy() *ServiceConnector {
	if in == nil {
		return nil
	}
	out := new(ServiceConnector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceConnector) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorList) DeepCopyInto(out *ServiceConnectorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceConnector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorList.
func (in *ServiceConnectorList) DeepCopy() *ServiceConnectorList {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceConnectorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorSpec) DeepCopyInto(out *ServiceConnectorSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ServiceConnectorSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorSpec.
func (in *ServiceConnectorSpec) DeepCopy() *ServiceConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorSpecResource) DeepCopyInto(out *ServiceConnectorSpecResource) {
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
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LifecyleDetails != nil {
		in, out := &in.LifecyleDetails, &out.LifecyleDetails
		*out = new(string)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(ServiceConnectorSpecSource)
		(*in).DeepCopyInto(*out)
	}
	if in.State != nil {
		in, out := &in.State, &out.State
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
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(ServiceConnectorSpecTarget)
		(*in).DeepCopyInto(*out)
	}
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]ServiceConnectorSpecTasks, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorSpecResource.
func (in *ServiceConnectorSpecResource) DeepCopy() *ServiceConnectorSpecResource {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorSpecSource) DeepCopyInto(out *ServiceConnectorSpecSource) {
	*out = *in
	if in.Cursor != nil {
		in, out := &in.Cursor, &out.Cursor
		*out = new(ServiceConnectorSpecSourceCursor)
		(*in).DeepCopyInto(*out)
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.LogSources != nil {
		in, out := &in.LogSources, &out.LogSources
		*out = make([]ServiceConnectorSpecSourceLogSources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StreamID != nil {
		in, out := &in.StreamID, &out.StreamID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorSpecSource.
func (in *ServiceConnectorSpecSource) DeepCopy() *ServiceConnectorSpecSource {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorSpecSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorSpecSourceCursor) DeepCopyInto(out *ServiceConnectorSpecSourceCursor) {
	*out = *in
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorSpecSourceCursor.
func (in *ServiceConnectorSpecSourceCursor) DeepCopy() *ServiceConnectorSpecSourceCursor {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorSpecSourceCursor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorSpecSourceLogSources) DeepCopyInto(out *ServiceConnectorSpecSourceLogSources) {
	*out = *in
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.LogGroupID != nil {
		in, out := &in.LogGroupID, &out.LogGroupID
		*out = new(string)
		**out = **in
	}
	if in.LogID != nil {
		in, out := &in.LogID, &out.LogID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorSpecSourceLogSources.
func (in *ServiceConnectorSpecSourceLogSources) DeepCopy() *ServiceConnectorSpecSourceLogSources {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorSpecSourceLogSources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorSpecTarget) DeepCopyInto(out *ServiceConnectorSpecTarget) {
	*out = *in
	if in.BatchRolloverSizeInMbs != nil {
		in, out := &in.BatchRolloverSizeInMbs, &out.BatchRolloverSizeInMbs
		*out = new(int64)
		**out = **in
	}
	if in.BatchRolloverTimeInMs != nil {
		in, out := &in.BatchRolloverTimeInMs, &out.BatchRolloverTimeInMs
		*out = new(int64)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]ServiceConnectorSpecTargetDimensions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableFormattedMessaging != nil {
		in, out := &in.EnableFormattedMessaging, &out.EnableFormattedMessaging
		*out = new(bool)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.LogGroupID != nil {
		in, out := &in.LogGroupID, &out.LogGroupID
		*out = new(string)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(string)
		**out = **in
	}
	if in.MetricNamespace != nil {
		in, out := &in.MetricNamespace, &out.MetricNamespace
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.ObjectNamePrefix != nil {
		in, out := &in.ObjectNamePrefix, &out.ObjectNamePrefix
		*out = new(string)
		**out = **in
	}
	if in.StreamID != nil {
		in, out := &in.StreamID, &out.StreamID
		*out = new(string)
		**out = **in
	}
	if in.TopicID != nil {
		in, out := &in.TopicID, &out.TopicID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorSpecTarget.
func (in *ServiceConnectorSpecTarget) DeepCopy() *ServiceConnectorSpecTarget {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorSpecTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorSpecTargetDimensions) DeepCopyInto(out *ServiceConnectorSpecTargetDimensions) {
	*out = *in
	if in.DimensionValue != nil {
		in, out := &in.DimensionValue, &out.DimensionValue
		*out = new(ServiceConnectorSpecTargetDimensionsDimensionValue)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorSpecTargetDimensions.
func (in *ServiceConnectorSpecTargetDimensions) DeepCopy() *ServiceConnectorSpecTargetDimensions {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorSpecTargetDimensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorSpecTargetDimensionsDimensionValue) DeepCopyInto(out *ServiceConnectorSpecTargetDimensionsDimensionValue) {
	*out = *in
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorSpecTargetDimensionsDimensionValue.
func (in *ServiceConnectorSpecTargetDimensionsDimensionValue) DeepCopy() *ServiceConnectorSpecTargetDimensionsDimensionValue {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorSpecTargetDimensionsDimensionValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorSpecTasks) DeepCopyInto(out *ServiceConnectorSpecTasks) {
	*out = *in
	if in.BatchSizeInKbs != nil {
		in, out := &in.BatchSizeInKbs, &out.BatchSizeInKbs
		*out = new(int64)
		**out = **in
	}
	if in.BatchTimeInSec != nil {
		in, out := &in.BatchTimeInSec, &out.BatchTimeInSec
		*out = new(int64)
		**out = **in
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(string)
		**out = **in
	}
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorSpecTasks.
func (in *ServiceConnectorSpecTasks) DeepCopy() *ServiceConnectorSpecTasks {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorSpecTasks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConnectorStatus) DeepCopyInto(out *ServiceConnectorStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConnectorStatus.
func (in *ServiceConnectorStatus) DeepCopy() *ServiceConnectorStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceConnectorStatus)
	in.DeepCopyInto(out)
	return out
}
