package tasks

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"

	"github.com/davveo/lemonShop/pkg/logger"
	"github.com/davveo/lemonShop/pkg/mq"
	"github.com/streadway/amqp"
)

const smsConsumerName = "sms_consumer"

type SmsConsumer struct {
	//监听交换机名称
	exchangeName string
	//监听队列名称
	queueName string
	//启动的消费者数量，默认一个
	taskNum int
	//rabbitMQ连接对象
	rabbitMQ *mq.RabbitMQ
}

func (smsConsumer *SmsConsumer) getRabbitMQ() *mq.RabbitMQ {
	return smsConsumer.rabbitMQ
}

func (smsConsumer *SmsConsumer) getName() string {
	return smsConsumerName
}

func (smsConsumer *SmsConsumer) start() (err error) {
	smsConsumer.rabbitMQ = mq.NewRabbitMQ()
	smsConsumer.rabbitMQ.SetConsumerConfig(
		smsConsumer.exchangeName,
		smsConsumer.queueName,
		smsConsumer.taskNum,
		smsConsumer.handler,
	)
	if err = smsConsumer.rabbitMQ.ConsumeStart(); err != nil {
		logger.GLogger.Error(fmt.Sprintf("[%s] 启动失败: %s", smsConsumerName, err.Error()))
		return
	}
	return
}

func (smsConsumer *SmsConsumer) handler(delivery amqp.Delivery) (err error) {
	logger.GLogger.Info(fmt.Sprintf("sms_consumer[%s]接收到数据：%s", delivery.ConsumerTag, string(delivery.Body)))
	var mqData map[string]interface{}
	if err = json.Unmarshal(delivery.Body, &mqData); err != nil {
		logger.GLogger.Error("队列json数据解析异常:",
			zap.Any("body", string(delivery.Body)), zap.Any("err", err.Error()))
		return
	}

	return
}
