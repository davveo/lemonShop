package app

import (
	"fmt"
	"github.com/davveo/lemonShop/app/router"
	"github.com/davveo/lemonShop/conf"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/db"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"syscall"
)

type Server struct {
	db      db.Repo
	cache   cache.Repo
	lg      *zap.Logger
	appConf *conf.AppConf
}

func NewServer(appConf *conf.AppConf, lg *zap.Logger, dbRepo db.Repo, cacheRepo cache.Repo) Server {
	return Server{
		db:      dbRepo,
		lg:      lg,
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
		s.lg.Debug(fmt.Sprintf("[info] pid is %d, start "+
			"http server listening %s", syscall.Getpid(), endPoint))
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}

func (s *Server) Clean() {
	if s.db != nil {
		if err := s.db.DbWClose(); err != nil {
			s.lg.Info("dbw close err, ", zap.Error(err))
			return
		}
		if err := s.db.DbRClose(); err != nil {
			s.lg.Info("dbr close err, ", zap.Error(err))
			return
		}
	}

	if s.cache != nil {
		if err := s.cache.Close(); err != nil {
			s.lg.Info("cache close err, ", zap.Error(err))
			return
		}
	}

	if err := s.lg.Sync(); err != nil {
		s.lg.Info("log sync err, ", zap.Error(err))
		return
	}
}
