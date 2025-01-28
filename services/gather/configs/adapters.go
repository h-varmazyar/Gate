package configs

type CoreAdapter struct {
	GrpcAddress string
}

type CoinexAdapter struct {
	BaseURL string `json:"base_url"`
}
