package models

// EsGroupbuyQuantityLog 团购商品库存日志表(es_groupbuy_quantity_log)
type EsGroupbuyQuantityLog struct {
	LogID    int    `gorm:"primaryKey;column:log_id" json:"-"` // 日志id
	OrderSn  string `gorm:"column:order_sn" json:"order_sn"`   // 订单编号
	GoodsID  int    `gorm:"column:goods_id" json:"goods_id"`   // 商品ID
	Quantity int    `gorm:"column:quantity" json:"quantity"`   // 团购售空数量
	OpTime   int64  `gorm:"column:op_time" json:"op_time"`     // 操作时间
	LogType  string `gorm:"column:log_type" json:"log_type"`   // 日志类型
	Reason   string `gorm:"column:reason" json:"reason"`       // 操作原因
	GbID     int    `gorm:"column:gb_id" json:"gb_id"`         // 团购商品id
}

// TableName get sql table name.获取数据库表名
func (m *EsGroupbuyQuantityLog) TableName() string {
	return "es_groupbuy_quantity_log"
}
