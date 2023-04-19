package bootstrap

import (
	"github.com/davveo/lemonShop/app"
	"github.com/davveo/lemonShop/config"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/db"
	"github.com/davveo/lemonShop/pkg/log"
)

func Bootstrap() {
	// 初始化配置
	if err := config.Init(); err != nil {
		panic(err)
	}

	// 初始化日志
	if err := log.Init(); err != nil {
		panic(err)
	}

	// 初始化数据
	if err := db.Init(); err != nil {
		panic(err)
	}

	// 初始化缓存
	if err := cache.Init(); err != nil {
		panic(err)
	}

	// 服务启动
	server := app.NewServer()
	server.Init()
}
