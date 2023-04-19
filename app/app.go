package app

import (
	"fmt"
	"github.com/davveo/lemonShop/app/router"
	"github.com/davveo/lemonShop/config"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"syscall"
)

type Server struct {
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Init() {
	gin.SetMode(config.Conf.Server.RunMode)

	route := gin.New()
	router.Init(route)

	endless.DefaultMaxHeaderBytes = 1 << 20
	endless.DefaultReadTimeOut = config.Conf.Server.ReadTimeout
	endless.DefaultWriteTimeOut = config.Conf.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.Conf.Server.HttpPort)

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
