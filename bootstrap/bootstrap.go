package bootstrap

import (
	"github.com/davveo/lemonShop/app"
	"github.com/davveo/lemonShop/conf"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/db"
	"github.com/davveo/lemonShop/pkg/logger"
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

	// 服务启动
	server := app.NewServer(appConf, gLogger, dbRepo, cacheRepo)
	server.Init()
}
