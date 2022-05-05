package automatedStrategy

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/api"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"sync"
	"time"
)

type automated struct {
	*repository.Strategy
	walletsService chipmunkApi.WalletsServiceClient
	ohlcService    chipmunkApi.OhlcServiceClient
	orderService   brokerageApi.OrderServiceClient
	checkTicker    *time.Ticker
	withTrading    bool
}

func NewAutomatedStrategy(strategy *brokerageApi.Strategy, withTrading bool) (*automated, error) {
	if strategy == nil {
		return nil, errors.NewWithSlug(context.Background(), codes.FailedPrecondition, "empty_strategy")
	}
	automated := new(automated)
	mapper.Struct(strategy, automated.Strategy)
	automated.withTrading = withTrading
	chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
	brokerageConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
	automated.walletsService = chipmunkApi.NewWalletsServiceClient(chipmunkConn)
	automated.orderService = brokerageApi.NewOrderServiceClient(brokerageConn)
	return automated, nil
}

func (s *automated) CheckForSignals(ctx context.Context, market *brokerageApi.Market) {
	var (
		err      error
		wallet   *brokerageApi.Wallet
		marketID uuid.UUID
		candles  *api.Candles
		strength float64
	)

	marketID, err = uuid.Parse(market.ID)
	if err != nil {
		log.WithError(err).Errorf("failed to parse market %v", market)
		return
	}

	s.checkTicker = time.NewTicker(time.Second)
	for range s.checkTicker.C {
		wallet, err = s.walletsService.ReturnByName(context.Background(), &chipmunkApi.ReturnWalletByDestReq{Destination: market.Destination.Name})
		if err != nil {
			log.WithError(err).Errorf("failed to fetch wallet info for %v", marketID)
			continue
		}
		if wallet.ActiveBalance <= 0 || wallet.ActiveBalance < wallet.TotalBalance/10 {
			continue
		}
		strength = 0
		candles, err = s.ohlcService.ReturnLastNCandles(ctx, &chipmunkApi.BufferedCandlesRequest{
			ResolutionID: "",
			MarketID:     marketID.String(),
			Count:        2,
		})
		if err != nil {
			continue
		}
		for _, strategyIndicator := range s.Indicators {
			switch strategyIndicator.Type {
			case chipmunkApi.IndicatorType_RSI:
				strength += s.checkRSI(candles.Candles, strategyIndicator.IndicatorID)
			case chipmunkApi.IndicatorType_Stochastic:
				strength += s.checkStochastic(candles.Candles, strategyIndicator.IndicatorID)
			case chipmunkApi.IndicatorType_BollingerBands:
				strength += s.checkBollingerBand(candles.Candles, strategyIndicator.IndicatorID, market.MakerFeeRate, market.TakerFeeRate)
			}
		}
		strength /= float64(len(s.Indicators))
		if strength >= 0.9 {
			price := candles.Candles[len(candles.Candles)-1].Close
			balance := float64(0)
			if wallet.ActiveBalance < wallet.TotalBalance/10 {
				balance = wallet.ActiveBalance
			} else {
				balance = wallet.TotalBalance / 10
			}
			order, err := s.orderService.NewOrder(context.Background(), &brokerageApi.NewOrderRequest{
				MarketID: market.ID,
				Type:     brokerageApi.Order_buy,
				Amount:   balance / price,
				Price:    price,
				Option:   brokerageApi.Order_NORMAL,
				Model:    brokerageApi.OrderModel_limit,
			})
			if err != nil {
				log.WithError(err).Errorf("failed to place new order for market %v", market.ID)
			}
			go s.manageBuyOrder(ctx, market, order)
		}
	}
}

func (s *automated) Stop() {
	s.checkTicker.Stop()
}

func (s *automated) manageBuyOrder(ctx context.Context, market *brokerageApi.Market, order *brokerageApi.Order) {
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
			order, err = s.orderService.CancelOrder(ctx, &brokerageApi.CancelOrderRequest{
				ServerOrderID: order.ServerOrderId,
				MarketID:      market.ID,
			})
			if err != nil {
				log.WithError(err).Errorf("failed to cancel order %v", order.ID)
				break
			}
			break LOOP

		case <-checkTicker.C:
			order, err = s.orderService.OrderStatus(ctx, &brokerageApi.OrderStatusRequest{})
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
			if order.Status == brokerageApi.Order_done {
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
		last      *api.Candles
		openOrder *brokerageApi.Order
		sellPrice float64
	)
	ticker := time.NewTicker(time.Second)
LOOP:
	for range ticker.C {
		if pool.Available == 0 {
			continue
		}
		if openOrder != nil {
			openOrder, err = s.orderService.OrderStatus(ctx, &brokerageApi.OrderStatusRequest{
				ServerOrderID: openOrder.ServerOrderId,
				MarketID:      pool.Market.ID,
			})
			pool.Lock.Lock()
			pool.Sold += openOrder.ExecutedAmount
			pool.Lock.Unlock()
			if pool.IsBaseOrderDone && pool.Total == pool.Sold {
				break LOOP
			}
		}
		if last, err = s.ohlcService.ReturnLastNCandles(ctx, &chipmunkApi.BufferedCandlesRequest{
			ResolutionID: s.ResolutionID.String(),
			MarketID:     pool.Market.ID,
			Count:        1,
		}); err != nil {
			log.WithError(err).Errorf("failed to get last candle of %v", pool.Market.Name)
			continue
		}

		if pool.AveragePrice*(0.98) > last.Candles[0].Close {
			sellPrice = pool.AveragePrice * 0.97
		}

		makerFee := pool.AveragePrice * pool.MakerFeeRate
		takerFee := pool.AveragePrice * (1 + s.MinProfitPercentage/100) * pool.TakerFeeRate
		profit := pool.AveragePrice*(1+s.MinProfitPercentage/100) + makerFee + takerFee

		if profit*0.9 < last.Candles[0].Close {
			sellPrice = profit
		}

		if sellPrice != 0 {
			updateOrder := false
			if openOrder != nil && openOrder.Amount != pool.Available {
				if _, err = s.orderService.CancelOrder(ctx, &brokerageApi.CancelOrderRequest{
					ServerOrderID: openOrder.ServerOrderId,
					MarketID:      pool.Market.ID,
				}); err != nil {
					log.WithError(err).Errorf("failed to cancel order %v", openOrder.ServerOrderId)
					continue
				}
				updateOrder = true
			}
			if updateOrder || openOrder == nil {
				if openOrder, err = s.orderService.NewOrder(ctx, &brokerageApi.NewOrderRequest{
					MarketID: pool.Market.Name,
					Type:     brokerageApi.Order_sell,
					Amount:   pool.Available,
					Price:    sellPrice,
					Option:   brokerageApi.Order_NORMAL,
					Model:    brokerageApi.OrderModel_limit,
				}); err != nil {
					log.WithError(err).Errorf("failed to update sell limit order")
				}
			}
		}
	}
	ticker.Stop()
}

func (s *automated) checkRSI(candles []*api.Candle, indicatorID uuid.UUID) float64 {
	if candles[0].IndicatorValues.RSIs[indicatorID.String()].RSI < 30 &&
		candles[1].IndicatorValues.RSIs[indicatorID.String()].RSI >= 30 {
		return 1
	}
	return 0
}

func (s *automated) checkStochastic(candles []*api.Candle, indicatorID uuid.UUID) float64 {
	if candles[1].IndicatorValues.Stochastics[indicatorID.String()].IndexD > 20 ||
		candles[0].IndicatorValues.Stochastics[indicatorID.String()].IndexK > 20 {
		return 0
	}

	if candles[0].IndicatorValues.Stochastics[indicatorID.String()].IndexK <
		candles[1].IndicatorValues.Stochastics[indicatorID.String()].IndexK {
		return 1
	}

	return 0
}

func (s *automated) checkBollingerBand(candles []*api.Candle, indicatorID uuid.UUID, makerFeeRate, takerFeeRate float64) float64 {
	if candles[0].Low > candles[0].IndicatorValues.BollingerBands[indicatorID.String()].LowerBand {
		return 0
	}

	price := candles[1].Close * (1 + makerFeeRate/100) * (1 + s.MinProfitPercentage/100) * (1 + takerFeeRate/100)
	if price < candles[1].IndicatorValues.BollingerBands[indicatorID.String()].UpperBand {
		return 1
	}
	return 0
}
