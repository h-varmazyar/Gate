// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: services/indicators/api/proto/src/misc.proto

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

type Source int32

const (
	Source_CUSTOM Source = 0
	Source_OHLC4  Source = 1
	Source_CLOSE  Source = 2
	Source_OPEN   Source = 3
	Source_HIGH   Source = 4
	Source_HLC3   Source = 5
	Source_LOW    Source = 6
	Source_HL2    Source = 7
)

// Enum value maps for Source.
var (
	Source_name = map[int32]string{
		0: "CUSTOM",
		1: "OHLC4",
		2: "CLOSE",
		3: "OPEN",
		4: "HIGH",
		5: "HLC3",
		6: "LOW",
		7: "HL2",
	}
	Source_value = map[string]int32{
		"CUSTOM": 0,
		"OHLC4":  1,
		"CLOSE":  2,
		"OPEN":   3,
		"HIGH":   4,
		"HLC3":   5,
		"LOW":    6,
		"HL2":    7,
	}
)

func (x Source) Enum() *Source {
	p := new(Source)
	*p = x
	return p
}

func (x Source) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source) Descriptor() protoreflect.EnumDescriptor {
	return file_services_indicators_api_proto_src_misc_proto_enumTypes[0].Descriptor()
}

func (Source) Type() protoreflect.EnumType {
	return &file_services_indicators_api_proto_src_misc_proto_enumTypes[0]
}

func (x Source) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Source.Descriptor instead.
func (Source) EnumDescriptor() ([]byte, []int) {
	return file_services_indicators_api_proto_src_misc_proto_rawDescGZIP(), []int{0}
}

type Type int32

const (
	Type_NOTHING         Type = 0
	Type_RSI             Type = 1
	Type_STOCHASTIC      Type = 2
	Type_SMA             Type = 3
	Type_EMA             Type = 4
	Type_BOLLINGER_BANDS Type = 5
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "NOTHING",
		1: "RSI",
		2: "STOCHASTIC",
		3: "SMA",
		4: "EMA",
		5: "BOLLINGER_BANDS",
	}
	Type_value = map[string]int32{
		"NOTHING":         0,
		"RSI":             1,
		"STOCHASTIC":      2,
		"SMA":             3,
		"EMA":             4,
		"BOLLINGER_BANDS": 5,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_services_indicators_api_proto_src_misc_proto_enumTypes[1].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_services_indicators_api_proto_src_misc_proto_enumTypes[1]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_services_indicators_api_proto_src_misc_proto_rawDescGZIP(), []int{1}
}

var File_services_indicators_api_proto_src_misc_proto protoreflect.FileDescriptor

var file_services_indicators_api_proto_src_misc_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x41, 0x70, 0x69, 0x2a, 0x5a, 0x0a, 0x06,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x48, 0x4c, 0x43, 0x34, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e,
	0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49, 0x47, 0x48, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04,
	0x48, 0x4c, 0x43, 0x33, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x4f, 0x57, 0x10, 0x06, 0x12,
	0x07, 0x0a, 0x03, 0x48, 0x4c, 0x32, 0x10, 0x07, 0x2a, 0x53, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x52, 0x53, 0x49, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x4f, 0x43, 0x48, 0x41,
	0x53, 0x54, 0x49, 0x43, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x4d, 0x41, 0x10, 0x03, 0x12,
	0x07, 0x0a, 0x03, 0x45, 0x4d, 0x41, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x4f, 0x4c, 0x4c,
	0x49, 0x4e, 0x47, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x4e, 0x44, 0x53, 0x10, 0x05, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61,
	0x72, 0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_services_indicators_api_proto_src_misc_proto_rawDescOnce sync.Once
	file_services_indicators_api_proto_src_misc_proto_rawDescData = file_services_indicators_api_proto_src_misc_proto_rawDesc
)

func file_services_indicators_api_proto_src_misc_proto_rawDescGZIP() []byte {
	file_services_indicators_api_proto_src_misc_proto_rawDescOnce.Do(func() {
		file_services_indicators_api_proto_src_misc_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_indicators_api_proto_src_misc_proto_rawDescData)
	})
	return file_services_indicators_api_proto_src_misc_proto_rawDescData
}

var file_services_indicators_api_proto_src_misc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_services_indicators_api_proto_src_misc_proto_goTypes = []interface{}{
	(Source)(0), // 0: indicatorApi.Source
	(Type)(0),   // 1: indicatorApi.Type
}
var file_services_indicators_api_proto_src_misc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_indicators_api_proto_src_misc_proto_init() }
func file_services_indicators_api_proto_src_misc_proto_init() {
	if File_services_indicators_api_proto_src_misc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_indicators_api_proto_src_misc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_services_indicators_api_proto_src_misc_proto_goTypes,
		DependencyIndexes: file_services_indicators_api_proto_src_misc_proto_depIdxs,
		EnumInfos:         file_services_indicators_api_proto_src_misc_proto_enumTypes,
	}.Build()
	File_services_indicators_api_proto_src_misc_proto = out.File
	file_services_indicators_api_proto_src_misc_proto_rawDesc = nil
	file_services_indicators_api_proto_src_misc_proto_goTypes = nil
	file_services_indicators_api_proto_src_misc_proto_depIdxs = nil
}
