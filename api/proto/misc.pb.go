// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: api/proto/src/misc.proto

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

type AuthType int32

const (
	AuthType_None             AuthType = 0
	AuthType_StaticToken      AuthType = 1
	AuthType_UsernamePassword AuthType = 2
)

// Enum value maps for AuthType.
var (
	AuthType_name = map[int32]string{
		0: "None",
		1: "StaticToken",
		2: "UsernamePassword",
	}
	AuthType_value = map[string]int32{
		"None":             0,
		"StaticToken":      1,
		"UsernamePassword": 2,
	}
)

func (x AuthType) Enum() *AuthType {
	p := new(AuthType)
	*p = x
	return p
}

func (x AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_src_misc_proto_enumTypes[0].Descriptor()
}

func (AuthType) Type() protoreflect.EnumType {
	return &file_api_proto_src_misc_proto_enumTypes[0]
}

func (x AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuthType.Descriptor instead.
func (AuthType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_src_misc_proto_rawDescGZIP(), []int{0}
}

type Platform int32

const (
	Platform_UnknownBrokerage Platform = 0
	Platform_Coinex           Platform = 1
	Platform_Nobitex          Platform = 2
	Platform_Mazdax           Platform = 3
	Platform_Binance          Platform = 4
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0: "UnknownBrokerage",
		1: "Coinex",
		2: "Nobitex",
		3: "Mazdax",
		4: "Binance",
	}
	Platform_value = map[string]int32{
		"UnknownBrokerage": 0,
		"Coinex":           1,
		"Nobitex":          2,
		"Mazdax":           3,
		"Binance":          4,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_src_misc_proto_enumTypes[1].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_api_proto_src_misc_proto_enumTypes[1]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_src_misc_proto_rawDescGZIP(), []int{1}
}

type Status int32

const (
	Status_Disable Status = 0
	Status_Enable  Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Disable",
		1: "Enable",
	}
	Status_value = map[string]int32{
		"Disable": 0,
		"Enable":  1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_src_misc_proto_enumTypes[2].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_proto_src_misc_proto_enumTypes[2]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_src_misc_proto_rawDescGZIP(), []int{2}
}

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"type"
	Type AuthType `protobuf:"varint,1,opt,name=Type,proto3,enum=api.AuthType" json:"type"`
	// @inject_tag: json:"username"
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"username"`
	// @inject_tag: json:"password"
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"password"`
	// @inject_tag: json:"access_id"
	AccessID string `protobuf:"bytes,4,opt,name=AccessID,proto3" json:"access_id"`
	// @inject_tag: json:"secret_key"
	SecretKey string `protobuf:"bytes,5,opt,name=SecretKey,proto3" json:"secret_key"`
	// @inject_tag: json:"nobitex_token"
	NobitexToken string `protobuf:"bytes,6,opt,name=NobitexToken,proto3" json:"nobitex_token"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_src_misc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_src_misc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_api_proto_src_misc_proto_rawDescGZIP(), []int{0}
}

func (x *Auth) GetType() AuthType {
	if x != nil {
		return x.Type
	}
	return AuthType_None
}

func (x *Auth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Auth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Auth) GetAccessID() string {
	if x != nil {
		return x.AccessID
	}
	return ""
}

func (x *Auth) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

func (x *Auth) GetNobitexToken() string {
	if x != nil {
		return x.NobitexToken
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_src_misc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_src_misc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_api_proto_src_misc_proto_rawDescGZIP(), []int{1}
}

var File_api_proto_src_misc_proto protoreflect.FileDescriptor

var file_api_proto_src_misc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22,
	0xbf, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a,
	0x0c, 0x4e, 0x6f, 0x62, 0x69, 0x74, 0x65, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x4e, 0x6f, 0x62, 0x69, 0x74, 0x65, 0x78, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x2a, 0x3b, 0x0a, 0x08, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x10, 0x01,
	0x12, 0x14, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x10, 0x02, 0x2a, 0x52, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x6f, 0x69, 0x6e,
	0x65, 0x78, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x6f, 0x62, 0x69, 0x74, 0x65, 0x78, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x61, 0x7a, 0x64, 0x61, 0x78, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x42, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x04, 0x2a, 0x21, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61,
	0x72, 0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_src_misc_proto_rawDescOnce sync.Once
	file_api_proto_src_misc_proto_rawDescData = file_api_proto_src_misc_proto_rawDesc
)

func file_api_proto_src_misc_proto_rawDescGZIP() []byte {
	file_api_proto_src_misc_proto_rawDescOnce.Do(func() {
		file_api_proto_src_misc_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_src_misc_proto_rawDescData)
	})
	return file_api_proto_src_misc_proto_rawDescData
}

var file_api_proto_src_misc_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_proto_src_misc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_src_misc_proto_goTypes = []interface{}{
	(AuthType)(0), // 0: api.AuthType
	(Platform)(0), // 1: api.Platform
	(Status)(0),   // 2: api.Status
	(*Auth)(nil),  // 3: api.Auth
	(*Void)(nil),  // 4: api.Void
}
var file_api_proto_src_misc_proto_depIdxs = []int32{
	0, // 0: api.Auth.Type:type_name -> api.AuthType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_src_misc_proto_init() }
func file_api_proto_src_misc_proto_init() {
	if File_api_proto_src_misc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_src_misc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_api_proto_src_misc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
			RawDescriptor: file_api_proto_src_misc_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_src_misc_proto_goTypes,
		DependencyIndexes: file_api_proto_src_misc_proto_depIdxs,
		EnumInfos:         file_api_proto_src_misc_proto_enumTypes,
		MessageInfos:      file_api_proto_src_misc_proto_msgTypes,
	}.Build()
	File_api_proto_src_misc_proto = out.File
	file_api_proto_src_misc_proto_rawDesc = nil
	file_api_proto_src_misc_proto_goTypes = nil
	file_api_proto_src_misc_proto_depIdxs = nil
}
