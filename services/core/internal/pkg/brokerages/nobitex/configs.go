package nobitex

type Configs struct {
	NobitexCallbackQueue        string `yaml:"nobitex_callback_queue"`
	ChipmunkOHLCQueue           string `yaml:"chipmunk_ohlc_queue"`
	NobitexPublicRateLimiterID  string `yaml:"nobitex_public_rate_limiter_id"`
	NobitexSpotApiRateLimiterID string `yaml:"nobitex_spot_api_rate_limiter_id"`
}
