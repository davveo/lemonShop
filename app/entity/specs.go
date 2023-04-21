package entity

type SpecsListRequest struct {
	KeyWord string `form:"keyword"`
	Page
}

type SpecsRequest struct {
	SpecId int64 `uri:"spec_id"`
}

type SpecsCreateRequest struct {
	SpecName string `form:"spec_name"`
	SpecMemo string `form:"spec_memo"`
}

type SpecsListResp struct {
	Data      interface{} `json:"data"`
	PageNo    int64       `form:"page_no" json:"page_no"`
	PageSize  int64       `form:"page_size" json:"page_size"`
	DataTotal int64       `json:"data_Total"`
}

type SpecResp struct {
	SpecID   int    `json:"spec_id"`   // 主键
	SpecName string `json:"spec_name"` // 规格项名称
	Disabled int    `json:"disabled"`  // 是否被删除0 删除   1  没有删除
	SpecMemo string `json:"spec_memo"` // 规格描述
	SellerID int    `json:"seller_id"` // 所属卖家
}

type SpecsCreateResp struct {
}
