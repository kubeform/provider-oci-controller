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

	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstance) DeepCopyInto(out *AnalyticsInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstance.
func (in *AnalyticsInstance) DeepCopy() *AnalyticsInstance {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyticsInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceList) DeepCopyInto(out *AnalyticsInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnalyticsInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceList.
func (in *AnalyticsInstanceList) DeepCopy() *AnalyticsInstanceList {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyticsInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstancePrivateAccessChannel) DeepCopyInto(out *AnalyticsInstancePrivateAccessChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstancePrivateAccessChannel.
func (in *AnalyticsInstancePrivateAccessChannel) DeepCopy() *AnalyticsInstancePrivateAccessChannel {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstancePrivateAccessChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyticsInstancePrivateAccessChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstancePrivateAccessChannelList) DeepCopyInto(out *AnalyticsInstancePrivateAccessChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnalyticsInstancePrivateAccessChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstancePrivateAccessChannelList.
func (in *AnalyticsInstancePrivateAccessChannelList) DeepCopy() *AnalyticsInstancePrivateAccessChannelList {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstancePrivateAccessChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyticsInstancePrivateAccessChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstancePrivateAccessChannelSpec) DeepCopyInto(out *AnalyticsInstancePrivateAccessChannelSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AnalyticsInstancePrivateAccessChannelSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstancePrivateAccessChannelSpec.
func (in *AnalyticsInstancePrivateAccessChannelSpec) DeepCopy() *AnalyticsInstancePrivateAccessChannelSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstancePrivateAccessChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstancePrivateAccessChannelSpecPrivateSourceDNSZones) DeepCopyInto(out *AnalyticsInstancePrivateAccessChannelSpecPrivateSourceDNSZones) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DnsZone != nil {
		in, out := &in.DnsZone, &out.DnsZone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstancePrivateAccessChannelSpecPrivateSourceDNSZones.
func (in *AnalyticsInstancePrivateAccessChannelSpecPrivateSourceDNSZones) DeepCopy() *AnalyticsInstancePrivateAccessChannelSpecPrivateSourceDNSZones {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstancePrivateAccessChannelSpecPrivateSourceDNSZones)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstancePrivateAccessChannelSpecResource) DeepCopyInto(out *AnalyticsInstancePrivateAccessChannelSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AnalyticsInstanceID != nil {
		in, out := &in.AnalyticsInstanceID, &out.AnalyticsInstanceID
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EgressSourceIPAddresses != nil {
		in, out := &in.EgressSourceIPAddresses, &out.EgressSourceIPAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.PrivateSourceDNSZones != nil {
		in, out := &in.PrivateSourceDNSZones, &out.PrivateSourceDNSZones
		*out = make([]AnalyticsInstancePrivateAccessChannelSpecPrivateSourceDNSZones, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VcnID != nil {
		in, out := &in.VcnID, &out.VcnID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstancePrivateAccessChannelSpecResource.
func (in *AnalyticsInstancePrivateAccessChannelSpecResource) DeepCopy() *AnalyticsInstancePrivateAccessChannelSpecResource {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstancePrivateAccessChannelSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstancePrivateAccessChannelStatus) DeepCopyInto(out *AnalyticsInstancePrivateAccessChannelStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstancePrivateAccessChannelStatus.
func (in *AnalyticsInstancePrivateAccessChannelStatus) DeepCopy() *AnalyticsInstancePrivateAccessChannelStatus {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstancePrivateAccessChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceSpec) DeepCopyInto(out *AnalyticsInstanceSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AnalyticsInstanceSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceSpec.
func (in *AnalyticsInstanceSpec) DeepCopy() *AnalyticsInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceSpecCapacity) DeepCopyInto(out *AnalyticsInstanceSpecCapacity) {
	*out = *in
	if in.CapacityType != nil {
		in, out := &in.CapacityType, &out.CapacityType
		*out = new(string)
		**out = **in
	}
	if in.CapacityValue != nil {
		in, out := &in.CapacityValue, &out.CapacityValue
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceSpecCapacity.
func (in *AnalyticsInstanceSpecCapacity) DeepCopy() *AnalyticsInstanceSpecCapacity {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceSpecCapacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceSpecNetworkEndpointDetails) DeepCopyInto(out *AnalyticsInstanceSpecNetworkEndpointDetails) {
	*out = *in
	if in.NetworkEndpointType != nil {
		in, out := &in.NetworkEndpointType, &out.NetworkEndpointType
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	if in.VcnID != nil {
		in, out := &in.VcnID, &out.VcnID
		*out = new(string)
		**out = **in
	}
	if in.WhitelistedIPS != nil {
		in, out := &in.WhitelistedIPS, &out.WhitelistedIPS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WhitelistedVcns != nil {
		in, out := &in.WhitelistedVcns, &out.WhitelistedVcns
		*out = make([]AnalyticsInstanceSpecNetworkEndpointDetailsWhitelistedVcns, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceSpecNetworkEndpointDetails.
func (in *AnalyticsInstanceSpecNetworkEndpointDetails) DeepCopy() *AnalyticsInstanceSpecNetworkEndpointDetails {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceSpecNetworkEndpointDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceSpecNetworkEndpointDetailsWhitelistedVcns) DeepCopyInto(out *AnalyticsInstanceSpecNetworkEndpointDetailsWhitelistedVcns) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.WhitelistedIPS != nil {
		in, out := &in.WhitelistedIPS, &out.WhitelistedIPS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceSpecNetworkEndpointDetailsWhitelistedVcns.
func (in *AnalyticsInstanceSpecNetworkEndpointDetailsWhitelistedVcns) DeepCopy() *AnalyticsInstanceSpecNetworkEndpointDetailsWhitelistedVcns {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceSpecNetworkEndpointDetailsWhitelistedVcns)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceSpecResource) DeepCopyInto(out *AnalyticsInstanceSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = new(AnalyticsInstanceSpecCapacity)
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
	if in.EmailNotification != nil {
		in, out := &in.EmailNotification, &out.EmailNotification
		*out = new(string)
		**out = **in
	}
	if in.FeatureSet != nil {
		in, out := &in.FeatureSet, &out.FeatureSet
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
	if in.IdcsAccessToken != nil {
		in, out := &in.IdcsAccessToken, &out.IdcsAccessToken
		*out = new(string)
		**out = **in
	}
	if in.LicenseType != nil {
		in, out := &in.LicenseType, &out.LicenseType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkEndpointDetails != nil {
		in, out := &in.NetworkEndpointDetails, &out.NetworkEndpointDetails
		*out = new(AnalyticsInstanceSpecNetworkEndpointDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.PrivateAccessChannels != nil {
		in, out := &in.PrivateAccessChannels, &out.PrivateAccessChannels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceURL != nil {
		in, out := &in.ServiceURL, &out.ServiceURL
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
	if in.VanityURLDetails != nil {
		in, out := &in.VanityURLDetails, &out.VanityURLDetails
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceSpecResource.
func (in *AnalyticsInstanceSpecResource) DeepCopy() *AnalyticsInstanceSpecResource {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceStatus) DeepCopyInto(out *AnalyticsInstanceStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceStatus.
func (in *AnalyticsInstanceStatus) DeepCopy() *AnalyticsInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceVanityURL) DeepCopyInto(out *AnalyticsInstanceVanityURL) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceVanityURL.
func (in *AnalyticsInstanceVanityURL) DeepCopy() *AnalyticsInstanceVanityURL {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceVanityURL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyticsInstanceVanityURL) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceVanityURLList) DeepCopyInto(out *AnalyticsInstanceVanityURLList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnalyticsInstanceVanityURL, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceVanityURLList.
func (in *AnalyticsInstanceVanityURLList) DeepCopy() *AnalyticsInstanceVanityURLList {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceVanityURLList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnalyticsInstanceVanityURLList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceVanityURLSpec) DeepCopyInto(out *AnalyticsInstanceVanityURLSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AnalyticsInstanceVanityURLSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceVanityURLSpec.
func (in *AnalyticsInstanceVanityURLSpec) DeepCopy() *AnalyticsInstanceVanityURLSpec {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceVanityURLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceVanityURLSpecResource) DeepCopyInto(out *AnalyticsInstanceVanityURLSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AnalyticsInstanceID != nil {
		in, out := &in.AnalyticsInstanceID, &out.AnalyticsInstanceID
		*out = new(string)
		**out = **in
	}
	if in.CaCertificate != nil {
		in, out := &in.CaCertificate, &out.CaCertificate
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Passphrase != nil {
		in, out := &in.Passphrase, &out.Passphrase
		*out = new(string)
		**out = **in
	}
	if in.PrivateKey != nil {
		in, out := &in.PrivateKey, &out.PrivateKey
		*out = new(string)
		**out = **in
	}
	if in.PublicCertificate != nil {
		in, out := &in.PublicCertificate, &out.PublicCertificate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceVanityURLSpecResource.
func (in *AnalyticsInstanceVanityURLSpecResource) DeepCopy() *AnalyticsInstanceVanityURLSpecResource {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceVanityURLSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnalyticsInstanceVanityURLStatus) DeepCopyInto(out *AnalyticsInstanceVanityURLStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnalyticsInstanceVanityURLStatus.
func (in *AnalyticsInstanceVanityURLStatus) DeepCopy() *AnalyticsInstanceVanityURLStatus {
	if in == nil {
		return nil
	}
	out := new(AnalyticsInstanceVanityURLStatus)
	in.DeepCopyInto(out)
	return out
}