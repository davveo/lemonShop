package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/davveo/lemonShop/pkg/errors"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, errors.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, errors.ERROR
	}
	if !check {
		return http.StatusBadRequest, errors.INVALID_PARAMS
	}

	return http.StatusOK, errors.SUCCESS
}
