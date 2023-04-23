package models

// EsComplainTopic [...]
type EsComplainTopic struct {
	TopicID     int    `gorm:"primaryKey;column:topic_id" json:"-"`     //  主键
	TopicName   string `gorm:"column:topic_name" json:"topic_name"`     // 投诉主题
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`   // 添加时间
	TopicRemark string `gorm:"column:topic_remark" json:"topic_remark"` // 主题说明
}

// TableName get sql table name.获取数据库表名
func (m *EsComplainTopic) TableName() string {
	return "es_complain_topic"
}
