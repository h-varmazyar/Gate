// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: services/brokerage/api/src/candle.proto

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

type OhlcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"resolution"
	Resolution *Resolution `protobuf:"bytes,1,opt,name=Resolution,proto3" json:"resolution"`
	// @inject_tag: json:"market"
	Market *Market `protobuf:"bytes,2,opt,name=Market,proto3" json:"market"`
	// @inject_tag: json:"from"
	From int64 `protobuf:"varint,3,opt,name=From,proto3" json:"from"`
	// @inject_tag: json:"to"
	To int64 `protobuf:"varint,4,opt,name=To,proto3" json:"to"`
}

func (x *OhlcRequest) Reset() {
	*x = OhlcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_candle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OhlcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OhlcRequest) ProtoMessage() {}

func (x *OhlcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_candle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OhlcRequest.ProtoReflect.Descriptor instead.
func (*OhlcRequest) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_candle_proto_rawDescGZIP(), []int{0}
}

func (x *OhlcRequest) GetResolution() *Resolution {
	if x != nil {
		return x.Resolution
	}
	return nil
}

func (x *OhlcRequest) GetMarket() *Market {
	if x != nil {
		return x.Market
	}
	return nil
}

func (x *OhlcRequest) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *OhlcRequest) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

var File_services_brokerage_api_src_candle_proto protoreflect.FileDescriptor

var file_services_brokerage_api_src_candle_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x1a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x0b, 0x4f, 0x68, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a,
	0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x52, 0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x46,
	0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x54, 0x6f, 0x32,
	0x40, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x4f, 0x48, 0x4c, 0x43, 0x12, 0x19, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x4f, 0x68, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x2d, 0x76, 0x61, 0x72, 0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_brokerage_api_src_candle_proto_rawDescOnce sync.Once
	file_services_brokerage_api_src_candle_proto_rawDescData = file_services_brokerage_api_src_candle_proto_rawDesc
)

func file_services_brokerage_api_src_candle_proto_rawDescGZIP() []byte {
	file_services_brokerage_api_src_candle_proto_rawDescOnce.Do(func() {
		file_services_brokerage_api_src_candle_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_brokerage_api_src_candle_proto_rawDescData)
	})
	return file_services_brokerage_api_src_candle_proto_rawDescData
}

var file_services_brokerage_api_src_candle_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_services_brokerage_api_src_candle_proto_goTypes = []interface{}{
	(*OhlcRequest)(nil), // 0: brokerageApi.OhlcRequest
	(*Resolution)(nil),  // 1: brokerageApi.Resolution
	(*Market)(nil),      // 2: brokerageApi.Market
	(*api.Candles)(nil), // 3: api.Candles
}
var file_services_brokerage_api_src_candle_proto_depIdxs = []int32{
	1, // 0: brokerageApi.OhlcRequest.Resolution:type_name -> brokerageApi.Resolution
	2, // 1: brokerageApi.OhlcRequest.Market:type_name -> brokerageApi.Market
	0, // 2: brokerageApi.CandleService.OHLC:input_type -> brokerageApi.OhlcRequest
	3, // 3: brokerageApi.CandleService.OHLC:output_type -> api.Candles
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_brokerage_api_src_candle_proto_init() }
func file_services_brokerage_api_src_candle_proto_init() {
	if File_services_brokerage_api_src_candle_proto != nil {
		return
	}
	file_services_brokerage_api_src_market_proto_init()
	file_services_brokerage_api_src_resolution_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_services_brokerage_api_src_candle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OhlcRequest); i {
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
			RawDescriptor: file_services_brokerage_api_src_candle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_brokerage_api_src_candle_proto_goTypes,
		DependencyIndexes: file_services_brokerage_api_src_candle_proto_depIdxs,
		MessageInfos:      file_services_brokerage_api_src_candle_proto_msgTypes,
	}.Build()
	File_services_brokerage_api_src_candle_proto = out.File
	file_services_brokerage_api_src_candle_proto_rawDesc = nil
	file_services_brokerage_api_src_candle_proto_goTypes = nil
	file_services_brokerage_api_src_candle_proto_depIdxs = nil
}
