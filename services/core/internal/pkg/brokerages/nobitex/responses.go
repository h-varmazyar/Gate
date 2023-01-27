package nobitex

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	api "github.com/h-varmazyar/Gate/api/proto"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	networkAPI "github.com/h-varmazyar/Gate/services/network/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Response struct {
	configs   *Configs
	ohlcQueue *amqpext.Queue
}

func NewResponse(configs *Configs, isAsync bool) (*Response, error) {
	r := &Response{
		configs: configs,
	}
	if isAsync {
		ohlcQueue, err := amqpext.Client.QueueDeclare(configs.ChipmunkOHLCQueue)
		if err != nil {
			return nil, err
		}
		r.ohlcQueue = ohlcQueue
	}
	return r, nil
}

func (r *Response) AsyncOHLC(_ context.Context, response *networkAPI.Response, metadata *brokerages.Metadata) {
	if response.Code != http.StatusOK {
		log.Errorf("ohlc request failed with code: %v - %v", response.Code, response.Body)
		return
	}
	data := struct {
		S string
		T []int64
		O []float64
		H []float64
		L []float64
		C []float64
		V []float64
	}{}
	if err := json.Unmarshal([]byte(response.Body), &data); err != nil {
		return
	}
	if data.S != "ok" {
		return
	}

	candles := make([]*chipmunkApi.Candle, 0)
	for i := 0; i < len(data.V); i++ {
		c := new(chipmunkApi.Candle)
		c.Time = data.T[i]
		c.Open = data.O[i]
		c.Close = data.C[i]
		c.High = data.H[i]
		c.Low = data.L[i]
		c.Volume = data.V[i]
		c.ResolutionID = metadata.ResolutionID
		c.MarketID = metadata.MarketID
		candles = append(candles, c)
	}
	message := &chipmunkApi.Candles{
		Elements: candles,
		Count:    int64(len(candles)),
	}
	if message.Count > 0 {
		if bytes, err := proto.Marshal(message); err != nil {
			log.WithError(err).Errorf("faled to marshal candles")
			return
		} else {
			if publishErr := r.ohlcQueue.Publish(bytes, grpcext.ProtobufContentType); publishErr != nil {
				log.WithError(err).Errorf("faled to publish candles")
			}
		}
	}
}

func (r *Response) AllMarkerStatistics(ctx context.Context, response *networkAPI.Response) (*coreApi.AllMarketStatisticsResp, error) {
	if response.Code != http.StatusOK {
		return nil, errors.New(ctx, codes.Canceled).AddDetails(response.Body)
	}
	data := struct {
		Status string
		Stats  map[string]struct {
			IsClosed   bool   `json:"isClosed"`
			BestSell   string `json:"bestSell"`
			BestBuy    string `json:"bestBuy"`
			VolumeSrc  string `json:"volumeSrc"`
			VolumeDest string `json:"volumeDst"`
			Latest     string `json:"latest"`
			DayLow     string `json:"dayLow"`
			DayHigh    string `json:"dayHigh"`
			DayOpen    string `json:"dayOpen"`
			DayClose   string `json:"dayClose"`
			DayChange  string `json:"dayChange"`
		}
	}{}
	if err := json.Unmarshal([]byte(response.Body), &data); err != nil {
		return nil, err
	}
	if data.Status != "ok" {
		return nil, errors.New(ctx, codes.Unknown)
	}
	date := time.Now().Unix()
	resp := new(coreApi.AllMarketStatisticsResp)
	resp.Platform = api.Platform_Nobitex
	resp.Date = date
	resp.AllStatistics = make(map[string]*coreApi.MarketStatistics)

	var err error
	for key, value := range data.Stats {
		marketStatistics := new(coreApi.MarketStatistics)
		marketStatistics.Date = date
		if marketStatistics.Volume, err = strconv.ParseFloat(value.VolumeSrc, 64); err != nil {
			log.WithError(err).Error("failed to parse volume")
			return nil, err
		}
		if marketStatistics.Close, err = strconv.ParseFloat(value.DayClose, 64); err != nil {
			log.WithError(err).Error("failed to parse close")
			return nil, err
		}
		if marketStatistics.Open, err = strconv.ParseFloat(value.DayOpen, 64); err != nil {
			log.WithError(err).Error("failed to parse open")
			return nil, err
		}
		if marketStatistics.High, err = strconv.ParseFloat(value.DayHigh, 64); err != nil {
			log.WithError(err).Error("failed to parse high")
			return nil, err
		}
		if marketStatistics.Low, err = strconv.ParseFloat(value.DayLow, 64); err != nil {
			log.WithError(err).Error("failed to parse low")
			return nil, err
		}
		resp.AllStatistics[strings.ToUpper(strings.Trim(key, "-"))] = marketStatistics
	}
	return resp, nil
}

func (r *Response) GetMarketInfo(ctx context.Context, response *networkAPI.Response) (*coreApi.MarketInfo, error) {
	return nil, errors.New(ctx, codes.Unimplemented)
}

func (r *Response) WalletsBalance(ctx context.Context, response *networkAPI.Response) (*chipmunkApi.Wallets, error) {
	if response.Code != http.StatusOK {
		return nil, errors.New(ctx, codes.Canceled).AddDetails(response.Body)
	}
	data := struct {
		Status  string
		Wallets []struct {
			Currency       string `json:"currency"`
			Balance        string `json:"balance"`
			BlockedBalance string `json:"blockedBalance"`
			ActiveBalance  string `json:"activeBalance"`
		}
	}{}

	if err := json.Unmarshal([]byte(response.Body), &data); err != nil {
		return nil, err
	}
	if data.Status != "ok" {
		return nil, errors.New(ctx, codes.Unknown)
	}
	wallets := make([]*chipmunkApi.Wallet, 0)
	var err error
	for _, wallet := range data.Wallets {
		w := &chipmunkApi.Wallet{}
		w.AssetName = wallet.Currency
		if w.ActiveBalance, err = strconv.ParseFloat(wallet.Balance, 64); err != nil {
			log.WithError(err).Error("failed to parse open")
			return nil, err
		}
		if w.BlockedBalance, err = strconv.ParseFloat(wallet.BlockedBalance, 64); err != nil {
			log.WithError(err).Error("failed to parse open")
			return nil, err
		}
		if w.TotalBalance, err = strconv.ParseFloat(wallet.Balance, 64); err != nil {
			log.WithError(err).Error("failed to parse open")
			return nil, err
		}
		wallets = append(wallets, w)
	}
	return &chipmunkApi.Wallets{
		Elements: wallets,
		Count:    int64(len(wallets)),
	}, nil
}
