package coinex

type Configs struct {
	CoinexCallbackQueue       string `yaml:"coinex_callback_queue"`
	ChipmunkOHLCQueue         string `yaml:"chipmunk_ohlc_queue"`
	CoinexPublicRateLimiterID string `yaml:"coinex_public_rate_limiter_id"`
}
