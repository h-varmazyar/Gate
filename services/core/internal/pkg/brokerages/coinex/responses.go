package coinex

import (
	"github.com/golang/protobuf/proto"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api"
	"github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Response struct {
	configs   *Configs
	ohlcQueue *amqpext.Queue
}

func NewResponse(configs *Configs) (*Response, error) {
	ohlcQueue, err := amqpext.Client.QueueDeclare(configs.ChipmunkOHLCQueue)
	if err != nil {
		return nil, err
	}
	r := &Response{
		configs:   configs,
		ohlcQueue: ohlcQueue,
	}
	return r, nil
}

func (r *Response) OHLC(response *brokerages.Response) {
	if response.Code != http.StatusOK {
		log.Errorf("ohlc request failed with code: %v - %v", response.Code, response.Body)
		return
	}
	data := make([][]interface{}, 0)
	if err := parseResponse(response.Body, &data); err != nil {
		log.WithError(err).Errorf("ohlc request parse failed: %v", response.Body)
		return
	}
	candles := make([]*chipmunkApi.Candle, 0)
	for _, item := range data {
		c := new(chipmunkApi.Candle)
		c.Time = int64(item[0].(float64))
		c.Open, _ = strconv.ParseFloat(item[1].(string), 64)
		c.Close, _ = strconv.ParseFloat(item[2].(string), 64)
		c.High, _ = strconv.ParseFloat(item[3].(string), 64)
		c.Low, _ = strconv.ParseFloat(item[4].(string), 64)
		c.Volume, _ = strconv.ParseFloat(item[5].(string), 64)
		c.Amount, _ = strconv.ParseFloat(item[6].(string), 64)
		candles = append(candles, c)
	}
	message := &chipmunkApi.Candles{
		Elements: candles,
		Count:    int64(len(candles)),
	}
	if bytes, err := proto.Marshal(message); err != nil {
		log.WithError(err).Errorf("faled to marshal candles")
		return
	} else {
		if publishErr := r.ohlcQueue.Publish(bytes, grpcext.ProtobufContentType); publishErr != nil {
			log.WithError(err).Errorf("faled to publish candles")
		}
	}
}
