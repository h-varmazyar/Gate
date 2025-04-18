// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: services/network/api/proto/src/IP.proto

package proto

import (
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

type IPCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"address"
	Address string `protobuf:"bytes,1,opt,name=Address,proto3" json:"address"`
	//	@inject_tag:	json:"port"
	Port uint32 `protobuf:"varint,2,opt,name=Port,proto3" json:"port"`
	//	@inject_tag:	json:"username"
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"username"`
	//	@inject_tag:	json:"password"
	Password string `protobuf:"bytes,4,opt,name=Password,proto3" json:"password"`
	//	@inject_tag:	json:"schema"
	Schema string `protobuf:"bytes,5,opt,name=Schema,proto3" json:"schema"`
}

func (x *IPCreateReq) Reset() {
	*x = IPCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_network_api_proto_src_IP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPCreateReq) ProtoMessage() {}

func (x *IPCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_network_api_proto_src_IP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPCreateReq.ProtoReflect.Descriptor instead.
func (*IPCreateReq) Descriptor() ([]byte, []int) {
	return file_services_network_api_proto_src_IP_proto_rawDescGZIP(), []int{0}
}

func (x *IPCreateReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *IPCreateReq) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *IPCreateReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IPCreateReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *IPCreateReq) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type IPReturnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
}

func (x *IPReturnReq) Reset() {
	*x = IPReturnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_network_api_proto_src_IP_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPReturnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPReturnReq) ProtoMessage() {}

func (x *IPReturnReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_network_api_proto_src_IP_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPReturnReq.ProtoReflect.Descriptor instead.
func (*IPReturnReq) Descriptor() ([]byte, []int) {
	return file_services_network_api_proto_src_IP_proto_rawDescGZIP(), []int{1}
}

func (x *IPReturnReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type IPListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IPListReq) Reset() {
	*x = IPListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_network_api_proto_src_IP_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPListReq) ProtoMessage() {}

func (x *IPListReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_network_api_proto_src_IP_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPListReq.ProtoReflect.Descriptor instead.
func (*IPListReq) Descriptor() ([]byte, []int) {
	return file_services_network_api_proto_src_IP_proto_rawDescGZIP(), []int{2}
}

type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	//	@inject_tag:	json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,2,opt,name=UpdatedAt,proto3" json:"updated_at"`
	//	@inject_tag:	json:"created_at"
	CreatedAt int64 `protobuf:"varint,3,opt,name=CreatedAt,proto3" json:"created_at"`
	//	@inject_tag:	json:"address"
	Address string `protobuf:"bytes,4,opt,name=Address,proto3" json:"address"`
	//	@inject_tag:	json:"port"
	Port uint32 `protobuf:"varint,5,opt,name=Port,proto3" json:"port"`
	//	@inject_tag:	json:"username"
	Username string `protobuf:"bytes,6,opt,name=Username,proto3" json:"username"`
	//	@inject_tag:	json:"password"
	Password string `protobuf:"bytes,7,opt,name=Password,proto3" json:"password"`
	//	@inject_tag:	json:"schema"
	Schema string `protobuf:"bytes,8,opt,name=Schema,proto3" json:"schema"`
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_network_api_proto_src_IP_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_services_network_api_proto_src_IP_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_services_network_api_proto_src_IP_proto_rawDescGZIP(), []int{3}
}

func (x *IP) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *IP) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *IP) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *IP) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *IP) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *IP) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IP) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *IP) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

type IPs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"elements"
	Elements []*IP `protobuf:"bytes,1,rep,name=Elements,proto3" json:"elements"`
}

func (x *IPs) Reset() {
	*x = IPs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_network_api_proto_src_IP_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPs) ProtoMessage() {}

func (x *IPs) ProtoReflect() protoreflect.Message {
	mi := &file_services_network_api_proto_src_IP_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPs.ProtoReflect.Descriptor instead.
func (*IPs) Descriptor() ([]byte, []int) {
	return file_services_network_api_proto_src_IP_proto_rawDescGZIP(), []int{4}
}

func (x *IPs) GetElements() []*IP {
	if x != nil {
		return x.Elements
	}
	return nil
}

var File_services_network_api_proto_src_IP_proto protoreflect.FileDescriptor

var file_services_network_api_proto_src_IP_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x49, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x41, 0x50, 0x49, 0x22, 0x8b, 0x01, 0x0a, 0x0b, 0x49, 0x50, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x22, 0x1d, 0x0a, 0x0b, 0x49, 0x50, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x0b, 0x0a, 0x09, 0x49, 0x50, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22,
	0xce, 0x01, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x22, 0x31, 0x0a, 0x03, 0x49, 0x50, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x41, 0x50, 0x49, 0x2e, 0x49, 0x50, 0x52, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x32, 0xa1, 0x01, 0x0a, 0x09, 0x49, 0x50, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x50, 0x49, 0x2e, 0x49, 0x50, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x50,
	0x49, 0x2e, 0x49, 0x50, 0x12, 0x31, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x17,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x50, 0x49, 0x2e, 0x49, 0x50, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x41, 0x50, 0x49, 0x2e, 0x49, 0x50, 0x12, 0x2e, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x50, 0x49, 0x2e, 0x49, 0x50, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x41, 0x50, 0x49, 0x2e, 0x49, 0x50, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72, 0x6d, 0x61, 0x7a, 0x79, 0x61,
	0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_network_api_proto_src_IP_proto_rawDescOnce sync.Once
	file_services_network_api_proto_src_IP_proto_rawDescData = file_services_network_api_proto_src_IP_proto_rawDesc
)

func file_services_network_api_proto_src_IP_proto_rawDescGZIP() []byte {
	file_services_network_api_proto_src_IP_proto_rawDescOnce.Do(func() {
		file_services_network_api_proto_src_IP_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_network_api_proto_src_IP_proto_rawDescData)
	})
	return file_services_network_api_proto_src_IP_proto_rawDescData
}

var file_services_network_api_proto_src_IP_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_network_api_proto_src_IP_proto_goTypes = []interface{}{
	(*IPCreateReq)(nil), // 0: networkAPI.IPCreateReq
	(*IPReturnReq)(nil), // 1: networkAPI.IPReturnReq
	(*IPListReq)(nil),   // 2: networkAPI.IPListReq
	(*IP)(nil),          // 3: networkAPI.IP
	(*IPs)(nil),         // 4: networkAPI.IPs
}
var file_services_network_api_proto_src_IP_proto_depIdxs = []int32{
	3, // 0: networkAPI.IPs.Elements:type_name -> networkAPI.IP
	0, // 1: networkAPI.IPService.Create:input_type -> networkAPI.IPCreateReq
	1, // 2: networkAPI.IPService.Return:input_type -> networkAPI.IPReturnReq
	2, // 3: networkAPI.IPService.List:input_type -> networkAPI.IPListReq
	3, // 4: networkAPI.IPService.Create:output_type -> networkAPI.IP
	3, // 5: networkAPI.IPService.Return:output_type -> networkAPI.IP
	4, // 6: networkAPI.IPService.List:output_type -> networkAPI.IPs
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_services_network_api_proto_src_IP_proto_init() }
func file_services_network_api_proto_src_IP_proto_init() {
	if File_services_network_api_proto_src_IP_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_network_api_proto_src_IP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPCreateReq); i {
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
		file_services_network_api_proto_src_IP_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPReturnReq); i {
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
		file_services_network_api_proto_src_IP_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPListReq); i {
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
		file_services_network_api_proto_src_IP_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
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
		file_services_network_api_proto_src_IP_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPs); i {
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
			RawDescriptor: file_services_network_api_proto_src_IP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_network_api_proto_src_IP_proto_goTypes,
		DependencyIndexes: file_services_network_api_proto_src_IP_proto_depIdxs,
		MessageInfos:      file_services_network_api_proto_src_IP_proto_msgTypes,
	}.Build()
	File_services_network_api_proto_src_IP_proto = out.File
	file_services_network_api_proto_src_IP_proto_rawDesc = nil
	file_services_network_api_proto_src_IP_proto_goTypes = nil
	file_services_network_api_proto_src_IP_proto_depIdxs = nil
}
