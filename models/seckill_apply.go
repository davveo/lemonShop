package models

// EsSeckillApply 限时抢购申请(es_seckill_apply)
type EsSeckillApply struct {
	ApplyID       int     `gorm:"primaryKey;column:apply_id" json:"-"`         // 主键ID
	SeckillID     int     `gorm:"column:seckill_id" json:"seckill_id"`         // 活动id
	TimeLine      int     `gorm:"column:time_line" json:"time_line"`           // 时刻
	StartDay      int64   `gorm:"column:start_day" json:"start_day"`           // 活动开始日期
	GoodsID       int     `gorm:"column:goods_id" json:"goods_id"`             // 商品ID
	GoodsName     string  `gorm:"column:goods_name" json:"goods_name"`         // 商品名称
	SkuID         int     `gorm:"column:sku_id" json:"sku_id"`                 // 商品规格ID
	SellerID      int     `gorm:"column:seller_id" json:"seller_id"`           // 商家ID
	ShopName      string  `gorm:"column:shop_name" json:"shop_name"`           // 商家名称
	Price         float64 `gorm:"column:price" json:"price"`                   // 价格
	SoldQuantity  int     `gorm:"column:sold_quantity" json:"sold_quantity"`   // 售空数量
	Status        string  `gorm:"column:status" json:"status"`                 // 申请状态
	FailReason    string  `gorm:"column:fail_reason" json:"fail_reason"`       // 驳回原因
	SalesNum      int     `gorm:"column:sales_num" json:"sales_num"`           // 已售数量
	OriginalPrice float64 `gorm:"column:original_price" json:"original_price"` // 商品原始价格
}


// TableName get sql table name.获取数据库表名
func (m *EsSeckillApply) TableName() string {
	return "es_seckill_apply"
}