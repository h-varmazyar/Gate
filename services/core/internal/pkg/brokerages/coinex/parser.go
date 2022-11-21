package coinex

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	"go/types"
	"reflect"
	"strconv"
)

type responseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func parseResponse(input string, response interface{}) error {
	tmp := new(responseModel)
	if err := json.Unmarshal([]byte(input), tmp); err != nil {
		return err
	}
	if tmp.Code != 0 {
		return errors.New(tmp.Message)
	}
	switch tmp.Data.(type) {
	case []interface{}:
		mapper.Slice(tmp.Data, response)
	case types.Struct:
	case map[string]interface{}:
		data, err := json.Marshal(tmp.Data)
		if err != nil {
			return err
		}
		return json.Unmarshal(data, response)
	default:
		return errors.New(fmt.Sprintf("cast data failed: %v", reflect.TypeOf(tmp.Data)))
	}
	return nil
}

func createOrder(data map[string]interface{}, market *chipmunkApi.Market) *eagleApi.Order {
	response := new(eagleApi.Order)
	response.ID = uuid.New().String()
	response.Market = market
	response.DestinationAsset = market.Destination
	response.SourceAsset = market.Source

	response.Amount, _ = strconv.ParseFloat(data["amount"].(string), 64)
	response.AveragePrice, _ = strconv.ParseFloat(data["avg_price"].(string), 64)
	response.CreatedAt, _ = strconv.ParseInt(data["create_time"].(string), 10, 64)
	response.ExecutedAmount, _ = strconv.ParseFloat(data["deal_amount"].(string), 64)
	response.Volume, _ = strconv.ParseFloat(data["deal_money"].(string), 64)
	response.FinishedAt, _ = strconv.ParseInt(data["finished_time"].(string), 10, 64)
	response.OrderNo, _ = strconv.ParseInt(data["id"].(string), 10, 64)
	response.MakerFeeRate, _ = strconv.ParseFloat(data["maker_fee_rate"].(string), 64)
	response.OrderType = eagleApi.OrderModel(eagleApi.OrderModel_value[data["order_type"].(string)])
	response.SellOrBuy = eagleApi.OrderType(eagleApi.OrderType_value[data["type"].(string)])
	response.Price, _ = strconv.ParseFloat(data["price"].(string), 64)
	response.TakerFeeRate, _ = strconv.ParseFloat(data["taker_fee_rate"].(string), 64)
	response.StockFee, _ = strconv.ParseFloat(data["stock_fee"].(string), 64)
	response.MoneyFee, _ = strconv.ParseFloat(data["money_fee"].(string), 64)
	response.AssetFee, _ = strconv.ParseFloat(data["asset_fee"].(string), 64)
	response.TransactionFee, _ = strconv.ParseFloat(data["fee_asset"].(string), 64)
	response.FeeDiscount, _ = strconv.ParseFloat(data["fee_discount"].(string), 64)
	response.UnExecutedAmount, _ = strconv.ParseFloat(data["left"].(string), 64)

	switch data["status"].(string) {
	case eagleApi.Order_not_deal.String():
		response.Status = eagleApi.Order_not_deal
	case eagleApi.Order_part_deal.String():
		response.Status = eagleApi.Order_part_deal
	case eagleApi.Order_done.String():
		response.Status = eagleApi.Order_done
	}
	return response
}
