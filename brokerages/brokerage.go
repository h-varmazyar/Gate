package brokerages

type BrokerageRequests interface {
	Validate() error
	//market endpoints
	OHLC(OHLCParams) *OHLCResponse
	MarketList() *MarketListResponse                             //coinex
	MarketInfo(MarketInfoParams) *MarketInfoResponse             //coinex
	MarketStatistics(MarketInfoParams) *MarketStatisticsResponse //coinex
	//account endpoints
	WalletList() *WalletListResponse
	//trading endpoints
	NewOrder(NewOrderParams) *OrderResponse
	CancelOrder(CancelOrderParams) *OrderResponse
	OrderStatus(OrderStatusParams) *OrderResponse
	OrderList(OrderListParams) *OrderListResponse

	Login(LoginParams) *BasicResponse

	UserInfo() *UserInfoResponse
	WalletInfo(WalletInfoParams) *WalletResponse
	OrderBook(OrderBookParams) *OrderBookResponse
	WalletBalance(WalletBalanceParams) *BalanceResponse
	RecentTrades(OrderBookParams) *RecentTradesResponse
	TransactionList(TransactionListParams) *TransactionListResponse
	UpdateOrderStatus(UpdateOrderStatusParams) *UpdateOrderStatusResponse
}

type BasicResponse struct {
	Error error
}
