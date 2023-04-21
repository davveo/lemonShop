package entity

type Page struct {
	PageNo    int64 `form:"page_no" json:"page_no"`
	PageSize  int64 `form:"page_size" json:"page_size"`
	DataTotal int64 `json:"data_Total"`
}
