package entity

type UserContext struct {
}

// todo Role.SELLER 买家卖家权限限制
// todo 从上下文里面写入
func (userContext *UserContext) GetSeller() (*Seller, error) {
	return &Seller{}, nil
}

func (userContext *UserContext) GetBuyer() (*Buyer, error) {
	return &Buyer{}, nil
}
