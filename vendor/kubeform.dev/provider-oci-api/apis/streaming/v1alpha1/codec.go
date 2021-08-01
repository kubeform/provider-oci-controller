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
		jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecCustomEncryptionKey{}).Type1()):     StreamPoolSpecCustomEncryptionKeyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecKafkaSettings{}).Type1()):           StreamPoolSpecKafkaSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecPrivateEndpointSettings{}).Type1()): StreamPoolSpecPrivateEndpointSettingsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecCustomEncryptionKey{}).Type1()):     StreamPoolSpecCustomEncryptionKeyCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecKafkaSettings{}).Type1()):           StreamPoolSpecKafkaSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecPrivateEndpointSettings{}).Type1()): StreamPoolSpecPrivateEndpointSettingsCodec{},
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
type StreamPoolSpecCustomEncryptionKeyCodec struct {
}

func (StreamPoolSpecCustomEncryptionKeyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*StreamPoolSpecCustomEncryptionKey)(ptr) == nil
}

func (StreamPoolSpecCustomEncryptionKeyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*StreamPoolSpecCustomEncryptionKey)(ptr)
	var objs []StreamPoolSpecCustomEncryptionKey
	if obj != nil {
		objs = []StreamPoolSpecCustomEncryptionKey{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecCustomEncryptionKey{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (StreamPoolSpecCustomEncryptionKeyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*StreamPoolSpecCustomEncryptionKey)(ptr) = StreamPoolSpecCustomEncryptionKey{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []StreamPoolSpecCustomEncryptionKey

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecCustomEncryptionKey{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*StreamPoolSpecCustomEncryptionKey)(ptr) = objs[0]
			} else {
				*(*StreamPoolSpecCustomEncryptionKey)(ptr) = StreamPoolSpecCustomEncryptionKey{}
			}
		} else {
			*(*StreamPoolSpecCustomEncryptionKey)(ptr) = StreamPoolSpecCustomEncryptionKey{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj StreamPoolSpecCustomEncryptionKey

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecCustomEncryptionKey{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*StreamPoolSpecCustomEncryptionKey)(ptr) = obj
		} else {
			*(*StreamPoolSpecCustomEncryptionKey)(ptr) = StreamPoolSpecCustomEncryptionKey{}
		}
	default:
		iter.ReportError("decode StreamPoolSpecCustomEncryptionKey", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type StreamPoolSpecKafkaSettingsCodec struct {
}

func (StreamPoolSpecKafkaSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*StreamPoolSpecKafkaSettings)(ptr) == nil
}

func (StreamPoolSpecKafkaSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*StreamPoolSpecKafkaSettings)(ptr)
	var objs []StreamPoolSpecKafkaSettings
	if obj != nil {
		objs = []StreamPoolSpecKafkaSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecKafkaSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (StreamPoolSpecKafkaSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*StreamPoolSpecKafkaSettings)(ptr) = StreamPoolSpecKafkaSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []StreamPoolSpecKafkaSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecKafkaSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*StreamPoolSpecKafkaSettings)(ptr) = objs[0]
			} else {
				*(*StreamPoolSpecKafkaSettings)(ptr) = StreamPoolSpecKafkaSettings{}
			}
		} else {
			*(*StreamPoolSpecKafkaSettings)(ptr) = StreamPoolSpecKafkaSettings{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj StreamPoolSpecKafkaSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecKafkaSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*StreamPoolSpecKafkaSettings)(ptr) = obj
		} else {
			*(*StreamPoolSpecKafkaSettings)(ptr) = StreamPoolSpecKafkaSettings{}
		}
	default:
		iter.ReportError("decode StreamPoolSpecKafkaSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type StreamPoolSpecPrivateEndpointSettingsCodec struct {
}

func (StreamPoolSpecPrivateEndpointSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*StreamPoolSpecPrivateEndpointSettings)(ptr) == nil
}

func (StreamPoolSpecPrivateEndpointSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*StreamPoolSpecPrivateEndpointSettings)(ptr)
	var objs []StreamPoolSpecPrivateEndpointSettings
	if obj != nil {
		objs = []StreamPoolSpecPrivateEndpointSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecPrivateEndpointSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (StreamPoolSpecPrivateEndpointSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*StreamPoolSpecPrivateEndpointSettings)(ptr) = StreamPoolSpecPrivateEndpointSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []StreamPoolSpecPrivateEndpointSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecPrivateEndpointSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*StreamPoolSpecPrivateEndpointSettings)(ptr) = objs[0]
			} else {
				*(*StreamPoolSpecPrivateEndpointSettings)(ptr) = StreamPoolSpecPrivateEndpointSettings{}
			}
		} else {
			*(*StreamPoolSpecPrivateEndpointSettings)(ptr) = StreamPoolSpecPrivateEndpointSettings{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj StreamPoolSpecPrivateEndpointSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(StreamPoolSpecPrivateEndpointSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*StreamPoolSpecPrivateEndpointSettings)(ptr) = obj
		} else {
			*(*StreamPoolSpecPrivateEndpointSettings)(ptr) = StreamPoolSpecPrivateEndpointSettings{}
		}
	default:
		iter.ReportError("decode StreamPoolSpecPrivateEndpointSettings", "unexpected JSON type")
	}
}