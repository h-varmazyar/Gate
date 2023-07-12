package callbacks

import (
	"github.com/golang/protobuf/proto"
	"github.com/h-varmazyar/Gate/pkg/amqpext"
	"github.com/h-varmazyar/Gate/pkg/grpcext"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	coinexCallback "github.com/h-varmazyar/Gate/services/core/internal/pkg/brokerages/coinex/callbacks"
	log "github.com/sirupsen/logrus"
)

const (
	ChipmunkOHLCQueue = "chipmunk_ohlc"
)

type OHLCCallbackHandler struct {
	ohlc              chan *coreApi.OHLCResponse
	chipmunkOHLCQueue *amqpext.Queue
}

func ListenOHLCCallbacks() error {
	ohlcQueue, err := amqpext.Client.QueueDeclare(ChipmunkOHLCQueue)
	if err != nil {
		log.WithError(err).Error("failed to declare coinex callback queue")
		return err
	}
	c := &OHLCCallbackHandler{
		ohlc:              make(chan *coreApi.OHLCResponse, 1000000),
		chipmunkOHLCQueue: ohlcQueue,
	}
	err = coinexCallback.ListenOHLCCallback(c.ohlc)
	if err != nil {
		return err
	}

	go c.listenToOHLC()

	return nil
}

func (c *OHLCCallbackHandler) listenToOHLC() {
	for ohlcResp := range c.ohlc {
		bytes, err := proto.Marshal(ohlcResp)
		if err != nil {
			log.WithError(err).Errorf("faled to marshal coinex async ohls message")
			return
		}

		err = c.chipmunkOHLCQueue.Publish(bytes, grpcext.ProtobufContentType)
		if err != nil {
			log.WithError(err).Errorf("faled to publish coinex async ohlc")
			return
		}
	}
}
