package automatedStrategy

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/eagle/configs"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/buffers"
	"github.com/h-varmazyar/Gate/services/eagle/internal/pkg/repository"
	"google.golang.org/grpc/codes"
)

type automated struct {
	*repository.Strategy
	makerFeeRate   float64
	takerFeeRate   float64
	walletsService chipmunkApi.WalletsServiceClient
}

func NewAutomatedStrategy(ctx context.Context, strategy *repository.Strategy, makerFeeRate, takerFeeRate float64) (*automated, error) {
	if makerFeeRate == 0 || takerFeeRate == 0 {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_fee_rate")
	}
	if strategy == nil {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "empty_strategy")
	}
	automated := new(automated)
	chipmunkConn := grpcext.NewConnection(configs.Variables.GrpcAddresses.Chipmunk)
	automated.walletsService = chipmunkApi.NewWalletsServiceClient(chipmunkConn)
	return automated, nil
}

func (s *automated) CheckForSignals(ctx context.Context, marketID uuid.UUID, marketName string) error {
	wallet, err := s.walletsService.ReturnByName(context.Background(), &chipmunkApi.ReturnWalletByDestReq{Destination: marketName})
	if err != nil {
		return errors.Cast(ctx, err).AddDetailF("failed to fetch wallet info for %v", marketID)
	}
	if wallet.ActiveBalance <= 0 {
		return nil
	}
	strength := float64(0)
	for _, strategyIndicator := range s.Indicators {
		switch strategyIndicator.Type {
		case chipmunkApi.IndicatorType_RSI:
			strength += s.checkRSI(marketID, strategyIndicator.IndicatorID)
		case chipmunkApi.IndicatorType_Stochastic:
			strength += s.checkStochastic(marketID, strategyIndicator.IndicatorID)
		case chipmunkApi.IndicatorType_BollingerBands:
			strength += s.checkBollingerBand(marketID, strategyIndicator.IndicatorID)
		}
	}
	strength /= 3
	if strength >= 0.9 {
		//buy
	}

}

func (s *automated) checkRSI(marketID, indicatorID uuid.UUID) float64 {
	candles := buffers.Markets.GetLastNCandles(marketID, 2)
	if candles[0].IndicatorValues.RSIs[indicatorID.String()].RSI < 30 &&
		candles[1].IndicatorValues.RSIs[indicatorID.String()].RSI >= 30 {
		return 1
	}
	return 0
}

func (s *automated) checkStochastic(marketID, indicatorID uuid.UUID) float64 {
	candles := buffers.Markets.GetLastNCandles(marketID, 2)
	if len(candles) != 2 {
		return 0
	}
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

func (s *automated) checkBollingerBand(marketID, indicatorID uuid.UUID) float64 {
	candles := buffers.Markets.GetLastNCandles(marketID, 2)
	if len(candles) != 2 {
		return 0
	}
	if candles[0].Low > candles[0].IndicatorValues.BollingerBands[indicatorID.String()].LowerBand {
		return 0
	}

	price := candles[1].Close * (1 + s.makerFeeRate/100) * float64(1+s.MinProfitPercentage/100) * (1 + s.takerFeeRate/100)
	if price < candles[1].IndicatorValues.BollingerBands[indicatorID.String()].UpperBand {
		return 1
	}
	return 0
}
