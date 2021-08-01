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
func (in *ComputationCustomTable) DeepCopyInto(out *ComputationCustomTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationCustomTable.
func (in *ComputationCustomTable) DeepCopy() *ComputationCustomTable {
	if in == nil {
		return nil
	}
	out := new(ComputationCustomTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputationCustomTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationCustomTableList) DeepCopyInto(out *ComputationCustomTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputationCustomTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationCustomTableList.
func (in *ComputationCustomTableList) DeepCopy() *ComputationCustomTableList {
	if in == nil {
		return nil
	}
	out := new(ComputationCustomTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputationCustomTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationCustomTableSpec) DeepCopyInto(out *ComputationCustomTableSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ComputationCustomTableSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationCustomTableSpec.
func (in *ComputationCustomTableSpec) DeepCopy() *ComputationCustomTableSpec {
	if in == nil {
		return nil
	}
	out := new(ComputationCustomTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationCustomTableSpecResource) DeepCopyInto(out *ComputationCustomTableSpecResource) {
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
	if in.SavedCustomTable != nil {
		in, out := &in.SavedCustomTable, &out.SavedCustomTable
		*out = new(ComputationCustomTableSpecSavedCustomTable)
		(*in).DeepCopyInto(*out)
	}
	if in.SavedReportID != nil {
		in, out := &in.SavedReportID, &out.SavedReportID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationCustomTableSpecResource.
func (in *ComputationCustomTableSpecResource) DeepCopy() *ComputationCustomTableSpecResource {
	if in == nil {
		return nil
	}
	out := new(ComputationCustomTableSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationCustomTableSpecSavedCustomTable) DeepCopyInto(out *ComputationCustomTableSpecSavedCustomTable) {
	*out = *in
	if in.ColumnGroupBy != nil {
		in, out := &in.ColumnGroupBy, &out.ColumnGroupBy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CompartmentDepth != nil {
		in, out := &in.CompartmentDepth, &out.CompartmentDepth
		*out = new(float64)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.GroupByTag != nil {
		in, out := &in.GroupByTag, &out.GroupByTag
		*out = make([]ComputationCustomTableSpecSavedCustomTableGroupByTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RowGroupBy != nil {
		in, out := &in.RowGroupBy, &out.RowGroupBy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationCustomTableSpecSavedCustomTable.
func (in *ComputationCustomTableSpecSavedCustomTable) DeepCopy() *ComputationCustomTableSpecSavedCustomTable {
	if in == nil {
		return nil
	}
	out := new(ComputationCustomTableSpecSavedCustomTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationCustomTableSpecSavedCustomTableGroupByTag) DeepCopyInto(out *ComputationCustomTableSpecSavedCustomTableGroupByTag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationCustomTableSpecSavedCustomTableGroupByTag.
func (in *ComputationCustomTableSpecSavedCustomTableGroupByTag) DeepCopy() *ComputationCustomTableSpecSavedCustomTableGroupByTag {
	if in == nil {
		return nil
	}
	out := new(ComputationCustomTableSpecSavedCustomTableGroupByTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationCustomTableStatus) DeepCopyInto(out *ComputationCustomTableStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationCustomTableStatus.
func (in *ComputationCustomTableStatus) DeepCopy() *ComputationCustomTableStatus {
	if in == nil {
		return nil
	}
	out := new(ComputationCustomTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQuery) DeepCopyInto(out *ComputationQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQuery.
func (in *ComputationQuery) DeepCopy() *ComputationQuery {
	if in == nil {
		return nil
	}
	out := new(ComputationQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputationQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQueryList) DeepCopyInto(out *ComputationQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputationQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQueryList.
func (in *ComputationQueryList) DeepCopy() *ComputationQueryList {
	if in == nil {
		return nil
	}
	out := new(ComputationQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputationQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQuerySpec) DeepCopyInto(out *ComputationQuerySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ComputationQuerySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQuerySpec.
func (in *ComputationQuerySpec) DeepCopy() *ComputationQuerySpec {
	if in == nil {
		return nil
	}
	out := new(ComputationQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQuerySpecQueryDefinition) DeepCopyInto(out *ComputationQuerySpecQueryDefinition) {
	*out = *in
	if in.CostAnalysisUi != nil {
		in, out := &in.CostAnalysisUi, &out.CostAnalysisUi
		*out = new(ComputationQuerySpecQueryDefinitionCostAnalysisUi)
		(*in).DeepCopyInto(*out)
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.ReportQuery != nil {
		in, out := &in.ReportQuery, &out.ReportQuery
		*out = new(ComputationQuerySpecQueryDefinitionReportQuery)
		(*in).DeepCopyInto(*out)
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQuerySpecQueryDefinition.
func (in *ComputationQuerySpecQueryDefinition) DeepCopy() *ComputationQuerySpecQueryDefinition {
	if in == nil {
		return nil
	}
	out := new(ComputationQuerySpecQueryDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQuerySpecQueryDefinitionCostAnalysisUi) DeepCopyInto(out *ComputationQuerySpecQueryDefinitionCostAnalysisUi) {
	*out = *in
	if in.Graph != nil {
		in, out := &in.Graph, &out.Graph
		*out = new(string)
		**out = **in
	}
	if in.IsCumulativeGraph != nil {
		in, out := &in.IsCumulativeGraph, &out.IsCumulativeGraph
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQuerySpecQueryDefinitionCostAnalysisUi.
func (in *ComputationQuerySpecQueryDefinitionCostAnalysisUi) DeepCopy() *ComputationQuerySpecQueryDefinitionCostAnalysisUi {
	if in == nil {
		return nil
	}
	out := new(ComputationQuerySpecQueryDefinitionCostAnalysisUi)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQuerySpecQueryDefinitionReportQuery) DeepCopyInto(out *ComputationQuerySpecQueryDefinitionReportQuery) {
	*out = *in
	if in.CompartmentDepth != nil {
		in, out := &in.CompartmentDepth, &out.CompartmentDepth
		*out = new(float64)
		**out = **in
	}
	if in.DateRangeName != nil {
		in, out := &in.DateRangeName, &out.DateRangeName
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.Forecast != nil {
		in, out := &in.Forecast, &out.Forecast
		*out = new(ComputationQuerySpecQueryDefinitionReportQueryForecast)
		(*in).DeepCopyInto(*out)
	}
	if in.Granularity != nil {
		in, out := &in.Granularity, &out.Granularity
		*out = new(string)
		**out = **in
	}
	if in.GroupBy != nil {
		in, out := &in.GroupBy, &out.GroupBy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GroupByTag != nil {
		in, out := &in.GroupByTag, &out.GroupByTag
		*out = make([]ComputationQuerySpecQueryDefinitionReportQueryGroupByTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsAggregateByTime != nil {
		in, out := &in.IsAggregateByTime, &out.IsAggregateByTime
		*out = new(bool)
		**out = **in
	}
	if in.QueryType != nil {
		in, out := &in.QueryType, &out.QueryType
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.TimeUsageEnded != nil {
		in, out := &in.TimeUsageEnded, &out.TimeUsageEnded
		*out = new(string)
		**out = **in
	}
	if in.TimeUsageStarted != nil {
		in, out := &in.TimeUsageStarted, &out.TimeUsageStarted
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQuerySpecQueryDefinitionReportQuery.
func (in *ComputationQuerySpecQueryDefinitionReportQuery) DeepCopy() *ComputationQuerySpecQueryDefinitionReportQuery {
	if in == nil {
		return nil
	}
	out := new(ComputationQuerySpecQueryDefinitionReportQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQuerySpecQueryDefinitionReportQueryForecast) DeepCopyInto(out *ComputationQuerySpecQueryDefinitionReportQueryForecast) {
	*out = *in
	if in.ForecastType != nil {
		in, out := &in.ForecastType, &out.ForecastType
		*out = new(string)
		**out = **in
	}
	if in.TimeForecastEnded != nil {
		in, out := &in.TimeForecastEnded, &out.TimeForecastEnded
		*out = new(string)
		**out = **in
	}
	if in.TimeForecastStarted != nil {
		in, out := &in.TimeForecastStarted, &out.TimeForecastStarted
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQuerySpecQueryDefinitionReportQueryForecast.
func (in *ComputationQuerySpecQueryDefinitionReportQueryForecast) DeepCopy() *ComputationQuerySpecQueryDefinitionReportQueryForecast {
	if in == nil {
		return nil
	}
	out := new(ComputationQuerySpecQueryDefinitionReportQueryForecast)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQuerySpecQueryDefinitionReportQueryGroupByTag) DeepCopyInto(out *ComputationQuerySpecQueryDefinitionReportQueryGroupByTag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQuerySpecQueryDefinitionReportQueryGroupByTag.
func (in *ComputationQuerySpecQueryDefinitionReportQueryGroupByTag) DeepCopy() *ComputationQuerySpecQueryDefinitionReportQueryGroupByTag {
	if in == nil {
		return nil
	}
	out := new(ComputationQuerySpecQueryDefinitionReportQueryGroupByTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQuerySpecResource) DeepCopyInto(out *ComputationQuerySpecResource) {
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
	if in.QueryDefinition != nil {
		in, out := &in.QueryDefinition, &out.QueryDefinition
		*out = new(ComputationQuerySpecQueryDefinition)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQuerySpecResource.
func (in *ComputationQuerySpecResource) DeepCopy() *ComputationQuerySpecResource {
	if in == nil {
		return nil
	}
	out := new(ComputationQuerySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationQueryStatus) DeepCopyInto(out *ComputationQueryStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationQueryStatus.
func (in *ComputationQueryStatus) DeepCopy() *ComputationQueryStatus {
	if in == nil {
		return nil
	}
	out := new(ComputationQueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationUsage) DeepCopyInto(out *ComputationUsage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationUsage.
func (in *ComputationUsage) DeepCopy() *ComputationUsage {
	if in == nil {
		return nil
	}
	out := new(ComputationUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputationUsage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationUsageList) DeepCopyInto(out *ComputationUsageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputationUsage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationUsageList.
func (in *ComputationUsageList) DeepCopy() *ComputationUsageList {
	if in == nil {
		return nil
	}
	out := new(ComputationUsageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ComputationUsageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationUsageSpec) DeepCopyInto(out *ComputationUsageSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ComputationUsageSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationUsageSpec.
func (in *ComputationUsageSpec) DeepCopy() *ComputationUsageSpec {
	if in == nil {
		return nil
	}
	out := new(ComputationUsageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationUsageSpecForecast) DeepCopyInto(out *ComputationUsageSpecForecast) {
	*out = *in
	if in.ForecastType != nil {
		in, out := &in.ForecastType, &out.ForecastType
		*out = new(string)
		**out = **in
	}
	if in.TimeForecastEnded != nil {
		in, out := &in.TimeForecastEnded, &out.TimeForecastEnded
		*out = new(string)
		**out = **in
	}
	if in.TimeForecastStarted != nil {
		in, out := &in.TimeForecastStarted, &out.TimeForecastStarted
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationUsageSpecForecast.
func (in *ComputationUsageSpecForecast) DeepCopy() *ComputationUsageSpecForecast {
	if in == nil {
		return nil
	}
	out := new(ComputationUsageSpecForecast)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationUsageSpecGroupByTag) DeepCopyInto(out *ComputationUsageSpecGroupByTag) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationUsageSpecGroupByTag.
func (in *ComputationUsageSpecGroupByTag) DeepCopy() *ComputationUsageSpecGroupByTag {
	if in == nil {
		return nil
	}
	out := new(ComputationUsageSpecGroupByTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationUsageSpecItems) DeepCopyInto(out *ComputationUsageSpecItems) {
	*out = *in
	if in.Ad != nil {
		in, out := &in.Ad, &out.Ad
		*out = new(string)
		**out = **in
	}
	if in.CompartmentID != nil {
		in, out := &in.CompartmentID, &out.CompartmentID
		*out = new(string)
		**out = **in
	}
	if in.CompartmentName != nil {
		in, out := &in.CompartmentName, &out.CompartmentName
		*out = new(string)
		**out = **in
	}
	if in.CompartmentPath != nil {
		in, out := &in.CompartmentPath, &out.CompartmentPath
		*out = new(string)
		**out = **in
	}
	if in.ComputedAmount != nil {
		in, out := &in.ComputedAmount, &out.ComputedAmount
		*out = new(float64)
		**out = **in
	}
	if in.ComputedQuantity != nil {
		in, out := &in.ComputedQuantity, &out.ComputedQuantity
		*out = new(float64)
		**out = **in
	}
	if in.Currency != nil {
		in, out := &in.Currency, &out.Currency
		*out = new(string)
		**out = **in
	}
	if in.Discount != nil {
		in, out := &in.Discount, &out.Discount
		*out = new(float64)
		**out = **in
	}
	if in.IsForecast != nil {
		in, out := &in.IsForecast, &out.IsForecast
		*out = new(bool)
		**out = **in
	}
	if in.ListRate != nil {
		in, out := &in.ListRate, &out.ListRate
		*out = new(float64)
		**out = **in
	}
	if in.Overage != nil {
		in, out := &in.Overage, &out.Overage
		*out = new(string)
		**out = **in
	}
	if in.OveragesFlag != nil {
		in, out := &in.OveragesFlag, &out.OveragesFlag
		*out = new(string)
		**out = **in
	}
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ResourceName != nil {
		in, out := &in.ResourceName, &out.ResourceName
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Shape != nil {
		in, out := &in.Shape, &out.Shape
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.SkuPartNumber != nil {
		in, out := &in.SkuPartNumber, &out.SkuPartNumber
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionID != nil {
		in, out := &in.SubscriptionID, &out.SubscriptionID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]ComputationUsageSpecItemsTags, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.TenantName != nil {
		in, out := &in.TenantName, &out.TenantName
		*out = new(string)
		**out = **in
	}
	if in.TimeUsageEnded != nil {
		in, out := &in.TimeUsageEnded, &out.TimeUsageEnded
		*out = new(string)
		**out = **in
	}
	if in.TimeUsageStarted != nil {
		in, out := &in.TimeUsageStarted, &out.TimeUsageStarted
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	if in.UnitPrice != nil {
		in, out := &in.UnitPrice, &out.UnitPrice
		*out = new(float64)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationUsageSpecItems.
func (in *ComputationUsageSpecItems) DeepCopy() *ComputationUsageSpecItems {
	if in == nil {
		return nil
	}
	out := new(ComputationUsageSpecItems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationUsageSpecItemsTags) DeepCopyInto(out *ComputationUsageSpecItemsTags) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationUsageSpecItemsTags.
func (in *ComputationUsageSpecItemsTags) DeepCopy() *ComputationUsageSpecItemsTags {
	if in == nil {
		return nil
	}
	out := new(ComputationUsageSpecItemsTags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationUsageSpecResource) DeepCopyInto(out *ComputationUsageSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.CompartmentDepth != nil {
		in, out := &in.CompartmentDepth, &out.CompartmentDepth
		*out = new(float64)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(string)
		**out = **in
	}
	if in.Forecast != nil {
		in, out := &in.Forecast, &out.Forecast
		*out = new(ComputationUsageSpecForecast)
		(*in).DeepCopyInto(*out)
	}
	if in.Granularity != nil {
		in, out := &in.Granularity, &out.Granularity
		*out = new(string)
		**out = **in
	}
	if in.GroupBy != nil {
		in, out := &in.GroupBy, &out.GroupBy
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GroupByTag != nil {
		in, out := &in.GroupByTag, &out.GroupByTag
		*out = make([]ComputationUsageSpecGroupByTag, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsAggregateByTime != nil {
		in, out := &in.IsAggregateByTime, &out.IsAggregateByTime
		*out = new(bool)
		**out = **in
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ComputationUsageSpecItems, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.QueryType != nil {
		in, out := &in.QueryType, &out.QueryType
		*out = new(string)
		**out = **in
	}
	if in.TenantID != nil {
		in, out := &in.TenantID, &out.TenantID
		*out = new(string)
		**out = **in
	}
	if in.TimeUsageEnded != nil {
		in, out := &in.TimeUsageEnded, &out.TimeUsageEnded
		*out = new(string)
		**out = **in
	}
	if in.TimeUsageStarted != nil {
		in, out := &in.TimeUsageStarted, &out.TimeUsageStarted
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationUsageSpecResource.
func (in *ComputationUsageSpecResource) DeepCopy() *ComputationUsageSpecResource {
	if in == nil {
		return nil
	}
	out := new(ComputationUsageSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputationUsageStatus) DeepCopyInto(out *ComputationUsageStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputationUsageStatus.
func (in *ComputationUsageStatus) DeepCopy() *ComputationUsageStatus {
	if in == nil {
		return nil
	}
	out := new(ComputationUsageStatus)
	in.DeepCopyInto(out)
	return out
}
