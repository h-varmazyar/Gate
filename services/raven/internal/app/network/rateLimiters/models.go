package rateLimiters

type RateLimiter struct {
	RequestCountLimit int64  `json:"request_count_limit"`
	TimeLimit         int64  `json:"time_limit"`
	Type              string `json:"type"  enums:"Spread,Immediate"`
}
