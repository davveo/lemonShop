package models

// EsAdminUser 平台管理员(es_admin_user)
type EsAdminUser struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`       // 平台管理员id
	Username   string `gorm:"column:username" json:"username"`     // 管理员名称
	Password   string `gorm:"column:password" json:"password"`     // 管理员密码
	Department string `gorm:"column:department" json:"department"` // 部门
	RoleID     int    `gorm:"column:role_id" json:"role_id"`       // 权限id
	DateLine   int64  `gorm:"column:date_line" json:"date_line"`   // 创建日期
	Remark     string `gorm:"column:remark" json:"remark"`         // 备注
	UserState  int    `gorm:"column:user_state" json:"user_state"` // 是否删除
	RealName   string `gorm:"column:real_name" json:"real_name"`   // 管理员真实姓名
	Face       string `gorm:"column:face" json:"face"`             // 头像
	Founder    int    `gorm:"column:founder" json:"founder"`       // 是否为超级管理员
}

// TableName get sql table name.获取数据库表名
func (m *EsAdminUser) TableName() string {
	return "es_admin_user"
}
