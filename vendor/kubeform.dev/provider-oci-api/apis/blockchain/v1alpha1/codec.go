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
		jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetails{}).Type1()):                         BlockchainPlatformSpecComponentDetailsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{}).Type1()):  BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParamCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{}).Type1()): BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParamCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecReplicas{}).Type1()):                                 BlockchainPlatformSpecReplicasCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(OsnSpecOcpuAllocationParam{}).Type1()):                                     OsnSpecOcpuAllocationParamCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeerSpecOcpuAllocationParam{}).Type1()):                                    PeerSpecOcpuAllocationParamCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetails{}).Type1()):                         BlockchainPlatformSpecComponentDetailsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{}).Type1()):  BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParamCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{}).Type1()): BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParamCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecReplicas{}).Type1()):                                 BlockchainPlatformSpecReplicasCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(OsnSpecOcpuAllocationParam{}).Type1()):                                     OsnSpecOcpuAllocationParamCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(PeerSpecOcpuAllocationParam{}).Type1()):                                    PeerSpecOcpuAllocationParamCodec{},
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
type BlockchainPlatformSpecComponentDetailsCodec struct {
}

func (BlockchainPlatformSpecComponentDetailsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BlockchainPlatformSpecComponentDetails)(ptr) == nil
}

func (BlockchainPlatformSpecComponentDetailsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BlockchainPlatformSpecComponentDetails)(ptr)
	var objs []BlockchainPlatformSpecComponentDetails
	if obj != nil {
		objs = []BlockchainPlatformSpecComponentDetails{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetails{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BlockchainPlatformSpecComponentDetailsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BlockchainPlatformSpecComponentDetails)(ptr) = BlockchainPlatformSpecComponentDetails{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BlockchainPlatformSpecComponentDetails

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetails{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BlockchainPlatformSpecComponentDetails)(ptr) = objs[0]
			} else {
				*(*BlockchainPlatformSpecComponentDetails)(ptr) = BlockchainPlatformSpecComponentDetails{}
			}
		} else {
			*(*BlockchainPlatformSpecComponentDetails)(ptr) = BlockchainPlatformSpecComponentDetails{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BlockchainPlatformSpecComponentDetails

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetails{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BlockchainPlatformSpecComponentDetails)(ptr) = obj
		} else {
			*(*BlockchainPlatformSpecComponentDetails)(ptr) = BlockchainPlatformSpecComponentDetails{}
		}
	default:
		iter.ReportError("decode BlockchainPlatformSpecComponentDetails", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParamCodec struct {
}

func (BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParamCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam)(ptr) == nil
}

func (BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParamCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam)(ptr)
	var objs []BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam
	if obj != nil {
		objs = []BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParamCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam)(ptr) = BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam)(ptr) = objs[0]
			} else {
				*(*BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam)(ptr) = BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{}
			}
		} else {
			*(*BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam)(ptr) = BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam)(ptr) = obj
		} else {
			*(*BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam)(ptr) = BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam{}
		}
	default:
		iter.ReportError("decode BlockchainPlatformSpecComponentDetailsOsnsOcpuAllocationParam", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParamCodec struct {
}

func (BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParamCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam)(ptr) == nil
}

func (BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParamCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam)(ptr)
	var objs []BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam
	if obj != nil {
		objs = []BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParamCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam)(ptr) = BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam)(ptr) = objs[0]
			} else {
				*(*BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam)(ptr) = BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{}
			}
		} else {
			*(*BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam)(ptr) = BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam)(ptr) = obj
		} else {
			*(*BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam)(ptr) = BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam{}
		}
	default:
		iter.ReportError("decode BlockchainPlatformSpecComponentDetailsPeersOcpuAllocationParam", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BlockchainPlatformSpecReplicasCodec struct {
}

func (BlockchainPlatformSpecReplicasCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BlockchainPlatformSpecReplicas)(ptr) == nil
}

func (BlockchainPlatformSpecReplicasCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BlockchainPlatformSpecReplicas)(ptr)
	var objs []BlockchainPlatformSpecReplicas
	if obj != nil {
		objs = []BlockchainPlatformSpecReplicas{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecReplicas{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BlockchainPlatformSpecReplicasCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BlockchainPlatformSpecReplicas)(ptr) = BlockchainPlatformSpecReplicas{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BlockchainPlatformSpecReplicas

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecReplicas{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BlockchainPlatformSpecReplicas)(ptr) = objs[0]
			} else {
				*(*BlockchainPlatformSpecReplicas)(ptr) = BlockchainPlatformSpecReplicas{}
			}
		} else {
			*(*BlockchainPlatformSpecReplicas)(ptr) = BlockchainPlatformSpecReplicas{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BlockchainPlatformSpecReplicas

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BlockchainPlatformSpecReplicas{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BlockchainPlatformSpecReplicas)(ptr) = obj
		} else {
			*(*BlockchainPlatformSpecReplicas)(ptr) = BlockchainPlatformSpecReplicas{}
		}
	default:
		iter.ReportError("decode BlockchainPlatformSpecReplicas", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type OsnSpecOcpuAllocationParamCodec struct {
}

func (OsnSpecOcpuAllocationParamCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*OsnSpecOcpuAllocationParam)(ptr) == nil
}

func (OsnSpecOcpuAllocationParamCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*OsnSpecOcpuAllocationParam)(ptr)
	var objs []OsnSpecOcpuAllocationParam
	if obj != nil {
		objs = []OsnSpecOcpuAllocationParam{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OsnSpecOcpuAllocationParam{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (OsnSpecOcpuAllocationParamCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*OsnSpecOcpuAllocationParam)(ptr) = OsnSpecOcpuAllocationParam{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []OsnSpecOcpuAllocationParam

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OsnSpecOcpuAllocationParam{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*OsnSpecOcpuAllocationParam)(ptr) = objs[0]
			} else {
				*(*OsnSpecOcpuAllocationParam)(ptr) = OsnSpecOcpuAllocationParam{}
			}
		} else {
			*(*OsnSpecOcpuAllocationParam)(ptr) = OsnSpecOcpuAllocationParam{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj OsnSpecOcpuAllocationParam

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OsnSpecOcpuAllocationParam{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*OsnSpecOcpuAllocationParam)(ptr) = obj
		} else {
			*(*OsnSpecOcpuAllocationParam)(ptr) = OsnSpecOcpuAllocationParam{}
		}
	default:
		iter.ReportError("decode OsnSpecOcpuAllocationParam", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type PeerSpecOcpuAllocationParamCodec struct {
}

func (PeerSpecOcpuAllocationParamCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*PeerSpecOcpuAllocationParam)(ptr) == nil
}

func (PeerSpecOcpuAllocationParamCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*PeerSpecOcpuAllocationParam)(ptr)
	var objs []PeerSpecOcpuAllocationParam
	if obj != nil {
		objs = []PeerSpecOcpuAllocationParam{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeerSpecOcpuAllocationParam{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (PeerSpecOcpuAllocationParamCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*PeerSpecOcpuAllocationParam)(ptr) = PeerSpecOcpuAllocationParam{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []PeerSpecOcpuAllocationParam

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeerSpecOcpuAllocationParam{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*PeerSpecOcpuAllocationParam)(ptr) = objs[0]
			} else {
				*(*PeerSpecOcpuAllocationParam)(ptr) = PeerSpecOcpuAllocationParam{}
			}
		} else {
			*(*PeerSpecOcpuAllocationParam)(ptr) = PeerSpecOcpuAllocationParam{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj PeerSpecOcpuAllocationParam

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(PeerSpecOcpuAllocationParam{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*PeerSpecOcpuAllocationParam)(ptr) = obj
		} else {
			*(*PeerSpecOcpuAllocationParam)(ptr) = PeerSpecOcpuAllocationParam{}
		}
	default:
		iter.ReportError("decode PeerSpecOcpuAllocationParam", "unexpected JSON type")
	}
}