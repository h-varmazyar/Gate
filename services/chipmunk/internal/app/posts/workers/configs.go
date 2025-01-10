package workers

type Configs struct {
	Running                         bool   `mapstructure:"running"`
	NetworkAddress                  string `mapstructure:"network_address"`
	SahamyabPostCollectorURL        string `mapstructure:"sahamyab_post_collector_url"`
	MaxSentimentDetectorTokenLength int64  `mapstructure:"max_sentiment_detector_token_length"`
	SentimentDetectorToken          string `mapstructure:"sentiment_detector_token"`
}
