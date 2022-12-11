// Copyright 2018 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/devtools/containeranalysis/v1beta1/image/image.proto

package image

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Instructions from Dockerfile.
type Layer_Directive int32

const (
	// Default value for unsupported/missing directive.
	Layer_DIRECTIVE_UNSPECIFIED Layer_Directive = 0
	// https://docs.docker.com/engine/reference/builder/
	Layer_MAINTAINER Layer_Directive = 1
	// https://docs.docker.com/engine/reference/builder/
	Layer_RUN Layer_Directive = 2
	// https://docs.docker.com/engine/reference/builder/
	Layer_CMD Layer_Directive = 3
	// https://docs.docker.com/engine/reference/builder/
	Layer_LABEL Layer_Directive = 4
	// https://docs.docker.com/engine/reference/builder/
	Layer_EXPOSE Layer_Directive = 5
	// https://docs.docker.com/engine/reference/builder/
	Layer_ENV Layer_Directive = 6
	// https://docs.docker.com/engine/reference/builder/
	Layer_ADD Layer_Directive = 7
	// https://docs.docker.com/engine/reference/builder/
	Layer_COPY Layer_Directive = 8
	// https://docs.docker.com/engine/reference/builder/
	Layer_ENTRYPOINT Layer_Directive = 9
	// https://docs.docker.com/engine/reference/builder/
	Layer_VOLUME Layer_Directive = 10
	// https://docs.docker.com/engine/reference/builder/
	Layer_USER Layer_Directive = 11
	// https://docs.docker.com/engine/reference/builder/
	Layer_WORKDIR Layer_Directive = 12
	// https://docs.docker.com/engine/reference/builder/
	Layer_ARG Layer_Directive = 13
	// https://docs.docker.com/engine/reference/builder/
	Layer_ONBUILD Layer_Directive = 14
	// https://docs.docker.com/engine/reference/builder/
	Layer_STOPSIGNAL Layer_Directive = 15
	// https://docs.docker.com/engine/reference/builder/
	Layer_HEALTHCHECK Layer_Directive = 16
	// https://docs.docker.com/engine/reference/builder/
	Layer_SHELL Layer_Directive = 17
)

// Enum value maps for Layer_Directive.
var (
	Layer_Directive_name = map[int32]string{
		0:  "DIRECTIVE_UNSPECIFIED",
		1:  "MAINTAINER",
		2:  "RUN",
		3:  "CMD",
		4:  "LABEL",
		5:  "EXPOSE",
		6:  "ENV",
		7:  "ADD",
		8:  "COPY",
		9:  "ENTRYPOINT",
		10: "VOLUME",
		11: "USER",
		12: "WORKDIR",
		13: "ARG",
		14: "ONBUILD",
		15: "STOPSIGNAL",
		16: "HEALTHCHECK",
		17: "SHELL",
	}
	Layer_Directive_value = map[string]int32{
		"DIRECTIVE_UNSPECIFIED": 0,
		"MAINTAINER":            1,
		"RUN":                   2,
		"CMD":                   3,
		"LABEL":                 4,
		"EXPOSE":                5,
		"ENV":                   6,
		"ADD":                   7,
		"COPY":                  8,
		"ENTRYPOINT":            9,
		"VOLUME":                10,
		"USER":                  11,
		"WORKDIR":               12,
		"ARG":                   13,
		"ONBUILD":               14,
		"STOPSIGNAL":            15,
		"HEALTHCHECK":           16,
		"SHELL":                 17,
	}
)

func (x Layer_Directive) Enum() *Layer_Directive {
	p := new(Layer_Directive)
	*p = x
	return p
}

func (x Layer_Directive) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Layer_Directive) Descriptor() protoreflect.EnumDescriptor {
	return file_google_devtools_containeranalysis_v1beta1_image_image_proto_enumTypes[0].Descriptor()
}

func (Layer_Directive) Type() protoreflect.EnumType {
	return &file_google_devtools_containeranalysis_v1beta1_image_image_proto_enumTypes[0]
}

func (x Layer_Directive) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Layer_Directive.Descriptor instead.
func (Layer_Directive) EnumDescriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescGZIP(), []int{0, 0}
}

// Layer holds metadata specific to a layer of a Docker image.
type Layer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The recovered Dockerfile directive used to construct this layer.
	Directive Layer_Directive `protobuf:"varint,1,opt,name=directive,proto3,enum=grafeas.v1beta1.image.Layer_Directive" json:"directive,omitempty"`
	// The recovered arguments to the Dockerfile directive.
	Arguments string `protobuf:"bytes,2,opt,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *Layer) Reset() {
	*x = Layer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Layer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Layer) ProtoMessage() {}

func (x *Layer) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Layer.ProtoReflect.Descriptor instead.
func (*Layer) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescGZIP(), []int{0}
}

func (x *Layer) GetDirective() Layer_Directive {
	if x != nil {
		return x.Directive
	}
	return Layer_DIRECTIVE_UNSPECIFIED
}

func (x *Layer) GetArguments() string {
	if x != nil {
		return x.Arguments
	}
	return ""
}

// A set of properties that uniquely identify a given Docker image.
type Fingerprint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The layer ID of the final layer in the Docker image's v1
	// representation.
	V1Name string `protobuf:"bytes,1,opt,name=v1_name,json=v1Name,proto3" json:"v1_name,omitempty"`
	// Required. The ordered list of v2 blobs that represent a given image.
	V2Blob []string `protobuf:"bytes,2,rep,name=v2_blob,json=v2Blob,proto3" json:"v2_blob,omitempty"`
	// Output only. The name of the image's v2 blobs computed via:
	//   [bottom] := v2_blob[bottom]
	//   [N] := sha256(v2_blob[N] + " " + v2_name[N+1])
	// Only the name of the final blob is kept.
	V2Name string `protobuf:"bytes,3,opt,name=v2_name,json=v2Name,proto3" json:"v2_name,omitempty"`
}

func (x *Fingerprint) Reset() {
	*x = Fingerprint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fingerprint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fingerprint) ProtoMessage() {}

func (x *Fingerprint) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fingerprint.ProtoReflect.Descriptor instead.
func (*Fingerprint) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescGZIP(), []int{1}
}

func (x *Fingerprint) GetV1Name() string {
	if x != nil {
		return x.V1Name
	}
	return ""
}

func (x *Fingerprint) GetV2Blob() []string {
	if x != nil {
		return x.V2Blob
	}
	return nil
}

func (x *Fingerprint) GetV2Name() string {
	if x != nil {
		return x.V2Name
	}
	return ""
}

// Basis describes the base image portion (Note) of the DockerImage
// relationship. Linked occurrences are derived from this or an
// equivalent image via:
//   FROM <Basis.resource_url>
// Or an equivalent reference, e.g. a tag of the resource_url.
type Basis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The resource_url for the resource representing the
	// basis of associated occurrence images.
	ResourceUrl string `protobuf:"bytes,1,opt,name=resource_url,json=resourceUrl,proto3" json:"resource_url,omitempty"`
	// Required. Immutable. The fingerprint of the base image.
	Fingerprint *Fingerprint `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
}

func (x *Basis) Reset() {
	*x = Basis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Basis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Basis) ProtoMessage() {}

func (x *Basis) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Basis.ProtoReflect.Descriptor instead.
func (*Basis) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescGZIP(), []int{2}
}

func (x *Basis) GetResourceUrl() string {
	if x != nil {
		return x.ResourceUrl
	}
	return ""
}

func (x *Basis) GetFingerprint() *Fingerprint {
	if x != nil {
		return x.Fingerprint
	}
	return nil
}

// Details of an image occurrence.
type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Immutable. The child image derived from the base image.
	DerivedImage *Derived `protobuf:"bytes,1,opt,name=derived_image,json=derivedImage,proto3" json:"derived_image,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescGZIP(), []int{3}
}

func (x *Details) GetDerivedImage() *Derived {
	if x != nil {
		return x.DerivedImage
	}
	return nil
}

// Derived describes the derived image portion (Occurrence) of the DockerImage
// relationship. This image would be produced from a Dockerfile with FROM
// <DockerImage.Basis in attached Note>.
type Derived struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The fingerprint of the derived image.
	Fingerprint *Fingerprint `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Output only. The number of layers by which this image differs from the
	// associated image basis.
	Distance int32 `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
	// This contains layer-specific metadata, if populated it has length
	// "distance" and is ordered with [distance] being the layer immediately
	// following the base image and [1] being the final layer.
	LayerInfo []*Layer `protobuf:"bytes,3,rep,name=layer_info,json=layerInfo,proto3" json:"layer_info,omitempty"`
	// Output only. This contains the base image URL for the derived image
	// occurrence.
	BaseResourceUrl string `protobuf:"bytes,4,opt,name=base_resource_url,json=baseResourceUrl,proto3" json:"base_resource_url,omitempty"`
}

func (x *Derived) Reset() {
	*x = Derived{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Derived) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Derived) ProtoMessage() {}

func (x *Derived) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Derived.ProtoReflect.Descriptor instead.
func (*Derived) Descriptor() ([]byte, []int) {
	return file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescGZIP(), []int{4}
}

func (x *Derived) GetFingerprint() *Fingerprint {
	if x != nil {
		return x.Fingerprint
	}
	return nil
}

func (x *Derived) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *Derived) GetLayerInfo() []*Layer {
	if x != nil {
		return x.LayerInfo
	}
	return nil
}

func (x *Derived) GetBaseResourceUrl() string {
	if x != nil {
		return x.BaseResourceUrl
	}
	return ""
}

var File_google_devtools_containeranalysis_v1beta1_image_image_proto protoreflect.FileDescriptor

var file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x22, 0xde, 0x02, 0x0a, 0x05, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x44,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xf0, 0x01, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x19, 0x0a, 0x15, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4d,
	0x41, 0x49, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x52,
	0x55, 0x4e, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x4d, 0x44, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x50, 0x4f,
	0x53, 0x45, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4e, 0x56, 0x10, 0x06, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x44, 0x44, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x4f, 0x50, 0x59, 0x10, 0x08,
	0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10, 0x09,
	0x12, 0x0a, 0x0a, 0x06, 0x56, 0x4f, 0x4c, 0x55, 0x4d, 0x45, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04,
	0x55, 0x53, 0x45, 0x52, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x4f, 0x52, 0x4b, 0x44, 0x49,
	0x52, 0x10, 0x0c, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x52, 0x47, 0x10, 0x0d, 0x12, 0x0b, 0x0a, 0x07,
	0x4f, 0x4e, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x10, 0x0e, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x4f,
	0x50, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x4c, 0x10, 0x0f, 0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x45, 0x41,
	0x4c, 0x54, 0x48, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x10, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x48,
	0x45, 0x4c, 0x4c, 0x10, 0x11, 0x22, 0x58, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x31, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x31, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x76, 0x32, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x76, 0x32, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x32, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x32, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x70, 0x0a, 0x05, 0x42, 0x61, 0x73, 0x69, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x44, 0x0a, 0x0b, 0x66,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x22, 0x4e, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x43, 0x0a, 0x0d,
	0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x72, 0x69,
	0x76, 0x65, 0x64, 0x52, 0x0c, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0xd4, 0x01, 0x0a, 0x07, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x12, 0x44, 0x0a,
	0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x3b, 0x0a, 0x0a, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x4c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x09, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x11,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x72, 0x6c, 0x42, 0x78, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x67,
	0x72, 0x61, 0x66, 0x65, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x50, 0x01, 0x5a, 0x54, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x3b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0xa2, 0x02, 0x03, 0x47,
	0x52, 0x41, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescOnce sync.Once
	file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescData = file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDesc
)

func file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescGZIP() []byte {
	file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescOnce.Do(func() {
		file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescData)
	})
	return file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDescData
}

var file_google_devtools_containeranalysis_v1beta1_image_image_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_devtools_containeranalysis_v1beta1_image_image_proto_goTypes = []interface{}{
	(Layer_Directive)(0), // 0: grafeas.v1beta1.image.Layer.Directive
	(*Layer)(nil),        // 1: grafeas.v1beta1.image.Layer
	(*Fingerprint)(nil),  // 2: grafeas.v1beta1.image.Fingerprint
	(*Basis)(nil),        // 3: grafeas.v1beta1.image.Basis
	(*Details)(nil),      // 4: grafeas.v1beta1.image.Details
	(*Derived)(nil),      // 5: grafeas.v1beta1.image.Derived
}
var file_google_devtools_containeranalysis_v1beta1_image_image_proto_depIdxs = []int32{
	0, // 0: grafeas.v1beta1.image.Layer.directive:type_name -> grafeas.v1beta1.image.Layer.Directive
	2, // 1: grafeas.v1beta1.image.Basis.fingerprint:type_name -> grafeas.v1beta1.image.Fingerprint
	5, // 2: grafeas.v1beta1.image.Details.derived_image:type_name -> grafeas.v1beta1.image.Derived
	2, // 3: grafeas.v1beta1.image.Derived.fingerprint:type_name -> grafeas.v1beta1.image.Fingerprint
	1, // 4: grafeas.v1beta1.image.Derived.layer_info:type_name -> grafeas.v1beta1.image.Layer
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_devtools_containeranalysis_v1beta1_image_image_proto_init() }
func file_google_devtools_containeranalysis_v1beta1_image_image_proto_init() {
	if File_google_devtools_containeranalysis_v1beta1_image_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Layer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fingerprint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Basis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Derived); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_containeranalysis_v1beta1_image_image_proto_goTypes,
		DependencyIndexes: file_google_devtools_containeranalysis_v1beta1_image_image_proto_depIdxs,
		EnumInfos:         file_google_devtools_containeranalysis_v1beta1_image_image_proto_enumTypes,
		MessageInfos:      file_google_devtools_containeranalysis_v1beta1_image_image_proto_msgTypes,
	}.Build()
	File_google_devtools_containeranalysis_v1beta1_image_image_proto = out.File
	file_google_devtools_containeranalysis_v1beta1_image_image_proto_rawDesc = nil
	file_google_devtools_containeranalysis_v1beta1_image_image_proto_goTypes = nil
	file_google_devtools_containeranalysis_v1beta1_image_image_proto_depIdxs = nil
}
