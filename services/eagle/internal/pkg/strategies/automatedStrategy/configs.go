package automatedStrategy

type Configs struct {
	RedisAddress       string `yaml:"redis_address"`
	CoreAddress        string `yaml:"core_address"`
	ChipmunkAddress    string `yaml:"chipmunk_address"`
	TelegramBotAddress string `yaml:"telegram_bot_address"`
	BroadcastChannelID int64  `yaml:"broadcast_channel_id"`
}
