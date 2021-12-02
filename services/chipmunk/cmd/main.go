package main

import (
	"github.com/mrNobody95/Gate/pkg/service"
	"github.com/mrNobody95/Gate/services/chipmunk/configs"
	"github.com/mrNobody95/Gate/services/chipmunk/internal/app/ohlc"
	"google.golang.org/grpc"
	"net"
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

func main() {
	service.Serve(configs.Variables.GrpcAddresses.Chipmunk, func(lst net.Listener) error {
		server := grpc.NewServer()
		ohlc.NewService().RegisterServer(server)
		return server.Serve(lst)
	})

	service.Start(configs.Variables.ServiceName, configs.Variables.Version)
}
