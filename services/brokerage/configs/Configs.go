package configs

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
* Date: 19.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type Configs struct {
	GrpcPort           uint16 `env:"GRPC_PORT,required"`
	NetworkGrpcPort    uint16 `env:"NETWORK_GRPC_PORT,required"`
	ChipmunkAddress    string `env:"CHIPMUNK_ADDRESS,required"`
	HttpPort           uint16 `env:"HTTP_PORT,required"`
	MaxLogsPerPage     int64  `env:"MAX_LOGS_PER_PAGE,required"`
	DatabaseConnection string `env:"DATABASE_CONNECTION,required,file"`
}
