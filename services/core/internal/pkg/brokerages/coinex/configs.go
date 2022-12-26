package coinex

type Configs struct {
	CoinexCallbackQueue        string `yaml:"coinex_callback_queue"`
	ChipmunkOHLCQueue          string `yaml:"chipmunk_ohlc_queue"`
	CoinexPublicRateLimiterID  string `yaml:"coinex_public_rate_limiter_id"`
	CoinexSpotApiRateLimiterID string `yaml:"coinex_spot_api_rate_limiter_id"`
}
