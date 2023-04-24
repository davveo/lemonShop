package models

// EsSssGoodsPv2019 [...]
type EsSssGoodsPv2019 struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`
	SellerID  int    `gorm:"column:seller_id" json:"seller_id"`
	GoodsID   int    `gorm:"column:goods_id" json:"goods_id"`
	GoodsName string `gorm:"column:goods_name" json:"goods_name"`
	VsYear    int    `gorm:"column:vs_year" json:"vs_year"`
	VsMonth   int    `gorm:"column:vs_month" json:"vs_month"`
	VsDay     int    `gorm:"column:vs_day" json:"vs_day"`
	VsNum     int    `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssGoodsPv2019) TableName() string {
	return "es_sss_goods_pv_2019"
}