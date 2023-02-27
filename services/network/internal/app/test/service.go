package test

import (
	"context"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
)

func TestNetwork(address string) {
	networkConn := grpcext.NewConnection(address)

	requestService := networkAPI.NewRequestServiceClient(networkConn)

	resp, err := requestService.Do(context.Background(), &networkAPI.Request{
		Method:   networkAPI.Request_GET,
		Endpoint: "https://api.podro.com/back4front/scms/products/WslLytH0",
		Type:     networkAPI.Request_Sync,
	})
	if err != nil {
		log.WithError(err).Error("failed to net do test")
		return
	}
	log.Infof("resp: %v", resp)
}
