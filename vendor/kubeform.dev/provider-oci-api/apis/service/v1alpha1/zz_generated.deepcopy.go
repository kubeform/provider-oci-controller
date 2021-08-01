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
func (in *CatalogPrivateApplication) DeepCopyInto(out *CatalogPrivateApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogPrivateApplication.
func (in *CatalogPrivateApplication) DeepCopy() *CatalogPrivateApplication {
	if in == nil {
		return nil
	}
	out := new(CatalogPrivateApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogPrivateApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogPrivateApplicationList) DeepCopyInto(out *CatalogPrivateApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CatalogPrivateApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogPrivateApplicationList.
func (in *CatalogPrivateApplicationList) DeepCopy() *CatalogPrivateApplicationList {
	if in == nil {
		return nil
	}
	out := new(CatalogPrivateApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogPrivateApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogPrivateApplicationSpec) DeepCopyInto(out *CatalogPrivateApplicationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CatalogPrivateApplicationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogPrivateApplicationSpec.
func (in *CatalogPrivateApplicationSpec) DeepCopy() *CatalogPrivateApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogPrivateApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogPrivateApplicationSpecLogo) DeepCopyInto(out *CatalogPrivateApplicationSpecLogo) {
	*out = *in
	if in.ContentURL != nil {
		in, out := &in.ContentURL, &out.ContentURL
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.MimeType != nil {
		in, out := &in.MimeType, &out.MimeType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogPrivateApplicationSpecLogo.
func (in *CatalogPrivateApplicationSpecLogo) DeepCopy() *CatalogPrivateApplicationSpecLogo {
	if in == nil {
		return nil
	}
	out := new(CatalogPrivateApplicationSpecLogo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogPrivateApplicationSpecPackageDetails) DeepCopyInto(out *CatalogPrivateApplicationSpecPackageDetails) {
	*out = *in
	if in.PackageType != nil {
		in, out := &in.PackageType, &out.PackageType
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.ZipFileBase64encoded != nil {
		in, out := &in.ZipFileBase64encoded, &out.ZipFileBase64encoded
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogPrivateApplicationSpecPackageDetails.
func (in *CatalogPrivateApplicationSpecPackageDetails) DeepCopy() *CatalogPrivateApplicationSpecPackageDetails {
	if in == nil {
		return nil
	}
	out := new(CatalogPrivateApplicationSpecPackageDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogPrivateApplicationSpecResource) DeepCopyInto(out *CatalogPrivateApplicationSpecResource) {
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
	if in.Logo != nil {
		in, out := &in.Logo, &out.Logo
		*out = new(CatalogPrivateApplicationSpecLogo)
		(*in).DeepCopyInto(*out)
	}
	if in.LogoFileBase64encoded != nil {
		in, out := &in.LogoFileBase64encoded, &out.LogoFileBase64encoded
		*out = new(string)
		**out = **in
	}
	if in.LongDescription != nil {
		in, out := &in.LongDescription, &out.LongDescription
		*out = new(string)
		**out = **in
	}
	if in.PackageDetails != nil {
		in, out := &in.PackageDetails, &out.PackageDetails
		*out = new(CatalogPrivateApplicationSpecPackageDetails)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogPrivateApplicationSpecResource.
func (in *CatalogPrivateApplicationSpecResource) DeepCopy() *CatalogPrivateApplicationSpecResource {
	if in == nil {
		return nil
	}
	out := new(CatalogPrivateApplicationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogPrivateApplicationStatus) DeepCopyInto(out *CatalogPrivateApplicationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogPrivateApplicationStatus.
func (in *CatalogPrivateApplicationStatus) DeepCopy() *CatalogPrivateApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogPrivateApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalog) DeepCopyInto(out *CatalogServiceCatalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalog.
func (in *CatalogServiceCatalog) DeepCopy() *CatalogServiceCatalog {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogServiceCatalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalogAssociation) DeepCopyInto(out *CatalogServiceCatalogAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalogAssociation.
func (in *CatalogServiceCatalogAssociation) DeepCopy() *CatalogServiceCatalogAssociation {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalogAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogServiceCatalogAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalogAssociationList) DeepCopyInto(out *CatalogServiceCatalogAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CatalogServiceCatalogAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalogAssociationList.
func (in *CatalogServiceCatalogAssociationList) DeepCopy() *CatalogServiceCatalogAssociationList {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalogAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogServiceCatalogAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalogAssociationSpec) DeepCopyInto(out *CatalogServiceCatalogAssociationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CatalogServiceCatalogAssociationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalogAssociationSpec.
func (in *CatalogServiceCatalogAssociationSpec) DeepCopy() *CatalogServiceCatalogAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalogAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalogAssociationSpecResource) DeepCopyInto(out *CatalogServiceCatalogAssociationSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.EntityID != nil {
		in, out := &in.EntityID, &out.EntityID
		*out = new(string)
		**out = **in
	}
	if in.EntityType != nil {
		in, out := &in.EntityType, &out.EntityType
		*out = new(string)
		**out = **in
	}
	if in.ServiceCatalogID != nil {
		in, out := &in.ServiceCatalogID, &out.ServiceCatalogID
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalogAssociationSpecResource.
func (in *CatalogServiceCatalogAssociationSpecResource) DeepCopy() *CatalogServiceCatalogAssociationSpecResource {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalogAssociationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalogAssociationStatus) DeepCopyInto(out *CatalogServiceCatalogAssociationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalogAssociationStatus.
func (in *CatalogServiceCatalogAssociationStatus) DeepCopy() *CatalogServiceCatalogAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalogAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalogList) DeepCopyInto(out *CatalogServiceCatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CatalogServiceCatalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalogList.
func (in *CatalogServiceCatalogList) DeepCopy() *CatalogServiceCatalogList {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogServiceCatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalogSpec) DeepCopyInto(out *CatalogServiceCatalogSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CatalogServiceCatalogSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalogSpec.
func (in *CatalogServiceCatalogSpec) DeepCopy() *CatalogServiceCatalogSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalogSpecResource) DeepCopyInto(out *CatalogServiceCatalogSpecResource) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalogSpecResource.
func (in *CatalogServiceCatalogSpecResource) DeepCopy() *CatalogServiceCatalogSpecResource {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalogSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogServiceCatalogStatus) DeepCopyInto(out *CatalogServiceCatalogStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogServiceCatalogStatus.
func (in *CatalogServiceCatalogStatus) DeepCopy() *CatalogServiceCatalogStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogServiceCatalogStatus)
	in.DeepCopyInto(out)
	return out
}
