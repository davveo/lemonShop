package models

// EsFullDiscount 满优惠活动(es_full_discount)
type EsFullDiscount struct {
	FdID          int     `gorm:"primaryKey;column:fd_id" json:"-"`            // 活动id
	FullMoney     float64 `gorm:"column:full_money" json:"full_money"`         // 优惠门槛金额
	MinusValue    float64 `gorm:"column:minus_value" json:"minus_value"`       // 减现金
	PointValue    int     `gorm:"column:point_value" json:"point_value"`       // 送积分
	IsFullMinus   int     `gorm:"column:is_full_minus" json:"is_full_minus"`   // 活动是否减现金
	IsFreeShip    int     `gorm:"column:is_free_ship" json:"is_free_ship"`     // 是否免邮
	IsSendPoint   int     `gorm:"column:is_send_point" json:"is_send_point"`   // 是否送积分
	IsSendGift    int     `gorm:"column:is_send_gift" json:"is_send_gift"`     // 是否有赠品
	IsSendBonus   int     `gorm:"column:is_send_bonus" json:"is_send_bonus"`   // 是否增优惠券
	GiftID        int     `gorm:"column:gift_id" json:"gift_id"`               // 赠品id
	BonusID       int     `gorm:"column:bonus_id" json:"bonus_id"`             // 优惠券id
	IsDiscount    int     `gorm:"column:is_discount" json:"is_discount"`       // 是否打折
	DiscountValue float64 `gorm:"column:discount_value" json:"discount_value"` // 打多少折
	StartTime     int64   `gorm:"column:start_time" json:"start_time"`         // 活动开始时间
	EndTime       int64   `gorm:"column:end_time" json:"end_time"`             // 活动结束时间
	Title         string  `gorm:"column:title" json:"title"`                   // 活动标题
	RangeType     int     `gorm:"column:range_type" json:"range_type"`         // 是否全部商品参与
	Disabled      int     `gorm:"column:disabled" json:"disabled"`             // 是否停用
	Description   string  `gorm:"column:description" json:"description"`       // 活动说明
	SellerID      int     `gorm:"column:seller_id" json:"seller_id"`           // 商家id
}

// TableName get sql table name.获取数据库表名
func (m *EsFullDiscount) TableName() string {
	return "es_full_discount"
}
