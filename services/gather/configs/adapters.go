package configs

type CoreAdapter struct {
	GrpcAddress string
}

type CoinexAdapter struct {
	BaseURL           string `json:"base_url"`
	SocksProxyAddress string `json:"socks_proxy_address"`
}
