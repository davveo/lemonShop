package entity

import "github.com/davveo/lemonShop/app/consts"

type Seller struct {
	User

	SellerId     int64  `json:"sellerId"`
	SellerName   string `json:"sellerName"`
	SelfOperated int    `json:"selfOperated"`
}

func NewSeller() {
	seller := Seller{}
	seller.Add(consts.Seller)
}
