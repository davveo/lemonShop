package models

// EsSssShopPv2018 [...]
type EsSssShopPv2018 struct {
	ID       int `gorm:"column:id" json:"id"`
	SellerID int `gorm:"column:seller_id" json:"seller_id"`
	VsYear   int `gorm:"column:vs_year" json:"vs_year"`
	VsMonth  int `gorm:"column:vs_month" json:"vs_month"`
	VsDay    int `gorm:"column:vs_day" json:"vs_day"`
	VsNum    int `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssShopPv2018) TableName() string {
	return "es_sss_shop_pv_2018"
}