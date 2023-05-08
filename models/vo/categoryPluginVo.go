package vo

import "github.com/davveo/lemonShop/models"

type CategoryPluginVo struct {
	models.EsCategory
	Id   int    `json:"id"`
	Text string `json:"text"`
}
