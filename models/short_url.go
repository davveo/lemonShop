package models

// EsShortURL 短链接(es_short_url)
type EsShortURL struct {
	ID  int64  `gorm:"primaryKey;column:id" json:"-"` // id
	URL string `gorm:"column:url" json:"url"`         // 跳转地址
	Su  string `gorm:"column:su" json:"su"`           // 短链接key
}

// TableName get sql table name.获取数据库表名
func (m *EsShortURL) TableName() string {
	return "es_short_url"
}
