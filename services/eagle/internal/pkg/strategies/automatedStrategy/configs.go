package automatedStrategy

type Configs struct {
	RedisAddress       string `mapstructure:"redis_address"`
	CoreAddress        string `mapstructure:"core_address"`
	ChipmunkAddress    string `mapstructure:"chipmunk_address"`
	TelegramBotAddress string `mapstructure:"telegram_bot_address"`
	BroadcastChannelID int64  `mapstructure:"broadcast_channel_id"`
}
