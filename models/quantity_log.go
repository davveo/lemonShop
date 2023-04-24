package models

// EsQuantityLog 库存日志表(es_quantity_log)
type EsQuantityLog struct {
	LogID          int    `gorm:"primaryKey;column:log_id" json:"-"`             // 日志id
	OrderSn        string `gorm:"column:order_sn" json:"order_sn"`               // 订单编号
	GoodsID        int    `gorm:"column:goods_id" json:"goods_id"`               // 商品id
	SkuID          int    `gorm:"column:sku_id" json:"sku_id"`                   // sku id
	Quantity       int    `gorm:"column:quantity" json:"quantity"`               // 库存数
	EnableQuantity int    `gorm:"column:enable_quantity" json:"enable_quantity"` // 可用库存
	OpTime         int    `gorm:"column:op_time" json:"op_time"`                 // 操作时间
	LogType        string `gorm:"column:log_type" json:"log_type"`               // 日志类型
	Reason         string `gorm:"column:reason" json:"reason"`                   // 原因
}

// TableName get sql table name.获取数据库表名
func (m *EsQuantityLog) TableName() string {
	return "es_quantity_log"
}