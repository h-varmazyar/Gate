// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: services/eagle/api/proto/src/signal.proto

package proto

import (
	api "github.com/h-varmazyar/Gate/api/proto"
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

type SignalStartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"brokerage_id"
	BrokerageID string `protobuf:"bytes,1,opt,name=BrokerageID,proto3" json:"brokerage_id"`
	// @inject_tag: json:"strategy_id"
	StrategyID string `protobuf:"bytes,2,opt,name=StrategyID,proto3" json:"strategy_id"`
	// @inject_tag: json:"with_trading"
	WithTrading bool `protobuf:"varint,3,opt,name=WithTrading,proto3" json:"with_trading"`
}

func (x *SignalStartReq) Reset() {
	*x = SignalStartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_proto_src_signal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalStartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalStartReq) ProtoMessage() {}

func (x *SignalStartReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_proto_src_signal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalStartReq.ProtoReflect.Descriptor instead.
func (*SignalStartReq) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_proto_src_signal_proto_rawDescGZIP(), []int{0}
}

func (x *SignalStartReq) GetBrokerageID() string {
	if x != nil {
		return x.BrokerageID
	}
	return ""
}

func (x *SignalStartReq) GetStrategyID() string {
	if x != nil {
		return x.StrategyID
	}
	return ""
}

func (x *SignalStartReq) GetWithTrading() bool {
	if x != nil {
		return x.WithTrading
	}
	return false
}

var File_services_eagle_api_proto_src_signal_proto protoreflect.FileDescriptor

var file_services_eagle_api_proto_src_signal_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x61, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x61, 0x67,
	0x6c, 0x65, 0x41, 0x70, 0x69, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d,
	0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x0e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x32,
	0x5b, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x18, 0x2e, 0x65, 0x61, 0x67, 0x6c,
	0x65, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69,
	0x64, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72,
	0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_eagle_api_proto_src_signal_proto_rawDescOnce sync.Once
	file_services_eagle_api_proto_src_signal_proto_rawDescData = file_services_eagle_api_proto_src_signal_proto_rawDesc
)

func file_services_eagle_api_proto_src_signal_proto_rawDescGZIP() []byte {
	file_services_eagle_api_proto_src_signal_proto_rawDescOnce.Do(func() {
		file_services_eagle_api_proto_src_signal_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_eagle_api_proto_src_signal_proto_rawDescData)
	})
	return file_services_eagle_api_proto_src_signal_proto_rawDescData
}

var file_services_eagle_api_proto_src_signal_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_eagle_api_proto_src_signal_proto_goTypes = []interface{}{
	(*SignalStartReq)(nil), // 0: eagleApi.SignalStartReq
	(*api.Void)(nil),       // 1: api.Void
}
var file_services_eagle_api_proto_src_signal_proto_depIdxs = []int32{
	0, // 0: eagleApi.SignalService.Start:input_type -> eagleApi.SignalStartReq
	1, // 1: eagleApi.SignalService.Stop:input_type -> api.Void
	1, // 2: eagleApi.SignalService.Start:output_type -> api.Void
	1, // 3: eagleApi.SignalService.Stop:output_type -> api.Void
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_eagle_api_proto_src_signal_proto_init() }
func file_services_eagle_api_proto_src_signal_proto_init() {
	if File_services_eagle_api_proto_src_signal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_eagle_api_proto_src_signal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalStartReq); i {
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
			RawDescriptor: file_services_eagle_api_proto_src_signal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_eagle_api_proto_src_signal_proto_goTypes,
		DependencyIndexes: file_services_eagle_api_proto_src_signal_proto_depIdxs,
		MessageInfos:      file_services_eagle_api_proto_src_signal_proto_msgTypes,
	}.Build()
	File_services_eagle_api_proto_src_signal_proto = out.File
	file_services_eagle_api_proto_src_signal_proto_rawDesc = nil
	file_services_eagle_api_proto_src_signal_proto_goTypes = nil
	file_services_eagle_api_proto_src_signal_proto_depIdxs = nil
}
