package models

// EsGroupbuyActive 团购活动表(es_groupbuy_active)
type EsGroupbuyActive struct {
	ActID        int    `gorm:"primaryKey;column:act_id" json:"-"`         // 活动Id
	ActName      string `gorm:"column:act_name" json:"act_name"`           // 活动名称
	StartTime    int64  `gorm:"column:start_time" json:"start_time"`       // 活动开启时间
	EndTime      int64  `gorm:"column:end_time" json:"end_time"`           // 团购结束时间
	JoinEndTime  int64  `gorm:"column:join_end_time" json:"join_end_time"` // 团购报名截止时间
	AddTime      int64  `gorm:"column:add_time" json:"add_time"`           // 团购添加时间
	ActTagID     int    `gorm:"column:act_tag_id" json:"act_tag_id"`       // 团购活动标签Id
	GoodsNum     int    `gorm:"column:goods_num" json:"goods_num"`         // 参与团购商品数量
	DeleteReason string `gorm:"column:delete_reason" json:"delete_reason"` // 删除原因
	DeleteTime   int64  `gorm:"column:delete_time" json:"delete_time"`     // 删除时间
	DeleteName   string `gorm:"column:delete_name" json:"delete_name"`     // 删除人
	DeleteStatus string `gorm:"column:delete_status" json:"delete_status"` // 是否删除 DELETED：已删除，NORMAL：正常
}

// TableName get sql table name.获取数据库表名
func (m *EsGroupbuyActive) TableName() string {
	return "es_groupbuy_active"
}
