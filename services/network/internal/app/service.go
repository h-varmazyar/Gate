package app

import (
	"context"
	"encoding/json"
	networkAPI "github.com/mrNobody95/Gate/services/network/api"
	"github.com/mrNobody95/Gate/services/network/internal/pkg/requests"
	"google.golang.org/grpc"
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
* Date: 10.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

func NewService() Service {
	return Service{}
}

type Service struct{}

func (s Service) RegisterServer(server *grpc.Server) {
	networkAPI.RegisterRequestServiceServer(server, s)
}

func (s Service) Do(ctx context.Context, req *networkAPI.Request) (*networkAPI.Response, error) {
	request := requests.New(req.Type, req.Endpoint)
	if err := request.SetAuth(req.Auth); err != nil {
		return nil, err
	}
	if err := request.SetHeaders(req.Headers); err != nil {
		return nil, err
	}
	if err := request.SetParams(req.Params); err != nil {
		return nil, err
	}
	responseBody, err := request.Do()
	if err != nil {
		return nil, err
	}
	response := new(networkAPI.Response)
	return response, json.Unmarshal(responseBody, response)
}
