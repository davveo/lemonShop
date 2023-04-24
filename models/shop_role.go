package models

// EsShopRole 店铺角色(es_shop_role)
type EsShopRole struct {
	RoleID       int    `gorm:"primaryKey;column:role_id" json:"-"`        // 角色主键
	RoleName     string `gorm:"column:role_name" json:"role_name"`         // 角色名称
	AuthIDs      string `gorm:"column:auth_ids" json:"auth_ids"`           // 角色
	RoleDescribe string `gorm:"column:role_describe" json:"role_describe"` // 角色描述
	ShopID       int    `gorm:"column:shop_id" json:"shop_id"`             // 店铺id
}

// TableName get sql table name.获取数据库表名
func (m *EsShopRole) TableName() string {
	return "es_shop_role"
}