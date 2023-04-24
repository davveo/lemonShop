package models

// EsMemberPointHistory 会员积分表(es_member_point_history)
type EsMemberPointHistory struct {
	ID              int    `gorm:"primaryKey;column:id" json:"-"`                     // 主键ID
	MemberID        int    `gorm:"column:member_id" json:"member_id"`                 // 会员ID
	GradePoint      int64  `gorm:"column:grade_point" json:"grade_point"`             // 等级积分
	Time            int64  `gorm:"column:time" json:"time"`                           // 操作时间
	Reason          string `gorm:"column:reason" json:"reason"`                       // 操作理由
	GradePointType  int    `gorm:"column:grade_point_type" json:"grade_point_type"`   // 等级积分类型
	Operator        string `gorm:"column:operator" json:"operator"`                   // 操作者
	ConsumPoint     int64  `gorm:"column:consum_point" json:"consum_point"`           // 消费积分
	ConsumPointType int    `gorm:"column:consum_point_type" json:"consum_point_type"` // 消费积分类型
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberPointHistory) TableName() string {
	return "es_member_point_history"
}
