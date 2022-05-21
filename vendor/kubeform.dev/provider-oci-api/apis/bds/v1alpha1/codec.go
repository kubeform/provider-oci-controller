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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicy{}).Type1()):                     AutoScalingConfigurationSpecPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetric{}).Type1()):          AutoScalingConfigurationSpecPolicyRulesMetricCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetricThreshold{}).Type1()): AutoScalingConfigurationSpecPolicyRulesMetricThresholdCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecCloudSQLDetails{}).Type1()):                         BdsInstanceSpecCloudSQLDetailsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecClusterDetails{}).Type1()):                          BdsInstanceSpecClusterDetailsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecMasterNode{}).Type1()):                              BdsInstanceSpecMasterNodeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecNetworkConfig{}).Type1()):                           BdsInstanceSpecNetworkConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecUtilNode{}).Type1()):                                BdsInstanceSpecUtilNodeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecWorkerNode{}).Type1()):                              BdsInstanceSpecWorkerNodeCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicy{}).Type1()):                     AutoScalingConfigurationSpecPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetric{}).Type1()):          AutoScalingConfigurationSpecPolicyRulesMetricCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetricThreshold{}).Type1()): AutoScalingConfigurationSpecPolicyRulesMetricThresholdCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecCloudSQLDetails{}).Type1()):                         BdsInstanceSpecCloudSQLDetailsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecClusterDetails{}).Type1()):                          BdsInstanceSpecClusterDetailsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecMasterNode{}).Type1()):                              BdsInstanceSpecMasterNodeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecNetworkConfig{}).Type1()):                           BdsInstanceSpecNetworkConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecUtilNode{}).Type1()):                                BdsInstanceSpecUtilNodeCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecWorkerNode{}).Type1()):                              BdsInstanceSpecWorkerNodeCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type AutoScalingConfigurationSpecPolicyCodec struct {
}

func (AutoScalingConfigurationSpecPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AutoScalingConfigurationSpecPolicy)(ptr) == nil
}

func (AutoScalingConfigurationSpecPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AutoScalingConfigurationSpecPolicy)(ptr)
	var objs []AutoScalingConfigurationSpecPolicy
	if obj != nil {
		objs = []AutoScalingConfigurationSpecPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AutoScalingConfigurationSpecPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AutoScalingConfigurationSpecPolicy)(ptr) = AutoScalingConfigurationSpecPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AutoScalingConfigurationSpecPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AutoScalingConfigurationSpecPolicy)(ptr) = objs[0]
			} else {
				*(*AutoScalingConfigurationSpecPolicy)(ptr) = AutoScalingConfigurationSpecPolicy{}
			}
		} else {
			*(*AutoScalingConfigurationSpecPolicy)(ptr) = AutoScalingConfigurationSpecPolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AutoScalingConfigurationSpecPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AutoScalingConfigurationSpecPolicy)(ptr) = obj
		} else {
			*(*AutoScalingConfigurationSpecPolicy)(ptr) = AutoScalingConfigurationSpecPolicy{}
		}
	default:
		iter.ReportError("decode AutoScalingConfigurationSpecPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AutoScalingConfigurationSpecPolicyRulesMetricCodec struct {
}

func (AutoScalingConfigurationSpecPolicyRulesMetricCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AutoScalingConfigurationSpecPolicyRulesMetric)(ptr) == nil
}

func (AutoScalingConfigurationSpecPolicyRulesMetricCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AutoScalingConfigurationSpecPolicyRulesMetric)(ptr)
	var objs []AutoScalingConfigurationSpecPolicyRulesMetric
	if obj != nil {
		objs = []AutoScalingConfigurationSpecPolicyRulesMetric{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetric{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AutoScalingConfigurationSpecPolicyRulesMetricCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AutoScalingConfigurationSpecPolicyRulesMetric)(ptr) = AutoScalingConfigurationSpecPolicyRulesMetric{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AutoScalingConfigurationSpecPolicyRulesMetric

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetric{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AutoScalingConfigurationSpecPolicyRulesMetric)(ptr) = objs[0]
			} else {
				*(*AutoScalingConfigurationSpecPolicyRulesMetric)(ptr) = AutoScalingConfigurationSpecPolicyRulesMetric{}
			}
		} else {
			*(*AutoScalingConfigurationSpecPolicyRulesMetric)(ptr) = AutoScalingConfigurationSpecPolicyRulesMetric{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AutoScalingConfigurationSpecPolicyRulesMetric

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetric{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AutoScalingConfigurationSpecPolicyRulesMetric)(ptr) = obj
		} else {
			*(*AutoScalingConfigurationSpecPolicyRulesMetric)(ptr) = AutoScalingConfigurationSpecPolicyRulesMetric{}
		}
	default:
		iter.ReportError("decode AutoScalingConfigurationSpecPolicyRulesMetric", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AutoScalingConfigurationSpecPolicyRulesMetricThresholdCodec struct {
}

func (AutoScalingConfigurationSpecPolicyRulesMetricThresholdCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AutoScalingConfigurationSpecPolicyRulesMetricThreshold)(ptr) == nil
}

func (AutoScalingConfigurationSpecPolicyRulesMetricThresholdCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AutoScalingConfigurationSpecPolicyRulesMetricThreshold)(ptr)
	var objs []AutoScalingConfigurationSpecPolicyRulesMetricThreshold
	if obj != nil {
		objs = []AutoScalingConfigurationSpecPolicyRulesMetricThreshold{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetricThreshold{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AutoScalingConfigurationSpecPolicyRulesMetricThresholdCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AutoScalingConfigurationSpecPolicyRulesMetricThreshold)(ptr) = AutoScalingConfigurationSpecPolicyRulesMetricThreshold{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AutoScalingConfigurationSpecPolicyRulesMetricThreshold

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetricThreshold{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AutoScalingConfigurationSpecPolicyRulesMetricThreshold)(ptr) = objs[0]
			} else {
				*(*AutoScalingConfigurationSpecPolicyRulesMetricThreshold)(ptr) = AutoScalingConfigurationSpecPolicyRulesMetricThreshold{}
			}
		} else {
			*(*AutoScalingConfigurationSpecPolicyRulesMetricThreshold)(ptr) = AutoScalingConfigurationSpecPolicyRulesMetricThreshold{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AutoScalingConfigurationSpecPolicyRulesMetricThreshold

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AutoScalingConfigurationSpecPolicyRulesMetricThreshold{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AutoScalingConfigurationSpecPolicyRulesMetricThreshold)(ptr) = obj
		} else {
			*(*AutoScalingConfigurationSpecPolicyRulesMetricThreshold)(ptr) = AutoScalingConfigurationSpecPolicyRulesMetricThreshold{}
		}
	default:
		iter.ReportError("decode AutoScalingConfigurationSpecPolicyRulesMetricThreshold", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BdsInstanceSpecCloudSQLDetailsCodec struct {
}

func (BdsInstanceSpecCloudSQLDetailsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BdsInstanceSpecCloudSQLDetails)(ptr) == nil
}

func (BdsInstanceSpecCloudSQLDetailsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BdsInstanceSpecCloudSQLDetails)(ptr)
	var objs []BdsInstanceSpecCloudSQLDetails
	if obj != nil {
		objs = []BdsInstanceSpecCloudSQLDetails{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecCloudSQLDetails{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BdsInstanceSpecCloudSQLDetailsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BdsInstanceSpecCloudSQLDetails)(ptr) = BdsInstanceSpecCloudSQLDetails{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BdsInstanceSpecCloudSQLDetails

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecCloudSQLDetails{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BdsInstanceSpecCloudSQLDetails)(ptr) = objs[0]
			} else {
				*(*BdsInstanceSpecCloudSQLDetails)(ptr) = BdsInstanceSpecCloudSQLDetails{}
			}
		} else {
			*(*BdsInstanceSpecCloudSQLDetails)(ptr) = BdsInstanceSpecCloudSQLDetails{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BdsInstanceSpecCloudSQLDetails

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecCloudSQLDetails{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BdsInstanceSpecCloudSQLDetails)(ptr) = obj
		} else {
			*(*BdsInstanceSpecCloudSQLDetails)(ptr) = BdsInstanceSpecCloudSQLDetails{}
		}
	default:
		iter.ReportError("decode BdsInstanceSpecCloudSQLDetails", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BdsInstanceSpecClusterDetailsCodec struct {
}

func (BdsInstanceSpecClusterDetailsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BdsInstanceSpecClusterDetails)(ptr) == nil
}

func (BdsInstanceSpecClusterDetailsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BdsInstanceSpecClusterDetails)(ptr)
	var objs []BdsInstanceSpecClusterDetails
	if obj != nil {
		objs = []BdsInstanceSpecClusterDetails{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecClusterDetails{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BdsInstanceSpecClusterDetailsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BdsInstanceSpecClusterDetails)(ptr) = BdsInstanceSpecClusterDetails{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BdsInstanceSpecClusterDetails

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecClusterDetails{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BdsInstanceSpecClusterDetails)(ptr) = objs[0]
			} else {
				*(*BdsInstanceSpecClusterDetails)(ptr) = BdsInstanceSpecClusterDetails{}
			}
		} else {
			*(*BdsInstanceSpecClusterDetails)(ptr) = BdsInstanceSpecClusterDetails{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BdsInstanceSpecClusterDetails

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecClusterDetails{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BdsInstanceSpecClusterDetails)(ptr) = obj
		} else {
			*(*BdsInstanceSpecClusterDetails)(ptr) = BdsInstanceSpecClusterDetails{}
		}
	default:
		iter.ReportError("decode BdsInstanceSpecClusterDetails", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BdsInstanceSpecMasterNodeCodec struct {
}

func (BdsInstanceSpecMasterNodeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BdsInstanceSpecMasterNode)(ptr) == nil
}

func (BdsInstanceSpecMasterNodeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BdsInstanceSpecMasterNode)(ptr)
	var objs []BdsInstanceSpecMasterNode
	if obj != nil {
		objs = []BdsInstanceSpecMasterNode{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecMasterNode{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BdsInstanceSpecMasterNodeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BdsInstanceSpecMasterNode)(ptr) = BdsInstanceSpecMasterNode{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BdsInstanceSpecMasterNode

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecMasterNode{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BdsInstanceSpecMasterNode)(ptr) = objs[0]
			} else {
				*(*BdsInstanceSpecMasterNode)(ptr) = BdsInstanceSpecMasterNode{}
			}
		} else {
			*(*BdsInstanceSpecMasterNode)(ptr) = BdsInstanceSpecMasterNode{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BdsInstanceSpecMasterNode

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecMasterNode{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BdsInstanceSpecMasterNode)(ptr) = obj
		} else {
			*(*BdsInstanceSpecMasterNode)(ptr) = BdsInstanceSpecMasterNode{}
		}
	default:
		iter.ReportError("decode BdsInstanceSpecMasterNode", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BdsInstanceSpecNetworkConfigCodec struct {
}

func (BdsInstanceSpecNetworkConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BdsInstanceSpecNetworkConfig)(ptr) == nil
}

func (BdsInstanceSpecNetworkConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BdsInstanceSpecNetworkConfig)(ptr)
	var objs []BdsInstanceSpecNetworkConfig
	if obj != nil {
		objs = []BdsInstanceSpecNetworkConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecNetworkConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BdsInstanceSpecNetworkConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BdsInstanceSpecNetworkConfig)(ptr) = BdsInstanceSpecNetworkConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BdsInstanceSpecNetworkConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecNetworkConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BdsInstanceSpecNetworkConfig)(ptr) = objs[0]
			} else {
				*(*BdsInstanceSpecNetworkConfig)(ptr) = BdsInstanceSpecNetworkConfig{}
			}
		} else {
			*(*BdsInstanceSpecNetworkConfig)(ptr) = BdsInstanceSpecNetworkConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BdsInstanceSpecNetworkConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecNetworkConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BdsInstanceSpecNetworkConfig)(ptr) = obj
		} else {
			*(*BdsInstanceSpecNetworkConfig)(ptr) = BdsInstanceSpecNetworkConfig{}
		}
	default:
		iter.ReportError("decode BdsInstanceSpecNetworkConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BdsInstanceSpecUtilNodeCodec struct {
}

func (BdsInstanceSpecUtilNodeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BdsInstanceSpecUtilNode)(ptr) == nil
}

func (BdsInstanceSpecUtilNodeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BdsInstanceSpecUtilNode)(ptr)
	var objs []BdsInstanceSpecUtilNode
	if obj != nil {
		objs = []BdsInstanceSpecUtilNode{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecUtilNode{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BdsInstanceSpecUtilNodeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BdsInstanceSpecUtilNode)(ptr) = BdsInstanceSpecUtilNode{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BdsInstanceSpecUtilNode

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecUtilNode{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BdsInstanceSpecUtilNode)(ptr) = objs[0]
			} else {
				*(*BdsInstanceSpecUtilNode)(ptr) = BdsInstanceSpecUtilNode{}
			}
		} else {
			*(*BdsInstanceSpecUtilNode)(ptr) = BdsInstanceSpecUtilNode{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BdsInstanceSpecUtilNode

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecUtilNode{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BdsInstanceSpecUtilNode)(ptr) = obj
		} else {
			*(*BdsInstanceSpecUtilNode)(ptr) = BdsInstanceSpecUtilNode{}
		}
	default:
		iter.ReportError("decode BdsInstanceSpecUtilNode", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BdsInstanceSpecWorkerNodeCodec struct {
}

func (BdsInstanceSpecWorkerNodeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BdsInstanceSpecWorkerNode)(ptr) == nil
}

func (BdsInstanceSpecWorkerNodeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BdsInstanceSpecWorkerNode)(ptr)
	var objs []BdsInstanceSpecWorkerNode
	if obj != nil {
		objs = []BdsInstanceSpecWorkerNode{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecWorkerNode{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BdsInstanceSpecWorkerNodeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BdsInstanceSpecWorkerNode)(ptr) = BdsInstanceSpecWorkerNode{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BdsInstanceSpecWorkerNode

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecWorkerNode{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BdsInstanceSpecWorkerNode)(ptr) = objs[0]
			} else {
				*(*BdsInstanceSpecWorkerNode)(ptr) = BdsInstanceSpecWorkerNode{}
			}
		} else {
			*(*BdsInstanceSpecWorkerNode)(ptr) = BdsInstanceSpecWorkerNode{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BdsInstanceSpecWorkerNode

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BdsInstanceSpecWorkerNode{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BdsInstanceSpecWorkerNode)(ptr) = obj
		} else {
			*(*BdsInstanceSpecWorkerNode)(ptr) = BdsInstanceSpecWorkerNode{}
		}
	default:
		iter.ReportError("decode BdsInstanceSpecWorkerNode", "unexpected JSON type")
	}
}