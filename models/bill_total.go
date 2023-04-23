package models

// EsBillTotal 总业绩单(es_bill_total)
type EsBillTotal struct {
	ID               int64   `gorm:"primaryKey;column:id" json:"-"`                       // id
	StartTime        int64   `gorm:"column:start_time" json:"start_time"`                 // 开始时间
	EndTime          int64   `gorm:"column:end_time" json:"end_time"`                     // 结束时间
	OrderCount       int64   `gorm:"column:order_count" json:"order_count"`               // 订单数量
	ReturnOrderCount int     `gorm:"column:return_order_count" json:"return_order_count"` // 退还订单数量
	FinalMoney       float64 `gorm:"column:final_money" json:"final_money"`               // 结算金额
	PushMoney        float64 `gorm:"column:push_money" json:"push_money"`                 // 提成金额
	ReturnPushMoney  float64 `gorm:"column:return_push_money" json:"return_push_money"`   // 退还提成金额
	OrderMoney       float64 `gorm:"column:order_money" json:"order_money"`               // 订单金额
	ReturnOrderMoney float64 `gorm:"column:return_order_money" json:"return_order_money"` // 退还订单金额
	Sn               string  `gorm:"column:sn" json:"sn"`                                 // 编号
}

// TableName get sql table name.获取数据库表名
func (m *EsBillTotal) TableName() string {
	return "es_bill_total"
}
