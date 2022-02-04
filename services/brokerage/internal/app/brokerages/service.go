package brokerages

import (
	"context"
	"github.com/mrNobody95/Gate/api"
	"github.com/mrNobody95/Gate/pkg/errors"
	"github.com/mrNobody95/Gate/pkg/grpcext"
	"github.com/mrNobody95/Gate/pkg/mapper"
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/brokerage/configs"
	"github.com/mrNobody95/Gate/services/brokerage/internal/pkg/repository"
	chipmunkApi "github.com/mrNobody95/Gate/services/chipmunk/api"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"strconv"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Service struct {
	ohlcService   chipmunkApi.OhlcServiceClient
	walletService chipmunkApi.WalletsServiceClient
}

var (
	GrpcService *Service
)

func NewService(configs *configs.Configs) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		chipmunkConnection := grpcext.NewConnection(configs.ChipmunkAddress)
		GrpcService.ohlcService = chipmunkApi.NewOhlcServiceClient(chipmunkConnection)
		GrpcService.walletService = chipmunkApi.NewWalletsServiceClient(chipmunkConnection)
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	brokerageApi.RegisterBrokerageServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, brokerage *brokerageApi.CreateBrokerageReq) (*brokerageApi.Brokerage, error) {
	//todo: validation on auth and other fields
	br := new(repository.Brokerage)
	if _, ok := brokerageApi.Names_value[brokerage.Name.String()]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong_name")
	} else {
		br.Name = brokerage.Name.String()
	}
	if _, ok := api.AuthType_value[brokerage.Auth.Type.String()]; !ok {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "wrong_auth_type")
	} else {
		br.AuthType = brokerage.Auth.Type.String()
	}
	br.Status = api.Status_Disable.String()
	resID, err := strconv.Atoi(brokerage.ResolutionID)
	if err != nil {
		return nil, errors.NewWithSlug(ctx, codes.FailedPrecondition, "invalid_resolution")
	}
	br.ResolutionID = uint(resID)
	mapper.Struct(brokerage, br)

	for _, market := range brokerage.Markets.Markets {
		tmp := new(repository.Market)
		tmp.ID = uint(market.ID)
		br.Markets = append(br.Markets, tmp)
	}
	if err := repository.Brokerages.Create(br); err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	mapper.Struct(br, response)
	return response, nil
}

func (s *Service) List(_ context.Context, _ *api.Void) (*brokerageApi.Brokerages, error) {
	bb, err := repository.Brokerages.List()
	if err != nil {
		return nil, err
	}

	response := new(brokerageApi.Brokerages)
	mapper.Slice(bb, &response.Brokerages)

	for i, brokerage := range bb {
		response.Brokerages[i].Status = api.Status(api.Status_value[brokerage.Status])
	}
	return response, err
}

func (s *Service) Get(_ context.Context, req *brokerageApi.BrokerageIDReq) (*brokerageApi.Brokerage, error) {
	brokerage, err := repository.Brokerages.ReturnByID(req.ID)
	if err != nil {
		return nil, err
	}
	response := new(brokerageApi.Brokerage)
	response.Name = brokerageApi.Names(brokerageApi.Names_value[brokerage.Name])
	mapper.Struct(brokerage, response)
	mapper.Slice(brokerage.Markets, &response.Markets.Markets)
	return response, err
}

func (s *Service) GetInternal(_ context.Context, req *brokerageApi.BrokerageIDReq) (*brokerageApi.Brokerage, error) {
	brokerage, err := repository.Brokerages.ReturnByID(req.ID)
	if err != nil {
		return nil, err
	}

	return &brokerageApi.Brokerage{
		ID: uint32(brokerage.ID),
		Auth: &api.Auth{
			Type:      api.AuthType(api.AuthType_value[brokerage.AuthType]),
			Username:  brokerage.Username,
			Password:  brokerage.Password,
			AccessID:  brokerage.AccessID,
			SecretKey: brokerage.SecretKey,
		},
		Name:   brokerageApi.Names(brokerageApi.Names_value[brokerage.Name]),
		Status: api.Status(api.Status_value[brokerage.Status]),
	}, err
}

func (s *Service) Delete(_ context.Context, req *brokerageApi.BrokerageIDReq) (*api.Void, error) {
	if err := repository.Brokerages.Delete(req.ID); err != nil {
		return nil, err
	}
	return new(api.Void), nil
}

func (s *Service) ChangeStatus(ctx context.Context, req *brokerageApi.StatusChangeRequest) (*brokerageApi.BrokerageStatus, error) {
	enables, err := repository.Brokerages.ReturnEnables()
	if err != nil {
		return nil, err
	}
	for _, enable := range enables {
		for _, market := range enable.Markets {
			if _, err := s.ohlcService.CancelWorker(ctx, &chipmunkApi.CancelWorkerRequest{
				ResolutionID: uint32(enable.ResolutionID),
				MarketID:     uint32(market.ID),
			}); err != nil {
				log.WithError(err).WithField("market", market.ID).WithField("brokerage", enable.ID)
			}
		}
	}
	brokerage, err := repository.Brokerages.ReturnByID(req.ID)
	if err != nil {
		return nil, err
	}
	switch brokerage.Status {
	case api.Status_Enable.String():
		brokerage.Status = api.Status_Disable.String()
	case api.Status_Disable.String():
		brokerage.Status = api.Status_Enable.String()
	}
	if err := repository.Brokerages.ChangeStatus(brokerage); err != nil {
		return nil, err
	}
	if brokerage.Status == api.Status_Enable.String() {
		if req.OHLC {
			resolution := new(brokerageApi.Resolution)
			mapper.Struct(brokerage.Resolution, resolution)

			for _, market := range brokerage.Markets {
				m := new(brokerageApi.Market)
				mapper.Struct(market, m)
				if _, err := s.ohlcService.AddMarket(ctx, &chipmunkApi.AddMarketRequest{
					Resolution: resolution,
					Market:     m,
				}); err != nil {
					log.WithError(err).WithField("market", market.ID).WithField("brokerage", brokerage.ID)
				}
			}
		}
		if req.Trading {
			_, err := s.walletService.StartWorker(ctx, &chipmunkApi.StartWorkerRequest{
				BrokerageID: uint32(brokerage.ID),
			})
			if err != nil {
				log.WithError(err).WithField("brokerage", brokerage.ID).Error("failed to start wallet worker")
			}
			//todo: add trading worker
		}
	}
	return &brokerageApi.BrokerageStatus{Status: api.Status(api.Status_value[brokerage.Status])}, nil
}
