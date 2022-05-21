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
func (in *EsxiHost) DeepCopyInto(out *EsxiHost) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsxiHost.
func (in *EsxiHost) DeepCopy() *EsxiHost {
	if in == nil {
		return nil
	}
	out := new(EsxiHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EsxiHost) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsxiHostList) DeepCopyInto(out *EsxiHostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EsxiHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsxiHostList.
func (in *EsxiHostList) DeepCopy() *EsxiHostList {
	if in == nil {
		return nil
	}
	out := new(EsxiHostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EsxiHostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsxiHostSpec) DeepCopyInto(out *EsxiHostSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(EsxiHostSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsxiHostSpec.
func (in *EsxiHostSpec) DeepCopy() *EsxiHostSpec {
	if in == nil {
		return nil
	}
	out := new(EsxiHostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsxiHostSpecResource) DeepCopyInto(out *EsxiHostSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.BillingContractEndDate != nil {
		in, out := &in.BillingContractEndDate, &out.BillingContractEndDate
		*out = new(string)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.ComputeAvailabilityDomain != nil {
		in, out := &in.ComputeAvailabilityDomain, &out.ComputeAvailabilityDomain
		*out = new(string)
		**out = **in
	}
	if in.ComputeInstanceID != nil {
		in, out := &in.ComputeInstanceID, &out.ComputeInstanceID
		*out = new(string)
		**out = **in
	}
	if in.CurrentSku != nil {
		in, out := &in.CurrentSku, &out.CurrentSku
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
	if in.FailedEsxiHostID != nil {
		in, out := &in.FailedEsxiHostID, &out.FailedEsxiHostID
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
	if in.GracePeriodEndDate != nil {
		in, out := &in.GracePeriodEndDate, &out.GracePeriodEndDate
		*out = new(string)
		**out = **in
	}
	if in.NextSku != nil {
		in, out := &in.NextSku, &out.NextSku
		*out = new(string)
		**out = **in
	}
	if in.ReplacementEsxiHostID != nil {
		in, out := &in.ReplacementEsxiHostID, &out.ReplacementEsxiHostID
		*out = new(string)
		**out = **in
	}
	if in.SddcID != nil {
		in, out := &in.SddcID, &out.SddcID
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsxiHostSpecResource.
func (in *EsxiHostSpecResource) DeepCopy() *EsxiHostSpecResource {
	if in == nil {
		return nil
	}
	out := new(EsxiHostSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EsxiHostStatus) DeepCopyInto(out *EsxiHostStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EsxiHostStatus.
func (in *EsxiHostStatus) DeepCopy() *EsxiHostStatus {
	if in == nil {
		return nil
	}
	out := new(EsxiHostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sddc) DeepCopyInto(out *Sddc) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sddc.
func (in *Sddc) DeepCopy() *Sddc {
	if in == nil {
		return nil
	}
	out := new(Sddc)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sddc) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SddcList) DeepCopyInto(out *SddcList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sddc, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SddcList.
func (in *SddcList) DeepCopy() *SddcList {
	if in == nil {
		return nil
	}
	out := new(SddcList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SddcList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SddcSpec) DeepCopyInto(out *SddcSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SddcSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SddcSpec.
func (in *SddcSpec) DeepCopy() *SddcSpec {
	if in == nil {
		return nil
	}
	out := new(SddcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SddcSpecHcxOnPremLicenses) DeepCopyInto(out *SddcSpecHcxOnPremLicenses) {
	*out = *in
	if in.ActivationKey != nil {
		in, out := &in.ActivationKey, &out.ActivationKey
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SystemName != nil {
		in, out := &in.SystemName, &out.SystemName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SddcSpecHcxOnPremLicenses.
func (in *SddcSpecHcxOnPremLicenses) DeepCopy() *SddcSpecHcxOnPremLicenses {
	if in == nil {
		return nil
	}
	out := new(SddcSpecHcxOnPremLicenses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SddcSpecResource) DeepCopyInto(out *SddcSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ActualEsxiHostsCount != nil {
		in, out := &in.ActualEsxiHostsCount, &out.ActualEsxiHostsCount
		*out = new(int64)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.ComputeAvailabilityDomain != nil {
		in, out := &in.ComputeAvailabilityDomain, &out.ComputeAvailabilityDomain
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
	if in.EsxiHostsCount != nil {
		in, out := &in.EsxiHostsCount, &out.EsxiHostsCount
		*out = new(int64)
		**out = **in
	}
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HcxAction != nil {
		in, out := &in.HcxAction, &out.HcxAction
		*out = new(string)
		**out = **in
	}
	if in.HcxFqdn != nil {
		in, out := &in.HcxFqdn, &out.HcxFqdn
		*out = new(string)
		**out = **in
	}
	if in.HcxInitialPassword != nil {
		in, out := &in.HcxInitialPassword, &out.HcxInitialPassword
		*out = new(string)
		**out = **in
	}
	if in.HcxOnPremKey != nil {
		in, out := &in.HcxOnPremKey, &out.HcxOnPremKey
		*out = new(string)
		**out = **in
	}
	if in.HcxOnPremLicenses != nil {
		in, out := &in.HcxOnPremLicenses, &out.HcxOnPremLicenses
		*out = make([]SddcSpecHcxOnPremLicenses, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HcxPrivateIPID != nil {
		in, out := &in.HcxPrivateIPID, &out.HcxPrivateIPID
		*out = new(string)
		**out = **in
	}
	if in.HcxVLANID != nil {
		in, out := &in.HcxVLANID, &out.HcxVLANID
		*out = new(string)
		**out = **in
	}
	if in.InitialSku != nil {
		in, out := &in.InitialSku, &out.InitialSku
		*out = new(string)
		**out = **in
	}
	if in.InstanceDisplayNamePrefix != nil {
		in, out := &in.InstanceDisplayNamePrefix, &out.InstanceDisplayNamePrefix
		*out = new(string)
		**out = **in
	}
	if in.IsHcxEnabled != nil {
		in, out := &in.IsHcxEnabled, &out.IsHcxEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsHcxEnterpriseEnabled != nil {
		in, out := &in.IsHcxEnterpriseEnabled, &out.IsHcxEnterpriseEnabled
		*out = new(bool)
		**out = **in
	}
	if in.IsHcxPendingDowngrade != nil {
		in, out := &in.IsHcxPendingDowngrade, &out.IsHcxPendingDowngrade
		*out = new(bool)
		**out = **in
	}
	if in.NsxEdgeUplink1vlanID != nil {
		in, out := &in.NsxEdgeUplink1vlanID, &out.NsxEdgeUplink1vlanID
		*out = new(string)
		**out = **in
	}
	if in.NsxEdgeUplink2vlanID != nil {
		in, out := &in.NsxEdgeUplink2vlanID, &out.NsxEdgeUplink2vlanID
		*out = new(string)
		**out = **in
	}
	if in.NsxEdgeUplinkIPID != nil {
		in, out := &in.NsxEdgeUplinkIPID, &out.NsxEdgeUplinkIPID
		*out = new(string)
		**out = **in
	}
	if in.NsxEdgeVtepVLANID != nil {
		in, out := &in.NsxEdgeVtepVLANID, &out.NsxEdgeVtepVLANID
		*out = new(string)
		**out = **in
	}
	if in.NsxManagerFqdn != nil {
		in, out := &in.NsxManagerFqdn, &out.NsxManagerFqdn
		*out = new(string)
		**out = **in
	}
	if in.NsxManagerInitialPassword != nil {
		in, out := &in.NsxManagerInitialPassword, &out.NsxManagerInitialPassword
		*out = new(string)
		**out = **in
	}
	if in.NsxManagerPrivateIPID != nil {
		in, out := &in.NsxManagerPrivateIPID, &out.NsxManagerPrivateIPID
		*out = new(string)
		**out = **in
	}
	if in.NsxManagerUsername != nil {
		in, out := &in.NsxManagerUsername, &out.NsxManagerUsername
		*out = new(string)
		**out = **in
	}
	if in.NsxOverlaySegmentName != nil {
		in, out := &in.NsxOverlaySegmentName, &out.NsxOverlaySegmentName
		*out = new(string)
		**out = **in
	}
	if in.NsxVtepVLANID != nil {
		in, out := &in.NsxVtepVLANID, &out.NsxVtepVLANID
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningSubnetID != nil {
		in, out := &in.ProvisioningSubnetID, &out.ProvisioningSubnetID
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningVLANID != nil {
		in, out := &in.ProvisioningVLANID, &out.ProvisioningVLANID
		*out = new(string)
		**out = **in
	}
	if in.RefreshHcxLicenseStatus != nil {
		in, out := &in.RefreshHcxLicenseStatus, &out.RefreshHcxLicenseStatus
		*out = new(bool)
		**out = **in
	}
	if in.ReplicationVLANID != nil {
		in, out := &in.ReplicationVLANID, &out.ReplicationVLANID
		*out = new(string)
		**out = **in
	}
	if in.ReservingHcxOnPremiseLicenseKeys != nil {
		in, out := &in.ReservingHcxOnPremiseLicenseKeys, &out.ReservingHcxOnPremiseLicenseKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SshAuthorizedKeys != nil {
		in, out := &in.SshAuthorizedKeys, &out.SshAuthorizedKeys
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
	if in.TimeHcxBillingCycleEnd != nil {
		in, out := &in.TimeHcxBillingCycleEnd, &out.TimeHcxBillingCycleEnd
		*out = new(string)
		**out = **in
	}
	if in.TimeHcxLicenseStatusUpdated != nil {
		in, out := &in.TimeHcxLicenseStatusUpdated, &out.TimeHcxLicenseStatusUpdated
		*out = new(string)
		**out = **in
	}
	if in.TimeUpdated != nil {
		in, out := &in.TimeUpdated, &out.TimeUpdated
		*out = new(string)
		**out = **in
	}
	if in.VcenterFqdn != nil {
		in, out := &in.VcenterFqdn, &out.VcenterFqdn
		*out = new(string)
		**out = **in
	}
	if in.VcenterInitialPassword != nil {
		in, out := &in.VcenterInitialPassword, &out.VcenterInitialPassword
		*out = new(string)
		**out = **in
	}
	if in.VcenterPrivateIPID != nil {
		in, out := &in.VcenterPrivateIPID, &out.VcenterPrivateIPID
		*out = new(string)
		**out = **in
	}
	if in.VcenterUsername != nil {
		in, out := &in.VcenterUsername, &out.VcenterUsername
		*out = new(string)
		**out = **in
	}
	if in.VmotionVLANID != nil {
		in, out := &in.VmotionVLANID, &out.VmotionVLANID
		*out = new(string)
		**out = **in
	}
	if in.VmwareSoftwareVersion != nil {
		in, out := &in.VmwareSoftwareVersion, &out.VmwareSoftwareVersion
		*out = new(string)
		**out = **in
	}
	if in.VsanVLANID != nil {
		in, out := &in.VsanVLANID, &out.VsanVLANID
		*out = new(string)
		**out = **in
	}
	if in.VsphereVLANID != nil {
		in, out := &in.VsphereVLANID, &out.VsphereVLANID
		*out = new(string)
		**out = **in
	}
	if in.WorkloadNetworkCIDR != nil {
		in, out := &in.WorkloadNetworkCIDR, &out.WorkloadNetworkCIDR
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SddcSpecResource.
func (in *SddcSpecResource) DeepCopy() *SddcSpecResource {
	if in == nil {
		return nil
	}
	out := new(SddcSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SddcStatus) DeepCopyInto(out *SddcStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SddcStatus.
func (in *SddcStatus) DeepCopy() *SddcStatus {
	if in == nil {
		return nil
	}
	out := new(SddcStatus)
	in.DeepCopyInto(out)
	return out
}