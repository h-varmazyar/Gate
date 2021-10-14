package core

import (
	"github.com/fatih/color"
	"github.com/gofrs/uuid"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/indicators"
	"github.com/mrNobody95/Gate/models"
	"github.com/mrNobody95/Gate/storage"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"time"
)

type MarketThread struct {
	*Node
	Market          *models.Market
	StartFrom       time.Time
	Resolution      *models.Resolution
	IndicatorConfig *indicators.Configuration
	CandlePool      *storage.CandlePool
}

func (thread *MarketThread) CollectPrimaryData() error {
	color.HiGreen("Collecting primary data for: %s", thread.Market.Name)

	if thread.IndicatorConfig == nil {
		thread.IndicatorConfig = indicators.DefaultConfig()
	}
	lastTime := thread.StartFrom
	var list []models.Candle
	for {
		if tmpList, err := models.LoadCandleList(thread.Market.Id, thread.PivotResolution.Id, lastTime); err != nil {
			if err == gorm.ErrRecordNotFound {
				break
			} else {
				return err
			}
		} else if len(tmpList) > 0 {
			list = append(list, tmpList...)
			lastTime = tmpList[len(tmpList)-1].Time
		} else {
			break
		}
	}
	for i := 0; i < len(list); i++ {
		list[i].FromDb = true
	}

	count := (time.Now().Unix() - lastTime.Unix()) / int64(thread.PivotResolution.Duration/time.Second)
	j := count / 500
	if count%500 != 0 {
		j++
	}
	for i := int64(1); i <= j; i++ {
		from := lastTime.Unix() + 500*(i-1)*int64(thread.PivotResolution.Duration/time.Second)
		to := lastTime.Unix() + 500*i*int64(thread.PivotResolution.Duration/time.Second)
		if candles, err := thread.makeOHLCRequest(thread.PivotResolution, from, to); err != nil {
			return err
		} else {
			if len(list) > 0 && list[len(list)-1].Time == candles[0].Time {
				candles[0].ID = list[len(list)-1].ID
				list = append(list[:len(list)-1], candles...)
			} else {
				list = append(list, candles...)
			}
		}
	}
	if err := thread.IndicatorConfig.CalculateIndicators(list); err != nil {
		return err
	}
	if err := thread.CandlePool.ImportNewCandles(list); err != nil {
		return err
	}
	return nil
}

func (thread *MarketThread) PeriodicOHLC() {
	color.HiGreen("Making Periodic ohlc for: %s", thread.Market.Name)
	for {
		start := time.Now()
		if candles, err := thread.makeOHLCRequest(thread.PivotResolution, thread.CandlePool.GetLastCandle().Time.Unix(), time.Now().Unix()); err != nil {
			log.Errorf("ohlc request failed: %s", err.Error())
		} else {
			for _, candle := range candles {
				if poolErr := thread.CandlePool.UpdateLastCandle(candle); poolErr != nil {
					log.WithError(poolErr).Errorf("update pool failed for market %s in timeframe %s",
						candle.Market.Name, candle.Resolution.Label)
					continue
				}
				thread.IndicatorConfig.UpdateIndicators(thread.CandlePool)
			}
		}
		if thread.EnableTrading || thread.FakeTrading {
			thread.checkForSignals()
		}
		end := time.Now()
		idealTime := thread.Strategy.IndicatorUpdatePeriod - end.Sub(start)
		if idealTime > 0 {
			time.Sleep(idealTime)
		}
	}
}

func (thread *MarketThread) makeOHLCRequest(resolution *models.Resolution, from, to int64) ([]models.Candle, error) {
	response := thread.Requests.OHLC(brokerages.OHLCParams{
		Resolution: resolution,
		Market:     thread.Market,
		From:       from,
		To:         to,
	})
	if response.Error != nil {
		return nil, response.Error
	}
	return response.Candles, nil
}

func (thread *MarketThread) checkForSignals() {
	walletResponse := thread.Requests.WalletInfo(brokerages.WalletInfoParams{
		WalletName: thread.Market.Destination.Symbol,
	})
	if walletResponse.Error != nil {
		return
	}
	if walletResponse.Wallet.ActiveBalance > 0 {
		candles := thread.CandlePool.GetLastNCandle(2)
		rsi := thread.Strategy.RsiSignal(candles[0], candles[1])
		bb := thread.Strategy.BollingerBandSignal(candles[0], candles[1], thread.Market.MakerFeeRate, thread.Market.TakerFeeRate)
		stochastic := thread.Strategy.StochasticSignal(candles[0], candles[1])
		if rsi && bb && stochastic {
			id, _ := uuid.NewV4()
			price := thread.CandlePool.GetLastCandle().Close
			newOrderResponse := thread.Requests.NewOrder(brokerages.NewOrderParams{
				OrderKind:  models.LimitOrderKind,
				ClientUUID: strings.ReplaceAll(id.String(), "-", ""),
				BuyOrSell:  models.Buy,
				Price:      price,
				Market:     *thread.Market,
				Amount:     walletResponse.Wallet.ActiveBalance / price,
				Option:     models.OptionNormal,
			})
			if newOrderResponse.Error != nil {
				log.WithError(newOrderResponse.Error).Error("new order failed")
				return
			}
			if newOrderResponse.Order.Status == models.NewOrderStatus {
				if err := newOrderResponse.Order.Create(); err != nil {
					log.WithError(err).Error("save new order failed")
					return
				}
				thread.checkOrder(newOrderResponse.Order)
			}
		}
	}
}

func (thread *MarketThread) cancelOrder(order models.Order) bool {
	isBuy := true
	if order.SellOrBuy == models.Sell {
		isBuy = false
	}
	response := thread.Requests.CancelOrder(brokerages.CancelOrderParams{
		ServerOrderId: order.ServerOrderId,
		Market:        *thread.Market,
		IsBuy:         isBuy,
		ClientUUID:    order.ClientUUID,
		AllOrders:     false,
	})
	if response.Error != nil {
		log.WithError(response.Error).Error("cancel order failed")
		return false
	}
	if err := response.Order.Update(); err != nil {
		log.WithError(response.Error).Error("cancel order failed")
		return false
	}
	return true
}

func (thread *MarketThread) checkOrder(order models.Order) {
	func() {
		checked := false
		for {
			start := time.Now()
			response := thread.Requests.OrderStatus(brokerages.OrderStatusParams{
				ServerOrderId: order.ServerOrderId,
				Market:        *thread.Market,
				ClientUUID:    order.ClientUUID,
			})
			if response.Error != nil {
				log.WithError(response.Error).Error("check order status failed")
			} else {
				if checked {
					switch response.Order.Status {
					case models.DoneOrderStatus:
						if err := response.Order.Update(); err != nil {
							log.WithError(err).Error("update done order failed")
						}
						checked = true
					case models.NewOrderStatus,
						models.PartlyExecutedOrderStatus,
						models.UnExecutedOrderStatus:
						if response.Order.CreatedAt.Add(time.Minute * 15).Before(time.Now()) {
							if err := response.Order.Update(); err != nil {
								log.WithError(err).Error("update un executed order failed")
							}
							if thread.cancelOrder(order) {
								checked = true
							}
						}
					}
				}
				if thread.checkPrice(response.Order) {
					return
				}
			}
			sleep := time.Second - time.Now().Sub(start)
			if sleep > 0 {
				time.Sleep(sleep)
			}
		}
	}()
}

func (thread *MarketThread) checkPrice(order models.Order) bool {
	candle := thread.CandlePool.GetLastCandle()
	upperPrice := order.AveragePrice * (1 + thread.Market.MakerFeeRate/100) * (1 + thread.Strategy.MinGainPercent/100) * (1 + thread.Market.TakerFeeRate/100)
	lowerPrice := order.AveragePrice * (1 - thread.Market.MakerFeeRate/100) * (1 - thread.Strategy.LossPercent/100) * (1 - thread.Market.TakerFeeRate/100)
	if candle.Close <= lowerPrice {
		id, _ := uuid.NewV4()
		response := thread.Requests.NewOrder(brokerages.NewOrderParams{
			OrderKind:  models.LimitOrderKind,
			ClientUUID: strings.ReplaceAll(id.String(), "-", ""),
			BuyOrSell:  models.Sell,
			Price:      candle.Close,
			Market:     *thread.Market,
			Amount:     order.ExecutedAmount,
			Option:     models.OptionNormal,
		})
		if response.Error != nil {
			log.WithError(response.Error).Error("check price order sell failed")
		} else {
			return true
		}
	}
	if candle.Close >= upperPrice {
		id, _ := uuid.NewV4()
		response := thread.Requests.NewOrder(brokerages.NewOrderParams{
			OrderKind:  models.LimitOrderKind,
			ClientUUID: strings.ReplaceAll(id.String(), "-", ""),
			BuyOrSell:  models.Buy,
			Price:      candle.Close,
			Market:     *thread.Market,
			Amount:     order.ExecutedAmount,
			Option:     models.OptionNormal,
		})
		if response.Error != nil {
			log.WithError(response.Error).Error("check price order buy failed")
		} else {
			return true
		}
	}
	return false
}
