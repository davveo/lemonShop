package ctrs

import (
	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/app/event_handler"
	"github.com/davveo/lemonShop/app/service"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/event"
	"github.com/davveo/lemonShop/pkg/reply"
	"github.com/gin-gonic/gin"
)

type CheckerController struct {
	specsService    service.SpecsService
	categoryService service.CategoryService
	eventManager    event.EventManager
}

func NewCheckerController() *CheckerController {
	return &CheckerController{
		specsService:    service.NewSpecsService(),
		categoryService: service.NewCategoryService(),
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
	// c.testMq(ctx)
	c.testCategory(ctx)
}

func (c *CheckerController) testMq(ctx *gin.Context) {
	key := ctx.DefaultQuery("key", "111")

	c.eventManager.Bind(consts.ORDER_CREATE, &event_handler.OrderEventHandler{}).
		Trigger(consts.ORDER_CREATE, key)

	reply.Reply(ctx, nil)
}

func (c *CheckerController) testCategory(ctx *gin.Context) {
	err := c.categoryService.Update(ctx, &entity.Category{CategoryID: 1})
	if err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	reply.Reply(ctx, nil)
}
