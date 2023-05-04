package tasks

import (
	"github.com/davveo/lemonShop/pkg/mq"
	"github.com/streadway/amqp"
)

type Consumer interface {
	getName() string
	start() error
	handler(delivery amqp.Delivery) error
	getRabbitMQ() *mq.RabbitMQ
}
