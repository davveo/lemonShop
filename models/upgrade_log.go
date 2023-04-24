package models

// EsUpgradeLog 升级日志(es_upgrade_log)
type EsUpgradeLog struct {
	ID         int64  `gorm:"primaryKey;column:id" json:"-"`           // id
	MemberID   int64  `gorm:"column:member_id" json:"member_id"`       // 会员id
	MemberName string `gorm:"column:member_name" json:"member_name"`   // 会员名称
	Type       string `gorm:"column:type" json:"type"`                 // 切换类型
	OldTplID   int64  `gorm:"column:old_tpl_id" json:"old_tpl_id"`     // 旧的模板id
	NewTplID   int64  `gorm:"column:new_tpl_id" json:"new_tpl_id"`     // 新的模板id
	NewTplName string `gorm:"column:new_tpl_name" json:"new_tpl_name"` // 新的模板名称
	OldTplName string `gorm:"column:old_tpl_name" json:"old_tpl_name"` // 旧的模板名称
	CreateTime int    `gorm:"column:create_time" json:"create_time"`   // 创建日期
}

// TableName get sql table name.获取数据库表名
func (m *EsUpgradeLog) TableName() string {
	return "es_upgrade_log"
}
