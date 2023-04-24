package models

// EsOrderOutStatus 订单出库状态(es_order_out_status)
type EsOrderOutStatus struct {
	OutID     int    `gorm:"primaryKey;column:out_id" json:"-"`   // 主键
	OrderSn   string `gorm:"column:order_sn" json:"order_sn"`     // 订单编号
	OutType   string `gorm:"column:out_type" json:"out_type"`     // 出库类型
	OutStatus string `gorm:"column:out_status" json:"out_status"` // 出库状态
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderOutStatus) TableName() string {
	return "es_order_out_status"
}
