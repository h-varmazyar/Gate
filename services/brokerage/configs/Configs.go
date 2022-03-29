package configs

type Configs struct {
	GrpcPort           uint16 `env:"GRPC_PORT,required"`
	NetworkGrpcPort    uint16 `env:"NETWORK_GRPC_PORT,required"`
	ChipmunkAddress    string `env:"CHIPMUNK_ADDRESS,required"`
	HttpPort           uint16 `env:"HTTP_PORT,required"`
	MaxLogsPerPage     int64  `env:"MAX_LOGS_PER_PAGE,required"`
	DatabaseConnection string `env:"DATABASE_CONNECTION,required,file"`
}
