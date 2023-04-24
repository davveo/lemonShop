package models

// EsRole 角色表(es_role)
type EsRole struct {
	RoleID       int    `gorm:"primaryKey;column:role_id" json:"-"`        // 主键ID
	RoleName     string `gorm:"column:role_name" json:"role_name"`         // 角色名称
	AuthIDs      string `gorm:"column:auth_ids" json:"auth_ids"`           // 角色介绍
	RoleDescribe string `gorm:"column:role_describe" json:"role_describe"` // 角色描述
}

// TableName get sql table name.获取数据库表名
func (m *EsRole) TableName() string {
	return "es_role"
}