package models

// EsPromotionGoods 有效活动商品对照表(es_promotion_goods)
type EsPromotionGoods struct {
	PgID          int     `gorm:"primaryKey;column:pg_id" json:"-"`            // 商品对照ID
	GoodsID       int     `gorm:"column:goods_id" json:"goods_id"`             // 商品id
	ProductID     int     `gorm:"column:product_id" json:"product_id"`         // 货品id
	StartTime     int64   `gorm:"column:start_time" json:"start_time"`         // 活动开始时间
	EndTime       int64   `gorm:"column:end_time" json:"end_time"`             // 活动结束时间
	ActivityID    int     `gorm:"column:activity_id" json:"activity_id"`       // 活动id
	PromotionType string  `gorm:"column:promotion_type" json:"promotion_type"` // 促销工具类型
	Title         string  `gorm:"column:title" json:"title"`                   // 活动标题
	Num           int     `gorm:"column:num" json:"num"`                       // 参与活动的商品数量
	Price         float64 `gorm:"column:price" json:"price"`                   // 活动时商品的价格
	SellerID      int     `gorm:"column:seller_id" json:"seller_id"`           // 商家ID
}

// TableName get sql table name.获取数据库表名
func (m *EsPromotionGoods) TableName() string {
	return "es_promotion_goods"
}
