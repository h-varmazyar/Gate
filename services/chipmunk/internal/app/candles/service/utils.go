package service

import (
	"context"
	"github.com/google/uuid"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/buffer"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/entity"
	indicatorsPkg "github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/indicators"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
	"time"
)

func (s *Service) validateDownloadPrimaryCandlesRequest(ctx context.Context, req *chipmunkApi.CandleWorkerStartReq) error {
	if req.Resolutions == nil || len(req.Resolutions.Elements) == 0 {
		err := errors.New(ctx, codes.FailedPrecondition).AddDetailF("invalid resolutions for Platform %v", req.Platform)
		s.logger.WithError(err).Errorf("failed to start candles candleReaderWorker")
		return err
	}
	if req.Markets == nil || len(req.Markets.Elements) == 0 {
		err := errors.New(ctx, codes.FailedPrecondition).AddDetailF("invalid markets for Platform %v", req.Platform)
		s.logger.WithError(err).Errorf("failed to start candles candleReaderWorker")
		return err
	}
	return nil
}

func (s *Service) preparePrimaryDataRequests(platform api.Platform, market *chipmunkApi.Market, resolutions *chipmunkApi.Resolutions, indicators []indicatorsPkg.Indicator) {
	for _, resolution := range resolutions.Elements {
		s.preparePrimaryDataRequestsByResolution(platform, market, resolution, indicators)
	}
}

func (s *Service) preparePrimaryDataRequestsByResolution(platform api.Platform, market *chipmunkApi.Market, resolution *chipmunkApi.Resolution, indicators []indicatorsPkg.Indicator) {
	from, err := s.prepareLocalCandles(market, resolution, indicators)
	if err != nil {
		return
	}

	s.makePrimaryDataRequests(platform, market, resolution, from)
}

func (s *Service) prepareLocalCandles(market *chipmunkApi.Market, resolution *chipmunkApi.Resolution, indicators []indicatorsPkg.Indicator) (time.Time, error) {
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
		if err = s.calculateIndicators(candles, indicators); err != nil {
			s.logger.WithError(err).Errorf("failed to calculate indicators for market %v in resolution %v", marketID, resolutionID)
			return time.Unix(0, 0), err
		}
		from = candles[len(candles)-1].Time.Add(time.Duration(resolution.Duration))

		for _, candle := range candles {
			buffer.CandleBuffer.Push(candle)
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

func (s *Service) calculateIndicators(candles []*entity.Candle, indicators []indicatorsPkg.Indicator) error {
	for _, indicator := range indicators {
		err := indicator.Calculate(candles)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) makePrimaryDataRequests(platform api.Platform, market *chipmunkApi.Market, resolution *chipmunkApi.Resolution, from time.Time) {
	_, err := s.functionsService.AsyncOHLC(context.Background(), &coreApi.OHLCReq{
		Resolution: resolution,
		Market:     market,
		From:       from.Unix(),
		To:         time.Now().Unix(),
		Platform:   platform,
	})
	if err != nil {
		s.logger.WithError(err).Errorf("failed to create primary async OHLC request %v in resolution %v and Platform %v", market.Name, resolution.Duration, platform)
	}
}

func (s *Service) loadIndicators(indicators []*chipmunkApi.Indicator) ([]indicatorsPkg.Indicator, error) {
	response := make([]indicatorsPkg.Indicator, 0)
	for _, i := range indicators {
		var indicatorCalculator indicatorsPkg.Indicator
		var err error
		indicator := new(entity.Indicator)
		mapper.Struct(i, indicator)
		indicator.ID, err = uuid.Parse(i.ID)
		if err != nil {
			return nil, err
		}
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
		response = append(response, indicatorCalculator)
	}
	return response, nil
}
