package entity

import "context"

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

var userKey string

func NewContext(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userKey).(*User)
	return u, ok
}
