package service

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	indicatorsPkg "github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	eagleApi "github.com/h-varmazyar/Gate/services/eagle/api/proto"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"time"
)

func (s *Service) validateDownloadPrimaryCandlesRequest(ctx context.Context, req *chipmunkApi.DownloadPrimaryCandlesReq) error {
	if req.Resolutions == nil || len(req.Resolutions.Elements) == 0 {
		err := errors.New(ctx, codes.FailedPrecondition).AddDetailF("invalid resolutions for Platform %v", req.Platform)
		s.logger.WithError(err).Errorf("failed to start candles primaryDataWorker")
		return err
	}
	if req.Markets == nil || len(req.Markets.Elements) == 0 {
		err := errors.New(ctx, codes.FailedPrecondition).AddDetailF("invalid markets for Platform %v", req.Platform)
		s.logger.WithError(err).Errorf("failed to start candles primaryDataWorker")
		return err
	}
	return nil
}

func (s *Service) prepareDownloadPrimaryCandles(req *chipmunkApi.DownloadPrimaryCandlesReq) (uuid.UUID, error) {
	strategyID, err := uuid.Parse(req.StrategyID)
	if err != nil {
		s.logger.WithError(err).Errorf("failed to load strategy id for Platform %v", req.Platform)
		return uuid.Nil, err
	}
	if !s.primaryDataWorker.Started {
		s.primaryDataWorker.Start()
	}
	return strategyID, nil
}

func (s *Service) preparePrimaryDataRequests(platform api.Platform, market *chipmunkApi.Market, resolutions *chipmunkApi.Resolutions, strategyID uuid.UUID) {
	for _, resolution := range resolutions.Elements {
		s.preparePrimaryDataRequestsByResolution(platform, market, resolution, strategyID)
	}
}

func (s *Service) preparePrimaryDataRequestsByResolution(platform api.Platform, market *chipmunkApi.Market, resolution *chipmunkApi.Resolution, strategyID uuid.UUID) {
	from, err := s.prepareLocalCandles(strategyID, market, resolution)
	if err != nil {
		return
	}

	s.makePrimaryDataRequests(platform, market, resolution, from)
}

func (s *Service) prepareLocalCandles(strategyID uuid.UUID, market *chipmunkApi.Market, resolution *chipmunkApi.Resolution) (time.Time, error) {
	marketID, err := uuid.Parse(market.ID)
	if err != nil {
		s.logger.WithError(err).Errorf("invalid market id %v", market)
		return time.Unix(0, 0), err
	}
	resolutionID, err := uuid.Parse(resolution.ID)
	if err != nil {
		s.logger.WithError(err).Errorf("invalid resolution id %v", resolution)
		return time.Unix(0, 0), err
	}
	var from time.Time
	candles, err := s.loadLocalCandles(marketID, resolutionID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			from = time.Unix(market.IssueDate, 0)
		} else {
			s.logger.WithError(err).Errorf("failed to load local candles for market %v in resolution %v", marketID, resolutionID)
			return time.Unix(0, 0), err
		}
	}

	for _, candle := range candles {
		candle.IndicatorValues = entity.NewIndicatorValues()
	}

	if len(candles) > 0 {
		if err = s.calculateIndicators(candles, strategyID); err != nil {
			s.logger.WithError(err).Errorf("failed to calculate indicators for market %v in resolution %v", marketID, resolutionID)
			return time.Unix(0, 0), err
		}
		from = candles[len(candles)-1].Time.Add(time.Duration(resolution.Duration))

		for _, candle := range candles {
			s.buffer.Push(candle)
		}
	} else {
		from = time.Unix(market.IssueDate, 0)
	}
	return from, nil
}

func (s *Service) loadLocalCandles(marketID, resolutionID uuid.UUID) ([]*entity.Candle, error) {
	end := false
	limit := 1000000
	candles := make([]*entity.Candle, 0)
	for offset := 0; !end; offset += limit {
		list, err := s.db.ReturnList(marketID, resolutionID, limit, offset)
		if err != nil {
			s.logger.WithError(err).Errorf("failed to fetch candle list")
			return nil, err
		}
		if len(list) < limit {
			end = true
		}
		candles = append(candles, list...)
	}
	return candles, nil
}

func (s *Service) calculateIndicators(candles []*entity.Candle, strategyID uuid.UUID) error {
	strategyIndicators, err := s.strategyService.Indicators(context.Background(), &eagleApi.StrategyIndicatorReq{StrategyID: strategyID.String()})
	if err != nil {
		return err
	}
	loadedIndicators, err := s.loadIndicators(strategyIndicators)
	if err != nil {
		return err
	}

	for _, indicator := range loadedIndicators {
		err := indicator.Calculate(candles)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) makePrimaryDataRequests(platform api.Platform, market *chipmunkApi.Market, resolution *chipmunkApi.Resolution, from time.Time) {
	for end := false; !end; {
		to := from.Add(time.Duration(resolution.Duration * 1000))
		if to.After(time.Now()) {
			to = time.Now()
			end = true
		}
		_, err := s.functionsService.AsyncOHLC(context.Background(), &coreApi.OHLCReq{
			Resolution: resolution,
			Market:     market,
			From:       from.Unix(),
			To:         to.Unix(),
			Platform:   platform,
		})
		if err != nil {
			s.logger.WithError(err).Errorf("failed to create async OHLC request for marker %v in resolution %v and Platform %v", market.Name, resolution.Duration, platform)
		}
		from = to.Add(time.Duration(resolution.Duration))
	}
}

func (s *Service) loadIndicators(strategyIndicators *eagleApi.StrategyIndicators) (map[uuid.UUID]indicatorsPkg.Indicator, error) {
	response := make(map[uuid.UUID]indicatorsPkg.Indicator)
	for _, strategyIndicator := range strategyIndicators.Elements {
		indicatorResp, err := s.indicatorService.Return(context.Background(), &chipmunkApi.IndicatorReturnReq{ID: strategyIndicator.IndicatorID})
		indicator := new(entity.Indicator)
		mapper.Struct(indicatorResp, indicator)
		if err != nil {
			return nil, err
		}
		var indicatorCalculator indicatorsPkg.Indicator
		switch indicator.Type {
		case chipmunkApi.Indicator_RSI:
			indicatorCalculator, err = indicatorsPkg.NewRSI(indicator.ID, indicator.Configs.RSI)
		case chipmunkApi.Indicator_Stochastic:
			indicatorCalculator, err = indicatorsPkg.NewStochastic(indicator.ID, indicator.Configs.Stochastic)
		case chipmunkApi.Indicator_MovingAverage:
			indicatorCalculator, err = indicatorsPkg.NewMovingAverage(indicator.ID, indicator.Configs.MovingAverage)
		case chipmunkApi.Indicator_BollingerBands:
			indicatorCalculator, err = indicatorsPkg.NewBollingerBands(indicator.ID, indicator.Configs.BollingerBands)
		}
		if err != nil {
			return nil, err
		}
		response[indicator.ID] = indicatorCalculator
	}
	return response, nil
}
