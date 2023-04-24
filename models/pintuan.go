package models

// EsPintuan [...]
type EsPintuan struct {
	PromotionID    int    `gorm:"primaryKey;column:promotion_id" json:"-"`       // id
	PromotionName  string `gorm:"column:promotion_name" json:"promotion_name"`   // 活动名称
	PromotionTitle string `gorm:"column:promotion_title" json:"promotion_title"` // 活动标题
	StartTime      int64  `gorm:"column:start_time" json:"start_time"`           // 活动开始时间
	EndTime        int64  `gorm:"column:end_time" json:"end_time"`               // 活动结束时间
	RequiredNum    int    `gorm:"column:required_num" json:"required_num"`       // 成团人数
	LimitNum       int    `gorm:"column:limit_num" json:"limit_num"`             // 限购数量
	EnableMocker   int16  `gorm:"column:enable_mocker" json:"enable_mocker"`     // 虚拟成团
	CreateTime     int64  `gorm:"column:create_time" json:"create_time"`         // 创建时间
	Status         string `gorm:"column:status" json:"status"`                   // 活动状态
	OptionStatus   string `gorm:"column:option_status" json:"option_status"`     // api请求操作状态
	SellerName     string `gorm:"column:seller_name" json:"seller_name"`         // 商家名称
	SellerID       int64  `gorm:"column:seller_id" json:"seller_id"`             // 商家ID
}

// TableName get sql table name.获取数据库表名
func (m *EsPintuan) TableName() string {
	return "es_pintuan"
}
