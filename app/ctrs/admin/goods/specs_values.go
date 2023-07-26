package goods

import (
	"github.com/davveo/lemonShop/app/service"
	"github.com/gin-gonic/gin"
)

type SpecsValueController struct {
	SpecsService service.SpecsService
}

func NewSpecsValueController() *SpecsValueController {
	return &SpecsValueController{
		SpecsService: service.NewSpecsService(),
	}
}

func (s *SpecsController) List(ctx *gin.Context) {

}
