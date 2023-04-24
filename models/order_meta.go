package models

// EsOrderMeta 订单扩展信息表(es_order_meta)
type EsOrderMeta struct {
	MetaID    int    `gorm:"primaryKey;column:meta_id" json:"-"`  // 主键ID
	OrderSn   string `gorm:"column:order_sn" json:"order_sn"`     // 订单编号
	MetaKey   string `gorm:"column:meta_key" json:"meta_key"`     // 扩展-键
	MetaValue string `gorm:"column:meta_value" json:"meta_value"` // 扩展-值
	Status    string `gorm:"column:status" json:"status"`         // 售后状态
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderMeta) TableName() string {
	return "es_order_meta"
}
