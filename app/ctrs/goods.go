package ctrs

import (
	"github.com/gin-gonic/gin"
)

type GoodsController struct {
}

func NewGoodsController() *GoodsController {
	return &GoodsController{}
}

func (g *GoodsController) Todo(ctx *gin.Context) {

}
