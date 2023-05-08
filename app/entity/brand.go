package entity

type Brand struct {
	BrandID int    `json:"brandId"` // 主键
	Name    string `json:"name"`    // 品牌名称
	Logo    string `json:"logo"`    // 品牌图标
}
