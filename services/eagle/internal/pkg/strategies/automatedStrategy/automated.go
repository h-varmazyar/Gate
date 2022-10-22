package automatedStrategy

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	telegramBotApi "github.com/h-varmazyar/Gate/services/telegramBot/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"sync"
	"time"
)

type automated struct {
	*repository.Strategy
	functionsService brokerageApi.FunctionsServiceClient
	walletsService   chipmunkApi.WalletsServiceClient
	candleService    chipmunkApi.CandleServiceClient
	botService       telegramBotApi.BotServiceClient
	//checkTicker      *time.Ticker
	withTrading bool
}

func NewAutomatedStrategy(strategy *eagleApi.Strategy, withTrading bool) (*automated, error) {
	if strategy == nil {
		return nil, errors.NewWithSlug(context.Background(), codes.FailedPrecondition, "empty_strategy")
	}
	automated := new(automated)
	automated.Strategy = new(repository.Strategy)
	mapper.Struct(strategy, automated.Strategy)
	mapper.Slice(strategy.Indicators, &automated.Indicators)
	automated.withTrading = withTrading

	log.Warnf("ind1: %v", strategy.Indicators)

	brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Brokerage)
	automated.functionsService = brokerageApi.NewFunctionsServiceClient(brokerageConn)

	chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
	automated.walletsService = chipmunkApi.NewWalletsServiceClient(chipmunkConn)
	automated.candleService = chipmunkApi.NewCandleServiceClient(chipmunkConn)

	telegramBotConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.TelegramBot)
	automated.botService = telegramBotApi.NewBotServiceClient(telegramBotConn)

	return automated, nil
}

func (s *automated) CheckForSignals(ctx context.Context, market *chipmunkApi.Market) {
	var (
		err       error
		reference *chipmunkApi.Reference
		marketID  uuid.UUID
		candles   *chipmunkApi.Candles
	)

	marketID, err = uuid.Parse(market.ID)
	if err != nil {
		log.WithError(err).Errorf("failed to parse markets %v", market)
		return
	}

	time.Sleep(time.Minute)
	checkTicker := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			checkTicker.Stop()
		case <-checkTicker.C:
			reference, err = s.walletsService.ReturnReference(context.Background(), &chipmunkApi.ReturnReferenceReq{ReferenceName: market.Destination.Name})
			if err != nil {
				log.WithError(err).Errorf("failed to fetch wallet info for market %v and reference %v", marketID, market.Destination.Name)
				continue
			}
			//if reference.ActiveBalance <= 0 || reference.ActiveBalance < reference.TotalBalance/10 {
			//	continue
			//}
			candles, err = s.candleService.ReturnLastNCandles(ctx, &chipmunkApi.BufferedCandlesRequest{
				ResolutionID: s.WorkingResolutionID.String(),
				MarketID:     marketID.String(),
				Count:        2,
			})
			if err != nil {
				log.WithError(err).Error("failed to fetch candles")
				continue
			}

			strength := s.calculateSignalStrength(candles.Elements, market)
			if strength >= 0.9 {
				price := candles.Elements[len(candles.Elements)-1].Close
				s.sendSignalToBot(ctx, market.Name, price, strength)
				if s.withTrading {
					balance := float64(0)
					if reference.ActiveBalance < reference.TotalBalance/10 {
						balance = reference.ActiveBalance
					} else {
						balance = reference.TotalBalance / 10
					}
					order, err := s.functionsService.NewOrder(context.Background(), &brokerageApi.NewOrderReq{
						Market: market,
						Type:   eagleApi.Order_buy,
						Amount: balance / price,
						Price:  price,
						Option: eagleApi.Order_NORMAL,
						Model:  eagleApi.OrderModel_limit,
					})
					if err != nil {
						log.WithError(err).Errorf("failed to place new order for markets %v", market.ID)
					}
					go s.manageBuyOrder(ctx, market, order)
				}
			}
		}
	}
}

func (s *automated) calculateSignalStrength(candles []*chipmunkApi.Candle, market *chipmunkApi.Market) float64 {
	strength := float64(0)
	rsi, stochastic, bb := float64(0), float64(0), float64(0)
	for _, strategyIndicator := range s.Indicators {
		switch strategyIndicator.Type {
		case chipmunkApi.Indicator_RSI:
			rsi = s.checkRSI(candles, strategyIndicator.IndicatorID)
			strength += rsi
		case chipmunkApi.Indicator_Stochastic:
			stochastic = s.checkStochastic(candles, strategyIndicator.IndicatorID)
			strength += stochastic
		case chipmunkApi.Indicator_BollingerBands:
			bb = s.checkBollingerBand(candles, strategyIndicator.IndicatorID, market.MakerFeeRate, market.TakerFeeRate)
			strength += bb
		}
	}
	strength = strength / float64(len(s.Indicators))
	log.Infof("market %v - total: %v - rsi: %v - stochastic: %v - bb: %v", market.Name, strength, rsi, stochastic, bb)
	return strength
}

func (s *automated) sendSignalToBot(ctx context.Context, market string, price float64, metadata interface{}) {
	text := fmt.Sprintf(`new buy signal raise in spot of coinex:
market: %s
enter price: %v
other data: %v
`, market, price, metadata)
	if _, err := s.botService.SendMessage(ctx, &telegramBotApi.Message{
		ChatID: configs.Variables.BroadcastChannelID,
		Text:   text,
	}); err != nil {
		log.WithError(err).Errorf("failed to send signal message to bot: %v", text)
	}
}

func (s *automated) manageBuyOrder(ctx context.Context, market *chipmunkApi.Market, order *eagleApi.Order) {
	var err error
	endTicker := time.NewTicker(time.Minute * 15)
	checkTicker := time.NewTicker(time.Second)
	pool := new(AssetBalancePool)
	pool.Lock = new(sync.Mutex)
	pool.Available = order.ExecutedAmount
	pool.Market = market
LOOP:
	for {
		select {
		case <-endTicker.C:
			order, err = s.functionsService.CancelOrder(ctx, &brokerageApi.CancelOrderReq{
				ServerOrderID: order.ServerOrderId,
				Market:        market,
			})
			if err != nil {
				log.WithError(err).Errorf("failed to cancel order %v", order.ID)
				break
			}
			break LOOP

		case <-checkTicker.C:
			order, err = s.functionsService.OrderStatus(ctx, &brokerageApi.OrderStatusReq{})
			if err != nil {
				log.WithError(err).Errorf("failed to get status of order %v", order.ID)
				break
			}
			if order.ExecutedAmount != pool.Available {
				pool.Lock.Lock()
				pool.Available = order.ExecutedAmount
				pool.AveragePrice = order.AveragePrice
				pool.Lock.Unlock()
			}
			if order.Status == eagleApi.Order_done {
				break LOOP
			}
			if !pool.Running {
				pool.Lock.Lock()
				pool.Running = true
				pool.Lock.Unlock()
				go s.manageBidOrder(ctx, pool)
			}
		}
	}
	pool.Lock.Lock()
	pool.IsBaseOrderDone = true
	pool.Total = order.ExecutedAmount
	pool.Lock.Unlock()
	checkTicker.Stop()
	endTicker.Stop()
}

func (s *automated) manageBidOrder(ctx context.Context, pool *AssetBalancePool) {
	var (
		err       error
		last      *brokerageApi.MarketStatisticsResp
		openOrder *eagleApi.Order
		sellPrice float64
	)
	ticker := time.NewTicker(time.Second)
LOOP:
	for range ticker.C {
		if pool.Available == 0 {
			continue
		}
		if openOrder != nil {
			openOrder, err = s.functionsService.OrderStatus(ctx, &brokerageApi.OrderStatusReq{
				ServerOrderID: openOrder.ServerOrderId,
				Market:        pool.Market,
			})
			pool.Lock.Lock()
			pool.Sold += openOrder.ExecutedAmount
			pool.Lock.Unlock()
			if pool.IsBaseOrderDone && pool.Total == pool.Sold {
				break LOOP
			}
		}
		if last, err = s.functionsService.MarketStatistics(ctx, &brokerageApi.MarketStatisticsReq{
			MarketName: pool.Market.Name,
		}); err != nil {
			log.WithError(err).Errorf("failed to get last candles of %v", pool.Market.Name)
			continue
		}

		if pool.AveragePrice*(0.98) > last.Close {
			sellPrice = pool.AveragePrice * 0.97
		}

		makerFee := pool.AveragePrice * pool.MakerFeeRate
		takerFee := pool.AveragePrice * (1 + s.MinProfitPerTradeRate/100) * pool.TakerFeeRate
		profit := pool.AveragePrice*(1+s.MinProfitPerTradeRate/100) + makerFee + takerFee

		if profit*0.9 < last.Close {
			sellPrice = profit
		}

		if sellPrice != 0 {
			updateOrder := false
			if openOrder != nil && openOrder.Amount != pool.Available {
				if _, err = s.functionsService.CancelOrder(ctx, &brokerageApi.CancelOrderReq{
					ServerOrderID: openOrder.ServerOrderId,
					Market:        pool.Market,
				}); err != nil {
					log.WithError(err).Errorf("failed to cancel order %v", openOrder.ServerOrderId)
					continue
				}
				updateOrder = true
			}
			if updateOrder || openOrder == nil {
				if openOrder, err = s.functionsService.NewOrder(ctx, &brokerageApi.NewOrderReq{
					Market: pool.Market,
					Type:   eagleApi.Order_sell,
					Amount: pool.Available,
					Price:  sellPrice,
					Option: eagleApi.Order_NORMAL,
					Model:  eagleApi.OrderModel_limit,
				}); err != nil {
					log.WithError(err).Errorf("failed to update sell limit order")
				}
			}
		}
	}
	ticker.Stop()
}

func (s *automated) checkRSI(candles []*chipmunkApi.Candle, indicatorID uuid.UUID) float64 {
	if chipmunkApi.GetRSIValue(candles[0].IndicatorValues[indicatorID.String()]).RSI < 30 &&
		chipmunkApi.GetRSIValue(candles[1].IndicatorValues[indicatorID.String()]).RSI >= 30 {
		return 1
	}
	return 0
}

func (s *automated) checkStochastic(candles []*chipmunkApi.Candle, indicatorID uuid.UUID) float64 {
	stochastic0 := chipmunkApi.GetStochasticValue(candles[0].IndicatorValues[indicatorID.String()])
	stochastic1 := chipmunkApi.GetStochasticValue(candles[1].IndicatorValues[indicatorID.String()])
	if stochastic1.IndexD > 20 || stochastic0.IndexK > 20 {
		return 0
	}
	if stochastic0.IndexK < stochastic1.IndexK {
		return 1
	}

	return 0
}

func (s *automated) checkBollingerBand(candles []*chipmunkApi.Candle, indicatorID uuid.UUID, makerFeeRate, takerFeeRate float64) float64 {
	bb0 := chipmunkApi.GetBollingerBandsValue(candles[0].IndicatorValues[indicatorID.String()])
	bb1 := chipmunkApi.GetBollingerBandsValue(candles[1].IndicatorValues[indicatorID.String()])
	if candles[0].Low > bb0.LowerBand {
		return 0
	}

	price := candles[1].Close * (1 + makerFeeRate/100) * (1 + s.MinProfitPerTradeRate/100) * (1 + takerFeeRate/100)
	if price < bb1.UpperBand {
		return 1
	}
	return 0
}
