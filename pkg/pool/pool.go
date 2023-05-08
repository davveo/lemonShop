package pool

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/davveo/lemonShop/conf"
	"github.com/davveo/lemonShop/pkg/logger"
)

func InitPool() error {
	if !conf.Conf.IsUseRoutinePool {
		return nil
	}

	logger.GLogger.Info("init go routine pool start")
	panicHandler := func(p interface{}) {
		logger.GLogger.Error(fmt.Sprintf("goroutine panic: %+v", p))
		// todo 增加通知或者告警
	}

	err := InitCommonPool(&Conf{
		PreAlloc:         true,
		Nonblocking:      false,
		DisablePurge:     false,
		PanicHandler:     panicHandler,
		Size:             conf.Conf.CommonGoRoutinePoolSize,
		MaxBlockingTasks: conf.Conf.CommonGoRoutinePoolMaxBlockingTasks,
		ExpiryDuration:   time.Duration(conf.Conf.CommonGoRoutinePoolMinuteExpire) * time.Minute,
	})
	if err != nil {
		logger.GLogger.Error("fail to init common pool", zap.Any("err", err))
		return err
	}

	//worker := handle.NewWorkerHandle(app.storeFactory)
	//err = InitWorkerPool(&Conf{
	//	Size:             conf.WorkerGoRoutinePoolSize,
	//	ExpiryDuration:   time.Duration(conf.WorkerGoRoutinePoolMinuteExpire) * time.Minute,
	//	PreAlloc:         true,
	//	Nonblocking:      false,
	//	MaxBlockingTasks: conf.WorkerGoRoutinePoolMaxBlockingTasks,
	//	DisablePurge:     false,
	//	PanicHandler:     panicHandler,
	//	PoolFunction:     worker.Consume,
	//})
	//if err != nil {
	//	logger.FatalErr("fail to init worker pool", err)
	//	os.Exit(1)
	//}

	return nil
}

//func test() error {
//	fmt.Println("hello world")
//	return nil
//}
//
//func mainCommonPool()  {
//	err := GetCommonPool().Submit(func() {
//		err := test()
//		if err != nil {
//			return
//		}
//	})
//	if err != nil {
//		return err
//	}
//}

//func worker(workerId int) error {
//	return nil
//}
//
//func mainWorkerPool() {
//	GetWorkerPool().Invoke(1)
//}
