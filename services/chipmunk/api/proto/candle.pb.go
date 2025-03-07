// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: services/chipmunk/api/proto/src/candle.proto

package proto

import (
	proto "github.com/h-varmazyar/Gate/api/proto"
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

type Candle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,1,opt,name=UpdatedAt,proto3" json:"updated_at"`
	//	@inject_tag:	json:"created_at"
	CreatedAt int64 `protobuf:"varint,2,opt,name=CreatedAt,proto3" json:"created_at"`
	//	@inject_tag:	json:"volume"
	Volume float64 `protobuf:"fixed64,3,opt,name=Volume,proto3" json:"volume"`
	//	@inject_tag:	json:"amount"
	Amount float64 `protobuf:"fixed64,4,opt,name=Amount,proto3" json:"amount"`
	//	@inject_tag:	json:"close"
	Close float64 `protobuf:"fixed64,5,opt,name=Close,proto3" json:"close"`
	//	@inject_tag:	json:"open"
	Open float64 `protobuf:"fixed64,6,opt,name=Open,proto3" json:"open"`
	//	@inject_tag:	json:"time"
	Time int64 `protobuf:"varint,7,opt,name=Time,proto3" json:"time"`
	//	@inject_tag:	json:"high"
	High float64 `protobuf:"fixed64,8,opt,name=High,proto3" json:"high"`
	//	@inject_tag:	json:"low"
	Low float64 `protobuf:"fixed64,9,opt,name=Low,proto3" json:"low"`
	//	@inject_tag:	json:"id"
	ID uint64 `protobuf:"varint,10,opt,name=ID,proto3" json:"id"`
	//	@inject_tag:	json:"market_id"
	MarketID string `protobuf:"bytes,11,opt,name=MarketID,proto3" json:"market_id"`
	//	@inject_tag:	json:"resolution_id"
	ResolutionID string `protobuf:"bytes,12,opt,name=ResolutionID,proto3" json:"resolution_id"`
}

func (x *Candle) Reset() {
	*x = Candle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candle) ProtoMessage() {}

func (x *Candle) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candle.ProtoReflect.Descriptor instead.
func (*Candle) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_candle_proto_rawDescGZIP(), []int{0}
}

func (x *Candle) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Candle) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Candle) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Candle) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Candle) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *Candle) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Candle) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Candle) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Candle) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Candle) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Candle) GetMarketID() string {
	if x != nil {
		return x.MarketID
	}
	return ""
}

func (x *Candle) GetResolutionID() string {
	if x != nil {
		return x.ResolutionID
	}
	return ""
}

type Candles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"elements"
	Elements []*Candle `protobuf:"bytes,1,rep,name=Elements,proto3" json:"elements"`
	//	@inject_tag:	json:"count"
	Count int64 `protobuf:"varint,2,opt,name=Count,proto3" json:"count"`
}

func (x *Candles) Reset() {
	*x = Candles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candles) ProtoMessage() {}

func (x *Candles) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candles.ProtoReflect.Descriptor instead.
func (*Candles) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_candle_proto_rawDescGZIP(), []int{1}
}

func (x *Candles) GetElements() []*Candle {
	if x != nil {
		return x.Elements
	}
	return nil
}

func (x *Candles) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CandleListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"resolution_id"
	ResolutionID string `protobuf:"bytes,1,opt,name=ResolutionID,proto3" json:"resolution_id"`
	//	@inject_tag:	json:"market_id"
	MarketID string `protobuf:"bytes,2,opt,name=MarketID,proto3" json:"market_id"`
	//	@inject_tag:	json:"count"
	Count int32 `protobuf:"varint,3,opt,name=Count,proto3" json:"count"`
}

func (x *CandleListReq) Reset() {
	*x = CandleListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandleListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandleListReq) ProtoMessage() {}

func (x *CandleListReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandleListReq.ProtoReflect.Descriptor instead.
func (*CandleListReq) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_candle_proto_rawDescGZIP(), []int{2}
}

func (x *CandleListReq) GetResolutionID() string {
	if x != nil {
		return x.ResolutionID
	}
	return ""
}

func (x *CandleListReq) GetMarketID() string {
	if x != nil {
		return x.MarketID
	}
	return ""
}

func (x *CandleListReq) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CandleUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"volume"
	Volume float64 `protobuf:"fixed64,1,opt,name=Volume,proto3" json:"volume"`
	//	@inject_tag:	json:"close"
	Close float64 `protobuf:"fixed64,2,opt,name=Close,proto3" json:"close"`
	//	@inject_tag:	json:"open"
	Open float64 `protobuf:"fixed64,3,opt,name=Open,proto3" json:"open"`
	//	@inject_tag:	json:"date"
	Date int64 `protobuf:"varint,4,opt,name=Date,proto3" json:"date"`
	//	@inject_tag:	json:"high"
	High float64 `protobuf:"fixed64,5,opt,name=High,proto3" json:"high"`
	//	@inject_tag:	json:"low"
	Low float64 `protobuf:"fixed64,6,opt,name=Low,proto3" json:"low"`
	//	@inject_tag:	json:"market"
	Market string `protobuf:"bytes,7,opt,name=Market,proto3" json:"market"`
	//	@inject_tag:	json:"platform"
	Platform proto.Platform `protobuf:"varint,8,opt,name=Platform,proto3,enum=api.Platform" json:"platform"`
}

func (x *CandleUpdateReq) Reset() {
	*x = CandleUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandleUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandleUpdateReq) ProtoMessage() {}

func (x *CandleUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandleUpdateReq.ProtoReflect.Descriptor instead.
func (*CandleUpdateReq) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_candle_proto_rawDescGZIP(), []int{3}
}

func (x *CandleUpdateReq) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *CandleUpdateReq) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *CandleUpdateReq) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *CandleUpdateReq) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *CandleUpdateReq) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *CandleUpdateReq) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *CandleUpdateReq) GetMarket() string {
	if x != nil {
		return x.Market
	}
	return ""
}

func (x *CandleUpdateReq) GetPlatform() proto.Platform {
	if x != nil {
		return x.Platform
	}
	return proto.Platform(0)
}

type CandlesAsyncUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//	@inject_tag:	json:"elements"
	Candles []*Candle `protobuf:"bytes,1,rep,name=Candles,proto3" json:"elements"`
	//	@inject_tag:	json:"market_id"
	MarketID string `protobuf:"bytes,2,opt,name=MarketID,proto3" json:"market_id"`
	//	@inject_tag:	json:"reference_id"
	ReferenceID string `protobuf:"bytes,3,opt,name=ReferenceID,proto3" json:"reference_id"`
	//	@inject_tag:	json:"error"
	Error string `protobuf:"bytes,4,opt,name=Error,proto3" json:"error"`
}

func (x *CandlesAsyncUpdate) Reset() {
	*x = CandlesAsyncUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandlesAsyncUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandlesAsyncUpdate) ProtoMessage() {}

func (x *CandlesAsyncUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_services_chipmunk_api_proto_src_candle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandlesAsyncUpdate.ProtoReflect.Descriptor instead.
func (*CandlesAsyncUpdate) Descriptor() ([]byte, []int) {
	return file_services_chipmunk_api_proto_src_candle_proto_rawDescGZIP(), []int{4}
}

func (x *CandlesAsyncUpdate) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

func (x *CandlesAsyncUpdate) GetMarketID() string {
	if x != nil {
		return x.MarketID
	}
	return ""
}

func (x *CandlesAsyncUpdate) GetReferenceID() string {
	if x != nil {
		return x.ReferenceID
	}
	return ""
}

func (x *CandlesAsyncUpdate) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_services_chipmunk_api_proto_src_candle_proto protoreflect.FileDescriptor

var file_services_chipmunk_api_proto_src_candle_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x1a, 0x18, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x69,
	0x67, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x48, 0x69, 0x67, 0x68, 0x12, 0x10,
	0x0a, 0x03, 0x4c, 0x6f, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x4c, 0x6f, 0x77,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x22, 0x50, 0x0a, 0x07, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x45,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x52, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x65, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x0f, 0x43, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4f,
	0x70, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x69, 0x67, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x48, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x77, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x4c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x12, 0x29, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x97, 0x01, 0x0a,
	0x12, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41,
	0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x07, 0x43, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x20,
	0x0a, 0x0b, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x86, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1a, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x43,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x63,
	0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x63,
	0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x63, 0x68, 0x69,
	0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d,
	0x76, 0x61, 0x72, 0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_services_chipmunk_api_proto_src_candle_proto_rawDescOnce sync.Once
	file_services_chipmunk_api_proto_src_candle_proto_rawDescData = file_services_chipmunk_api_proto_src_candle_proto_rawDesc
)

func file_services_chipmunk_api_proto_src_candle_proto_rawDescGZIP() []byte {
	file_services_chipmunk_api_proto_src_candle_proto_rawDescOnce.Do(func() {
		file_services_chipmunk_api_proto_src_candle_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_chipmunk_api_proto_src_candle_proto_rawDescData)
	})
	return file_services_chipmunk_api_proto_src_candle_proto_rawDescData
}

var file_services_chipmunk_api_proto_src_candle_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_chipmunk_api_proto_src_candle_proto_goTypes = []interface{}{
	(*Candle)(nil),             // 0: chipmunkApi.Candle
	(*Candles)(nil),            // 1: chipmunkApi.Candles
	(*CandleListReq)(nil),      // 2: chipmunkApi.CandleListReq
	(*CandleUpdateReq)(nil),    // 3: chipmunkApi.CandleUpdateReq
	(*CandlesAsyncUpdate)(nil), // 4: chipmunkApi.CandlesAsyncUpdate
	(proto.Platform)(0),        // 5: api.Platform
}
var file_services_chipmunk_api_proto_src_candle_proto_depIdxs = []int32{
	0, // 0: chipmunkApi.Candles.Elements:type_name -> chipmunkApi.Candle
	5, // 1: chipmunkApi.CandleUpdateReq.Platform:type_name -> api.Platform
	0, // 2: chipmunkApi.CandlesAsyncUpdate.Candles:type_name -> chipmunkApi.Candle
	2, // 3: chipmunkApi.CandleService.List:input_type -> chipmunkApi.CandleListReq
	3, // 4: chipmunkApi.CandleService.Update:input_type -> chipmunkApi.CandleUpdateReq
	1, // 5: chipmunkApi.CandleService.List:output_type -> chipmunkApi.Candles
	0, // 6: chipmunkApi.CandleService.Update:output_type -> chipmunkApi.Candle
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_chipmunk_api_proto_src_candle_proto_init() }
func file_services_chipmunk_api_proto_src_candle_proto_init() {
	if File_services_chipmunk_api_proto_src_candle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_chipmunk_api_proto_src_candle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candle); i {
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
		file_services_chipmunk_api_proto_src_candle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candles); i {
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
		file_services_chipmunk_api_proto_src_candle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandleListReq); i {
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
		file_services_chipmunk_api_proto_src_candle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandleUpdateReq); i {
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
		file_services_chipmunk_api_proto_src_candle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandlesAsyncUpdate); i {
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
			RawDescriptor: file_services_chipmunk_api_proto_src_candle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_chipmunk_api_proto_src_candle_proto_goTypes,
		DependencyIndexes: file_services_chipmunk_api_proto_src_candle_proto_depIdxs,
		MessageInfos:      file_services_chipmunk_api_proto_src_candle_proto_msgTypes,
	}.Build()
	File_services_chipmunk_api_proto_src_candle_proto = out.File
	file_services_chipmunk_api_proto_src_candle_proto_rawDesc = nil
	file_services_chipmunk_api_proto_src_candle_proto_goTypes = nil
	file_services_chipmunk_api_proto_src_candle_proto_depIdxs = nil
}
