// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: services/chipmunk/api/proto/src/resolution.proto

package proto

import (
	proto "github.com/h-varmazyar/Gate/api/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Resolution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"platform"
	Platform proto.Platform `protobuf:"varint,1,opt,name=Platform,proto3,enum=api.Platform" json:"platform"`
	//	@inject_tag:	json:"duration"
	Duration int64 `protobuf:"varint,2,opt,name=Duration,proto3" json:"duration"`
	//	@inject_tag:	json:"label"
	Label string `protobuf:"bytes,3,opt,name=Label,proto3" json:"label"`
	//	@inject_tag:	json:"value"
	Value string `protobuf:"bytes,4,opt,name=Value,proto3" json:"value"`
	//	@inject_tag:	json:"id"
	ID string `protobuf:"bytes,5,opt,name=ID,proto3" json:"id"`
	//	@inject_tag:	json:"created_at"
	CreatedAt int64 `protobuf:"varint,6,opt,name=CreatedAt,proto3" json:"created_at"`
	//	@inject_tag:	json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,7,opt,name=UpdatedAt,proto3" json:"updated_at"`
}

func (x *Resolution) Reset() {
	*x = Resolution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resolution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resolution) ProtoMessage() {}

func (x *Resolution) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resolution.ProtoReflect.Descriptor instead.
func (*Resolution) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_resolution_proto_rawDescGZIP(), []int{0}
}

func (x *Resolution) GetPlatform() proto.Platform {
	if x != nil {
		return x.Platform
	}
	return proto.Platform(0)
}

func (x *Resolution) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Resolution) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Resolution) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Resolution) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Resolution) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Resolution) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Resolutions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"elements"
	Elements []*Resolution `protobuf:"bytes,1,rep,name=Elements,proto3" json:"elements"`
}

func (x *Resolutions) Reset() {
	*x = Resolutions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resolutions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resolutions) ProtoMessage() {}

func (x *Resolutions) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resolutions.ProtoReflect.Descriptor instead.
func (*Resolutions) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_resolution_proto_rawDescGZIP(), []int{1}
}

func (x *Resolutions) GetElements() []*Resolution {
	if x != nil {
		return x.Elements
	}
	return nil
}

type ResolutionReturnByIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
}

func (x *ResolutionReturnByIDReq) Reset() {
	*x = ResolutionReturnByIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolutionReturnByIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolutionReturnByIDReq) ProtoMessage() {}

func (x *ResolutionReturnByIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolutionReturnByIDReq.ProtoReflect.Descriptor instead.
func (*ResolutionReturnByIDReq) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_resolution_proto_rawDescGZIP(), []int{2}
}

func (x *ResolutionReturnByIDReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ResolutionReturnByDurationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"platform"
	Platform proto.Platform `protobuf:"varint,1,opt,name=Platform,proto3,enum=api.Platform" json:"platform"`
	//	@inject_tag:	json:"duration"
	Duration int64 `protobuf:"varint,2,opt,name=Duration,proto3" json:"duration"`
}

func (x *ResolutionReturnByDurationReq) Reset() {
	*x = ResolutionReturnByDurationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolutionReturnByDurationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolutionReturnByDurationReq) ProtoMessage() {}

func (x *ResolutionReturnByDurationReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolutionReturnByDurationReq.ProtoReflect.Descriptor instead.
func (*ResolutionReturnByDurationReq) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_resolution_proto_rawDescGZIP(), []int{3}
}

func (x *ResolutionReturnByDurationReq) GetPlatform() proto.Platform {
	if x != nil {
		return x.Platform
	}
	return proto.Platform(0)
}

func (x *ResolutionReturnByDurationReq) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type ResolutionListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"platform"
	Platform proto.Platform `protobuf:"varint,1,opt,name=Platform,proto3,enum=api.Platform" json:"platform"`
}

func (x *ResolutionListReq) Reset() {
	*x = ResolutionListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolutionListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolutionListReq) ProtoMessage() {}

func (x *ResolutionListReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolutionListReq.ProtoReflect.Descriptor instead.
func (*ResolutionListReq) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_resolution_proto_rawDescGZIP(), []int{4}
}

func (x *ResolutionListReq) GetPlatform() proto.Platform {
	if x != nil {
		return x.Platform
	}
	return proto.Platform(0)
}

var File_services_chipmunk_api_proto_src_resolution_proto protoreflect.FileDescriptor

var file_services_chipmunk_api_proto_src_resolution_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x1a,
	0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d,
	0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x0a, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x42, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x29, 0x0a, 0x17, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x66, 0x0a, 0x1d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x79, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e,
	0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x32, 0xa6,
	0x02, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x68,
	0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x4b, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x24, 0x2e,
	0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x57, 0x0a, 0x10,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42,
	0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x63,
	0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e,
	0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e,
	0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72, 0x6d, 0x61, 0x7a, 0x79, 0x61,
	0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_chipmunk_api_proto_src_resolution_proto_rawDescOnce sync.Once
	file_services_chipmunk_api_proto_src_resolution_proto_rawDescData = file_services_chipmunk_api_proto_src_resolution_proto_rawDesc
)

func file_services_chipmunk_api_proto_src_resolution_proto_rawDescGZIP() []byte {
	file_services_chipmunk_api_proto_src_resolution_proto_rawDescOnce.Do(func() {
		file_services_chipmunk_api_proto_src_resolution_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_chipmunk_api_proto_src_resolution_proto_rawDescData)
	})
	return file_services_chipmunk_api_proto_src_resolution_proto_rawDescData
}

var file_services_chipmunk_api_proto_src_resolution_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_chipmunk_api_proto_src_resolution_proto_goTypes = []interface{}{
	(*Resolution)(nil),                    // 0: chipmunkApi.Resolution
	(*Resolutions)(nil),                   // 1: chipmunkApi.Resolutions
	(*ResolutionReturnByIDReq)(nil),       // 2: chipmunkApi.ResolutionReturnByIDReq
	(*ResolutionReturnByDurationReq)(nil), // 3: chipmunkApi.ResolutionReturnByDurationReq
	(*ResolutionListReq)(nil),             // 4: chipmunkApi.ResolutionListReq
	(proto.Platform)(0),                   // 5: api.Platform
	(*proto.Void)(nil),                    // 6: api.Void
}
var file_services_chipmunk_api_proto_src_resolution_proto_depIdxs = []int32{
	5, // 0: chipmunkApi.Resolution.Platform:type_name -> api.Platform
	0, // 1: chipmunkApi.Resolutions.Elements:type_name -> chipmunkApi.Resolution
	5, // 2: chipmunkApi.ResolutionReturnByDurationReq.Platform:type_name -> api.Platform
	5, // 3: chipmunkApi.ResolutionListReq.Platform:type_name -> api.Platform
	0, // 4: chipmunkApi.ResolutionService.Set:input_type -> chipmunkApi.Resolution
	2, // 5: chipmunkApi.ResolutionService.ReturnByID:input_type -> chipmunkApi.ResolutionReturnByIDReq
	3, // 6: chipmunkApi.ResolutionService.ReturnByDuration:input_type -> chipmunkApi.ResolutionReturnByDurationReq
	4, // 7: chipmunkApi.ResolutionService.List:input_type -> chipmunkApi.ResolutionListReq
	6, // 8: chipmunkApi.ResolutionService.Set:output_type -> api.Void
	0, // 9: chipmunkApi.ResolutionService.ReturnByID:output_type -> chipmunkApi.Resolution
	0, // 10: chipmunkApi.ResolutionService.ReturnByDuration:output_type -> chipmunkApi.Resolution
	1, // 11: chipmunkApi.ResolutionService.List:output_type -> chipmunkApi.Resolutions
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_chipmunk_api_proto_src_resolution_proto_init() }
func file_services_chipmunk_api_proto_src_resolution_proto_init() {
	if File_services_chipmunk_api_proto_src_resolution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resolution); i {
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
		file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resolutions); i {
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
		file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolutionReturnByIDReq); i {
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
		file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolutionReturnByDurationReq); i {
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
		file_services_chipmunk_api_proto_src_resolution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolutionListReq); i {
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
			RawDescriptor: file_services_chipmunk_api_proto_src_resolution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_chipmunk_api_proto_src_resolution_proto_goTypes,
		DependencyIndexes: file_services_chipmunk_api_proto_src_resolution_proto_depIdxs,
		MessageInfos:      file_services_chipmunk_api_proto_src_resolution_proto_msgTypes,
	}.Build()
	File_services_chipmunk_api_proto_src_resolution_proto = out.File
	file_services_chipmunk_api_proto_src_resolution_proto_rawDesc = nil
	file_services_chipmunk_api_proto_src_resolution_proto_goTypes = nil
	file_services_chipmunk_api_proto_src_resolution_proto_depIdxs = nil
}
