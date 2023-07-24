package tasks

import "github.com/davveo/lemonShop/app/consts"

func Init() error {
	taskManager := TaskManager{}
	taskManager.AddTask(&OrderConsumer{
		exchangeName: consts.ORDER_CREATE,
		queueName:    consts.ORDER_CREATE + "_QUEUE",
		taskNum:      1,
	}, &SmsConsumer{
		exchangeName: consts.SEND_MESSAGE,
		queueName:    consts.SEND_MESSAGE + "_QUEUE",
		taskNum:      1,
	})
	return taskManager.start()
}
