package core

import (
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/gather/configs"
)

type Core struct {
	functionService coreApi.FunctionsServiceClient
}

func NewCore(cfg configs.CoreAdapter) Core {
	coreConnection := grpcext.NewConnection(cfg.GrpcAddress)
	return Core{
		functionService: coreApi.NewFunctionsServiceClient(coreConnection),
	}
}
