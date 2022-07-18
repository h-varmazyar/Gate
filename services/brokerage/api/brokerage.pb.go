// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: services/brokerage/api/src/brokerage.proto

package api

import (
	api "github.com/h-varmazyar/Gate/api"
	api1 "github.com/h-varmazyar/Gate/services/chipmunk/api"
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

type Platform int32

const (
	Platform_UnknownBrokerage Platform = 0
	Platform_Coinex           Platform = 1
	Platform_Nobitex          Platform = 2
	Platform_All              Platform = 99
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0:  "UnknownBrokerage",
		1:  "Coinex",
		2:  "Nobitex",
		99: "All",
	}
	Platform_value = map[string]int32{
		"UnknownBrokerage": 0,
		"Coinex":           1,
		"Nobitex":          2,
		"All":              99,
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
	return file_services_brokerage_api_src_brokerage_proto_enumTypes[0].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_services_brokerage_api_src_brokerage_proto_enumTypes[0]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{0}
}

type CreateBrokerageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"title"`
	// @inject_tag: json:"description"
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"description"`
	// @inject_tag: json:"auth"
	Auth *api.Auth `protobuf:"bytes,3,opt,name=Auth,proto3" json:"auth"`
	// @inject_tag: json:"platform"
	Platform Platform `protobuf:"varint,4,opt,name=Platform,proto3,enum=brokerageApi.Platform" json:"platform"`
	// @inject_tag: json:"markets"
	Markets *api1.Markets `protobuf:"bytes,5,opt,name=Markets,proto3" json:"markets"`
	// @inject_tag: json:"resolution_id"
	ResolutionID string `protobuf:"bytes,6,opt,name=ResolutionID,proto3" json:"resolution_id"`
	// @inject_tag: json:"strategy_id"
	StrategyID string `protobuf:"bytes,7,opt,name=StrategyID,proto3" json:"strategy_id"`
}

func (x *CreateBrokerageReq) Reset() {
	*x = CreateBrokerageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBrokerageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBrokerageReq) ProtoMessage() {}

func (x *CreateBrokerageReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBrokerageReq.ProtoReflect.Descriptor instead.
func (*CreateBrokerageReq) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBrokerageReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateBrokerageReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBrokerageReq) GetAuth() *api.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *CreateBrokerageReq) GetPlatform() Platform {
	if x != nil {
		return x.Platform
	}
	return Platform_UnknownBrokerage
}

func (x *CreateBrokerageReq) GetMarkets() *api1.Markets {
	if x != nil {
		return x.Markets
	}
	return nil
}

func (x *CreateBrokerageReq) GetResolutionID() string {
	if x != nil {
		return x.ResolutionID
	}
	return ""
}

func (x *CreateBrokerageReq) GetStrategyID() string {
	if x != nil {
		return x.StrategyID
	}
	return ""
}

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
	Auth *api.Auth `protobuf:"bytes,4,opt,name=Auth,proto3" json:"auth"`
	// @inject_tag: json:"platform"
	Platform Platform `protobuf:"varint,5,opt,name=Platform,proto3,enum=brokerageApi.Platform" json:"platform"`
	// @inject_tag: json:"markets"
	Markets *api1.Markets `protobuf:"bytes,6,opt,name=Markets,proto3" json:"markets"`
	// @inject_tag: json:"resolution_id"
	ResolutionID string `protobuf:"bytes,7,opt,name=ResolutionID,proto3" json:"resolution_id"`
	// @inject_tag: json:"resolution"
	Resolution *api1.Resolution `protobuf:"bytes,8,opt,name=Resolution,proto3" json:"resolution"`
	// @inject_tag: json:"strategy_id"
	StrategyID string `protobuf:"bytes,9,opt,name=StrategyID,proto3" json:"strategy_id"`
	// @inject_tag: json:"status"
	Status api.Status `protobuf:"varint,10,opt,name=Status,proto3,enum=api.Status" json:"status"`
}

func (x *Brokerage) Reset() {
	*x = Brokerage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brokerage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brokerage) ProtoMessage() {}

func (x *Brokerage) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[1]
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
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{1}
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

func (x *Brokerage) GetAuth() *api.Auth {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *Brokerage) GetPlatform() Platform {
	if x != nil {
		return x.Platform
	}
	return Platform_UnknownBrokerage
}

func (x *Brokerage) GetMarkets() *api1.Markets {
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

func (x *Brokerage) GetResolution() *api1.Resolution {
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

func (x *Brokerage) GetStatus() api.Status {
	if x != nil {
		return x.Status
	}
	return api.Status(0)
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
		mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Brokerages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Brokerages) ProtoMessage() {}

func (x *Brokerages) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[2]
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
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{2}
}

func (x *Brokerages) GetElements() []*Brokerage {
	if x != nil {
		return x.Elements
	}
	return nil
}

type BrokerageStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"status"
	Status api.Status `protobuf:"varint,1,opt,name=Status,proto3,enum=api.Status" json:"status"`
}

func (x *BrokerageStatus) Reset() {
	*x = BrokerageStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerageStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerageStatus) ProtoMessage() {}

func (x *BrokerageStatus) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerageStatus.ProtoReflect.Descriptor instead.
func (*BrokerageStatus) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{3}
}

func (x *BrokerageStatus) GetStatus() api.Status {
	if x != nil {
		return x.Status
	}
	return api.Status(0)
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
		mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusChangeRequest) ProtoMessage() {}

func (x *StatusChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[4]
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
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{4}
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

type ReturnBrokerageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"with_markets"
	WithMarkets bool `protobuf:"varint,2,opt,name=WithMarkets,proto3" json:"with_markets"`
}

func (x *ReturnBrokerageReq) Reset() {
	*x = ReturnBrokerageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnBrokerageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnBrokerageReq) ProtoMessage() {}

func (x *ReturnBrokerageReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnBrokerageReq.ProtoReflect.Descriptor instead.
func (*ReturnBrokerageReq) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{5}
}

func (x *ReturnBrokerageReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ReturnBrokerageReq) GetWithMarkets() bool {
	if x != nil {
		return x.WithMarkets
	}
	return false
}

type DeleteBrokerageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
}

func (x *DeleteBrokerageReq) Reset() {
	*x = DeleteBrokerageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBrokerageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBrokerageReq) ProtoMessage() {}

func (x *DeleteBrokerageReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBrokerageReq.ProtoReflect.Descriptor instead.
func (*DeleteBrokerageReq) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBrokerageReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ActiveBrokerageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"with_markets"
	WithMarkets bool `protobuf:"varint,1,opt,name=WithMarkets,proto3" json:"with_markets"`
	// @inject_tag: json:"with_resolutions"
	WithResolutions bool `protobuf:"varint,2,opt,name=WithResolutions,proto3" json:"with_resolutions"`
}

func (x *ActiveBrokerageReq) Reset() {
	*x = ActiveBrokerageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveBrokerageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveBrokerageReq) ProtoMessage() {}

func (x *ActiveBrokerageReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveBrokerageReq.ProtoReflect.Descriptor instead.
func (*ActiveBrokerageReq) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{7}
}

func (x *ActiveBrokerageReq) GetWithMarkets() bool {
	if x != nil {
		return x.WithMarkets
	}
	return false
}

func (x *ActiveBrokerageReq) GetWithResolutions() bool {
	if x != nil {
		return x.WithResolutions
	}
	return false
}

var File_services_brokerage_api_src_brokerage_proto protoreflect.FileDescriptor

var file_services_brokerage_api_src_brokerage_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e,
	0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x32, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x16, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x12, 0x2e, 0x0a, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b,
	0x41, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x07, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x22, 0xf8, 0x02, 0x0a, 0x09, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x32, 0x0a,
	0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x2e, 0x0a, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69,
	0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x70,
	0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x12, 0x23,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x41, 0x0a, 0x0a, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x33, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41,
	0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x08, 0x45, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x36, 0x0a, 0x0f, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x53,
	0x0a, 0x13, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x4f, 0x48, 0x4c, 0x43, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x4f,
	0x48, 0x4c, 0x43, 0x22, 0x46, 0x0a, 0x12, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x69, 0x74,
	0x68, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x57, 0x69, 0x74, 0x68, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x60, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x57, 0x69, 0x74, 0x68, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x57, 0x69,
	0x74, 0x68, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x57, 0x69, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2a, 0x42, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x14, 0x0a, 0x10, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x6f, 0x69, 0x6e, 0x65, 0x78, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x6f, 0x62, 0x69, 0x74, 0x65, 0x78, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x10, 0x63, 0x32, 0x97, 0x03, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x43, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x12, 0x20, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x18,
	0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x50, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x68, 0x2d, 0x76, 0x61, 0x72, 0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_brokerage_api_src_brokerage_proto_rawDescOnce sync.Once
	file_services_brokerage_api_src_brokerage_proto_rawDescData = file_services_brokerage_api_src_brokerage_proto_rawDesc
)

func file_services_brokerage_api_src_brokerage_proto_rawDescGZIP() []byte {
	file_services_brokerage_api_src_brokerage_proto_rawDescOnce.Do(func() {
		file_services_brokerage_api_src_brokerage_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_brokerage_api_src_brokerage_proto_rawDescData)
	})
	return file_services_brokerage_api_src_brokerage_proto_rawDescData
}

var file_services_brokerage_api_src_brokerage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_brokerage_api_src_brokerage_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_brokerage_api_src_brokerage_proto_goTypes = []interface{}{
	(Platform)(0),               // 0: brokerageApi.Platform
	(*CreateBrokerageReq)(nil),  // 1: brokerageApi.CreateBrokerageReq
	(*Brokerage)(nil),           // 2: brokerageApi.Brokerage
	(*Brokerages)(nil),          // 3: brokerageApi.Brokerages
	(*BrokerageStatus)(nil),     // 4: brokerageApi.BrokerageStatus
	(*StatusChangeRequest)(nil), // 5: brokerageApi.StatusChangeRequest
	(*ReturnBrokerageReq)(nil),  // 6: brokerageApi.ReturnBrokerageReq
	(*DeleteBrokerageReq)(nil),  // 7: brokerageApi.DeleteBrokerageReq
	(*ActiveBrokerageReq)(nil),  // 8: brokerageApi.ActiveBrokerageReq
	(*api.Auth)(nil),            // 9: api.Auth
	(*api1.Markets)(nil),        // 10: chipmunkApi.Markets
	(*api1.Resolution)(nil),     // 11: chipmunkApi.Resolution
	(api.Status)(0),             // 12: api.Status
	(*api.Void)(nil),            // 13: api.Void
}
var file_services_brokerage_api_src_brokerage_proto_depIdxs = []int32{
	9,  // 0: brokerageApi.CreateBrokerageReq.Auth:type_name -> api.Auth
	0,  // 1: brokerageApi.CreateBrokerageReq.Platform:type_name -> brokerageApi.Platform
	10, // 2: brokerageApi.CreateBrokerageReq.Markets:type_name -> chipmunkApi.Markets
	9,  // 3: brokerageApi.Brokerage.Auth:type_name -> api.Auth
	0,  // 4: brokerageApi.Brokerage.Platform:type_name -> brokerageApi.Platform
	10, // 5: brokerageApi.Brokerage.Markets:type_name -> chipmunkApi.Markets
	11, // 6: brokerageApi.Brokerage.Resolution:type_name -> chipmunkApi.Resolution
	12, // 7: brokerageApi.Brokerage.Status:type_name -> api.Status
	2,  // 8: brokerageApi.Brokerages.Elements:type_name -> brokerageApi.Brokerage
	12, // 9: brokerageApi.BrokerageStatus.Status:type_name -> api.Status
	1,  // 10: brokerageApi.BrokerageService.Create:input_type -> brokerageApi.CreateBrokerageReq
	8,  // 11: brokerageApi.BrokerageService.Active:input_type -> brokerageApi.ActiveBrokerageReq
	6,  // 12: brokerageApi.BrokerageService.Return:input_type -> brokerageApi.ReturnBrokerageReq
	13, // 13: brokerageApi.BrokerageService.List:input_type -> api.Void
	7,  // 14: brokerageApi.BrokerageService.Delete:input_type -> brokerageApi.DeleteBrokerageReq
	5,  // 15: brokerageApi.BrokerageService.ChangeStatus:input_type -> brokerageApi.StatusChangeRequest
	2,  // 16: brokerageApi.BrokerageService.Create:output_type -> brokerageApi.Brokerage
	2,  // 17: brokerageApi.BrokerageService.Active:output_type -> brokerageApi.Brokerage
	2,  // 18: brokerageApi.BrokerageService.Return:output_type -> brokerageApi.Brokerage
	3,  // 19: brokerageApi.BrokerageService.List:output_type -> brokerageApi.Brokerages
	13, // 20: brokerageApi.BrokerageService.Delete:output_type -> api.Void
	4,  // 21: brokerageApi.BrokerageService.ChangeStatus:output_type -> brokerageApi.BrokerageStatus
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_services_brokerage_api_src_brokerage_proto_init() }
func file_services_brokerage_api_src_brokerage_proto_init() {
	if File_services_brokerage_api_src_brokerage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_brokerage_api_src_brokerage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBrokerageReq); i {
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
		file_services_brokerage_api_src_brokerage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_brokerage_api_src_brokerage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_brokerage_api_src_brokerage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrokerageStatus); i {
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
		file_services_brokerage_api_src_brokerage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_services_brokerage_api_src_brokerage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnBrokerageReq); i {
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
		file_services_brokerage_api_src_brokerage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBrokerageReq); i {
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
		file_services_brokerage_api_src_brokerage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveBrokerageReq); i {
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
			RawDescriptor: file_services_brokerage_api_src_brokerage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_brokerage_api_src_brokerage_proto_goTypes,
		DependencyIndexes: file_services_brokerage_api_src_brokerage_proto_depIdxs,
		EnumInfos:         file_services_brokerage_api_src_brokerage_proto_enumTypes,
		MessageInfos:      file_services_brokerage_api_src_brokerage_proto_msgTypes,
	}.Build()
	File_services_brokerage_api_src_brokerage_proto = out.File
	file_services_brokerage_api_src_brokerage_proto_rawDesc = nil
	file_services_brokerage_api_src_brokerage_proto_goTypes = nil
	file_services_brokerage_api_src_brokerage_proto_depIdxs = nil
}
