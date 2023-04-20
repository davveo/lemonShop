package app

import (
	"fmt"
	"github.com/davveo/lemonShop/app/router"
	"github.com/davveo/lemonShop/conf"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/db"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"syscall"
)

type Server struct {
	appConf *conf.AppConf
	db      db.Repo
	cache   cache.Repo
}

func NewServer(appConf *conf.AppConf,
	dbRepo db.Repo, cacheRepo cache.Repo) Server {
	return Server{
		db:      dbRepo,
		appConf: appConf,
		cache:   cacheRepo,
	}
}

func (s *Server) Init() {
	defer s.Clean() // 资源清理

	gin.SetMode(conf.Conf.Server.RunMode)

	route := gin.New()
	router.Init(route)

	endless.DefaultMaxHeaderBytes = 1 << 20
	endless.DefaultReadTimeOut = conf.Conf.Server.ReadTimeout
	endless.DefaultWriteTimeOut = conf.Conf.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", conf.Conf.Server.HttpPort)

	server := endless.NewServer(endPoint, route)
	server.BeforeBegin = func(add string) {
		log.Printf("[info] pid is %d, start "+
			"http server listening %s", syscall.Getpid(), endPoint)
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

func (s *Server) Clean() {
	if s.db != nil {
		if err := s.db.DbWClose(); err != nil {
			log.Printf("dbw close err, %+v", err.Error())
		}
		if err := s.db.DbRClose(); err != nil {
			log.Printf("dbw close err, %+v", err.Error())
		}
	}

	if s.cache != nil {
		if err := s.cache.Close(); err != nil {
			log.Printf("ca close err, %+v", err.Error())
		}
	}
}
