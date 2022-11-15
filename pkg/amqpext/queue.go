package amqpext

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type Queue struct {
	client     *client
	key        string
	name       string
	exchange   string
	deliveries chan amqp.Delivery
	done       chan error
}

func (q *Queue) Consume(consumer string) <-chan amqp.Delivery {
	q.client.locker.RLock()
	defer q.client.locker.RUnlock()
	q.deliveries = make(chan amqp.Delivery)
	go q.pickup(consumer)
	return q.deliveries
}

func (q *Queue) Publish(bin []byte, contentType string) error {
	q.client.locker.RLock()
	defer q.client.locker.RUnlock()
	return q.client.channel.Publish(q.exchange, q.key, false, false, amqp.Publishing{
		ContentType:  contentType,
		DeliveryMode: amqp.Transient,
		Timestamp:    time.Now(),
		Body:         bin,
	})
}

func (q *Queue) pickup(consumer string) {
	//for range time.NewTicker(time.Second).C {
	if q.client.isSuspend {
		close(q.deliveries)
		return
	}
	q.client.locker.Lock()
	channel, err := q.client.channel.Consume(q.name, consumer, false, false, false, false, nil)
	if err != nil {
		log.WithError(err).WithField("queue name", q.name).WithField("consumer", consumer).Error("cannot consume from rabbitMQ")
	}
	q.client.locker.Unlock()
	for delivery := range channel {
		q.deliveries <- delivery
	}
	q.done <- nil
	//}
}

func (q *Queue) Close(consumer string) error {
	if err := q.client.channel.Cancel(consumer, true); err != nil {
		return err
	}
	return <-q.done
}
