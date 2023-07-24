package shop

import "github.com/davveo/lemonShop/models"

var _ ShopRoleService = (*shopRoleService)(nil)

type (
	ShopRoleService interface {
		i()
		GetByRoleId(roleId int) (*models.EsShopRole, error)
	}

	shopRoleService struct {
	}
)

func (s shopRoleService) GetByRoleId(roleId int) (*models.EsShopRole, error) {
	//TODO implement me
	panic("implement me")
}

func (s shopRoleService) i() {}

func NewShopRoleService() ShopRoleService {
	return &shopRoleService{}
}
