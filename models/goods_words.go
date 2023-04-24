package models

// EsGoodsWords 商品分词表(es_goods_words)
type EsGoodsWords struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`     // 主键
	Words    string `gorm:"column:words" json:"words"`         // 分词
	GoodsNum int64  `gorm:"column:goods_num" json:"goods_num"` // 数量
	Quanpin  string `gorm:"column:quanpin" json:"quanpin"`     // 全拼字母
	Szm      string `gorm:"column:szm" json:"szm"`             // 首字母
	Type     string `gorm:"column:type" json:"type"`           // 类型，系统(SYSTEM)，平台(PLATFORM)
	Sort     int    `gorm:"column:sort" json:"sort"`           // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsGoodsWords) TableName() string {
	return "es_goods_words"
}
