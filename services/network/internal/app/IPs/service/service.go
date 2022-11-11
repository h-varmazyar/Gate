package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/Gate/pkg/mapper"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	"github.com/h-varmazyar/Gate/services/network/internal/app/IPs/repository"
	"github.com/h-varmazyar/Gate/services/network/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Service struct {
	logger  *log.Logger
	configs *Configs
	db      repository.IPRepository
}

var (
	GrpcService *Service
)

func NewService(_ context.Context, logger *log.Logger, configs *Configs, db repository.IPRepository) *Service {
	if GrpcService == nil {
		GrpcService = new(Service)
		GrpcService.logger = logger
		GrpcService.configs = configs
		GrpcService.db = db
	}
	return GrpcService
}

func (s *Service) RegisterServer(server *grpc.Server) {
	networkAPI.RegisterIPServiceServer(server, s)
}

func (s *Service) Create(_ context.Context, req *networkAPI.IPCreateReq) (*networkAPI.IP, error) {
	ip := new(entity.IP)
	mapper.Struct(req, ip)
	err := s.db.Create(ip)
	if err != nil {
		return nil, err
	}
	response := new(networkAPI.IP)
	mapper.Struct(ip, response)
	return response, nil
}

func (s *Service) Return(_ context.Context, req *networkAPI.IPReturnReq) (*networkAPI.IP, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}

	var ip *entity.IP
	ip, err = s.db.Return(id)
	if err != nil {
		return nil, err
	}
	response := new(networkAPI.IP)
	mapper.Struct(ip, response)
	return response, nil
}

func (s *Service) List(_ context.Context, _ *networkAPI.IPListReq) (*networkAPI.IPs, error) {
	ips, err := s.db.List()
	if err != nil {
		return nil, err
	}
	response := new(networkAPI.IPs)
	mapper.Slice(ips, &response.Elements)
	return response, nil
}
