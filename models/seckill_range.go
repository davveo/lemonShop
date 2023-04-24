package models

// EsSeckillRange 限时抢购时刻(es_seckill_range)
type EsSeckillRange struct {
	RangeID   int `gorm:"primaryKey;column:range_id" json:"-"` // 主键
	SeckillID int `gorm:"column:seckill_id" json:"seckill_id"` // 限时抢购活动id
	RangeTime int `gorm:"column:range_time" json:"range_time"` // 整点时刻
}

// TableName get sql table name.获取数据库表名
func (m *EsSeckillRange) TableName() string {
	return "es_seckill_range"
}