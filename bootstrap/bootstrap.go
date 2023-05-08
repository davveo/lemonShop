package bootstrap

import (
	"github.com/davveo/lemonShop/app"
	"github.com/davveo/lemonShop/app/tasks"
	"github.com/davveo/lemonShop/conf"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/db"
	"github.com/davveo/lemonShop/pkg/es"
	"github.com/davveo/lemonShop/pkg/logger"
	"github.com/davveo/lemonShop/pkg/pool"
)

func Bootstrap() {
	// 初始化配置
	appConf, err := conf.Init()
	if err != nil {
		panic(err)
	}

	// 初始化日志
	gLogger, err := logger.Init()
	if err != nil {
		panic(err)
	}

	// 初始化数据
	dbRepo, err := db.Init()
	if err != nil {
		panic(err)
	}

	// 初始化缓存
	cacheRepo, err := cache.Init()
	if err != nil {
		panic(err)
	}

	// 启动mq消费者
	err = tasks.Init()
	if err != nil {
		panic(err)
	}

	// 启动es
	err = es.InitESClient()
	if err != nil {
		panic(err)
	}

	err = pool.InitPool()
	if err != nil {
		panic(err)
	}

	// 服务启动
	server := app.NewServer(appConf, gLogger, dbRepo, cacheRepo)
	server.Init()
}
