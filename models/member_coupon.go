package models

// EsMemberCoupon [...]
type EsMemberCoupon struct {
	McID       int    `gorm:"primaryKey;column:mc_id" json:"-"`
	CouponID   int    `gorm:"column:coupon_id" json:"coupon_id"`
	MemberID   int    `gorm:"column:member_id" json:"member_id"`
	UsedTime   int64  `gorm:"column:used_time" json:"used_time"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	OrderID    int    `gorm:"column:order_id" json:"order_id"`
	OrderSn    string `gorm:"column:order_sn" json:"order_sn"`
	MemberName string `gorm:"column:member_name" json:"member_name"`
	CouponSn   string `gorm:"column:coupon_sn" json:"coupon_sn"`
	Title      string `gorm:"column:title" json:"title"`
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberCoupon) TableName() string {
	return "es_member_coupon"
}
