package tasks

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"

	"github.com/davveo/lemonShop/pkg/logger"
	"github.com/davveo/lemonShop/pkg/mq"
	"github.com/streadway/amqp"
)

const orderConsumerName = "order_consumer"

type OrderConsumer struct {
	//监听交换机名称
	exchangeName string
	//监听队列名称
	queueName string
	//启动的消费者数量，默认一个
	taskNum int
	//rabbitMQ连接对象
	rabbitMQ *mq.RabbitMQ
}

func (orderConsumer *OrderConsumer) getRabbitMQ() *mq.RabbitMQ {
	return orderConsumer.rabbitMQ
}

func (orderConsumer *OrderConsumer) getName() string {
	return orderConsumerName
}

func (orderConsumer *OrderConsumer) start() (err error) {
	orderConsumer.rabbitMQ = mq.NewRabbitMQ()
	orderConsumer.rabbitMQ.SetConsumerConfig(
		orderConsumer.exchangeName,
		orderConsumer.queueName,
		orderConsumer.taskNum,
		orderConsumer.handler,
	)
	if err = orderConsumer.rabbitMQ.ConsumeStart(); err != nil {
		logger.GLogger.Error(fmt.Sprintf("[%s] 启动失败: %s", orderConsumerName, err.Error()))
		return
	}
	return
}

func (orderConsumer *OrderConsumer) handler(delivery amqp.Delivery) (err error) {
	logger.GLogger.Info(fmt.Sprintf("order_consumer[%s]接收到数据：%s", delivery.ConsumerTag, string(delivery.Body)))
	var mqData map[string]interface{}
	if err = json.Unmarshal(delivery.Body, &mqData); err != nil {
		logger.GLogger.Error("队列json数据解析异常:",
			zap.Any("body", string(delivery.Body)), zap.Any("err", err.Error()))
		return
	}
	return
}
