// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: services/eagle/api/src/strategy.proto

package api

import (
	api1 "github.com/h-varmazyar/Gate/api"
	api "github.com/h-varmazyar/Gate/services/chipmunk/api"
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

type Strategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"created_at,omitempty"
	CreatedAt int64 `protobuf:"varint,2,opt,name=CreatedAt,proto3" json:"created_at,omitempty"`
	// @inject_tag: json:"updated_at,omitempty"
	UpdatedAt int64 `protobuf:"varint,3,opt,name=UpdatedAt,proto3" json:"updated_at,omitempty"`
	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"name"`
	// @inject_tag: json:"description"
	Description string `protobuf:"bytes,5,opt,name=Description,proto3" json:"description"`
	// @inject_tag: json:"min_daily_profit_rate"
	MinDailyProfitRate float64 `protobuf:"fixed64,6,opt,name=MinDailyProfitRate,proto3" json:"min_daily_profit_rate"`
	// @inject_tag: json:"min_profit_per_trade_rate"
	MinProfitPerTradeRate float64 `protobuf:"fixed64,7,opt,name=MinProfitPerTradeRate,proto3" json:"min_profit_per_trade_rate"`
	// @inject_tag: json:"max_fund_per_trade"
	MaxFundPerTrade float64 `protobuf:"fixed64,8,opt,name=MaxFundPerTrade,proto3" json:"max_fund_per_trade"`
	// @inject_tag: json:"max_fund_per_trade_rate"
	MaxFundPerTradeRate float64 `protobuf:"fixed64,9,opt,name=MaxFundPerTradeRate,proto3" json:"max_fund_per_trade_rate"`
	// @inject_tag: json:"working_resolution"
	WorkingResolution *api.Resolution `protobuf:"bytes,10,opt,name=WorkingResolution,proto3" json:"working_resolution"`
	// @inject_tag: json:"indicators"
	Indicators []*StrategyIndicator `protobuf:"bytes,11,rep,name=Indicators,proto3" json:"indicators"`
}

func (x *Strategy) Reset() {
	*x = Strategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strategy) ProtoMessage() {}

func (x *Strategy) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strategy.ProtoReflect.Descriptor instead.
func (*Strategy) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *Strategy) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Strategy) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Strategy) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Strategy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Strategy) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Strategy) GetMinDailyProfitRate() float64 {
	if x != nil {
		return x.MinDailyProfitRate
	}
	return 0
}

func (x *Strategy) GetMinProfitPerTradeRate() float64 {
	if x != nil {
		return x.MinProfitPerTradeRate
	}
	return 0
}

func (x *Strategy) GetMaxFundPerTrade() float64 {
	if x != nil {
		return x.MaxFundPerTrade
	}
	return 0
}

func (x *Strategy) GetMaxFundPerTradeRate() float64 {
	if x != nil {
		return x.MaxFundPerTradeRate
	}
	return 0
}

func (x *Strategy) GetWorkingResolution() *api.Resolution {
	if x != nil {
		return x.WorkingResolution
	}
	return nil
}

func (x *Strategy) GetIndicators() []*StrategyIndicator {
	if x != nil {
		return x.Indicators
	}
	return nil
}

type Strategies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"elements"
	Elements []*Strategy `protobuf:"bytes,1,rep,name=Elements,proto3" json:"elements"`
}

func (x *Strategies) Reset() {
	*x = Strategies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_strategy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Strategies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Strategies) ProtoMessage() {}

func (x *Strategies) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_strategy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Strategies.ProtoReflect.Descriptor instead.
func (*Strategies) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_strategy_proto_rawDescGZIP(), []int{1}
}

func (x *Strategies) GetElements() []*Strategy {
	if x != nil {
		return x.Elements
	}
	return nil
}

type CreateStrategyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"name"
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"name"`
	// @inject_tag: json:"description"
	Description string `protobuf:"bytes,2,opt,name=Description,proto3" json:"description"`
	// @inject_tag: json:"min_daily_profit_rate"
	MinDailyProfitRate string `protobuf:"bytes,3,opt,name=MinDailyProfitRate,proto3" json:"min_daily_profit_rate"`
	// @inject_tag: json:"min_profit_per_trade_rate"
	MinProfitPerTradeRate string `protobuf:"bytes,4,opt,name=MinProfitPerTradeRate,proto3" json:"min_profit_per_trade_rate"`
	// @inject_tag: json:"max_fund_per_trade"
	MaxFundPerTrade string `protobuf:"bytes,5,opt,name=MaxFundPerTrade,proto3" json:"max_fund_per_trade"`
	// @inject_tag: json:"indicators"
	Indicators []*StrategyIndicator `protobuf:"bytes,6,rep,name=Indicators,proto3" json:"indicators"`
}

func (x *CreateStrategyReq) Reset() {
	*x = CreateStrategyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_strategy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStrategyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStrategyReq) ProtoMessage() {}

func (x *CreateStrategyReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_strategy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStrategyReq.ProtoReflect.Descriptor instead.
func (*CreateStrategyReq) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_strategy_proto_rawDescGZIP(), []int{2}
}

func (x *CreateStrategyReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStrategyReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateStrategyReq) GetMinDailyProfitRate() string {
	if x != nil {
		return x.MinDailyProfitRate
	}
	return ""
}

func (x *CreateStrategyReq) GetMinProfitPerTradeRate() string {
	if x != nil {
		return x.MinProfitPerTradeRate
	}
	return ""
}

func (x *CreateStrategyReq) GetMaxFundPerTrade() string {
	if x != nil {
		return x.MaxFundPerTrade
	}
	return ""
}

func (x *CreateStrategyReq) GetIndicators() []*StrategyIndicator {
	if x != nil {
		return x.Indicators
	}
	return nil
}

type StrategyIndicatorReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"strategy_id"
	StrategyID string `protobuf:"bytes,1,opt,name=StrategyID,proto3" json:"strategy_id"`
}

func (x *StrategyIndicatorReq) Reset() {
	*x = StrategyIndicatorReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_strategy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyIndicatorReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyIndicatorReq) ProtoMessage() {}

func (x *StrategyIndicatorReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_strategy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyIndicatorReq.ProtoReflect.Descriptor instead.
func (*StrategyIndicatorReq) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_strategy_proto_rawDescGZIP(), []int{3}
}

func (x *StrategyIndicatorReq) GetStrategyID() string {
	if x != nil {
		return x.StrategyID
	}
	return ""
}

type StrategyIndicator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"strategy_id"
	StrategyID string `protobuf:"bytes,1,opt,name=StrategyID,proto3" json:"strategy_id"`
	// @inject_tag: json:"indicator_id"
	IndicatorID string `protobuf:"bytes,2,opt,name=IndicatorID,proto3" json:"indicator_id"`
	// @inject_tag: json:"type"
	Type api.IndicatorType `protobuf:"varint,3,opt,name=Type,proto3,enum=chipmunkApi.IndicatorType" json:"type"`
}

func (x *StrategyIndicator) Reset() {
	*x = StrategyIndicator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_strategy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyIndicator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyIndicator) ProtoMessage() {}

func (x *StrategyIndicator) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_strategy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyIndicator.ProtoReflect.Descriptor instead.
func (*StrategyIndicator) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_strategy_proto_rawDescGZIP(), []int{4}
}

func (x *StrategyIndicator) GetStrategyID() string {
	if x != nil {
		return x.StrategyID
	}
	return ""
}

func (x *StrategyIndicator) GetIndicatorID() string {
	if x != nil {
		return x.IndicatorID
	}
	return ""
}

func (x *StrategyIndicator) GetType() api.IndicatorType {
	if x != nil {
		return x.Type
	}
	return api.IndicatorType(0)
}

type StrategyIndicators struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"elements"
	Elements []*StrategyIndicator `protobuf:"bytes,1,rep,name=Elements,proto3" json:"elements"`
}

func (x *StrategyIndicators) Reset() {
	*x = StrategyIndicators{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_strategy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyIndicators) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyIndicators) ProtoMessage() {}

func (x *StrategyIndicators) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_strategy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyIndicators.ProtoReflect.Descriptor instead.
func (*StrategyIndicators) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_strategy_proto_rawDescGZIP(), []int{5}
}

func (x *StrategyIndicators) GetElements() []*StrategyIndicator {
	if x != nil {
		return x.Elements
	}
	return nil
}

type ReturnStrategyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
}

func (x *ReturnStrategyReq) Reset() {
	*x = ReturnStrategyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_strategy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnStrategyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnStrategyReq) ProtoMessage() {}

func (x *ReturnStrategyReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_strategy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnStrategyReq.ProtoReflect.Descriptor instead.
func (*ReturnStrategyReq) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_strategy_proto_rawDescGZIP(), []int{6}
}

func (x *ReturnStrategyReq) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

var File_services_eagle_api_src_strategy_proto protoreflect.FileDescriptor

var file_services_eagle_api_src_strategy_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x61, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70,
	0x69, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x70,
	0x6d, 0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x69, 0x6e, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x03, 0x0a,
	0x08, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x4d,
	0x69, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x4d, 0x69, 0x6e, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x4d,
	0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x50, 0x65, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x4d, 0x69, 0x6e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x74, 0x50, 0x65, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x61, 0x78, 0x46, 0x75, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x4d, 0x61, 0x78, 0x46,
	0x75, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x4d,
	0x61, 0x78, 0x46, 0x75, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x13, 0x4d, 0x61, 0x78, 0x46, 0x75, 0x6e,
	0x64, 0x50, 0x65, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x45, 0x0a,
	0x11, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x11, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65,
	0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x6e, 0x64, 0x69,
	0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x73, 0x22, 0x3c, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x12,
	0x2e, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x96, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x4d,
	0x69, 0x6e, 0x44, 0x61, 0x69, 0x6c, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x4d, 0x69, 0x6e, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x4d,
	0x69, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x50, 0x65, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x4d, 0x69, 0x6e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x74, 0x50, 0x65, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x61, 0x78, 0x46, 0x75, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4d, 0x61, 0x78, 0x46,
	0x75, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x49, 0x6e,
	0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x36, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x44,
	0x22, 0x86, 0x01, 0x0a, 0x11, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e,
	0x6b, 0x41, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4d, 0x0a, 0x12, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12,
	0x37, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x32, 0xfc, 0x01,
	0x0a, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x65, 0x61,
	0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65,
	0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x39, 0x0a, 0x06,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12, 0x1b, 0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x14, 0x2e, 0x65, 0x61, 0x67,
	0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x12, 0x4a, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x1e,
	0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c,
	0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72,
	0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_eagle_api_src_strategy_proto_rawDescOnce sync.Once
	file_services_eagle_api_src_strategy_proto_rawDescData = file_services_eagle_api_src_strategy_proto_rawDesc
)

func file_services_eagle_api_src_strategy_proto_rawDescGZIP() []byte {
	file_services_eagle_api_src_strategy_proto_rawDescOnce.Do(func() {
		file_services_eagle_api_src_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_eagle_api_src_strategy_proto_rawDescData)
	})
	return file_services_eagle_api_src_strategy_proto_rawDescData
}

var file_services_eagle_api_src_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_services_eagle_api_src_strategy_proto_goTypes = []interface{}{
	(*Strategy)(nil),             // 0: eagleApi.Strategy
	(*Strategies)(nil),           // 1: eagleApi.Strategies
	(*CreateStrategyReq)(nil),    // 2: eagleApi.CreateStrategyReq
	(*StrategyIndicatorReq)(nil), // 3: eagleApi.StrategyIndicatorReq
	(*StrategyIndicator)(nil),    // 4: eagleApi.StrategyIndicator
	(*StrategyIndicators)(nil),   // 5: eagleApi.StrategyIndicators
	(*ReturnStrategyReq)(nil),    // 6: eagleApi.ReturnStrategyReq
	(*api.Resolution)(nil),       // 7: chipmunkApi.Resolution
	(api.IndicatorType)(0),       // 8: chipmunkApi.Indicator.type
	(*api1.Void)(nil),            // 9: api.Void
}
var file_services_eagle_api_src_strategy_proto_depIdxs = []int32{
	7,  // 0: eagleApi.Strategy.WorkingResolution:type_name -> chipmunkApi.Resolution
	4,  // 1: eagleApi.Strategy.Indicators:type_name -> eagleApi.StrategyIndicator
	0,  // 2: eagleApi.Strategies.Elements:type_name -> eagleApi.Strategy
	4,  // 3: eagleApi.CreateStrategyReq.Indicators:type_name -> eagleApi.StrategyIndicator
	8,  // 4: eagleApi.StrategyIndicator.Type:type_name -> chipmunkApi.Indicator.type
	4,  // 5: eagleApi.StrategyIndicators.Elements:type_name -> eagleApi.StrategyIndicator
	2,  // 6: eagleApi.StrategyService.Create:input_type -> eagleApi.CreateStrategyReq
	6,  // 7: eagleApi.StrategyService.Return:input_type -> eagleApi.ReturnStrategyReq
	9,  // 8: eagleApi.StrategyService.List:input_type -> api.Void
	3,  // 9: eagleApi.StrategyService.Indicators:input_type -> eagleApi.StrategyIndicatorReq
	0,  // 10: eagleApi.StrategyService.Create:output_type -> eagleApi.Strategy
	0,  // 11: eagleApi.StrategyService.Return:output_type -> eagleApi.Strategy
	1,  // 12: eagleApi.StrategyService.List:output_type -> eagleApi.Strategies
	5,  // 13: eagleApi.StrategyService.Indicators:output_type -> eagleApi.StrategyIndicators
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_services_eagle_api_src_strategy_proto_init() }
func file_services_eagle_api_src_strategy_proto_init() {
	if File_services_eagle_api_src_strategy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_eagle_api_src_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strategy); i {
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
		file_services_eagle_api_src_strategy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Strategies); i {
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
		file_services_eagle_api_src_strategy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStrategyReq); i {
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
		file_services_eagle_api_src_strategy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyIndicatorReq); i {
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
		file_services_eagle_api_src_strategy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyIndicator); i {
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
		file_services_eagle_api_src_strategy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyIndicators); i {
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
		file_services_eagle_api_src_strategy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnStrategyReq); i {
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
			RawDescriptor: file_services_eagle_api_src_strategy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_eagle_api_src_strategy_proto_goTypes,
		DependencyIndexes: file_services_eagle_api_src_strategy_proto_depIdxs,
		MessageInfos:      file_services_eagle_api_src_strategy_proto_msgTypes,
	}.Build()
	File_services_eagle_api_src_strategy_proto = out.File
	file_services_eagle_api_src_strategy_proto_rawDesc = nil
	file_services_eagle_api_src_strategy_proto_goTypes = nil
	file_services_eagle_api_src_strategy_proto_depIdxs = nil
}
