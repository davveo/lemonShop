package models

// EsOrderLog 订单日志表(es_order_log)
type EsOrderLog struct {
	LogID   int    `gorm:"primaryKey;column:log_id" json:"-"` // 主键ID
	OrderSn string `gorm:"column:order_sn" json:"order_sn"`   // 订单编号
	OpID    int    `gorm:"column:op_id" json:"op_id"`         // 操作者id
	OpName  string `gorm:"column:op_name" json:"op_name"`     // 操作者名称
	Message string `gorm:"column:message" json:"message"`     // 日志信息
	OpTime  int64  `gorm:"column:op_time" json:"op_time"`     // 操作时间
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderLog) TableName() string {
	return "es_order_log"
}
