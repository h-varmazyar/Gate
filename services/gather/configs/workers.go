package configs

import "time"

type WorkerMarketUpdate struct {
	RunningTime string
	S3Bucket    string
	S3AccessKey string
	S3SecretKey string
	S3Endpoint  string
	S3Region    string
}

type WorkerLastCandle struct {
	RunningInterval time.Duration
}

type WorkerSahamyabArchive struct {
	SocksProxyAddress string
	GeminiAPIKey      string
}

type WorkerSahamyabStream struct {
}

type WorkerTicker struct {
	RunningInterval time.Duration
}

type WorkerWarmup struct {
	NeedWarmup bool
}
