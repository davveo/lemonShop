package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/davveo/lemonShop/pkg/log"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		log.Info(err.Key, err.Message)
	}

	return
}
