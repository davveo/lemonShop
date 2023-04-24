package models

// EsShopLogiRel 店铺物流公司对照表(es_shop_logi_rel)
type EsShopLogiRel struct {
	LogiID int `gorm:"column:logi_id" json:"logi_id"` // 物流公司ID
	ShopID int `gorm:"column:shop_id" json:"shop_id"` // 店铺ID
}

// TableName get sql table name.获取数据库表名
func (m *EsShopLogiRel) TableName() string {
	return "es_shop_logi_rel"
}