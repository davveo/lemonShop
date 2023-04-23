package models

// EsCustomWords 自定义分词表(es_custom_words)
type EsCustomWords struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`         // 主键
	Name       string `gorm:"column:name" json:"name"`               // 名称
	AddTime    int64  `gorm:"column:add_time" json:"add_time"`       // 添加时间
	Disabled   int16  `gorm:"column:disabled" json:"disabled"`       // 显示 1  隐藏 0
	ModifyTime int64  `gorm:"column:modify_time" json:"modify_time"` // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *EsCustomWords) TableName() string {
	return "es_custom_words"
}
