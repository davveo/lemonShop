package member

import "github.com/davveo/lemonShop/models/vo"

var _ ShopClient = (*shopClient)(nil)

type ShopClient interface {
	i()
	GetShop(shopId int) (*vo.ShopVo, error)
}

type shopClient struct {
}

func (s shopClient) GetShop(shopId int) (*vo.ShopVo, error) {
	//TODO implement me
	panic("implement me")
}

func (s shopClient) i() {}

func NewShopClient() ShopClient {
	return &shopClient{}
}
