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
func (in *AcceptedAgreement) DeepCopyInto(out *AcceptedAgreement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceptedAgreement.
func (in *AcceptedAgreement) DeepCopy() *AcceptedAgreement {
	if in == nil {
		return nil
	}
	out := new(AcceptedAgreement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcceptedAgreement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceptedAgreementList) DeepCopyInto(out *AcceptedAgreementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AcceptedAgreement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceptedAgreementList.
func (in *AcceptedAgreementList) DeepCopy() *AcceptedAgreementList {
	if in == nil {
		return nil
	}
	out := new(AcceptedAgreementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AcceptedAgreementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceptedAgreementSpec) DeepCopyInto(out *AcceptedAgreementSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AcceptedAgreementSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceptedAgreementSpec.
func (in *AcceptedAgreementSpec) DeepCopy() *AcceptedAgreementSpec {
	if in == nil {
		return nil
	}
	out := new(AcceptedAgreementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceptedAgreementSpecResource) DeepCopyInto(out *AcceptedAgreementSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AgreementID != nil {
		in, out := &in.AgreementID, &out.AgreementID
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
	if in.ListingID != nil {
		in, out := &in.ListingID, &out.ListingID
		*out = new(string)
		**out = **in
	}
	if in.PackageVersion != nil {
		in, out := &in.PackageVersion, &out.PackageVersion
		*out = new(string)
		**out = **in
	}
	if in.Signature != nil {
		in, out := &in.Signature, &out.Signature
		*out = new(string)
		**out = **in
	}
	if in.TimeAccepted != nil {
		in, out := &in.TimeAccepted, &out.TimeAccepted
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceptedAgreementSpecResource.
func (in *AcceptedAgreementSpecResource) DeepCopy() *AcceptedAgreementSpecResource {
	if in == nil {
		return nil
	}
	out := new(AcceptedAgreementSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceptedAgreementStatus) DeepCopyInto(out *AcceptedAgreementStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceptedAgreementStatus.
func (in *AcceptedAgreementStatus) DeepCopy() *AcceptedAgreementStatus {
	if in == nil {
		return nil
	}
	out := new(AcceptedAgreementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListingPackageAgreement) DeepCopyInto(out *ListingPackageAgreement) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListingPackageAgreement.
func (in *ListingPackageAgreement) DeepCopy() *ListingPackageAgreement {
	if in == nil {
		return nil
	}
	out := new(ListingPackageAgreement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ListingPackageAgreement) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListingPackageAgreementList) DeepCopyInto(out *ListingPackageAgreementList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ListingPackageAgreement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListingPackageAgreementList.
func (in *ListingPackageAgreementList) DeepCopy() *ListingPackageAgreementList {
	if in == nil {
		return nil
	}
	out := new(ListingPackageAgreementList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ListingPackageAgreementList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListingPackageAgreementSpec) DeepCopyInto(out *ListingPackageAgreementSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ListingPackageAgreementSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListingPackageAgreementSpec.
func (in *ListingPackageAgreementSpec) DeepCopy() *ListingPackageAgreementSpec {
	if in == nil {
		return nil
	}
	out := new(ListingPackageAgreementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListingPackageAgreementSpecResource) DeepCopyInto(out *ListingPackageAgreementSpecResource) {
	*out = *in
	if in.AgreementID != nil {
		in, out := &in.AgreementID, &out.AgreementID
		*out = new(string)
		**out = **in
	}
	if in.Author != nil {
		in, out := &in.Author, &out.Author
		*out = new(string)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.ContentURL != nil {
		in, out := &in.ContentURL, &out.ContentURL
		*out = new(string)
		**out = **in
	}
	if in.ListingID != nil {
		in, out := &in.ListingID, &out.ListingID
		*out = new(string)
		**out = **in
	}
	if in.PackageVersion != nil {
		in, out := &in.PackageVersion, &out.PackageVersion
		*out = new(string)
		**out = **in
	}
	if in.Prompt != nil {
		in, out := &in.Prompt, &out.Prompt
		*out = new(string)
		**out = **in
	}
	if in.Signature != nil {
		in, out := &in.Signature, &out.Signature
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListingPackageAgreementSpecResource.
func (in *ListingPackageAgreementSpecResource) DeepCopy() *ListingPackageAgreementSpecResource {
	if in == nil {
		return nil
	}
	out := new(ListingPackageAgreementSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListingPackageAgreementStatus) DeepCopyInto(out *ListingPackageAgreementStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListingPackageAgreementStatus.
func (in *ListingPackageAgreementStatus) DeepCopy() *ListingPackageAgreementStatus {
	if in == nil {
		return nil
	}
	out := new(ListingPackageAgreementStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Publication) DeepCopyInto(out *Publication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Publication.
func (in *Publication) DeepCopy() *Publication {
	if in == nil {
		return nil
	}
	out := new(Publication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Publication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationList) DeepCopyInto(out *PublicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Publication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationList.
func (in *PublicationList) DeepCopy() *PublicationList {
	if in == nil {
		return nil
	}
	out := new(PublicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationSpec) DeepCopyInto(out *PublicationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PublicationSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationSpec.
func (in *PublicationSpec) DeepCopy() *PublicationSpec {
	if in == nil {
		return nil
	}
	out := new(PublicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationSpecIcon) DeepCopyInto(out *PublicationSpecIcon) {
	*out = *in
	if in.ContentURL != nil {
		in, out := &in.ContentURL, &out.ContentURL
		*out = new(string)
		**out = **in
	}
	if in.FileExtension != nil {
		in, out := &in.FileExtension, &out.FileExtension
		*out = new(string)
		**out = **in
	}
	if in.MimeType != nil {
		in, out := &in.MimeType, &out.MimeType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationSpecIcon.
func (in *PublicationSpecIcon) DeepCopy() *PublicationSpecIcon {
	if in == nil {
		return nil
	}
	out := new(PublicationSpecIcon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationSpecPackageDetails) DeepCopyInto(out *PublicationSpecPackageDetails) {
	*out = *in
	if in.Eula != nil {
		in, out := &in.Eula, &out.Eula
		*out = make([]PublicationSpecPackageDetailsEula, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.OperatingSystem != nil {
		in, out := &in.OperatingSystem, &out.OperatingSystem
		*out = new(PublicationSpecPackageDetailsOperatingSystem)
		(*in).DeepCopyInto(*out)
	}
	if in.PackageType != nil {
		in, out := &in.PackageType, &out.PackageType
		*out = new(string)
		**out = **in
	}
	if in.PackageVersion != nil {
		in, out := &in.PackageVersion, &out.PackageVersion
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationSpecPackageDetails.
func (in *PublicationSpecPackageDetails) DeepCopy() *PublicationSpecPackageDetails {
	if in == nil {
		return nil
	}
	out := new(PublicationSpecPackageDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationSpecPackageDetailsEula) DeepCopyInto(out *PublicationSpecPackageDetailsEula) {
	*out = *in
	if in.EulaType != nil {
		in, out := &in.EulaType, &out.EulaType
		*out = new(string)
		**out = **in
	}
	if in.LicenseText != nil {
		in, out := &in.LicenseText, &out.LicenseText
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationSpecPackageDetailsEula.
func (in *PublicationSpecPackageDetailsEula) DeepCopy() *PublicationSpecPackageDetailsEula {
	if in == nil {
		return nil
	}
	out := new(PublicationSpecPackageDetailsEula)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationSpecPackageDetailsOperatingSystem) DeepCopyInto(out *PublicationSpecPackageDetailsOperatingSystem) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationSpecPackageDetailsOperatingSystem.
func (in *PublicationSpecPackageDetailsOperatingSystem) DeepCopy() *PublicationSpecPackageDetailsOperatingSystem {
	if in == nil {
		return nil
	}
	out := new(PublicationSpecPackageDetailsOperatingSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationSpecResource) DeepCopyInto(out *PublicationSpecResource) {
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
	if in.FreeformTags != nil {
		in, out := &in.FreeformTags, &out.FreeformTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Icon != nil {
		in, out := &in.Icon, &out.Icon
		*out = new(PublicationSpecIcon)
		(*in).DeepCopyInto(*out)
	}
	if in.IsAgreementAcknowledged != nil {
		in, out := &in.IsAgreementAcknowledged, &out.IsAgreementAcknowledged
		*out = new(bool)
		**out = **in
	}
	if in.ListingType != nil {
		in, out := &in.ListingType, &out.ListingType
		*out = new(string)
		**out = **in
	}
	if in.LongDescription != nil {
		in, out := &in.LongDescription, &out.LongDescription
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PackageDetails != nil {
		in, out := &in.PackageDetails, &out.PackageDetails
		*out = new(PublicationSpecPackageDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.PackageType != nil {
		in, out := &in.PackageType, &out.PackageType
		*out = new(string)
		**out = **in
	}
	if in.ShortDescription != nil {
		in, out := &in.ShortDescription, &out.ShortDescription
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.SupportContacts != nil {
		in, out := &in.SupportContacts, &out.SupportContacts
		*out = make([]PublicationSpecSupportContacts, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SupportedOperatingSystems != nil {
		in, out := &in.SupportedOperatingSystems, &out.SupportedOperatingSystems
		*out = make([]PublicationSpecSupportedOperatingSystems, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TimeCreated != nil {
		in, out := &in.TimeCreated, &out.TimeCreated
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationSpecResource.
func (in *PublicationSpecResource) DeepCopy() *PublicationSpecResource {
	if in == nil {
		return nil
	}
	out := new(PublicationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationSpecSupportContacts) DeepCopyInto(out *PublicationSpecSupportContacts) {
	*out = *in
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Phone != nil {
		in, out := &in.Phone, &out.Phone
		*out = new(string)
		**out = **in
	}
	if in.Subject != nil {
		in, out := &in.Subject, &out.Subject
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationSpecSupportContacts.
func (in *PublicationSpecSupportContacts) DeepCopy() *PublicationSpecSupportContacts {
	if in == nil {
		return nil
	}
	out := new(PublicationSpecSupportContacts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationSpecSupportedOperatingSystems) DeepCopyInto(out *PublicationSpecSupportedOperatingSystems) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationSpecSupportedOperatingSystems.
func (in *PublicationSpecSupportedOperatingSystems) DeepCopy() *PublicationSpecSupportedOperatingSystems {
	if in == nil {
		return nil
	}
	out := new(PublicationSpecSupportedOperatingSystems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicationStatus) DeepCopyInto(out *PublicationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicationStatus.
func (in *PublicationStatus) DeepCopy() *PublicationStatus {
	if in == nil {
		return nil
	}
	out := new(PublicationStatus)
	in.DeepCopyInto(out)
	return out
}
