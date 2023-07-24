package member

import "github.com/gin-gonic/gin"

type Controller struct {
}

func NewMemberController() *Controller {
	return &Controller{}
}

func (m *Controller) Todo(ctx *gin.Context) {

}
