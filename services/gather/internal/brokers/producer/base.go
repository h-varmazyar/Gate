package producer

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

type Producer struct {
	logger *log.Logger
	nc     *nats.Conn
}

func NewProducer(logger *log.Logger, nc *nats.Conn) *Producer {
	return &Producer{
		logger: logger,
		nc:     nc,
	}
}

func (p *Producer) PublishCandleUpdates(payload CandlePayload) error {
	subject := fmt.Sprintf("candles.market.%v.%v", payload.MarketID, payload.ResolutionID)
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err := p.nc.Publish(subject, data); err != nil {
		return err
	}
	p.logger.Infof("published: %v", string(data))
	return nil
}

func (p *Producer) PublishTicker(payload TickerPayload) error {
	subject := fmt.Sprintf("tickers.market.%v", payload.MarketID)
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err := p.nc.Publish(subject, data); err != nil {
		return err
	}
	p.logger.Infof("published: %v", string(data))
	return nil
}

func (p *Producer) PublishPost(payload PostPayload) error {
	subject := fmt.Sprintf("post.provider.%v", payload.Provider)
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err := p.nc.Publish(subject, data); err != nil {
		return err
	}
	p.logger.Infof("published: %v", string(data))
	return nil
}
