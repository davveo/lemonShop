package models

// EsMinus 单品立减(es_minus)
type EsMinus struct {
	MinusID              int     `gorm:"primaryKey;column:minus_id" json:"-"`                         // 单品立减活动id
	SingleReductionValue float64 `gorm:"column:single_reduction_value" json:"single_reduction_value"` // 单品立减金额
	StartTime            int64   `gorm:"column:start_time" json:"start_time"`                         // 起始时间
	StartTimeStr         string  `gorm:"column:start_time_str" json:"start_time_str"`                 // 起始时间字符串
	EndTime              int64   `gorm:"column:end_time" json:"end_time"`                             // 结束时间
	EndTimeStr           string  `gorm:"column:end_time_str" json:"end_time_str"`                     // 结束时间字符串
	Title                string  `gorm:"column:title" json:"title"`                                   // 单品立减活动标题
	RangeType            int     `gorm:"column:range_type" json:"range_type"`                         // 是否选择全部商品
	Disabled             int     `gorm:"column:disabled" json:"disabled"`                             // 是否停用
	Description          string  `gorm:"column:description" json:"description"`                       // 描述
	SellerID             int     `gorm:"column:seller_id" json:"seller_id"`                           // 商家id
}

// TableName get sql table name.获取数据库表名
func (m *EsMinus) TableName() string {
	return "es_minus"
}
