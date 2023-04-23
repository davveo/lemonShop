package models

// EsCategory 商品分类(es_category)
type EsCategory struct {
	CategoryID    int    `gorm:"primaryKey;column:category_id" json:"-"`      // 主键
	Name          string `gorm:"column:name" json:"name"`                     // 分类名称
	ParentID      int    `gorm:"column:parent_id" json:"parent_id"`           // 分类父id
	CategoryPath  string `gorm:"column:category_path" json:"category_path"`   // 分类父子路径
	GoodsCount    int    `gorm:"column:goods_count" json:"goods_count"`       // 该分类下商品数量
	CategoryOrder int    `gorm:"column:category_order" json:"category_order"` // 分类排序
	Image         string `gorm:"column:image" json:"image"`                   // 分类图标
}

// TableName get sql table name.获取数据库表名
func (m *EsCategory) TableName() string {
	return "es_category"
}
