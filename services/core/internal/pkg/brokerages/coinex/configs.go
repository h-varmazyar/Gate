package coinex

type Configs struct {
	CoinexCallbackQueue string `yaml:"coinex_callback_queue"`
	ChipmunkOHLCQueue   string `yaml:"chipmunk_ohlc_queue"`
}
