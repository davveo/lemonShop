package models

// EsSeckill 限时抢购入库(es_seckill)
type EsSeckill struct {
	SeckillID     int    `gorm:"primaryKey;column:seckill_id" json:"-"`       // 主键id
	SeckillName   string `gorm:"column:seckill_name" json:"seckill_name"`     // 活动名称
	StartDay      int64  `gorm:"column:start_day" json:"start_day"`           // 活动日期
	ApplyEndTime  int64  `gorm:"column:apply_end_time" json:"apply_end_time"` // 报名截至时间
	SeckillRule   string `gorm:"column:seckill_rule" json:"seckill_rule"`     // 申请规则
	SellerIDs     string `gorm:"column:seller_ids" json:"seller_ids"`         // 商家id
	SeckillStatus string `gorm:"column:seckill_status" json:"seckill_status"` // 状态
	DeleteStatus  string `gorm:"column:delete_status" json:"delete_status"`   // 是否删除 DELETED：已删除，NORMAL：正常
}

// TableName get sql table name.获取数据库表名
func (m *EsSeckill) TableName() string {
	return "es_seckill"
}
