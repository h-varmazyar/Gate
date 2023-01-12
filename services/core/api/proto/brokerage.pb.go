// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: services/core/api/proto/src/brokerage.proto

package proto

import (
	proto "github.com/h-varmazyar/Gate/api/proto"
	proto1 "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
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

type Brokerage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"title"`
	// @inject_tag: json:"description"
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"description"`
	// @inject_tag: json:"auth"
	Auth *proto.Auth `protobuf:"bytes,4,opt,name=Auth,proto3" json:"auth"`
	// @inject_tag: json:"platform"
	Platform proto.Platform `protobuf:"varint,5,opt,name=Platform,proto3,enum=api.Platform" json:"platform"`
	// @inject_tag: json:"markets"
	Markets *proto1.Markets `protobuf:"bytes,6,opt,name=Markets,proto3" json:"markets"`
	// @inject_tag: json:"resolution_id"
	ResolutionID string `protobuf:"bytes,7,opt,name=ResolutionID,proto3" json:"resolution_id"`
	// @inject_tag: json:"resolution"
	Resolution *proto1.Resolution `protobuf:"bytes,8,opt,name=Resolution,proto3" json:"resolution"`
	// @inject_tag: json:"strategy_id"
	StrategyID string `protobuf:"bytes,9,opt,name=StrategyID,proto3" json:"strategy_id"`
	// @inject_tag: json:"status"
	Status proto.Status `protobuf:"varint,10,opt,name=Status,proto3,enum=api.Status" json:"status"`
}

func (x *Brokerage) Reset() {
	*x = Brokerage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brokerage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brokerage) ProtoMessage() {}

func (x *Brokerage) ProtoReflect() protoreflect.Message {
	mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Brokerage.ProtoReflect.Descriptor instead.
func (*Brokerage) Descriptor() ([]byte, []int) {
	return file_services_core_api_proto_src_brokerage_proto_rawDescGZIP(), []int{0}
}

func (x *Brokerage) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Brokerage) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Brokerage) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Brokerage) GetAuth() *proto.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Brokerage) GetPlatform() proto.Platform {
	if x != nil {
		return x.Platform
	}
	return proto.Platform(0)
}

func (x *Brokerage) GetMarkets() *proto1.Markets {
	if x != nil {
		return x.Markets
	}
	return nil
}

func (x *Brokerage) GetResolutionID() string {
	if x != nil {
		return x.ResolutionID
	}
	return ""
}

func (x *Brokerage) GetResolution() *proto1.Resolution {
	if x != nil {
		return x.Resolution
	}
	return nil
}

func (x *Brokerage) GetStrategyID() string {
	if x != nil {
		return x.StrategyID
	}
	return ""
}

func (x *Brokerage) GetStatus() proto.Status {
	if x != nil {
		return x.Status
	}
	return proto.Status(0)
}

type Brokerages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"elements"
	Elements []*Brokerage `protobuf:"bytes,1,rep,name=Elements,proto3" json:"elements"`
}

func (x *Brokerages) Reset() {
	*x = Brokerages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brokerages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brokerages) ProtoMessage() {}

func (x *Brokerages) ProtoReflect() protoreflect.Message {
	mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Brokerages.ProtoReflect.Descriptor instead.
func (*Brokerages) Descriptor() ([]byte, []int) {
	return file_services_core_api_proto_src_brokerage_proto_rawDescGZIP(), []int{1}
}

func (x *Brokerages) GetElements() []*Brokerage {
	if x != nil {
		return x.Elements
	}
	return nil
}

type BrokerageCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"title"`
	// @inject_tag: json:"description"
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"description"`
	// @inject_tag: json:"auth"
	Auth *proto.Auth `protobuf:"bytes,3,opt,name=Auth,proto3" json:"auth"`
	// @inject_tag: json:"platform"
	Platform proto.Platform `protobuf:"varint,4,opt,name=Platform,proto3,enum=api.Platform" json:"platform"`
	// @inject_tag: json:"markets"
	Markets *proto1.Markets `protobuf:"bytes,5,opt,name=Markets,proto3" json:"markets"`
	// @inject_tag: json:"strategy_id"
	StrategyID string `protobuf:"bytes,7,opt,name=StrategyID,proto3" json:"strategy_id"`
}

func (x *BrokerageCreateReq) Reset() {
	*x = BrokerageCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerageCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerageCreateReq) ProtoMessage() {}

func (x *BrokerageCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerageCreateReq.ProtoReflect.Descriptor instead.
func (*BrokerageCreateReq) Descriptor() ([]byte, []int) {
	return file_services_core_api_proto_src_brokerage_proto_rawDescGZIP(), []int{2}
}

func (x *BrokerageCreateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BrokerageCreateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BrokerageCreateReq) GetAuth() *proto.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *BrokerageCreateReq) GetPlatform() proto.Platform {
	if x != nil {
		return x.Platform
	}
	return proto.Platform(0)
}

func (x *BrokerageCreateReq) GetMarkets() *proto1.Markets {
	if x != nil {
		return x.Markets
	}
	return nil
}

func (x *BrokerageCreateReq) GetStrategyID() string {
	if x != nil {
		return x.StrategyID
	}
	return ""
}

type StatusChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"trading"
	Trading bool `protobuf:"varint,2,opt,name=Trading,proto3" json:"trading"`
	// @inject_tag: json:"ohlc"
	OHLC bool `protobuf:"varint,3,opt,name=OHLC,proto3" json:"ohlc"`
}

func (x *StatusChangeRequest) Reset() {
	*x = StatusChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusChangeRequest) ProtoMessage() {}

func (x *StatusChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusChangeRequest.ProtoReflect.Descriptor instead.
func (*StatusChangeRequest) Descriptor() ([]byte, []int) {
	return file_services_core_api_proto_src_brokerage_proto_rawDescGZIP(), []int{3}
}

func (x *StatusChangeRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *StatusChangeRequest) GetTrading() bool {
	if x != nil {
		return x.Trading
	}
	return false
}

func (x *StatusChangeRequest) GetOHLC() bool {
	if x != nil {
		return x.OHLC
	}
	return false
}

type BrokerageReturnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"with_markets"
	WithMarkets bool `protobuf:"varint,2,opt,name=WithMarkets,proto3" json:"with_markets"`
}

func (x *BrokerageReturnReq) Reset() {
	*x = BrokerageReturnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerageReturnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerageReturnReq) ProtoMessage() {}

func (x *BrokerageReturnReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerageReturnReq.ProtoReflect.Descriptor instead.
func (*BrokerageReturnReq) Descriptor() ([]byte, []int) {
	return file_services_core_api_proto_src_brokerage_proto_rawDescGZIP(), []int{4}
}

func (x *BrokerageReturnReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *BrokerageReturnReq) GetWithMarkets() bool {
	if x != nil {
		return x.WithMarkets
	}
	return false
}

type BrokerageDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
}

func (x *BrokerageDeleteReq) Reset() {
	*x = BrokerageDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerageDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerageDeleteReq) ProtoMessage() {}

func (x *BrokerageDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerageDeleteReq.ProtoReflect.Descriptor instead.
func (*BrokerageDeleteReq) Descriptor() ([]byte, []int) {
	return file_services_core_api_proto_src_brokerage_proto_rawDescGZIP(), []int{5}
}

func (x *BrokerageDeleteReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type BrokerageStartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"with_trading"
	WithTrading bool `protobuf:"varint,2,opt,name=WithTrading,proto3" json:"with_trading"`
}

func (x *BrokerageStartReq) Reset() {
	*x = BrokerageStartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerageStartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerageStartReq) ProtoMessage() {}

func (x *BrokerageStartReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerageStartReq.ProtoReflect.Descriptor instead.
func (*BrokerageStartReq) Descriptor() ([]byte, []int) {
	return file_services_core_api_proto_src_brokerage_proto_rawDescGZIP(), []int{6}
}

func (x *BrokerageStartReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *BrokerageStartReq) GetWithTrading() bool {
	if x != nil {
		return x.WithTrading
	}
	return false
}

type BrokerageStopReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
}

func (x *BrokerageStopReq) Reset() {
	*x = BrokerageStopReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerageStopReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerageStopReq) ProtoMessage() {}

func (x *BrokerageStopReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_core_api_proto_src_brokerage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerageStopReq.ProtoReflect.Descriptor instead.
func (*BrokerageStopReq) Descriptor() ([]byte, []int) {
	return file_services_core_api_proto_src_brokerage_proto_rawDescGZIP(), []int{7}
}

func (x *BrokerageStopReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_services_core_api_proto_src_brokerage_proto protoreflect.FileDescriptor

var file_services_core_api_proto_src_brokerage_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63,
	0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e,
	0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xef, 0x02, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x29, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x2e, 0x0a, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x12, 0x23, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x3c, 0x0a, 0x0a, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x2e, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0xe6, 0x01, 0x0a, 0x12, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x29,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52,
	0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2e, 0x0a, 0x07, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x69,
	0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73,
	0x52, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x22, 0x53, 0x0a, 0x13, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x48,
	0x4c, 0x43, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x4f, 0x48, 0x4c, 0x43, 0x22, 0x46,
	0x0a, 0x12, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x57, 0x69, 0x74, 0x68, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x45, 0x0a, 0x11,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x22, 0x22, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32, 0xd2, 0x02, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69,
	0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x35, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x41,
	0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x6f, 0x70,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72,
	0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_core_api_proto_src_brokerage_proto_rawDescOnce sync.Once
	file_services_core_api_proto_src_brokerage_proto_rawDescData = file_services_core_api_proto_src_brokerage_proto_rawDesc
)

func file_services_core_api_proto_src_brokerage_proto_rawDescGZIP() []byte {
	file_services_core_api_proto_src_brokerage_proto_rawDescOnce.Do(func() {
		file_services_core_api_proto_src_brokerage_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_core_api_proto_src_brokerage_proto_rawDescData)
	})
	return file_services_core_api_proto_src_brokerage_proto_rawDescData
}

var file_services_core_api_proto_src_brokerage_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_core_api_proto_src_brokerage_proto_goTypes = []interface{}{
	(*Brokerage)(nil),           // 0: coreApi.Brokerage
	(*Brokerages)(nil),          // 1: coreApi.Brokerages
	(*BrokerageCreateReq)(nil),  // 2: coreApi.BrokerageCreateReq
	(*StatusChangeRequest)(nil), // 3: coreApi.StatusChangeRequest
	(*BrokerageReturnReq)(nil),  // 4: coreApi.BrokerageReturnReq
	(*BrokerageDeleteReq)(nil),  // 5: coreApi.BrokerageDeleteReq
	(*BrokerageStartReq)(nil),   // 6: coreApi.BrokerageStartReq
	(*BrokerageStopReq)(nil),    // 7: coreApi.BrokerageStopReq
	(*proto.Auth)(nil),          // 8: api.Auth
	(proto.Platform)(0),         // 9: api.Platform
	(*proto1.Markets)(nil),      // 10: chipmunkApi.Markets
	(*proto1.Resolution)(nil),   // 11: chipmunkApi.Resolution
	(proto.Status)(0),           // 12: api.Status
	(*proto.Void)(nil),          // 13: api.Void
}
var file_services_core_api_proto_src_brokerage_proto_depIdxs = []int32{
	8,  // 0: coreApi.Brokerage.Auth:type_name -> api.Auth
	9,  // 1: coreApi.Brokerage.Platform:type_name -> api.Platform
	10, // 2: coreApi.Brokerage.Markets:type_name -> chipmunkApi.Markets
	11, // 3: coreApi.Brokerage.Resolution:type_name -> chipmunkApi.Resolution
	12, // 4: coreApi.Brokerage.Status:type_name -> api.Status
	0,  // 5: coreApi.Brokerages.Elements:type_name -> coreApi.Brokerage
	8,  // 6: coreApi.BrokerageCreateReq.Auth:type_name -> api.Auth
	9,  // 7: coreApi.BrokerageCreateReq.Platform:type_name -> api.Platform
	10, // 8: coreApi.BrokerageCreateReq.Markets:type_name -> chipmunkApi.Markets
	2,  // 9: coreApi.BrokerageService.Create:input_type -> coreApi.BrokerageCreateReq
	6,  // 10: coreApi.BrokerageService.Start:input_type -> coreApi.BrokerageStartReq
	7,  // 11: coreApi.BrokerageService.Stop:input_type -> coreApi.BrokerageStopReq
	4,  // 12: coreApi.BrokerageService.Return:input_type -> coreApi.BrokerageReturnReq
	5,  // 13: coreApi.BrokerageService.Delete:input_type -> coreApi.BrokerageDeleteReq
	13, // 14: coreApi.BrokerageService.List:input_type -> api.Void
	0,  // 15: coreApi.BrokerageService.Create:output_type -> coreApi.Brokerage
	0,  // 16: coreApi.BrokerageService.Start:output_type -> coreApi.Brokerage
	0,  // 17: coreApi.BrokerageService.Stop:output_type -> coreApi.Brokerage
	0,  // 18: coreApi.BrokerageService.Return:output_type -> coreApi.Brokerage
	13, // 19: coreApi.BrokerageService.Delete:output_type -> api.Void
	1,  // 20: coreApi.BrokerageService.List:output_type -> coreApi.Brokerages
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_services_core_api_proto_src_brokerage_proto_init() }
func file_services_core_api_proto_src_brokerage_proto_init() {
	if File_services_core_api_proto_src_brokerage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_core_api_proto_src_brokerage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Brokerage); i {
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
		file_services_core_api_proto_src_brokerage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Brokerages); i {
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
		file_services_core_api_proto_src_brokerage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerageCreateReq); i {
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
		file_services_core_api_proto_src_brokerage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusChangeRequest); i {
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
		file_services_core_api_proto_src_brokerage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerageReturnReq); i {
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
		file_services_core_api_proto_src_brokerage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerageDeleteReq); i {
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
		file_services_core_api_proto_src_brokerage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerageStartReq); i {
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
		file_services_core_api_proto_src_brokerage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerageStopReq); i {
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
			RawDescriptor: file_services_core_api_proto_src_brokerage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_core_api_proto_src_brokerage_proto_goTypes,
		DependencyIndexes: file_services_core_api_proto_src_brokerage_proto_depIdxs,
		MessageInfos:      file_services_core_api_proto_src_brokerage_proto_msgTypes,
	}.Build()
	File_services_core_api_proto_src_brokerage_proto = out.File
	file_services_core_api_proto_src_brokerage_proto_rawDesc = nil
	file_services_core_api_proto_src_brokerage_proto_goTypes = nil
	file_services_core_api_proto_src_brokerage_proto_depIdxs = nil
}
