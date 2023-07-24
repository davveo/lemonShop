package entity

import "github.com/davveo/lemonShop/app/consts"

type Seller struct {
	User

	SellerId     int    `json:"sellerId"`
	SellerName   string `json:"sellerName"`
	SelfOperated int    `json:"selfOperated"`
}

func NewSeller() Seller {
	seller := Seller{}
	seller.Add(consts.Seller)
	return seller
}
