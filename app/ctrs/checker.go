package ctrs

import (
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/reply"
	"github.com/gin-gonic/gin"
)

type CheckerController struct {
}

func NewCheckerController() *CheckerController {
	return &CheckerController{}
}

func (c *CheckerController) Check(ctx *gin.Context) {
	data, err := cache.Cache.Get("test")
	if err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	reply.Reply(ctx, data)
}
