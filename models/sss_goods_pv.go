package models

// EsSssGoodsPv 商品访问量统计表(es_sss_goods_pv)
type EsSssGoodsPv struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`       // 主键id
	SellerID  int    `gorm:"column:seller_id" json:"seller_id"`   // 店铺id
	GoodsID   int    `gorm:"column:goods_id" json:"goods_id"`     // 商品id
	GoodsName string `gorm:"column:goods_name" json:"goods_name"` // 商品名称
	VsYear    int    `gorm:"column:vs_year" json:"vs_year"`       // 年份
	VsNum     int    `gorm:"column:vs_num" json:"vs_num"`         // 访问量
	VsMonth   int    `gorm:"column:vs_month" json:"vs_month"`     // 月份
	VsDay     int    `gorm:"column:vs_day" json:"vs_day"`         // 天
}

// TableName get sql table name.获取数据库表名
func (m *EsSssGoodsPv) TableName() string {
	return "es_sss_goods_pv"
}