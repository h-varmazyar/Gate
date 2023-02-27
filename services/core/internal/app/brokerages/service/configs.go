package service

type Configs struct {
	ChipmunkGrpcAddress string `mapstructure:"chipmunk_grpc_address"`
	EagleGrpcAddress    string `mapstructure:"eagle_grpc_address"`
}
