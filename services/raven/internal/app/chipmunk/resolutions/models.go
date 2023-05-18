package resolutions

type ResolutionSetReq struct {
	Platform string `enums:"Coinex,UnknownBrokerage,Nobitex,Mazdax,Binance"`
	Duration int64
	Label    string
	Value    string
}
