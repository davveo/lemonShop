package models

// EsShop 店铺(es_shop)
type EsShop struct {
	ShopID         int    `gorm:"primaryKey;column:shop_id" json:"-"`            // 店铺Id
	MemberID       int    `gorm:"column:member_id" json:"member_id"`             // 会员Id
	MemberName     string `gorm:"column:member_name" json:"member_name"`         // 会员名称
	ShopName       string `gorm:"column:shop_name" json:"shop_name"`             // 店铺名称
	ShopDisable    string `gorm:"column:shop_disable" json:"shop_disable"`       // 店铺状态
	ShopCreatetime int64  `gorm:"column:shop_createtime" json:"shop_createtime"` // 店铺创建时间
	ShopEndtime    int64  `gorm:"column:shop_endtime" json:"shop_endtime"`       // 店铺关闭时间
}

// TableName get sql table name.获取数据库表名
func (m *EsShop) TableName() string {
	return "es_shop"
}