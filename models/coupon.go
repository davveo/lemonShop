package models

// EsCoupon 优惠券(es_coupon)
type EsCoupon struct {
	Title                string  `gorm:"column:title" json:"title"`                                   // 优惠券标题
	CouponID             int     `gorm:"primaryKey;column:coupon_id" json:"-"`                        // 主键
	CouponPrice          float64 `gorm:"column:coupon_price" json:"coupon_price"`                     // 优惠券面额
	CouponThresholdPrice float64 `gorm:"column:coupon_threshold_price" json:"coupon_threshold_price"` // 优惠券门槛价格
	StartTime            int64   `gorm:"column:start_time" json:"start_time"`                         // 使用起始时间
	EndTime              int64   `gorm:"column:end_time" json:"end_time"`                             // 使用截止时间
	CreateNum            int     `gorm:"column:create_num" json:"create_num"`                         // 发行量
	LimitNum             int     `gorm:"column:limit_num" json:"limit_num"`                           // 每人限领数量
	UsedNum              int     `gorm:"column:used_num" json:"used_num"`                             // 已被使用的数量
	SellerID             int     `gorm:"column:seller_id" json:"seller_id"`                           // 店铺ID
	ReceivedNum          int     `gorm:"column:received_num" json:"received_num"`                     // 已被领取的数量
	SellerName           string  `gorm:"column:seller_name" json:"seller_name"`                       // 店铺名称
	Type                 string  `gorm:"column:type" json:"type"`                                     // 优惠券类型，分为免费领取和活动赠送
	UseScope             string  `gorm:"column:use_scope" json:"use_scope"`                           // 使用范围，全品，分类，部分商品
	ScopeID              string  `gorm:"column:scope_id" json:"scope_id"`                             // 全品或者商家优惠券时为0<br />分类时为分类id<br />部分商品时为商品ID集合
	ShopCommission       int     `gorm:"column:shop_commission" json:"shop_commission"`               // 店铺承担比例，正整数
	ScopeDescription     string  `gorm:"column:scope_description" json:"scope_description"`           // 范围描述
	ActivityDescription  string  `gorm:"column:activity_description" json:"activity_description"`     // 活动说明
}

// TableName get sql table name.获取数据库表名
func (m *EsCoupon) TableName() string {
	return "es_coupon"
}
