package entity

import "github.com/davveo/lemonShop/app/consts"

type Buyer struct {
	User
}

func NewBuyer() {
	buyer := Buyer{}
	buyer.Add(consts.Buyer)
}
