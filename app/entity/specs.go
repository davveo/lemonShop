package entity

type Specs struct {
	SpecID   int    `form:"spec_id" uri:"spec_id" json:"spec_id"` // 主键
	SpecName string `form:"spec_name" json:"spec_name"`           // 规格项名称
	Disabled int    `form:"disabled" json:"disabled"`             // 是否被删除0 删除   1  没有删除
	SpecMemo string `form:"spec_memo" json:"spec_memo"`           // 规格描述
	SellerID int    `form:"seller_id" json:"seller_id"`           // 所属卖家
}

type SpecsListRequest struct {
	KeyWord string `form:"keyword"`
	Page
}

type SpecsDeleteRequest struct {
	SpecId string `uri:"spec_id"`
}
