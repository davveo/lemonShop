package entity

type SpecValues struct {
	SpecValueID int    `json:"spec_value_id"` // 主键
	SpecID      int    `json:"spec_id"`       // 规格项id
	SpecValue   string `json:"spec_value"`    // 规格值名字
	SellerID    int    `json:"seller_id"`     // 所属卖家
	SpecName    string `json:"spec_name"`     // 规格名字
}
