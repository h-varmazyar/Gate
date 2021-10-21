package models

import (
	"time"
)

type Order struct {
	ClientUUID       string `gorm:"primarykey"`
	ServerOrderId    int64
	Amount           float64
	CreatedAt        time.Time
	FinishedAt       time.Time
	ExecutedAmount   float64
	UnExecutedAmount float64
	Status           OrderStatus
	ExecutedPrice    float64
	Market           Market `gorm:"foreignKey:MarketRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	MarketRefer      uint16
	MakerFeeRate     float64
	TakerFeeRate     float64
	SellOrBuy        OrderType
	OrderKind        OrderKind
	AveragePrice     float64
	TransactionFee   float64
	SourceAsset      *Asset `gorm:"foreignKey:SourceRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DestinationAsset *Asset `gorm:"foreignKey:DestinationRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FeeAsset         *Asset `gorm:"foreignKey:FeeAssetRefer;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SourceRefer      uint
	DestinationRefer uint
	FeeAssetRefer    uint
	FeeDiscount      float64
	AssetFee         float64
	MoneyFee         float64
	StockFee         float64
	User             string

	//Fee                 float64
	//Qty                 int
	//User                string
	//Price               float64
	//TotalPrice          string
	//MatchedVolume       string
	//UnMatchedVolume     string
}

type OrderType string
type OrderKind string
type OrderOption string
type OrderStatus string

//type Asset string

const (
	Ask  OrderType = "ask"
	Bid  OrderType = "bid"
	Buy  OrderType = "buy"
	Sell OrderType = "sell"
)

const (
	LimitOrderKind         OrderKind = "limit"
	MarketOrderKind        OrderKind = "market"
	MultipleLimitOrderKind OrderKind = "batch"
	IOCOrderKind           OrderKind = "ioc"
	StopLimitOrderKind     OrderKind = "stop-limit"
	AllOrderKind           OrderKind = "all"
)

const (
	NewOrderStatus            OrderStatus = "new"
	DoneOrderStatus           OrderStatus = "done"
	CanceledOrderStatus       OrderStatus = "cancel"
	UnExecutedOrderStatus     OrderStatus = "not_deal"
	PartlyExecutedOrderStatus OrderStatus = "part_deal"
)

const (
	OptionIOC       OrderOption = "IOC"
	OptionFOK       OrderOption = "FOK"
	OptionNormal    OrderOption = "NORMAL"
	OptionMakerOnly OrderOption = "MAKER_ONLY"
)

func (order *Order) Create() error {
	return db.Model(&Order{}).Create(order).Error
}

func (order *Order) Update() error {
	return db.Updates(order).Error
}
