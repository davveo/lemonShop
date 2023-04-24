package models

// EsSssOrderGoodsData 订单商品表(es_sss_order_goods_data)
type EsSssOrderGoodsData struct {
	ID           int     `gorm:"primaryKey;column:id" json:"-"`             // id
	OrderSn      string  `gorm:"column:order_sn" json:"order_sn"`           // 订单编号
	GoodsID      int     `gorm:"column:goods_id" json:"goods_id"`           // 商品id
	GoodsName    string  `gorm:"column:goods_name" json:"goods_name"`       // 商品名称
	GoodsNum     int     `gorm:"column:goods_num" json:"goods_num"`         // 数量
	Price        float64 `gorm:"column:price" json:"price"`                 // 商品单价
	SubTotal     float64 `gorm:"column:sub_total" json:"sub_total"`         // 小计
	CategoryPath string  `gorm:"column:category_path" json:"category_path"` // 分类path
	CategoryID   int     `gorm:"column:category_id" json:"category_id"`     // 分类id
	CreateTime   int     `gorm:"column:create_time" json:"create_time"`     // 创建时间
	IndustryID   int64   `gorm:"column:industry_id" json:"industry_id"`     // 行业id
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderGoodsData) TableName() string {
	return "es_sss_order_goods_data"
}
