package models


// EsClerk 店员(es_clerk)
type EsClerk struct {
	ClerkID    int    `gorm:"primaryKey;column:clerk_id" json:"-"`   // 店员id
	MemberID   int    `gorm:"column:member_id" json:"member_id"`     // 会员id
	ClerkName  string `gorm:"column:clerk_name" json:"clerk_name"`   // 店员名称
	Founder    int    `gorm:"column:founder" json:"founder"`         // 是否为超级管理员，1为超级管理员 0为其他管理员
	RoleID     int    `gorm:"column:role_id" json:"role_id"`         // 权限id
	UserState  int    `gorm:"column:user_state" json:"user_state"`   // 店员状态，0为禁用，1为正常
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` // 创建日期
	ShopID     int    `gorm:"column:shop_id" json:"shop_id"`         // 店铺id
}

// TableName get sql table name.获取数据库表名
func (m *EsClerk) TableName() string {
	return "es_clerk"
}
