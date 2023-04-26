package ctrs

import (
	"github.com/davveo/lemonShop/app/service"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/reply"
	"github.com/gin-gonic/gin"
)

type CheckerController struct {
	SpecsService service.SpecsService
}

func NewCheckerController() *CheckerController {
	return &CheckerController{
		SpecsService: service.NewSpecsService(),
	}
}

func (c *CheckerController) Check(ctx *gin.Context) {
	data, err := cache.Cache.Get("test")
	if err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	reply.Reply(ctx, data)
}

func (c *CheckerController) Test(ctx *gin.Context) {
	spec, err := c.SpecsService.QuerySellerSpec(ctx, 311)
	if err != nil {
		return
	}
	reply.Reply(ctx, spec)
}
