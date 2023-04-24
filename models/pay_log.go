package models

// EsPayLog 收款单(es_pay_log)
type EsPayLog struct {
	PayLogID      int     `gorm:"primaryKey;column:pay_log_id" json:"-"`         // 收款单ID
	OrderSn       string  `gorm:"column:order_sn" json:"order_sn"`               // 订单编号
	PayWay        string  `gorm:"column:pay_way" json:"pay_way"`                 // 支付方式
	PayType       string  `gorm:"column:pay_type" json:"pay_type"`               // 支付类型
	PayTime       int64   `gorm:"column:pay_time" json:"pay_time"`               // 付款时间
	PayMoney      float64 `gorm:"column:pay_money" json:"pay_money"`             // 付款金额
	PayMemberName string  `gorm:"column:pay_member_name" json:"pay_member_name"` // 付款会员名
	PayStatus     string  `gorm:"column:pay_status" json:"pay_status"`           // 付款状态
	PayLogSn      string  `gorm:"column:pay_log_sn" json:"pay_log_sn"`           // 流水号
	PayOrderNo    string  `gorm:"column:pay_order_no" json:"pay_order_no"`       // 支付方式返回流水号
}

// TableName get sql table name.获取数据库表名
func (m *EsPayLog) TableName() string {
	return "es_pay_log"
}
