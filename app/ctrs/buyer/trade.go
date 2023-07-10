package buyer

import (
	"github.com/gin-gonic/gin"
)

type TradeController struct {
}

func NewTradeController() *TradeController {
	return &TradeController{}
}

func (t *TradeController) Todo(ctx *gin.Context) {

}
