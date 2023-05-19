package nobitex

import (
	"context"
	"encoding/json"
	"fmt"
	api "github.com/h-varmazyar/Gate/api/proto"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
)

type Requests struct {
	Auth    *api.Auth
	configs *Configs
}

func NewRequest(configs *Configs, auth *api.Auth) brokerages.Requests {
	return &Requests{
		configs: configs,
		Auth:    auth,
	}
}

func (r *Requests) AsyncOHLC(_ context.Context, inputs *brokerages.OHLCParams) (*networkAPI.Request, error) {
	if inputs.Resolution == nil {
		return nil, brokerages.ErrNilResolution
	}
	if inputs.Market == nil {
		return nil, brokerages.ErrNilMarket
	}
	if inputs.From.IsZero() {
		return nil, brokerages.ErrWrongStartTime
	}
	if inputs.To.IsZero() {
		return nil, brokerages.ErrWrongEndTime
	}

	request := new(networkAPI.Request)
	request.Method = networkAPI.Request_GET
	request.CallbackQueue = r.configs.NobitexCallbackQueue
	resolutionSeconds := inputs.Resolution.Duration
	count := int64(inputs.To.Sub(inputs.From)) / resolutionSeconds
	if (int64(inputs.To.Sub(inputs.From)) % resolutionSeconds) > 0 {
		count++
	}
	if count > 500 {
		return nil, brokerages.ErrInvalidCandleCount
	}

	request.Params = []*networkAPI.KV{
		networkAPI.NewKV("symbol", inputs.Market.Name),
		networkAPI.NewKV("resolution", inputs.Resolution.Value),
		networkAPI.NewKV("from", fmt.Sprintf("%v", inputs.From.Unix())),
		networkAPI.NewKV("to", fmt.Sprintf("%v", inputs.To.Unix()))}
	request.RateLimiterID = r.configs.NobitexPublicRateLimiterID
	request.Endpoint = "https://api.nobitex.ir/market/udf/history"

	metadataBytes, _ := json.Marshal(&brokerages.Metadata{
		Method:       brokerages.MethodOHLC,
		Platform:     api.Platform_Nobitex,
		MarketID:     inputs.Market.ID,
		ResolutionID: inputs.Resolution.ID,
	})
	request.Metadata = string(metadataBytes)
	return request, nil
}

func (r *Requests) AllMarketStatistics(_ context.Context, _ *brokerages.AllMarketStatisticsParams) (*networkAPI.Request, error) {
	return nil, brokerages.ErrUnImplemented
}

func (r *Requests) GetMarketInfo(_ context.Context, _ *brokerages.MarketInfoParams) (*networkAPI.Request, error) {
	return nil, brokerages.ErrNotAvailable
}

func (r *Requests) WalletsBalance(_ context.Context, _ *brokerages.WalletsBalanceParams) (*networkAPI.Request, error) {
	return nil, brokerages.ErrNotAvailable
	//request := &networkAPI.Request{
	//	Method:   networkAPI.Request_GET,
	//	Endpoint: "https://api.nobitex.ir/users/wallets/list",
	//	Headers: []*networkAPI.KV{
	//		networkAPI.NewKV("Authorization", r.generateAuthorization()),
	//	},
	//	Type: networkAPI.Request_Sync,
	//}
	//
	//return request, nil
}

func (r *Requests) OHLC(ctx context.Context, inputs *brokerages.OHLCParams, runner brokerages.Handler) ([]*chipmunkApi.Candle, error) {
	return nil, brokerages.ErrNotAvailable
	//request := new(networkAPI.Request)
	//request.Method = networkAPI.Request_GET
	//resolutionSeconds := inputs.Resolution.Duration / 1e6
	//count := int64(inputs.To.Sub(inputs.From)) / resolutionSeconds
	//if (int64(inputs.To.Sub(inputs.From)) % resolutionSeconds) > 0 {
	//	count++
	//}
	//if int64(count) >= 1000 {
	//	request.Params = []*networkAPI.KV{
	//		api2.NewKV("market", inputs.Market.Name),
	//		api2.NewKV("interval", inputs.Resolution.Value),
	//		api2.NewKV("start_time", fmt.Sprintf("%v", inputs.From.Unix())),
	//		api2.NewKV("end_time", fmt.Sprintf("%v", inputs.To.Unix()))}
	//	request.Endpoint = "https://www.coinex.com/res/market/kline"
	//} else {
	//	request.Params = []*networkAPI.KV{
	//		api2.NewKV("markets", inputs.Market.Name),
	//		api2.NewKV("type", inputs.Resolution.Label),
	//		api2.NewKV("limit", fmt.Sprintf("%v", int64(count))),
	//	}
	//	request.Endpoint = "https://api.coinex.com/v1/market/kline"
	//}
	//resp, err := runner(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.Code != http.StatusOK {
	//	return nil, errors.NewWithSlug(ctx, codes.NotFound, resp.Body)
	//}
	//data := make([][]interface{}, 0)
	//if err := parseResponse(resp.Body, &data); err != nil {
	//	return nil, err
	//}
	//rateLimiters := make([]*chipmunkApi.Candle, 0)
	//for _, item := range data {
	//	c := new(chipmunkApi.Candle)
	//	var err error
	//	c.Time = int64(item[0].(float64))
	//	c.Open, err = strconv.ParseFloat(item[1].(string), 64)
	//	if err != nil {
	//		continue
	//	}
	//	c.Close, err = strconv.ParseFloat(item[2].(string), 64)
	//	if err != nil {
	//		continue
	//	}
	//	c.High, err = strconv.ParseFloat(item[3].(string), 64)
	//	if err != nil {
	//		continue
	//	}
	//	c.Low, err = strconv.ParseFloat(item[4].(string), 64)
	//	if err != nil {
	//		continue
	//	}
	//	c.Volume, err = strconv.ParseFloat(item[5].(string), 64)
	//	if err != nil {
	//		continue
	//	}
	//	c.Amount, err = strconv.ParseFloat(item[6].(string), 64)
	//	if err != nil {
	//		continue
	//	}
	//	rateLimiters = append(rateLimiters, c)
	//}
	//return rateLimiters, nil
}

func (r *Requests) UpdateMarket(ctx context.Context, runner brokerages.Handler) ([]*chipmunkApi.Market, error) {
	return nil, brokerages.ErrNotAvailable
	//request := new(networkAPI.Request)
	//request.Method = networkAPI.Request_GET
	//request.Endpoint = "https://api.coinex.com/v1/market/info"
	//
	//resp, err := runner(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.Code != http.StatusOK {
	//	return nil, errors.NewWithSlug(ctx, codes.NotFound, resp.Body)
	//}
	//tmp := new(responseModel)
	//if err := json.Unmarshal([]byte(resp.Body), tmp); err != nil {
	//	return nil, err
	//}
	//if tmp.Code != 0 {
	//	log.WithError(err)
	//	return nil, errors.New(ctx, codes.Canceled)
	//}
	//data := tmp.Data.(map[string]interface{})
	//markets := make([]*chipmunkApi.Market, 0)
	//for _, value := range data {
	//	item := value.(map[string]interface{})
	//	m := new(chipmunkApi.Market)
	//	m.Platform = api.Platform_Coinex
	//	m.PricingDecimal = item["pricing_decimal"].(float64)
	//	m.TradingDecimal = item["trading_decimal"].(float64)
	//	num, err := strconv.ParseFloat(item["taker_fee_rate"].(string), 64)
	//	if err != nil {
	//		log.WithError(err).WithField("taker_fee_rate", item["taker_fee_rate"]).Error("failed to add markets")
	//		continue
	//	}
	//	m.TakerFeeRate = num
	//	num, err = strconv.ParseFloat(item["maker_fee_rate"].(string), 64)
	//	if err != nil {
	//		log.WithError(err).WithField("maker_fee_rate", item["maker_fee_rate"]).Error("failed to add markets")
	//		continue
	//	}
	//	m.MakerFeeRate = num
	//	num, err = strconv.ParseFloat(item["min_amount"].(string), 64)
	//	if err != nil {
	//		log.WithError(err).WithField("min_amount", item["min_amount"]).Error("failed to add markets")
	//		continue
	//	}
	//	m.MinAmount = num
	//	m.IsAMM = false
	//	m.Name = item["name"].(string)
	//	m.Status = api.Status_Enable
	//	markets = append(markets, m)
	//}
	//return markets, nil
}

func (r *Requests) MarketStatistics(ctx context.Context, inputs *brokerages.MarketStatisticsParams, runner brokerages.Handler) (*chipmunkApi.Candle, error) {
	return nil, brokerages.ErrNotAvailable
	//var market string
	//if inputs.Market == "" {
	//	market = strings.ToUpper(fmt.Sprint(inputs.Source, inputs.Destination))
	//} else {
	//	market = inputs.Market
	//}
	//request := new(networkAPI.Request)
	//request.Method = networkAPI.Request_GET
	//request.Params = []*networkAPI.KV{
	//	api2.NewKV("market", market),
	//}
	//request.Endpoint = "https://api.coinex.com/v1/market/ticker"
	//
	//resp, err := runner(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.Code != http.StatusOK {
	//	return nil, errors.NewWithSlug(ctx, codes.NotFound, resp.Body)
	//}
	//data := struct {
	//	Date   float64 `json:"date"`
	//	Ticker struct {
	//		Buy        string `json:"buy"`
	//		BuyAmount  string `json:"buy_amount"`
	//		Open       string `json:"open"`
	//		High       string `json:"high"`
	//		Low        string `json:"low"`
	//		Last       string `json:"last"`
	//		Sell       string `json:"sell"`
	//		SellAmount string `json:"sell_amount"`
	//		Volume     string `json:"vol"`
	//	} `json:"ticker"`
	//}{}
	//if err := parseResponse(resp.Body, &data); err != nil {
	//	return nil, err
	//}
	//candle := new(chipmunkApi.Candle)
	//candle.CreatedAt = time.Now().Unix()
	//candle.UpdatedAt = time.Now().Unix()
	//candle.Time = int64(data.Date)
	//if candle.Volume, err = strconv.ParseFloat(data.Ticker.Volume, 64); err != nil {
	//	log.WithError(err).Error("failed to parse volume")
	//	return nil, err
	//}
	//if candle.Close, err = strconv.ParseFloat(data.Ticker.Last, 64); err != nil {
	//	log.WithError(err).Error("failed to parse close")
	//	return nil, err
	//}
	//if candle.Open, err = strconv.ParseFloat(data.Ticker.Open, 64); err != nil {
	//	log.WithError(err).Error("failed to parse open")
	//	return nil, err
	//}
	//if candle.High, err = strconv.ParseFloat(data.Ticker.High, 64); err != nil {
	//	log.WithError(err).Error("failed to parse high")
	//	return nil, err
	//}
	//if candle.Low, err = strconv.ParseFloat(data.Ticker.Low, 64); err != nil {
	//	log.WithError(err).Error("failed to parse low")
	//	return nil, err
	//}
	//return candle, nil
}

func (r *Requests) MarketList(ctx context.Context, runner brokerages.Handler) (*chipmunkApi.Markets, error) {
	return nil, brokerages.ErrNotAvailable
	//request := new(networkAPI.Request)
	//request.Method = networkAPI.Request_GET
	//request.Endpoint = "https://api.coinex.com/v1/market/info"
	//
	//resp, err := runner(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.Code != http.StatusOK {
	//	return nil, errors.NewWithSlug(ctx, codes.NotFound, resp.Body)
	//}
	//data := make(map[string]struct {
	//	Name           string  `json:"name"`
	//	MinAmount      string  `json:"min_amount"`
	//	MakerFeeRate   string  `json:"maker_fee_rate"`
	//	TakerFeeRate   string  `json:"taker_fee_rate"`
	//	PricingName    string  `json:"pricing_name"`
	//	PricingDecimal float64 `json:"pricing_decimal"`
	//	TradingName    string  `json:"trading_name"`
	//	TradingDecimal float64 `json:"trading_decimal"`
	//})
	//if err := parseResponse(resp.Body, &data); err != nil {
	//	return nil, err
	//}
	//markets := new(chipmunkApi.Markets)
	//markets.Elements = make([]*chipmunkApi.Market, 0)
	//
	//for key, value := range data {
	//	market := &chipmunkApi.Market{
	//		PricingDecimal: value.PricingDecimal,
	//		TradingDecimal: value.TradingDecimal,
	//		IsAMM:          true,
	//		Name:           key,
	//		Status:         api.Status_Enable,
	//		Platform:       api.Platform_Coinex,
	//		Source:         &chipmunkApi.Asset{Name: value.TradingName},
	//		Destination:    &chipmunkApi.Asset{Name: value.PricingName},
	//	}
	//
	//	if market.TakerFeeRate, err = strconv.ParseFloat(value.TakerFeeRate, 64); err != nil {
	//		log.WithError(err).Error("failed to parse taker fee rate")
	//		return nil, err
	//	}
	//	if market.MakerFeeRate, err = strconv.ParseFloat(value.MakerFeeRate, 64); err != nil {
	//		log.WithError(err).Error("failed to parse high")
	//		return nil, err
	//	}
	//	if market.MinAmount, err = strconv.ParseFloat(value.MinAmount, 64); err != nil {
	//		log.WithError(err).Error("failed to parse low")
	//		return nil, err
	//	}
	//	markets.Elements = append(markets.Elements, market)
	//}
	//return markets, nil
}

func (r *Requests) NewOrder(ctx context.Context, inputs *brokerages.NewOrderParams, runner brokerages.Handler) (*eagleApi.Order, error) {
	return nil, brokerages.ErrNotAvailable
	//request := new(networkAPI.Request)
	//request.Method = networkAPI.Request_POST
	//request.Params = []*networkAPI.KV{
	//	api2.NewKV("access_id", r.Auth.AccessID),
	//	api2.NewKV("tonce", fmt.Sprintf("%d", time.Now().UnixNano()/1e6)),
	//}
	//request.Headers = []*networkAPI.KV{
	//	api2.NewKV("authorization", r.generateAuthorization(request.Params)),
	//	api2.NewKV("tonce", fmt.Sprintf("%d", time.Now().UnixNano()/1e6)),
	//}
	//switch inputs.OrderModel {
	//case eagleApi.OrderModel_limit:
	//	request.Endpoint = "https://api.coinex.com/v1/order/limit"
	//	request.Params = append(request.Params, api2.NewKV("markets", inputs.Market.Name))
	//	request.Params = append(request.Params, api2.NewKV("type", inputs.BuyOrSell.String()))
	//	request.Params = append(request.Params, api2.NewKV("amount", fmt.Sprintf("%f", inputs.Amount)))
	//	request.Params = append(request.Params, api2.NewKV("price", fmt.Sprintf("%f", inputs.Price)))
	//	request.Params = append(request.Params, api2.NewKV("source_id", fmt.Sprintf("%d", rand.Int())))
	//	request.Params = append(request.Params, api2.NewKV("option", inputs.Option.String()))
	//	request.Params = append(request.Params, api2.NewKV("client_id", strings.ReplaceAll(uuid.New().String(), "-", "")))
	//	request.Params = append(request.Params, api2.NewKV("hide", inputs.HideOrder))
	//case eagleApi.OrderModel_market:
	//	request.Endpoint = "https://api.coinex.com/v1/order/market"
	//	request.Params = append(request.Params, api2.NewKV("markets", inputs.Market.Name))
	//	request.Params = append(request.Params, api2.NewKV("type", inputs.BuyOrSell.String()))
	//	request.Params = append(request.Params, api2.NewKV("amount", fmt.Sprintf("%f", inputs.Amount)))
	//	request.Params = append(request.Params, api2.NewKV("option", inputs.Option.String()))
	//	request.Params = append(request.Params, api2.NewKV("client_id", strings.ReplaceAll(uuid.New().String(), "-", "")))
	//case eagleApi.OrderModel_stopLimit:
	//	request.Endpoint = "https://api.coinex.com/v1/order/stop/limit"
	//	request.Params = append(request.Params, api2.NewKV("markets", inputs.Market.Name))
	//	request.Params = append(request.Params, api2.NewKV("type", inputs.BuyOrSell.String()))
	//	request.Params = append(request.Params, api2.NewKV("amount", fmt.Sprintf("%f", inputs.Amount)))
	//	request.Params = append(request.Params, api2.NewKV("price", fmt.Sprintf("%f", inputs.Price)))
	//	request.Params = append(request.Params, api2.NewKV("source_id", fmt.Sprintf("%d", rand.Int())))
	//	request.Params = append(request.Params, api2.NewKV("option", inputs.Option.String()))
	//	request.Params = append(request.Params, api2.NewKV("client_id", strings.ReplaceAll(uuid.New().String(), "-", "")))
	//	request.Params = append(request.Params, api2.NewKV("hide", inputs.HideOrder))
	//	request.Params = append(request.Params, api2.NewKV("stop_price", fmt.Sprintf("%f", inputs.StopPrice)))
	//case eagleApi.OrderModel_ioc:
	//	request.Endpoint = "https://api.coinex.com/v1/order/ioc"
	//	request.Params = append(request.Params, api2.NewKV("markets", inputs.Market.Name))
	//	request.Params = append(request.Params, api2.NewKV("type", inputs.BuyOrSell.String()))
	//	request.Params = append(request.Params, api2.NewKV("amount", fmt.Sprintf("%f", inputs.Amount)))
	//	request.Params = append(request.Params, api2.NewKV("price", fmt.Sprintf("%f", inputs.Price)))
	//	request.Params = append(request.Params, api2.NewKV("source_id", fmt.Sprintf("%d", rand.Int())))
	//	request.Params = append(request.Params, api2.NewKV("client_id", strings.ReplaceAll(uuid.New().String(), "-", "")))
	//default:
	//	return nil, errors.New(ctx, codes.Unimplemented)
	//}
	//
	//resp, err := runner(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.Code != http.StatusOK {
	//	return nil, errors.NewWithSlug(ctx, codes.Unknown, resp.Body)
	//}
	//data := make(map[string]interface{})
	//if err := parseResponse(resp.Body, &data); err != nil {
	//	return nil, err
	//}
	//
	//return createOrder(data, inputs.Market), nil
}

func (r *Requests) CancelOrder(ctx context.Context, inputs *brokerages.CancelOrderParams, runner brokerages.Handler) (*eagleApi.Order, error) {
	return nil, brokerages.ErrNotAvailable
	//request := new(networkAPI.Request)
	//request.Method = networkAPI.Request_DELETE
	//request.Params = []*networkAPI.KV{
	//	api2.NewKV("access_id", r.Auth.AccessID),
	//	api2.NewKV("tonce", fmt.Sprintf("%d", time.Now().UnixNano()/1e6)),
	//}
	//request.Headers = []*networkAPI.KV{
	//	api2.NewKV("authorization", r.generateAuthorization(request.Params)),
	//	api2.NewKV("tonce", fmt.Sprintf("%d", time.Now().UnixNano()/1e6)),
	//}
	//request.Params = []*networkAPI.KV{
	//	api2.NewKV("id", inputs.ServerOrderId),
	//	api2.NewKV("markets", inputs.Market.Name),
	//	api2.NewKV("account_id", 0)}
	//request.Endpoint = "https://api.coinex.com/v1/order/pending"
	//
	//resp, err := runner(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.Code != http.StatusOK {
	//	return nil, errors.NewWithSlug(ctx, codes.Unknown, resp.Body)
	//}
	//data := make(map[string]interface{})
	//if err := parseResponse(resp.Body, &data); err != nil {
	//	return nil, err
	//}
	//
	//return createOrder(data, inputs.Market), nil
}

func (r *Requests) OrderStatus(ctx context.Context, inputs *brokerages.OrderStatusParams, runner brokerages.Handler) (*eagleApi.Order, error) {
	return nil, brokerages.ErrNotAvailable
	//request := new(networkAPI.Request)
	//request.Method = networkAPI.Request_GET
	//request.Params = []*networkAPI.KV{
	//	api2.NewKV("access_id", r.Auth.AccessID),
	//	api2.NewKV("tonce", fmt.Sprintf("%d", time.Now().UnixNano()/1e6)),
	//}
	//request.Headers = []*networkAPI.KV{
	//	api2.NewKV("authorization", r.generateAuthorization(request.Params)),
	//	api2.NewKV("tonce", fmt.Sprintf("%d", time.Now().UnixNano()/1e6)),
	//}
	//request.Params = []*networkAPI.KV{
	//	api2.NewKV("id", inputs.ServerOrderId),
	//	api2.NewKV("markets", inputs.Market.Name)}
	//request.Endpoint = "https://api.coinex.com/v1/order/status"
	//
	//resp, err := runner(ctx, request)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.Code != http.StatusOK {
	//	return nil, errors.NewWithSlug(ctx, codes.Unknown, resp.Body)
	//}
	//data := make(map[string]interface{})
	//if err := parseResponse(resp.Body, &data); err != nil {
	//	return nil, err
	//}
	//
	//return createOrder(data, inputs.Market), nil
}

//
//func (r *Requests) generateAuthorization(params []*networkAPI.KV) string {
//	urlParameters := url.Values{}
//	for _, param := range params {
//		urlParameters.Add(param.Key, parseValue(param))
//	}
//	queryParamsString := urlParameters.Encode()
//	toEncodeParamsString := queryParamsString + "&secret_key=" + r.Auth.SecretKey
//	w := md5.New()
//	_, _ = io.WriteString(w, toEncodeParamsString)
//	return strings.ToUpper(fmt.Sprintf("%x", w.Sum(nil)))
//}
//
//func parseValue(param *networkAPI.KV) string {
//	value := ""
//	switch v := param.Value.(type) {
//	case *networkAPI.KV_String_:
//		value = v.String_
//	case *networkAPI.KV_Bool:
//		value = fmt.Sprint(v.Bool)
//	case *networkAPI.KV_Float32:
//		value = fmt.Sprint(v.Float32)
//	case *networkAPI.KV_Float64:
//		value = fmt.Sprint(v.Float64)
//	case *networkAPI.KV_Integer:
//		value = fmt.Sprint(v.Integer)
//	}
//	return value
//}
