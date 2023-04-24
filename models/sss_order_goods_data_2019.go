package models

// EsSssOrderGoodsData2019 [...]
type EsSssOrderGoodsData2019 struct {
	ID           int     `gorm:"primaryKey;column:id" json:"-"`
	OrderSn      string  `gorm:"column:order_sn" json:"order_sn"`
	GoodsID      int     `gorm:"column:goods_id" json:"goods_id"`
	GoodsName    string  `gorm:"column:goods_name" json:"goods_name"`
	GoodsNum     int     `gorm:"column:goods_num" json:"goods_num"`
	Price        float64 `gorm:"column:price" json:"price"`
	SubTotal     float64 `gorm:"column:sub_total" json:"sub_total"`
	CategoryPath string  `gorm:"column:category_path" json:"category_path"`
	CategoryID   int     `gorm:"column:category_id" json:"category_id"`
	CreateTime   int     `gorm:"column:create_time" json:"create_time"`
	IndustryID   int     `gorm:"column:industry_id" json:"industry_id"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderGoodsData2019) TableName() string {
	return "es_sss_order_goods_data_2019"
}