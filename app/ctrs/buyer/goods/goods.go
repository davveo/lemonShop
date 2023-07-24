package goods

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewGoodsController() *Controller {
	return &Controller{}
}

func (g *Controller) Todo(ctx *gin.Context) {

}
