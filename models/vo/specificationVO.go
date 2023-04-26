package vo

import "github.com/davveo/lemonShop/models"

type SpecificationVO struct {
	models.EsSpecification
	ValueList []*models.EsSpecValues `json:"valueList"`
}
