package models

// EsAsLog 售后服务日志记录表(es_as_log)
type EsAsLog struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`       // 主键ID
	Sn        string `gorm:"column:sn" json:"sn"`                 // 售后/退款编号
	LogTime   int64  `gorm:"column:log_time" json:"log_time"`     // 创建时间
	LogDetail string `gorm:"column:log_detail" json:"log_detail"` // 详细信息
	Operator  string `gorm:"column:operator" json:"operator"`     // 操作者
}

// TableName get sql table name.获取数据库表名
func (m *EsAsLog) TableName() string {
	return "es_as_log"
}
