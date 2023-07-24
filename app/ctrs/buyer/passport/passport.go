package passport

import "github.com/gin-gonic/gin"

type Controller struct {
}

func NewPassportController() *Controller {
	return &Controller{}
}

func (p *Controller) Todo(ctx *gin.Context) {

}
