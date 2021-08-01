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
		jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecNetworkPolicy{}).Type1()):  AuthenticationPolicySpecNetworkPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecPasswordPolicy{}).Type1()): AuthenticationPolicySpecPasswordPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TagSpecValidator{}).Type1()):                       TagSpecValidatorCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(UserSpecCapabilities{}).Type1()):                   UserSpecCapabilitiesCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecNetworkPolicy{}).Type1()):  AuthenticationPolicySpecNetworkPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecPasswordPolicy{}).Type1()): AuthenticationPolicySpecPasswordPolicyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TagSpecValidator{}).Type1()):                       TagSpecValidatorCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(UserSpecCapabilities{}).Type1()):                   UserSpecCapabilitiesCodec{},
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
type AuthenticationPolicySpecNetworkPolicyCodec struct {
}

func (AuthenticationPolicySpecNetworkPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AuthenticationPolicySpecNetworkPolicy)(ptr) == nil
}

func (AuthenticationPolicySpecNetworkPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AuthenticationPolicySpecNetworkPolicy)(ptr)
	var objs []AuthenticationPolicySpecNetworkPolicy
	if obj != nil {
		objs = []AuthenticationPolicySpecNetworkPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecNetworkPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AuthenticationPolicySpecNetworkPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AuthenticationPolicySpecNetworkPolicy)(ptr) = AuthenticationPolicySpecNetworkPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AuthenticationPolicySpecNetworkPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecNetworkPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AuthenticationPolicySpecNetworkPolicy)(ptr) = objs[0]
			} else {
				*(*AuthenticationPolicySpecNetworkPolicy)(ptr) = AuthenticationPolicySpecNetworkPolicy{}
			}
		} else {
			*(*AuthenticationPolicySpecNetworkPolicy)(ptr) = AuthenticationPolicySpecNetworkPolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AuthenticationPolicySpecNetworkPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecNetworkPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AuthenticationPolicySpecNetworkPolicy)(ptr) = obj
		} else {
			*(*AuthenticationPolicySpecNetworkPolicy)(ptr) = AuthenticationPolicySpecNetworkPolicy{}
		}
	default:
		iter.ReportError("decode AuthenticationPolicySpecNetworkPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AuthenticationPolicySpecPasswordPolicyCodec struct {
}

func (AuthenticationPolicySpecPasswordPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AuthenticationPolicySpecPasswordPolicy)(ptr) == nil
}

func (AuthenticationPolicySpecPasswordPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AuthenticationPolicySpecPasswordPolicy)(ptr)
	var objs []AuthenticationPolicySpecPasswordPolicy
	if obj != nil {
		objs = []AuthenticationPolicySpecPasswordPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecPasswordPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AuthenticationPolicySpecPasswordPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AuthenticationPolicySpecPasswordPolicy)(ptr) = AuthenticationPolicySpecPasswordPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AuthenticationPolicySpecPasswordPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecPasswordPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AuthenticationPolicySpecPasswordPolicy)(ptr) = objs[0]
			} else {
				*(*AuthenticationPolicySpecPasswordPolicy)(ptr) = AuthenticationPolicySpecPasswordPolicy{}
			}
		} else {
			*(*AuthenticationPolicySpecPasswordPolicy)(ptr) = AuthenticationPolicySpecPasswordPolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AuthenticationPolicySpecPasswordPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AuthenticationPolicySpecPasswordPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AuthenticationPolicySpecPasswordPolicy)(ptr) = obj
		} else {
			*(*AuthenticationPolicySpecPasswordPolicy)(ptr) = AuthenticationPolicySpecPasswordPolicy{}
		}
	default:
		iter.ReportError("decode AuthenticationPolicySpecPasswordPolicy", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TagSpecValidatorCodec struct {
}

func (TagSpecValidatorCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TagSpecValidator)(ptr) == nil
}

func (TagSpecValidatorCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TagSpecValidator)(ptr)
	var objs []TagSpecValidator
	if obj != nil {
		objs = []TagSpecValidator{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TagSpecValidator{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TagSpecValidatorCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TagSpecValidator)(ptr) = TagSpecValidator{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TagSpecValidator

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TagSpecValidator{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TagSpecValidator)(ptr) = objs[0]
			} else {
				*(*TagSpecValidator)(ptr) = TagSpecValidator{}
			}
		} else {
			*(*TagSpecValidator)(ptr) = TagSpecValidator{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TagSpecValidator

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TagSpecValidator{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TagSpecValidator)(ptr) = obj
		} else {
			*(*TagSpecValidator)(ptr) = TagSpecValidator{}
		}
	default:
		iter.ReportError("decode TagSpecValidator", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type UserSpecCapabilitiesCodec struct {
}

func (UserSpecCapabilitiesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*UserSpecCapabilities)(ptr) == nil
}

func (UserSpecCapabilitiesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*UserSpecCapabilities)(ptr)
	var objs []UserSpecCapabilities
	if obj != nil {
		objs = []UserSpecCapabilities{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(UserSpecCapabilities{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (UserSpecCapabilitiesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*UserSpecCapabilities)(ptr) = UserSpecCapabilities{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []UserSpecCapabilities

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(UserSpecCapabilities{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*UserSpecCapabilities)(ptr) = objs[0]
			} else {
				*(*UserSpecCapabilities)(ptr) = UserSpecCapabilities{}
			}
		} else {
			*(*UserSpecCapabilities)(ptr) = UserSpecCapabilities{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj UserSpecCapabilities

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(UserSpecCapabilities{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*UserSpecCapabilities)(ptr) = obj
		} else {
			*(*UserSpecCapabilities)(ptr) = UserSpecCapabilities{}
		}
	default:
		iter.ReportError("decode UserSpecCapabilities", "unexpected JSON type")
	}
}
