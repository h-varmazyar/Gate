package amqpext

import (
	"fmt"
	"github.com/h-varmazyar/Gate/pkg/envext"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

type Configs struct {
	Connection string `env:"RABBITMQ_CONNECTION,file,required"`
}

type client struct {
	Connection string
	locker     sync.RWMutex
	conn       *amqp.Connection
	channel    *amqp.Channel
	isSuspend  bool
}

var (
	configs *Configs
	Client  *client
)

func init() {
	configs = new(Configs)
	if err := envext.Load(configs); err != nil {
		log.WithError(err).Fatal("failed to load rabbitMQ configs")
	}
	Client = &client{
		Connection: configs.Connection,
	}
	if err := Client.connect(); err != nil {
		log.WithError(err).Fatal("cannot connect to rabbitMQ")
	}
	go Client.check()
}

func (c *client) QueueDeclare(key string) (*Queue, error) {
	c.locker.Lock()
	defer c.locker.Unlock()
	q := &Queue{
		client:   c,
		key:      key,
		name:     fmt.Sprintf("%v_name", key),
		exchange: fmt.Sprintf("%v_exchange", key),
		done:     make(chan error),
	}
	if err := c.channel.ExchangeDeclare(q.exchange, amqp.ExchangeTopic, true, false, false, false, nil); err != nil {
		return nil, err
	}
	if _, err := c.channel.QueueDeclare(q.name, true, false, false, false, nil); err != nil {
		return nil, err
	}
	if err := c.channel.QueueBind(q.name, q.key, q.exchange, false, nil); err != nil {
		return nil, err
	}
	return q, nil
}

func (c *client) connect() error {
	if c.conn != nil && !c.conn.IsClosed() {
		return nil
	}

	conn, err := amqp.Dial(c.Connection)
	if err != nil {
		return err
	}
	channel, chanErr := conn.Channel()
	if chanErr != nil {
		return chanErr
	}

	c.locker.Lock()
	defer c.locker.Unlock()
	if c.channel != nil {
		_ = c.channel.Close()
	}
	if c.conn != nil {
		_ = c.conn.Close()
	}
	c.conn = conn
	c.channel = channel
	return nil
}

func (c *client) check() {
	for range time.NewTicker(time.Second).C {
		if c.isSuspend {
			break
		}
		if err := c.connect(); err != nil {
			log.WithError(err).Error("cannot reconnect to rabbitMQ")
		}
	}
}

func (c *client) Close() error {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.isSuspend = true
	var errors []error
	if err := c.channel.Close(); err != nil {
		errors = append(errors, err)
	}
	if err := c.conn.Close(); err != nil {
		errors = append(errors, err)
	}
	if len(errors) > 0 {
		return errors[0]
	}
	return nil
}
