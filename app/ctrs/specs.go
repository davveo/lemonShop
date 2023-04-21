package ctrs

import (
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/app/service"
	"github.com/davveo/lemonShop/pkg/reply"
	"github.com/gin-gonic/gin"
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

//func (s *SpecsController) CreateSpecs(ctx *gin.Context) {
//	var request entity.SpecsCreateRequest
//	if err := ctx.ShouldBind(&request); err != nil {
//		reply.ReplyInternalErr(ctx, err.Error())
//		return
//	}
//
//	// extra params
//	mapData["seller_id"] = "0"
//	mapData["disabled"] = 1
//
//	spec, err := model.CreateSpecFactory("").Add(mapData)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"code":    http.StatusInternalServerError,
//			"message": err.Error(),
//		})
//		return
//	}
//	ctx.JSON(http.StatusOK, spec)
//}

//func UpdateSpecs(ctx *gin.Context) {
//	var (
//		err      error
//		specMemo = ctx.DefaultPostForm("spec_memo", "")
//		specName = ctx.DefaultPostForm("spec_name", "")
//		disabled = ctx.DefaultPostForm("disabled", "")
//		specId   = ctx.DefaultPostForm("spec_id", "")
//		sellerId = ctx.DefaultPostForm("seller_id", "")
//	)
//
//	specRequest := request.SpecsRequest{
//		SpecMemo: specMemo, SpecName: specName,
//	}
//	mapData := transfer.StructToMap(specRequest)
//
//	// extra params
//	mapData["spec_id"] = specId
//	mapData["disabled"] = disabled
//	mapData["seller_id"] = sellerId
//
//	spec, err := model.CreateSpecFactory("").Edit(mapData)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"code":    http.StatusInternalServerError,
//			"message": err.Error(),
//		})
//		return
//	}
//	ctx.JSON(http.StatusOK, spec)
//}
//
//func DeleteSpecs(ctx *gin.Context) {
//
//	specIds := ctx.Param("spec_id")
//	split := strings.Split(specIds, ",")
//	err := model.CreateSpecFactory("").Delete(transfer.StringToInt(split))
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"code":    http.StatusInternalServerError,
//			"message": err.Error(),
//		})
//		return
//	}
//	ctx.JSON(http.StatusOK, nil)
//}
