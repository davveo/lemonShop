package models

// EsSensitiveWords 敏感词(es_sensitive_words)
type EsSensitiveWords struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`         // 主键
	WordName   string `gorm:"column:word_name" json:"word_name"`     // 敏感词名称
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` // 创建时间
	IsDelete   int16  `gorm:"column:is_delete" json:"is_delete"`     // 删除状态  1正常 0 删除
}

// TableName get sql table name.获取数据库表名
func (m *EsSensitiveWords) TableName() string {
	return "es_sensitive_words"
}