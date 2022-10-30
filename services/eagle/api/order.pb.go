// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: services/eagle/api/src/order.proto

package api

import (
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

type OrderModel int32

const (
	OrderModel_all           OrderModel = 0
	OrderModel_market        OrderModel = 1
	OrderModel_multipleLimit OrderModel = 2
	OrderModel_ioc           OrderModel = 3
	OrderModel_stopLimit     OrderModel = 4
	OrderModel_limit         OrderModel = 5
)

// Enum value maps for OrderModel.
var (
	OrderModel_name = map[int32]string{
		0: "all",
		1: "market",
		2: "multipleLimit",
		3: "ioc",
		4: "stopLimit",
		5: "limit",
	}
	OrderModel_value = map[string]int32{
		"all":           0,
		"market":        1,
		"multipleLimit": 2,
		"ioc":           3,
		"stopLimit":     4,
		"limit":         5,
	}
)

func (x OrderModel) Enum() *OrderModel {
	p := new(OrderModel)
	*p = x
	return p
}

func (x OrderModel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderModel) Descriptor() protoreflect.EnumDescriptor {
	return file_services_eagle_api_src_order_proto_enumTypes[0].Descriptor()
}

func (OrderModel) Type() protoreflect.EnumType {
	return &file_services_eagle_api_src_order_proto_enumTypes[0]
}

func (x OrderModel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderModel.Descriptor instead.
func (OrderModel) EnumDescriptor() ([]byte, []int) {
	return file_services_eagle_api_src_order_proto_rawDescGZIP(), []int{0}
}

type OrderType int32

const (
	Order_ask  OrderType = 0
	Order_bid  OrderType = 1
	Order_buy  OrderType = 2
	Order_sell OrderType = 3
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0: "ask",
		1: "bid",
		2: "buy",
		3: "sell",
	}
	OrderType_value = map[string]int32{
		"ask":  0,
		"bid":  1,
		"buy":  2,
		"sell": 3,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_services_eagle_api_src_order_proto_enumTypes[1].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_services_eagle_api_src_order_proto_enumTypes[1]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_services_eagle_api_src_order_proto_rawDescGZIP(), []int{0, 0}
}

type OrderStatus int32

const (
	Order_new       OrderStatus = 0
	Order_done      OrderStatus = 1
	Order_part_deal OrderStatus = 2
	Order_not_deal  OrderStatus = 3
	Order_cancel    OrderStatus = 4
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "new",
		1: "done",
		2: "part_deal",
		3: "not_deal",
		4: "cancel",
	}
	OrderStatus_value = map[string]int32{
		"new":       0,
		"done":      1,
		"part_deal": 2,
		"not_deal":  3,
		"cancel":    4,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_services_eagle_api_src_order_proto_enumTypes[2].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_services_eagle_api_src_order_proto_enumTypes[2]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_services_eagle_api_src_order_proto_rawDescGZIP(), []int{0, 1}
}

type OrderOption int32

const (
	Order_IOC        OrderOption = 0
	Order_FOK        OrderOption = 1
	Order_NORMAL     OrderOption = 2
	Order_MAKER_ONLY OrderOption = 3
)

// Enum value maps for OrderOption.
var (
	OrderOption_name = map[int32]string{
		0: "IOC",
		1: "FOK",
		2: "NORMAL",
		3: "MAKER_ONLY",
	}
	OrderOption_value = map[string]int32{
		"IOC":        0,
		"FOK":        1,
		"NORMAL":     2,
		"MAKER_ONLY": 3,
	}
)

func (x OrderOption) Enum() *OrderOption {
	p := new(OrderOption)
	*p = x
	return p
}

func (x OrderOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderOption) Descriptor() protoreflect.EnumDescriptor {
	return file_services_eagle_api_src_order_proto_enumTypes[3].Descriptor()
}

func (OrderOption) Type() protoreflect.EnumType {
	return &file_services_eagle_api_src_order_proto_enumTypes[3]
}

func (x OrderOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderOption.Descriptor instead.
func (OrderOption) EnumDescriptor() ([]byte, []int) {
	return file_services_eagle_api_src_order_proto_rawDescGZIP(), []int{0, 2}
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"server_order_id"
	ServerOrderId int64 `protobuf:"varint,2,opt,name=ServerOrderId,proto3" json:"server_order_id"`
	// @inject_tag: json:"amount"
	Amount float64 `protobuf:"fixed64,3,opt,name=Amount,proto3" json:"amount"`
	// @inject_tag: json:"finished_at"
	FinishedAt int64 `protobuf:"varint,4,opt,name=FinishedAt,proto3" json:"finished_at"`
	// @inject_tag: json:"executed_amount"
	ExecutedAmount float64 `protobuf:"fixed64,5,opt,name=ExecutedAmount,proto3" json:"executed_amount"`
	// @inject_tag: json:"un_executed_amount"
	UnExecutedAmount float64 `protobuf:"fixed64,6,opt,name=UnExecutedAmount,proto3" json:"un_executed_amount"`
	// @inject_tag: json:"status"
	Status OrderStatus `protobuf:"varint,7,opt,name=Status,proto3,enum=eagleApi.OrderStatus" json:"status"`
	// @inject_tag: json:"price"
	Price float64 `protobuf:"fixed64,8,opt,name=Price,proto3" json:"price"`
	// @inject_tag: json:"market"
	Market *api.Market `protobuf:"bytes,9,opt,name=Market,proto3" json:"market"`
	// @inject_tag: json:"maker_fee_rate"
	MakerFeeRate float64 `protobuf:"fixed64,10,opt,name=MakerFeeRate,proto3" json:"maker_fee_rate"`
	// @inject_tag: json:"taker_fee_rate"
	TakerFeeRate float64 `protobuf:"fixed64,11,opt,name=TakerFeeRate,proto3" json:"taker_fee_rate"`
	// @inject_tag: json:"sell_or_buy"
	SellOrBuy OrderType `protobuf:"varint,12,opt,name=SellOrBuy,proto3,enum=eagleApi.OrderType" json:"sell_or_buy"`
	// @inject_tag: json:"order_type"
	OrderType OrderModel `protobuf:"varint,13,opt,name=OrderType,proto3,enum=eagleApi.OrderModel" json:"order_type"`
	// @inject_tag: json:"average_price"
	AveragePrice float64 `protobuf:"fixed64,14,opt,name=AveragePrice,proto3" json:"average_price"`
	// @inject_tag: json:"transaction_fee"
	TransactionFee float64 `protobuf:"fixed64,15,opt,name=TransactionFee,proto3" json:"transaction_fee"`
	// @inject_tag: json:"source_asset"
	SourceAsset *api.Asset `protobuf:"bytes,16,opt,name=SourceAsset,proto3" json:"source_asset"`
	// @inject_tag: json:"destination_asset"
	DestinationAsset *api.Asset `protobuf:"bytes,17,opt,name=DestinationAsset,proto3" json:"destination_asset"`
	// @inject_tag: json:"fee_asset"
	FeeAsset *api.Asset `protobuf:"bytes,18,opt,name=FeeAsset,proto3" json:"fee_asset"`
	// @inject_tag: json:"fee_discount"
	FeeDiscount float64 `protobuf:"fixed64,19,opt,name=FeeDiscount,proto3" json:"fee_discount"`
	// @inject_tag: json:"asset_fee"
	AssetFee float64 `protobuf:"fixed64,20,opt,name=AssetFee,proto3" json:"asset_fee"`
	// @inject_tag: json:"money_fee"
	MoneyFee float64 `protobuf:"fixed64,21,opt,name=MoneyFee,proto3" json:"money_fee"`
	// @inject_tag: json:"volume"
	Volume float64 `protobuf:"fixed64,22,opt,name=Volume,proto3" json:"volume"`
	// @inject_tag: json:"order_no"
	OrderNo int64 `protobuf:"varint,23,opt,name=OrderNo,proto3" json:"order_no"`
	// @inject_tag: json:"stock_fee"
	StockFee float64 `protobuf:"fixed64,24,opt,name=StockFee,proto3" json:"stock_fee"`
	// @inject_tag: json:"created_at"
	CreatedAt int64 `protobuf:"varint,25,opt,name=CreatedAt,proto3" json:"created_at"`
	// @inject_tag: json:"updated_at"
	UpdatedAt int64 `protobuf:"varint,26,opt,name=UpdatedAt,proto3" json:"updated_at"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Order) GetServerOrderId() int64 {
	if x != nil {
		return x.ServerOrderId
	}
	return 0
}

func (x *Order) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order) GetFinishedAt() int64 {
	if x != nil {
		return x.FinishedAt
	}
	return 0
}

func (x *Order) GetExecutedAmount() float64 {
	if x != nil {
		return x.ExecutedAmount
	}
	return 0
}

func (x *Order) GetUnExecutedAmount() float64 {
	if x != nil {
		return x.UnExecutedAmount
	}
	return 0
}

func (x *Order) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return Order_new
}

func (x *Order) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetMarket() *api.Market {
	if x != nil {
		return x.Market
	}
	return nil
}

func (x *Order) GetMakerFeeRate() float64 {
	if x != nil {
		return x.MakerFeeRate
	}
	return 0
}

func (x *Order) GetTakerFeeRate() float64 {
	if x != nil {
		return x.TakerFeeRate
	}
	return 0
}

func (x *Order) GetSellOrBuy() OrderType {
	if x != nil {
		return x.SellOrBuy
	}
	return Order_ask
}

func (x *Order) GetOrderType() OrderModel {
	if x != nil {
		return x.OrderType
	}
	return OrderModel_all
}

func (x *Order) GetAveragePrice() float64 {
	if x != nil {
		return x.AveragePrice
	}
	return 0
}

func (x *Order) GetTransactionFee() float64 {
	if x != nil {
		return x.TransactionFee
	}
	return 0
}

func (x *Order) GetSourceAsset() *api.Asset {
	if x != nil {
		return x.SourceAsset
	}
	return nil
}

func (x *Order) GetDestinationAsset() *api.Asset {
	if x != nil {
		return x.DestinationAsset
	}
	return nil
}

func (x *Order) GetFeeAsset() *api.Asset {
	if x != nil {
		return x.FeeAsset
	}
	return nil
}

func (x *Order) GetFeeDiscount() float64 {
	if x != nil {
		return x.FeeDiscount
	}
	return 0
}

func (x *Order) GetAssetFee() float64 {
	if x != nil {
		return x.AssetFee
	}
	return 0
}

func (x *Order) GetMoneyFee() float64 {
	if x != nil {
		return x.MoneyFee
	}
	return 0
}

func (x *Order) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Order) GetOrderNo() int64 {
	if x != nil {
		return x.OrderNo
	}
	return 0
}

func (x *Order) GetStockFee() float64 {
	if x != nil {
		return x.StockFee
	}
	return 0
}

func (x *Order) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Order) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"elements"
	Elements []*Order `protobuf:"bytes,1,rep,name=Elements,proto3" json:"elements"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_order_proto_rawDescGZIP(), []int{1}
}

func (x *Orders) GetElements() []*Order {
	if x != nil {
		return x.Elements
	}
	return nil
}

type OrderListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderListRequest) Reset() {
	*x = OrderListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_eagle_api_src_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListRequest) ProtoMessage() {}

func (x *OrderListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_eagle_api_src_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListRequest.ProtoReflect.Descriptor instead.
func (*OrderListRequest) Descriptor() ([]byte, []int) {
	return file_services_eagle_api_src_order_proto_rawDescGZIP(), []int{2}
}

var File_services_eagle_api_src_order_proto protoreflect.FileDescriptor

var file_services_eagle_api_src_order_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x61, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x1a, 0x26,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e,
	0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x08,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a,
	0x10, 0x55, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x55, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x65, 0x61, 0x67, 0x6c,
	0x65, 0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x4d, 0x61, 0x6b, 0x65, 0x72, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x4d, 0x61, 0x6b, 0x65, 0x72, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x6b, 0x65, 0x72, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x54, 0x61, 0x6b, 0x65, 0x72, 0x46, 0x65, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x42, 0x75,
	0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41,
	0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x52, 0x09, 0x53,
	0x65, 0x6c, 0x6c, 0x4f, 0x72, 0x42, 0x75, 0x79, 0x12, 0x32, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x65, 0x61,
	0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x65, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x52, 0x0b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x3e,
	0x0a, 0x10, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d,
	0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x10, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x2e,
	0x0a, 0x08, 0x46, 0x65, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x68, 0x69, 0x70, 0x6d, 0x75, 0x6e, 0x6b, 0x41, 0x70, 0x69, 0x2e, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x52, 0x08, 0x46, 0x65, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x46, 0x65, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x46, 0x65, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x65, 0x65, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x41, 0x73, 0x73, 0x65, 0x74, 0x46, 0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x46, 0x65, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x46, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x61, 0x73,
	0x6b, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03,
	0x62, 0x75, 0x79, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x73, 0x65, 0x6c, 0x6c, 0x10, 0x03, 0x22,
	0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x65, 0x77,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x65, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x6e,
	0x6f, 0x74, 0x5f, 0x64, 0x65, 0x61, 0x6c, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x10, 0x04, 0x22, 0x36, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x07, 0x0a, 0x03, 0x49, 0x4f, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x4f, 0x4b, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x4d, 0x41, 0x4b, 0x45, 0x52, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x03, 0x22, 0x35, 0x0a,
	0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x61, 0x67, 0x6c,
	0x65, 0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x08, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0x57, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x07, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x69, 0x6f, 0x63, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x70, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x10,
	0x05, 0x32, 0x49, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a,
	0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65, 0x61, 0x67,
	0x6c, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x2d, 0x76, 0x61, 0x72,
	0x6d, 0x61, 0x7a, 0x79, 0x61, 0x72, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_eagle_api_src_order_proto_rawDescOnce sync.Once
	file_services_eagle_api_src_order_proto_rawDescData = file_services_eagle_api_src_order_proto_rawDesc
)

func file_services_eagle_api_src_order_proto_rawDescGZIP() []byte {
	file_services_eagle_api_src_order_proto_rawDescOnce.Do(func() {
		file_services_eagle_api_src_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_eagle_api_src_order_proto_rawDescData)
	})
	return file_services_eagle_api_src_order_proto_rawDescData
}

var file_services_eagle_api_src_order_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_services_eagle_api_src_order_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_eagle_api_src_order_proto_goTypes = []interface{}{
	(OrderModel)(0),          // 0: eagleApi.OrderModel
	(OrderType)(0),           // 1: eagleApi.Order.type
	(OrderStatus)(0),         // 2: eagleApi.Order.status
	(OrderOption)(0),         // 3: eagleApi.Order.option
	(*Order)(nil),            // 4: eagleApi.Order
	(*Orders)(nil),           // 5: eagleApi.Orders
	(*OrderListRequest)(nil), // 6: eagleApi.OrderListRequest
	(*api.Market)(nil),       // 7: chipmunkApi.Market
	(*api.Asset)(nil),        // 8: chipmunkApi.Asset
}
var file_services_eagle_api_src_order_proto_depIdxs = []int32{
	2, // 0: eagleApi.Order.Status:type_name -> eagleApi.Order.status
	7, // 1: eagleApi.Order.Market:type_name -> chipmunkApi.Market
	1, // 2: eagleApi.Order.SellOrBuy:type_name -> eagleApi.Order.type
	0, // 3: eagleApi.Order.OrderType:type_name -> eagleApi.OrderModel
	8, // 4: eagleApi.Order.SourceAsset:type_name -> chipmunkApi.Asset
	8, // 5: eagleApi.Order.DestinationAsset:type_name -> chipmunkApi.Asset
	8, // 6: eagleApi.Order.FeeAsset:type_name -> chipmunkApi.Asset
	4, // 7: eagleApi.Orders.Elements:type_name -> eagleApi.Order
	6, // 8: eagleApi.OrderService.OrderList:input_type -> eagleApi.OrderListRequest
	5, // 9: eagleApi.OrderService.OrderList:output_type -> eagleApi.Orders
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_services_eagle_api_src_order_proto_init() }
func file_services_eagle_api_src_order_proto_init() {
	if File_services_eagle_api_src_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_eagle_api_src_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_services_eagle_api_src_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
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
		file_services_eagle_api_src_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListRequest); i {
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
			RawDescriptor: file_services_eagle_api_src_order_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_eagle_api_src_order_proto_goTypes,
		DependencyIndexes: file_services_eagle_api_src_order_proto_depIdxs,
		EnumInfos:         file_services_eagle_api_src_order_proto_enumTypes,
		MessageInfos:      file_services_eagle_api_src_order_proto_msgTypes,
	}.Build()
	File_services_eagle_api_src_order_proto = out.File
	file_services_eagle_api_src_order_proto_rawDesc = nil
	file_services_eagle_api_src_order_proto_goTypes = nil
	file_services_eagle_api_src_order_proto_depIdxs = nil
}
