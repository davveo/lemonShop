package buyer

import "github.com/gin-gonic/gin"

type MemberController struct {
}

func NewMemberController() *MemberController {
	return &MemberController{}
}

func (m *MemberController) Todo(ctx *gin.Context) {

}
