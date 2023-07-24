package trade

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewTradeController() *Controller {
	return &Controller{}
}

func (t *Controller) Todo(ctx *gin.Context) {

}
