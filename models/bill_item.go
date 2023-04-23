package models

// EsBillItem 结算单项表(es_bill_item)
type EsBillItem struct {
	ID               int     `gorm:"primaryKey;column:id" json:"-"`                     // 主键id
	OrderSn          string  `gorm:"column:order_sn" json:"order_sn"`                   // 订单编号
	Price            float64 `gorm:"column:price" json:"price"`                         // 订单价格
	DiscountPrice    float64 `gorm:"column:discount_price" json:"discount_price"`       // 优惠价格
	ItemType         string  `gorm:"column:item_type" json:"item_type"`                 // 单项类型 收款/退款
	AddTime          int     `gorm:"column:add_time" json:"add_time"`                   // 加入时间
	BillID           int     `gorm:"column:bill_id" json:"bill_id"`                     // 所属账单id
	MemberID         int     `gorm:"column:member_id" json:"member_id"`                 // 会员id
	Status           int     `gorm:"column:status" json:"status"`                       // 状态
	SellerID         int     `gorm:"column:seller_id" json:"seller_id"`                 // 店铺id
	OrderTime        int64   `gorm:"column:order_time" json:"order_time"`               // 下单时间
	RefundSn         string  `gorm:"column:refund_sn" json:"refund_sn"`                 // 退款单号
	MemberName       string  `gorm:"column:member_name" json:"member_name"`             // 会员名称
	ShipName         string  `gorm:"column:ship_name" json:"ship_name"`                 // 收货人
	PaymentType      string  `gorm:"column:payment_type" json:"payment_type"`           // 支付方式
	RefundTime       int64   `gorm:"column:refund_time" json:"refund_time"`             // 退货时间
	SiteCouponPrice  float64 `gorm:"column:site_coupon_price" json:"site_coupon_price"` // 使用平台优惠券金额
	CouponCommission int     `gorm:"column:coupon_commission" json:"coupon_commission"` // 优惠券佣金比例
}

// TableName get sql table name.获取数据库表名
func (m *EsBillItem) TableName() string {
	return "es_bill_item"
}
