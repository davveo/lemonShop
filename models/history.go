package models

// EsHistory [...]
type EsHistory struct {
	ID         int     `gorm:"primaryKey;column:id" json:"-"`         // 主键
	GoodsID    int     `gorm:"column:goods_id" json:"goods_id"`       // 商品id
	GoodsName  string  `gorm:"column:goods_name" json:"goods_name"`   // 商品名称
	GoodsPrice float64 `gorm:"column:goods_price" json:"goods_price"` // 商品价格
	GoodsImg   string  `gorm:"column:goods_img" json:"goods_img"`     // 商品主图
	MemberID   int     `gorm:"column:member_id" json:"member_id"`     // 会员id
	MemberName string  `gorm:"column:member_name" json:"member_name"` // 会员名称
	CreateTime int64   `gorm:"column:create_time" json:"create_time"` // 创建时间，按天存
	UpdateTime int64   `gorm:"column:update_time" json:"update_time"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *EsHistory) TableName() string {
	return "es_history"
}
