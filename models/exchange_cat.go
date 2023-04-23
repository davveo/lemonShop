package models

// EsExchangeCat 积分兑换分类(es_exchange_cat)
type EsExchangeCat struct {
	CategoryID    int    `gorm:"primaryKey;column:category_id" json:"-"`      // 分类id
	Name          string `gorm:"column:name" json:"name"`                     // 分类名称
	ParentID      int    `gorm:"column:parent_id" json:"parent_id"`           // 父分类
	CategoryPath  string `gorm:"column:category_path" json:"category_path"`   // 分类id路径
	GoodsCount    int    `gorm:"column:goods_count" json:"goods_count"`       // 商品数量
	CategoryOrder int    `gorm:"column:category_order" json:"category_order"` // 分类排序
	ListShow      int    `gorm:"column:list_show" json:"list_show"`           // 是否在页面上显示
	Image         string `gorm:"column:image" json:"image"`                   // 分类图片
}

// TableName get sql table name.获取数据库表名
func (m *EsExchangeCat) TableName() string {
	return "es_exchange_cat"
}
