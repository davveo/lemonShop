package models

// EsMemberShopScore 店铺评分(es_member_shop_score)
type EsMemberShopScore struct {
	ScoreID          int    `gorm:"primaryKey;column:score_id" json:"-"`               // 主键
	MemberID         int    `gorm:"column:member_id" json:"member_id"`                 // 会员id
	OrderSn          string `gorm:"column:order_sn" json:"order_sn"`                   // 订单编号
	DeliveryScore    int16  `gorm:"column:delivery_score" json:"delivery_score"`       // 发货速度评分
	DescriptionScore int16  `gorm:"column:description_score" json:"description_score"` // 描述相符度评分
	ServiceScore     int16  `gorm:"column:service_score" json:"service_score"`         // 服务评分
	SellerID         int    `gorm:"column:seller_id" json:"seller_id"`                 // 卖家
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberShopScore) TableName() string {
	return "es_member_shop_score"
}
