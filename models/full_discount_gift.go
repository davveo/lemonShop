package models

// EsFullDiscountGift 满优惠赠品(es_full_discount_gift)
type EsFullDiscountGift struct {
	GiftID      int     `gorm:"primaryKey;column:gift_id" json:"-"`      // 赠品id
	GiftName    string  `gorm:"column:gift_name" json:"gift_name"`       // 赠品名称
	GiftPrice   float64 `gorm:"column:gift_price" json:"gift_price"`     // 赠品金额
	GiftImg     string  `gorm:"column:gift_img" json:"gift_img"`         // 赠品图片
	GiftType    int     `gorm:"column:gift_type" json:"gift_type"`       // 赠品类型
	ActualStore int     `gorm:"column:actual_store" json:"actual_store"` // 库存
	EnableStore int     `gorm:"column:enable_store" json:"enable_store"` // 可用库存
	CreateTime  int64   `gorm:"column:create_time" json:"create_time"`   // 活动时间
	GoodsID     int     `gorm:"column:goods_id" json:"goods_id"`         // 活动商品id
	Disabled    int     `gorm:"column:disabled" json:"disabled"`         // 是否禁用
	SellerID    int     `gorm:"column:seller_id" json:"seller_id"`       // 店铺id
}

// TableName get sql table name.获取数据库表名
func (m *EsFullDiscountGift) TableName() string {
	return "es_full_discount_gift"
}
