package tasks

import (
	"fmt"

	"github.com/davveo/lemonShop/pkg/logger"
)

type TaskManager struct {
	taskPool []Consumer
}

func (taskManager *TaskManager) AddTask(consumer Consumer) {
	isExist := false
	for _, task := range taskManager.taskPool {
		if consumer.getName() == task.getName() {
			isExist = true
			break
		}
	}
	if !isExist {
		logger.GLogger.Info(fmt.Sprintf("mq添加消费者: %s", consumer.getName()))
		taskManager.taskPool = append(taskManager.taskPool, consumer)
	}
}

func (taskManager *TaskManager) start() error {
	for _, task := range taskManager.taskPool {

		if err := task.start(); err != nil {
			logger.GLogger.Error(err.Error())
			return err
		}
	}
	return nil
}

// 同一进程下操作
func (taskManager *TaskManager) stop() {
	for _, task := range taskManager.taskPool {
		//停止所有消费者
		task.getRabbitMQ().ConsumeStop()
	}
}

// 同一进程下操作
func (taskManager *TaskManager) StopAll() {
	for _, task := range taskManager.taskPool {
		//停止所有消费者
		task.getRabbitMQ().ConsumeStop()
	}
}

// 停止任务
func (taskManager *TaskManager) StopConsumer(taskName string) {
	for _, task := range taskManager.taskPool {
		if taskName == task.getName() {
			//停止消费者
			//delete() 删除map
			//taskManager.taskPool = taskManager.taskPool[1:4]
			break
		}
	}
}
