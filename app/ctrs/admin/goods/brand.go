package goods

import (
	"github.com/gin-gonic/gin"
)

type BrandController struct {
}

func NewBrandController() *BrandController {
	return &BrandController{}
}

func (b *BrandController) List(ctx *gin.Context) {

}
