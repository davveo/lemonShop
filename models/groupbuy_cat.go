package models

// EsGroupbuyCat 团购分类(es_groupbuy_cat)
type EsGroupbuyCat struct {
	CatID    int    `gorm:"primaryKey;column:cat_id" json:"-"` // 分类id
	ParentID int    `gorm:"column:parent_id" json:"parent_id"` // 父类id
	CatName  string `gorm:"column:cat_name" json:"cat_name"`   // 分类名称
	CatPath  string `gorm:"column:cat_path" json:"cat_path"`   // 分类结构目录
	CatOrder int    `gorm:"column:cat_order" json:"cat_order"` // 分类排序
}

// TableName get sql table name.获取数据库表名
func (m *EsGroupbuyCat) TableName() string {
	return "es_groupbuy_cat"
}
