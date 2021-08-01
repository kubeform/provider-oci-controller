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
func (in *Model) DeepCopyInto(out *Model) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Model.
func (in *Model) DeepCopy() *Model {
	if in == nil {
		return nil
	}
	out := new(Model)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Model) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeployment) DeepCopyInto(out *ModelDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeployment.
func (in *ModelDeployment) DeepCopy() *ModelDeployment {
	if in == nil {
		return nil
	}
	out := new(ModelDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentList) DeepCopyInto(out *ModelDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentList.
func (in *ModelDeploymentList) DeepCopy() *ModelDeploymentList {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpec) DeepCopyInto(out *ModelDeploymentSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ModelDeploymentSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpec.
func (in *ModelDeploymentSpec) DeepCopy() *ModelDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpecCategoryLogDetails) DeepCopyInto(out *ModelDeploymentSpecCategoryLogDetails) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = new(ModelDeploymentSpecCategoryLogDetailsAccess)
		(*in).DeepCopyInto(*out)
	}
	if in.Predict != nil {
		in, out := &in.Predict, &out.Predict
		*out = new(ModelDeploymentSpecCategoryLogDetailsPredict)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpecCategoryLogDetails.
func (in *ModelDeploymentSpecCategoryLogDetails) DeepCopy() *ModelDeploymentSpecCategoryLogDetails {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpecCategoryLogDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpecCategoryLogDetailsAccess) DeepCopyInto(out *ModelDeploymentSpecCategoryLogDetailsAccess) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpecCategoryLogDetailsAccess.
func (in *ModelDeploymentSpecCategoryLogDetailsAccess) DeepCopy() *ModelDeploymentSpecCategoryLogDetailsAccess {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpecCategoryLogDetailsAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpecCategoryLogDetailsPredict) DeepCopyInto(out *ModelDeploymentSpecCategoryLogDetailsPredict) {
	*out = *in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpecCategoryLogDetailsPredict.
func (in *ModelDeploymentSpecCategoryLogDetailsPredict) DeepCopy() *ModelDeploymentSpecCategoryLogDetailsPredict {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpecCategoryLogDetailsPredict)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpecModelDeploymentConfigurationDetails) DeepCopyInto(out *ModelDeploymentSpecModelDeploymentConfigurationDetails) {
	*out = *in
	if in.DeploymentType != nil {
		in, out := &in.DeploymentType, &out.DeploymentType
		*out = new(string)
		**out = **in
	}
	if in.ModelConfigurationDetails != nil {
		in, out := &in.ModelConfigurationDetails, &out.ModelConfigurationDetails
		*out = new(ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetails)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpecModelDeploymentConfigurationDetails.
func (in *ModelDeploymentSpecModelDeploymentConfigurationDetails) DeepCopy() *ModelDeploymentSpecModelDeploymentConfigurationDetails {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpecModelDeploymentConfigurationDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetails) DeepCopyInto(out *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetails) {
	*out = *in
	if in.BandwidthMbps != nil {
		in, out := &in.BandwidthMbps, &out.BandwidthMbps
		*out = new(int64)
		**out = **in
	}
	if in.InstanceConfiguration != nil {
		in, out := &in.InstanceConfiguration, &out.InstanceConfiguration
		*out = new(ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ModelID != nil {
		in, out := &in.ModelID, &out.ModelID
		*out = new(string)
		**out = **in
	}
	if in.ScalingPolicy != nil {
		in, out := &in.ScalingPolicy, &out.ScalingPolicy
		*out = new(ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetails.
func (in *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetails) DeepCopy() *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetails {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfiguration) DeepCopyInto(out *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfiguration) {
	*out = *in
	if in.InstanceShapeName != nil {
		in, out := &in.InstanceShapeName, &out.InstanceShapeName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfiguration.
func (in *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfiguration) DeepCopy() *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfiguration {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsInstanceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicy) DeepCopyInto(out *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicy) {
	*out = *in
	if in.InstanceCount != nil {
		in, out := &in.InstanceCount, &out.InstanceCount
		*out = new(int64)
		**out = **in
	}
	if in.PolicyType != nil {
		in, out := &in.PolicyType, &out.PolicyType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicy.
func (in *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicy) DeepCopy() *ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpecModelDeploymentConfigurationDetailsModelConfigurationDetailsScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentSpecResource) DeepCopyInto(out *ModelDeploymentSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CategoryLogDetails != nil {
		in, out := &in.CategoryLogDetails, &out.CategoryLogDetails
		*out = new(ModelDeploymentSpecCategoryLogDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
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
	if in.LifecycleDetails != nil {
		in, out := &in.LifecycleDetails, &out.LifecycleDetails
		*out = new(string)
		**out = **in
	}
	if in.ModelDeploymentConfigurationDetails != nil {
		in, out := &in.ModelDeploymentConfigurationDetails, &out.ModelDeploymentConfigurationDetails
		*out = new(ModelDeploymentSpecModelDeploymentConfigurationDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.ModelDeploymentURL != nil {
		in, out := &in.ModelDeploymentURL, &out.ModelDeploymentURL
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentSpecResource.
func (in *ModelDeploymentSpecResource) DeepCopy() *ModelDeploymentSpecResource {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelDeploymentStatus) DeepCopyInto(out *ModelDeploymentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelDeploymentStatus.
func (in *ModelDeploymentStatus) DeepCopy() *ModelDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(ModelDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelList) DeepCopyInto(out *ModelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Model, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelList.
func (in *ModelList) DeepCopy() *ModelList {
	if in == nil {
		return nil
	}
	out := new(ModelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelProvenance) DeepCopyInto(out *ModelProvenance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelProvenance.
func (in *ModelProvenance) DeepCopy() *ModelProvenance {
	if in == nil {
		return nil
	}
	out := new(ModelProvenance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelProvenance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelProvenanceList) DeepCopyInto(out *ModelProvenanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ModelProvenance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelProvenanceList.
func (in *ModelProvenanceList) DeepCopy() *ModelProvenanceList {
	if in == nil {
		return nil
	}
	out := new(ModelProvenanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModelProvenanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelProvenanceSpec) DeepCopyInto(out *ModelProvenanceSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ModelProvenanceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelProvenanceSpec.
func (in *ModelProvenanceSpec) DeepCopy() *ModelProvenanceSpec {
	if in == nil {
		return nil
	}
	out := new(ModelProvenanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelProvenanceSpecResource) DeepCopyInto(out *ModelProvenanceSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.GitBranch != nil {
		in, out := &in.GitBranch, &out.GitBranch
		*out = new(string)
		**out = **in
	}
	if in.GitCommit != nil {
		in, out := &in.GitCommit, &out.GitCommit
		*out = new(string)
		**out = **in
	}
	if in.ModelID != nil {
		in, out := &in.ModelID, &out.ModelID
		*out = new(string)
		**out = **in
	}
	if in.RepositoryURL != nil {
		in, out := &in.RepositoryURL, &out.RepositoryURL
		*out = new(string)
		**out = **in
	}
	if in.ScriptDir != nil {
		in, out := &in.ScriptDir, &out.ScriptDir
		*out = new(string)
		**out = **in
	}
	if in.TrainingScript != nil {
		in, out := &in.TrainingScript, &out.TrainingScript
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelProvenanceSpecResource.
func (in *ModelProvenanceSpecResource) DeepCopy() *ModelProvenanceSpecResource {
	if in == nil {
		return nil
	}
	out := new(ModelProvenanceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelProvenanceStatus) DeepCopyInto(out *ModelProvenanceStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelProvenanceStatus.
func (in *ModelProvenanceStatus) DeepCopy() *ModelProvenanceStatus {
	if in == nil {
		return nil
	}
	out := new(ModelProvenanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSpec) DeepCopyInto(out *ModelSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ModelSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSpec.
func (in *ModelSpec) DeepCopy() *ModelSpec {
	if in == nil {
		return nil
	}
	out := new(ModelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelSpecResource) DeepCopyInto(out *ModelSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ArtifactContentDisposition != nil {
		in, out := &in.ArtifactContentDisposition, &out.ArtifactContentDisposition
		*out = new(string)
		**out = **in
	}
	if in.ArtifactContentLength != nil {
		in, out := &in.ArtifactContentLength, &out.ArtifactContentLength
		*out = new(string)
		**out = **in
	}
	if in.ArtifactContentMd5 != nil {
		in, out := &in.ArtifactContentMd5, &out.ArtifactContentMd5
		*out = new(string)
		**out = **in
	}
	if in.ArtifactLastModified != nil {
		in, out := &in.ArtifactLastModified, &out.ArtifactLastModified
		*out = new(string)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
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
	if in.EmptyModel != nil {
		in, out := &in.EmptyModel, &out.EmptyModel
		*out = new(bool)
		**out = **in
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ModelArtifact != nil {
		in, out := &in.ModelArtifact, &out.ModelArtifact
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelSpecResource.
func (in *ModelSpecResource) DeepCopy() *ModelSpecResource {
	if in == nil {
		return nil
	}
	out := new(ModelSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelStatus) DeepCopyInto(out *ModelStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelStatus.
func (in *ModelStatus) DeepCopy() *ModelStatus {
	if in == nil {
		return nil
	}
	out := new(ModelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookSession) DeepCopyInto(out *NotebookSession) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookSession.
func (in *NotebookSession) DeepCopy() *NotebookSession {
	if in == nil {
		return nil
	}
	out := new(NotebookSession)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotebookSession) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookSessionList) DeepCopyInto(out *NotebookSessionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NotebookSession, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookSessionList.
func (in *NotebookSessionList) DeepCopy() *NotebookSessionList {
	if in == nil {
		return nil
	}
	out := new(NotebookSessionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NotebookSessionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookSessionSpec) DeepCopyInto(out *NotebookSessionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(NotebookSessionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookSessionSpec.
func (in *NotebookSessionSpec) DeepCopy() *NotebookSessionSpec {
	if in == nil {
		return nil
	}
	out := new(NotebookSessionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookSessionSpecNotebookSessionConfigurationDetails) DeepCopyInto(out *NotebookSessionSpecNotebookSessionConfigurationDetails) {
	*out = *in
	if in.BlockStorageSizeInGbs != nil {
		in, out := &in.BlockStorageSizeInGbs, &out.BlockStorageSizeInGbs
		*out = new(int64)
		**out = **in
	}
	if in.NotebookSessionShapeConfigDetails != nil {
		in, out := &in.NotebookSessionShapeConfigDetails, &out.NotebookSessionShapeConfigDetails
		*out = new(NotebookSessionSpecNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookSessionSpecNotebookSessionConfigurationDetails.
func (in *NotebookSessionSpecNotebookSessionConfigurationDetails) DeepCopy() *NotebookSessionSpecNotebookSessionConfigurationDetails {
	if in == nil {
		return nil
	}
	out := new(NotebookSessionSpecNotebookSessionConfigurationDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookSessionSpecNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetails) DeepCopyInto(out *NotebookSessionSpecNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetails) {
	*out = *in
	if in.MemoryInGbs != nil {
		in, out := &in.MemoryInGbs, &out.MemoryInGbs
		*out = new(float64)
		**out = **in
	}
	if in.Ocpus != nil {
		in, out := &in.Ocpus, &out.Ocpus
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookSessionSpecNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetails.
func (in *NotebookSessionSpecNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetails) DeepCopy() *NotebookSessionSpecNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetails {
	if in == nil {
		return nil
	}
	out := new(NotebookSessionSpecNotebookSessionConfigurationDetailsNotebookSessionShapeConfigDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookSessionSpecResource) DeepCopyInto(out *NotebookSessionSpecResource) {
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
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
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
	if in.LifecycleDetails != nil {
		in, out := &in.LifecycleDetails, &out.LifecycleDetails
		*out = new(string)
		**out = **in
	}
	if in.NotebookSessionConfigurationDetails != nil {
		in, out := &in.NotebookSessionConfigurationDetails, &out.NotebookSessionConfigurationDetails
		*out = new(NotebookSessionSpecNotebookSessionConfigurationDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.NotebookSessionURL != nil {
		in, out := &in.NotebookSessionURL, &out.NotebookSessionURL
		*out = new(string)
		**out = **in
	}
	if in.ProjectID != nil {
		in, out := &in.ProjectID, &out.ProjectID
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookSessionSpecResource.
func (in *NotebookSessionSpecResource) DeepCopy() *NotebookSessionSpecResource {
	if in == nil {
		return nil
	}
	out := new(NotebookSessionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotebookSessionStatus) DeepCopyInto(out *NotebookSessionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotebookSessionStatus.
func (in *NotebookSessionStatus) DeepCopy() *NotebookSessionStatus {
	if in == nil {
		return nil
	}
	out := new(NotebookSessionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ProjectSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpecResource) DeepCopyInto(out *ProjectSpecResource) {
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
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpecResource.
func (in *ProjectSpecResource) DeepCopy() *ProjectSpecResource {
	if in == nil {
		return nil
	}
	out := new(ProjectSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}