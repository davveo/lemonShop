package models

// EsBillMember 个人业绩单(es_bill_member)
type EsBillMember struct {
	ID               int     `gorm:"primaryKey;column:id" json:"-"`                       // id
	TotalID          int     `gorm:"column:total_id" json:"total_id"`                     // 总结算单id
	MemberID         int     `gorm:"column:member_id" json:"member_id"`                   // 会员id
	StartTime        int     `gorm:"column:start_time" json:"start_time"`                 // 开始时间
	EndTime          int     `gorm:"column:end_time" json:"end_time"`                     // 结束时间
	FinalMoney       float64 `gorm:"column:final_money" json:"final_money"`               // 最终结算金额
	PushMoney        float64 `gorm:"column:push_money" json:"push_money"`                 // 提成金额
	OrderMoney       float64 `gorm:"column:order_money" json:"order_money"`               // 订单金额
	ReturnOrderMoney float64 `gorm:"column:return_order_money" json:"return_order_money"` // 订单返还金额
	ReturnPushMoney  float64 `gorm:"column:return_push_money" json:"return_push_money"`   // 返还提成金额
	MemberName       string  `gorm:"column:member_name" json:"member_name"`               // 会员名称
	Sn               string  `gorm:"column:sn" json:"sn"`                                 // 编号
	OrderCount       int     `gorm:"column:order_count" json:"order_count"`               // 订单数
	ReturnOrderCount int     `gorm:"column:return_order_count" json:"return_order_count"` // 返还订单数
}

// TableName get sql table name.获取数据库表名
func (m *EsBillMember) TableName() string {
	return "es_bill_member"
}
