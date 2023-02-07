package service

type Configs struct {
	NetworkAddress string `mapstructure:"network_address"`
	EagleAddress   string `mapstructure:"eagle_address"`
	CoreAddress    string `mapstructure:"core_address"`
}
