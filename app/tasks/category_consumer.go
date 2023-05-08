package tasks

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"

	"github.com/davveo/lemonShop/pkg/logger"
	"github.com/davveo/lemonShop/pkg/mq"
	"github.com/streadway/amqp"
)

const categoryConsumerName = "category_consumer"

type CategoryConsumer struct {
	//监听交换机名称
	exchangeName string
	//监听队列名称
	queueName string
	//启动的消费者数量，默认一个
	taskNum int
	//rabbitMQ连接对象
	rabbitMQ *mq.RabbitMQ
}

func (categoryConsumer *CategoryConsumer) getRabbitMQ() *mq.RabbitMQ {
	return categoryConsumer.rabbitMQ
}

func (categoryConsumer *CategoryConsumer) getName() string {
	return categoryConsumerName
}

func (categoryConsumer *CategoryConsumer) start() (err error) {
	categoryConsumer.rabbitMQ = mq.NewRabbitMQ()
	categoryConsumer.rabbitMQ.SetConsumerConfig(
		categoryConsumer.exchangeName,
		categoryConsumer.queueName,
		categoryConsumer.taskNum,
		categoryConsumer.handler,
	)
	if err = categoryConsumer.rabbitMQ.ConsumeStart(); err != nil {
		logger.GLogger.Error(fmt.Sprintf("[%s] 启动失败: %s", orderConsumerName, err.Error()))
		return
	}
	return
}

func (categoryConsumer *CategoryConsumer) handler(delivery amqp.Delivery) (err error) {
	logger.GLogger.Info(fmt.Sprintf("categoryConsumer[%s]接收到数据：%s", delivery.ConsumerTag, string(delivery.Body)))
	var mqData map[string]interface{}
	if err = json.Unmarshal(delivery.Body, &mqData); err != nil {
		logger.GLogger.Error("队列json数据解析异常:",
			zap.Any("body", string(delivery.Body)), zap.Any("err", err.Error()))
		return
	}
	//todo 补充消费者逻辑

	return
}
