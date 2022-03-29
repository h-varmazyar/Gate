// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: services/brokerage/api/src/brokerage.proto

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

type Names int32

const (
	Names_UnknownBrokerage Names = 0
	Names_Coinex           Names = 1
	Names_Nobitex          Names = 2
	Names_All              Names = 99
)

// Enum value maps for Names.
var (
	Names_name = map[int32]string{
		0:  "UnknownBrokerage",
		1:  "Coinex",
		2:  "Nobitex",
		99: "All",
	}
	Names_value = map[string]int32{
		"UnknownBrokerage": 0,
		"Coinex":           1,
		"Nobitex":          2,
		"All":              99,
	}
)

func (x Names) Enum() *Names {
	p := new(Names)
	*p = x
	return p
}

func (x Names) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Names) Descriptor() protoreflect.EnumDescriptor {
	return file_services_brokerage_api_src_brokerage_proto_enumTypes[0].Descriptor()
}

func (Names) Type() protoreflect.EnumType {
	return &file_services_brokerage_api_src_brokerage_proto_enumTypes[0]
}

func (x Names) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Names.Descriptor instead.
func (Names) EnumDescriptor() ([]byte, []int) {
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
	// @inject_tag: json:"name"
	Name Names `protobuf:"varint,4,opt,name=Name,proto3,enum=brokerageApi.Names" json:"name"`
	// @inject_tag: json:"markets"
	Markets *Markets `protobuf:"bytes,5,opt,name=Markets,proto3" json:"markets"`
	// @inject_tag: json:"resolution_id"
	ResolutionID string `protobuf:"bytes,6,opt,name=ResolutionID,proto3" json:"resolution_id"`
	// @inject_tag: json:"strategy_id"
	StrategyID uint32 `protobuf:"varint,7,opt,name=StrategyID,proto3" json:"strategy_id"`
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

func (x *CreateBrokerageReq) GetName() Names {
	if x != nil {
		return x.Name
	}
	return Names_UnknownBrokerage
}

func (x *CreateBrokerageReq) GetMarkets() *Markets {
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

func (x *CreateBrokerageReq) GetStrategyID() uint32 {
	if x != nil {
		return x.StrategyID
	}
	return 0
}

type Brokerage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"title"
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"title"`
	// @inject_tag: json:"description"
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"description"`
	// @inject_tag: json:"auth"
	Auth *api.Auth `protobuf:"bytes,4,opt,name=Auth,proto3" json:"auth"`
	// @inject_tag: json:"name"
	Name Names `protobuf:"varint,5,opt,name=Name,proto3,enum=brokerageApi.Names" json:"name"`
	// @inject_tag: json:"markets"
	Markets *Markets `protobuf:"bytes,6,opt,name=Markets,proto3" json:"markets"`
	// @inject_tag: json:"resolution_id"
	ResolutionID string `protobuf:"bytes,7,opt,name=ResolutionID,proto3" json:"resolution_id"`
	// @inject_tag: json:"resolution"
	Resolution *Resolution `protobuf:"bytes,8,opt,name=Resolution,proto3" json:"resolution"`
	// @inject_tag: json:"strategy_id"
	StrategyID uint32 `protobuf:"varint,9,opt,name=StrategyID,proto3" json:"strategy_id"`
	// @inject_tag: json:"status"
	Status api.Status `protobuf:"varint,10,opt,name=Status,proto3,enum=api.Status" json:"status"`
	// @inject_tag: json:"strategy"
	Strategy *Strategy `protobuf:"bytes,11,opt,name=Strategy,proto3" json:"strategy"`
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

func (x *Brokerage) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
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

func (x *Brokerage) GetName() Names {
	if x != nil {
		return x.Name
	}
	return Names_UnknownBrokerage
}

func (x *Brokerage) GetMarkets() *Markets {
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

func (x *Brokerage) GetResolution() *Resolution {
	if x != nil {
		return x.Resolution
	}
	return nil
}

func (x *Brokerage) GetStrategyID() uint32 {
	if x != nil {
		return x.StrategyID
	}
	return 0
}

func (x *Brokerage) GetStatus() api.Status {
	if x != nil {
		return x.Status
	}
	return api.Status_Disable
}

func (x *Brokerage) GetStrategy() *Strategy {
	if x != nil {
		return x.Strategy
	}
	return nil
}

type Brokerages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"brokerages"
	Brokerages []*Brokerage `protobuf:"bytes,1,rep,name=Brokerages,proto3" json:"brokerages"`
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

func (x *Brokerages) GetBrokerages() []*Brokerage {
	if x != nil {
		return x.Brokerages
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
	return api.Status_Disable
}

type StatusChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
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

func (x *StatusChangeRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
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

type BrokerageIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
}

func (x *BrokerageIDReq) Reset() {
	*x = BrokerageIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_brokerage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrokerageIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerageIDReq) ProtoMessage() {}

func (x *BrokerageIDReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BrokerageIDReq.ProtoReflect.Descriptor instead.
func (*BrokerageIDReq) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_brokerage_proto_rawDescGZIP(), []int{5}
}

func (x *BrokerageIDReq) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_services_brokerage_api_src_brokerage_proto protoreflect.FileDescriptor

var file_services_brokerage_api_src_brokerage_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x89, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x27, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x07,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x22, 0xa3, 0x03, 0x0a, 0x09,
	0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1d, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x27, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x73, 0x52, 0x07, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x38,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a,
	0x08, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x08, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x22, 0x45, 0x0a, 0x0a, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x12,
	0x37, 0x0a, 0x0a, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41,
	0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0a, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x0f, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x53, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x54, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x48, 0x4c, 0x43, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x4f, 0x48, 0x4c, 0x43, 0x22, 0x20, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x2a, 0x3f, 0x0a, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x10, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x42, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x6f, 0x69, 0x6e, 0x65, 0x78,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x6f, 0x62, 0x69, 0x74, 0x65, 0x78, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x10, 0x63, 0x32, 0x8d, 0x03, 0x0a, 0x10, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x2b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x1a, 0x18, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41,
	0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x73, 0x12, 0x44, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1c, 0x2e, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x50, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72, 0x6d, 0x61, 0x7a, 0x79,
	0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_services_brokerage_api_src_brokerage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_brokerage_api_src_brokerage_proto_goTypes = []interface{}{
	(Names)(0),                  // 0: brokerageApi.Names
	(*CreateBrokerageReq)(nil),  // 1: brokerageApi.CreateBrokerageReq
	(*Brokerage)(nil),           // 2: brokerageApi.Brokerage
	(*Brokerages)(nil),          // 3: brokerageApi.Brokerages
	(*BrokerageStatus)(nil),     // 4: brokerageApi.BrokerageStatus
	(*StatusChangeRequest)(nil), // 5: brokerageApi.StatusChangeRequest
	(*BrokerageIDReq)(nil),      // 6: brokerageApi.BrokerageIDReq
	(*api.Auth)(nil),            // 7: api.Auth
	(*Markets)(nil),             // 8: brokerageApi.Markets
	(*Resolution)(nil),          // 9: brokerageApi.Resolution
	(api.Status)(0),             // 10: api.Status
	(*Strategy)(nil),            // 11: brokerageApi.Strategy
	(*api.Void)(nil),            // 12: api.Void
}
var file_services_brokerage_api_src_brokerage_proto_depIdxs = []int32{
	7,  // 0: brokerageApi.CreateBrokerageReq.Auth:type_name -> api.Auth
	0,  // 1: brokerageApi.CreateBrokerageReq.Name:type_name -> brokerageApi.Names
	8,  // 2: brokerageApi.CreateBrokerageReq.Markets:type_name -> brokerageApi.Markets
	7,  // 3: brokerageApi.Brokerage.Auth:type_name -> api.Auth
	0,  // 4: brokerageApi.Brokerage.Name:type_name -> brokerageApi.Names
	8,  // 5: brokerageApi.Brokerage.Markets:type_name -> brokerageApi.Markets
	9,  // 6: brokerageApi.Brokerage.Resolution:type_name -> brokerageApi.Resolution
	10, // 7: brokerageApi.Brokerage.Status:type_name -> api.Status
	11, // 8: brokerageApi.Brokerage.Strategy:type_name -> brokerageApi.Strategy
	2,  // 9: brokerageApi.Brokerages.Brokerages:type_name -> brokerageApi.Brokerage
	10, // 10: brokerageApi.BrokerageStatus.Status:type_name -> api.Status
	1,  // 11: brokerageApi.BrokerageService.Create:input_type -> brokerageApi.CreateBrokerageReq
	6,  // 12: brokerageApi.BrokerageService.Get:input_type -> brokerageApi.BrokerageIDReq
	12, // 13: brokerageApi.BrokerageService.List:input_type -> api.Void
	6,  // 14: brokerageApi.BrokerageService.GetInternal:input_type -> brokerageApi.BrokerageIDReq
	6,  // 15: brokerageApi.BrokerageService.Delete:input_type -> brokerageApi.BrokerageIDReq
	5,  // 16: brokerageApi.BrokerageService.ChangeStatus:input_type -> brokerageApi.StatusChangeRequest
	2,  // 17: brokerageApi.BrokerageService.Create:output_type -> brokerageApi.Brokerage
	2,  // 18: brokerageApi.BrokerageService.Get:output_type -> brokerageApi.Brokerage
	3,  // 19: brokerageApi.BrokerageService.List:output_type -> brokerageApi.Brokerages
	2,  // 20: brokerageApi.BrokerageService.GetInternal:output_type -> brokerageApi.Brokerage
	12, // 21: brokerageApi.BrokerageService.Delete:output_type -> api.Void
	4,  // 22: brokerageApi.BrokerageService.ChangeStatus:output_type -> brokerageApi.BrokerageStatus
	17, // [17:23] is the sub-list for method output_type
	11, // [11:17] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_services_brokerage_api_src_brokerage_proto_init() }
func file_services_brokerage_api_src_brokerage_proto_init() {
	if File_services_brokerage_api_src_brokerage_proto != nil {
		return
	}
	file_services_brokerage_api_src_market_proto_init()
	file_services_brokerage_api_src_resolution_proto_init()
	file_services_brokerage_api_src_strategy_proto_init()
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
			switch v := v.(*BrokerageIDReq); i {
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
			NumMessages:   6,
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
