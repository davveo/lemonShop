package vo

import "github.com/davveo/lemonShop/models"

type CategoryVo struct {
	models.EsCategory
	Children  []*CategoryVo `json:"children"`
	BrandList []*BrandVo    `json:"brandList"`
}
