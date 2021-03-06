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
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecBillingAddress{}).Type1()):                                GatewaySubscriptionSpecBillingAddressCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGateway{}).Type1()):                                GatewaySubscriptionSpecPaymentGatewayCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{}).Type1()):             GatewaySubscriptionSpecPaymentGatewayMerchantDefinedDataCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscription{}).Type1()):                                  GatewaySubscriptionSpecSubscriptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionBillingAddress{}).Type1()):                    GatewaySubscriptionSpecSubscriptionBillingAddressCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGateway{}).Type1()):                    GatewaySubscriptionSpecSubscriptionPaymentGatewayCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{}).Type1()): GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedDataCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionTaxInfo{}).Type1()):                           GatewaySubscriptionSpecSubscriptionTaxInfoCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecTaxInfo{}).Type1()):                                       GatewaySubscriptionSpecTaxInfoCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecBillingAddress{}).Type1()):                                GatewaySubscriptionSpecBillingAddressCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGateway{}).Type1()):                                GatewaySubscriptionSpecPaymentGatewayCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{}).Type1()):             GatewaySubscriptionSpecPaymentGatewayMerchantDefinedDataCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscription{}).Type1()):                                  GatewaySubscriptionSpecSubscriptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionBillingAddress{}).Type1()):                    GatewaySubscriptionSpecSubscriptionBillingAddressCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGateway{}).Type1()):                    GatewaySubscriptionSpecSubscriptionPaymentGatewayCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{}).Type1()): GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedDataCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionTaxInfo{}).Type1()):                           GatewaySubscriptionSpecSubscriptionTaxInfoCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecTaxInfo{}).Type1()):                                       GatewaySubscriptionSpecTaxInfoCodec{},
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
type GatewaySubscriptionSpecBillingAddressCodec struct {
}

func (GatewaySubscriptionSpecBillingAddressCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySubscriptionSpecBillingAddress)(ptr) == nil
}

func (GatewaySubscriptionSpecBillingAddressCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySubscriptionSpecBillingAddress)(ptr)
	var objs []GatewaySubscriptionSpecBillingAddress
	if obj != nil {
		objs = []GatewaySubscriptionSpecBillingAddress{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecBillingAddress{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySubscriptionSpecBillingAddressCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySubscriptionSpecBillingAddress)(ptr) = GatewaySubscriptionSpecBillingAddress{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySubscriptionSpecBillingAddress

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecBillingAddress{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySubscriptionSpecBillingAddress)(ptr) = objs[0]
			} else {
				*(*GatewaySubscriptionSpecBillingAddress)(ptr) = GatewaySubscriptionSpecBillingAddress{}
			}
		} else {
			*(*GatewaySubscriptionSpecBillingAddress)(ptr) = GatewaySubscriptionSpecBillingAddress{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySubscriptionSpecBillingAddress

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecBillingAddress{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySubscriptionSpecBillingAddress)(ptr) = obj
		} else {
			*(*GatewaySubscriptionSpecBillingAddress)(ptr) = GatewaySubscriptionSpecBillingAddress{}
		}
	default:
		iter.ReportError("decode GatewaySubscriptionSpecBillingAddress", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySubscriptionSpecPaymentGatewayCodec struct {
}

func (GatewaySubscriptionSpecPaymentGatewayCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySubscriptionSpecPaymentGateway)(ptr) == nil
}

func (GatewaySubscriptionSpecPaymentGatewayCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySubscriptionSpecPaymentGateway)(ptr)
	var objs []GatewaySubscriptionSpecPaymentGateway
	if obj != nil {
		objs = []GatewaySubscriptionSpecPaymentGateway{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGateway{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySubscriptionSpecPaymentGatewayCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySubscriptionSpecPaymentGateway)(ptr) = GatewaySubscriptionSpecPaymentGateway{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySubscriptionSpecPaymentGateway

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGateway{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySubscriptionSpecPaymentGateway)(ptr) = objs[0]
			} else {
				*(*GatewaySubscriptionSpecPaymentGateway)(ptr) = GatewaySubscriptionSpecPaymentGateway{}
			}
		} else {
			*(*GatewaySubscriptionSpecPaymentGateway)(ptr) = GatewaySubscriptionSpecPaymentGateway{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySubscriptionSpecPaymentGateway

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGateway{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySubscriptionSpecPaymentGateway)(ptr) = obj
		} else {
			*(*GatewaySubscriptionSpecPaymentGateway)(ptr) = GatewaySubscriptionSpecPaymentGateway{}
		}
	default:
		iter.ReportError("decode GatewaySubscriptionSpecPaymentGateway", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySubscriptionSpecPaymentGatewayMerchantDefinedDataCodec struct {
}

func (GatewaySubscriptionSpecPaymentGatewayMerchantDefinedDataCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData)(ptr) == nil
}

func (GatewaySubscriptionSpecPaymentGatewayMerchantDefinedDataCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData)(ptr)
	var objs []GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData
	if obj != nil {
		objs = []GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySubscriptionSpecPaymentGatewayMerchantDefinedDataCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData)(ptr) = GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData)(ptr) = objs[0]
			} else {
				*(*GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData)(ptr) = GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{}
			}
		} else {
			*(*GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData)(ptr) = GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData)(ptr) = obj
		} else {
			*(*GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData)(ptr) = GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData{}
		}
	default:
		iter.ReportError("decode GatewaySubscriptionSpecPaymentGatewayMerchantDefinedData", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySubscriptionSpecSubscriptionCodec struct {
}

func (GatewaySubscriptionSpecSubscriptionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySubscriptionSpecSubscription)(ptr) == nil
}

func (GatewaySubscriptionSpecSubscriptionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySubscriptionSpecSubscription)(ptr)
	var objs []GatewaySubscriptionSpecSubscription
	if obj != nil {
		objs = []GatewaySubscriptionSpecSubscription{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscription{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySubscriptionSpecSubscriptionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySubscriptionSpecSubscription)(ptr) = GatewaySubscriptionSpecSubscription{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySubscriptionSpecSubscription

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscription{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySubscriptionSpecSubscription)(ptr) = objs[0]
			} else {
				*(*GatewaySubscriptionSpecSubscription)(ptr) = GatewaySubscriptionSpecSubscription{}
			}
		} else {
			*(*GatewaySubscriptionSpecSubscription)(ptr) = GatewaySubscriptionSpecSubscription{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySubscriptionSpecSubscription

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscription{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySubscriptionSpecSubscription)(ptr) = obj
		} else {
			*(*GatewaySubscriptionSpecSubscription)(ptr) = GatewaySubscriptionSpecSubscription{}
		}
	default:
		iter.ReportError("decode GatewaySubscriptionSpecSubscription", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySubscriptionSpecSubscriptionBillingAddressCodec struct {
}

func (GatewaySubscriptionSpecSubscriptionBillingAddressCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySubscriptionSpecSubscriptionBillingAddress)(ptr) == nil
}

func (GatewaySubscriptionSpecSubscriptionBillingAddressCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySubscriptionSpecSubscriptionBillingAddress)(ptr)
	var objs []GatewaySubscriptionSpecSubscriptionBillingAddress
	if obj != nil {
		objs = []GatewaySubscriptionSpecSubscriptionBillingAddress{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionBillingAddress{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySubscriptionSpecSubscriptionBillingAddressCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySubscriptionSpecSubscriptionBillingAddress)(ptr) = GatewaySubscriptionSpecSubscriptionBillingAddress{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySubscriptionSpecSubscriptionBillingAddress

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionBillingAddress{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySubscriptionSpecSubscriptionBillingAddress)(ptr) = objs[0]
			} else {
				*(*GatewaySubscriptionSpecSubscriptionBillingAddress)(ptr) = GatewaySubscriptionSpecSubscriptionBillingAddress{}
			}
		} else {
			*(*GatewaySubscriptionSpecSubscriptionBillingAddress)(ptr) = GatewaySubscriptionSpecSubscriptionBillingAddress{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySubscriptionSpecSubscriptionBillingAddress

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionBillingAddress{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySubscriptionSpecSubscriptionBillingAddress)(ptr) = obj
		} else {
			*(*GatewaySubscriptionSpecSubscriptionBillingAddress)(ptr) = GatewaySubscriptionSpecSubscriptionBillingAddress{}
		}
	default:
		iter.ReportError("decode GatewaySubscriptionSpecSubscriptionBillingAddress", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySubscriptionSpecSubscriptionPaymentGatewayCodec struct {
}

func (GatewaySubscriptionSpecSubscriptionPaymentGatewayCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySubscriptionSpecSubscriptionPaymentGateway)(ptr) == nil
}

func (GatewaySubscriptionSpecSubscriptionPaymentGatewayCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySubscriptionSpecSubscriptionPaymentGateway)(ptr)
	var objs []GatewaySubscriptionSpecSubscriptionPaymentGateway
	if obj != nil {
		objs = []GatewaySubscriptionSpecSubscriptionPaymentGateway{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGateway{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySubscriptionSpecSubscriptionPaymentGatewayCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySubscriptionSpecSubscriptionPaymentGateway)(ptr) = GatewaySubscriptionSpecSubscriptionPaymentGateway{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySubscriptionSpecSubscriptionPaymentGateway

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGateway{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySubscriptionSpecSubscriptionPaymentGateway)(ptr) = objs[0]
			} else {
				*(*GatewaySubscriptionSpecSubscriptionPaymentGateway)(ptr) = GatewaySubscriptionSpecSubscriptionPaymentGateway{}
			}
		} else {
			*(*GatewaySubscriptionSpecSubscriptionPaymentGateway)(ptr) = GatewaySubscriptionSpecSubscriptionPaymentGateway{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySubscriptionSpecSubscriptionPaymentGateway

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGateway{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySubscriptionSpecSubscriptionPaymentGateway)(ptr) = obj
		} else {
			*(*GatewaySubscriptionSpecSubscriptionPaymentGateway)(ptr) = GatewaySubscriptionSpecSubscriptionPaymentGateway{}
		}
	default:
		iter.ReportError("decode GatewaySubscriptionSpecSubscriptionPaymentGateway", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedDataCodec struct {
}

func (GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedDataCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData)(ptr) == nil
}

func (GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedDataCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData)(ptr)
	var objs []GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData
	if obj != nil {
		objs = []GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedDataCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData)(ptr) = GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData)(ptr) = objs[0]
			} else {
				*(*GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData)(ptr) = GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{}
			}
		} else {
			*(*GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData)(ptr) = GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData)(ptr) = obj
		} else {
			*(*GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData)(ptr) = GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData{}
		}
	default:
		iter.ReportError("decode GatewaySubscriptionSpecSubscriptionPaymentGatewayMerchantDefinedData", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySubscriptionSpecSubscriptionTaxInfoCodec struct {
}

func (GatewaySubscriptionSpecSubscriptionTaxInfoCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySubscriptionSpecSubscriptionTaxInfo)(ptr) == nil
}

func (GatewaySubscriptionSpecSubscriptionTaxInfoCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySubscriptionSpecSubscriptionTaxInfo)(ptr)
	var objs []GatewaySubscriptionSpecSubscriptionTaxInfo
	if obj != nil {
		objs = []GatewaySubscriptionSpecSubscriptionTaxInfo{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionTaxInfo{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySubscriptionSpecSubscriptionTaxInfoCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySubscriptionSpecSubscriptionTaxInfo)(ptr) = GatewaySubscriptionSpecSubscriptionTaxInfo{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySubscriptionSpecSubscriptionTaxInfo

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionTaxInfo{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySubscriptionSpecSubscriptionTaxInfo)(ptr) = objs[0]
			} else {
				*(*GatewaySubscriptionSpecSubscriptionTaxInfo)(ptr) = GatewaySubscriptionSpecSubscriptionTaxInfo{}
			}
		} else {
			*(*GatewaySubscriptionSpecSubscriptionTaxInfo)(ptr) = GatewaySubscriptionSpecSubscriptionTaxInfo{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySubscriptionSpecSubscriptionTaxInfo

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecSubscriptionTaxInfo{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySubscriptionSpecSubscriptionTaxInfo)(ptr) = obj
		} else {
			*(*GatewaySubscriptionSpecSubscriptionTaxInfo)(ptr) = GatewaySubscriptionSpecSubscriptionTaxInfo{}
		}
	default:
		iter.ReportError("decode GatewaySubscriptionSpecSubscriptionTaxInfo", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type GatewaySubscriptionSpecTaxInfoCodec struct {
}

func (GatewaySubscriptionSpecTaxInfoCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*GatewaySubscriptionSpecTaxInfo)(ptr) == nil
}

func (GatewaySubscriptionSpecTaxInfoCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*GatewaySubscriptionSpecTaxInfo)(ptr)
	var objs []GatewaySubscriptionSpecTaxInfo
	if obj != nil {
		objs = []GatewaySubscriptionSpecTaxInfo{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecTaxInfo{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (GatewaySubscriptionSpecTaxInfoCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*GatewaySubscriptionSpecTaxInfo)(ptr) = GatewaySubscriptionSpecTaxInfo{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []GatewaySubscriptionSpecTaxInfo

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecTaxInfo{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*GatewaySubscriptionSpecTaxInfo)(ptr) = objs[0]
			} else {
				*(*GatewaySubscriptionSpecTaxInfo)(ptr) = GatewaySubscriptionSpecTaxInfo{}
			}
		} else {
			*(*GatewaySubscriptionSpecTaxInfo)(ptr) = GatewaySubscriptionSpecTaxInfo{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj GatewaySubscriptionSpecTaxInfo

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(GatewaySubscriptionSpecTaxInfo{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*GatewaySubscriptionSpecTaxInfo)(ptr) = obj
		} else {
			*(*GatewaySubscriptionSpecTaxInfo)(ptr) = GatewaySubscriptionSpecTaxInfo{}
		}
	default:
		iter.ReportError("decode GatewaySubscriptionSpecTaxInfo", "unexpected JSON type")
	}
}
