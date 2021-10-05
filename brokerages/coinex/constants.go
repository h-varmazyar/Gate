package coinex

import (
	"github.com/mrNobody95/Gate/models"
)

const (
	ResponseSuccess               = 0
	ResponseError                 = 1
	ResponseParameterError        = 2
	ResponseInternalError         = 3
	ResponseIpNotAllow            = 23
	ResponseAccessIdNotExist      = 24
	ResponseSignatureError        = 25
	ResponseServiceUnavailable    = 35
	ResponseServiceTimeout        = 36
	ResponseUnpairedAccount       = 40
	ResponseTransferToSubRejected = 49
	ResponseInsufficientBalance   = 107
	ResponseForbidTrading         = 115
	ResponseTonceCheckError       = 227
	ResponseOrderNumberNotExist   = 600
	ResponseOtherUserOrder        = 601
	ResponseBellowMinLimit        = 602
	ResponsePriceDeviationLarge   = 606
	ResponseMergeDepthError       = 651
)

const (
	USD  models.Currency = "usd"
	RLS  models.Currency = "rls"
	BTC  models.Currency = "btc"
	ETH  models.Currency = "eth"
	LTC  models.Currency = "ltc"
	USDT models.Currency = "usdt"
	XRP  models.Currency = "xrp"
	BCH  models.Currency = "bch"
	BNB  models.Currency = "bnb"
	EOS  models.Currency = "eos"
	DOGE models.Currency = "doge"
	XLM  models.Currency = "xlm"
	TRX  models.Currency = "trx"
	ADA  models.Currency = "ada"
	XMR  models.Currency = "xmr"
	XEM  models.Currency = "xem"
	IOTA models.Currency = "iota"
	ETC  models.Currency = "etc"
	DASH models.Currency = "dash"
	ZEC  models.Currency = "zec"
	NEO  models.Currency = "neo"
)
