// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: services/chipmunk/api/src/resolution.proto

package api

import (
	api "github.com/h-varmazyar/Gate/api"
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

	// @inject_tag: json:"brokerage_name"
	BrokerageName string `protobuf:"bytes,1,opt,name=BrokerageName,proto3" json:"brokerage_name"`
	// @inject_tag: json:"duration"
	Duration int64 `protobuf:"varint,2,opt,name=Duration,proto3" json:"duration"`
	// @inject_tag: json:"label"
	Label string `protobuf:"bytes,3,opt,name=Label,proto3" json:"label"`
	// @inject_tag: json:"value"
	Value string `protobuf:"bytes,4,opt,name=Value,proto3" json:"value"`
	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,5,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"created_at"
	CreatedAt int64 `protobuf:"varint,6,opt,name=CreatedAt,proto3" json:"created_at"`
	// @inject_tag: json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,7,opt,name=UpdatedAt,proto3" json:"updated_at"`
}

func (x *Resolution) Reset() {
	*x = Resolution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resolution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resolution) ProtoMessage() {}

func (x *Resolution) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[0]
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
	return file_services_chipmunk_api_src_resolution_proto_rawDescGZIP(), []int{0}
}

func (x *Resolution) GetBrokerageName() string {
	if x != nil {
		return x.BrokerageName
	}
	return ""
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

	// @inject_tag: json:"elements"
	Elements []*Resolution `protobuf:"bytes,1,rep,name=Elements,proto3" json:"elements"`
}

func (x *Resolutions) Reset() {
	*x = Resolutions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resolutions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resolutions) ProtoMessage() {}

func (x *Resolutions) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[1]
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
	return file_services_chipmunk_api_src_resolution_proto_rawDescGZIP(), []int{1}
}

func (x *Resolutions) GetElements() []*Resolution {
	if x != nil {
		return x.Elements
	}
	return nil
}

type GetResolutionByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
}

func (x *GetResolutionByIDRequest) Reset() {
	*x = GetResolutionByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResolutionByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResolutionByIDRequest) ProtoMessage() {}

func (x *GetResolutionByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResolutionByIDRequest.ProtoReflect.Descriptor instead.
func (*GetResolutionByIDRequest) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_src_resolution_proto_rawDescGZIP(), []int{2}
}

func (x *GetResolutionByIDRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetResolutionByDurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"brokerage_name"
	BrokerageName string `protobuf:"bytes,1,opt,name=BrokerageName,proto3" json:"brokerage_name"`
	// @inject_tag: json:"duration"
	Duration int64 `protobuf:"varint,2,opt,name=Duration,proto3" json:"duration"`
}

func (x *GetResolutionByDurationRequest) Reset() {
	*x = GetResolutionByDurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResolutionByDurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResolutionByDurationRequest) ProtoMessage() {}

func (x *GetResolutionByDurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResolutionByDurationRequest.ProtoReflect.Descriptor instead.
func (*GetResolutionByDurationRequest) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_src_resolution_proto_rawDescGZIP(), []int{3}
}

func (x *GetResolutionByDurationRequest) GetBrokerageName() string {
	if x != nil {
		return x.BrokerageName
	}
	return ""
}

func (x *GetResolutionByDurationRequest) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type GetResolutionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"brokerage_name"
	BrokerageName string `protobuf:"bytes,1,opt,name=BrokerageName,proto3" json:"brokerage_name"`
}

func (x *GetResolutionListRequest) Reset() {
	*x = GetResolutionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResolutionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResolutionListRequest) ProtoMessage() {}

func (x *GetResolutionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_src_resolution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResolutionListRequest.ProtoReflect.Descriptor instead.
func (*GetResolutionListRequest) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_src_resolution_proto_rawDescGZIP(), []int{4}
}

func (x *GetResolutionListRequest) GetBrokerageName() string {
	if x != nil {
		return x.BrokerageName
	}
	return ""
}

var File_services_chipmunk_api_src_resolution_proto protoreflect.FileDescriptor

var file_services_chipmunk_api_src_resolution_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x68,
	0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x42, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75,
	0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x62, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x42,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xa9, 0x02, 0x0a,
	0x11, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x70,
	0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x49, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x63, 0x68, 0x69, 0x70,
	0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e,
	0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x47, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75,
	0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72, 0x6d, 0x61, 0x7a, 0x79,
	0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_chipmunk_api_src_resolution_proto_rawDescOnce sync.Once
	file_services_chipmunk_api_src_resolution_proto_rawDescData = file_services_chipmunk_api_src_resolution_proto_rawDesc
)

func file_services_chipmunk_api_src_resolution_proto_rawDescGZIP() []byte {
	file_services_chipmunk_api_src_resolution_proto_rawDescOnce.Do(func() {
		file_services_chipmunk_api_src_resolution_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_chipmunk_api_src_resolution_proto_rawDescData)
	})
	return file_services_chipmunk_api_src_resolution_proto_rawDescData
}

var file_services_chipmunk_api_src_resolution_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_chipmunk_api_src_resolution_proto_goTypes = []interface{}{
	(*Resolution)(nil),                     // 0: chipmunkApi.Resolution
	(*Resolutions)(nil),                    // 1: chipmunkApi.Resolutions
	(*GetResolutionByIDRequest)(nil),       // 2: chipmunkApi.GetResolutionByIDRequest
	(*GetResolutionByDurationRequest)(nil), // 3: chipmunkApi.GetResolutionByDurationRequest
	(*GetResolutionListRequest)(nil),       // 4: chipmunkApi.GetResolutionListRequest
	(*api.Void)(nil),                       // 5: api.Void
}
var file_services_chipmunk_api_src_resolution_proto_depIdxs = []int32{
	0, // 0: chipmunkApi.Resolutions.Elements:type_name -> chipmunkApi.Resolution
	0, // 1: chipmunkApi.ResolutionService.Set:input_type -> chipmunkApi.Resolution
	2, // 2: chipmunkApi.ResolutionService.GetByID:input_type -> chipmunkApi.GetResolutionByIDRequest
	3, // 3: chipmunkApi.ResolutionService.GetByDuration:input_type -> chipmunkApi.GetResolutionByDurationRequest
	4, // 4: chipmunkApi.ResolutionService.List:input_type -> chipmunkApi.GetResolutionListRequest
	5, // 5: chipmunkApi.ResolutionService.Set:output_type -> api.Void
	0, // 6: chipmunkApi.ResolutionService.GetByID:output_type -> chipmunkApi.Resolution
	0, // 7: chipmunkApi.ResolutionService.GetByDuration:output_type -> chipmunkApi.Resolution
	1, // 8: chipmunkApi.ResolutionService.List:output_type -> chipmunkApi.Resolutions
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_chipmunk_api_src_resolution_proto_init() }
func file_services_chipmunk_api_src_resolution_proto_init() {
	if File_services_chipmunk_api_src_resolution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_chipmunk_api_src_resolution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_chipmunk_api_src_resolution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_chipmunk_api_src_resolution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResolutionByIDRequest); i {
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
		file_services_chipmunk_api_src_resolution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResolutionByDurationRequest); i {
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
		file_services_chipmunk_api_src_resolution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResolutionListRequest); i {
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
			RawDescriptor: file_services_chipmunk_api_src_resolution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_chipmunk_api_src_resolution_proto_goTypes,
		DependencyIndexes: file_services_chipmunk_api_src_resolution_proto_depIdxs,
		MessageInfos:      file_services_chipmunk_api_src_resolution_proto_msgTypes,
	}.Build()
	File_services_chipmunk_api_src_resolution_proto = out.File
	file_services_chipmunk_api_src_resolution_proto_rawDesc = nil
	file_services_chipmunk_api_src_resolution_proto_goTypes = nil
	file_services_chipmunk_api_src_resolution_proto_depIdxs = nil
}
