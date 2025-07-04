package configs

import "time"

type WorkerMarketUpdate struct {
	Running     bool
	NeedWarmup  bool
	RunningTime string
	S3Bucket    string
	S3AccessKey string
	S3SecretKey string
	S3Endpoint  string
	S3Region    string
}

type WorkerLastCandle struct {
	Running         bool
	RunningInterval time.Duration
}

type WorkerSahamyabArchive struct {
	Running bool
}

type WorkerPostSentimentCheck struct {
	Running           bool
	SocksProxyAddress string
	GeminiAPIKey      string
}

type WorkerSahamyabStream struct {
	Running bool
}

type WorkerTicker struct {
	Running         bool
	RunningInterval time.Duration
}
