package models

// EsHalfPrice 第二件半价(es_half_price)
type EsHalfPrice struct {
	HpID        int    `gorm:"primaryKey;column:hp_id" json:"-"`      // 第二件半价活动ID
	StartTime   int64  `gorm:"column:start_time" json:"start_time"`   // 活动开始时间
	EndTime     int64  `gorm:"column:end_time" json:"end_time"`       // 活动结束时间
	Title       string `gorm:"column:title" json:"title"`             // 活动标题
	RangeType   int    `gorm:"column:range_type" json:"range_type"`   // 是否全部商品参与
	Disabled    int    `gorm:"column:disabled" json:"disabled"`       // 是否停用
	Description string `gorm:"column:description" json:"description"` // 活动说明
	SellerID    int    `gorm:"column:seller_id" json:"seller_id"`     // 商家id
}

// TableName get sql table name.获取数据库表名
func (m *EsHalfPrice) TableName() string {
	return "es_half_price"
}
