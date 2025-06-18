package configs

type CoreAdapter struct {
	GrpcAddress string
}

type CoinexAdapter struct {
	APIBaseURL        string `json:"api_base_url"`
	BaseURL           string `json:"base_url"`
	SocksProxyAddress string `json:"socks_proxy_address"`
}

type SahamyabAdapter struct {
	GuestBaseURL string `json:"guest_base_url"`
}
