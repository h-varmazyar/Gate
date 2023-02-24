package coinex

type Configs struct {
	CoinexCallbackQueue        string `mapstructure:"coinex_callback_queue"`
	ChipmunkOHLCQueue          string `mapstructure:"chipmunk_ohlc_queue"`
	CoinexPublicRateLimiterID  string `mapstructure:"coinex_public_rate_limiter_id"`
	CoinexSpotApiRateLimiterID string `mapstructure:"coinex_spot_api_rate_limiter_id"`
}
