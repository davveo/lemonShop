package models

// EsExchange 积分兑换(es_exchange)
type EsExchange struct {
	ExchangeID     int     `gorm:"primaryKey;column:exchange_id" json:"-"`        // 主键
	GoodsID        int     `gorm:"column:goods_id" json:"goods_id"`               // 商品id
	CategoryID     int     `gorm:"column:category_id" json:"category_id"`         // 商品所属积分分类
	EnableExchange int     `gorm:"column:enable_exchange" json:"enable_exchange"` // 是否允许兑换
	ExchangeMoney  float64 `gorm:"column:exchange_money" json:"exchange_money"`   // 兑换所需金额
	ExchangePoint  int     `gorm:"column:exchange_point" json:"exchange_point"`   // 兑换所需积分
	GoodsName      string  `gorm:"column:goods_name" json:"goods_name"`           // 商品名称
	GoodsPrice     float64 `gorm:"column:goods_price" json:"goods_price"`         // 商品原价
	GoodsImg       string  `gorm:"column:goods_img" json:"goods_img"`             // 商品图片
}

// TableName get sql table name.获取数据库表名
func (m *EsExchange) TableName() string {
	return "es_exchange"
}
