package coinex

type Configs struct {
	CoinexPublicRateLimiterID  string `mapstructure:"coinex_public_rate_limiter_id"`
	CoinexSpotApiRateLimiterID string `mapstructure:"coinex_spot_api_rate_limiter_id"`
}
