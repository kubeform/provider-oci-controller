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
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ClusterSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecEndpointConfig) DeepCopyInto(out *ClusterSpecEndpointConfig) {
	*out = *in
	if in.IsPublicIPEnabled != nil {
		in, out := &in.IsPublicIPEnabled, &out.IsPublicIPEnabled
		*out = new(bool)
		**out = **in
	}
	if in.NsgIDS != nil {
		in, out := &in.NsgIDS, &out.NsgIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecEndpointConfig.
func (in *ClusterSpecEndpointConfig) DeepCopy() *ClusterSpecEndpointConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecEndpointConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecEndpoints) DeepCopyInto(out *ClusterSpecEndpoints) {
	*out = *in
	if in.Kubernetes != nil {
		in, out := &in.Kubernetes, &out.Kubernetes
		*out = new(string)
		**out = **in
	}
	if in.PrivateEndpoint != nil {
		in, out := &in.PrivateEndpoint, &out.PrivateEndpoint
		*out = new(string)
		**out = **in
	}
	if in.PublicEndpoint != nil {
		in, out := &in.PublicEndpoint, &out.PublicEndpoint
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecEndpoints.
func (in *ClusterSpecEndpoints) DeepCopy() *ClusterSpecEndpoints {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecEndpoints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecImagePolicyConfig) DeepCopyInto(out *ClusterSpecImagePolicyConfig) {
	*out = *in
	if in.IsPolicyEnabled != nil {
		in, out := &in.IsPolicyEnabled, &out.IsPolicyEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KeyDetails != nil {
		in, out := &in.KeyDetails, &out.KeyDetails
		*out = make([]ClusterSpecImagePolicyConfigKeyDetails, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecImagePolicyConfig.
func (in *ClusterSpecImagePolicyConfig) DeepCopy() *ClusterSpecImagePolicyConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecImagePolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecImagePolicyConfigKeyDetails) DeepCopyInto(out *ClusterSpecImagePolicyConfigKeyDetails) {
	*out = *in
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecImagePolicyConfigKeyDetails.
func (in *ClusterSpecImagePolicyConfigKeyDetails) DeepCopy() *ClusterSpecImagePolicyConfigKeyDetails {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecImagePolicyConfigKeyDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecMetadata) DeepCopyInto(out *ClusterSpecMetadata) {
	*out = *in
	if in.CreatedByUserID != nil {
		in, out := &in.CreatedByUserID, &out.CreatedByUserID
		*out = new(string)
		**out = **in
	}
	if in.CreatedByWorkRequestID != nil {
		in, out := &in.CreatedByWorkRequestID, &out.CreatedByWorkRequestID
		*out = new(string)
		**out = **in
	}
	if in.DeletedByUserID != nil {
		in, out := &in.DeletedByUserID, &out.DeletedByUserID
		*out = new(string)
		**out = **in
	}
	if in.DeletedByWorkRequestID != nil {
		in, out := &in.DeletedByWorkRequestID, &out.DeletedByWorkRequestID
		*out = new(string)
		**out = **in
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	if in.TimeDeleted != nil {
		in, out := &in.TimeDeleted, &out.TimeDeleted
		*out = new(string)
		**out = **in
	}
	if in.TimeUpdated != nil {
		in, out := &in.TimeUpdated, &out.TimeUpdated
		*out = new(string)
		**out = **in
	}
	if in.UpdatedByUserID != nil {
		in, out := &in.UpdatedByUserID, &out.UpdatedByUserID
		*out = new(string)
		**out = **in
	}
	if in.UpdatedByWorkRequestID != nil {
		in, out := &in.UpdatedByWorkRequestID, &out.UpdatedByWorkRequestID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecMetadata.
func (in *ClusterSpecMetadata) DeepCopy() *ClusterSpecMetadata {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecOptions) DeepCopyInto(out *ClusterSpecOptions) {
	*out = *in
	if in.AddOns != nil {
		in, out := &in.AddOns, &out.AddOns
		*out = new(ClusterSpecOptionsAddOns)
		(*in).DeepCopyInto(*out)
	}
	if in.AdmissionControllerOptions != nil {
		in, out := &in.AdmissionControllerOptions, &out.AdmissionControllerOptions
		*out = new(ClusterSpecOptionsAdmissionControllerOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.KubernetesNetworkConfig != nil {
		in, out := &in.KubernetesNetworkConfig, &out.KubernetesNetworkConfig
		*out = new(ClusterSpecOptionsKubernetesNetworkConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistentVolumeConfig != nil {
		in, out := &in.PersistentVolumeConfig, &out.PersistentVolumeConfig
		*out = new(ClusterSpecOptionsPersistentVolumeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceLbConfig != nil {
		in, out := &in.ServiceLbConfig, &out.ServiceLbConfig
		*out = new(ClusterSpecOptionsServiceLbConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceLbSubnetIDS != nil {
		in, out := &in.ServiceLbSubnetIDS, &out.ServiceLbSubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecOptions.
func (in *ClusterSpecOptions) DeepCopy() *ClusterSpecOptions {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecOptionsAddOns) DeepCopyInto(out *ClusterSpecOptionsAddOns) {
	*out = *in
	if in.IsKubernetesDashboardEnabled != nil {
		in, out := &in.IsKubernetesDashboardEnabled, &out.IsKubernetesDashboardEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsTillerEnabled != nil {
		in, out := &in.IsTillerEnabled, &out.IsTillerEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecOptionsAddOns.
func (in *ClusterSpecOptionsAddOns) DeepCopy() *ClusterSpecOptionsAddOns {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecOptionsAddOns)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecOptionsAdmissionControllerOptions) DeepCopyInto(out *ClusterSpecOptionsAdmissionControllerOptions) {
	*out = *in
	if in.IsPodSecurityPolicyEnabled != nil {
		in, out := &in.IsPodSecurityPolicyEnabled, &out.IsPodSecurityPolicyEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecOptionsAdmissionControllerOptions.
func (in *ClusterSpecOptionsAdmissionControllerOptions) DeepCopy() *ClusterSpecOptionsAdmissionControllerOptions {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecOptionsAdmissionControllerOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecOptionsKubernetesNetworkConfig) DeepCopyInto(out *ClusterSpecOptionsKubernetesNetworkConfig) {
	*out = *in
	if in.PodsCIDR != nil {
		in, out := &in.PodsCIDR, &out.PodsCIDR
		*out = new(string)
		**out = **in
	}
	if in.ServicesCIDR != nil {
		in, out := &in.ServicesCIDR, &out.ServicesCIDR
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecOptionsKubernetesNetworkConfig.
func (in *ClusterSpecOptionsKubernetesNetworkConfig) DeepCopy() *ClusterSpecOptionsKubernetesNetworkConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecOptionsKubernetesNetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecOptionsPersistentVolumeConfig) DeepCopyInto(out *ClusterSpecOptionsPersistentVolumeConfig) {
	*out = *in
	if in.DefinedTags != nil {
		in, out := &in.DefinedTags, &out.DefinedTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecOptionsPersistentVolumeConfig.
func (in *ClusterSpecOptionsPersistentVolumeConfig) DeepCopy() *ClusterSpecOptionsPersistentVolumeConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecOptionsPersistentVolumeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecOptionsServiceLbConfig) DeepCopyInto(out *ClusterSpecOptionsServiceLbConfig) {
	*out = *in
	if in.DefinedTags != nil {
		in, out := &in.DefinedTags, &out.DefinedTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecOptionsServiceLbConfig.
func (in *ClusterSpecOptionsServiceLbConfig) DeepCopy() *ClusterSpecOptionsServiceLbConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecOptionsServiceLbConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecResource) DeepCopyInto(out *ClusterSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AvailableKubernetesUpgrades != nil {
		in, out := &in.AvailableKubernetesUpgrades, &out.AvailableKubernetesUpgrades
		*out = make([]string, len(*in))
		copy(*out, *in)
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
	if in.EndpointConfig != nil {
		in, out := &in.EndpointConfig, &out.EndpointConfig
		*out = new(ClusterSpecEndpointConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = new(ClusterSpecEndpoints)
		(*in).DeepCopyInto(*out)
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ImagePolicyConfig != nil {
		in, out := &in.ImagePolicyConfig, &out.ImagePolicyConfig
		*out = new(ClusterSpecImagePolicyConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.LifecycleDetails != nil {
		in, out := &in.LifecycleDetails, &out.LifecycleDetails
		*out = new(string)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(ClusterSpecMetadata)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(ClusterSpecOptions)
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
	if in.VcnID != nil {
		in, out := &in.VcnID, &out.VcnID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecResource.
func (in *ClusterSpecResource) DeepCopy() *ClusterSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePool) DeepCopyInto(out *NodePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePool.
func (in *NodePool) DeepCopy() *NodePool {
	if in == nil {
		return nil
	}
	out := new(NodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolList) DeepCopyInto(out *NodePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolList.
func (in *NodePoolList) DeepCopy() *NodePoolList {
	if in == nil {
		return nil
	}
	out := new(NodePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpec) DeepCopyInto(out *NodePoolSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(NodePoolSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpec.
func (in *NodePoolSpec) DeepCopy() *NodePoolSpec {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpecInitialNodeLabels) DeepCopyInto(out *NodePoolSpecInitialNodeLabels) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpecInitialNodeLabels.
func (in *NodePoolSpecInitialNodeLabels) DeepCopy() *NodePoolSpecInitialNodeLabels {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpecInitialNodeLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpecNodeConfigDetails) DeepCopyInto(out *NodePoolSpecNodeConfigDetails) {
	*out = *in
	if in.DefinedTags != nil {
		in, out := &in.DefinedTags, &out.DefinedTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.IsPvEncryptionInTransitEnabled != nil {
		in, out := &in.IsPvEncryptionInTransitEnabled, &out.IsPvEncryptionInTransitEnabled
		*out = new(bool)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.NsgIDS != nil {
		in, out := &in.NsgIDS, &out.NsgIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PlacementConfigs != nil {
		in, out := &in.PlacementConfigs, &out.PlacementConfigs
		*out = make([]NodePoolSpecNodeConfigDetailsPlacementConfigs, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpecNodeConfigDetails.
func (in *NodePoolSpecNodeConfigDetails) DeepCopy() *NodePoolSpecNodeConfigDetails {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpecNodeConfigDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpecNodeConfigDetailsPlacementConfigs) DeepCopyInto(out *NodePoolSpecNodeConfigDetailsPlacementConfigs) {
	*out = *in
	if in.AvailabilityDomain != nil {
		in, out := &in.AvailabilityDomain, &out.AvailabilityDomain
		*out = new(string)
		**out = **in
	}
	if in.CapacityReservationID != nil {
		in, out := &in.CapacityReservationID, &out.CapacityReservationID
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpecNodeConfigDetailsPlacementConfigs.
func (in *NodePoolSpecNodeConfigDetailsPlacementConfigs) DeepCopy() *NodePoolSpecNodeConfigDetailsPlacementConfigs {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpecNodeConfigDetailsPlacementConfigs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpecNodeShapeConfig) DeepCopyInto(out *NodePoolSpecNodeShapeConfig) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpecNodeShapeConfig.
func (in *NodePoolSpecNodeShapeConfig) DeepCopy() *NodePoolSpecNodeShapeConfig {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpecNodeShapeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpecNodeSource) DeepCopyInto(out *NodePoolSpecNodeSource) {
	*out = *in
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.SourceName != nil {
		in, out := &in.SourceName, &out.SourceName
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpecNodeSource.
func (in *NodePoolSpecNodeSource) DeepCopy() *NodePoolSpecNodeSource {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpecNodeSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpecNodeSourceDetails) DeepCopyInto(out *NodePoolSpecNodeSourceDetails) {
	*out = *in
	if in.BootVolumeSizeInGbs != nil {
		in, out := &in.BootVolumeSizeInGbs, &out.BootVolumeSizeInGbs
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.SourceType != nil {
		in, out := &in.SourceType, &out.SourceType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpecNodeSourceDetails.
func (in *NodePoolSpecNodeSourceDetails) DeepCopy() *NodePoolSpecNodeSourceDetails {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpecNodeSourceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpecNodes) DeepCopyInto(out *NodePoolSpecNodes) {
	*out = *in
	if in.AvailabilityDomain != nil {
		in, out := &in.AvailabilityDomain, &out.AvailabilityDomain
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
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(NodePoolSpecNodesError)
		(*in).DeepCopyInto(*out)
	}
	if in.FaultDomain != nil {
		in, out := &in.FaultDomain, &out.FaultDomain
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
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.LifecycleDetails != nil {
		in, out := &in.LifecycleDetails, &out.LifecycleDetails
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodePoolID != nil {
		in, out := &in.NodePoolID, &out.NodePoolID
		*out = new(string)
		**out = **in
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(string)
		**out = **in
	}
	if in.PublicIP != nil {
		in, out := &in.PublicIP, &out.PublicIP
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
	if in.SystemTags != nil {
		in, out := &in.SystemTags, &out.SystemTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpecNodes.
func (in *NodePoolSpecNodes) DeepCopy() *NodePoolSpecNodes {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpecNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpecNodesError) DeepCopyInto(out *NodePoolSpecNodesError) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpecNodesError.
func (in *NodePoolSpecNodesError) DeepCopy() *NodePoolSpecNodesError {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpecNodesError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpecResource) DeepCopyInto(out *NodePoolSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
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
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.InitialNodeLabels != nil {
		in, out := &in.InitialNodeLabels, &out.InitialNodeLabels
		*out = make([]NodePoolSpecInitialNodeLabels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeConfigDetails != nil {
		in, out := &in.NodeConfigDetails, &out.NodeConfigDetails
		*out = new(NodePoolSpecNodeConfigDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeImageID != nil {
		in, out := &in.NodeImageID, &out.NodeImageID
		*out = new(string)
		**out = **in
	}
	if in.NodeImageName != nil {
		in, out := &in.NodeImageName, &out.NodeImageName
		*out = new(string)
		**out = **in
	}
	if in.NodeMetadata != nil {
		in, out := &in.NodeMetadata, &out.NodeMetadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeShape != nil {
		in, out := &in.NodeShape, &out.NodeShape
		*out = new(string)
		**out = **in
	}
	if in.NodeShapeConfig != nil {
		in, out := &in.NodeShapeConfig, &out.NodeShapeConfig
		*out = new(NodePoolSpecNodeShapeConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSource != nil {
		in, out := &in.NodeSource, &out.NodeSource
		*out = new(NodePoolSpecNodeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSourceDetails != nil {
		in, out := &in.NodeSourceDetails, &out.NodeSourceDetails
		*out = new(NodePoolSpecNodeSourceDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]NodePoolSpecNodes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.QuantityPerSubnet != nil {
		in, out := &in.QuantityPerSubnet, &out.QuantityPerSubnet
		*out = new(int64)
		**out = **in
	}
	if in.SshPublicKey != nil {
		in, out := &in.SshPublicKey, &out.SshPublicKey
		*out = new(string)
		**out = **in
	}
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SystemTags != nil {
		in, out := &in.SystemTags, &out.SystemTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpecResource.
func (in *NodePoolSpecResource) DeepCopy() *NodePoolSpecResource {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolStatus) DeepCopyInto(out *NodePoolStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolStatus.
func (in *NodePoolStatus) DeepCopy() *NodePoolStatus {
	if in == nil {
		return nil
	}
	out := new(NodePoolStatus)
	in.DeepCopyInto(out)
	return out
}
