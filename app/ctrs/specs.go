package ctrs

import (
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/app/service"
	"github.com/davveo/lemonShop/pkg/common"
	"github.com/davveo/lemonShop/pkg/reply"
	"github.com/gin-gonic/gin"
	"strings"
)

type SpecsController struct {
	SpecsService service.SpecsService
}

func NewSpecsController() *SpecsController {
	return &SpecsController{
		SpecsService: service.NewSpecsService(),
	}
}

func (s *SpecsController) SpecsList(ctx *gin.Context) {
	var req entity.SpecsListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	data, err := s.SpecsService.List(ctx, &req)
	if err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	reply.Reply(ctx, data)
}

func (s *SpecsController) Specs(ctx *gin.Context) {
	var req entity.SpecsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	obj, err := s.SpecsService.Get(ctx, &req)
	if err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	reply.Reply(ctx, obj)
}

func (s *SpecsController) Create(ctx *gin.Context) {
	var req entity.SpecsCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}

	spec, err := s.SpecsService.Create(ctx, &req)
	if err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	reply.Reply(ctx, spec)
}

func (s *SpecsController) Update(ctx *gin.Context) {
	var req entity.SpecsUpdateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}

	spec, err := s.SpecsService.Update(ctx, &req)
	if err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	reply.Reply(ctx, spec)
}

func (s *SpecsController) Delete(ctx *gin.Context) {
	var req entity.SpecsDeleteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	strList := strings.Split(req.SpecId, ",")
	specsList := common.StringToInt64(strList)
	if err := s.SpecsService.Delete(ctx, specsList); err != nil {
		reply.ReplyInternalErr(ctx, err.Error())
		return
	}
	reply.Reply(ctx, nil)
}
