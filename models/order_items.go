package models

// EsOrderItems 订单货物表(es_order_items)
type EsOrderItems struct {
	ItemID        int     `gorm:"primaryKey;column:item_id" json:"-"`          // 主键ID
	GoodsID       int     `gorm:"column:goods_id" json:"goods_id"`             // 商品ID
	ProductID     int     `gorm:"column:product_id" json:"product_id"`         // 货品ID
	Num           int     `gorm:"column:num" json:"num"`                       // 销售量
	ShipNum       int     `gorm:"column:ship_num" json:"ship_num"`             // 发货量
	TradeSn       string  `gorm:"column:trade_sn" json:"trade_sn"`             // 交易编号
	OrderSn       string  `gorm:"column:order_sn" json:"order_sn"`             // 订单编号
	Image         string  `gorm:"column:image" json:"image"`                   // 图片
	Name          string  `gorm:"column:name" json:"name"`                     // 商品名称
	Price         float64 `gorm:"column:price" json:"price"`                   // 销售金额
	CatID         int     `gorm:"column:cat_id" json:"cat_id"`                 // 分类ID
	State         int     `gorm:"column:state" json:"state"`                   // 状态
	SnapshotID    int     `gorm:"column:snapshot_id" json:"snapshot_id"`       // 快照id
	SpecJSON      string  `gorm:"column:spec_json" json:"spec_json"`           // 规格json
	PromotionType string  `gorm:"column:promotion_type" json:"promotion_type"` // 促销类型
	PromotionID   int     `gorm:"column:promotion_id" json:"promotion_id"`     // 促销id
	RefundPrice   float64 `gorm:"column:refund_price" json:"refund_price"`     // 订单项可退款金额
	CommentStatus string  `gorm:"column:comment_status" json:"comment_status"` // 评论状态
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderItems) TableName() string {
	return "es_order_items"
}
