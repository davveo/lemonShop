package models

// EsKeywordSearchHistory [...]
type EsKeywordSearchHistory struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`
	Keyword    string `gorm:"column:keyword" json:"keyword"`         // 关键字
	Count      int    `gorm:"column:count" json:"count"`             // 搜索次数
	AddTime    int64  `gorm:"column:add_time" json:"add_time"`       // 新增时间
	ModifyTime int64  `gorm:"column:modify_time" json:"modify_time"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *EsKeywordSearchHistory) TableName() string {
	return "es_keyword_search_history"
}
